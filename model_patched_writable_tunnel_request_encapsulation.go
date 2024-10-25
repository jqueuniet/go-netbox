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

// PatchedWritableTunnelRequestEncapsulation * `ipsec-transport` - IPsec - Transport * `ipsec-tunnel` - IPsec - Tunnel * `ip-ip` - IP-in-IP * `gre` - GRE
type PatchedWritableTunnelRequestEncapsulation string

// List of PatchedWritableTunnelRequest_encapsulation
const (
	PATCHEDWRITABLETUNNELREQUESTENCAPSULATION_IPSEC_TRANSPORT PatchedWritableTunnelRequestEncapsulation = "ipsec-transport"
	PATCHEDWRITABLETUNNELREQUESTENCAPSULATION_IPSEC_TUNNEL    PatchedWritableTunnelRequestEncapsulation = "ipsec-tunnel"
	PATCHEDWRITABLETUNNELREQUESTENCAPSULATION_IP_IP           PatchedWritableTunnelRequestEncapsulation = "ip-ip"
	PATCHEDWRITABLETUNNELREQUESTENCAPSULATION_GRE             PatchedWritableTunnelRequestEncapsulation = "gre"
)

// All allowed values of PatchedWritableTunnelRequestEncapsulation enum
var AllowedPatchedWritableTunnelRequestEncapsulationEnumValues = []PatchedWritableTunnelRequestEncapsulation{
	"ipsec-transport",
	"ipsec-tunnel",
	"ip-ip",
	"gre",
}

func (v *PatchedWritableTunnelRequestEncapsulation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableTunnelRequestEncapsulation(value)
	for _, existing := range AllowedPatchedWritableTunnelRequestEncapsulationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableTunnelRequestEncapsulation", value)
}

// NewPatchedWritableTunnelRequestEncapsulationFromValue returns a pointer to a valid PatchedWritableTunnelRequestEncapsulation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableTunnelRequestEncapsulationFromValue(v string) (*PatchedWritableTunnelRequestEncapsulation, error) {
	ev := PatchedWritableTunnelRequestEncapsulation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableTunnelRequestEncapsulation: valid values are %v", v, AllowedPatchedWritableTunnelRequestEncapsulationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableTunnelRequestEncapsulation) IsValid() bool {
	for _, existing := range AllowedPatchedWritableTunnelRequestEncapsulationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableTunnelRequest_encapsulation value
func (v PatchedWritableTunnelRequestEncapsulation) Ptr() *PatchedWritableTunnelRequestEncapsulation {
	return &v
}

type NullablePatchedWritableTunnelRequestEncapsulation struct {
	value *PatchedWritableTunnelRequestEncapsulation
	isSet bool
}

func (v NullablePatchedWritableTunnelRequestEncapsulation) Get() *PatchedWritableTunnelRequestEncapsulation {
	return v.value
}

func (v *NullablePatchedWritableTunnelRequestEncapsulation) Set(val *PatchedWritableTunnelRequestEncapsulation) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableTunnelRequestEncapsulation) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableTunnelRequestEncapsulation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableTunnelRequestEncapsulation(val *PatchedWritableTunnelRequestEncapsulation) *NullablePatchedWritableTunnelRequestEncapsulation {
	return &NullablePatchedWritableTunnelRequestEncapsulation{value: val, isSet: true}
}

func (v NullablePatchedWritableTunnelRequestEncapsulation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableTunnelRequestEncapsulation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
