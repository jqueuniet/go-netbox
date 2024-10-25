/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.7.8 (3.7)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// checks if the ComponentNestedModuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentNestedModuleRequest{}

// ComponentNestedModuleRequest Used by device component serializers.
type ComponentNestedModuleRequest struct {
	Device               int32 `json:"device"`
	AdditionalProperties map[string]interface{}
}

type _ComponentNestedModuleRequest ComponentNestedModuleRequest

// NewComponentNestedModuleRequest instantiates a new ComponentNestedModuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentNestedModuleRequest(device int32) *ComponentNestedModuleRequest {
	this := ComponentNestedModuleRequest{}
	this.Device = device
	return &this
}

// NewComponentNestedModuleRequestWithDefaults instantiates a new ComponentNestedModuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentNestedModuleRequestWithDefaults() *ComponentNestedModuleRequest {
	this := ComponentNestedModuleRequest{}
	return &this
}

// GetDevice returns the Device field value
func (o *ComponentNestedModuleRequest) GetDevice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *ComponentNestedModuleRequest) GetDeviceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *ComponentNestedModuleRequest) SetDevice(v int32) {
	o.Device = v
}

func (o ComponentNestedModuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentNestedModuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device"] = o.Device

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ComponentNestedModuleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"device",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varComponentNestedModuleRequest := _ComponentNestedModuleRequest{}

	err = json.Unmarshal(data, &varComponentNestedModuleRequest)

	if err != nil {
		return err
	}

	*o = ComponentNestedModuleRequest(varComponentNestedModuleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComponentNestedModuleRequest struct {
	value *ComponentNestedModuleRequest
	isSet bool
}

func (v NullableComponentNestedModuleRequest) Get() *ComponentNestedModuleRequest {
	return v.value
}

func (v *NullableComponentNestedModuleRequest) Set(val *ComponentNestedModuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentNestedModuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentNestedModuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentNestedModuleRequest(val *ComponentNestedModuleRequest) *NullableComponentNestedModuleRequest {
	return &NullableComponentNestedModuleRequest{value: val, isSet: true}
}

func (v NullableComponentNestedModuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentNestedModuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
