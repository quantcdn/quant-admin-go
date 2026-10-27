# V2KvEntryDeleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Status** | **string** |  | 
**Message** | **string** |  | 
**Conditional** | Pointer to **bool** | Format 2 stores. True when the delete had if_version. | [optional] 

## Methods

### NewV2KvEntryDeleteResponse

`func NewV2KvEntryDeleteResponse(key string, status string, message string, ) *V2KvEntryDeleteResponse`

NewV2KvEntryDeleteResponse instantiates a new V2KvEntryDeleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvEntryDeleteResponseWithDefaults

`func NewV2KvEntryDeleteResponseWithDefaults() *V2KvEntryDeleteResponse`

NewV2KvEntryDeleteResponseWithDefaults instantiates a new V2KvEntryDeleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *V2KvEntryDeleteResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *V2KvEntryDeleteResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *V2KvEntryDeleteResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetStatus

`func (o *V2KvEntryDeleteResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V2KvEntryDeleteResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V2KvEntryDeleteResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *V2KvEntryDeleteResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2KvEntryDeleteResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2KvEntryDeleteResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetConditional

`func (o *V2KvEntryDeleteResponse) GetConditional() bool`

GetConditional returns the Conditional field if non-nil, zero value otherwise.

### GetConditionalOk

`func (o *V2KvEntryDeleteResponse) GetConditionalOk() (*bool, bool)`

GetConditionalOk returns a tuple with the Conditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditional

`func (o *V2KvEntryDeleteResponse) SetConditional(v bool)`

SetConditional sets Conditional field to given value.

### HasConditional

`func (o *V2KvEntryDeleteResponse) HasConditional() bool`

HasConditional returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


