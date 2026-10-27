# V2KvIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Collection** | **string** |  | 
**Fields** | **[]string** |  | 
**SortBy** | Pointer to **string** |  | [optional] 
**Unique** | Pointer to **bool** |  | [optional] 
**Count** | Pointer to **bool** |  | [optional] 
**Slot** | **int32** | The secondary index slot that serves the index. | 
**State** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Error** | Pointer to **string** | FAILED only: the code of the first problem the build found. | [optional] 
**Conflicts** | Pointer to **[]string** | FAILED only: up to 20 keys (16 KiB) that caused the failure, without duplicates. | [optional] 
**Job** | Pointer to [**V2KvIndexJob**](V2KvIndexJob.md) |  | [optional] 

## Methods

### NewV2KvIndex

`func NewV2KvIndex(name string, collection string, fields []string, slot int32, state string, createdAt time.Time, ) *V2KvIndex`

NewV2KvIndex instantiates a new V2KvIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvIndexWithDefaults

`func NewV2KvIndexWithDefaults() *V2KvIndex`

NewV2KvIndexWithDefaults instantiates a new V2KvIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2KvIndex) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2KvIndex) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2KvIndex) SetName(v string)`

SetName sets Name field to given value.


### GetCollection

`func (o *V2KvIndex) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *V2KvIndex) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *V2KvIndex) SetCollection(v string)`

SetCollection sets Collection field to given value.


### GetFields

`func (o *V2KvIndex) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *V2KvIndex) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *V2KvIndex) SetFields(v []string)`

SetFields sets Fields field to given value.


### GetSortBy

`func (o *V2KvIndex) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *V2KvIndex) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *V2KvIndex) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *V2KvIndex) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetUnique

`func (o *V2KvIndex) GetUnique() bool`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *V2KvIndex) GetUniqueOk() (*bool, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *V2KvIndex) SetUnique(v bool)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *V2KvIndex) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetCount

`func (o *V2KvIndex) GetCount() bool`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V2KvIndex) GetCountOk() (*bool, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V2KvIndex) SetCount(v bool)`

SetCount sets Count field to given value.

### HasCount

`func (o *V2KvIndex) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSlot

`func (o *V2KvIndex) GetSlot() int32`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *V2KvIndex) GetSlotOk() (*int32, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *V2KvIndex) SetSlot(v int32)`

SetSlot sets Slot field to given value.


### GetState

`func (o *V2KvIndex) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V2KvIndex) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V2KvIndex) SetState(v string)`

SetState sets State field to given value.


### GetCreatedAt

`func (o *V2KvIndex) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2KvIndex) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2KvIndex) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetError

`func (o *V2KvIndex) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2KvIndex) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2KvIndex) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V2KvIndex) HasError() bool`

HasError returns a boolean if a field has been set.

### GetConflicts

`func (o *V2KvIndex) GetConflicts() []string`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *V2KvIndex) GetConflictsOk() (*[]string, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *V2KvIndex) SetConflicts(v []string)`

SetConflicts sets Conflicts field to given value.

### HasConflicts

`func (o *V2KvIndex) HasConflicts() bool`

HasConflicts returns a boolean if a field has been set.

### GetJob

`func (o *V2KvIndex) GetJob() V2KvIndexJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *V2KvIndex) GetJobOk() (*V2KvIndexJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *V2KvIndex) SetJob(v V2KvIndexJob)`

SetJob sets Job field to given value.

### HasJob

`func (o *V2KvIndex) HasJob() bool`

HasJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


