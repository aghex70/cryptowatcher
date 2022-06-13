import os

PROJECT_NAME = "fetcher"

BINANCE_WEBSOCKET_CONFIGURATION = {
    "binance": {
        "source": "binance",
        "url": "wss://stream.binance.com:9443/ws/btcusdt@trade",
        "subscription": {
            "method": "SUBSCRIBE",
            "params": [
                "btcusdt@trade",
                "ethusdt@trade",
            ],
            "id": 2,
        },
    },
}

CELERY_BROKER_URL = os.environ["CELERY_BROKER_URL"]
CELERY_RESULT_BACKEND = os.environ["CELERY_RESULT_BACKEND"]
