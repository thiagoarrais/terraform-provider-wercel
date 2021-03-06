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

// EnvironmentVariable struct for EnvironmentVariable
type EnvironmentVariable struct {
	Type      string                     `json:"type"`
	Id        string                     `json:"id"`
	Key       string                     `json:"key"`
	Value     *string                    `json:"value,omitempty"`
	Target    *EnvironmentVariableTarget `json:"target,omitempty"`
	CreatedAt *int64                     `json:"createdAt,omitempty"`
	UpdatedAt *int64                     `json:"updatedAt,omitempty"`
}

// NewEnvironmentVariable instantiates a new EnvironmentVariable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentVariable(type_ string, id string, key string) *EnvironmentVariable {
	this := EnvironmentVariable{}
	this.Type = type_
	this.Id = id
	this.Key = key
	return &this
}

// NewEnvironmentVariableWithDefaults instantiates a new EnvironmentVariable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentVariableWithDefaults() *EnvironmentVariable {
	this := EnvironmentVariable{}
	return &this
}

// GetType returns the Type field value
func (o *EnvironmentVariable) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EnvironmentVariable) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *EnvironmentVariable) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EnvironmentVariable) SetId(v string) {
	o.Id = v
}

// GetKey returns the Key field value
func (o *EnvironmentVariable) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *EnvironmentVariable) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EnvironmentVariable) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EnvironmentVariable) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EnvironmentVariable) SetValue(v string) {
	o.Value = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *EnvironmentVariable) GetTarget() EnvironmentVariableTarget {
	if o == nil || o.Target == nil {
		var ret EnvironmentVariableTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetTargetOk() (*EnvironmentVariableTarget, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *EnvironmentVariable) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given EnvironmentVariableTarget and assigns it to the Target field.
func (o *EnvironmentVariable) SetTarget(v EnvironmentVariableTarget) {
	o.Target = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EnvironmentVariable) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EnvironmentVariable) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *EnvironmentVariable) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EnvironmentVariable) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EnvironmentVariable) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *EnvironmentVariable) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

func (o EnvironmentVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentVariable struct {
	value *EnvironmentVariable
	isSet bool
}

func (v NullableEnvironmentVariable) Get() *EnvironmentVariable {
	return v.value
}

func (v *NullableEnvironmentVariable) Set(val *EnvironmentVariable) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentVariable) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentVariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentVariable(val *EnvironmentVariable) *NullableEnvironmentVariable {
	return &NullableEnvironmentVariable{value: val, isSet: true}
}

func (v NullableEnvironmentVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentVariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
