# GetLanguages200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | **string** | The language code of the given language. | 
**Name** | **string** | Name of the language in English. | 
**SupportsFormality** | Pointer to **bool** | Denotes formality support in case of a target language listing. | [optional] 

## Methods

### NewGetLanguages200ResponseInner

`func NewGetLanguages200ResponseInner(language string, name string, ) *GetLanguages200ResponseInner`

NewGetLanguages200ResponseInner instantiates a new GetLanguages200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLanguages200ResponseInnerWithDefaults

`func NewGetLanguages200ResponseInnerWithDefaults() *GetLanguages200ResponseInner`

NewGetLanguages200ResponseInnerWithDefaults instantiates a new GetLanguages200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *GetLanguages200ResponseInner) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *GetLanguages200ResponseInner) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *GetLanguages200ResponseInner) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetName

`func (o *GetLanguages200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetLanguages200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetLanguages200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetSupportsFormality

`func (o *GetLanguages200ResponseInner) GetSupportsFormality() bool`

GetSupportsFormality returns the SupportsFormality field if non-nil, zero value otherwise.

### GetSupportsFormalityOk

`func (o *GetLanguages200ResponseInner) GetSupportsFormalityOk() (*bool, bool)`

GetSupportsFormalityOk returns a tuple with the SupportsFormality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsFormality

`func (o *GetLanguages200ResponseInner) SetSupportsFormality(v bool)`

SetSupportsFormality sets SupportsFormality field to given value.

### HasSupportsFormality

`func (o *GetLanguages200ResponseInner) HasSupportsFormality() bool`

HasSupportsFormality returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


