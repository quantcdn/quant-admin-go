# V2KvArchiveStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Store** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | purged: the latest purge is COMPLETE. purging: a purge is not final. blocked: the latest purge is BLOCKED. retained: a lock or the quiet period is in the future. eligible: otherwise. | [optional] 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**EarliestPurgeAt** | Pointer to **NullableTime** |  | [optional] 
**RetainedUntil** | Pointer to **NullableTime** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **NullableString** |  | [optional] 
**Purge** | Pointer to [**NullableV2KvArchivePurge**](V2KvArchivePurge.md) |  | [optional] 

## Methods

### NewV2KvArchiveStatus

`func NewV2KvArchiveStatus() *V2KvArchiveStatus`

NewV2KvArchiveStatus instantiates a new V2KvArchiveStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvArchiveStatusWithDefaults

`func NewV2KvArchiveStatusWithDefaults() *V2KvArchiveStatus`

NewV2KvArchiveStatusWithDefaults instantiates a new V2KvArchiveStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStore

`func (o *V2KvArchiveStatus) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *V2KvArchiveStatus) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *V2KvArchiveStatus) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *V2KvArchiveStatus) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetStatus

`func (o *V2KvArchiveStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V2KvArchiveStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V2KvArchiveStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V2KvArchiveStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDeletedAt

`func (o *V2KvArchiveStatus) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *V2KvArchiveStatus) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *V2KvArchiveStatus) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *V2KvArchiveStatus) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *V2KvArchiveStatus) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *V2KvArchiveStatus) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetEarliestPurgeAt

`func (o *V2KvArchiveStatus) GetEarliestPurgeAt() time.Time`

GetEarliestPurgeAt returns the EarliestPurgeAt field if non-nil, zero value otherwise.

### GetEarliestPurgeAtOk

`func (o *V2KvArchiveStatus) GetEarliestPurgeAtOk() (*time.Time, bool)`

GetEarliestPurgeAtOk returns a tuple with the EarliestPurgeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestPurgeAt

`func (o *V2KvArchiveStatus) SetEarliestPurgeAt(v time.Time)`

SetEarliestPurgeAt sets EarliestPurgeAt field to given value.

### HasEarliestPurgeAt

`func (o *V2KvArchiveStatus) HasEarliestPurgeAt() bool`

HasEarliestPurgeAt returns a boolean if a field has been set.

### SetEarliestPurgeAtNil

`func (o *V2KvArchiveStatus) SetEarliestPurgeAtNil(b bool)`

 SetEarliestPurgeAtNil sets the value for EarliestPurgeAt to be an explicit nil

### UnsetEarliestPurgeAt
`func (o *V2KvArchiveStatus) UnsetEarliestPurgeAt()`

UnsetEarliestPurgeAt ensures that no value is present for EarliestPurgeAt, not even an explicit nil
### GetRetainedUntil

`func (o *V2KvArchiveStatus) GetRetainedUntil() time.Time`

GetRetainedUntil returns the RetainedUntil field if non-nil, zero value otherwise.

### GetRetainedUntilOk

`func (o *V2KvArchiveStatus) GetRetainedUntilOk() (*time.Time, bool)`

GetRetainedUntilOk returns a tuple with the RetainedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedUntil

`func (o *V2KvArchiveStatus) SetRetainedUntil(v time.Time)`

SetRetainedUntil sets RetainedUntil field to given value.

### HasRetainedUntil

`func (o *V2KvArchiveStatus) HasRetainedUntil() bool`

HasRetainedUntil returns a boolean if a field has been set.

### SetRetainedUntilNil

`func (o *V2KvArchiveStatus) SetRetainedUntilNil(b bool)`

 SetRetainedUntilNil sets the value for RetainedUntil to be an explicit nil

### UnsetRetainedUntil
`func (o *V2KvArchiveStatus) UnsetRetainedUntil()`

UnsetRetainedUntil ensures that no value is present for RetainedUntil, not even an explicit nil
### GetSource

`func (o *V2KvArchiveStatus) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *V2KvArchiveStatus) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *V2KvArchiveStatus) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *V2KvArchiveStatus) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetProject

`func (o *V2KvArchiveStatus) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *V2KvArchiveStatus) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *V2KvArchiveStatus) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *V2KvArchiveStatus) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *V2KvArchiveStatus) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *V2KvArchiveStatus) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetPurge

`func (o *V2KvArchiveStatus) GetPurge() V2KvArchivePurge`

GetPurge returns the Purge field if non-nil, zero value otherwise.

### GetPurgeOk

`func (o *V2KvArchiveStatus) GetPurgeOk() (*V2KvArchivePurge, bool)`

GetPurgeOk returns a tuple with the Purge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurge

`func (o *V2KvArchiveStatus) SetPurge(v V2KvArchivePurge)`

SetPurge sets Purge field to given value.

### HasPurge

`func (o *V2KvArchiveStatus) HasPurge() bool`

HasPurge returns a boolean if a field has been set.

### SetPurgeNil

`func (o *V2KvArchiveStatus) SetPurgeNil(b bool)`

 SetPurgeNil sets the value for Purge to be an explicit nil

### UnsetPurge
`func (o *V2KvArchiveStatus) UnsetPurge()`

UnsetPurge ensures that no value is present for Purge, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


