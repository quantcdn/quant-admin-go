# PurgeOrgResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | **string** |  | 
**Application** | Pointer to **string** | scope environment only | [optional] 
**Environment** | Pointer to **string** | scope environment only | [optional] 
**Confirm** | Pointer to **bool** | scope all only; must be true | [optional] 
**Cursor** | Pointer to **string** | scope environment only; resume a partial purge | [optional] 

## Methods

### NewPurgeOrgResourceRequest

`func NewPurgeOrgResourceRequest(scope string, ) *PurgeOrgResourceRequest`

NewPurgeOrgResourceRequest instantiates a new PurgeOrgResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurgeOrgResourceRequestWithDefaults

`func NewPurgeOrgResourceRequestWithDefaults() *PurgeOrgResourceRequest`

NewPurgeOrgResourceRequestWithDefaults instantiates a new PurgeOrgResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *PurgeOrgResourceRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PurgeOrgResourceRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PurgeOrgResourceRequest) SetScope(v string)`

SetScope sets Scope field to given value.


### GetApplication

`func (o *PurgeOrgResourceRequest) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *PurgeOrgResourceRequest) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *PurgeOrgResourceRequest) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *PurgeOrgResourceRequest) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetEnvironment

`func (o *PurgeOrgResourceRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PurgeOrgResourceRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PurgeOrgResourceRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PurgeOrgResourceRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetConfirm

`func (o *PurgeOrgResourceRequest) GetConfirm() bool`

GetConfirm returns the Confirm field if non-nil, zero value otherwise.

### GetConfirmOk

`func (o *PurgeOrgResourceRequest) GetConfirmOk() (*bool, bool)`

GetConfirmOk returns a tuple with the Confirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirm

`func (o *PurgeOrgResourceRequest) SetConfirm(v bool)`

SetConfirm sets Confirm field to given value.

### HasConfirm

`func (o *PurgeOrgResourceRequest) HasConfirm() bool`

HasConfirm returns a boolean if a field has been set.

### GetCursor

`func (o *PurgeOrgResourceRequest) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *PurgeOrgResourceRequest) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *PurgeOrgResourceRequest) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *PurgeOrgResourceRequest) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


