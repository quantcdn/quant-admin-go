# V2KvEntryPutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Value** | Pointer to **interface{}** | The stored JSON value. | [optional] 
**Status** | **string** |  | 
**Message** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**ExpiresAt** | Pointer to **int64** | Epoch seconds. Absent without ttl. | [optional] 
**Version** | Pointer to **string** | Format 2 stores only. The new version. | [optional] 
**Conditional** | Pointer to **bool** | Format 2 stores. True when the write had ifAbsent or ifVersion. | [optional] 

## Methods

### NewV2KvEntryPutResponse

`func NewV2KvEntryPutResponse(key string, status string, message string, ) *V2KvEntryPutResponse`

NewV2KvEntryPutResponse instantiates a new V2KvEntryPutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvEntryPutResponseWithDefaults

`func NewV2KvEntryPutResponseWithDefaults() *V2KvEntryPutResponse`

NewV2KvEntryPutResponseWithDefaults instantiates a new V2KvEntryPutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *V2KvEntryPutResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *V2KvEntryPutResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *V2KvEntryPutResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *V2KvEntryPutResponse) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V2KvEntryPutResponse) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V2KvEntryPutResponse) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *V2KvEntryPutResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *V2KvEntryPutResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *V2KvEntryPutResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetStatus

`func (o *V2KvEntryPutResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V2KvEntryPutResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V2KvEntryPutResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *V2KvEntryPutResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2KvEntryPutResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2KvEntryPutResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCreatedAt

`func (o *V2KvEntryPutResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2KvEntryPutResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2KvEntryPutResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V2KvEntryPutResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *V2KvEntryPutResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V2KvEntryPutResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V2KvEntryPutResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V2KvEntryPutResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *V2KvEntryPutResponse) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *V2KvEntryPutResponse) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *V2KvEntryPutResponse) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *V2KvEntryPutResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetVersion

`func (o *V2KvEntryPutResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V2KvEntryPutResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V2KvEntryPutResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V2KvEntryPutResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetConditional

`func (o *V2KvEntryPutResponse) GetConditional() bool`

GetConditional returns the Conditional field if non-nil, zero value otherwise.

### GetConditionalOk

`func (o *V2KvEntryPutResponse) GetConditionalOk() (*bool, bool)`

GetConditionalOk returns a tuple with the Conditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditional

`func (o *V2KvEntryPutResponse) SetConditional(v bool)`

SetConditional sets Conditional field to given value.

### HasConditional

`func (o *V2KvEntryPutResponse) HasConditional() bool`

HasConditional returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


