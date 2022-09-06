# TranslateDocument200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | Pointer to **string** | A unique ID assigned to the uploaded document and the translation process. Must be used when referring to this particular document in subsequent API requests. | [optional] 
**DocumentKey** | Pointer to **string** | A unique key that is used to encrypt the uploaded document as well as the resulting translation on the server side. Must be provided with every subsequent API request regarding this particular document. | [optional] 

## Methods

### NewTranslateDocument200Response

`func NewTranslateDocument200Response() *TranslateDocument200Response`

NewTranslateDocument200Response instantiates a new TranslateDocument200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranslateDocument200ResponseWithDefaults

`func NewTranslateDocument200ResponseWithDefaults() *TranslateDocument200Response`

NewTranslateDocument200ResponseWithDefaults instantiates a new TranslateDocument200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *TranslateDocument200Response) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *TranslateDocument200Response) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *TranslateDocument200Response) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *TranslateDocument200Response) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetDocumentKey

`func (o *TranslateDocument200Response) GetDocumentKey() string`

GetDocumentKey returns the DocumentKey field if non-nil, zero value otherwise.

### GetDocumentKeyOk

`func (o *TranslateDocument200Response) GetDocumentKeyOk() (*string, bool)`

GetDocumentKeyOk returns a tuple with the DocumentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentKey

`func (o *TranslateDocument200Response) SetDocumentKey(v string)`

SetDocumentKey sets DocumentKey field to given value.

### HasDocumentKey

`func (o *TranslateDocument200Response) HasDocumentKey() bool`

HasDocumentKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


