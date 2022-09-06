# \TranslatingDocumentsApi

All URIs are relative to *https://api.deepl.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadDocument**](TranslatingDocumentsApi.md#DownloadDocument) | **Post** /document/{document_id}/result | Download Translated Document
[**GetDocumentStatus**](TranslatingDocumentsApi.md#GetDocumentStatus) | **Post** /document/{document_id} | Check Document Status
[**TranslateDocument**](TranslatingDocumentsApi.md#TranslateDocument) | **Post** /document | Upload and Translate a Document



## DownloadDocument

> *os.File DownloadDocument(ctx, documentId).GetDocumentStatusRequest(getDocumentStatusRequest).Execute()

Download Translated Document



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
    documentId := "04DE5AD98A02647D83285A36021911C6" // string | The document ID that was sent to the client when the document was uploaded to the API.
    getDocumentStatusRequest := TODO // GetDocumentStatusRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslatingDocumentsApi.DownloadDocument(context.Background(), documentId).GetDocumentStatusRequest(getDocumentStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslatingDocumentsApi.DownloadDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadDocument`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `TranslatingDocumentsApi.DownloadDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** | The document ID that was sent to the client when the document was uploaded to the API. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getDocumentStatusRequest** | [**GetDocumentStatusRequest**](GetDocumentStatusRequest.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[auth_header](../README.md#auth_header)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentStatus

> GetDocumentStatus200Response GetDocumentStatus(ctx, documentId).GetDocumentStatusRequest(getDocumentStatusRequest).Execute()

Check Document Status



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
    documentId := "04DE5AD98A02647D83285A36021911C6" // string | The document ID that was sent to the client when the document was uploaded to the API.
    getDocumentStatusRequest := TODO // GetDocumentStatusRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslatingDocumentsApi.GetDocumentStatus(context.Background(), documentId).GetDocumentStatusRequest(getDocumentStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslatingDocumentsApi.GetDocumentStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentStatus`: GetDocumentStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `TranslatingDocumentsApi.GetDocumentStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** | The document ID that was sent to the client when the document was uploaded to the API. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getDocumentStatusRequest** | [**GetDocumentStatusRequest**](GetDocumentStatusRequest.md) |  | 

### Return type

[**GetDocumentStatus200Response**](GetDocumentStatus200Response.md)

### Authorization

[auth_header](../README.md#auth_header)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslateDocument

> TranslateDocument200Response TranslateDocument(ctx).TargetLang(targetLang).File(file).SourceLang(sourceLang).Filename(filename).Formality(formality).GlossaryId(glossaryId).Execute()

Upload and Translate a Document



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
    targetLang := TODO // TargetLanguage | 
    file := os.NewFile(1234, "some_file") // *os.File | The document file to be translated. The file name should be included in this part's content disposition. As an alternative, the filename parameter can be used. The following file types and extensions are supported:   * `docx` - Microsoft Word Document   * `pptx` - Microsoft PowerPoint Document   * `pdf` - Portable Document Format   * `htm / html` - HTML Document   * `txt` - Plain Text Document  Please note that in order to translate PDF documents you need to give one-time consent to using the Adobe API via the account interface.
    sourceLang := openapiclient.SourceLanguage("BG") // SourceLanguage |  (optional)
    filename := "filename_example" // string | The name of the uploaded file. Can be used as an alternative to including the file name in the file part's content disposition. (optional)
    formality := TODO // Formality |  (optional) (default to "default")
    glossaryId := "glossaryId_example" // string | A unique ID assigned to a glossary. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslatingDocumentsApi.TranslateDocument(context.Background()).TargetLang(targetLang).File(file).SourceLang(sourceLang).Filename(filename).Formality(formality).GlossaryId(glossaryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslatingDocumentsApi.TranslateDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TranslateDocument`: TranslateDocument200Response
    fmt.Fprintf(os.Stdout, "Response from `TranslatingDocumentsApi.TranslateDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTranslateDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetLang** | [**TargetLanguage**](TargetLanguage.md) |  | 
 **file** | ***os.File** | The document file to be translated. The file name should be included in this part&#39;s content disposition. As an alternative, the filename parameter can be used. The following file types and extensions are supported:   * &#x60;docx&#x60; - Microsoft Word Document   * &#x60;pptx&#x60; - Microsoft PowerPoint Document   * &#x60;pdf&#x60; - Portable Document Format   * &#x60;htm / html&#x60; - HTML Document   * &#x60;txt&#x60; - Plain Text Document  Please note that in order to translate PDF documents you need to give one-time consent to using the Adobe API via the account interface. | 
 **sourceLang** | [**SourceLanguage**](SourceLanguage.md) |  | 
 **filename** | **string** | The name of the uploaded file. Can be used as an alternative to including the file name in the file part&#39;s content disposition. | 
 **formality** | [**Formality**](Formality.md) |  | [default to &quot;default&quot;]
 **glossaryId** | **string** | A unique ID assigned to a glossary. | 

### Return type

[**TranslateDocument200Response**](TranslateDocument200Response.md)

### Authorization

[auth_header](../README.md#auth_header)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

