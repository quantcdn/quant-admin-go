# ListOrchestrationBatches200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Batches** | Pointer to [**[]ListOrchestrationBatches200ResponseBatchesInner**](ListOrchestrationBatches200ResponseBatchesInner.md) |  | [optional] 
**NextCursor** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListOrchestrationBatches200Response

`func NewListOrchestrationBatches200Response() *ListOrchestrationBatches200Response`

NewListOrchestrationBatches200Response instantiates a new ListOrchestrationBatches200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrchestrationBatches200ResponseWithDefaults

`func NewListOrchestrationBatches200ResponseWithDefaults() *ListOrchestrationBatches200Response`

NewListOrchestrationBatches200ResponseWithDefaults instantiates a new ListOrchestrationBatches200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatches

`func (o *ListOrchestrationBatches200Response) GetBatches() []ListOrchestrationBatches200ResponseBatchesInner`

GetBatches returns the Batches field if non-nil, zero value otherwise.

### GetBatchesOk

`func (o *ListOrchestrationBatches200Response) GetBatchesOk() (*[]ListOrchestrationBatches200ResponseBatchesInner, bool)`

GetBatchesOk returns a tuple with the Batches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatches

`func (o *ListOrchestrationBatches200Response) SetBatches(v []ListOrchestrationBatches200ResponseBatchesInner)`

SetBatches sets Batches field to given value.

### HasBatches

`func (o *ListOrchestrationBatches200Response) HasBatches() bool`

HasBatches returns a boolean if a field has been set.

### GetNextCursor

`func (o *ListOrchestrationBatches200Response) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ListOrchestrationBatches200Response) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ListOrchestrationBatches200Response) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ListOrchestrationBatches200Response) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *ListOrchestrationBatches200Response) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *ListOrchestrationBatches200Response) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


