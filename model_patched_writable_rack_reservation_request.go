/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.5.8 (3.5)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the PatchedWritableRackReservationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableRackReservationRequest{}

// PatchedWritableRackReservationRequest Adds support for custom fields and tags.
type PatchedWritableRackReservationRequest struct {
	Rack                 *int32                 `json:"rack,omitempty"`
	Units                []int32                `json:"units,omitempty"`
	User                 *int32                 `json:"user,omitempty"`
	Tenant               NullableInt32          `json:"tenant,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableRackReservationRequest PatchedWritableRackReservationRequest

// NewPatchedWritableRackReservationRequest instantiates a new PatchedWritableRackReservationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableRackReservationRequest() *PatchedWritableRackReservationRequest {
	this := PatchedWritableRackReservationRequest{}
	return &this
}

// NewPatchedWritableRackReservationRequestWithDefaults instantiates a new PatchedWritableRackReservationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableRackReservationRequestWithDefaults() *PatchedWritableRackReservationRequest {
	this := PatchedWritableRackReservationRequest{}
	return &this
}

// GetRack returns the Rack field value if set, zero value otherwise.
func (o *PatchedWritableRackReservationRequest) GetRack() int32 {
	if o == nil || IsNil(o.Rack) {
		var ret int32
		return ret
	}
	return *o.Rack
}

// GetRackOk returns a tuple with the Rack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableRackReservationRequest) GetRackOk() (*int32, bool) {
	if o == nil || IsNil(o.Rack) {
		return nil, false
	}
	return o.Rack, true
}

// HasRack returns a boolean if a field has been set.
func (o *PatchedWritableRackReservationRequest) HasRack() bool {
	if o != nil && !IsNil(o.Rack) {
		return true
	}

	return false
}

// SetRack gets a reference to the given int32 and assigns it to the Rack field.
func (o *PatchedWritableRackReservationRequest) SetRack(v int32) {
	o.Rack = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *PatchedWritableRackReservationRequest) GetUnits() []int32 {
	if o == nil || IsNil(o.Units) {
		var ret []int32
		return ret
	}
	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableRackReservationRequest) GetUnitsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *PatchedWritableRackReservationRequest) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given []int32 and assigns it to the Units field.
func (o *PatchedWritableRackReservationRequest) SetUnits(v []int32) {
	o.Units = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *PatchedWritableRackReservationRequest) GetUser() int32 {
	if o == nil || IsNil(o.User) {
		var ret int32
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableRackReservationRequest) GetUserOk() (*int32, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *PatchedWritableRackReservationRequest) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given int32 and assigns it to the User field.
func (o *PatchedWritableRackReservationRequest) SetUser(v int32) {
	o.User = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableRackReservationRequest) GetTenant() int32 {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret int32
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableRackReservationRequest) GetTenantOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedWritableRackReservationRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableInt32 and assigns it to the Tenant field.
func (o *PatchedWritableRackReservationRequest) SetTenant(v int32) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PatchedWritableRackReservationRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PatchedWritableRackReservationRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableRackReservationRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableRackReservationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableRackReservationRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableRackReservationRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritableRackReservationRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableRackReservationRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritableRackReservationRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritableRackReservationRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableRackReservationRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableRackReservationRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableRackReservationRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableRackReservationRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableRackReservationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableRackReservationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableRackReservationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableRackReservationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableRackReservationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableRackReservationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rack) {
		toSerialize["rack"] = o.Rack
	}
	if !IsNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedWritableRackReservationRequest) UnmarshalJSON(bytes []byte) (err error) {
	varPatchedWritableRackReservationRequest := _PatchedWritableRackReservationRequest{}

	if err = json.Unmarshal(bytes, &varPatchedWritableRackReservationRequest); err == nil {
		*o = PatchedWritableRackReservationRequest(varPatchedWritableRackReservationRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "rack")
		delete(additionalProperties, "units")
		delete(additionalProperties, "user")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableRackReservationRequest struct {
	value *PatchedWritableRackReservationRequest
	isSet bool
}

func (v NullablePatchedWritableRackReservationRequest) Get() *PatchedWritableRackReservationRequest {
	return v.value
}

func (v *NullablePatchedWritableRackReservationRequest) Set(val *PatchedWritableRackReservationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableRackReservationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableRackReservationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableRackReservationRequest(val *PatchedWritableRackReservationRequest) *NullablePatchedWritableRackReservationRequest {
	return &NullablePatchedWritableRackReservationRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableRackReservationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableRackReservationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}