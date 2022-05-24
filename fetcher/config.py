websockets_configuration = {
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
