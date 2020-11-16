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

// ProjectCreation struct for ProjectCreation
type ProjectCreation struct {
	Name          string             `json:"name"`
	GitRepository *GitRepositoryLink `json:"gitRepository,omitempty"`
}

// NewProjectCreation instantiates a new ProjectCreation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectCreation(name string) *ProjectCreation {
	this := ProjectCreation{}
	this.Name = name
	return &this
}

// NewProjectCreationWithDefaults instantiates a new ProjectCreation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectCreationWithDefaults() *ProjectCreation {
	this := ProjectCreation{}
	return &this
}

// GetName returns the Name field value
func (o *ProjectCreation) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectCreation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectCreation) SetName(v string) {
	o.Name = v
}

// GetGitRepository returns the GitRepository field value if set, zero value otherwise.
func (o *ProjectCreation) GetGitRepository() GitRepositoryLink {
	if o == nil || o.GitRepository == nil {
		var ret GitRepositoryLink
		return ret
	}
	return *o.GitRepository
}

// GetGitRepositoryOk returns a tuple with the GitRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreation) GetGitRepositoryOk() (*GitRepositoryLink, bool) {
	if o == nil || o.GitRepository == nil {
		return nil, false
	}
	return o.GitRepository, true
}

// HasGitRepository returns a boolean if a field has been set.
func (o *ProjectCreation) HasGitRepository() bool {
	if o != nil && o.GitRepository != nil {
		return true
	}

	return false
}

// SetGitRepository gets a reference to the given GitRepositoryLink and assigns it to the GitRepository field.
func (o *ProjectCreation) SetGitRepository(v GitRepositoryLink) {
	o.GitRepository = &v
}

func (o ProjectCreation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.GitRepository != nil {
		toSerialize["gitRepository"] = o.GitRepository
	}
	return json.Marshal(toSerialize)
}

type NullableProjectCreation struct {
	value *ProjectCreation
	isSet bool
}

func (v NullableProjectCreation) Get() *ProjectCreation {
	return v.value
}

func (v *NullableProjectCreation) Set(val *ProjectCreation) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectCreation) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectCreation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectCreation(val *ProjectCreation) *NullableProjectCreation {
	return &NullableProjectCreation{value: val, isSet: true}
}

func (v NullableProjectCreation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectCreation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
