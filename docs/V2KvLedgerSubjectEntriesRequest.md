# V2KvLedgerSubjectEntriesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectId** | [**V2KvLedgerSubjectId**](V2KvLedgerSubjectId.md) |  | 
**Collection** | Pointer to **string** | Read only this collection of the ledger. | [optional] 
**Since** | Pointer to **time.Time** | Only entries at or after this time. | [optional] 
**Until** | Pointer to **time.Time** | Only entries before this time. | [optional] 
**Limit** | Pointer to **int32** | Entries per page, 1 to 1000. The default is 100. | [optional] 
**NextToken** | Pointer to **string** | The nextToken of the previous page. | [optional] 
**Personal** | Pointer to **string** | sealed (the default): personal fields read \&quot;[PERSONAL]\&quot; and the subject field shows the subject token. decrypted: personal fields in clear; needs the ledger_reveal permission and is logged as kv.ledger.reveal. | [optional] 

## Methods

### NewV2KvLedgerSubjectEntriesRequest

`func NewV2KvLedgerSubjectEntriesRequest(subjectId V2KvLedgerSubjectId, ) *V2KvLedgerSubjectEntriesRequest`

NewV2KvLedgerSubjectEntriesRequest instantiates a new V2KvLedgerSubjectEntriesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerSubjectEntriesRequestWithDefaults

`func NewV2KvLedgerSubjectEntriesRequestWithDefaults() *V2KvLedgerSubjectEntriesRequest`

NewV2KvLedgerSubjectEntriesRequestWithDefaults instantiates a new V2KvLedgerSubjectEntriesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectId

`func (o *V2KvLedgerSubjectEntriesRequest) GetSubjectId() V2KvLedgerSubjectId`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *V2KvLedgerSubjectEntriesRequest) GetSubjectIdOk() (*V2KvLedgerSubjectId, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *V2KvLedgerSubjectEntriesRequest) SetSubjectId(v V2KvLedgerSubjectId)`

SetSubjectId sets SubjectId field to given value.


### GetCollection

`func (o *V2KvLedgerSubjectEntriesRequest) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *V2KvLedgerSubjectEntriesRequest) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *V2KvLedgerSubjectEntriesRequest) SetCollection(v string)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *V2KvLedgerSubjectEntriesRequest) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetSince

`func (o *V2KvLedgerSubjectEntriesRequest) GetSince() time.Time`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *V2KvLedgerSubjectEntriesRequest) GetSinceOk() (*time.Time, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *V2KvLedgerSubjectEntriesRequest) SetSince(v time.Time)`

SetSince sets Since field to given value.

### HasSince

`func (o *V2KvLedgerSubjectEntriesRequest) HasSince() bool`

HasSince returns a boolean if a field has been set.

### GetUntil

`func (o *V2KvLedgerSubjectEntriesRequest) GetUntil() time.Time`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *V2KvLedgerSubjectEntriesRequest) GetUntilOk() (*time.Time, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *V2KvLedgerSubjectEntriesRequest) SetUntil(v time.Time)`

SetUntil sets Until field to given value.

### HasUntil

`func (o *V2KvLedgerSubjectEntriesRequest) HasUntil() bool`

HasUntil returns a boolean if a field has been set.

### GetLimit

`func (o *V2KvLedgerSubjectEntriesRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *V2KvLedgerSubjectEntriesRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *V2KvLedgerSubjectEntriesRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *V2KvLedgerSubjectEntriesRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNextToken

`func (o *V2KvLedgerSubjectEntriesRequest) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *V2KvLedgerSubjectEntriesRequest) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *V2KvLedgerSubjectEntriesRequest) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *V2KvLedgerSubjectEntriesRequest) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetPersonal

`func (o *V2KvLedgerSubjectEntriesRequest) GetPersonal() string`

GetPersonal returns the Personal field if non-nil, zero value otherwise.

### GetPersonalOk

`func (o *V2KvLedgerSubjectEntriesRequest) GetPersonalOk() (*string, bool)`

GetPersonalOk returns a tuple with the Personal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonal

`func (o *V2KvLedgerSubjectEntriesRequest) SetPersonal(v string)`

SetPersonal sets Personal field to given value.

### HasPersonal

`func (o *V2KvLedgerSubjectEntriesRequest) HasPersonal() bool`

HasPersonal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


