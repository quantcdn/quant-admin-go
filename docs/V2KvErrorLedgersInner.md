# V2KvErrorLedgersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The ledger name. | 
**State** | **string** | The ledger state (PENDING, ACTIVE or DROPPED). | 
**Reason** | **string** | not_dropped: drop the ledger first (after its retention ends). not_drained: the ledger is DROPPED but its last entries are not archived yet. | 

## Methods

### NewV2KvErrorLedgersInner

`func NewV2KvErrorLedgersInner(name string, state string, reason string, ) *V2KvErrorLedgersInner`

NewV2KvErrorLedgersInner instantiates a new V2KvErrorLedgersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvErrorLedgersInnerWithDefaults

`func NewV2KvErrorLedgersInnerWithDefaults() *V2KvErrorLedgersInner`

NewV2KvErrorLedgersInnerWithDefaults instantiates a new V2KvErrorLedgersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2KvErrorLedgersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2KvErrorLedgersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2KvErrorLedgersInner) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *V2KvErrorLedgersInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V2KvErrorLedgersInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V2KvErrorLedgersInner) SetState(v string)`

SetState sets State field to given value.


### GetReason

`func (o *V2KvErrorLedgersInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V2KvErrorLedgersInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V2KvErrorLedgersInner) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


