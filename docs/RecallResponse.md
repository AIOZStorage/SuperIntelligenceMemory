# RecallResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | Pointer to **bool** |  | [optional] 
**Results** | Pointer to [**[]RecallMatch**](RecallMatch.md) |  | [optional] 

## Methods

### NewRecallResponse

`func NewRecallResponse() *RecallResponse`

NewRecallResponse instantiates a new RecallResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecallResponseWithDefaults

`func NewRecallResponseWithDefaults() *RecallResponse`

NewRecallResponseWithDefaults instantiates a new RecallResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *RecallResponse) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *RecallResponse) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *RecallResponse) SetFound(v bool)`

SetFound sets Found field to given value.

### HasFound

`func (o *RecallResponse) HasFound() bool`

HasFound returns a boolean if a field has been set.

### GetResults

`func (o *RecallResponse) GetResults() []RecallMatch`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RecallResponse) GetResultsOk() (*[]RecallMatch, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RecallResponse) SetResults(v []RecallMatch)`

SetResults sets Results field to given value.

### HasResults

`func (o *RecallResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


