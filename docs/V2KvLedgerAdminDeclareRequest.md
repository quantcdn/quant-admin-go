# V2KvLedgerAdminDeclareRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique in the store. | 
**Collection** | **string** | A collection name, or a pattern where * as a whole segment matches exactly one segment. Must not overlap another ledger. At most 512 UTF-8 bytes; every concrete collection appended to the ledger is at most 512 bytes too. | 
**Chains** | **int32** | Parallel hash chains (shards) per collection. | 
**RetentionDays** | **int32** | S3 Object Lock retention of the archive, in days. The ledger cannot be dropped before it ends. | 
**LockMode** | **string** | S3 Object Lock mode of the archive. governance can become compliance later, never back. | 
**PersonalFields** | Pointer to **[]string** | Top-level value fields that are encrypted per data subject, so a shred (KV_ledgers_shred) can erase them. Needs subjectField (else 400 LEDGER_SUBJECT_REQUIRED). Any field name except at. An index cannot use a personal field (400 INDEX_FIELD_ENCRYPTED). | [optional] 
**SubjectField** | Pointer to **string** | The top-level value field that names the data subject. It is not a personal field. May be set without personal fields; such a ledger is pseudonymised and can be shredded. $subject and $subjectTag are reserved value field names. | [optional] 
**SubjectMode** | Pointer to **string** | How the subject field is stored. pseudonym (the default): a random token, and after a shred nobody can link the remaining entries to the person. clear: the value as written; after a shred the id and the activity stay. Needs subjectField. Cannot change. | [optional] 
**SubjectCase** | Pointer to **string** | exact (the default): ids that differ in case are different subjects. fold: case is ignored; use it for email addresses and usernames. Needs subjectField. Cannot change. | [optional] 
**Anchor** | Pointer to [**V2KvLedgerAnchor**](V2KvLedgerAnchor.md) |  | [optional] 
**WriterPolicy** | Pointer to [**V2KvLedgerAdminWriterPolicy**](V2KvLedgerAdminWriterPolicy.md) |  | [optional] 

## Methods

### NewV2KvLedgerAdminDeclareRequest

`func NewV2KvLedgerAdminDeclareRequest(name string, collection string, chains int32, retentionDays int32, lockMode string, ) *V2KvLedgerAdminDeclareRequest`

NewV2KvLedgerAdminDeclareRequest instantiates a new V2KvLedgerAdminDeclareRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerAdminDeclareRequestWithDefaults

`func NewV2KvLedgerAdminDeclareRequestWithDefaults() *V2KvLedgerAdminDeclareRequest`

NewV2KvLedgerAdminDeclareRequestWithDefaults instantiates a new V2KvLedgerAdminDeclareRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2KvLedgerAdminDeclareRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2KvLedgerAdminDeclareRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2KvLedgerAdminDeclareRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCollection

`func (o *V2KvLedgerAdminDeclareRequest) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *V2KvLedgerAdminDeclareRequest) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *V2KvLedgerAdminDeclareRequest) SetCollection(v string)`

SetCollection sets Collection field to given value.


### GetChains

`func (o *V2KvLedgerAdminDeclareRequest) GetChains() int32`

GetChains returns the Chains field if non-nil, zero value otherwise.

### GetChainsOk

`func (o *V2KvLedgerAdminDeclareRequest) GetChainsOk() (*int32, bool)`

GetChainsOk returns a tuple with the Chains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChains

`func (o *V2KvLedgerAdminDeclareRequest) SetChains(v int32)`

SetChains sets Chains field to given value.


### GetRetentionDays

