# OrgResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgName** | Pointer to **string** |  | [optional] 
**ResourceId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** | org for resources managed by these endpoints. app rows are compatibility records for an application-managed cache and cannot be mutated here. | [optional] 
**Config** | Pointer to **map[string]interface{}** | Type-specific settings, such as dataStorageMaxGb for a cache | [optional] 
**Physical** | Pointer to **map[string]interface{}** | Provisioned detail: bucket and region for object storage, cache identifier and endpoint for a cache | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewOrgResource

`func NewOrgResource() *OrgResource`

NewOrgResource instantiates a new OrgResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgResourceWithDefaults

`func NewOrgResourceWithDefaults() *OrgResource`

NewOrgResourceWithDefaults instantiates a new OrgResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgName

`func (o *OrgResource) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *OrgResource) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *OrgResource) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *OrgResource) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetResourceId

`func (o *OrgResource) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *OrgResource) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *OrgResource) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *OrgResource) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetType

`func (o *OrgResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrgResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrgResource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrgResource) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *OrgResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrgResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *OrgResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrgResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrgResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrgResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetScope

`func (o *OrgResource) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OrgResource) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OrgResource) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OrgResource) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetConfig

`func (o *OrgResource) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *OrgResource) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *OrgResource) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *OrgResource) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetPhysical

`func (o *OrgResource) GetPhysical() map[string]interface{}`

GetPhysical returns the Physical field if non-nil, zero value otherwise.

### GetPhysicalOk

`func (o *OrgResource) GetPhysicalOk() (*map[string]interface{}, bool)`

GetPhysicalOk returns a tuple with the Physical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysical

`func (o *OrgResource) SetPhysical(v map[string]interface{})`

SetPhysical sets Physical field to given value.

### HasPhysical

`func (o *OrgResource) HasPhysical() bool`

HasPhysical returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrgResource) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrgResource) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrgResource) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrgResource) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


