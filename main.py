# -*- coding: utf-8 -*-
"""
    @Author: kong
    @File  : net_set_config.py
    @Date  : 2023-04-20 03:00:35
    @GitHub: https://github.com/kongxiaoaaa
    @notes : 内网认证客户端
"""
from net.net_check import check_start_config
from net.net_set import UpDateWIFI
from net.config import log
import sys


FLAG = False
check_start_config()
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
