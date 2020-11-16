# ProjectCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**GitRepository** | Pointer to [**GitRepositoryLink**](GitRepositoryLink.md) |  | [optional] 

## Methods

### NewProjectCreation

`func NewProjectCreation(name string, ) *ProjectCreation`

NewProjectCreation instantiates a new ProjectCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCreationWithDefaults

`func NewProjectCreationWithDefaults() *ProjectCreation`

NewProjectCreationWithDefaults instantiates a new ProjectCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectCreation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectCreation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectCreation) SetName(v string)`

SetName sets Name field to given value.


### GetGitRepository

`func (o *ProjectCreation) GetGitRepository() GitRepositoryLink`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *ProjectCreation) GetGitRepositoryOk() (*GitRepositoryLink, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *ProjectCreation) SetGitRepository(v GitRepositoryLink)`

SetGitRepository sets GitRepository field to given value.

### HasGitRepository

`func (o *ProjectCreation) HasGitRepository() bool`

HasGitRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


