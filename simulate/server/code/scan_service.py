# -*- coding: utf-8 -*-
"""
    @Author: kong
    @GitHub: https://github.com/kongxiaoaaa
    @notes : arp存活主机扫描
"""
import json
import subprocess
from multiprocessing import Process, Queue
from pathlib import Path
from typing import List

from tqdm import tqdm

from .scan_config import resource_dir, log, res_filename


class Parser:
    """ 扫描解析 """

    @staticmethod
    def parser_info(result: str):
        """解析单次扫描结果"""
        data_dict = dict()
        result_info = result.split(' ')

        data_dict.update({
            "scan_ip": result_info[-3],
            "mac": result_info[2],  # 防止信息泄露 | 将隐藏 MAC 地址
            "ttl": result_info[-1],
        })
        return data_dict

    @staticmethod
    def write_info(content: dict, file_name: str = res_filename):
        """将数据写入 json文件中"""
        info_name = str(resource_dir / file_name)
        with open(info_name, "w", encoding="utf-8") as file_obj:
            file_obj.write(json.dumps(content, ensure_ascii=False))
            log.info("存储成功!")


class ScanIP(Parser):

    def __init__(self, ip_list: List):
        super().__init__()
        self._RESULT: dict = {}
        self.queue = Queue()
        self.scan_ips: List = ip_list

    def run_demo_exe(self, scan_ip: str) -> None:
        """调用subprocess启动程序

        Args:
            - `scan_ip`: 扫描地址段
        Return:
            None
        """
        file_load: Path = resource_dir / "arp-scan.exe"
        popen_obj = subprocess.Popen(str(file_load) + " -t " + scan_ip,
                                     stdout=subprocess.PIPE,
                                     stderr=subprocess.STDOUT,
                                     encoding="utf-8")
        try:
            count, info = 0, {}
            popen_obj.wait(timeout=10)

            if results := popen_obj.communicate()[0]:
                key = scan_ip.split('.')[-2]
                info.update({key: []})

                for result in results.split('\n'):
                    if result != '' and "Reply" in result:
                        info[key].append(self.parser_info(result))
                    count += 1
                self.queue.put(info)
        except Exception as err:
            log.error(f"process timeout {err}")
            popen_obj.kill()

    def run(self):
        """多进程运行"""
        process_list = []
        bar = tqdm(self.scan_ips)

        for scan in bar:
            p = Process(target=self.run_demo_exe, args=(scan, ))
            bar.set_description_str(f"启动任务 [{scan}]")
            p.start()
            process_list.append(p)

        log.info("请等待所有任务结束!")
        for i in process_list:
            i.join()

        # 从队列中读取结果
        while not self.queue.empty():
            result = self.queue.get()
            self._RESULT.update(result)
        log.info("扫描结束!")

        if not self._RESULT:
            log.error("似乎没有扫描到哦~")
            return False
        else:
            self.write_info(self._RESULT)
            return True
