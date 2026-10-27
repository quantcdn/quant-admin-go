# V2KvUsageTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ItemCount** | **int32** |  | 
**SizeBytes** | **int32** | Table data only. | 
**IndexSizeBytes** | **int32** | Sum of the secondary index sizes (0 when none). | 

## Methods

### NewV2KvUsageTable

`func NewV2KvUsageTable(name string, itemCount int32, sizeBytes int32, indexSizeBytes int32, ) *V2KvUsageTable`

NewV2KvUsageTable instantiates a new V2KvUsageTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvUsageTableWithDefaults

`func NewV2KvUsageTableWithDefaults() *V2KvUsageTable`

NewV2KvUsageTableWithDefaults instantiates a new V2KvUsageTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2KvUsageTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2KvUsageTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2KvUsageTable) SetName(v string)`

SetName sets Name field to given value.


### GetItemCount

`func (o *V2KvUsageTable) GetItemCount() int32`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *V2KvUsageTable) GetItemCountOk() (*int32, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *V2KvUsageTable) SetItemCount(v int32)`

SetItemCount sets ItemCount field to given value.


### GetSizeBytes

`func (o *V2KvUsageTable) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *V2KvUsageTable) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *V2KvUsageTable) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.


### GetIndexSizeBytes

`func (o *V2KvUsageTable) GetIndexSizeBytes() int32`

GetIndexSizeBytes returns the IndexSizeBytes field if non-nil, zero value otherwise.

### GetIndexSizeBytesOk

`func (o *V2KvUsageTable) GetIndexSizeBytesOk() (*int32, bool)`

GetIndexSizeBytesOk returns a tuple with the IndexSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexSizeBytes

`func (o *V2KvUsageTable) SetIndexSizeBytes(v int32)`

SetIndexSizeBytes sets IndexSizeBytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


