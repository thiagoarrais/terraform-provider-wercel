# GitRepositoryLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Repo** | **string** |  | 
**Sourceless** | Pointer to **bool** |  | [optional] 

## Methods

### NewGitRepositoryLink

`func NewGitRepositoryLink(type_ string, repo string, ) *GitRepositoryLink`

NewGitRepositoryLink instantiates a new GitRepositoryLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitRepositoryLinkWithDefaults

`func NewGitRepositoryLinkWithDefaults() *GitRepositoryLink`

NewGitRepositoryLinkWithDefaults instantiates a new GitRepositoryLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GitRepositoryLink) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GitRepositoryLink) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GitRepositoryLink) SetType(v string)`

SetType sets Type field to given value.


### GetRepo

`func (o *GitRepositoryLink) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *GitRepositoryLink) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *GitRepositoryLink) SetRepo(v string)`

SetRepo sets Repo field to given value.


### GetSourceless

`func (o *GitRepositoryLink) GetSourceless() bool`

GetSourceless returns the Sourceless field if non-nil, zero value otherwise.

### GetSourcelessOk

`func (o *GitRepositoryLink) GetSourcelessOk() (*bool, bool)`

GetSourcelessOk returns a tuple with the Sourceless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceless

`func (o *GitRepositoryLink) SetSourceless(v bool)`

SetSourceless sets Sourceless field to given value.

### HasSourceless

`func (o *GitRepositoryLink) HasSourceless() bool`

HasSourceless returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


