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

// BriefJobStatusValue * `pending` - Pending * `scheduled` - Scheduled * `running` - Running * `completed` - Completed * `errored` - Errored * `failed` - Failed
type BriefJobStatusValue string

// List of BriefJob_status_value
const (
	BRIEFJOBSTATUSVALUE_PENDING BriefJobStatusValue = "pending"
	BRIEFJOBSTATUSVALUE_SCHEDULED BriefJobStatusValue = "scheduled"
	BRIEFJOBSTATUSVALUE_RUNNING BriefJobStatusValue = "running"
	BRIEFJOBSTATUSVALUE_COMPLETED BriefJobStatusValue = "completed"
	BRIEFJOBSTATUSVALUE_ERRORED BriefJobStatusValue = "errored"
	BRIEFJOBSTATUSVALUE_FAILED BriefJobStatusValue = "failed"
)

// All allowed values of BriefJobStatusValue enum
var AllowedBriefJobStatusValueEnumValues = []BriefJobStatusValue{
	"pending",
	"scheduled",
	"running",
	"completed",
	"errored",
	"failed",
}

func (v *BriefJobStatusValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BriefJobStatusValue(value)
	for _, existing := range AllowedBriefJobStatusValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BriefJobStatusValue", value)
}

// NewBriefJobStatusValueFromValue returns a pointer to a valid BriefJobStatusValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBriefJobStatusValueFromValue(v string) (*BriefJobStatusValue, error) {
	ev := BriefJobStatusValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BriefJobStatusValue: valid values are %v", v, AllowedBriefJobStatusValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BriefJobStatusValue) IsValid() bool {
	for _, existing := range AllowedBriefJobStatusValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BriefJob_status_value value
func (v BriefJobStatusValue) Ptr() *BriefJobStatusValue {
	return &v
}

type NullableBriefJobStatusValue struct {
	value *BriefJobStatusValue
	isSet bool
}

func (v NullableBriefJobStatusValue) Get() *BriefJobStatusValue {
	return v.value
}

func (v *NullableBriefJobStatusValue) Set(val *BriefJobStatusValue) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefJobStatusValue) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefJobStatusValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefJobStatusValue(val *BriefJobStatusValue) *NullableBriefJobStatusValue {
	return &NullableBriefJobStatusValue{value: val, isSet: true}
}

func (v NullableBriefJobStatusValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefJobStatusValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

