#!/usr/bin/env python3

from flask import Flask
from flask import request


app = Flask(__name__)


@app.route("/", methods=["GET"])
def get_my_ip():
    if request.headers.get("X-Appengine-User-Ip"):
        ip = request.headers.get("X-Appengine-User-Ip")
    else:
        ip = request.remote_addr
    return ip, 200


if __name__ == '__main__':
    app.run(host='0.0.0.0')
