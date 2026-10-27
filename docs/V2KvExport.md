# V2KvExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ULID of the export. | 
**Phase** | **string** |  | 
**RequestedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** | Moves on every step. An export in progress whose updatedAt is older than 960 seconds can be resumed. | 
**ExportTime** | Pointer to **time.Time** | The point in time the export reads. | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**ItemCount** | Pointer to **int32** |  | [optional] 
**Skipped** | Pointer to [**V2KvExportSkipped**](V2KvExportSkipped.md) |  | [optional] 
**SizeBytes** | Pointer to **int32** | Total size of the compressed files. | [optional] 
**Parts** | Pointer to **int32** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Downloads** | Pointer to [**[]V2KvExportDownload**](V2KvExportDownload.md) | Phase done only. | [optional] 

## Methods

### NewV2KvExport

`func NewV2KvExport(id string, phase string, requestedAt time.Time, updatedAt time.Time, ) *V2KvExport`

NewV2KvExport instantiates a new V2KvExport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvExportWithDefaults

`func NewV2KvExportWithDefaults() *V2KvExport`

NewV2KvExportWithDefaults instantiates a new V2KvExport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2KvExport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2KvExport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2KvExport) SetId(v string)`

SetId sets Id field to given value.


### GetPhase

`func (o *V2KvExport) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *V2KvExport) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *V2KvExport) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetRequestedAt

`func (o *V2KvExport) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *V2KvExport) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *V2KvExport) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.


### GetUpdatedAt

`func (o *V2KvExport) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V2KvExport) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V2KvExport) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetExportTime

`func (o *V2KvExport) GetExportTime() time.Time`

GetExportTime returns the ExportTime field if non-nil, zero value otherwise.

### GetExportTimeOk

`func (o *V2KvExport) GetExportTimeOk() (*time.Time, bool)`

GetExportTimeOk returns a tuple with the ExportTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTime

`func (o *V2KvExport) SetExportTime(v time.Time)`

SetExportTime sets ExportTime field to given value.

### HasExportTime

`func (o *V2KvExport) HasExportTime() bool`

HasExportTime returns a boolean if a field has been set.

### GetCompletedAt

`func (o *V2KvExport) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *V2KvExport) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *V2KvExport) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *V2KvExport) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetItemCount

`func (o *V2KvExport) GetItemCount() int32`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *V2KvExport) GetItemCountOk() (*int32, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *V2KvExport) SetItemCount(v int32)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *V2KvExport) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### GetSkipped

`func (o *V2KvExport) GetSkipped() V2KvExportSkipped`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *V2KvExport) GetSkippedOk() (*V2KvExportSkipped, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *V2KvExport) SetSkipped(v V2KvExportSkipped)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *V2KvExport) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetSizeBytes

`func (o *V2KvExport) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *V2KvExport) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *V2KvExport) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *V2KvExport) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetParts

`func (o *V2KvExport) GetParts() int32`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *V2KvExport) GetPartsOk() (*int32, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *V2KvExport) SetParts(v int32)`

SetParts sets Parts field to given value.

### HasParts

`func (o *V2KvExport) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetExpiresAt

`func (o *V2KvExport) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *V2KvExport) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *V2KvExport) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *V2KvExport) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetError

`func (o *V2KvExport) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2KvExport) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2KvExport) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V2KvExport) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDownloads

`func (o *V2KvExport) GetDownloads() []V2KvExportDownload`

GetDownloads returns the Downloads field if non-nil, zero value otherwise.

### GetDownloadsOk

`func (o *V2KvExport) GetDownloadsOk() (*[]V2KvExportDownload, bool)`

GetDownloadsOk returns a tuple with the Downloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloads

`func (o *V2KvExport) SetDownloads(v []V2KvExportDownload)`

SetDownloads sets Downloads field to given value.

### HasDownloads

`func (o *V2KvExport) HasDownloads() bool`

HasDownloads returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


