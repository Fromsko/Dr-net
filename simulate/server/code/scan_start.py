# -*- coding: utf-8 -*-
"""
    @Author: kong
    @File  : main.py
    @Date  : 2023-04-13 04:55:29
    @GitHub: https://github.com/kongxiaoaaa
    @notes : 基于IP地址和MAC地址 模拟认证[服务端]
"""
from datetime import datetime
from random import choice

from fastapi import FastAPI, Request
from fastapi.responses import RedirectResponse
from starlette.exceptions import HTTPException as StarletteHTTPException
from uvicorn import run

from scan_config import log, Config

app = FastAPI(docs_url=None, redoc_url=None, openapi_url=None)
result = {}


def _register_app_startup_event(application: FastAPI) -> callable:
    """注册启动插件"""

    def start_up():
        global result
        result = Config().load_scan_data()

    return start_up


class UberSuperHandler(StarletteHTTPException):
    pass


async def function_for_uber_super_handler(request, exc):
    return RedirectResponse("/")


app.add_event_handler("startup", _register_app_startup_event(app))
app.add_exception_handler(UberSuperHandler, function_for_uber_super_handler)


@app.get("/api/v1/ip")
async def choice_area_ip():
    """ 随机选择ip """
    choice_num_ip = choice(choice(result))
    return {
        "content": choice_num_ip,
        "time": str(datetime.now().strftime("%Y-%m-%d %H:%M:%S"))
    }


@app.get("/api/v1/all_ip")
async def all_ip_dict():
    """ 返回全部ip信息 """
    return {
        "content": choice(result),
        "time": str(datetime.now().strftime("%Y-%m-%d %H:%M:%S"))
    }


@app.exception_handler(StarletteHTTPException)
async def custom_http_exception_handler(request, exc):
    return RedirectResponse("/api/v1/ip")


@app.get("/api/v1/test")
async def choice_one_ip(request: Request):
    origin_ip: str = request.client.host
    one_ip = choice(choice(result))
    log.info(f"获取到的地址是: {one_ip}")
    return {'content': one_ip, 'origin_ip': origin_ip}


if __name__ == "__main__":
    run("scan_start:app", host="0.0.0.0", port=9000, reload=True)
