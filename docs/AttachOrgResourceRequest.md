# AttachOrgResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | **string** |  | 
**Environment** | **string** |  | 
**EnvVarPrefix** | Pointer to **string** | Namespaces every injected variable, so MEDIA yields MEDIA_S3_BUCKET | [optional] 
**AccessLevel** | Pointer to **string** | Cache only. scoped injects an RBAC user limited to this environment&#39;s CACHE_PREFIX (plain and {hash-tag} forms) with FLUSHALL and FLUSHDB denied. admin injects the cache-wide credential for integrations that require FLUSHDB, such as Laravel Cache::flush() or the WordPress object cache without selective flush; it can read, write and flush every attached environment&#39;s keys. | [optional] [default to "scoped"]

## Methods

### NewAttachOrgResourceRequest

`func NewAttachOrgResourceRequest(application string, environment string, ) *AttachOrgResourceRequest`

NewAttachOrgResourceRequest instantiates a new AttachOrgResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachOrgResourceRequestWithDefaults

`func NewAttachOrgResourceRequestWithDefaults() *AttachOrgResourceRequest`

NewAttachOrgResourceRequestWithDefaults instantiates a new AttachOrgResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *AttachOrgResourceRequest) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *AttachOrgResourceRequest) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *AttachOrgResourceRequest) SetApplication(v string)`

SetApplication sets Application field to given value.


### GetEnvironment

`func (o *AttachOrgResourceRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AttachOrgResourceRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AttachOrgResourceRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetEnvVarPrefix

`func (o *AttachOrgResourceRequest) GetEnvVarPrefix() string`

GetEnvVarPrefix returns the EnvVarPrefix field if non-nil, zero value otherwise.

### GetEnvVarPrefixOk

`func (o *AttachOrgResourceRequest) GetEnvVarPrefixOk() (*string, bool)`

GetEnvVarPrefixOk returns a tuple with the EnvVarPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVarPrefix

`func (o *AttachOrgResourceRequest) SetEnvVarPrefix(v string)`

SetEnvVarPrefix sets EnvVarPrefix field to given value.

### HasEnvVarPrefix

`func (o *AttachOrgResourceRequest) HasEnvVarPrefix() bool`

HasEnvVarPrefix returns a boolean if a field has been set.

### GetAccessLevel

`func (o *AttachOrgResourceRequest) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *AttachOrgResourceRequest) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *AttachOrgResourceRequest) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *AttachOrgResourceRequest) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


