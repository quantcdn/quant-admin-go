# V2KvLedgerSubjectEntries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ledger** | **string** |  | 
**SubjectState** | **string** |  | 
**Entries** | [**[]V2KvEntry**](V2KvEntry.md) | Empty when subjectState is shredded. | 
**Complete** | **bool** | False when the time budget ended first: send the request again with nextToken. | 
**NextToken** | Pointer to **string** | Present when more entries remain. | [optional] 
**Shreds** | Pointer to **[]string** | Present when subjectState is shredded: the shred ids (KV_ledgers_shreds_receipt). | [optional] 
**EarlierGenerationsShredded** | **int32** | The number of earlier key generations of this subject that were shredded. | 

## Methods

### NewV2KvLedgerSubjectEntries

`func NewV2KvLedgerSubjectEntries(ledger string, subjectState string, entries []V2KvEntry, complete bool, earlierGenerationsShredded int32, ) *V2KvLedgerSubjectEntries`

NewV2KvLedgerSubjectEntries instantiates a new V2KvLedgerSubjectEntries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerSubjectEntriesWithDefaults

`func NewV2KvLedgerSubjectEntriesWithDefaults() *V2KvLedgerSubjectEntries`

NewV2KvLedgerSubjectEntriesWithDefaults instantiates a new V2KvLedgerSubjectEntries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLedger

`func (o *V2KvLedgerSubjectEntries) GetLedger() string`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *V2KvLedgerSubjectEntries) GetLedgerOk() (*string, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *V2KvLedgerSubjectEntries) SetLedger(v string)`

SetLedger sets Ledger field to given value.


### GetSubjectState

`func (o *V2KvLedgerSubjectEntries) GetSubjectState() string`

GetSubjectState returns the SubjectState field if non-nil, zero value otherwise.

### GetSubjectStateOk

`func (o *V2KvLedgerSubjectEntries) GetSubjectStateOk() (*string, bool)`

GetSubjectStateOk returns a tuple with the SubjectState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectState

`func (o *V2KvLedgerSubjectEntries) SetSubjectState(v string)`

SetSubjectState sets SubjectState field to given value.


### GetEntries

`func (o *V2KvLedgerSubjectEntries) GetEntries() []V2KvEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *V2KvLedgerSubjectEntries) GetEntriesOk() (*[]V2KvEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *V2KvLedgerSubjectEntries) SetEntries(v []V2KvEntry)`

SetEntries sets Entries field to given value.


### GetComplete

`func (o *V2KvLedgerSubjectEntries) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *V2KvLedgerSubjectEntries) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *V2KvLedgerSubjectEntries) SetComplete(v bool)`

SetComplete sets Complete field to given value.


### GetNextToken

`func (o *V2KvLedgerSubjectEntries) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *V2KvLedgerSubjectEntries) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *V2KvLedgerSubjectEntries) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *V2KvLedgerSubjectEntries) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetShreds

`func (o *V2KvLedgerSubjectEntries) GetShreds() []string`

GetShreds returns the Shreds field if non-nil, zero value otherwise.

### GetShredsOk

`func (o *V2KvLedgerSubjectEntries) GetShredsOk() (*[]string, bool)`

GetShredsOk returns a tuple with the Shreds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShreds

`func (o *V2KvLedgerSubjectEntries) SetShreds(v []string)`

SetShreds sets Shreds field to given value.

### HasShreds

`func (o *V2KvLedgerSubjectEntries) HasShreds() bool`

HasShreds returns a boolean if a field has been set.

### GetEarlierGenerationsShredded

`func (o *V2KvLedgerSubjectEntries) GetEarlierGenerationsShredded() int32`

GetEarlierGenerationsShredded returns the EarlierGenerationsShredded field if non-nil, zero value otherwise.

### GetEarlierGenerationsShreddedOk

`func (o *V2KvLedgerSubjectEntries) GetEarlierGenerationsShreddedOk() (*int32, bool)`

GetEarlierGenerationsShreddedOk returns a tuple with the EarlierGenerationsShredded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlierGenerationsShredded

`func (o *V2KvLedgerSubjectEntries) SetEarlierGenerationsShredded(v int32)`

SetEarlierGenerationsShredded sets EarlierGenerationsShredded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


