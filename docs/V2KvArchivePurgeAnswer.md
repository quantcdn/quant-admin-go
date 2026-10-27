# V2KvArchivePurgeAnswer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PurgeId** | Pointer to **string** | Set on a start (202). | [optional] 
**State** | Pointer to **string** | Set on a start (202). | [optional] 
**DryRun** | Pointer to **bool** |  | [optional] 
**Store** | Pointer to **string** |  | [optional] 
**Eligible** | Pointer to **bool** |  | [optional] 
**Partial** | Pointer to **bool** |  | [optional] 
**Versions** | Pointer to **int32** |  | [optional] 
**Bytes** | Pointer to **int32** |  | [optional] 
**LedgerObjects** | Pointer to **int32** |  | [optional] 
**DigestObjects** | Pointer to **int32** |  | [optional] 
**AnchorObjects** | Pointer to **int32** |  | [optional] 
**DeleteMarkers** | Pointer to **int32** |  | [optional] 
**LatestRetainUntil** | Pointer to **NullableTime** |  | [optional] 
**EarliestPurgeAt** | Pointer to **NullableTime** |  | [optional] 
**QuietUntil** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewV2KvArchivePurgeAnswer

`func NewV2KvArchivePurgeAnswer() *V2KvArchivePurgeAnswer`

NewV2KvArchivePurgeAnswer instantiates a new V2KvArchivePurgeAnswer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvArchivePurgeAnswerWithDefaults

`func NewV2KvArchivePurgeAnswerWithDefaults() *V2KvArchivePurgeAnswer`

NewV2KvArchivePurgeAnswerWithDefaults instantiates a new V2KvArchivePurgeAnswer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurgeId

`func (o *V2KvArchivePurgeAnswer) GetPurgeId() string`

GetPurgeId returns the PurgeId field if non-nil, zero value otherwise.

### GetPurgeIdOk

`func (o *V2KvArchivePurgeAnswer) GetPurgeIdOk() (*string, bool)`

GetPurgeIdOk returns a tuple with the PurgeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeId

`func (o *V2KvArchivePurgeAnswer) SetPurgeId(v string)`

SetPurgeId sets PurgeId field to given value.

### HasPurgeId

`func (o *V2KvArchivePurgeAnswer) HasPurgeId() bool`

HasPurgeId returns a boolean if a field has been set.

### GetState

`func (o *V2KvArchivePurgeAnswer) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V2KvArchivePurgeAnswer) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V2KvArchivePurgeAnswer) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *V2KvArchivePurgeAnswer) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDryRun

`func (o *V2KvArchivePurgeAnswer) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *V2KvArchivePurgeAnswer) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *V2KvArchivePurgeAnswer) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *V2KvArchivePurgeAnswer) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetStore

`func (o *V2KvArchivePurgeAnswer) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *V2KvArchivePurgeAnswer) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *V2KvArchivePurgeAnswer) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *V2KvArchivePurgeAnswer) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetEligible

`func (o *V2KvArchivePurgeAnswer) GetEligible() bool`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *V2KvArchivePurgeAnswer) GetEligibleOk() (*bool, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *V2KvArchivePurgeAnswer) SetEligible(v bool)`

SetEligible sets Eligible field to given value.

### HasEligible

`func (o *V2KvArchivePurgeAnswer) HasEligible() bool`

HasEligible returns a boolean if a field has been set.

### GetPartial

`func (o *V2KvArchivePurgeAnswer) GetPartial() bool`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *V2KvArchivePurgeAnswer) GetPartialOk() (*bool, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *V2KvArchivePurgeAnswer) SetPartial(v bool)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *V2KvArchivePurgeAnswer) HasPartial() bool`

HasPartial returns a boolean if a field has been set.

### GetVersions

`func (o *V2KvArchivePurgeAnswer) GetVersions() int32`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *V2KvArchivePurgeAnswer) GetVersionsOk() (*int32, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *V2KvArchivePurgeAnswer) SetVersions(v int32)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *V2KvArchivePurgeAnswer) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetBytes

`func (o *V2KvArchivePurgeAnswer) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *V2KvArchivePurgeAnswer) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *V2KvArchivePurgeAnswer) SetBytes(v int32)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *V2KvArchivePurgeAnswer) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetLedgerObjects

`func (o *V2KvArchivePurgeAnswer) GetLedgerObjects() int32`

GetLedgerObjects returns the LedgerObjects field if non-nil, zero value otherwise.

### GetLedgerObjectsOk

`func (o *V2KvArchivePurgeAnswer) GetLedgerObjectsOk() (*int32, bool)`

GetLedgerObjectsOk returns a tuple with the LedgerObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerObjects

`func (o *V2KvArchivePurgeAnswer) SetLedgerObjects(v int32)`

SetLedgerObjects sets LedgerObjects field to given value.

### HasLedgerObjects

`func (o *V2KvArchivePurgeAnswer) HasLedgerObjects() bool`

HasLedgerObjects returns a boolean if a field has been set.

### GetDigestObjects

`func (o *V2KvArchivePurgeAnswer) GetDigestObjects() int32`

GetDigestObjects returns the DigestObjects field if non-nil, zero value otherwise.

### GetDigestObjectsOk

`func (o *V2KvArchivePurgeAnswer) GetDigestObjectsOk() (*int32, bool)`

GetDigestObjectsOk returns a tuple with the DigestObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestObjects

`func (o *V2KvArchivePurgeAnswer) SetDigestObjects(v int32)`

SetDigestObjects sets DigestObjects field to given value.

