# RecallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** |  | [optional] 
**Question** | Pointer to **string** |  | [optional] 
**ReturnDistance** | Pointer to **bool** |  | [optional] 
**VectorBucketName** | Pointer to **string** |  | [optional] 

## Methods

### NewRecallRequest

`func NewRecallRequest() *RecallRequest`

NewRecallRequest instantiates a new RecallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecallRequestWithDefaults

`func NewRecallRequestWithDefaults() *RecallRequest`

NewRecallRequestWithDefaults instantiates a new RecallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *RecallRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RecallRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RecallRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *RecallRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetQuestion

`func (o *RecallRequest) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *RecallRequest) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *RecallRequest) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *RecallRequest) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetReturnDistance

`func (o *RecallRequest) GetReturnDistance() bool`

GetReturnDistance returns the ReturnDistance field if non-nil, zero value otherwise.

### GetReturnDistanceOk

`func (o *RecallRequest) GetReturnDistanceOk() (*bool, bool)`

GetReturnDistanceOk returns a tuple with the ReturnDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnDistance

`func (o *RecallRequest) SetReturnDistance(v bool)`

SetReturnDistance sets ReturnDistance field to given value.

### HasReturnDistance

`func (o *RecallRequest) HasReturnDistance() bool`

HasReturnDistance returns a boolean if a field has been set.

### GetVectorBucketName

`func (o *RecallRequest) GetVectorBucketName() string`

GetVectorBucketName returns the VectorBucketName field if non-nil, zero value otherwise.

### GetVectorBucketNameOk

`func (o *RecallRequest) GetVectorBucketNameOk() (*string, bool)`

GetVectorBucketNameOk returns a tuple with the VectorBucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorBucketName

`func (o *RecallRequest) SetVectorBucketName(v string)`

SetVectorBucketName sets VectorBucketName field to given value.

### HasVectorBucketName

`func (o *RecallRequest) HasVectorBucketName() bool`

HasVectorBucketName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


