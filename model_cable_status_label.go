/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// CableStatusLabel the model 'CableStatusLabel'
type CableStatusLabel string

// List of Cable_status_label
const (
	CABLESTATUSLABEL_CONNECTED CableStatusLabel = "Connected"
	CABLESTATUSLABEL_PLANNED CableStatusLabel = "Planned"
	CABLESTATUSLABEL_DECOMMISSIONING CableStatusLabel = "Decommissioning"
)

// All allowed values of CableStatusLabel enum
var AllowedCableStatusLabelEnumValues = []CableStatusLabel{
	"Connected",
	"Planned",
	"Decommissioning",
}

func (v *CableStatusLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CableStatusLabel(value)
	for _, existing := range AllowedCableStatusLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CableStatusLabel", value)
}

// NewCableStatusLabelFromValue returns a pointer to a valid CableStatusLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCableStatusLabelFromValue(v string) (*CableStatusLabel, error) {
	ev := CableStatusLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CableStatusLabel: valid values are %v", v, AllowedCableStatusLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CableStatusLabel) IsValid() bool {
	for _, existing := range AllowedCableStatusLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Cable_status_label value
func (v CableStatusLabel) Ptr() *CableStatusLabel {
	return &v
}

type NullableCableStatusLabel struct {
	value *CableStatusLabel
	isSet bool
}

func (v NullableCableStatusLabel) Get() *CableStatusLabel {
	return v.value
}

func (v *NullableCableStatusLabel) Set(val *CableStatusLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableCableStatusLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableCableStatusLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCableStatusLabel(val *CableStatusLabel) *NullableCableStatusLabel {
	return &NullableCableStatusLabel{value: val, isSet: true}
}

func (v NullableCableStatusLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCableStatusLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

