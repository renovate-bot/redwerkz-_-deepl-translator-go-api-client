# \MetaInformationApi

All URIs are relative to *https://api.deepl.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLanguages**](MetaInformationApi.md#GetLanguages) | **Get** /languages | Retrieve Supported Languages
[**GetUsage**](MetaInformationApi.md#GetUsage) | **Get** /usage | Check Usage and Limits



## GetLanguages

> []GetLanguages200ResponseInner GetLanguages(ctx).Type_(type_).Execute()

Retrieve Supported Languages



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    type_ := "target" // string | Sets whether source or target languages should be listed. Possible options are:  * `source` (default): For languages that can be used in the `source_lang` parameter of [translate](https://www.deepl.com/docs-api/translating-text/translate-text) requests.  * `target`: For languages that can be used in the `target_lang` parameter of [translate](https://www.deepl.com/docs-api/translating-text/translate-text) requests. (optional) (default to "source")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaInformationApi.GetLanguages(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaInformationApi.GetLanguages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLanguages`: []GetLanguages200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `MetaInformationApi.GetLanguages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLanguagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Sets whether source or target languages should be listed. Possible options are:  * &#x60;source&#x60; (default): For languages that can be used in the &#x60;source_lang&#x60; parameter of [translate](https://www.deepl.com/docs-api/translating-text/translate-text) requests.  * &#x60;target&#x60;: For languages that can be used in the &#x60;target_lang&#x60; parameter of [translate](https://www.deepl.com/docs-api/translating-text/translate-text) requests. | [default to &quot;source&quot;]

### Return type

[**[]GetLanguages200ResponseInner**](GetLanguages200ResponseInner.md)

### Authorization

[auth_header](../README.md#auth_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsage

> GetUsage200Response GetUsage(ctx).Execute()

Check Usage and Limits



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaInformationApi.GetUsage(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaInformationApi.GetUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsage`: GetUsage200Response
    fmt.Fprintf(os.Stdout, "Response from `MetaInformationApi.GetUsage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageRequest struct via the builder pattern


### Return type

[**GetUsage200Response**](GetUsage200Response.md)

### Authorization

[auth_header](../README.md#auth_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

