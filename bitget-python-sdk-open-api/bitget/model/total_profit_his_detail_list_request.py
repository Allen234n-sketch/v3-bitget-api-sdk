# coding: utf-8

"""
    Bitget Open API

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 2.0.0
    Generated by: https://openapi-generator.tech
"""

from datetime import date, datetime  # noqa: F401
import decimal  # noqa: F401
import functools  # noqa: F401
import io  # noqa: F401
import re  # noqa: F401
import typing  # noqa: F401
import typing_extensions  # noqa: F401
import uuid  # noqa: F401

import frozendict  # noqa: F401

from bitget import schemas  # noqa: F401


class TotalProfitHisDetailListRequest(
    schemas.DictSchema
):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """


    class MetaOapg:
        required = {
            "date",
            "coinName",
        }
        
        class properties:
            coinName = schemas.StrSchema
            date = schemas.StrSchema
            pageNo = schemas.StrSchema
            pageSize = schemas.StrSchema
            __annotations__ = {
                "coinName": coinName,
                "date": date,
                "pageNo": pageNo,
                "pageSize": pageSize,
            }
    
    date: MetaOapg.properties.date
    coinName: MetaOapg.properties.coinName
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["coinName"]) -> MetaOapg.properties.coinName: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["date"]) -> MetaOapg.properties.date: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["pageNo"]) -> MetaOapg.properties.pageNo: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["pageSize"]) -> MetaOapg.properties.pageSize: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["coinName", "date", "pageNo", "pageSize", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["coinName"]) -> MetaOapg.properties.coinName: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["date"]) -> MetaOapg.properties.date: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["pageNo"]) -> typing.Union[MetaOapg.properties.pageNo, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["pageSize"]) -> typing.Union[MetaOapg.properties.pageSize, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["coinName", "date", "pageNo", "pageSize", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *args: typing.Union[dict, frozendict.frozendict, ],
        date: typing.Union[MetaOapg.properties.date, str, ],
        coinName: typing.Union[MetaOapg.properties.coinName, str, ],
        pageNo: typing.Union[MetaOapg.properties.pageNo, str, schemas.Unset] = schemas.unset,
        pageSize: typing.Union[MetaOapg.properties.pageSize, str, schemas.Unset] = schemas.unset,
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'TotalProfitHisDetailListRequest':
        return super().__new__(
            cls,
            *args,
            date=date,
            coinName=coinName,
            pageNo=pageNo,
            pageSize=pageSize,
            _configuration=_configuration,
            **kwargs,
        )