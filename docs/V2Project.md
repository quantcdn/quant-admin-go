# V2Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Project ID | 
**Uuid** | **string** | Project UUID | 
**Name** | **string** | Project name | 
**MachineName** | **string** | Project machine name | 
**WriteToken** | Pointer to **string** | Write token for API access | [optional] 
**SearchEnabled** | Pointer to **bool** | Whether the project&#39;s search index is enabled. Requires an organisation search subscription. | [optional] 
**SearchConfig** | Pointer to [**NullableV2SearchConfig**](V2SearchConfig.md) |  | [optional] 

## Methods

### NewV2Project

`func NewV2Project(id int32, uuid string, name string, machineName string, ) *V2Project`

NewV2Project instantiates a new V2Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ProjectWithDefaults

`func NewV2ProjectWithDefaults() *V2Project`

NewV2ProjectWithDefaults instantiates a new V2Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2Project) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2Project) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2Project) SetId(v int32)`

SetId sets Id field to given value.


### GetUuid

`func (o *V2Project) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *V2Project) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *V2Project) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *V2Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2Project) SetName(v string)`

SetName sets Name field to given value.


### GetMachineName

`func (o *V2Project) GetMachineName() string`

GetMachineName returns the MachineName field if non-nil, zero value otherwise.

### GetMachineNameOk

`func (o *V2Project) GetMachineNameOk() (*string, bool)`

GetMachineNameOk returns a tuple with the MachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineName

`func (o *V2Project) SetMachineName(v string)`

SetMachineName sets MachineName field to given value.


### GetWriteToken

`func (o *V2Project) GetWriteToken() string`

GetWriteToken returns the WriteToken field if non-nil, zero value otherwise.

### GetWriteTokenOk

`func (o *V2Project) GetWriteTokenOk() (*string, bool)`

GetWriteTokenOk returns a tuple with the WriteToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteToken

`func (o *V2Project) SetWriteToken(v string)`

SetWriteToken sets WriteToken field to given value.

### HasWriteToken

`func (o *V2Project) HasWriteToken() bool`

HasWriteToken returns a boolean if a field has been set.

### GetSearchEnabled

`func (o *V2Project) GetSearchEnabled() bool`

GetSearchEnabled returns the SearchEnabled field if non-nil, zero value otherwise.

### GetSearchEnabledOk

`func (o *V2Project) GetSearchEnabledOk() (*bool, bool)`

GetSearchEnabledOk returns a tuple with the SearchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchEnabled

`func (o *V2Project) SetSearchEnabled(v bool)`

SetSearchEnabled sets SearchEnabled field to given value.

### HasSearchEnabled

`func (o *V2Project) HasSearchEnabled() bool`

HasSearchEnabled returns a boolean if a field has been set.

### GetSearchConfig

`func (o *V2Project) GetSearchConfig() V2SearchConfig`

GetSearchConfig returns the SearchConfig field if non-nil, zero value otherwise.

### GetSearchConfigOk

`func (o *V2Project) GetSearchConfigOk() (*V2SearchConfig, bool)`

GetSearchConfigOk returns a tuple with the SearchConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchConfig

`func (o *V2Project) SetSearchConfig(v V2SearchConfig)`

SetSearchConfig sets SearchConfig field to given value.

### HasSearchConfig

`func (o *V2Project) HasSearchConfig() bool`

HasSearchConfig returns a boolean if a field has been set.

### SetSearchConfigNil

`func (o *V2Project) SetSearchConfigNil(b bool)`

 SetSearchConfigNil sets the value for SearchConfig to be an explicit nil

### UnsetSearchConfig
`func (o *V2Project) UnsetSearchConfig()`

UnsetSearchConfig ensures that no value is present for SearchConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


