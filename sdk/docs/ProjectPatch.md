# ProjectPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**RootDirectory** | **NullableString** |  | 

## Methods

### NewProjectPatch

`func NewProjectPatch(rootDirectory NullableString, ) *ProjectPatch`

NewProjectPatch instantiates a new ProjectPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectPatchWithDefaults

`func NewProjectPatchWithDefaults() *ProjectPatch`

NewProjectPatchWithDefaults instantiates a new ProjectPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRootDirectory

`func (o *ProjectPatch) GetRootDirectory() string`

GetRootDirectory returns the RootDirectory field if non-nil, zero value otherwise.

### GetRootDirectoryOk

`func (o *ProjectPatch) GetRootDirectoryOk() (*string, bool)`

GetRootDirectoryOk returns a tuple with the RootDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDirectory

`func (o *ProjectPatch) SetRootDirectory(v string)`

SetRootDirectory sets RootDirectory field to given value.


### SetRootDirectoryNil

`func (o *ProjectPatch) SetRootDirectoryNil(b bool)`

 SetRootDirectoryNil sets the value for RootDirectory to be an explicit nil

### UnsetRootDirectory
`func (o *ProjectPatch) UnsetRootDirectory()`

UnsetRootDirectory ensures that no value is present for RootDirectory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


