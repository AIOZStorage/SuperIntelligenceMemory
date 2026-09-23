# Memory

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Recall**](Memory.md#Recall) | **Post** /Recall | Recall the most relevant stored statement for a question
[**Record**](Memory.md#Record) | **Post** /Record | Store a statement in memory



## Recall

> Recall(recallRequest RecallRequest) (*RecallResponse, error)

> RecallWithContext(ctx context.Context, recallRequest RecallRequest) (*RecallResponse, error)


Recall the most relevant stored statement for a question



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    superintelligencememory "github.com/AIOZStorage/SuperIntelligenceMemory"
)

func main() {
    apiClientOpts := superintelligencememory.
        DefaultApiClientOptions().
        BaseURL(YOUR_API_BASE_URL).
        HTTPClient(
            &http.Client{
                Timeout: YOUR_TIMEOUT,
            },
        ).
        Credentials(YOUR_ACCESS_KEY_ID, YOUR_SECRET_ACCESS_KEY).
        Region(YOUR_REGION)
    client := superintelligencememory.MustNewApiClient(ctx, apiClientOpts)
        
    recallRequest := *superintelligencememory.NewRecallRequest() // RecallRequest | Question to recall against

    
    res, err := client.Memory.Recall(recallRequest)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Memory.Recall``: %v\n", err)
    }
    // response from `Recall`: RecallResponse
    fmt.Fprintf(os.Stdout, "Response from `Memory.Recall`: %v\n", res)
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recallRequest** | [**RecallRequest**](RecallRequest.md) | Question to recall against | 

### Return type

[**RecallResponse**](RecallResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Record

> Record(recordRequest RecordRequest) (*RecordResponse, error)

> RecordWithContext(ctx context.Context, recordRequest RecordRequest) (*RecordResponse, error)


Store a statement in memory



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    superintelligencememory "github.com/AIOZStorage/SuperIntelligenceMemory"
)

func main() {
    apiClientOpts := superintelligencememory.
        DefaultApiClientOptions().
        BaseURL(YOUR_API_BASE_URL).
        HTTPClient(
            &http.Client{
                Timeout: YOUR_TIMEOUT,
            },
        ).
        Credentials(YOUR_ACCESS_KEY_ID, YOUR_SECRET_ACCESS_KEY).
        Region(YOUR_REGION)
    client := superintelligencememory.MustNewApiClient(ctx, apiClientOpts)
        
    recordRequest := *superintelligencememory.NewRecordRequest() // RecordRequest | Statement to record

    
    res, err := client.Memory.Record(recordRequest)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Memory.Record``: %v\n", err)
    }
    // response from `Record`: RecordResponse
    fmt.Fprintf(os.Stdout, "Response from `Memory.Record`: %v\n", res)
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRequest** | [**RecordRequest**](RecordRequest.md) | Statement to record | 

### Return type

[**RecordResponse**](RecordResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

