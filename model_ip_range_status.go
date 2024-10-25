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

// checks if the IPRangeStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPRangeStatus{}

// IPRangeStatus struct for IPRangeStatus
type IPRangeStatus struct {
	Value                *IPRangeStatusValue `json:"value,omitempty"`
	Label                *IPRangeStatusLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IPRangeStatus IPRangeStatus

// NewIPRangeStatus instantiates a new IPRangeStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPRangeStatus() *IPRangeStatus {
	this := IPRangeStatus{}
	return &this
}

// NewIPRangeStatusWithDefaults instantiates a new IPRangeStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPRangeStatusWithDefaults() *IPRangeStatus {
	this := IPRangeStatus{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IPRangeStatus) GetValue() IPRangeStatusValue {
	if o == nil || IsNil(o.Value) {
		var ret IPRangeStatusValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRangeStatus) GetValueOk() (*IPRangeStatusValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IPRangeStatus) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given IPRangeStatusValue and assigns it to the Value field.
func (o *IPRangeStatus) SetValue(v IPRangeStatusValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *IPRangeStatus) GetLabel() IPRangeStatusLabel {
	if o == nil || IsNil(o.Label) {
		var ret IPRangeStatusLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRangeStatus) GetLabelOk() (*IPRangeStatusLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *IPRangeStatus) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given IPRangeStatusLabel and assigns it to the Label field.
func (o *IPRangeStatus) SetLabel(v IPRangeStatusLabel) {
	o.Label = &v
}

func (o IPRangeStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPRangeStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IPRangeStatus) UnmarshalJSON(data []byte) (err error) {
	varIPRangeStatus := _IPRangeStatus{}

	err = json.Unmarshal(data, &varIPRangeStatus)

	if err != nil {
		return err
	}

	*o = IPRangeStatus(varIPRangeStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIPRangeStatus struct {
	value *IPRangeStatus
	isSet bool
}

func (v NullableIPRangeStatus) Get() *IPRangeStatus {
	return v.value
}

func (v *NullableIPRangeStatus) Set(val *IPRangeStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIPRangeStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIPRangeStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPRangeStatus(val *IPRangeStatus) *NullableIPRangeStatus {
	return &NullableIPRangeStatus{value: val, isSet: true}
}

func (v NullableIPRangeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPRangeStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
