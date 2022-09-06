# Glossary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlossaryId** | Pointer to **string** | A unique ID assigned to a glossary. | [optional] 
**Name** | Pointer to **string** | Name associated with the glossary. | [optional] 
**Ready** | Pointer to **bool** | Indicates if the newly created glossary can already be used in &#x60;translate&#x60; requests. If the created glossary is not yet ready, you have to wait and check the &#x60;ready&#x60; status of the glossary before using it in a &#x60;translate&#x60; request. | [optional] 
**SourceLang** | Pointer to [**GlossarySourceLanguage**](GlossarySourceLanguage.md) |  | [optional] 
**TargetLang** | Pointer to [**GlossaryTargetLanguage**](GlossaryTargetLanguage.md) |  | [optional] 
**CreationTime** | Pointer to **time.Time** | The creation time of the glossary in the ISO 8601-1:2019 format (e.g.: &#x60;2021-08-03T14:16:18.329Z&#x60;). | [optional] 
**EntryCount** | Pointer to **int32** | The number of entries in the glossary. | [optional] 

## Methods

### NewGlossary

`func NewGlossary() *Glossary`

NewGlossary instantiates a new Glossary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlossaryWithDefaults

`func NewGlossaryWithDefaults() *Glossary`

NewGlossaryWithDefaults instantiates a new Glossary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlossaryId

`func (o *Glossary) GetGlossaryId() string`

GetGlossaryId returns the GlossaryId field if non-nil, zero value otherwise.

### GetGlossaryIdOk

`func (o *Glossary) GetGlossaryIdOk() (*string, bool)`

GetGlossaryIdOk returns a tuple with the GlossaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlossaryId

`func (o *Glossary) SetGlossaryId(v string)`

SetGlossaryId sets GlossaryId field to given value.

### HasGlossaryId

`func (o *Glossary) HasGlossaryId() bool`

HasGlossaryId returns a boolean if a field has been set.

### GetName

`func (o *Glossary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Glossary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Glossary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Glossary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReady

`func (o *Glossary) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *Glossary) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *Glossary) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *Glossary) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetSourceLang

`func (o *Glossary) GetSourceLang() GlossarySourceLanguage`

GetSourceLang returns the SourceLang field if non-nil, zero value otherwise.

### GetSourceLangOk

`func (o *Glossary) GetSourceLangOk() (*GlossarySourceLanguage, bool)`

GetSourceLangOk returns a tuple with the SourceLang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLang

`func (o *Glossary) SetSourceLang(v GlossarySourceLanguage)`

SetSourceLang sets SourceLang field to given value.

### HasSourceLang

`func (o *Glossary) HasSourceLang() bool`

HasSourceLang returns a boolean if a field has been set.

### GetTargetLang

`func (o *Glossary) GetTargetLang() GlossaryTargetLanguage`

GetTargetLang returns the TargetLang field if non-nil, zero value otherwise.

### GetTargetLangOk

`func (o *Glossary) GetTargetLangOk() (*GlossaryTargetLanguage, bool)`

GetTargetLangOk returns a tuple with the TargetLang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLang

`func (o *Glossary) SetTargetLang(v GlossaryTargetLanguage)`

SetTargetLang sets TargetLang field to given value.

### HasTargetLang

`func (o *Glossary) HasTargetLang() bool`

HasTargetLang returns a boolean if a field has been set.

### GetCreationTime

`func (o *Glossary) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Glossary) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Glossary) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Glossary) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEntryCount

`func (o *Glossary) GetEntryCount() int32`

GetEntryCount returns the EntryCount field if non-nil, zero value otherwise.

### GetEntryCountOk

`func (o *Glossary) GetEntryCountOk() (*int32, bool)`

GetEntryCountOk returns a tuple with the EntryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryCount

`func (o *Glossary) SetEntryCount(v int32)`

SetEntryCount sets EntryCount field to given value.

### HasEntryCount

`func (o *Glossary) HasEntryCount() bool`

HasEntryCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


