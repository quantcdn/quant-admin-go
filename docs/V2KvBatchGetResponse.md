# V2KvBatchGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]V2KvEntry**](V2KvEntry.md) |  | 

## Methods

### NewV2KvBatchGetResponse

`func NewV2KvBatchGetResponse(items []*V2KvEntry, ) *V2KvBatchGetResponse`

NewV2KvBatchGetResponse instantiates a new V2KvBatchGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvBatchGetResponseWithDefaults

`func NewV2KvBatchGetResponseWithDefaults() *V2KvBatchGetResponse`

NewV2KvBatchGetResponseWithDefaults instantiates a new V2KvBatchGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *V2KvBatchGetResponse) GetItems() []*V2KvEntry`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *V2KvBatchGetResponse) GetItemsOk() (*[]*V2KvEntry, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *V2KvBatchGetResponse) SetItems(v []*V2KvEntry)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


