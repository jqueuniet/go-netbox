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

// IKEProposalAuthenticationAlgorithmValue * `hmac-sha1` - SHA-1 HMAC * `hmac-sha256` - SHA-256 HMAC * `hmac-sha384` - SHA-384 HMAC * `hmac-sha512` - SHA-512 HMAC * `hmac-md5` - MD5 HMAC
type IKEProposalAuthenticationAlgorithmValue string

// List of IKEProposal_authentication_algorithm_value
const (
	IKEPROPOSALAUTHENTICATIONALGORITHMVALUE_SHA1   IKEProposalAuthenticationAlgorithmValue = "hmac-sha1"
	IKEPROPOSALAUTHENTICATIONALGORITHMVALUE_SHA256 IKEProposalAuthenticationAlgorithmValue = "hmac-sha256"
	IKEPROPOSALAUTHENTICATIONALGORITHMVALUE_SHA384 IKEProposalAuthenticationAlgorithmValue = "hmac-sha384"
	IKEPROPOSALAUTHENTICATIONALGORITHMVALUE_SHA512 IKEProposalAuthenticationAlgorithmValue = "hmac-sha512"
	IKEPROPOSALAUTHENTICATIONALGORITHMVALUE_MD5    IKEProposalAuthenticationAlgorithmValue = "hmac-md5"
)

// All allowed values of IKEProposalAuthenticationAlgorithmValue enum
var AllowedIKEProposalAuthenticationAlgorithmValueEnumValues = []IKEProposalAuthenticationAlgorithmValue{
	"hmac-sha1",
	"hmac-sha256",
	"hmac-sha384",
	"hmac-sha512",
	"hmac-md5",
}

func (v *IKEProposalAuthenticationAlgorithmValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IKEProposalAuthenticationAlgorithmValue(value)
	for _, existing := range AllowedIKEProposalAuthenticationAlgorithmValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IKEProposalAuthenticationAlgorithmValue", value)
}

// NewIKEProposalAuthenticationAlgorithmValueFromValue returns a pointer to a valid IKEProposalAuthenticationAlgorithmValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIKEProposalAuthenticationAlgorithmValueFromValue(v string) (*IKEProposalAuthenticationAlgorithmValue, error) {
	ev := IKEProposalAuthenticationAlgorithmValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IKEProposalAuthenticationAlgorithmValue: valid values are %v", v, AllowedIKEProposalAuthenticationAlgorithmValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IKEProposalAuthenticationAlgorithmValue) IsValid() bool {
	for _, existing := range AllowedIKEProposalAuthenticationAlgorithmValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IKEProposal_authentication_algorithm_value value
func (v IKEProposalAuthenticationAlgorithmValue) Ptr() *IKEProposalAuthenticationAlgorithmValue {
	return &v
}

type NullableIKEProposalAuthenticationAlgorithmValue struct {
	value *IKEProposalAuthenticationAlgorithmValue
	isSet bool
}

func (v NullableIKEProposalAuthenticationAlgorithmValue) Get() *IKEProposalAuthenticationAlgorithmValue {
	return v.value
}

func (v *NullableIKEProposalAuthenticationAlgorithmValue) Set(val *IKEProposalAuthenticationAlgorithmValue) {
	v.value = val
	v.isSet = true
}

func (v NullableIKEProposalAuthenticationAlgorithmValue) IsSet() bool {
	return v.isSet
}

func (v *NullableIKEProposalAuthenticationAlgorithmValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIKEProposalAuthenticationAlgorithmValue(val *IKEProposalAuthenticationAlgorithmValue) *NullableIKEProposalAuthenticationAlgorithmValue {
	return &NullableIKEProposalAuthenticationAlgorithmValue{value: val, isSet: true}
}

func (v NullableIKEProposalAuthenticationAlgorithmValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIKEProposalAuthenticationAlgorithmValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
