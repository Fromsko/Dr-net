# -*- coding: utf-8 -*-
"""
    @Author: skong
    @Date  : 2023-04-25
    @GitHub: https://github.com/kongxiaoaaa
    @notes : 测试文件
"""
import sys

from snet_work.net_set import UpDateWIFI
from snet_work.tool.config import log

FLAG = False
update = UpDateWIFI()

if len(args := sys.argv) > 1:
    if "ip" == args[1]:
        update.run_set()

# 重载配置
while not FLAG:
    if login_info := update.login_info():
        FLAG = True
        log.info(f"登录用户: {login_info[0]}")
        log.info(f"登录IP地址: {login_info[1]}")
    else:
        update.run_set()
