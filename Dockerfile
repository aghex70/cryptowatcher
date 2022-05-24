FROM python:3.10

RUN apt-get -y update &&\
    apt-get install -y libldap2-dev libssl-dev libsasl2-dev  &&\
    apt-get clean

WORKDIR /app

ADD requirements.txt /app/fetcher/
RUN pip install --upgrade pip && pip install --no-cache-dir -r /app/fetcher/requirements.txt

COPY . /app/fetcher/