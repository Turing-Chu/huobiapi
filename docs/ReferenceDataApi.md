# \ReferenceDataApi

All URIs are relative to *http://api.testnet.huobi.pro/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2MarketStatusGet**](ReferenceDataApi.md#V2MarketStatusGet) | **Get** /v2/market-status | The endpoint returns current market status
[**V2SummaryJsonGet**](ReferenceDataApi.md#V2SummaryJsonGet) | **Get** /v2/summary.json | Get system status



## V2MarketStatusGet

> ModelsMarketStatus V2MarketStatusGet(ctx, )

The endpoint returns current market status

The enum values of market status includes: 1 - normal (order submission & cancellation are allowed)，

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ModelsMarketStatus**](models.MarketStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SummaryJsonGet

> map[string]interface{} V2SummaryJsonGet(ctx, )

Get system status

This endpoint allows users to get system status, Incidents and planned maintenance.

### Required Parameters

This endpoint does not need any parameter.

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

