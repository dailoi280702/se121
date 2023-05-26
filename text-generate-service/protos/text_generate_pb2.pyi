from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class GenerateBlogSummarizationReq(_message.Message):
    __slots__ = ["body", "title"]
    BODY_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    body: str
    title: str
    def __init__(self, title: _Optional[str] = ..., body: _Optional[str] = ...) -> None: ...

class GenerateReviewReq(_message.Message):
    __slots__ = ["brand", "fuelType", "horsePower", "name", "series", "torque", "transmission"]
    BRAND_FIELD_NUMBER: _ClassVar[int]
    FUELTYPE_FIELD_NUMBER: _ClassVar[int]
    HORSEPOWER_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    SERIES_FIELD_NUMBER: _ClassVar[int]
    TORQUE_FIELD_NUMBER: _ClassVar[int]
    TRANSMISSION_FIELD_NUMBER: _ClassVar[int]
    brand: str
    fuelType: str
    horsePower: int
    name: str
    series: str
    torque: int
    transmission: str
    def __init__(self, name: _Optional[str] = ..., brand: _Optional[str] = ..., series: _Optional[str] = ..., horsePower: _Optional[int] = ..., torque: _Optional[int] = ..., transmission: _Optional[str] = ..., fuelType: _Optional[str] = ...) -> None: ...

class ResString(_message.Message):
    __slots__ = ["text"]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    text: str
    def __init__(self, text: _Optional[str] = ...) -> None: ...
