# DeploymentCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Target** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**GitSource** | Pointer to [**DeploymentCreationGitSource**](DeploymentCreation_gitSource.md) |  | [optional] 

## Methods

### NewDeploymentCreation

`func NewDeploymentCreation(name string, ) *DeploymentCreation`

NewDeploymentCreation instantiates a new DeploymentCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentCreationWithDefaults

`func NewDeploymentCreationWithDefaults() *DeploymentCreation`

NewDeploymentCreationWithDefaults instantiates a new DeploymentCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeploymentCreation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentCreation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentCreation) SetName(v string)`

SetName sets Name field to given value.


### GetTarget

`func (o *DeploymentCreation) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DeploymentCreation) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DeploymentCreation) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *DeploymentCreation) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetSource

`func (o *DeploymentCreation) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DeploymentCreation) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DeploymentCreation) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DeploymentCreation) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetGitSource

`func (o *DeploymentCreation) GetGitSource() DeploymentCreationGitSource`

GetGitSource returns the GitSource field if non-nil, zero value otherwise.

### GetGitSourceOk

`func (o *DeploymentCreation) GetGitSourceOk() (*DeploymentCreationGitSource, bool)`

GetGitSourceOk returns a tuple with the GitSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitSource

`func (o *DeploymentCreation) SetGitSource(v DeploymentCreationGitSource)`

SetGitSource sets GitSource field to given value.

### HasGitSource

`func (o *DeploymentCreation) HasGitSource() bool`

HasGitSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


