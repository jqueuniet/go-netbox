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

// checks if the CustomFieldFilterLogic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomFieldFilterLogic{}

// CustomFieldFilterLogic struct for CustomFieldFilterLogic
type CustomFieldFilterLogic struct {
	Value                *CustomFieldFilterLogicValue `json:"value,omitempty"`
	Label                *CustomFieldFilterLogicLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomFieldFilterLogic CustomFieldFilterLogic

// NewCustomFieldFilterLogic instantiates a new CustomFieldFilterLogic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFieldFilterLogic() *CustomFieldFilterLogic {
	this := CustomFieldFilterLogic{}
	return &this
}

// NewCustomFieldFilterLogicWithDefaults instantiates a new CustomFieldFilterLogic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFieldFilterLogicWithDefaults() *CustomFieldFilterLogic {
	this := CustomFieldFilterLogic{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CustomFieldFilterLogic) GetValue() CustomFieldFilterLogicValue {
	if o == nil || IsNil(o.Value) {
		var ret CustomFieldFilterLogicValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldFilterLogic) GetValueOk() (*CustomFieldFilterLogicValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CustomFieldFilterLogic) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given CustomFieldFilterLogicValue and assigns it to the Value field.
func (o *CustomFieldFilterLogic) SetValue(v CustomFieldFilterLogicValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CustomFieldFilterLogic) GetLabel() CustomFieldFilterLogicLabel {
	if o == nil || IsNil(o.Label) {
		var ret CustomFieldFilterLogicLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldFilterLogic) GetLabelOk() (*CustomFieldFilterLogicLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CustomFieldFilterLogic) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given CustomFieldFilterLogicLabel and assigns it to the Label field.
func (o *CustomFieldFilterLogic) SetLabel(v CustomFieldFilterLogicLabel) {
	o.Label = &v
}

func (o CustomFieldFilterLogic) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFieldFilterLogic) ToMap() (map[string]interface{}, error) {
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

func (o *CustomFieldFilterLogic) UnmarshalJSON(data []byte) (err error) {
	varCustomFieldFilterLogic := _CustomFieldFilterLogic{}

	err = json.Unmarshal(data, &varCustomFieldFilterLogic)

	if err != nil {
		return err
	}

	*o = CustomFieldFilterLogic(varCustomFieldFilterLogic)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomFieldFilterLogic struct {
	value *CustomFieldFilterLogic
	isSet bool
}

func (v NullableCustomFieldFilterLogic) Get() *CustomFieldFilterLogic {
	return v.value
}

func (v *NullableCustomFieldFilterLogic) Set(val *CustomFieldFilterLogic) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldFilterLogic) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldFilterLogic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldFilterLogic(val *CustomFieldFilterLogic) *NullableCustomFieldFilterLogic {
	return &NullableCustomFieldFilterLogic{value: val, isSet: true}
}

func (v NullableCustomFieldFilterLogic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldFilterLogic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
