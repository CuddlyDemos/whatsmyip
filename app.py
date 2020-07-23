#!/usr/bin/env python3

from flask import Flask, request, jsonify
import os


app = Flask(__name__)


@app.route("/", methods=["GET"])
def main():
    data = {
        'remote_addr': request.remote_addr
    }
    for k, v in request.headers.items():
        data[k] = v
    return jsonify(data), 200


if __name__ == '__main__':
    port = int(os.getenv('PORT', "5000"))
    app.run(debug=True, host='0.0.0.0', port=port)
