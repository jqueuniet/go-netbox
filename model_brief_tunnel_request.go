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

// checks if the BriefTunnelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefTunnelRequest{}

// BriefTunnelRequest Adds support for custom fields and tags.
type BriefTunnelRequest struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefTunnelRequest BriefTunnelRequest

// NewBriefTunnelRequest instantiates a new BriefTunnelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefTunnelRequest(name string) *BriefTunnelRequest {
	this := BriefTunnelRequest{}
	this.Name = name
	return &this
}

// NewBriefTunnelRequestWithDefaults instantiates a new BriefTunnelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefTunnelRequestWithDefaults() *BriefTunnelRequest {
	this := BriefTunnelRequest{}
	return &this
}

// GetName returns the Name field value
func (o *BriefTunnelRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefTunnelRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefTunnelRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefTunnelRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefTunnelRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefTunnelRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefTunnelRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BriefTunnelRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefTunnelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefTunnelRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBriefTunnelRequest := _BriefTunnelRequest{}

	err = json.Unmarshal(data, &varBriefTunnelRequest)

	if err != nil {
		return err
	}

	*o = BriefTunnelRequest(varBriefTunnelRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefTunnelRequest struct {
	value *BriefTunnelRequest
	isSet bool
}

func (v NullableBriefTunnelRequest) Get() *BriefTunnelRequest {
	return v.value
}

func (v *NullableBriefTunnelRequest) Set(val *BriefTunnelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefTunnelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefTunnelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefTunnelRequest(val *BriefTunnelRequest) *NullableBriefTunnelRequest {
	return &NullableBriefTunnelRequest{value: val, isSet: true}
}

func (v NullableBriefTunnelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefTunnelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


