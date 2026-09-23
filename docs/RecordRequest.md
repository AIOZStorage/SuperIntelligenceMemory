# RecordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statement** | Pointer to **string** |  | [optional] 
**VectorBucketName** | Pointer to **string** |  | [optional] 

## Methods

### NewRecordRequest

`func NewRecordRequest() *RecordRequest`

NewRecordRequest instantiates a new RecordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordRequestWithDefaults

`func NewRecordRequestWithDefaults() *RecordRequest`

NewRecordRequestWithDefaults instantiates a new RecordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatement

`func (o *RecordRequest) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *RecordRequest) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *RecordRequest) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *RecordRequest) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetVectorBucketName

`func (o *RecordRequest) GetVectorBucketName() string`

GetVectorBucketName returns the VectorBucketName field if non-nil, zero value otherwise.

### GetVectorBucketNameOk

`func (o *RecordRequest) GetVectorBucketNameOk() (*string, bool)`

GetVectorBucketNameOk returns a tuple with the VectorBucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorBucketName

`func (o *RecordRequest) SetVectorBucketName(v string)`

SetVectorBucketName sets VectorBucketName field to given value.

### HasVectorBucketName

`func (o *RecordRequest) HasVectorBucketName() bool`

HasVectorBucketName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


