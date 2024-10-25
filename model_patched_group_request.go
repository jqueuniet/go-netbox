/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.7.8 (3.7)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the PatchedGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedGroupRequest{}

// PatchedGroupRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedGroupRequest struct {
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedGroupRequest PatchedGroupRequest

// NewPatchedGroupRequest instantiates a new PatchedGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedGroupRequest() *PatchedGroupRequest {
	this := PatchedGroupRequest{}
	return &this
}

// NewPatchedGroupRequestWithDefaults instantiates a new PatchedGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedGroupRequestWithDefaults() *PatchedGroupRequest {
	this := PatchedGroupRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedGroupRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGroupRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedGroupRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedGroupRequest) SetName(v string) {
	o.Name = &v
}

func (o PatchedGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedGroupRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedGroupRequest := _PatchedGroupRequest{}

	err = json.Unmarshal(data, &varPatchedGroupRequest)

	if err != nil {
		return err
	}

	*o = PatchedGroupRequest(varPatchedGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedGroupRequest struct {
	value *PatchedGroupRequest
	isSet bool
}

func (v NullablePatchedGroupRequest) Get() *PatchedGroupRequest {
	return v.value
}

func (v *NullablePatchedGroupRequest) Set(val *PatchedGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedGroupRequest(val *PatchedGroupRequest) *NullablePatchedGroupRequest {
	return &NullablePatchedGroupRequest{value: val, isSet: true}
}

func (v NullablePatchedGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
