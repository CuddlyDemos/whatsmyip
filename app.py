#!/usr/bin/env python3

from flask import Flask, request, jsonify
import os


app = Flask(__name__)
app.config["JSONIFY_PRETTYPRINT_REGULAR"] = True


@app.route("/", methods=["GET"])
def main():
    data = {}
    for k, v in request.headers.items():
        data[k] = v

    return jsonify(data)


if __name__ == '__main__':
    port = int(os.getenv('PORT', "5000"))
    app.run(debug=True, host='0.0.0.0', port=port)
