# AiAuditRecordsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
**OrgId** | **string** |  | 
**StartDate** | **string** |  | 
**EndDate** | **string** |  | 
**Count** | **int32** |  | 
**Records** | [**[]AiAuditRecordRow**](AiAuditRecordRow.md) |  | 
**Scanned** | **int32** | Keys examined by this call | 
**SkippedArchived** | **int32** | Records older than 365 days (archived) that this call left out | 
**HasMore** | **bool** |  | 
**NextCursor** | Pointer to **string** |  | [optional] 

## Methods

### NewAiAuditRecordsPage

`func NewAiAuditRecordsPage(message string, error_ bool, orgId string, startDate string, endDate string, count int32, records []AiAuditRecordRow, scanned int32, skippedArchived int32, hasMore bool, ) *AiAuditRecordsPage`

NewAiAuditRecordsPage instantiates a new AiAuditRecordsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAuditRecordsPageWithDefaults

`func NewAiAuditRecordsPageWithDefaults() *AiAuditRecordsPage`

NewAiAuditRecordsPageWithDefaults instantiates a new AiAuditRecordsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AiAuditRecordsPage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AiAuditRecordsPage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AiAuditRecordsPage) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *AiAuditRecordsPage) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AiAuditRecordsPage) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AiAuditRecordsPage) SetError(v bool)`

SetError sets Error field to given value.


### GetOrgId

`func (o *AiAuditRecordsPage) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AiAuditRecordsPage) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AiAuditRecordsPage) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetStartDate

`func (o *AiAuditRecordsPage) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AiAuditRecordsPage) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AiAuditRecordsPage) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *AiAuditRecordsPage) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AiAuditRecordsPage) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AiAuditRecordsPage) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetCount

`func (o *AiAuditRecordsPage) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AiAuditRecordsPage) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AiAuditRecordsPage) SetCount(v int32)`

SetCount sets Count field to given value.


### GetRecords

`func (o *AiAuditRecordsPage) GetRecords() []AiAuditRecordRow`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *AiAuditRecordsPage) GetRecordsOk() (*[]AiAuditRecordRow, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *AiAuditRecordsPage) SetRecords(v []AiAuditRecordRow)`

SetRecords sets Records field to given value.


### GetScanned

`func (o *AiAuditRecordsPage) GetScanned() int32`

GetScanned returns the Scanned field if non-nil, zero value otherwise.

### GetScannedOk

`func (o *AiAuditRecordsPage) GetScannedOk() (*int32, bool)`

GetScannedOk returns a tuple with the Scanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanned

`func (o *AiAuditRecordsPage) SetScanned(v int32)`

SetScanned sets Scanned field to given value.


### GetSkippedArchived

`func (o *AiAuditRecordsPage) GetSkippedArchived() int32`

GetSkippedArchived returns the SkippedArchived field if non-nil, zero value otherwise.

### GetSkippedArchivedOk

`func (o *AiAuditRecordsPage) GetSkippedArchivedOk() (*int32, bool)`

GetSkippedArchivedOk returns a tuple with the SkippedArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedArchived

`func (o *AiAuditRecordsPage) SetSkippedArchived(v int32)`

SetSkippedArchived sets SkippedArchived field to given value.


### GetHasMore

`func (o *AiAuditRecordsPage) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *AiAuditRecordsPage) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *AiAuditRecordsPage) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetNextCursor

`func (o *AiAuditRecordsPage) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *AiAuditRecordsPage) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *AiAuditRecordsPage) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *AiAuditRecordsPage) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


