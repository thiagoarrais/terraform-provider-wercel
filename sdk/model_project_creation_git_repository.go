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

// ProjectCreationGitRepository struct for ProjectCreationGitRepository
type ProjectCreationGitRepository struct {
	Type       string `json:"type"`
	Repo       string `json:"repo"`
	Sourceless *bool  `json:"sourceless,omitempty"`
}

// NewProjectCreationGitRepository instantiates a new ProjectCreationGitRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectCreationGitRepository(type_ string, repo string) *ProjectCreationGitRepository {
	this := ProjectCreationGitRepository{}
	this.Type = type_
	this.Repo = repo
	return &this
}

// NewProjectCreationGitRepositoryWithDefaults instantiates a new ProjectCreationGitRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectCreationGitRepositoryWithDefaults() *ProjectCreationGitRepository {
	this := ProjectCreationGitRepository{}
	return &this
}

// GetType returns the Type field value
func (o *ProjectCreationGitRepository) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProjectCreationGitRepository) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ProjectCreationGitRepository) SetType(v string) {
	o.Type = v
}

// GetRepo returns the Repo field value
func (o *ProjectCreationGitRepository) GetRepo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repo
}

// GetRepoOk returns a tuple with the Repo field value
// and a boolean to check if the value has been set.
func (o *ProjectCreationGitRepository) GetRepoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repo, true
}

// SetRepo sets field value
func (o *ProjectCreationGitRepository) SetRepo(v string) {
	o.Repo = v
}

// GetSourceless returns the Sourceless field value if set, zero value otherwise.
func (o *ProjectCreationGitRepository) GetSourceless() bool {
	if o == nil || o.Sourceless == nil {
		var ret bool
		return ret
	}
	return *o.Sourceless
}

// GetSourcelessOk returns a tuple with the Sourceless field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreationGitRepository) GetSourcelessOk() (*bool, bool) {
	if o == nil || o.Sourceless == nil {
		return nil, false
	}
	return o.Sourceless, true
}

// HasSourceless returns a boolean if a field has been set.
func (o *ProjectCreationGitRepository) HasSourceless() bool {
	if o != nil && o.Sourceless != nil {
		return true
	}

	return false
}

// SetSourceless gets a reference to the given bool and assigns it to the Sourceless field.
func (o *ProjectCreationGitRepository) SetSourceless(v bool) {
	o.Sourceless = &v
}

func (o ProjectCreationGitRepository) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["repo"] = o.Repo
	}
	if o.Sourceless != nil {
		toSerialize["sourceless"] = o.Sourceless
	}
	return json.Marshal(toSerialize)
}

type NullableProjectCreationGitRepository struct {
	value *ProjectCreationGitRepository
	isSet bool
}

func (v NullableProjectCreationGitRepository) Get() *ProjectCreationGitRepository {
	return v.value
}

func (v *NullableProjectCreationGitRepository) Set(val *ProjectCreationGitRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectCreationGitRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectCreationGitRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectCreationGitRepository(val *ProjectCreationGitRepository) *NullableProjectCreationGitRepository {
	return &NullableProjectCreationGitRepository{value: val, isSet: true}
}

func (v NullableProjectCreationGitRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectCreationGitRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}