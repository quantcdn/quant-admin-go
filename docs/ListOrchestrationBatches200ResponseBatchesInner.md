# ListOrchestrationBatches200ResponseBatchesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | Pointer to **string** |  | [optional] 
**OrchestrationId** | Pointer to **string** |  | [optional] 
**Iteration** | Pointer to **int32** |  | [optional] 
**ItemCount** | Pointer to **int32** |  | [optional] 
**CompletedCount** | Pointer to **int32** |  | [optional] 
**FailedCount** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListOrchestrationBatches200ResponseBatchesInner

`func NewListOrchestrationBatches200ResponseBatchesInner() *ListOrchestrationBatches200ResponseBatchesInner`

NewListOrchestrationBatches200ResponseBatchesInner instantiates a new ListOrchestrationBatches200ResponseBatchesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrchestrationBatches200ResponseBatchesInnerWithDefaults

`func NewListOrchestrationBatches200ResponseBatchesInnerWithDefaults() *ListOrchestrationBatches200ResponseBatchesInner`

NewListOrchestrationBatches200ResponseBatchesInnerWithDefaults instantiates a new ListOrchestrationBatches200ResponseBatchesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *ListOrchestrationBatches200ResponseBatchesInner) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *ListOrchestrationBatches200ResponseBatchesInner) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *ListOrchestrationBatches200ResponseBatchesInner) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *ListOrchestrationBatches200ResponseBatchesInner) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetOrchestrationId

`func (o *ListOrchestrationBatches200ResponseBatchesInner) GetOrchestrationId() string`

GetOrchestrationId returns the OrchestrationId field if non-nil, zero value otherwise.

### GetOrchestrationIdOk

`func (o *ListOrchestrationBatches200ResponseBatchesInner) GetOrchestrationIdOk() (*string, bool)`

GetOrchestrationIdOk returns a tuple with the OrchestrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestrationId

`func (o *ListOrchestrationBatches200ResponseBatchesInner) SetOrchestrationId(v string)`

SetOrchestrationId sets OrchestrationId field to given value.

### HasOrchestrationId

`func (o *ListOrchestrationBatches200ResponseBatchesInner) HasOrchestrationId() bool`

HasOrchestrationId returns a boolean if a field has been set.

### GetIteration

`func (o *ListOrchestrationBatches200ResponseBatchesInner) GetIteration() int32`

GetIteration returns the Iteration field if non-nil, zero value otherwise.

### GetIterationOk

`func (o *ListOrchestrationBatches200ResponseBatchesInner) GetIterationOk() (*int32, bool)`

GetIterationOk returns a tuple with the Iteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIteration

`func (o *ListOrchestrationBatches200ResponseBatchesInner) SetIteration(v int32)`

SetIteration sets Iteration field to given value.

### HasIteration

`func (o *ListOrchestrationBatches200ResponseBatchesInner) HasIteration() bool`

HasIteration returns a boolean if a field has been set.

### GetItemCount

`func (o *ListOrchestrationBatches200ResponseBatchesInner) GetItemCount() int32`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *ListOrchestrationBatches200ResponseBatchesInner) GetItemCountOk() (*int32, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *ListOrchestrationBatches200ResponseBatchesInner) SetItemCount(v int32)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *ListOrchestrationBatches200ResponseBatchesInner) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### GetCompletedCount

`func (o *ListOrchestrationBatches200ResponseBatchesInner) GetCompletedCount() int32`

GetCompletedCount returns the CompletedCount field if non-nil, zero value otherwise.

### GetCompletedCountOk

`func (o *ListOrchestrationBatches200ResponseBatchesInner) GetCompletedCountOk() (*int32, bool)`

GetCompletedCountOk returns a tuple with the CompletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedCount

`func (o *ListOrchestrationBatches200ResponseBatchesInner) SetCompletedCount(v int32)`

SetCompletedCount sets CompletedCount field to given value.

### HasCompletedCount

`func (o *ListOrchestrationBatches200ResponseBatchesInner) HasCompletedCount() bool`

HasCompletedCount returns a boolean if a field has been set.

### GetFailedCount

`func (o *ListOrchestrationBatches200ResponseBatchesInner) GetFailedCount() int32`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *ListOrchestrationBatches200ResponseBatchesInner) GetFailedCountOk() (*int32, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *ListOrchestrationBatches200ResponseBatchesInner) SetFailedCount(v int32)`

SetFailedCount sets FailedCount field to given value.

### HasFailedCount

`func (o *ListOrchestrationBatches200ResponseBatchesInner) HasFailedCount() bool`

HasFailedCount returns a boolean if a field has been set.

### GetStatus

`func (o *ListOrchestrationBatches200ResponseBatchesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListOrchestrationBatches200ResponseBatchesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListOrchestrationBatches200ResponseBatchesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListOrchestrationBatches200ResponseBatchesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartedAt

`func (o *ListOrchestrationBatches200ResponseBatchesInner) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ListOrchestrationBatches200ResponseBatchesInner) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ListOrchestrationBatches200ResponseBatchesInner) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ListOrchestrationBatches200ResponseBatchesInner) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *ListOrchestrationBatches200ResponseBatchesInner) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ListOrchestrationBatches200ResponseBatchesInner) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ListOrchestrationBatches200ResponseBatchesInner) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ListOrchestrationBatches200ResponseBatchesInner) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetError

`func (o *ListOrchestrationBatches200ResponseBatchesInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListOrchestrationBatches200ResponseBatchesInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListOrchestrationBatches200ResponseBatchesInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ListOrchestrationBatches200ResponseBatchesInner) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ListOrchestrationBatches200ResponseBatchesInner) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ListOrchestrationBatches200ResponseBatchesInner) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


