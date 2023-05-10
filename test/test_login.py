import requests
from user_agent import generate_user_agent

header = {
    "User-Agent": generate_user_agent()
}


def find_mac(username: str, user_ip: str):
    data = {
        "user_account": username,
        "login_method": 0,
        "find_mac": 0,
        "wlan_user_ip": user_ip,
        "wlan_user_mac": "000000000000",
        "jsVersion": "4.2"
    }
    resp = requests.post(
        url="http://10.253.0.1:801/eportal/portal/mac/find",
        headers=header,
        data=data,
    )
    if resp.ok:
        print(resp.text)
    else:
        print("Error: 访问失败")


def short_login():
    data = {
        "program_index": "ezEpAA1657160826",
        "page_index": "hJoSZN1657160869",
        "login_method": 0,
        "wlan_user_ip": "10.27.134.81",
        "wlan_user_ipv6": "",
        "wlan_user_mac": "000000000000",
        "wlan_ac_ip": "",
        "wlan_ac_name": "",
        "gw_id": "000000000000",
        "gw_port": "",
        "gw_address": "",
        "login_type": 1
    }
    short_login = "http://10.253.0.1:801/eportal/portal/login/short_login"
    resp = requests.post(
        url=short_login,
        headers=header,
        data=data,
    )
    if resp.ok:
        print(resp.text)
    else:
        print("Error: 访问失败")


def logout(
    username,
    parser_ip: str = "169576017",
    _host: str = "10.253.0.1:801"
):
    """模拟退出登录"""
    def _(ip: str):
        if not ip.find("."):
            return ip
        ip = ip.split('.')
        ip = (int(ip[0]) << 24 |
              int(ip[1]) << 16 |
              int(ip[2]) << 8 |
              int(ip[3])) & 0xFFFFFFFF
        return ip

    logout_url = f"http://{_host}/eportal/portal/mac/unbind?callback=dr1002&"\
        f"user_account={username}%40cmcc&wlan_user_mac=000000000000"\
        f"&wlan_user_ip={_(parser_ip)}&jsVersion=4.2&v=3056&lang=zh"
    resp = requests.get(
        url=logout_url,
        headers=header,
    )
    if resp.ok:
        print(resp.text)
        print("Success logout")
    else:
        print("Error: 访问失败")


if __name__ == "__main__":
    # find_mac(
    #     "202216470112@cmcc",
    #     "10.27.134.81"
    # )
    logout(
        "202216470112@cmcc",
        "10.27.134.81"
    )

# ding@9000
# 123456
