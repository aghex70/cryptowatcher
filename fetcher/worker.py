import os

from celery import Celery

app = Celery(
    broker=os.environ["CELERY_BROKER_URL"],
    backend=os.environ["CELERY_RESULT_BACKEND"],
    include=("fetcher.tasks",),
)
