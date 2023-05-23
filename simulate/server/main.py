# -*- encoding: utf-8 -*-
"""
    Author  : kong
    GitHub  : https://github.com/kongxiaoaaa
    Notes   : 程序入口
"""
import os
import sys
from pathlib import Path

from code import ScanIP, log, Config

if __name__ == '__main__':
    config = Config()
    config.load_scan_data()
    file_load = Path(__file__).parent / 'code' / 'scan_start.py'

    scan_ips = [f"10.27.{i}.1/24" for i in range(130, 200)]
    scanner = ScanIP(scan_ips)

    if scanner.run():
        os.system(f"{sys.executable} {file_load}")
    else:
        log.error("没有获取到数据, 所以无法启动服务!")
