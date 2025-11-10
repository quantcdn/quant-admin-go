# V2ProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
**Name** | **string** | Project name | 
**MachineName** | **string** | Project machine name | 
**Region** | Pointer to **string** | Project region | [optional] 
**AllowQueryParams** | Pointer to **bool** | Allow query parameters | [optional] 
**DisableRevisions** | Pointer to **bool** | Disable revisions | [optional] 
**BasicAuthUsername** | Pointer to **string** | Basic auth username | [optional] 
**BasicAuthPassword** | Pointer to **string** | Basic auth password | [optional] 

## Methods

### NewV2ProjectRequest

`func NewV2ProjectRequest(message string, error_ bool, name string, machineName string, ) *V2ProjectRequest`

NewV2ProjectRequest instantiates a new V2ProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ProjectRequestWithDefaults

`func NewV2ProjectRequestWithDefaults() *V2ProjectRequest`

NewV2ProjectRequestWithDefaults instantiates a new V2ProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *V2ProjectRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2ProjectRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2ProjectRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2ProjectRequest) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2ProjectRequest) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2ProjectRequest) SetError(v bool)`

SetError sets Error field to given value.


### GetName

`func (o *V2ProjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2ProjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2ProjectRequest) SetName(v string)`

SetName sets Name field to given value.


### GetMachineName

`func (o *V2ProjectRequest) GetMachineName() string`

GetMachineName returns the MachineName field if non-nil, zero value otherwise.

### GetMachineNameOk

`func (o *V2ProjectRequest) GetMachineNameOk() (*string, bool)`

GetMachineNameOk returns a tuple with the MachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineName

`func (o *V2ProjectRequest) SetMachineName(v string)`

SetMachineName sets MachineName field to given value.


### GetRegion

`func (o *V2ProjectRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *V2ProjectRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *V2ProjectRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *V2ProjectRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetAllowQueryParams

`func (o *V2ProjectRequest) GetAllowQueryParams() bool`

GetAllowQueryParams returns the AllowQueryParams field if non-nil, zero value otherwise.

### GetAllowQueryParamsOk

`func (o *V2ProjectRequest) GetAllowQueryParamsOk() (*bool, bool)`

GetAllowQueryParamsOk returns a tuple with the AllowQueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQueryParams

`func (o *V2ProjectRequest) SetAllowQueryParams(v bool)`

SetAllowQueryParams sets AllowQueryParams field to given value.

### HasAllowQueryParams

`func (o *V2ProjectRequest) HasAllowQueryParams() bool`

HasAllowQueryParams returns a boolean if a field has been set.

### GetDisableRevisions

`func (o *V2ProjectRequest) GetDisableRevisions() bool`

GetDisableRevisions returns the DisableRevisions field if non-nil, zero value otherwise.

### GetDisableRevisionsOk

`func (o *V2ProjectRequest) GetDisableRevisionsOk() (*bool, bool)`

GetDisableRevisionsOk returns a tuple with the DisableRevisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRevisions

`func (o *V2ProjectRequest) SetDisableRevisions(v bool)`

SetDisableRevisions sets DisableRevisions field to given value.

### HasDisableRevisions

`func (o *V2ProjectRequest) HasDisableRevisions() bool`

HasDisableRevisions returns a boolean if a field has been set.

### GetBasicAuthUsername

`func (o *V2ProjectRequest) GetBasicAuthUsername() string`

GetBasicAuthUsername returns the BasicAuthUsername field if non-nil, zero value otherwise.

### GetBasicAuthUsernameOk

`func (o *V2ProjectRequest) GetBasicAuthUsernameOk() (*string, bool)`

GetBasicAuthUsernameOk returns a tuple with the BasicAuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthUsername

`func (o *V2ProjectRequest) SetBasicAuthUsername(v string)`

SetBasicAuthUsername sets BasicAuthUsername field to given value.

### HasBasicAuthUsername

`func (o *V2ProjectRequest) HasBasicAuthUsername() bool`

HasBasicAuthUsername returns a boolean if a field has been set.

### GetBasicAuthPassword

`func (o *V2ProjectRequest) GetBasicAuthPassword() string`

GetBasicAuthPassword returns the BasicAuthPassword field if non-nil, zero value otherwise.

### GetBasicAuthPasswordOk

`func (o *V2ProjectRequest) GetBasicAuthPasswordOk() (*string, bool)`

GetBasicAuthPasswordOk returns a tuple with the BasicAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthPassword

`func (o *V2ProjectRequest) SetBasicAuthPassword(v string)`

SetBasicAuthPassword sets BasicAuthPassword field to given value.

### HasBasicAuthPassword

`func (o *V2ProjectRequest) HasBasicAuthPassword() bool`

HasBasicAuthPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


