FROM python:3.10

ENV PYTHONUNBUFFERED 0
WORKDIR /app

ADD requirements.txt /app/interceptor-py/
RUN pip install --upgrade pip && \
    pip install --no-cache-dir -r /app/interceptor-py/requirements.txt

COPY . /app/interceptor-py/
CMD ["python3", "/app/interceptor-py/main.py", "--queue=receiver"]