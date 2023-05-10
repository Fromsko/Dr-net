# -*- coding: utf-8 -*-
"""
    @Author: kong
    @File  : net_check.py
    @GitHub: https://github.com/kongxiaoaaa
    @notes : 检测登陆状况
"""
import os
import sys
from ast import literal_eval
from datetime import datetime
from re import findall
from time import sleep

from requests import api, Response, HTTPError, ReadTimeout

from .config import header, log, Config, net_dir


def check_login() -> dict:
    """寻找登陆状态
        - `_host` 基础地址[str]
        ~~~
        - `result` 结果[dict]
    """
    fetch_url = "http://10.253.0.1/drcom/chkstatus?" \
                "callback=dr1002&jsVersion=4.X&v=4197&lang=zh"
    result = {"code": 200, "data": None, "time": str(datetime.now())}

    try:
        resp: Response = api.get(fetch_url, headers=header, timeout=9)
        resp.raise_for_status()

        if (res := findall(pattern=r"dr1002\((.*?)\)",
                           string=resp.text)) is not None:
            assert isinstance(res, list)

            result['data'] = literal_eval(res[0]) or ''
            log.info(
                f"当前信息:=> {result['data'].get('NID', '')} "
                f"{result['data'].get('v46ip', '')}"
            )

    except (TimeoutError, HTTPError, AssertionError, ReadTimeout) as err:
        result.update(**{'data': '', 'code': 404, 'error': str(err.__class__)})
        log.error(err.__class__)
        return result
    return result


def check_start_config(new_ip: str = None):
    """启动时: 文件清除

        判断 [scan-all.json] 文件时间 <= 1
            - yes: 重新生成文件
            - no: 执行服务器启动
    """
    config = Config()
    config.check_file_timeout()
    server_file = net_dir / "net_server.py"

    if config.scan_flag:  # 文件超时
        # 重新生成 scan-all.json
        scan_file = net_dir / "net_scan.py"
        os.system(f"{sys.executable} {scan_file} {new_ip or ''}")

    os.system(f"start {sys.executable} {server_file}")
    log.info("请等待几秒服务器需要启动")
    sleep(3)


if __name__ == "__main__":
    check_start_config()
