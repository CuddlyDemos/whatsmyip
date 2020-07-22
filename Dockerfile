FROM python:alpine

ADD requirements.txt /requirements.txt

RUN pip3 install -r /requirements.txt && \
    rm -rf /root/.cache

ADD app.py /app.py

ENV PORT=5000

EXPOSE 5000

USER nobody

CMD exec gunicorn --bind :$PORT --workers 1 --threads 8 --timeout 0 app:app
