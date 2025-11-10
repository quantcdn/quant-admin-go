# V1Meta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**SeqNum** | Pointer to **int32** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**PublishedRevision** | Pointer to **int32** |  | [optional] 
**PublishedMd5** | Pointer to **string** |  | [optional] 
**ByteLength** | Pointer to **int32** |  | [optional] 
**RevisionCount** | Pointer to **int32** |  | [optional] 
**HighestRevisionNumber** | Pointer to **int32** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DeletedTimestamp** | Pointer to **int32** |  | [optional] 
**Md5** | Pointer to **string** |  | [optional] 
**RevisionNumber** | Pointer to **int32** |  | [optional] 
**DateTimestamp** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1Meta

`func NewV1Meta() *V1Meta`

NewV1Meta instantiates a new V1Meta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MetaWithDefaults

`func NewV1MetaWithDefaults() *V1Meta`

NewV1MetaWithDefaults instantiates a new V1Meta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *V1Meta) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V1Meta) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V1Meta) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *V1Meta) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetType

`func (o *V1Meta) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1Meta) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1Meta) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1Meta) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSeqNum

`func (o *V1Meta) GetSeqNum() int32`

GetSeqNum returns the SeqNum field if non-nil, zero value otherwise.

### GetSeqNumOk

`func (o *V1Meta) GetSeqNumOk() (*int32, bool)`

GetSeqNumOk returns a tuple with the SeqNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqNum

`func (o *V1Meta) SetSeqNum(v int32)`

SetSeqNum sets SeqNum field to given value.

### HasSeqNum

`func (o *V1Meta) HasSeqNum() bool`

HasSeqNum returns a boolean if a field has been set.

### GetPublished

`func (o *V1Meta) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *V1Meta) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *V1Meta) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *V1Meta) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetPublishedRevision

`func (o *V1Meta) GetPublishedRevision() int32`

GetPublishedRevision returns the PublishedRevision field if non-nil, zero value otherwise.

### GetPublishedRevisionOk

`func (o *V1Meta) GetPublishedRevisionOk() (*int32, bool)`

GetPublishedRevisionOk returns a tuple with the PublishedRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedRevision

`func (o *V1Meta) SetPublishedRevision(v int32)`

SetPublishedRevision sets PublishedRevision field to given value.

### HasPublishedRevision

`func (o *V1Meta) HasPublishedRevision() bool`

HasPublishedRevision returns a boolean if a field has been set.

### GetPublishedMd5

`func (o *V1Meta) GetPublishedMd5() string`

GetPublishedMd5 returns the PublishedMd5 field if non-nil, zero value otherwise.

### GetPublishedMd5Ok

`func (o *V1Meta) GetPublishedMd5Ok() (*string, bool)`

GetPublishedMd5Ok returns a tuple with the PublishedMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedMd5

`func (o *V1Meta) SetPublishedMd5(v string)`

SetPublishedMd5 sets PublishedMd5 field to given value.

### HasPublishedMd5

`func (o *V1Meta) HasPublishedMd5() bool`

HasPublishedMd5 returns a boolean if a field has been set.

### GetByteLength

`func (o *V1Meta) GetByteLength() int32`

GetByteLength returns the ByteLength field if non-nil, zero value otherwise.

### GetByteLengthOk

`func (o *V1Meta) GetByteLengthOk() (*int32, bool)`

GetByteLengthOk returns a tuple with the ByteLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteLength

`func (o *V1Meta) SetByteLength(v int32)`

SetByteLength sets ByteLength field to given value.

### HasByteLength

`func (o *V1Meta) HasByteLength() bool`

HasByteLength returns a boolean if a field has been set.

### GetRevisionCount

`func (o *V1Meta) GetRevisionCount() int32`

GetRevisionCount returns the RevisionCount field if non-nil, zero value otherwise.

### GetRevisionCountOk

`func (o *V1Meta) GetRevisionCountOk() (*int32, bool)`

GetRevisionCountOk returns a tuple with the RevisionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionCount

`func (o *V1Meta) SetRevisionCount(v int32)`

SetRevisionCount sets RevisionCount field to given value.

### HasRevisionCount

`func (o *V1Meta) HasRevisionCount() bool`

HasRevisionCount returns a boolean if a field has been set.

### GetHighestRevisionNumber

`func (o *V1Meta) GetHighestRevisionNumber() int32`

GetHighestRevisionNumber returns the HighestRevisionNumber field if non-nil, zero value otherwise.

### GetHighestRevisionNumberOk

`func (o *V1Meta) GetHighestRevisionNumberOk() (*int32, bool)`

GetHighestRevisionNumberOk returns a tuple with the HighestRevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestRevisionNumber

`func (o *V1Meta) SetHighestRevisionNumber(v int32)`

SetHighestRevisionNumber sets HighestRevisionNumber field to given value.

### HasHighestRevisionNumber

`func (o *V1Meta) HasHighestRevisionNumber() bool`

HasHighestRevisionNumber returns a boolean if a field has been set.

### GetDeleted

`func (o *V1Meta) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *V1Meta) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *V1Meta) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *V1Meta) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDeletedTimestamp

`func (o *V1Meta) GetDeletedTimestamp() int32`

GetDeletedTimestamp returns the DeletedTimestamp field if non-nil, zero value otherwise.

### GetDeletedTimestampOk

`func (o *V1Meta) GetDeletedTimestampOk() (*int32, bool)`

GetDeletedTimestampOk returns a tuple with the DeletedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedTimestamp

`func (o *V1Meta) SetDeletedTimestamp(v int32)`

SetDeletedTimestamp sets DeletedTimestamp field to given value.

### HasDeletedTimestamp

`func (o *V1Meta) HasDeletedTimestamp() bool`

HasDeletedTimestamp returns a boolean if a field has been set.

### GetMd5

`func (o *V1Meta) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *V1Meta) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *V1Meta) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *V1Meta) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetRevisionNumber

`func (o *V1Meta) GetRevisionNumber() int32`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *V1Meta) GetRevisionNumberOk() (*int32, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *V1Meta) SetRevisionNumber(v int32)`

SetRevisionNumber sets RevisionNumber field to given value.

### HasRevisionNumber

`func (o *V1Meta) HasRevisionNumber() bool`

HasRevisionNumber returns a boolean if a field has been set.

### GetDateTimestamp

`func (o *V1Meta) GetDateTimestamp() int32`

GetDateTimestamp returns the DateTimestamp field if non-nil, zero value otherwise.

### GetDateTimestampOk

`func (o *V1Meta) GetDateTimestampOk() (*int32, bool)`

GetDateTimestampOk returns a tuple with the DateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimestamp

`func (o *V1Meta) SetDateTimestamp(v int32)`

SetDateTimestamp sets DateTimestamp field to given value.

### HasDateTimestamp

`func (o *V1Meta) HasDateTimestamp() bool`

HasDateTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


