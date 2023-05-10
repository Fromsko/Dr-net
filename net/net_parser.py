# -*- coding: utf-8 -*-
"""
    @Author: kong
    @File  : net_parser.py
    @GitHub: https://github.com/kongxiaoaaa
    @notes : 解析IP数据
"""
import json

from config import res_filename, resource_dir


class Parser:
    """ 扫描解析 """

    @staticmethod
    def parser_info(result: str):
        """解析单次扫描结果"""
        data_dict = dict()
        result_info = result.split(' ')

        data_dict.update({
            "scan_ip": result_info[-3],
            # "mac": info[2],  # 防止信息泄露 | 将隐藏 MAC 地址
            "ttl": result_info[-1],
        })
        return data_dict

    @staticmethod
    def write_info(content: dict, file_name: str = res_filename):
        """将数据写入 json文件中"""
        if not isinstance(content, dict):
            raise TypeError("存储格式的类型不匹配,应提供 dict 格式的数据.")

        if not file_name.endswith(".json"):
            file_name = file_name + ".json"

        info_name = str(resource_dir / file_name)

        with open(info_name, "w", encoding="utf-8") as file_obj:
            file_obj.write(
                json.dumps(content, ensure_ascii=False)
            )
        return True
