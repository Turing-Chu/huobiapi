# \ReferenceDataApi

All URIs are relative to *https://api.testnet.huobi.pro*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CurrencyAndChainsV2**](ReferenceDataApi.md#CurrencyAndChainsV2) | **Get** /v2/reference/currencies | Currency &amp; Chains
[**MarketStatusV2**](ReferenceDataApi.md#MarketStatusV2) | **Get** /v2/market-status | The endpoint returns current market status
[**Summary**](ReferenceDataApi.md#Summary) | **Get** /api/v2/summary.json | Get system status
[**SupportCurrenciesV1**](ReferenceDataApi.md#SupportCurrenciesV1) | **Get** /v1/common/currencys | Get all Supported Currencies
[**SupportedTradingSymbolV1**](ReferenceDataApi.md#SupportedTradingSymbolV1) | **Get** /v1/common/symbols | Get all Supported Trading Symbol
[**TimestampV1**](ReferenceDataApi.md#TimestampV1) | **Get** /v1/common/timestamp | Get Current Timestamp



## CurrencyAndChainsV2

> InlineResponse2004 CurrencyAndChainsV2(ctx, currency, authorizedUser)

Currency & Chains

API user could query static reference information for each currency, as well as its corresponding chain(s). (Public Endpoint)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string**| btc, ltc, etc ...(available currencies in Huobi Global) | 
**authorizedUser** | **bool**| true or false (if not filled, default value is true) | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketStatusV2

> InlineResponse2003 MarketStatusV2(ctx, )

The endpoint returns current market status

The enum values of market status includes: 1 - normal (order submission & cancellation are allowed)，

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Summary

> InlineResponse200 Summary(ctx, )

Get system status

This endpoint allows users to get system status, Incidents and planned maintenance.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportCurrenciesV1

> InlineResponse2001 SupportCurrenciesV1(ctx, )

Get all Supported Currencies

This endpoint returns all Huobi's supported trading currencies.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportedTradingSymbolV1

> InlineResponse2002 SupportedTradingSymbolV1(ctx, )

Get all Supported Trading Symbol

This endpoint returns all Huobi's supported trading symbol.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TimestampV1

> int32 TimestampV1(ctx, )

Get Current Timestamp

This endpoint returns the current timestamp, i.e. the number of milliseconds that have elapsed since 00:00:00 UTC on 1 January 1970.

### Required Parameters

This endpoint does not need any parameter.

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

