from typing import List, Union

from app import tasks
from config import settings
from config.logger import logger

# from grpc_server import trades_pb2, trades_pb2_grpc
from grpc_server import trades_pb2, trades_pb2_grpc


class FetcherService(trades_pb2_grpc.FetcherServiceServicer):
    def FetchTrades(
        self, request: trades_pb2.FetchTradesRequest, context
    ) -> trades_pb2.FetchTradesResponse:
        logger.info("FetchTrades called. Request: %s", request)
        result = self._handle_fetch_request(request)
        return trades_pb2.FetchTradesResponse(
            tasks=[trades_pb2.Task(id=result)],
        )

    def StopFetchTrades(
        self, request: trades_pb2.Empty, context
    ) -> trades_pb2.StopFetchTradesResponse:
        logger.info("StopFetchTrades called. Request: %s", request)
        result = self._handle_stop_fetch_request()
        return trades_pb2.StopFetchTradesResponse(
            success=result,
        )

    @staticmethod
    def _handle_fetch_request(
        request: trades_pb2.FetchTradesRequest,
    ) -> Union[str, None]:
        """Handle a FetchTrades request.
        :param request: The request to handle.
        :return: Task_id.
        """
        source = request.source
        configuration = settings.BINANCE_WEBSOCKET_CONFIGURATION.get(source)
        if not configuration:
            logger.error("Websocket configuration not found. Source: %s", source)
            return None

        source, url, subscription = (
            configuration.get(k) for k in ("source", "url", "subscription")
        )

        logger.info("Calling retrieve orders task. Source: %s, url: %s", source, url)
        task_id = tasks.retrieve_orders.apply_async((url, source, subscription)).id
        logger.info(
            "Retrieve order task created. Task_id: %s, source: %s, url: %s",
            task_id,
            source,
            url,
        )
        return task_id

    @staticmethod
    def _handle_stop_fetch_request() -> bool:
        """Handle a StopFetchTrades request.
        :return True.
        """
        logger.info("Stopping fetch trades.")
        tasks.stop_orders_fetch.apply_async()
        return True
