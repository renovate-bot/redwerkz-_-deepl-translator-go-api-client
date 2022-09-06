/*
DeepL API

The DeepL API provides programmatic access to DeepL’s machine translation technology.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GetDocumentStatus200Response struct for GetDocumentStatus200Response
type GetDocumentStatus200Response struct {
	// A unique ID assigned to the uploaded document and the requested translation process. The same ID that was used when requesting the translation status.
	DocumentId string `json:"document_id"`
	// A short description of the state the document translation process is currently in. Possible values are:  * `queued` - the translation job is waiting in line to be processed  * `translating` - the translation is currently ongoing  * `done` - the translation is done and the translated document is ready for download  * `error` - an irrecoverable error occurred while translating the document
	Status string `json:"status"`
	// Estimated number of seconds until the translation is done. This parameter is only included while `status` is `\"translating\"`.
	SecondsRemaining *int32 `json:"seconds_remaining,omitempty"`
	// The number of characters billed to your account.
	BilledCharacters *int32 `json:"billed_characters,omitempty"`
	// A short description of the error, if available. Note that the content is subject to change. This parameter may be included if an error occurred during translation.
	ErrorMessage *string `json:"error_message,omitempty"`
}

// NewGetDocumentStatus200Response instantiates a new GetDocumentStatus200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDocumentStatus200Response(documentId string, status string) *GetDocumentStatus200Response {
	this := GetDocumentStatus200Response{}
	this.DocumentId = documentId
	this.Status = status
	return &this
}

// NewGetDocumentStatus200ResponseWithDefaults instantiates a new GetDocumentStatus200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDocumentStatus200ResponseWithDefaults() *GetDocumentStatus200Response {
	this := GetDocumentStatus200Response{}
	return &this
}

// GetDocumentId returns the DocumentId field value
func (o *GetDocumentStatus200Response) GetDocumentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value
// and a boolean to check if the value has been set.
func (o *GetDocumentStatus200Response) GetDocumentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentId, true
}

// SetDocumentId sets field value
func (o *GetDocumentStatus200Response) SetDocumentId(v string) {
	o.DocumentId = v
}

// GetStatus returns the Status field value
func (o *GetDocumentStatus200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetDocumentStatus200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetDocumentStatus200Response) SetStatus(v string) {
	o.Status = v
}

// GetSecondsRemaining returns the SecondsRemaining field value if set, zero value otherwise.
func (o *GetDocumentStatus200Response) GetSecondsRemaining() int32 {
	if o == nil || o.SecondsRemaining == nil {
		var ret int32
		return ret
	}
	return *o.SecondsRemaining
}

// GetSecondsRemainingOk returns a tuple with the SecondsRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDocumentStatus200Response) GetSecondsRemainingOk() (*int32, bool) {
	if o == nil || o.SecondsRemaining == nil {
		return nil, false
	}
	return o.SecondsRemaining, true
}

// HasSecondsRemaining returns a boolean if a field has been set.
func (o *GetDocumentStatus200Response) HasSecondsRemaining() bool {
	if o != nil && o.SecondsRemaining != nil {
		return true
	}

	return false
}

// SetSecondsRemaining gets a reference to the given int32 and assigns it to the SecondsRemaining field.
func (o *GetDocumentStatus200Response) SetSecondsRemaining(v int32) {
	o.SecondsRemaining = &v
}

// GetBilledCharacters returns the BilledCharacters field value if set, zero value otherwise.
func (o *GetDocumentStatus200Response) GetBilledCharacters() int32 {
	if o == nil || o.BilledCharacters == nil {
		var ret int32
		return ret
	}
	return *o.BilledCharacters
}

// GetBilledCharactersOk returns a tuple with the BilledCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDocumentStatus200Response) GetBilledCharactersOk() (*int32, bool) {
	if o == nil || o.BilledCharacters == nil {
		return nil, false
	}
	return o.BilledCharacters, true
}

// HasBilledCharacters returns a boolean if a field has been set.
func (o *GetDocumentStatus200Response) HasBilledCharacters() bool {
	if o != nil && o.BilledCharacters != nil {
		return true
	}

	return false
}

// SetBilledCharacters gets a reference to the given int32 and assigns it to the BilledCharacters field.
func (o *GetDocumentStatus200Response) SetBilledCharacters(v int32) {
	o.BilledCharacters = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *GetDocumentStatus200Response) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDocumentStatus200Response) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *GetDocumentStatus200Response) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *GetDocumentStatus200Response) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o GetDocumentStatus200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["document_id"] = o.DocumentId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.SecondsRemaining != nil {
		toSerialize["seconds_remaining"] = o.SecondsRemaining
	}
	if o.BilledCharacters != nil {
		toSerialize["billed_characters"] = o.BilledCharacters
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	return json.Marshal(toSerialize)
}

type NullableGetDocumentStatus200Response struct {
	value *GetDocumentStatus200Response
	isSet bool
}

func (v NullableGetDocumentStatus200Response) Get() *GetDocumentStatus200Response {
	return v.value
}

func (v *NullableGetDocumentStatus200Response) Set(val *GetDocumentStatus200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDocumentStatus200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDocumentStatus200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDocumentStatus200Response(val *GetDocumentStatus200Response) *NullableGetDocumentStatus200Response {
	return &NullableGetDocumentStatus200Response{value: val, isSet: true}
}

func (v NullableGetDocumentStatus200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDocumentStatus200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


