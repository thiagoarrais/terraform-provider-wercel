/*
 * Vercel API
 *
 * Unnofficial OpenAPI description for the Vercel API
 *
 * API version: 0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// EnvironmentVariableCreation struct for EnvironmentVariableCreation
type EnvironmentVariableCreation struct {
	Type   string   `json:"type"`
	Key    string   `json:"key"`
	Value  string   `json:"value"`
	Target []string `json:"target"`
}

// NewEnvironmentVariableCreation instantiates a new EnvironmentVariableCreation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentVariableCreation(type_ string, key string, value string, target []string) *EnvironmentVariableCreation {
	this := EnvironmentVariableCreation{}
	this.Type = type_
	this.Key = key
	this.Value = value
	this.Target = target
	return &this
}

// NewEnvironmentVariableCreationWithDefaults instantiates a new EnvironmentVariableCreation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentVariableCreationWithDefaults() *EnvironmentVariableCreation {
	this := EnvironmentVariableCreation{}
	return &this
}

// GetType returns the Type field value
func (o *EnvironmentVariableCreation) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableCreation) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EnvironmentVariableCreation) SetType(v string) {
	o.Type = v
}

// GetKey returns the Key field value
func (o *EnvironmentVariableCreation) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableCreation) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *EnvironmentVariableCreation) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *EnvironmentVariableCreation) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableCreation) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *EnvironmentVariableCreation) SetValue(v string) {
	o.Value = v
}

// GetTarget returns the Target field value
func (o *EnvironmentVariableCreation) GetTarget() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableCreation) GetTargetOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *EnvironmentVariableCreation) SetTarget(v []string) {
	o.Target = v
}

func (o EnvironmentVariableCreation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentVariableCreation struct {
	value *EnvironmentVariableCreation
	isSet bool
}

func (v NullableEnvironmentVariableCreation) Get() *EnvironmentVariableCreation {
	return v.value
}

func (v *NullableEnvironmentVariableCreation) Set(val *EnvironmentVariableCreation) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentVariableCreation) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentVariableCreation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentVariableCreation(val *EnvironmentVariableCreation) *NullableEnvironmentVariableCreation {
	return &NullableEnvironmentVariableCreation{value: val, isSet: true}
}

func (v NullableEnvironmentVariableCreation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentVariableCreation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
