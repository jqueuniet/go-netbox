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

// checks if the NestedModuleType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedModuleType{}

// NestedModuleType Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedModuleType struct {
	Id                   int32              `json:"id"`
	Url                  string             `json:"url"`
	Display              string             `json:"display"`
	Manufacturer         NestedManufacturer `json:"manufacturer"`
	Model                string             `json:"model"`
	AdditionalProperties map[string]interface{}
}

type _NestedModuleType NestedModuleType

// NewNestedModuleType instantiates a new NestedModuleType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedModuleType(id int32, url string, display string, manufacturer NestedManufacturer, model string) *NestedModuleType {
	this := NestedModuleType{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Manufacturer = manufacturer
	this.Model = model
	return &this
}

// NewNestedModuleTypeWithDefaults instantiates a new NestedModuleType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedModuleTypeWithDefaults() *NestedModuleType {
	this := NestedModuleType{}
	return &this
}

// GetId returns the Id field value
func (o *NestedModuleType) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedModuleType) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedModuleType) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedModuleType) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedModuleType) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedModuleType) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedModuleType) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedModuleType) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedModuleType) SetDisplay(v string) {
	o.Display = v
}

// GetManufacturer returns the Manufacturer field value
func (o *NestedModuleType) GetManufacturer() NestedManufacturer {
	if o == nil {
		var ret NestedManufacturer
		return ret
	}

	return o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value
// and a boolean to check if the value has been set.
func (o *NestedModuleType) GetManufacturerOk() (*NestedManufacturer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manufacturer, true
}

// SetManufacturer sets field value
func (o *NestedModuleType) SetManufacturer(v NestedManufacturer) {
	o.Manufacturer = v
}

// GetModel returns the Model field value
func (o *NestedModuleType) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *NestedModuleType) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *NestedModuleType) SetModel(v string) {
	o.Model = v
}

func (o NestedModuleType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedModuleType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["manufacturer"] = o.Manufacturer
	toSerialize["model"] = o.Model

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedModuleType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"manufacturer",
		"model",
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

	varNestedModuleType := _NestedModuleType{}

	err = json.Unmarshal(data, &varNestedModuleType)

	if err != nil {
		return err
	}

	*o = NestedModuleType(varNestedModuleType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "model")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedModuleType struct {
	value *NestedModuleType
	isSet bool
}

func (v NullableNestedModuleType) Get() *NestedModuleType {
	return v.value
}

func (v *NullableNestedModuleType) Set(val *NestedModuleType) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedModuleType) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedModuleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedModuleType(val *NestedModuleType) *NullableNestedModuleType {
	return &NullableNestedModuleType{value: val, isSet: true}
}

func (v NullableNestedModuleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedModuleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
