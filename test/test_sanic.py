from sanic import response
from sanic import Sanic
from sanic.response import json

import logging
from logging.handlers import RotatingFileHandler

app = Sanic("my-hello-world-app")

# 配置访问日志记录器
access_logger = logging.getLogger("access")
access_logger.setLevel(logging.INFO)
handler = RotatingFileHandler("access.log",
                              maxBytes=10 * 1024 * 1024,
                              backupCount=10)
handler.setFormatter(
    logging.Formatter("%(asctime)s [%(levelname)s] %(message)s"))
access_logger.addHandler(handler)


# 注册请求开始的回调函数
@app.listener('before_server_start')
async def setup_access_log(app, loop):

    def log_request(request):
        access_logger.info(f"{request.ip} - {request.method} {request.url}")

    app.request_middleware.append(log_request)


@app.route("/")
async def index(request):
    return response.text("Hello, Sanic!")


@app.route("/user/<name>")
async def user(request, name):
    return response.text(f"Hello, {name}!")


@app.route('/api/power')
async def test_spider(request):
    home = request.args.get('home', '515')
    return json({'spider': f'spider to {home}'})


if __name__ == '__main__':
    app.run(host="0.0.0.0", port=9090, access_log=False, debug=True)
