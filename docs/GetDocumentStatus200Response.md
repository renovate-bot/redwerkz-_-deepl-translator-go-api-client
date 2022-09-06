# GetDocumentStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | **string** | A unique ID assigned to the uploaded document and the requested translation process. The same ID that was used when requesting the translation status. | 
**Status** | **string** | A short description of the state the document translation process is currently in. Possible values are:  * &#x60;queued&#x60; - the translation job is waiting in line to be processed  * &#x60;translating&#x60; - the translation is currently ongoing  * &#x60;done&#x60; - the translation is done and the translated document is ready for download  * &#x60;error&#x60; - an irrecoverable error occurred while translating the document | 
**SecondsRemaining** | Pointer to **int32** | Estimated number of seconds until the translation is done. This parameter is only included while &#x60;status&#x60; is &#x60;\&quot;translating\&quot;&#x60;. | [optional] 
**BilledCharacters** | Pointer to **int32** | The number of characters billed to your account. | [optional] 
**ErrorMessage** | Pointer to **string** | A short description of the error, if available. Note that the content is subject to change. This parameter may be included if an error occurred during translation. | [optional] 

## Methods

### NewGetDocumentStatus200Response

`func NewGetDocumentStatus200Response(documentId string, status string, ) *GetDocumentStatus200Response`

NewGetDocumentStatus200Response instantiates a new GetDocumentStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDocumentStatus200ResponseWithDefaults

`func NewGetDocumentStatus200ResponseWithDefaults() *GetDocumentStatus200Response`

NewGetDocumentStatus200ResponseWithDefaults instantiates a new GetDocumentStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *GetDocumentStatus200Response) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *GetDocumentStatus200Response) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *GetDocumentStatus200Response) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetStatus

`func (o *GetDocumentStatus200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDocumentStatus200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDocumentStatus200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSecondsRemaining

`func (o *GetDocumentStatus200Response) GetSecondsRemaining() int32`

GetSecondsRemaining returns the SecondsRemaining field if non-nil, zero value otherwise.

### GetSecondsRemainingOk

`func (o *GetDocumentStatus200Response) GetSecondsRemainingOk() (*int32, bool)`

GetSecondsRemainingOk returns a tuple with the SecondsRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsRemaining

`func (o *GetDocumentStatus200Response) SetSecondsRemaining(v int32)`

SetSecondsRemaining sets SecondsRemaining field to given value.

### HasSecondsRemaining

`func (o *GetDocumentStatus200Response) HasSecondsRemaining() bool`

HasSecondsRemaining returns a boolean if a field has been set.

### GetBilledCharacters

`func (o *GetDocumentStatus200Response) GetBilledCharacters() int32`

GetBilledCharacters returns the BilledCharacters field if non-nil, zero value otherwise.

### GetBilledCharactersOk

`func (o *GetDocumentStatus200Response) GetBilledCharactersOk() (*int32, bool)`

GetBilledCharactersOk returns a tuple with the BilledCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledCharacters

`func (o *GetDocumentStatus200Response) SetBilledCharacters(v int32)`

SetBilledCharacters sets BilledCharacters field to given value.

### HasBilledCharacters

`func (o *GetDocumentStatus200Response) HasBilledCharacters() bool`

HasBilledCharacters returns a boolean if a field has been set.

### GetErrorMessage

`func (o *GetDocumentStatus200Response) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetDocumentStatus200Response) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetDocumentStatus200Response) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetDocumentStatus200Response) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


