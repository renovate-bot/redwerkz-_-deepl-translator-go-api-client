# \ManagingGlossariesApi

All URIs are relative to *https://api.deepl.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGlossary**](ManagingGlossariesApi.md#CreateGlossary) | **Post** /glossaries | Create a Glossary
[**DeleteGlossary**](ManagingGlossariesApi.md#DeleteGlossary) | **Delete** /glossaries/{glossary_id} | Delete a Glossary
[**GetGlossary**](ManagingGlossariesApi.md#GetGlossary) | **Get** /glossaries/{glossary_id} | Retrieve Glossary Details
[**GetGlossaryEntries**](ManagingGlossariesApi.md#GetGlossaryEntries) | **Get** /glossaries/{glossary_id}/entries | Retrieve Glossary Entries
[**ListGlossaries**](ManagingGlossariesApi.md#ListGlossaries) | **Get** /glossaries | List all Glossaries
[**ListGlossaryLanguages**](ManagingGlossariesApi.md#ListGlossaryLanguages) | **Get** /glossary-language-pairs | List Language Pairs Supported by Glossaries



## CreateGlossary

> Glossary CreateGlossary(ctx).Name(name).SourceLang(sourceLang).TargetLang(targetLang).Entries(entries).EntriesFormat(entriesFormat).Execute()

Create a Glossary

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
    name := "name_example" // string | Name to be associated with the glossary.
    sourceLang := openapiclient.GlossarySourceLanguage("de") // GlossarySourceLanguage | 
    targetLang := openapiclient.GlossaryTargetLanguage("de") // GlossaryTargetLanguage | 
    entries := "entries_example" // string | The entries of the glossary. The entries have to be specified in the format provided by the `entries_format` parameter.
    entriesFormat := "entriesFormat_example" // string | The format in which the glossary entries are provided. Formats currently available: * `tsv` - Tab-separated values. Entries have to be specified as tab-separated values with the \\\"source entry\\\" being the text in the source language of the glossary and the \\\"target entry\\\" being the text in the target language of the glossary.    In addition the following restrictions apply:     * Duplicate source entries are not allowed.     * Source-target entry pairs are separated by a newline.     * Source entries and target entries are separated by a tab.     * Source entries and target entries are not empty.     * Source and target entries must not contain any [C0 or C1 control characters](https://en.wikipedia.org/wiki/C0_and_C1_control_codes) (including e.g. `\\\"\\\\t\\\"` or `\\\"\\\\n\\\"`) or any [Unicode newline](https://en.wikipedia.org/wiki/Newline#Unicode).     * Source and target entries must not contain any leading or trailing Unicode whitespace.    Valid glossary entries in the TSV format could be created in a programming language with backslash escape sequences (e.g. Python, JavaScript, etc.) like this:    `\\\"sourceEntry1\\\\ttargetEntry1\\\\nsourceEntry2\\\\targetEntry2\\\"`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagingGlossariesApi.CreateGlossary(context.Background()).Name(name).SourceLang(sourceLang).TargetLang(targetLang).Entries(entries).EntriesFormat(entriesFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagingGlossariesApi.CreateGlossary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGlossary`: Glossary
    fmt.Fprintf(os.Stdout, "Response from `ManagingGlossariesApi.CreateGlossary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGlossaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name to be associated with the glossary. | 
 **sourceLang** | [**GlossarySourceLanguage**](GlossarySourceLanguage.md) |  | 
 **targetLang** | [**GlossaryTargetLanguage**](GlossaryTargetLanguage.md) |  | 
 **entries** | **string** | The entries of the glossary. The entries have to be specified in the format provided by the &#x60;entries_format&#x60; parameter. | 
 **entriesFormat** | **string** | The format in which the glossary entries are provided. Formats currently available: * &#x60;tsv&#x60; - Tab-separated values. Entries have to be specified as tab-separated values with the \\\&quot;source entry\\\&quot; being the text in the source language of the glossary and the \\\&quot;target entry\\\&quot; being the text in the target language of the glossary.    In addition the following restrictions apply:     * Duplicate source entries are not allowed.     * Source-target entry pairs are separated by a newline.     * Source entries and target entries are separated by a tab.     * Source entries and target entries are not empty.     * Source and target entries must not contain any [C0 or C1 control characters](https://en.wikipedia.org/wiki/C0_and_C1_control_codes) (including e.g. &#x60;\\\&quot;\\\\t\\\&quot;&#x60; or &#x60;\\\&quot;\\\\n\\\&quot;&#x60;) or any [Unicode newline](https://en.wikipedia.org/wiki/Newline#Unicode).     * Source and target entries must not contain any leading or trailing Unicode whitespace.    Valid glossary entries in the TSV format could be created in a programming language with backslash escape sequences (e.g. Python, JavaScript, etc.) like this:    &#x60;\\\&quot;sourceEntry1\\\\ttargetEntry1\\\\nsourceEntry2\\\\targetEntry2\\\&quot;&#x60; | 

### Return type

[**Glossary**](Glossary.md)

### Authorization

[auth_header](../README.md#auth_header)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGlossary

> DeleteGlossary(ctx, glossaryId).Execute()

Delete a Glossary



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
    glossaryId := "glossaryId_example" // string | A unique ID assigned to the glossary.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagingGlossariesApi.DeleteGlossary(context.Background(), glossaryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagingGlossariesApi.DeleteGlossary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**glossaryId** | **string** | A unique ID assigned to the glossary. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGlossaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[auth_header](../README.md#auth_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlossary

> Glossary GetGlossary(ctx, glossaryId).Execute()

Retrieve Glossary Details



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
    glossaryId := "glossaryId_example" // string | A unique ID assigned to the glossary.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagingGlossariesApi.GetGlossary(context.Background(), glossaryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagingGlossariesApi.GetGlossary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlossary`: Glossary
    fmt.Fprintf(os.Stdout, "Response from `ManagingGlossariesApi.GetGlossary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**glossaryId** | **string** | A unique ID assigned to the glossary. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlossaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Glossary**](Glossary.md)

### Authorization

[auth_header](../README.md#auth_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlossaryEntries

> GetGlossaryEntries(ctx, glossaryId).Accept(accept).Execute()

Retrieve Glossary Entries



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
    glossaryId := "glossaryId_example" // string | A unique ID assigned to the glossary.
    accept := "text/tab-separated-values" // string | The requested format of the returned glossary entries. Currently, supports only `text/tab-separated-values`. (optional) (default to "text/tab-separated-values")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagingGlossariesApi.GetGlossaryEntries(context.Background(), glossaryId).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagingGlossariesApi.GetGlossaryEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**glossaryId** | **string** | A unique ID assigned to the glossary. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlossaryEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** | The requested format of the returned glossary entries. Currently, supports only &#x60;text/tab-separated-values&#x60;. | [default to &quot;text/tab-separated-values&quot;]

### Return type

 (empty response body)

### Authorization

[auth_header](../README.md#auth_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/tab-separated-values, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlossaries

> ListGlossaries200Response ListGlossaries(ctx).Execute()

List all Glossaries



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
    resp, r, err := apiClient.ManagingGlossariesApi.ListGlossaries(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagingGlossariesApi.ListGlossaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGlossaries`: ListGlossaries200Response
    fmt.Fprintf(os.Stdout, "Response from `ManagingGlossariesApi.ListGlossaries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGlossariesRequest struct via the builder pattern


### Return type

[**ListGlossaries200Response**](ListGlossaries200Response.md)

### Authorization

[auth_header](../README.md#auth_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlossaryLanguages

> ListGlossaryLanguages200Response ListGlossaryLanguages(ctx).Execute()

List Language Pairs Supported by Glossaries



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
    resp, r, err := apiClient.ManagingGlossariesApi.ListGlossaryLanguages(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagingGlossariesApi.ListGlossaryLanguages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGlossaryLanguages`: ListGlossaryLanguages200Response
    fmt.Fprintf(os.Stdout, "Response from `ManagingGlossariesApi.ListGlossaryLanguages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGlossaryLanguagesRequest struct via the builder pattern


### Return type

[**ListGlossaryLanguages200Response**](ListGlossaryLanguages200Response.md)

### Authorization

[auth_header](../README.md#auth_header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

