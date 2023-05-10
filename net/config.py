# -*- coding: utf-8 -*-
"""
    @Author: kong
    @File  : config.py
    @Date  : 2023-04-20 02:54:55
    @GitHub: https://github.com/kongxiaoaaa
    @notes : 配置文件
"""
import json
import sys
from datetime import datetime
from pathlib import Path
from typing import Any

from loguru import logger
from user_agent import generate_user_agent

# 根路径
base_dir = Path(__file__).parents[1]
# 网络文件
net_dir: Path = base_dir / 'net'
# 日志路径
log_dir: Path = base_dir.parent / 'log'
# 资源路径
resource_dir: Path = base_dir / 'res'
# 资源文件
res_filename: str = "scan-all.json"
# 请求头
header: dict = {"User-Agent": generate_user_agent()}


class Config:
    """配置类"""
    scan_flag: bool = False

    @property
    def load_scan_data(self):
        """导入扫描信息"""
        try:
            with open(
                    str(resource_dir / res_filename),
                    mode="r",
                    encoding="utf-8"
            ) as f:
                content: dict = json.loads(f.read())
        except FileNotFoundError:
            err_info = f"-[{res_filename}]- File is not Found!"
            raise FileNotFoundError(err_info)
        return content

    def check_file_timeout(self):
        """文件修改时间检测 or 文件存在检测"""
        dir_create_time = datetime.fromtimestamp(resource_dir.stat().st_mtime)
        delta_time = datetime.now() - dir_create_time

        # 判断时间差是否超过一天
        if delta_time.days >= 1:
            self.scan_flag = True
            log.error("文件创建时间超过了 1 天")
        else:
            if (resource_dir / res_filename).exists():
                self.scan_flag = False
                log.debug("文件创建没有超过一天")
            else:
                self.scan_flag = True
                log.error("资源文件不存在")


class Log:
    def __init__(self, log_name="net_log.log", bind_name="Login_info"):
        if not log_name.endswith(".log"):
            log_name: str = log_name + ".log"

        # 移除标准控制台
        logger.remove()

        # 添加控制台输出
        logger.add(
            sys.stdout,
            colorize=True,
            format="<green><b>{time:YYYY-MM-DD HH:mm:ss}</b> </green> "
                   "| <blue>{level}</blue> | {file} - {message}"
        )

        logger.add(
            str(log_dir / log_name),
            level='DEBUG',
            format='{time:YYYY-MM-DD HH:mm:ss} - |{level}| {file} - {message}',
            rotation="10 MB"
        )
        logger.bind(name=bind_name)

    @staticmethod
    def info(message: Any) -> None:
        logger.info(message)

    @staticmethod
    def debug(message: Any) -> None:
        logger.debug(message)

    @staticmethod
    def error(message: Any) -> None:
        logger.error(message)


log = Log()
