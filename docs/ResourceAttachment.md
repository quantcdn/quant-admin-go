# ResourceAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | Pointer to **string** |  | [optional] 
**EnvName** | Pointer to **string** |  | [optional] 
**EnvVarPrefix** | Pointer to **string** | Namespaces every injected variable, so MEDIA yields MEDIA_S3_BUCKET | [optional] 
**AccessKeyId** | Pointer to **string** | Object storage only. The secret half is written to the environment&#39;s secrets and never returned. | [optional] 
**CacheUserId** | Pointer to **string** | Cache only. This environment&#39;s own RBAC user, limited to its CACHE_PREFIX with FLUSHALL and FLUSHDB denied, so it cannot touch another environment&#39;s keys. | [optional] 
**AccessLevel** | Pointer to **string** | Cache only. scoped: the environment holds its own RBAC user. admin: it holds the cache-wide credential and can read, write and flush every attached environment&#39;s keys. Absent on attachments made before access levels existed (treated as scoped). | [optional] 
**InjectedKeys** | Pointer to **[]string** | The exact variable names this attachment wrote, removed precisely on detach | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Note** | Pointer to **string** | When the credentials take effect | [optional] 

## Methods

### NewResourceAttachment

`func NewResourceAttachment() *ResourceAttachment`

NewResourceAttachment instantiates a new ResourceAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAttachmentWithDefaults

`func NewResourceAttachmentWithDefaults() *ResourceAttachment`

NewResourceAttachmentWithDefaults instantiates a new ResourceAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *ResourceAttachment) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ResourceAttachment) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ResourceAttachment) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *ResourceAttachment) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetEnvName

`func (o *ResourceAttachment) GetEnvName() string`

GetEnvName returns the EnvName field if non-nil, zero value otherwise.

### GetEnvNameOk

`func (o *ResourceAttachment) GetEnvNameOk() (*string, bool)`

GetEnvNameOk returns a tuple with the EnvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvName

`func (o *ResourceAttachment) SetEnvName(v string)`

SetEnvName sets EnvName field to given value.

### HasEnvName

`func (o *ResourceAttachment) HasEnvName() bool`

HasEnvName returns a boolean if a field has been set.

### GetEnvVarPrefix

`func (o *ResourceAttachment) GetEnvVarPrefix() string`

GetEnvVarPrefix returns the EnvVarPrefix field if non-nil, zero value otherwise.

### GetEnvVarPrefixOk

`func (o *ResourceAttachment) GetEnvVarPrefixOk() (*string, bool)`

GetEnvVarPrefixOk returns a tuple with the EnvVarPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVarPrefix

`func (o *ResourceAttachment) SetEnvVarPrefix(v string)`

SetEnvVarPrefix sets EnvVarPrefix field to given value.

### HasEnvVarPrefix

`func (o *ResourceAttachment) HasEnvVarPrefix() bool`

HasEnvVarPrefix returns a boolean if a field has been set.

### GetAccessKeyId

`func (o *ResourceAttachment) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *ResourceAttachment) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *ResourceAttachment) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *ResourceAttachment) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetCacheUserId

`func (o *ResourceAttachment) GetCacheUserId() string`

GetCacheUserId returns the CacheUserId field if non-nil, zero value otherwise.

### GetCacheUserIdOk

`func (o *ResourceAttachment) GetCacheUserIdOk() (*string, bool)`

GetCacheUserIdOk returns a tuple with the CacheUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheUserId

`func (o *ResourceAttachment) SetCacheUserId(v string)`

SetCacheUserId sets CacheUserId field to given value.

### HasCacheUserId

`func (o *ResourceAttachment) HasCacheUserId() bool`

HasCacheUserId returns a boolean if a field has been set.

### GetAccessLevel

`func (o *ResourceAttachment) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *ResourceAttachment) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *ResourceAttachment) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *ResourceAttachment) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetInjectedKeys

`func (o *ResourceAttachment) GetInjectedKeys() []string`

GetInjectedKeys returns the InjectedKeys field if non-nil, zero value otherwise.

### GetInjectedKeysOk

`func (o *ResourceAttachment) GetInjectedKeysOk() (*[]string, bool)`

GetInjectedKeysOk returns a tuple with the InjectedKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectedKeys

`func (o *ResourceAttachment) SetInjectedKeys(v []string)`

SetInjectedKeys sets InjectedKeys field to given value.

### HasInjectedKeys

`func (o *ResourceAttachment) HasInjectedKeys() bool`

HasInjectedKeys returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ResourceAttachment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResourceAttachment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResourceAttachment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ResourceAttachment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetNote

`func (o *ResourceAttachment) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ResourceAttachment) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ResourceAttachment) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ResourceAttachment) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


