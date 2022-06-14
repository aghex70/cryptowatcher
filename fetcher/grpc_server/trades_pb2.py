# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: trades.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database

# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(
    b'\n\x0ctrades.proto\x12\nfetcher.v1"$\n\x12\x46\x65tchTradesRequest\x12\x0e\n\x06source\x18\x01 \x01(\t"\x12\n\x04Task\x12\n\n\x02id\x18\x01 \x01(\t"6\n\x13\x46\x65tchTradesResponse\x12\x1f\n\x05tasks\x18\x01 \x03(\x0b\x32\x10.fetcher.v1.Task"*\n\x17StopFetchTradesResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08"\x07\n\x05\x45mpty2\xaf\x01\n\x0e\x46\x65tcherService\x12P\n\x0b\x46\x65tchTrades\x12\x1e.fetcher.v1.FetchTradesRequest\x1a\x1f.fetcher.v1.FetchTradesResponse"\x00\x12K\n\x0fStopFetchTrades\x12\x11.fetcher.v1.Empty\x1a#.fetcher.v1.StopFetchTradesResponse"\x00\x42\x08Z\x06\x63w-apib\x06proto3'
)


_FETCHTRADESREQUEST = DESCRIPTOR.message_types_by_name["FetchTradesRequest"]
_TASK = DESCRIPTOR.message_types_by_name["Task"]
_FETCHTRADESRESPONSE = DESCRIPTOR.message_types_by_name["FetchTradesResponse"]
_STOPFETCHTRADESRESPONSE = DESCRIPTOR.message_types_by_name["StopFetchTradesResponse"]
_EMPTY = DESCRIPTOR.message_types_by_name["Empty"]
FetchTradesRequest = _reflection.GeneratedProtocolMessageType(
    "FetchTradesRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _FETCHTRADESREQUEST,
        "__module__": "trades_pb2"
        # @@protoc_insertion_point(class_scope:fetcher.v1.FetchTradesRequest)
    },
)
_sym_db.RegisterMessage(FetchTradesRequest)

Task = _reflection.GeneratedProtocolMessageType(
    "Task",
    (_message.Message,),
    {
        "DESCRIPTOR": _TASK,
        "__module__": "trades_pb2"
        # @@protoc_insertion_point(class_scope:fetcher.v1.Task)
    },
)
_sym_db.RegisterMessage(Task)

FetchTradesResponse = _reflection.GeneratedProtocolMessageType(
    "FetchTradesResponse",
    (_message.Message,),
    {
        "DESCRIPTOR": _FETCHTRADESRESPONSE,
        "__module__": "trades_pb2"
        # @@protoc_insertion_point(class_scope:fetcher.v1.FetchTradesResponse)
    },
)
_sym_db.RegisterMessage(FetchTradesResponse)

StopFetchTradesResponse = _reflection.GeneratedProtocolMessageType(
    "StopFetchTradesResponse",
    (_message.Message,),
    {
        "DESCRIPTOR": _STOPFETCHTRADESRESPONSE,
        "__module__": "trades_pb2"
        # @@protoc_insertion_point(class_scope:fetcher.v1.StopFetchTradesResponse)
    },
)
_sym_db.RegisterMessage(StopFetchTradesResponse)

Empty = _reflection.GeneratedProtocolMessageType(
    "Empty",
    (_message.Message,),
    {
        "DESCRIPTOR": _EMPTY,
        "__module__": "trades_pb2"
        # @@protoc_insertion_point(class_scope:fetcher.v1.Empty)
    },
)
_sym_db.RegisterMessage(Empty)

_FETCHERSERVICE = DESCRIPTOR.services_by_name["FetcherService"]
if _descriptor._USE_C_DESCRIPTORS == False:

    DESCRIPTOR._options = None
    DESCRIPTOR._serialized_options = b"Z\006cw-api"
    _FETCHTRADESREQUEST._serialized_start = 28
    _FETCHTRADESREQUEST._serialized_end = 64
    _TASK._serialized_start = 66
    _TASK._serialized_end = 84
    _FETCHTRADESRESPONSE._serialized_start = 86
    _FETCHTRADESRESPONSE._serialized_end = 140
    _STOPFETCHTRADESRESPONSE._serialized_start = 142
    _STOPFETCHTRADESRESPONSE._serialized_end = 184
    _EMPTY._serialized_start = 186
    _EMPTY._serialized_end = 193
    _FETCHERSERVICE._serialized_start = 196
    _FETCHERSERVICE._serialized_end = 371
# @@protoc_insertion_point(module_scope)
