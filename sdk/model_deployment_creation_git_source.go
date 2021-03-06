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

// DeploymentCreationGitSource struct for DeploymentCreationGitSource
type DeploymentCreationGitSource struct {
	Type      *string `json:"type,omitempty"`
	Ref       *string `json:"ref,omitempty"`
	ProjectId *string `json:"projectId,omitempty"`
	RepoId    *int32  `json:"repoId,omitempty"`
}

// NewDeploymentCreationGitSource instantiates a new DeploymentCreationGitSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentCreationGitSource() *DeploymentCreationGitSource {
	this := DeploymentCreationGitSource{}
	return &this
}

// NewDeploymentCreationGitSourceWithDefaults instantiates a new DeploymentCreationGitSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentCreationGitSourceWithDefaults() *DeploymentCreationGitSource {
	this := DeploymentCreationGitSource{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeploymentCreationGitSource) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentCreationGitSource) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeploymentCreationGitSource) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DeploymentCreationGitSource) SetType(v string) {
	o.Type = &v
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *DeploymentCreationGitSource) GetRef() string {
	if o == nil || o.Ref == nil {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentCreationGitSource) GetRefOk() (*string, bool) {
	if o == nil || o.Ref == nil {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *DeploymentCreationGitSource) HasRef() bool {
	if o != nil && o.Ref != nil {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *DeploymentCreationGitSource) SetRef(v string) {
	o.Ref = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *DeploymentCreationGitSource) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentCreationGitSource) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DeploymentCreationGitSource) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *DeploymentCreationGitSource) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetRepoId returns the RepoId field value if set, zero value otherwise.
func (o *DeploymentCreationGitSource) GetRepoId() int32 {
	if o == nil || o.RepoId == nil {
		var ret int32
		return ret
	}
	return *o.RepoId
}

// GetRepoIdOk returns a tuple with the RepoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentCreationGitSource) GetRepoIdOk() (*int32, bool) {
	if o == nil || o.RepoId == nil {
		return nil, false
	}
	return o.RepoId, true
}

// HasRepoId returns a boolean if a field has been set.
func (o *DeploymentCreationGitSource) HasRepoId() bool {
	if o != nil && o.RepoId != nil {
		return true
	}

	return false
}

// SetRepoId gets a reference to the given int32 and assigns it to the RepoId field.
func (o *DeploymentCreationGitSource) SetRepoId(v int32) {
	o.RepoId = &v
}

func (o DeploymentCreationGitSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Ref != nil {
		toSerialize["ref"] = o.Ref
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.RepoId != nil {
		toSerialize["repoId"] = o.RepoId
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentCreationGitSource struct {
	value *DeploymentCreationGitSource
	isSet bool
}

func (v NullableDeploymentCreationGitSource) Get() *DeploymentCreationGitSource {
	return v.value
}

func (v *NullableDeploymentCreationGitSource) Set(val *DeploymentCreationGitSource) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentCreationGitSource) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentCreationGitSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentCreationGitSource(val *DeploymentCreationGitSource) *NullableDeploymentCreationGitSource {
	return &NullableDeploymentCreationGitSource{value: val, isSet: true}
}

func (v NullableDeploymentCreationGitSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentCreationGitSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