`func (o *V2KvLedgerAdminDeclareRequest) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *V2KvLedgerAdminDeclareRequest) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *V2KvLedgerAdminDeclareRequest) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.


### GetLockMode

`func (o *V2KvLedgerAdminDeclareRequest) GetLockMode() string`

GetLockMode returns the LockMode field if non-nil, zero value otherwise.

### GetLockModeOk

`func (o *V2KvLedgerAdminDeclareRequest) GetLockModeOk() (*string, bool)`

GetLockModeOk returns a tuple with the LockMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockMode

`func (o *V2KvLedgerAdminDeclareRequest) SetLockMode(v string)`

SetLockMode sets LockMode field to given value.


### GetPersonalFields

`func (o *V2KvLedgerAdminDeclareRequest) GetPersonalFields() []string`

GetPersonalFields returns the PersonalFields field if non-nil, zero value otherwise.

### GetPersonalFieldsOk

`func (o *V2KvLedgerAdminDeclareRequest) GetPersonalFieldsOk() (*[]string, bool)`

GetPersonalFieldsOk returns a tuple with the PersonalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalFields

`func (o *V2KvLedgerAdminDeclareRequest) SetPersonalFields(v []string)`

SetPersonalFields sets PersonalFields field to given value.

### HasPersonalFields

`func (o *V2KvLedgerAdminDeclareRequest) HasPersonalFields() bool`

HasPersonalFields returns a boolean if a field has been set.

### GetSubjectField

`func (o *V2KvLedgerAdminDeclareRequest) GetSubjectField() string`

GetSubjectField returns the SubjectField field if non-nil, zero value otherwise.

### GetSubjectFieldOk

`func (o *V2KvLedgerAdminDeclareRequest) GetSubjectFieldOk() (*string, bool)`

GetSubjectFieldOk returns a tuple with the SubjectField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectField

`func (o *V2KvLedgerAdminDeclareRequest) SetSubjectField(v string)`

SetSubjectField sets SubjectField field to given value.

### HasSubjectField

`func (o *V2KvLedgerAdminDeclareRequest) HasSubjectField() bool`

HasSubjectField returns a boolean if a field has been set.

### GetSubjectMode

`func (o *V2KvLedgerAdminDeclareRequest) GetSubjectMode() string`

GetSubjectMode returns the SubjectMode field if non-nil, zero value otherwise.

### GetSubjectModeOk

`func (o *V2KvLedgerAdminDeclareRequest) GetSubjectModeOk() (*string, bool)`

GetSubjectModeOk returns a tuple with the SubjectMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectMode

`func (o *V2KvLedgerAdminDeclareRequest) SetSubjectMode(v string)`

SetSubjectMode sets SubjectMode field to given value.

### HasSubjectMode

`func (o *V2KvLedgerAdminDeclareRequest) HasSubjectMode() bool`

HasSubjectMode returns a boolean if a field has been set.

### GetSubjectCase

`func (o *V2KvLedgerAdminDeclareRequest) GetSubjectCase() string`

GetSubjectCase returns the SubjectCase field if non-nil, zero value otherwise.

### GetSubjectCaseOk

`func (o *V2KvLedgerAdminDeclareRequest) GetSubjectCaseOk() (*string, bool)`

GetSubjectCaseOk returns a tuple with the SubjectCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCase

`func (o *V2KvLedgerAdminDeclareRequest) SetSubjectCase(v string)`

SetSubjectCase sets SubjectCase field to given value.

### HasSubjectCase

`func (o *V2KvLedgerAdminDeclareRequest) HasSubjectCase() bool`

HasSubjectCase returns a boolean if a field has been set.

### GetAnchor

`func (o *V2KvLedgerAdminDeclareRequest) GetAnchor() V2KvLedgerAnchor`

GetAnchor returns the Anchor field if non-nil, zero value otherwise.

### GetAnchorOk

`func (o *V2KvLedgerAdminDeclareRequest) GetAnchorOk() (*V2KvLedgerAnchor, bool)`

GetAnchorOk returns a tuple with the Anchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchor

`func (o *V2KvLedgerAdminDeclareRequest) SetAnchor(v V2KvLedgerAnchor)`

SetAnchor sets Anchor field to given value.

### HasAnchor

`func (o *V2KvLedgerAdminDeclareRequest) HasAnchor() bool`

HasAnchor returns a boolean if a field has been set.

### GetWriterPolicy

`func (o *V2KvLedgerAdminDeclareRequest) GetWriterPolicy() V2KvLedgerAdminWriterPolicy`

GetWriterPolicy returns the WriterPolicy field if non-nil, zero value otherwise.

### GetWriterPolicyOk

`func (o *V2KvLedgerAdminDeclareRequest) GetWriterPolicyOk() (*V2KvLedgerAdminWriterPolicy, bool)`

GetWriterPolicyOk returns a tuple with the WriterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriterPolicy

`func (o *V2KvLedgerAdminDeclareRequest) SetWriterPolicy(v V2KvLedgerAdminWriterPolicy)`

SetWriterPolicy sets WriterPolicy field to given value.

### HasWriterPolicy

`func (o *V2KvLedgerAdminDeclareRequest) HasWriterPolicy() bool`

HasWriterPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


