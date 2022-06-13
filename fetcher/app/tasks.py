from __future__ import absolute_import, unicode_literals

import json
import os
from typing import Dict, Optional

import celery
import motor.motor_asyncio
import websockets
from celery import Celery
from config import settings
from config.logger import logger

app = Celery(
    broker=settings.CELERY_BROKER_URL,
    backend=settings.CELERY_RESULT_BACKEND,
    include=("fetcher.tasks",),
)


class HandlerTask(celery.Task):
    def run(self, *args, **kwargs):
        raise NotImplementedError

    def on_failure(self, exc, task_id, args, kwargs, einfo):
        if self.request.retries != self.max_retries:
            countdown = self.default_retry_delay ** (self.request.retries + 1)
            logger.error(exc)
            logger.error(
                "Error occurred in task. Proceeding to retry in %s seconds. Task: %s",
                countdown,
                self.name,
            )
            self.retry(countdown)
        else:
            logger.error(exc)
            logger.error(
                "Error occurred in task. Max retries reached. Task: %s", self.name
            )


@app.task(base=HandlerTask, bind=True, name="stop_orders_fetch")
def stop_orders_fetch(self) -> None:
    """Stop fetching trades from websocket."""
    logger.info("Executing task. Task: %s", self.name)
    task_id = self.request.id
    inspect = celery.app.control.Inspect(app=app)
    active_tasks = inspect.active()
    logger.info("Retrieved active tasks. Tasks: %s", active_tasks)
    task_ids = [
        task.get("id")
        for tasks in active_tasks.values()
        for task in tasks
        if task.get("id") != task_id
    ]

    logger.info("Retrieved tasks to revoke. Task_ids: %s", task_ids)
    control = celery.app.control.Control(app=app)
    for task_id in task_ids:
        logger.info("Revoking task. Task_id: %s", task_id)
        control.revoke(task_id, terminate=True)
        logger.info("Task revoked. Task_id: %s", task_id)


@app.task(base=HandlerTask, bind=True, name="retrieve_orders")
def retrieve_orders(
    self, url: str, source: str, subscription: Optional[Dict] = None
) -> None:
    """Connect to the websocket and retrieve orders
    :param url: URL of websocket.
    :param source: Name of the exchange.
    :param subscription: Subscription event to be used if necessary.
    """
    import asyncio

    logger.info("Executing task. Task: %s", self.name)

    try:
        logger.info(f"Opening websocket connection. Source: %s, url: %s", source, url)
        loop = asyncio.get_event_loop()
    except RuntimeError:
        loop = asyncio.new_event_loop()
        asyncio.set_event_loop(loop)

    logger.info("Creating driver. URI: %s", os.environ["MONGO_URI"])
    client = motor.motor_asyncio.AsyncIOMotorClient(
        os.environ["MONGO_URI"], io_loop=loop
    )

    loop.run_until_complete(fetch_and_insert(client, url, source, subscription))


async def fetch_and_insert(
    client: motor.motor_asyncio.AsyncIOMotorClient,
    url: str,
    source: str,
    subscription: Optional[Dict] = None,
) -> None:
    """Fetch events from websocket and insert them into database.
    :param client: Asynchronous MongoDB client.
    :param url: URL of websocket.
    :param source: Name of the exchange.
    :param subscription: Subscription event to be used if necessary.
    """
    db = client.events
    db_collection = db.trades

    logger.info("Connecting to websocket. URL: %s", url)
    async with websockets.connect(url) as ws:
        if subscription:
            logger.info("Subscribing to websocket. Subscription: %s", subscription)
            await ws.send(json.dumps(subscription))

        while True:
            message = await ws.recv()
            logger.info("Received message. Message: %s", message)
            await db_collection.insert_one(json.loads(message) | {"source": source})
