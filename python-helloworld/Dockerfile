# Multi-Stage build
# Build stage
FROM python:3.9.7-slim-buster AS builder

WORKDIR /tmp

RUN apt-get update && apt-get install -y gcc

COPY requirements.txt .

RUN pip3 install -r requirements.txt

# Multi-Stage build
# Running stage
FROM python:3.9.7-slim-buster

WORKDIR /opt/helloworld

RUN apt-get update && apt-get install -y wget

COPY --from=builder /usr/local/lib/python3.9/site-packages /usr/local/lib/python3.9/site-packages
COPY --from=builder /usr/local/bin/uwsgi /usr/local/bin/uwsgi

COPY src/ .

ENTRYPOINT ["uwsgi","uwsgi.ini"]