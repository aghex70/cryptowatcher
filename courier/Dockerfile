FROM python:3.10

ENV PYTHONUNBUFFERED 0
WORKDIR /app

ADD requirements.txt /app/receiver-py/
RUN pip install --upgrade pip && \
    pip install --no-cache-dir -r /app/receiver-py/requirements.txt

COPY . /app/courier-py/
CMD ["python3", "/app/courier-py/main.py", "--queue=sender"]