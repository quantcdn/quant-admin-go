# V2OrganizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
**Name** | **string** | Organization name | 
**MachineName** | **string** | Organization machine name | 

## Methods

### NewV2OrganizationRequest

`func NewV2OrganizationRequest(message string, error_ bool, name string, machineName string, ) *V2OrganizationRequest`

NewV2OrganizationRequest instantiates a new V2OrganizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2OrganizationRequestWithDefaults

`func NewV2OrganizationRequestWithDefaults() *V2OrganizationRequest`

NewV2OrganizationRequestWithDefaults instantiates a new V2OrganizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *V2OrganizationRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2OrganizationRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2OrganizationRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2OrganizationRequest) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2OrganizationRequest) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2OrganizationRequest) SetError(v bool)`

SetError sets Error field to given value.


### GetName

`func (o *V2OrganizationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2OrganizationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2OrganizationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetMachineName

`func (o *V2OrganizationRequest) GetMachineName() string`

GetMachineName returns the MachineName field if non-nil, zero value otherwise.

### GetMachineNameOk

`func (o *V2OrganizationRequest) GetMachineNameOk() (*string, bool)`

GetMachineNameOk returns a tuple with the MachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineName

`func (o *V2OrganizationRequest) SetMachineName(v string)`

SetMachineName sets MachineName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


