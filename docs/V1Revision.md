# V1Revision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Md5** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ByteLength** | Pointer to **int32** |  | [optional] 
**RevisionNumber** | Pointer to **int32** |  | [optional] 
**DateTimestamp** | Pointer to **int32** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DeletedTimestamp** | Pointer to **int32** |  | [optional] 
**Transitions** | Pointer to [**[]V1Transition**](V1Transition.md) |  | [optional] 
**Info** | Pointer to [**V1Info**](V1Info.md) |  | [optional] 

## Methods

### NewV1Revision

`func NewV1Revision() *V1Revision`

NewV1Revision instantiates a new V1Revision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RevisionWithDefaults

`func NewV1RevisionWithDefaults() *V1Revision`

NewV1RevisionWithDefaults instantiates a new V1Revision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMd5

`func (o *V1Revision) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *V1Revision) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *V1Revision) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *V1Revision) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetType

`func (o *V1Revision) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1Revision) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1Revision) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1Revision) HasType() bool`

HasType returns a boolean if a field has been set.

### GetByteLength

`func (o *V1Revision) GetByteLength() int32`

GetByteLength returns the ByteLength field if non-nil, zero value otherwise.

### GetByteLengthOk

`func (o *V1Revision) GetByteLengthOk() (*int32, bool)`

GetByteLengthOk returns a tuple with the ByteLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteLength

`func (o *V1Revision) SetByteLength(v int32)`

SetByteLength sets ByteLength field to given value.

### HasByteLength

`func (o *V1Revision) HasByteLength() bool`

HasByteLength returns a boolean if a field has been set.

### GetRevisionNumber

`func (o *V1Revision) GetRevisionNumber() int32`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *V1Revision) GetRevisionNumberOk() (*int32, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *V1Revision) SetRevisionNumber(v int32)`

SetRevisionNumber sets RevisionNumber field to given value.

### HasRevisionNumber

`func (o *V1Revision) HasRevisionNumber() bool`

HasRevisionNumber returns a boolean if a field has been set.

### GetDateTimestamp

`func (o *V1Revision) GetDateTimestamp() int32`

GetDateTimestamp returns the DateTimestamp field if non-nil, zero value otherwise.

### GetDateTimestampOk

`func (o *V1Revision) GetDateTimestampOk() (*int32, bool)`

GetDateTimestampOk returns a tuple with the DateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimestamp

`func (o *V1Revision) SetDateTimestamp(v int32)`

SetDateTimestamp sets DateTimestamp field to given value.

### HasDateTimestamp

`func (o *V1Revision) HasDateTimestamp() bool`

HasDateTimestamp returns a boolean if a field has been set.

### GetDeleted

`func (o *V1Revision) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *V1Revision) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *V1Revision) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *V1Revision) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDeletedTimestamp

`func (o *V1Revision) GetDeletedTimestamp() int32`

GetDeletedTimestamp returns the DeletedTimestamp field if non-nil, zero value otherwise.

### GetDeletedTimestampOk

`func (o *V1Revision) GetDeletedTimestampOk() (*int32, bool)`

GetDeletedTimestampOk returns a tuple with the DeletedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedTimestamp

`func (o *V1Revision) SetDeletedTimestamp(v int32)`

SetDeletedTimestamp sets DeletedTimestamp field to given value.

### HasDeletedTimestamp

`func (o *V1Revision) HasDeletedTimestamp() bool`

HasDeletedTimestamp returns a boolean if a field has been set.

### GetTransitions

`func (o *V1Revision) GetTransitions() []V1Transition`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *V1Revision) GetTransitionsOk() (*[]V1Transition, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *V1Revision) SetTransitions(v []V1Transition)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *V1Revision) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.

### GetInfo

`func (o *V1Revision) GetInfo() V1Info`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *V1Revision) GetInfoOk() (*V1Info, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *V1Revision) SetInfo(v V1Info)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *V1Revision) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


