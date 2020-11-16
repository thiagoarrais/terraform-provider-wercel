# \DeploymentsApi

All URIs are relative to *https://api.vercel.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewDeployment**](DeploymentsApi.md#CreateNewDeployment) | **Post** /v13/now/deployments | 



## CreateNewDeployment

> CreateNewDeployment(ctx).DeploymentCreation(deploymentCreation).Execute()



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
    deploymentCreation := *openapiclient.NewDeploymentCreation("Name_example") // DeploymentCreation |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.CreateNewDeployment(context.Background()).DeploymentCreation(deploymentCreation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.CreateNewDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentCreation** | [**DeploymentCreation**](DeploymentCreation.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

