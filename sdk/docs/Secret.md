# Secret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** |  | 
**Name** | **string** |  | 
**Created** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Value** | Pointer to [**SecretValue**](Secret_value.md) |  | [optional] 

## Methods

### NewSecret

`func NewSecret(uid string, name string, ) *Secret`

NewSecret instantiates a new Secret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretWithDefaults

`func NewSecretWithDefaults() *Secret`

NewSecretWithDefaults instantiates a new Secret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *Secret) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Secret) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Secret) SetUid(v string)`

SetUid sets Uid field to given value.


### GetName

`func (o *Secret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Secret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Secret) SetName(v string)`

SetName sets Name field to given value.


### GetCreated

`func (o *Secret) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Secret) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Secret) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Secret) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUserId

`func (o *Secret) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Secret) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Secret) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Secret) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetValue

`func (o *Secret) GetValue() SecretValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Secret) GetValueOk() (*SecretValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Secret) SetValue(v SecretValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *Secret) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


