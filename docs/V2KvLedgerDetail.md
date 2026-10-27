# V2KvLedgerDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Collection** | **string** | A collection, or a pattern where * is one whole segment. | 
**Chains** | **int32** |  | 
**RetentionDays** | **int32** |  | 
**LockMode** | **string** |  | 
**PersonalFields** | **[]string** |  | 
**SubjectField** | Pointer to **string** |  | [optional] 
**SubjectMode** | Pointer to **string** | Present when subjectField is set. | [optional] 
**SubjectCase** | Pointer to **string** | Present when subjectField is set. | [optional] 
**Anchor** | [**V2KvLedgerAnchor**](V2KvLedgerAnchor.md) |  | 
**WriterPolicy** | Pointer to [**V2KvLedgerAdminWriterPolicy**](V2KvLedgerAdminWriterPolicy.md) |  | [optional] 
**State** | **string** | PENDING for 75 seconds after a declare, then ACTIVE or FAILED. DROPPED after KV_ledgers_delete. | 
**CreatedAt** | **time.Time** |  | 
**ActiveAfter** | **time.Time** | When a PENDING ledger is checked again and becomes ACTIVE or FAILED. | 
**DefinitionHash** | **string** | Hex SHA-256 of the canonical (RFC 8785) definition. | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Error** | Pointer to **string** | FAILED only: why the ledger failed. | [optional] 
**DroppedAt** | Pointer to **time.Time** |  | [optional] 
**RetainUntil** | Pointer to **time.Time** | The archive retention end; an ACTIVE ledger can be dropped after it. | [optional] 
**LastDigestAt** | Pointer to **time.Time** |  | [optional] 
**LastVerify** | Pointer to [**V2KvLedgerLastVerify**](V2KvLedgerLastVerify.md) |  | [optional] 

## Methods

### NewV2KvLedgerDetail

`func NewV2KvLedgerDetail(name string, collection string, chains int32, retentionDays int32, lockMode string, personalFields []string, anchor V2KvLedgerAnchor, state string, createdAt time.Time, activeAfter time.Time, definitionHash string, ) *V2KvLedgerDetail`

NewV2KvLedgerDetail instantiates a new V2KvLedgerDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerDetailWithDefaults

`func NewV2KvLedgerDetailWithDefaults() *V2KvLedgerDetail`

NewV2KvLedgerDetailWithDefaults instantiates a new V2KvLedgerDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2KvLedgerDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2KvLedgerDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2KvLedgerDetail) SetName(v string)`

SetName sets Name field to given value.


### GetCollection

`func (o *V2KvLedgerDetail) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *V2KvLedgerDetail) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *V2KvLedgerDetail) SetCollection(v string)`

SetCollection sets Collection field to given value.


### GetChains

`func (o *V2KvLedgerDetail) GetChains() int32`

GetChains returns the Chains field if non-nil, zero value otherwise.

### GetChainsOk

`func (o *V2KvLedgerDetail) GetChainsOk() (*int32, bool)`

GetChainsOk returns a tuple with the Chains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChains

`func (o *V2KvLedgerDetail) SetChains(v int32)`

SetChains sets Chains field to given value.


### GetRetentionDays

`func (o *V2KvLedgerDetail) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *V2KvLedgerDetail) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *V2KvLedgerDetail) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.


### GetLockMode

`func (o *V2KvLedgerDetail) GetLockMode() string`

GetLockMode returns the LockMode field if non-nil, zero value otherwise.

### GetLockModeOk

`func (o *V2KvLedgerDetail) GetLockModeOk() (*string, bool)`

GetLockModeOk returns a tuple with the LockMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockMode

`func (o *V2KvLedgerDetail) SetLockMode(v string)`

SetLockMode sets LockMode field to given value.


### GetPersonalFields

`func (o *V2KvLedgerDetail) GetPersonalFields() []string`

GetPersonalFields returns the PersonalFields field if non-nil, zero value otherwise.

### GetPersonalFieldsOk

