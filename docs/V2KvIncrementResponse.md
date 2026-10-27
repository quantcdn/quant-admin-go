# V2KvIncrementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Value** | **int64** | The value after this increment. | 
**Operation** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**ExpiresAt** | Pointer to **int64** | Epoch seconds. | [optional] 
**Version** | Pointer to **string** | Format 2 stores only. | [optional] 

## Methods

### NewV2KvIncrementResponse

`func NewV2KvIncrementResponse(key string, value int64, operation string, updatedAt time.Time, ) *V2KvIncrementResponse`

NewV2KvIncrementResponse instantiates a new V2KvIncrementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvIncrementResponseWithDefaults

`func NewV2KvIncrementResponseWithDefaults() *V2KvIncrementResponse`

NewV2KvIncrementResponseWithDefaults instantiates a new V2KvIncrementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *V2KvIncrementResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *V2KvIncrementResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *V2KvIncrementResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *V2KvIncrementResponse) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V2KvIncrementResponse) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V2KvIncrementResponse) SetValue(v int64)`

SetValue sets Value field to given value.


### GetOperation

`func (o *V2KvIncrementResponse) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *V2KvIncrementResponse) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *V2KvIncrementResponse) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetUpdatedAt

`func (o *V2KvIncrementResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V2KvIncrementResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V2KvIncrementResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetExpiresAt

`func (o *V2KvIncrementResponse) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *V2KvIncrementResponse) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *V2KvIncrementResponse) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *V2KvIncrementResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetVersion

`func (o *V2KvIncrementResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V2KvIncrementResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V2KvIncrementResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V2KvIncrementResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


