# GetUsage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CharacterCount** | Pointer to **int64** | Characters translated so far in the current billing period. | [optional] 
**CharacterLimit** | Pointer to **int64** | Current maximum number of characters that can be translated per billing period. | [optional] 
**DocumentLimit** | Pointer to **int64** | Documents translated so far in the current billing period. | [optional] 
**DocumentCount** | Pointer to **int64** | Current maximum number of documents that can be translated per billing period. | [optional] 
**TeamDocumentLimit** | Pointer to **int64** | Documents translated by all users in the team so far in the current billing period. | [optional] 
**TeamDocumentCount** | Pointer to **int64** | Current maximum number of documents that can be translated by the team per billing period. | [optional] 

## Methods

### NewGetUsage200Response

`func NewGetUsage200Response() *GetUsage200Response`

NewGetUsage200Response instantiates a new GetUsage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUsage200ResponseWithDefaults

`func NewGetUsage200ResponseWithDefaults() *GetUsage200Response`

NewGetUsage200ResponseWithDefaults instantiates a new GetUsage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharacterCount

`func (o *GetUsage200Response) GetCharacterCount() int64`

GetCharacterCount returns the CharacterCount field if non-nil, zero value otherwise.

### GetCharacterCountOk

`func (o *GetUsage200Response) GetCharacterCountOk() (*int64, bool)`

GetCharacterCountOk returns a tuple with the CharacterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterCount

`func (o *GetUsage200Response) SetCharacterCount(v int64)`

SetCharacterCount sets CharacterCount field to given value.

### HasCharacterCount

`func (o *GetUsage200Response) HasCharacterCount() bool`

HasCharacterCount returns a boolean if a field has been set.

### GetCharacterLimit

`func (o *GetUsage200Response) GetCharacterLimit() int64`

GetCharacterLimit returns the CharacterLimit field if non-nil, zero value otherwise.

### GetCharacterLimitOk

`func (o *GetUsage200Response) GetCharacterLimitOk() (*int64, bool)`

GetCharacterLimitOk returns a tuple with the CharacterLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterLimit

`func (o *GetUsage200Response) SetCharacterLimit(v int64)`

SetCharacterLimit sets CharacterLimit field to given value.

### HasCharacterLimit

`func (o *GetUsage200Response) HasCharacterLimit() bool`

HasCharacterLimit returns a boolean if a field has been set.

### GetDocumentLimit

`func (o *GetUsage200Response) GetDocumentLimit() int64`

GetDocumentLimit returns the DocumentLimit field if non-nil, zero value otherwise.

### GetDocumentLimitOk

`func (o *GetUsage200Response) GetDocumentLimitOk() (*int64, bool)`

GetDocumentLimitOk returns a tuple with the DocumentLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentLimit

`func (o *GetUsage200Response) SetDocumentLimit(v int64)`

SetDocumentLimit sets DocumentLimit field to given value.

### HasDocumentLimit

`func (o *GetUsage200Response) HasDocumentLimit() bool`

HasDocumentLimit returns a boolean if a field has been set.

### GetDocumentCount

`func (o *GetUsage200Response) GetDocumentCount() int64`

GetDocumentCount returns the DocumentCount field if non-nil, zero value otherwise.

### GetDocumentCountOk

`func (o *GetUsage200Response) GetDocumentCountOk() (*int64, bool)`

GetDocumentCountOk returns a tuple with the DocumentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentCount

`func (o *GetUsage200Response) SetDocumentCount(v int64)`

SetDocumentCount sets DocumentCount field to given value.

### HasDocumentCount

`func (o *GetUsage200Response) HasDocumentCount() bool`

HasDocumentCount returns a boolean if a field has been set.

### GetTeamDocumentLimit

`func (o *GetUsage200Response) GetTeamDocumentLimit() int64`

GetTeamDocumentLimit returns the TeamDocumentLimit field if non-nil, zero value otherwise.

### GetTeamDocumentLimitOk

`func (o *GetUsage200Response) GetTeamDocumentLimitOk() (*int64, bool)`

GetTeamDocumentLimitOk returns a tuple with the TeamDocumentLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamDocumentLimit

`func (o *GetUsage200Response) SetTeamDocumentLimit(v int64)`

SetTeamDocumentLimit sets TeamDocumentLimit field to given value.

### HasTeamDocumentLimit

`func (o *GetUsage200Response) HasTeamDocumentLimit() bool`

HasTeamDocumentLimit returns a boolean if a field has been set.

### GetTeamDocumentCount

`func (o *GetUsage200Response) GetTeamDocumentCount() int64`

GetTeamDocumentCount returns the TeamDocumentCount field if non-nil, zero value otherwise.

### GetTeamDocumentCountOk

`func (o *GetUsage200Response) GetTeamDocumentCountOk() (*int64, bool)`

GetTeamDocumentCountOk returns a tuple with the TeamDocumentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamDocumentCount

`func (o *GetUsage200Response) SetTeamDocumentCount(v int64)`

SetTeamDocumentCount sets TeamDocumentCount field to given value.

### HasTeamDocumentCount

`func (o *GetUsage200Response) HasTeamDocumentCount() bool`

HasTeamDocumentCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