`func (o *V2KvLedgerDetail) GetPersonalFieldsOk() (*[]string, bool)`

GetPersonalFieldsOk returns a tuple with the PersonalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalFields

`func (o *V2KvLedgerDetail) SetPersonalFields(v []string)`

SetPersonalFields sets PersonalFields field to given value.


### GetSubjectField

`func (o *V2KvLedgerDetail) GetSubjectField() string`

GetSubjectField returns the SubjectField field if non-nil, zero value otherwise.

### GetSubjectFieldOk

`func (o *V2KvLedgerDetail) GetSubjectFieldOk() (*string, bool)`

GetSubjectFieldOk returns a tuple with the SubjectField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectField

`func (o *V2KvLedgerDetail) SetSubjectField(v string)`

SetSubjectField sets SubjectField field to given value.

### HasSubjectField

`func (o *V2KvLedgerDetail) HasSubjectField() bool`

HasSubjectField returns a boolean if a field has been set.

### GetSubjectMode

`func (o *V2KvLedgerDetail) GetSubjectMode() string`

GetSubjectMode returns the SubjectMode field if non-nil, zero value otherwise.

### GetSubjectModeOk

`func (o *V2KvLedgerDetail) GetSubjectModeOk() (*string, bool)`

GetSubjectModeOk returns a tuple with the SubjectMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectMode

`func (o *V2KvLedgerDetail) SetSubjectMode(v string)`

SetSubjectMode sets SubjectMode field to given value.

### HasSubjectMode

`func (o *V2KvLedgerDetail) HasSubjectMode() bool`

HasSubjectMode returns a boolean if a field has been set.

### GetSubjectCase

`func (o *V2KvLedgerDetail) GetSubjectCase() string`

GetSubjectCase returns the SubjectCase field if non-nil, zero value otherwise.

### GetSubjectCaseOk

`func (o *V2KvLedgerDetail) GetSubjectCaseOk() (*string, bool)`

GetSubjectCaseOk returns a tuple with the SubjectCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCase

`func (o *V2KvLedgerDetail) SetSubjectCase(v string)`

SetSubjectCase sets SubjectCase field to given value.

### HasSubjectCase

`func (o *V2KvLedgerDetail) HasSubjectCase() bool`

HasSubjectCase returns a boolean if a field has been set.

### GetAnchor

`func (o *V2KvLedgerDetail) GetAnchor() V2KvLedgerAnchor`

GetAnchor returns the Anchor field if non-nil, zero value otherwise.

### GetAnchorOk

`func (o *V2KvLedgerDetail) GetAnchorOk() (*V2KvLedgerAnchor, bool)`

GetAnchorOk returns a tuple with the Anchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchor

`func (o *V2KvLedgerDetail) SetAnchor(v V2KvLedgerAnchor)`

SetAnchor sets Anchor field to given value.


### GetWriterPolicy

`func (o *V2KvLedgerDetail) GetWriterPolicy() V2KvLedgerAdminWriterPolicy`

GetWriterPolicy returns the WriterPolicy field if non-nil, zero value otherwise.

### GetWriterPolicyOk

`func (o *V2KvLedgerDetail) GetWriterPolicyOk() (*V2KvLedgerAdminWriterPolicy, bool)`

GetWriterPolicyOk returns a tuple with the WriterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriterPolicy

`func (o *V2KvLedgerDetail) SetWriterPolicy(v V2KvLedgerAdminWriterPolicy)`

SetWriterPolicy sets WriterPolicy field to given value.

### HasWriterPolicy

`func (o *V2KvLedgerDetail) HasWriterPolicy() bool`

HasWriterPolicy returns a boolean if a field has been set.

### GetState

`func (o *V2KvLedgerDetail) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V2KvLedgerDetail) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V2KvLedgerDetail) SetState(v string)`

SetState sets State field to given value.


### GetCreatedAt

