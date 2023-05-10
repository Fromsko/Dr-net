# -*- coding: utf-8 -*-
"""
    @Author: kong
    @Date  : 2023-04-25
    @GitHub: https://github.com/kongxiaoaaa
    @notes : 网络设置
"""
import os
import re
import sys
from random import choice
from ast import literal_eval
from time import sleep

from requests import api, ConnectionError
from wmi import WMI

from .net_check import check_login
from .config import header, log


class UpDateWIFI:
    """随机修改指定ip段的本机ip"""

    def __init__(self, speac=None):
        self.speac_ip = speac
        self.wmi_service = WMI()
        self.fetch_url = 'http://localhost/api/ip'
        self.header = header

        # 获取到本地有网卡信息
        self.col_nic_configs = (
            self.wmi_service.Win32_NetworkAdapterConfiguration(IPEnabled=True))

    def get_adapter(self):
        flag = 0

        # 遍历所有网卡，找到要修改的那个，这里我是用原ip的第一段正则出来的
        for obj in self.col_nic_configs:
            ip = re.findall(r"10.\d+.\d+.\d+", obj.IPAddress[0])
            if len(ip) > 0:
                return flag
            else:
                flag = flag + 1

    def run_set(self):
        """配置修改方法"""
        try:
            adapter = self.col_nic_configs[self.get_adapter()]
        except TypeError:
            log.error("请检测是否正确连接网络 ==> HNIU")
            sys.exit(0)

        # 子网掩码 和 DNS
        arr_subnet_masks = '255.255.240.0'
        arr_dns_servers = ['223.5.5.5', '114.114.114.114']

        # IP 地址
        try:
            api_resp = api.get(url=self.fetch_url,
                               headers=self.header,
                               timeout=9).text
        except ConnectionError:
            log.error("Server start failed!")
            sys.exit(1)

        arr_ip_addresses = literal_eval(api_resp)['content']['scan_ip']

        if self.speac_ip is not None:
            arr_ip_addresses = self.speac_ip

        # 网关
        arr_default_gateways = [
            f"10.{choice(['14', '15', '26', '10', '4'])}.128.1"
        ]
        arr_gateway_cost_metrics = [1]  # 这里要设置成1，代表非自动选择
        ip_res = adapter.EnableStatic(IPAddress=[arr_ip_addresses],
                                      SubnetMask=[arr_subnet_masks])
        if ip_res[0] == 0:
            log.info("IP地址设置成功!")
            log.info(f"IP => {arr_ip_addresses}")
        else:
            if ip_res[0] == 1:
                log.info('设置IP成功,需要重启计算机!')
            else:
                log.error('修改IP失败: IP设置发生错误')
            return False

        way_res = adapter.SetGateways(
            DefaultIPGateway=arr_default_gateways,
            GatewayCostMetric=arr_gateway_cost_metrics)
        if way_res[0] == 0:
            log.info('设置网关成功')
            log.info(f"Gateway => {arr_default_gateways[0]}")
        else:
            log.error('tip:修改网关失败: 网关设置发生错误')
            return False

        dns_res = adapter.SetDNSServerSearchOrder(
            DNSServerSearchOrder=arr_dns_servers)

        if dns_res[0] == 0:
            log.info('设置DNS成功,等待3秒刷新缓存')
            sleep(3)
            # 刷新DNS缓存使DNS生效
            os.system('ipconfig /flushdns')
        else:
            log.error('tip:修改DNS失败: DNS设置发生错误')
            return False

    @staticmethod
    def _ping():
        """ping某ip看是否可以通"""
        sleep(3)
        dialog = "请求 内网服务器 http://10.253.0.1/ {} !"
        res = os.popen('ping -n 2 -w 1 http://10.253.0.1/').read()
        if '请求超时' in res:  # 注意乱码编码问题
            log.error(dialog.format("失败"))
        else:
            log.info(dialog.format("成功"))
            net_info = check_login()
            return net_info
        return False

    def login_info(self):
        """检测登陆信息"""
        try:
            if net_info := self._ping():
                if (info := net_info['data']) != '' or net_info['code'] == 200:
                    resp_ip = info['v4ip']
                    user = info.get('uid', '')
                    return resp_ip, user
        except KeyError as err:
            log.error(f"登录失败 - {err} - 正在尝试更改配置!")
            return False
        return False
