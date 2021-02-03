# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Link** | Pointer to [**ProjectLink**](Project_link.md) |  | [optional] 
**Alias** | Pointer to [**[]ProjectAlias**](ProjectAlias.md) |  | [optional] 
**RootDirectory** | Pointer to **string** |  | [optional] 

## Methods

### NewProject

`func NewProject(id string, name string, ) *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Project) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Project) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Project) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Project) SetName(v string)`

SetName sets Name field to given value.


### GetLink

`func (o *Project) GetLink() ProjectLink`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *Project) GetLinkOk() (*ProjectLink, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *Project) SetLink(v ProjectLink)`

SetLink sets Link field to given value.

### HasLink

`func (o *Project) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetAlias

`func (o *Project) GetAlias() []ProjectAlias`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *Project) GetAliasOk() (*[]ProjectAlias, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *Project) SetAlias(v []ProjectAlias)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *Project) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetRootDirectory

`func (o *Project) GetRootDirectory() string`

GetRootDirectory returns the RootDirectory field if non-nil, zero value otherwise.

### GetRootDirectoryOk

`func (o *Project) GetRootDirectoryOk() (*string, bool)`

GetRootDirectoryOk returns a tuple with the RootDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDirectory

`func (o *Project) SetRootDirectory(v string)`

SetRootDirectory sets RootDirectory field to given value.

### HasRootDirectory

`func (o *Project) HasRootDirectory() bool`

HasRootDirectory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


