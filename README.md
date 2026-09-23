<h1 align="center">AIOZ S3 Vector Memory Go Client</h1>

`superintelligencememory` is the Go client for `s3vector_gw`'s S3 Vector memory API
(`POST /Record` / `POST /Recall`). Every request is authenticated with AWS
SigV4 (service `s3vectors`) using a static access key ID / secret access key
pair -- there is no login/token-exchange step or JWT/API-key auth mode.

## Table of contents

- [Getting started](#getting-started)
  - [Installation](#installation)
  - [Code sample](#code-sample)
- [Documentation](#documentation)
  - [API Endpoints](#api-endpoints)
    - [Memory](#memory)
  - [Models](#models)

## Getting started

### Installation
```bash
go get github.com/AIOZStorage/SuperIntelligenceMemory
```

### Code sample

```golang
package main

import (
    "context"
    "fmt"
    "net/http"
    "time"

    superintelligencememory "github.com/AIOZStorage/SuperIntelligenceMemory"
)

func main() {
    ctx := context.Background()

    apiClientOpts := superintelligencememory.
        DefaultApiClientOptions().
        BaseURL("https://vector-api.aiozstorage.network").
        HTTPClient(
            &http.Client{
                Timeout: 15 * time.Second,
            },
        ).
        Credentials("YOUR_ACCESS_KEY_ID", "YOUR_SECRET_ACCESS_KEY").
        Region("us-east-1")
    apiClient, err := superintelligencememory.NewApiClient(ctx, apiClientOpts)
    if err != nil {
        // Handle error
    }
    // or
    // apiClient := superintelligencememory.MustNewApiClient(ctx, apiClientOpts)

    out, err := apiClient.Memory.RecordWithContext(ctx, superintelligencememory.RecordRequest{
        VectorBucketName: superintelligencememory.PtrString("my-bucket"),
        Statement:        superintelligencememory.PtrString("Paris is the capital of France."),
    })
    if err != nil {
        // A non-nil err from a non-2xx response is *superintelligencememory.AgentMemoryAPIError
        // ({Type, Message}, decoded from the {__type, message} error envelope
        // every s3vector_gw error response carries).
        fmt.Println(err)
        return
    }
    fmt.Println(out)
}
```

`RecordWithContext`/`RecallWithContext` take the request struct **by
value** (`RecordRequest`, not `*RecordRequest`); `apiClient.Memory` also
exposes non-context `Record`/`Recall` convenience methods that call the
`*WithContext` variants with `context.Background()`.

## Documentation

### API Endpoints

All URIs are relative to *http://localhost*


#### Memory


##### Retrieve an instance of the Memory API:
```golang
apiClientOpts := superintelligencememory.
        DefaultApiClientOptions().
        BaseURL(S3VECTOR_BASE_URL).
        HTTPClient(
            &http.Client{
                Timeout: 15 * time.Second,
            },
        ).
        Credentials(S3VECTOR_ACCESS_KEY_ID, S3VECTOR_SECRET_ACCESS_KEY).
        Region(S3VECTOR_REGION)
client := superintelligencememory.MustNewApiClient(ctx, apiClientOpts)
memoryApi := client.Memory
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Recall**](docs/Memory.md#Recall) | **Post** `/Recall` | Recall the most relevant stored statement for a question
[**Record**](docs/Memory.md#Record) | **Post** `/Record` | Store a statement in memory




### Models

 - [APIErrorResponse](docs/APIErrorResponse.md)
 - [RecallMatch](docs/RecallMatch.md)
 - [RecallRequest](docs/RecallRequest.md)
 - [RecallResponse](docs/RecallResponse.md)
 - [RecordRequest](docs/RecordRequest.md)
 - [RecordResponse](docs/RecordResponse.md)

