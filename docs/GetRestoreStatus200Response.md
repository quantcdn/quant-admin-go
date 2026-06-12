# GetRestoreStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestoreId** | Pointer to **string** |  | [optional] 
**OrgName** | Pointer to **string** |  | [optional] 
**AppName** | Pointer to **string** |  | [optional] 
**EnvName** | Pointer to **string** |  | [optional] 
**BackupId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**TaskArn** | Pointer to **NullableString** |  | [optional] 
**Ttl** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewGetRestoreStatus200Response

`func NewGetRestoreStatus200Response() *GetRestoreStatus200Response`

NewGetRestoreStatus200Response instantiates a new GetRestoreStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRestoreStatus200ResponseWithDefaults

`func NewGetRestoreStatus200ResponseWithDefaults() *GetRestoreStatus200Response`

NewGetRestoreStatus200ResponseWithDefaults instantiates a new GetRestoreStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestoreId

`func (o *GetRestoreStatus200Response) GetRestoreId() string`

GetRestoreId returns the RestoreId field if non-nil, zero value otherwise.

### GetRestoreIdOk

`func (o *GetRestoreStatus200Response) GetRestoreIdOk() (*string, bool)`

GetRestoreIdOk returns a tuple with the RestoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreId

`func (o *GetRestoreStatus200Response) SetRestoreId(v string)`

SetRestoreId sets RestoreId field to given value.

### HasRestoreId

`func (o *GetRestoreStatus200Response) HasRestoreId() bool`

HasRestoreId returns a boolean if a field has been set.

### GetOrgName

`func (o *GetRestoreStatus200Response) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *GetRestoreStatus200Response) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *GetRestoreStatus200Response) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *GetRestoreStatus200Response) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetAppName

`func (o *GetRestoreStatus200Response) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *GetRestoreStatus200Response) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *GetRestoreStatus200Response) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *GetRestoreStatus200Response) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetEnvName

`func (o *GetRestoreStatus200Response) GetEnvName() string`

GetEnvName returns the EnvName field if non-nil, zero value otherwise.

### GetEnvNameOk

`func (o *GetRestoreStatus200Response) GetEnvNameOk() (*string, bool)`

GetEnvNameOk returns a tuple with the EnvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvName

`func (o *GetRestoreStatus200Response) SetEnvName(v string)`

SetEnvName sets EnvName field to given value.

### HasEnvName

`func (o *GetRestoreStatus200Response) HasEnvName() bool`

HasEnvName returns a boolean if a field has been set.

### GetBackupId

`func (o *GetRestoreStatus200Response) GetBackupId() string`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *GetRestoreStatus200Response) GetBackupIdOk() (*string, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *GetRestoreStatus200Response) SetBackupId(v string)`

SetBackupId sets BackupId field to given value.

### HasBackupId

`func (o *GetRestoreStatus200Response) HasBackupId() bool`

HasBackupId returns a boolean if a field has been set.

### GetStatus

`func (o *GetRestoreStatus200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRestoreStatus200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRestoreStatus200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetRestoreStatus200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartedAt

`func (o *GetRestoreStatus200Response) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GetRestoreStatus200Response) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GetRestoreStatus200Response) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GetRestoreStatus200Response) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *GetRestoreStatus200Response) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GetRestoreStatus200Response) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GetRestoreStatus200Response) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GetRestoreStatus200Response) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *GetRestoreStatus200Response) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *GetRestoreStatus200Response) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetErrorMessage

`func (o *GetRestoreStatus200Response) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetRestoreStatus200Response) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetRestoreStatus200Response) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetRestoreStatus200Response) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *GetRestoreStatus200Response) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *GetRestoreStatus200Response) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetTaskArn

`func (o *GetRestoreStatus200Response) GetTaskArn() string`

GetTaskArn returns the TaskArn field if non-nil, zero value otherwise.

### GetTaskArnOk

`func (o *GetRestoreStatus200Response) GetTaskArnOk() (*string, bool)`

GetTaskArnOk returns a tuple with the TaskArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskArn

`func (o *GetRestoreStatus200Response) SetTaskArn(v string)`

SetTaskArn sets TaskArn field to given value.

### HasTaskArn

`func (o *GetRestoreStatus200Response) HasTaskArn() bool`

HasTaskArn returns a boolean if a field has been set.

### SetTaskArnNil

`func (o *GetRestoreStatus200Response) SetTaskArnNil(b bool)`

 SetTaskArnNil sets the value for TaskArn to be an explicit nil

### UnsetTaskArn
`func (o *GetRestoreStatus200Response) UnsetTaskArn()`

UnsetTaskArn ensures that no value is present for TaskArn, not even an explicit nil
### GetTtl

`func (o *GetRestoreStatus200Response) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetRestoreStatus200Response) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetRestoreStatus200Response) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetRestoreStatus200Response) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### SetTtlNil

`func (o *GetRestoreStatus200Response) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *GetRestoreStatus200Response) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


