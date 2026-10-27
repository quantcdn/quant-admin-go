# V2KvLedgerArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]V2KvLedgerArchiveObject**](V2KvLedgerArchiveObject.md) |  | 
**Digests** | [**[]V2KvLedgerArchiveDigest**](V2KvLedgerArchiveDigest.md) | First page (or digest_cursor pages) only. | 
**Timestamps** | [**[]V2KvLedgerArchiveFile**](V2KvLedgerArchiveFile.md) | RFC 3161 tokens and anchor binding files. First page (or digest_cursor pages) only. | 
**DigestsTruncated** | Pointer to **bool** | True when more than 2,000 digest files exist: continue with next_digest_cursor. | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**NextCursor** | **NullableString** | Pass as cursor for the next page of objects. | 
**NextDigestCursor** | **NullableString** | Pass as digest_cursor for the next page of digests. | 

## Methods

### NewV2KvLedgerArchive

`func NewV2KvLedgerArchive(objects []V2KvLedgerArchiveObject, digests []V2KvLedgerArchiveDigest, timestamps []V2KvLedgerArchiveFile, nextCursor NullableString, nextDigestCursor NullableString, ) *V2KvLedgerArchive`

NewV2KvLedgerArchive instantiates a new V2KvLedgerArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerArchiveWithDefaults

`func NewV2KvLedgerArchiveWithDefaults() *V2KvLedgerArchive`

NewV2KvLedgerArchiveWithDefaults instantiates a new V2KvLedgerArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *V2KvLedgerArchive) GetObjects() []V2KvLedgerArchiveObject`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *V2KvLedgerArchive) GetObjectsOk() (*[]V2KvLedgerArchiveObject, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *V2KvLedgerArchive) SetObjects(v []V2KvLedgerArchiveObject)`

SetObjects sets Objects field to given value.


### GetDigests

`func (o *V2KvLedgerArchive) GetDigests() []V2KvLedgerArchiveDigest`

GetDigests returns the Digests field if non-nil, zero value otherwise.

### GetDigestsOk

`func (o *V2KvLedgerArchive) GetDigestsOk() (*[]V2KvLedgerArchiveDigest, bool)`

GetDigestsOk returns a tuple with the Digests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigests

`func (o *V2KvLedgerArchive) SetDigests(v []V2KvLedgerArchiveDigest)`

SetDigests sets Digests field to given value.


### GetTimestamps

`func (o *V2KvLedgerArchive) GetTimestamps() []V2KvLedgerArchiveFile`

GetTimestamps returns the Timestamps field if non-nil, zero value otherwise.

### GetTimestampsOk

`func (o *V2KvLedgerArchive) GetTimestampsOk() (*[]V2KvLedgerArchiveFile, bool)`

GetTimestampsOk returns a tuple with the Timestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamps

`func (o *V2KvLedgerArchive) SetTimestamps(v []V2KvLedgerArchiveFile)`

SetTimestamps sets Timestamps field to given value.


### GetDigestsTruncated

`func (o *V2KvLedgerArchive) GetDigestsTruncated() bool`

GetDigestsTruncated returns the DigestsTruncated field if non-nil, zero value otherwise.

### GetDigestsTruncatedOk

`func (o *V2KvLedgerArchive) GetDigestsTruncatedOk() (*bool, bool)`

GetDigestsTruncatedOk returns a tuple with the DigestsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestsTruncated

`func (o *V2KvLedgerArchive) SetDigestsTruncated(v bool)`

SetDigestsTruncated sets DigestsTruncated field to given value.

### HasDigestsTruncated

`func (o *V2KvLedgerArchive) HasDigestsTruncated() bool`

HasDigestsTruncated returns a boolean if a field has been set.

### GetNote

`func (o *V2KvLedgerArchive) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *V2KvLedgerArchive) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *V2KvLedgerArchive) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *V2KvLedgerArchive) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetNextCursor

`func (o *V2KvLedgerArchive) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *V2KvLedgerArchive) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *V2KvLedgerArchive) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *V2KvLedgerArchive) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *V2KvLedgerArchive) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetNextDigestCursor

`func (o *V2KvLedgerArchive) GetNextDigestCursor() string`

GetNextDigestCursor returns the NextDigestCursor field if non-nil, zero value otherwise.

### GetNextDigestCursorOk

`func (o *V2KvLedgerArchive) GetNextDigestCursorOk() (*string, bool)`

GetNextDigestCursorOk returns a tuple with the NextDigestCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDigestCursor

`func (o *V2KvLedgerArchive) SetNextDigestCursor(v string)`

SetNextDigestCursor sets NextDigestCursor field to given value.


### SetNextDigestCursorNil

`func (o *V2KvLedgerArchive) SetNextDigestCursorNil(b bool)`

 SetNextDigestCursorNil sets the value for NextDigestCursor to be an explicit nil

### UnsetNextDigestCursor
`func (o *V2KvLedgerArchive) UnsetNextDigestCursor()`

UnsetNextDigestCursor ensures that no value is present for NextDigestCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


