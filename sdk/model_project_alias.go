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

// ProjectAlias struct for ProjectAlias
type ProjectAlias struct {
	Domain string `json:"domain"`
}

// NewProjectAlias instantiates a new ProjectAlias object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectAlias(domain string) *ProjectAlias {
	this := ProjectAlias{}
	this.Domain = domain
	return &this
}

// NewProjectAliasWithDefaults instantiates a new ProjectAlias object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectAliasWithDefaults() *ProjectAlias {
	this := ProjectAlias{}
	return &this
}

// GetDomain returns the Domain field value
func (o *ProjectAlias) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *ProjectAlias) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *ProjectAlias) SetDomain(v string) {
	o.Domain = v
}

func (o ProjectAlias) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["domain"] = o.Domain
	}
	return json.Marshal(toSerialize)
}

type NullableProjectAlias struct {
	value *ProjectAlias
	isSet bool
}

func (v NullableProjectAlias) Get() *ProjectAlias {
	return v.value
}

func (v *NullableProjectAlias) Set(val *ProjectAlias) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectAlias) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectAlias) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectAlias(val *ProjectAlias) *NullableProjectAlias {
	return &NullableProjectAlias{value: val, isSet: true}
}

func (v NullableProjectAlias) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectAlias) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
