# TranslateText200ResponseTranslationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetectedSourceLanguage** | Pointer to [**NullableSourceLanguage**](SourceLanguage.md) |  | [optional] 
**Text** | Pointer to **string** | The translated text. | [optional] 

## Methods

### NewTranslateText200ResponseTranslationsInner

`func NewTranslateText200ResponseTranslationsInner() *TranslateText200ResponseTranslationsInner`

NewTranslateText200ResponseTranslationsInner instantiates a new TranslateText200ResponseTranslationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranslateText200ResponseTranslationsInnerWithDefaults

`func NewTranslateText200ResponseTranslationsInnerWithDefaults() *TranslateText200ResponseTranslationsInner`

NewTranslateText200ResponseTranslationsInnerWithDefaults instantiates a new TranslateText200ResponseTranslationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetectedSourceLanguage

`func (o *TranslateText200ResponseTranslationsInner) GetDetectedSourceLanguage() SourceLanguage`

GetDetectedSourceLanguage returns the DetectedSourceLanguage field if non-nil, zero value otherwise.

### GetDetectedSourceLanguageOk

`func (o *TranslateText200ResponseTranslationsInner) GetDetectedSourceLanguageOk() (*SourceLanguage, bool)`

GetDetectedSourceLanguageOk returns a tuple with the DetectedSourceLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedSourceLanguage

`func (o *TranslateText200ResponseTranslationsInner) SetDetectedSourceLanguage(v SourceLanguage)`

SetDetectedSourceLanguage sets DetectedSourceLanguage field to given value.

### HasDetectedSourceLanguage

`func (o *TranslateText200ResponseTranslationsInner) HasDetectedSourceLanguage() bool`

HasDetectedSourceLanguage returns a boolean if a field has been set.

### SetDetectedSourceLanguageNil

`func (o *TranslateText200ResponseTranslationsInner) SetDetectedSourceLanguageNil(b bool)`

 SetDetectedSourceLanguageNil sets the value for DetectedSourceLanguage to be an explicit nil

### UnsetDetectedSourceLanguage
`func (o *TranslateText200ResponseTranslationsInner) UnsetDetectedSourceLanguage()`

UnsetDetectedSourceLanguage ensures that no value is present for DetectedSourceLanguage, not even an explicit nil
### GetText

`func (o *TranslateText200ResponseTranslationsInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TranslateText200ResponseTranslationsInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TranslateText200ResponseTranslationsInner) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TranslateText200ResponseTranslationsInner) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


