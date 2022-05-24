FROM python:3.10

WORKDIR /app

ADD requirements.txt /app/fetcher/
RUN pip install --upgrade pip && pip install --no-cache-dir -r /app/fetcher/requirements.txt

COPY . /app/fetcher/
#EXPOSE 50051