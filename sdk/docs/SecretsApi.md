# \SecretsApi

All URIs are relative to *https://api.vercel.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewSecret**](SecretsApi.md#CreateNewSecret) | **Post** /v2/now/secrets | 
[**RemoveSecret**](SecretsApi.md#RemoveSecret) | **Delete** /v2/now/secrets/{name} | 



## CreateNewSecret

> Secret CreateNewSecret(ctx).SecretCreation(secretCreation).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    secretCreation := *openapiclient.NewSecretCreation("Name_example", "Value_example") // SecretCreation |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecretsApi.CreateNewSecret(context.Background()).SecretCreation(secretCreation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.CreateNewSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewSecret`: Secret
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.CreateNewSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **secretCreation** | [**SecretCreation**](SecretCreation.md) |  | 

### Return type

[**Secret**](Secret.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSecret

> Secret RemoveSecret(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecretsApi.RemoveSecret(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.RemoveSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveSecret`: Secret
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.RemoveSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Secret**](Secret.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

