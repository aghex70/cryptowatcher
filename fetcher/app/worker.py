from celery import Celery

try:
    from config import settings
except ModuleNotFoundError:
    from fetcher.config import settings

app = Celery(
    broker=settings.CELERY_BROKER_URL,
    backend=settings.CELERY_RESULT_BACKEND,
    include=("fetcher.app.tasks",),
)