### HasDigestObjects

`func (o *V2KvArchivePurgeAnswer) HasDigestObjects() bool`

HasDigestObjects returns a boolean if a field has been set.

### GetAnchorObjects

`func (o *V2KvArchivePurgeAnswer) GetAnchorObjects() int32`

GetAnchorObjects returns the AnchorObjects field if non-nil, zero value otherwise.

### GetAnchorObjectsOk

`func (o *V2KvArchivePurgeAnswer) GetAnchorObjectsOk() (*int32, bool)`

GetAnchorObjectsOk returns a tuple with the AnchorObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorObjects

`func (o *V2KvArchivePurgeAnswer) SetAnchorObjects(v int32)`

SetAnchorObjects sets AnchorObjects field to given value.

### HasAnchorObjects

`func (o *V2KvArchivePurgeAnswer) HasAnchorObjects() bool`

HasAnchorObjects returns a boolean if a field has been set.

### GetDeleteMarkers

`func (o *V2KvArchivePurgeAnswer) GetDeleteMarkers() int32`

GetDeleteMarkers returns the DeleteMarkers field if non-nil, zero value otherwise.

### GetDeleteMarkersOk

`func (o *V2KvArchivePurgeAnswer) GetDeleteMarkersOk() (*int32, bool)`

GetDeleteMarkersOk returns a tuple with the DeleteMarkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteMarkers

`func (o *V2KvArchivePurgeAnswer) SetDeleteMarkers(v int32)`

SetDeleteMarkers sets DeleteMarkers field to given value.

### HasDeleteMarkers

`func (o *V2KvArchivePurgeAnswer) HasDeleteMarkers() bool`

HasDeleteMarkers returns a boolean if a field has been set.

### GetLatestRetainUntil

`func (o *V2KvArchivePurgeAnswer) GetLatestRetainUntil() time.Time`

GetLatestRetainUntil returns the LatestRetainUntil field if non-nil, zero value otherwise.

### GetLatestRetainUntilOk

`func (o *V2KvArchivePurgeAnswer) GetLatestRetainUntilOk() (*time.Time, bool)`

GetLatestRetainUntilOk returns a tuple with the LatestRetainUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRetainUntil

`func (o *V2KvArchivePurgeAnswer) SetLatestRetainUntil(v time.Time)`

SetLatestRetainUntil sets LatestRetainUntil field to given value.

### HasLatestRetainUntil

`func (o *V2KvArchivePurgeAnswer) HasLatestRetainUntil() bool`

HasLatestRetainUntil returns a boolean if a field has been set.

### SetLatestRetainUntilNil

`func (o *V2KvArchivePurgeAnswer) SetLatestRetainUntilNil(b bool)`

 SetLatestRetainUntilNil sets the value for LatestRetainUntil to be an explicit nil

### UnsetLatestRetainUntil
`func (o *V2KvArchivePurgeAnswer) UnsetLatestRetainUntil()`

UnsetLatestRetainUntil ensures that no value is present for LatestRetainUntil, not even an explicit nil
### GetEarliestPurgeAt

`func (o *V2KvArchivePurgeAnswer) GetEarliestPurgeAt() time.Time`

GetEarliestPurgeAt returns the EarliestPurgeAt field if non-nil, zero value otherwise.

### GetEarliestPurgeAtOk

`func (o *V2KvArchivePurgeAnswer) GetEarliestPurgeAtOk() (*time.Time, bool)`

GetEarliestPurgeAtOk returns a tuple with the EarliestPurgeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestPurgeAt

`func (o *V2KvArchivePurgeAnswer) SetEarliestPurgeAt(v time.Time)`

SetEarliestPurgeAt sets EarliestPurgeAt field to given value.

### HasEarliestPurgeAt

`func (o *V2KvArchivePurgeAnswer) HasEarliestPurgeAt() bool`

HasEarliestPurgeAt returns a boolean if a field has been set.

### SetEarliestPurgeAtNil

`func (o *V2KvArchivePurgeAnswer) SetEarliestPurgeAtNil(b bool)`

 SetEarliestPurgeAtNil sets the value for EarliestPurgeAt to be an explicit nil

### UnsetEarliestPurgeAt
`func (o *V2KvArchivePurgeAnswer) UnsetEarliestPurgeAt()`

UnsetEarliestPurgeAt ensures that no value is present for EarliestPurgeAt, not even an explicit nil
### GetQuietUntil

`func (o *V2KvArchivePurgeAnswer) GetQuietUntil() time.Time`

GetQuietUntil returns the QuietUntil field if non-nil, zero value otherwise.

### GetQuietUntilOk

`func (o *V2KvArchivePurgeAnswer) GetQuietUntilOk() (*time.Time, bool)`

GetQuietUntilOk returns a tuple with the QuietUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuietUntil

`func (o *V2KvArchivePurgeAnswer) SetQuietUntil(v time.Time)`

SetQuietUntil sets QuietUntil field to given value.

### HasQuietUntil

`func (o *V2KvArchivePurgeAnswer) HasQuietUntil() bool`

HasQuietUntil returns a boolean if a field has been set.

### SetQuietUntilNil

`func (o *V2KvArchivePurgeAnswer) SetQuietUntilNil(b bool)`

 SetQuietUntilNil sets the value for QuietUntil to be an explicit nil

### UnsetQuietUntil
`func (o *V2KvArchivePurgeAnswer) UnsetQuietUntil()`

UnsetQuietUntil ensures that no value is present for QuietUntil, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


