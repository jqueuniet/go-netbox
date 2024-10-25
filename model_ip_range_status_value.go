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

// IPRangeStatusValue * `active` - Active * `reserved` - Reserved * `deprecated` - Deprecated
type IPRangeStatusValue string

// List of IPRange_status_value
const (
	IPRANGESTATUSVALUE_ACTIVE     IPRangeStatusValue = "active"
	IPRANGESTATUSVALUE_RESERVED   IPRangeStatusValue = "reserved"
	IPRANGESTATUSVALUE_DEPRECATED IPRangeStatusValue = "deprecated"
)

// All allowed values of IPRangeStatusValue enum
var AllowedIPRangeStatusValueEnumValues = []IPRangeStatusValue{
	"active",
	"reserved",
	"deprecated",
}

func (v *IPRangeStatusValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IPRangeStatusValue(value)
	for _, existing := range AllowedIPRangeStatusValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IPRangeStatusValue", value)
}

// NewIPRangeStatusValueFromValue returns a pointer to a valid IPRangeStatusValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIPRangeStatusValueFromValue(v string) (*IPRangeStatusValue, error) {
	ev := IPRangeStatusValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IPRangeStatusValue: valid values are %v", v, AllowedIPRangeStatusValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IPRangeStatusValue) IsValid() bool {
	for _, existing := range AllowedIPRangeStatusValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IPRange_status_value value
func (v IPRangeStatusValue) Ptr() *IPRangeStatusValue {
	return &v
}

type NullableIPRangeStatusValue struct {
	value *IPRangeStatusValue
	isSet bool
}

func (v NullableIPRangeStatusValue) Get() *IPRangeStatusValue {
	return v.value
}

func (v *NullableIPRangeStatusValue) Set(val *IPRangeStatusValue) {
	v.value = val
	v.isSet = true
}

func (v NullableIPRangeStatusValue) IsSet() bool {
	return v.isSet
}

func (v *NullableIPRangeStatusValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPRangeStatusValue(val *IPRangeStatusValue) *NullableIPRangeStatusValue {
	return &NullableIPRangeStatusValue{value: val, isSet: true}
}

func (v NullableIPRangeStatusValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPRangeStatusValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
