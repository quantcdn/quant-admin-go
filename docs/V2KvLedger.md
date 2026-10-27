# V2KvLedger

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

## Methods

### NewV2KvLedger

`func NewV2KvLedger(name string, collection string, chains int32, retentionDays int32, lockMode string, personalFields []string, anchor V2KvLedgerAnchor, state string, createdAt time.Time, activeAfter time.Time, definitionHash string, ) *V2KvLedger`

NewV2KvLedger instantiates a new V2KvLedger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerWithDefaults

`func NewV2KvLedgerWithDefaults() *V2KvLedger`

NewV2KvLedgerWithDefaults instantiates a new V2KvLedger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2KvLedger) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2KvLedger) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2KvLedger) SetName(v string)`

SetName sets Name field to given value.


### GetCollection

`func (o *V2KvLedger) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *V2KvLedger) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *V2KvLedger) SetCollection(v string)`

SetCollection sets Collection field to given value.


### GetChains

`func (o *V2KvLedger) GetChains() int32`

GetChains returns the Chains field if non-nil, zero value otherwise.

### GetChainsOk

`func (o *V2KvLedger) GetChainsOk() (*int32, bool)`

GetChainsOk returns a tuple with the Chains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChains

`func (o *V2KvLedger) SetChains(v int32)`

SetChains sets Chains field to given value.


### GetRetentionDays

`func (o *V2KvLedger) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *V2KvLedger) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *V2KvLedger) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.


### GetLockMode

`func (o *V2KvLedger) GetLockMode() string`

GetLockMode returns the LockMode field if non-nil, zero value otherwise.

### GetLockModeOk

`func (o *V2KvLedger) GetLockModeOk() (*string, bool)`

GetLockModeOk returns a tuple with the LockMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockMode

`func (o *V2KvLedger) SetLockMode(v string)`

SetLockMode sets LockMode field to given value.


### GetPersonalFields

`func (o *V2KvLedger) GetPersonalFields() []string`

GetPersonalFields returns the PersonalFields field if non-nil, zero value otherwise.

### GetPersonalFieldsOk

`func (o *V2KvLedger) GetPersonalFieldsOk() (*[]string, bool)`

GetPersonalFieldsOk returns a tuple with the PersonalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalFields

`func (o *V2KvLedger) SetPersonalFields(v []string)`

SetPersonalFields sets PersonalFields field to given value.


### GetSubjectField

`func (o *V2KvLedger) GetSubjectField() string`

GetSubjectField returns the SubjectField field if non-nil, zero value otherwise.

### GetSubjectFieldOk

`func (o *V2KvLedger) GetSubjectFieldOk() (*string, bool)`

GetSubjectFieldOk returns a tuple with the SubjectField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectField

`func (o *V2KvLedger) SetSubjectField(v string)`

SetSubjectField sets SubjectField field to given value.

### HasSubjectField

`func (o *V2KvLedger) HasSubjectField() bool`

HasSubjectField returns a boolean if a field has been set.

### GetSubjectMode

`func (o *V2KvLedger) GetSubjectMode() string`

GetSubjectMode returns the SubjectMode field if non-nil, zero value otherwise.

### GetSubjectModeOk

`func (o *V2KvLedger) GetSubjectModeOk() (*string, bool)`

GetSubjectModeOk returns a tuple with the SubjectMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectMode

`func (o *V2KvLedger) SetSubjectMode(v string)`

SetSubjectMode sets SubjectMode field to given value.

### HasSubjectMode

`func (o *V2KvLedger) HasSubjectMode() bool`

HasSubjectMode returns a boolean if a field has been set.

### GetSubjectCase

`func (o *V2KvLedger) GetSubjectCase() string`

GetSubjectCase returns the SubjectCase field if non-nil, zero value otherwise.

### GetSubjectCaseOk

`func (o *V2KvLedger) GetSubjectCaseOk() (*string, bool)`

GetSubjectCaseOk returns a tuple with the SubjectCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCase

`func (o *V2KvLedger) SetSubjectCase(v string)`

SetSubjectCase sets SubjectCase field to given value.

### HasSubjectCase

`func (o *V2KvLedger) HasSubjectCase() bool`

HasSubjectCase returns a boolean if a field has been set.

### GetAnchor

`func (o *V2KvLedger) GetAnchor() V2KvLedgerAnchor`

GetAnchor returns the Anchor field if non-nil, zero value otherwise.

### GetAnchorOk

`func (o *V2KvLedger) GetAnchorOk() (*V2KvLedgerAnchor, bool)`

GetAnchorOk returns a tuple with the Anchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchor

`func (o *V2KvLedger) SetAnchor(v V2KvLedgerAnchor)`

SetAnchor sets Anchor field to given value.


### GetWriterPolicy

`func (o *V2KvLedger) GetWriterPolicy() V2KvLedgerAdminWriterPolicy`

GetWriterPolicy returns the WriterPolicy field if non-nil, zero value otherwise.

### GetWriterPolicyOk

`func (o *V2KvLedger) GetWriterPolicyOk() (*V2KvLedgerAdminWriterPolicy, bool)`

GetWriterPolicyOk returns a tuple with the WriterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriterPolicy

`func (o *V2KvLedger) SetWriterPolicy(v V2KvLedgerAdminWriterPolicy)`

SetWriterPolicy sets WriterPolicy field to given value.

### HasWriterPolicy

`func (o *V2KvLedger) HasWriterPolicy() bool`

HasWriterPolicy returns a boolean if a field has been set.

### GetState

`func (o *V2KvLedger) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V2KvLedger) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V2KvLedger) SetState(v string)`

SetState sets State field to given value.


### GetCreatedAt

`func (o *V2KvLedger) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2KvLedger) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2KvLedger) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetActiveAfter

`func (o *V2KvLedger) GetActiveAfter() time.Time`

GetActiveAfter returns the ActiveAfter field if non-nil, zero value otherwise.

### GetActiveAfterOk

`func (o *V2KvLedger) GetActiveAfterOk() (*time.Time, bool)`

GetActiveAfterOk returns a tuple with the ActiveAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAfter

`func (o *V2KvLedger) SetActiveAfter(v time.Time)`

SetActiveAfter sets ActiveAfter field to given value.


### GetDefinitionHash

`func (o *V2KvLedger) GetDefinitionHash() string`

GetDefinitionHash returns the DefinitionHash field if non-nil, zero value otherwise.

### GetDefinitionHashOk

`func (o *V2KvLedger) GetDefinitionHashOk() (*string, bool)`

GetDefinitionHashOk returns a tuple with the DefinitionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionHash

`func (o *V2KvLedger) SetDefinitionHash(v string)`

SetDefinitionHash sets DefinitionHash field to given value.


### GetUpdatedAt

`func (o *V2KvLedger) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V2KvLedger) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V2KvLedger) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V2KvLedger) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetError

`func (o *V2KvLedger) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2KvLedger) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2KvLedger) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V2KvLedger) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDroppedAt

`func (o *V2KvLedger) GetDroppedAt() time.Time`

GetDroppedAt returns the DroppedAt field if non-nil, zero value otherwise.

### GetDroppedAtOk

`func (o *V2KvLedger) GetDroppedAtOk() (*time.Time, bool)`

GetDroppedAtOk returns a tuple with the DroppedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroppedAt

`func (o *V2KvLedger) SetDroppedAt(v time.Time)`

SetDroppedAt sets DroppedAt field to given value.

### HasDroppedAt

`func (o *V2KvLedger) HasDroppedAt() bool`

HasDroppedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


