/*
Terraform Private Registry

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tfclient

import (
	"encoding/json"
)

// ProviderPlatformResponse struct for ProviderPlatformResponse
type ProviderPlatformResponse struct {
	Data *ProviderPlatformResponseData `json:"data,omitempty"`
}

// NewProviderPlatformResponse instantiates a new ProviderPlatformResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderPlatformResponse() *ProviderPlatformResponse {
	this := ProviderPlatformResponse{}
	return &this
}

// NewProviderPlatformResponseWithDefaults instantiates a new ProviderPlatformResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderPlatformResponseWithDefaults() *ProviderPlatformResponse {
	this := ProviderPlatformResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ProviderPlatformResponse) GetData() ProviderPlatformResponseData {
	if o == nil || o.Data == nil {
		var ret ProviderPlatformResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderPlatformResponse) GetDataOk() (*ProviderPlatformResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ProviderPlatformResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ProviderPlatformResponseData and assigns it to the Data field.
func (o *ProviderPlatformResponse) SetData(v ProviderPlatformResponseData) {
	o.Data = &v
}

func (o ProviderPlatformResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableProviderPlatformResponse struct {
	value *ProviderPlatformResponse
	isSet bool
}

func (v NullableProviderPlatformResponse) Get() *ProviderPlatformResponse {
	return v.value
}

func (v *NullableProviderPlatformResponse) Set(val *ProviderPlatformResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderPlatformResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderPlatformResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderPlatformResponse(val *ProviderPlatformResponse) *NullableProviderPlatformResponse {
	return &NullableProviderPlatformResponse{value: val, isSet: true}
}

func (v NullableProviderPlatformResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderPlatformResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}