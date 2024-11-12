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

// checks if the BriefTenantRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefTenantRequest{}

// BriefTenantRequest Adds support for custom fields and tags.
type BriefTenantRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefTenantRequest BriefTenantRequest

// NewBriefTenantRequest instantiates a new BriefTenantRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefTenantRequest(name string, slug string) *BriefTenantRequest {
	this := BriefTenantRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewBriefTenantRequestWithDefaults instantiates a new BriefTenantRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefTenantRequestWithDefaults() *BriefTenantRequest {
	this := BriefTenantRequest{}
	return &this
}

// GetName returns the Name field value
func (o *BriefTenantRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefTenantRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefTenantRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *BriefTenantRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *BriefTenantRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *BriefTenantRequest) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefTenantRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefTenantRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefTenantRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefTenantRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BriefTenantRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefTenantRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefTenantRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"slug",
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

	varBriefTenantRequest := _BriefTenantRequest{}

	err = json.Unmarshal(data, &varBriefTenantRequest)

	if err != nil {
		return err
	}

	*o = BriefTenantRequest(varBriefTenantRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefTenantRequest struct {
	value *BriefTenantRequest
	isSet bool
}

func (v NullableBriefTenantRequest) Get() *BriefTenantRequest {
	return v.value
}

func (v *NullableBriefTenantRequest) Set(val *BriefTenantRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefTenantRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefTenantRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefTenantRequest(val *BriefTenantRequest) *NullableBriefTenantRequest {
	return &NullableBriefTenantRequest{value: val, isSet: true}
}

func (v NullableBriefTenantRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefTenantRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

