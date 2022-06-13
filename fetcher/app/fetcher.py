from typing import List, Union

from app import tasks
from config import settings
from config.logger import logger
from grpc_server import fetcher_pb2, fetcher_pb2_grpc


class FetcherService(fetcher_pb2_grpc.FetcherServiceServicer):
    def FetchTrades(
        self, request: fetcher_pb2.FetchTradesRequest, context
    ) -> fetcher_pb2.FetchTradesResponse:
        logger.info("FetchTrades called. Request: %s", request)
        result = self._handle_fetch_request(request)
        return fetcher_pb2.FetchTradesResponse(
            tasks=[result],
        )

    def StopFetchTrades(
        self, request: fetcher_pb2.StopFetchTradesRequest, context
    ) -> fetcher_pb2.StopFetchTradesResponse:
        logger.info("StopFetchTrades called. Request: %s", request)
        result = self._handle_stop_fetch_request()
        return fetcher_pb2.StopFetchTradesResponse(
            success=result,
        )

    @staticmethod
    def _handle_fetch_request(
        request: fetcher_pb2.FetchTradesRequest,
    ) -> Union[List[int], None]:
        """Handle a FetchTrades request.
        :param request: The request to handle.
        :return: A list of task_ids.
        """
        task_ids = []
        for req in request.sources:
            source = req.source
            configuration = settings.BINANCE_WEBSOCKET_CONFIGURATION.get(source)
            if not configuration:
                logger.error("Websocket configuration not found. Source: %s", source)
                return None

            source, url, subscription = (
                configuration.get(k) for k in ("source", "url", "subscription")
            )

            logger.info(
                "Calling retrieve orders task. Source: %s, url: %s", source, url
            )
            task_id = tasks.retrieve_orders.apply_async((url, source, subscription))
            logger.info(
                "Retrieve order task created. Task_id: %s, source: %s, url: %s",
                task_id,
                source,
                url,
            )
            task_ids.append(task_id)
        return task_ids

    @staticmethod
    def _handle_stop_fetch_request() -> bool:
        """Handle a StopFetchTrades request.
        :return True.
        """
        logger.info("Stopping fetch trades.")
        tasks.stop_orders_fetch.apply_async()
        return True
