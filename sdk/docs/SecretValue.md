# SecretValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewSecretValue

`func NewSecretValue() *SecretValue`

NewSecretValue instantiates a new SecretValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretValueWithDefaults

`func NewSecretValueWithDefaults() *SecretValue`

NewSecretValueWithDefaults instantiates a new SecretValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SecretValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecretValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecretValue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SecretValue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetData

`func (o *SecretValue) GetData() []int32`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SecretValue) GetDataOk() (*[]int32, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SecretValue) SetData(v []int32)`

SetData sets Data field to given value.

### HasData

`func (o *SecretValue) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


