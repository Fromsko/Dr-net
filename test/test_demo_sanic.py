from sanic import response, Request, Sanic

app = Sanic("net_login")


@app.route("/")
async def index(request: Request):
    """根路由"""
    return response.json({
        "status": 200,
        "client_ip": request.ip
    })


@app.route("/user/<name>")
async def user(request: Request, name: str):
    return response.text(f"Hello, {name}!")


@app.route('/api/power')
async def test_spider(request):
    home = request.args.get('home', '515')
    return response.json({'spider': f'spider to {home}'})


if __name__ == '__main__':
    app.run(host="0.0.0.0", port=9090, access_log=False, debug=True)