`func (o *V2KvLedgerDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2KvLedgerDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2KvLedgerDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetActiveAfter

`func (o *V2KvLedgerDetail) GetActiveAfter() time.Time`

GetActiveAfter returns the ActiveAfter field if non-nil, zero value otherwise.

### GetActiveAfterOk

`func (o *V2KvLedgerDetail) GetActiveAfterOk() (*time.Time, bool)`

GetActiveAfterOk returns a tuple with the ActiveAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAfter

`func (o *V2KvLedgerDetail) SetActiveAfter(v time.Time)`

SetActiveAfter sets ActiveAfter field to given value.


### GetDefinitionHash

`func (o *V2KvLedgerDetail) GetDefinitionHash() string`

GetDefinitionHash returns the DefinitionHash field if non-nil, zero value otherwise.

### GetDefinitionHashOk

`func (o *V2KvLedgerDetail) GetDefinitionHashOk() (*string, bool)`

GetDefinitionHashOk returns a tuple with the DefinitionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionHash

`func (o *V2KvLedgerDetail) SetDefinitionHash(v string)`

SetDefinitionHash sets DefinitionHash field to given value.


### GetUpdatedAt

`func (o *V2KvLedgerDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V2KvLedgerDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V2KvLedgerDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V2KvLedgerDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetError

`func (o *V2KvLedgerDetail) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2KvLedgerDetail) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2KvLedgerDetail) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V2KvLedgerDetail) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDroppedAt

`func (o *V2KvLedgerDetail) GetDroppedAt() time.Time`

GetDroppedAt returns the DroppedAt field if non-nil, zero value otherwise.

### GetDroppedAtOk

`func (o *V2KvLedgerDetail) GetDroppedAtOk() (*time.Time, bool)`

GetDroppedAtOk returns a tuple with the DroppedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroppedAt

`func (o *V2KvLedgerDetail) SetDroppedAt(v time.Time)`

SetDroppedAt sets DroppedAt field to given value.

### HasDroppedAt

`func (o *V2KvLedgerDetail) HasDroppedAt() bool`

HasDroppedAt returns a boolean if a field has been set.

### GetRetainUntil

`func (o *V2KvLedgerDetail) GetRetainUntil() time.Time`

GetRetainUntil returns the RetainUntil field if non-nil, zero value otherwise.

### GetRetainUntilOk

`func (o *V2KvLedgerDetail) GetRetainUntilOk() (*time.Time, bool)`

GetRetainUntilOk returns a tuple with the RetainUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainUntil

`func (o *V2KvLedgerDetail) SetRetainUntil(v time.Time)`

SetRetainUntil sets RetainUntil field to given value.

### HasRetainUntil

`func (o *V2KvLedgerDetail) HasRetainUntil() bool`

HasRetainUntil returns a boolean if a field has been set.

### GetLastDigestAt

`func (o *V2KvLedgerDetail) GetLastDigestAt() time.Time`

GetLastDigestAt returns the LastDigestAt field if non-nil, zero value otherwise.

### GetLastDigestAtOk

`func (o *V2KvLedgerDetail) GetLastDigestAtOk() (*time.Time, bool)`

GetLastDigestAtOk returns a tuple with the LastDigestAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDigestAt

`func (o *V2KvLedgerDetail) SetLastDigestAt(v time.Time)`

SetLastDigestAt sets LastDigestAt field to given value.

### HasLastDigestAt

`func (o *V2KvLedgerDetail) HasLastDigestAt() bool`

HasLastDigestAt returns a boolean if a field has been set.

### GetLastVerify

`func (o *V2KvLedgerDetail) GetLastVerify() V2KvLedgerLastVerify`

GetLastVerify returns the LastVerify field if non-nil, zero value otherwise.

### GetLastVerifyOk

`func (o *V2KvLedgerDetail) GetLastVerifyOk() (*V2KvLedgerLastVerify, bool)`

GetLastVerifyOk returns a tuple with the LastVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVerify

`func (o *V2KvLedgerDetail) SetLastVerify(v V2KvLedgerLastVerify)`

SetLastVerify sets LastVerify field to given value.

### HasLastVerify

`func (o *V2KvLedgerDetail) HasLastVerify() bool`

HasLastVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


