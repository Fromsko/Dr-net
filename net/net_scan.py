# -*- coding: utf-8 -*-
"""
    @Author: kong
    @Date  : 2023-04-19 17:05:22
    @GitHub: https://github.com/kongxiaoaaa
    @notes : arp存活主机扫描
"""

import subprocess
import sys
from multiprocessing import Process, Queue
from pathlib import Path
from typing import List

from config import resource_dir, log
from net_parser import Parser


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
                        info[key].append(
                            self.parser_info(result)
                        )
                    count += 1
                self.queue.put(info)
        except Exception as err:
            log.error("process timeout" + str(err))
            popen_obj.kill()

    def run(self):
        """多进程运行"""
        process_list = []

        for scan in self.scan_ips:
            p = Process(target=self.run_demo_exe, args=(scan,))
            log.info(f"start => {scan}")
            p.start()
            process_list.append(p)

        for i in process_list:
            i.join()

        # 从队列中读取结果
        while not self.queue.empty():
            result = self.queue.get()
            self._RESULT.update(result)

        log.info("Scan End!")
        return self._RESULT


if __name__ == "__main__":
    scan_ips = [f"10.27.{i}.1/24" for i in range(130, 200)]

    if len(args := sys.argv) > 1:
        try:
            arg = int(args[1])
            scan_ips = [f"10.{arg}.{i}.1/24" for i in range(130, 200)]
        except ValueError:
            log.error("You Input Error!")
            sys.exit(0)

    scanner = ScanIP(scan_ips)
    scan_data = scanner.run()
    if scanner.write_info(scan_data):
        log.info("Finish Saved => scan-all.json")
