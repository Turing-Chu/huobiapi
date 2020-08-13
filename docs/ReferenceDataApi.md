# \ReferenceDataApi

All URIs are relative to *https://api.testnet.huobi.pro*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CurrencyAndChainsV2**](ReferenceDataApi.md#CurrencyAndChainsV2) | **Get** /v2/reference/currencies | Currency &amp; Chains
[**TimestampV1**](ReferenceDataApi.md#TimestampV1) | **Get** /v1/common/timestamp | Get Current Timestamp



## CurrencyAndChainsV2

> InlineResponse200 CurrencyAndChainsV2(ctx, currency, authorizedUser)

Currency & Chains

API user could query static reference information for each currency, as well as its corresponding chain(s). (Public Endpoint)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string**| btc, ltc, etc ...(available currencies in Huobi Global) | 
**authorizedUser** | **bool**| true or false (if not filled, default value is true) | 

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

