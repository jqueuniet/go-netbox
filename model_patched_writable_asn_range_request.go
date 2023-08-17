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

// checks if the PatchedWritableASNRangeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableASNRangeRequest{}

// PatchedWritableASNRangeRequest Adds support for custom fields and tags.
type PatchedWritableASNRangeRequest struct {
	Name                 *string                `json:"name,omitempty"`
	Slug                 *string                `json:"slug,omitempty"`
	Rir                  *int32                 `json:"rir,omitempty"`
	Start                *int64                 `json:"start,omitempty"`
	End                  *int64                 `json:"end,omitempty"`
	Tenant               NullableInt32          `json:"tenant,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableASNRangeRequest PatchedWritableASNRangeRequest

// NewPatchedWritableASNRangeRequest instantiates a new PatchedWritableASNRangeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableASNRangeRequest() *PatchedWritableASNRangeRequest {
	this := PatchedWritableASNRangeRequest{}
	return &this
}

// NewPatchedWritableASNRangeRequestWithDefaults instantiates a new PatchedWritableASNRangeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableASNRangeRequestWithDefaults() *PatchedWritableASNRangeRequest {
	this := PatchedWritableASNRangeRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableASNRangeRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableASNRangeRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableASNRangeRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableASNRangeRequest) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *PatchedWritableASNRangeRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableASNRangeRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *PatchedWritableASNRangeRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *PatchedWritableASNRangeRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetRir returns the Rir field value if set, zero value otherwise.
func (o *PatchedWritableASNRangeRequest) GetRir() int32 {
	if o == nil || IsNil(o.Rir) {
		var ret int32
		return ret
	}
	return *o.Rir
}

// GetRirOk returns a tuple with the Rir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableASNRangeRequest) GetRirOk() (*int32, bool) {
	if o == nil || IsNil(o.Rir) {
		return nil, false
	}
	return o.Rir, true
}

// HasRir returns a boolean if a field has been set.
func (o *PatchedWritableASNRangeRequest) HasRir() bool {
	if o != nil && !IsNil(o.Rir) {
		return true
	}

	return false
}

// SetRir gets a reference to the given int32 and assigns it to the Rir field.
func (o *PatchedWritableASNRangeRequest) SetRir(v int32) {
	o.Rir = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *PatchedWritableASNRangeRequest) GetStart() int64 {
	if o == nil || IsNil(o.Start) {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableASNRangeRequest) GetStartOk() (*int64, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *PatchedWritableASNRangeRequest) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *PatchedWritableASNRangeRequest) SetStart(v int64) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *PatchedWritableASNRangeRequest) GetEnd() int64 {
	if o == nil || IsNil(o.End) {
		var ret int64
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableASNRangeRequest) GetEndOk() (*int64, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *PatchedWritableASNRangeRequest) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int64 and assigns it to the End field.
func (o *PatchedWritableASNRangeRequest) SetEnd(v int64) {
	o.End = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableASNRangeRequest) GetTenant() int32 {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret int32
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableASNRangeRequest) GetTenantOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedWritableASNRangeRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableInt32 and assigns it to the Tenant field.
func (o *PatchedWritableASNRangeRequest) SetTenant(v int32) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PatchedWritableASNRangeRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PatchedWritableASNRangeRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableASNRangeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableASNRangeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableASNRangeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableASNRangeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableASNRangeRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableASNRangeRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableASNRangeRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableASNRangeRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableASNRangeRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableASNRangeRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableASNRangeRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableASNRangeRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableASNRangeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableASNRangeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Rir) {
		toSerialize["rir"] = o.Rir
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
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

func (o *PatchedWritableASNRangeRequest) UnmarshalJSON(bytes []byte) (err error) {
	varPatchedWritableASNRangeRequest := _PatchedWritableASNRangeRequest{}

	if err = json.Unmarshal(bytes, &varPatchedWritableASNRangeRequest); err == nil {
		*o = PatchedWritableASNRangeRequest(varPatchedWritableASNRangeRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "rir")
		delete(additionalProperties, "start")
		delete(additionalProperties, "end")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableASNRangeRequest struct {
	value *PatchedWritableASNRangeRequest
	isSet bool
}

func (v NullablePatchedWritableASNRangeRequest) Get() *PatchedWritableASNRangeRequest {
	return v.value
}

func (v *NullablePatchedWritableASNRangeRequest) Set(val *PatchedWritableASNRangeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableASNRangeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableASNRangeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableASNRangeRequest(val *PatchedWritableASNRangeRequest) *NullablePatchedWritableASNRangeRequest {
	return &NullablePatchedWritableASNRangeRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableASNRangeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableASNRangeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}