# -*- coding: utf-8 -*-
"""
    @Author: kong
    @File  : main.py
    @Date  : 2023-04-13 04:55:29
    @GitHub: https://github.com/kongxiaoaaa
    @notes : 基于IP地址和MAC地址 模拟认证[服务端]
"""
import sys
from random import choice

from fastapi import FastAPI, Request
from fastapi.responses import RedirectResponse
from uvicorn import run

from config import Config, log

try:
    result = Config().load_scan_data()
except FileNotFoundError:
    log.error("File not found!")
    sys.exit(1)
else:
    data_info = list(result.values())

app = FastAPI(docs_url=None, redoc_url=None, openapi_url=None)


@app.exception_handler(404)
async def redirect(request: Request, exc):
    return RedirectResponse('/api/rand_ip')


@app.api_route('/api/ip', methods=['GET', 'POST'])
async def choice_one_ip(request: Request):
    origin_ip: str = request.client.host
    one_ip = choice(choice(data_info))
    log.info(f"获取到的地址是: {one_ip}")
    return {'content': one_ip, 'origin_ip': origin_ip}


@app.get('/api/area_ip')
async def choice_area_ip():
    """ 随机选择ip """
    rand_ip: dict = choice(data_info)
    return rand_ip


@app.get('/api/all_ip')
async def all_ip_dict():
    """ 返回全部ip信息 """
    return data_info


if __name__ == "__main__":
    run("net_server:app", host="0.0.0.0", port=80, reload=True)
