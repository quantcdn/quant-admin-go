# V2KvTransactResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditional** | **bool** |  | 
**Versions** | **map[string]string** | The new version of each set and increment key. | 
**Replayed** | Pointer to **bool** | True when an earlier request with the same idempotencyKey applied the operations. | [optional] 

## Methods

### NewV2KvTransactResponse

`func NewV2KvTransactResponse(conditional bool, versions map[string]string, ) *V2KvTransactResponse`

NewV2KvTransactResponse instantiates a new V2KvTransactResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvTransactResponseWithDefaults

`func NewV2KvTransactResponseWithDefaults() *V2KvTransactResponse`

NewV2KvTransactResponseWithDefaults instantiates a new V2KvTransactResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditional

`func (o *V2KvTransactResponse) GetConditional() bool`

GetConditional returns the Conditional field if non-nil, zero value otherwise.

### GetConditionalOk

`func (o *V2KvTransactResponse) GetConditionalOk() (*bool, bool)`

GetConditionalOk returns a tuple with the Conditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditional

`func (o *V2KvTransactResponse) SetConditional(v bool)`

SetConditional sets Conditional field to given value.


### GetVersions

`func (o *V2KvTransactResponse) GetVersions() map[string]string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *V2KvTransactResponse) GetVersionsOk() (*map[string]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *V2KvTransactResponse) SetVersions(v map[string]string)`

SetVersions sets Versions field to given value.


### GetReplayed

`func (o *V2KvTransactResponse) GetReplayed() bool`

GetReplayed returns the Replayed field if non-nil, zero value otherwise.

### GetReplayedOk

`func (o *V2KvTransactResponse) GetReplayedOk() (*bool, bool)`

GetReplayedOk returns a tuple with the Replayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayed

`func (o *V2KvTransactResponse) SetReplayed(v bool)`

SetReplayed sets Replayed field to given value.

### HasReplayed

`func (o *V2KvTransactResponse) HasReplayed() bool`

HasReplayed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


