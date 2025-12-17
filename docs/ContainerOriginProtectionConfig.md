# ContainerOriginProtectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether origin protection is enabled. Defaults to true if this config object is provided. | [optional] [default to true]
**IpAllow** | Pointer to **[]string** | List of IP addresses or CIDR ranges that can bypass origin protection for direct access (e.g., VPN IPs) | [optional] 

## Methods

### NewContainerOriginProtectionConfig

`func NewContainerOriginProtectionConfig() *ContainerOriginProtectionConfig`

NewContainerOriginProtectionConfig instantiates a new ContainerOriginProtectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerOriginProtectionConfigWithDefaults

`func NewContainerOriginProtectionConfigWithDefaults() *ContainerOriginProtectionConfig`

NewContainerOriginProtectionConfigWithDefaults instantiates a new ContainerOriginProtectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ContainerOriginProtectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ContainerOriginProtectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ContainerOriginProtectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ContainerOriginProtectionConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIpAllow

`func (o *ContainerOriginProtectionConfig) GetIpAllow() []string`

GetIpAllow returns the IpAllow field if non-nil, zero value otherwise.

### GetIpAllowOk

`func (o *ContainerOriginProtectionConfig) GetIpAllowOk() (*[]string, bool)`

GetIpAllowOk returns a tuple with the IpAllow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllow

`func (o *ContainerOriginProtectionConfig) SetIpAllow(v []string)`

SetIpAllow sets IpAllow field to given value.

### HasIpAllow

`func (o *ContainerOriginProtectionConfig) HasIpAllow() bool`

HasIpAllow returns a boolean if a field has been set.

### SetIpAllowNil

`func (o *ContainerOriginProtectionConfig) SetIpAllowNil(b bool)`

 SetIpAllowNil sets the value for IpAllow to be an explicit nil

### UnsetIpAllow
`func (o *ContainerOriginProtectionConfig) UnsetIpAllow()`

UnsetIpAllow ensures that no value is present for IpAllow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


