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

// Project struct for Project
type Project struct {
	Id            string          `json:"id"`
	Name          string          `json:"name"`
	Link          *ProjectLink    `json:"link,omitempty"`
	Alias         *[]ProjectAlias `json:"alias,omitempty"`
	RootDirectory *string         `json:"rootDirectory,omitempty"`
}

// NewProject instantiates a new Project object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProject(id string, name string) *Project {
	this := Project{}
	this.Id = id
	this.Name = name
	return &this
}

// NewProjectWithDefaults instantiates a new Project object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectWithDefaults() *Project {
	this := Project{}
	return &this
}

// GetId returns the Id field value
func (o *Project) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Project) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Project) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Project) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Project) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Project) SetName(v string) {
	o.Name = v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *Project) GetLink() ProjectLink {
	if o == nil || o.Link == nil {
		var ret ProjectLink
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetLinkOk() (*ProjectLink, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *Project) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given ProjectLink and assigns it to the Link field.
func (o *Project) SetLink(v ProjectLink) {
	o.Link = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *Project) GetAlias() []ProjectAlias {
	if o == nil || o.Alias == nil {
		var ret []ProjectAlias
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetAliasOk() (*[]ProjectAlias, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *Project) HasAlias() bool {
	if o != nil && o.Alias != nil {
		return true
	}

	return false
}

// SetAlias gets a reference to the given []ProjectAlias and assigns it to the Alias field.
func (o *Project) SetAlias(v []ProjectAlias) {
	o.Alias = &v
}

// GetRootDirectory returns the RootDirectory field value if set, zero value otherwise.
func (o *Project) GetRootDirectory() string {
	if o == nil || o.RootDirectory == nil {
		var ret string
		return ret
	}
	return *o.RootDirectory
}

// GetRootDirectoryOk returns a tuple with the RootDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetRootDirectoryOk() (*string, bool) {
	if o == nil || o.RootDirectory == nil {
		return nil, false
	}
	return o.RootDirectory, true
}

// HasRootDirectory returns a boolean if a field has been set.
func (o *Project) HasRootDirectory() bool {
	if o != nil && o.RootDirectory != nil {
		return true
	}

	return false
}

// SetRootDirectory gets a reference to the given string and assigns it to the RootDirectory field.
func (o *Project) SetRootDirectory(v string) {
	o.RootDirectory = &v
}

func (o Project) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if o.RootDirectory != nil {
		toSerialize["rootDirectory"] = o.RootDirectory
	}
	return json.Marshal(toSerialize)
}

type NullableProject struct {
	value *Project
	isSet bool
}

func (v NullableProject) Get() *Project {
	return v.value
}

func (v *NullableProject) Set(val *Project) {
	v.value = val
	v.isSet = true
}

func (v NullableProject) IsSet() bool {
	return v.isSet
}

func (v *NullableProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProject(val *Project) *NullableProject {
	return &NullableProject{value: val, isSet: true}
}

func (v NullableProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
