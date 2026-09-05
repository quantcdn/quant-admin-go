# CreateOrgResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** | Lowercase letters, numbers and hyphens, 2-41 characters, starting with a letter or number | 
**Config** | Pointer to [**CreateOrgResourceRequestConfig**](CreateOrgResourceRequestConfig.md) |  | [optional] 

## Methods

### NewCreateOrgResourceRequest

`func NewCreateOrgResourceRequest(type_ string, name string, ) *CreateOrgResourceRequest`

NewCreateOrgResourceRequest instantiates a new CreateOrgResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrgResourceRequestWithDefaults

`func NewCreateOrgResourceRequestWithDefaults() *CreateOrgResourceRequest`

NewCreateOrgResourceRequestWithDefaults instantiates a new CreateOrgResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateOrgResourceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOrgResourceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOrgResourceRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CreateOrgResourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrgResourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrgResourceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetConfig

`func (o *CreateOrgResourceRequest) GetConfig() CreateOrgResourceRequestConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateOrgResourceRequest) GetConfigOk() (*CreateOrgResourceRequestConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateOrgResourceRequest) SetConfig(v CreateOrgResourceRequestConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateOrgResourceRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


