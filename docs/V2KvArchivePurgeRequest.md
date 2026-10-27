# V2KvArchivePurgeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | true: report only. false: start the purge. | [optional] [default to true]
**Confirm** | **string** | The store name, exactly. | 
**Reason** | Pointer to **NullableString** | Optional note kept with the purge and the receipt. No control characters. It is never echoed in an error message. | [optional] 

## Methods

### NewV2KvArchivePurgeRequest

`func NewV2KvArchivePurgeRequest(confirm string, ) *V2KvArchivePurgeRequest`

NewV2KvArchivePurgeRequest instantiates a new V2KvArchivePurgeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvArchivePurgeRequestWithDefaults

`func NewV2KvArchivePurgeRequestWithDefaults() *V2KvArchivePurgeRequest`

NewV2KvArchivePurgeRequestWithDefaults instantiates a new V2KvArchivePurgeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *V2KvArchivePurgeRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *V2KvArchivePurgeRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *V2KvArchivePurgeRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *V2KvArchivePurgeRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetConfirm

`func (o *V2KvArchivePurgeRequest) GetConfirm() string`

GetConfirm returns the Confirm field if non-nil, zero value otherwise.

### GetConfirmOk

`func (o *V2KvArchivePurgeRequest) GetConfirmOk() (*string, bool)`

GetConfirmOk returns a tuple with the Confirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirm

`func (o *V2KvArchivePurgeRequest) SetConfirm(v string)`

SetConfirm sets Confirm field to given value.


### GetReason

`func (o *V2KvArchivePurgeRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V2KvArchivePurgeRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V2KvArchivePurgeRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V2KvArchivePurgeRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *V2KvArchivePurgeRequest) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *V2KvArchivePurgeRequest) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


