from celery import Celery
from config import settings

app = Celery(
    broker=settings.CELERY_BROKER_URL,
    backend=settings.CELERY_RESULT_BACKEND,
    include=("fetcher.tasks",),
)
