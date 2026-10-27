# V2KvLedgerAppendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | **string** | A concrete collection (no *) that a declared ledger covers. At most 512 UTF-8 bytes. | 
**Value** | **map[string]interface{}** | The entry, a JSON object (not a list). At most 350,000 bytes as JSON. Top-level at is reserved (the KV service sets it) and $event is reserved for shred entries; a top-level field whose value is an object with a $enc or $shredded key fails VALIDATION. When the ledger has personal fields, value must hold the subject field (else LEDGER_SUBJECT_REQUIRED). | 
**IdempotencyKey** | Pointer to **string** | Optional. A retry with the same key and the same collection and value within 24 hours returns the first receipt with replayed true and appends nothing. The same key with another collection or value fails VALIDATION. A key is unique within the store; derive one key per append. Send it to make a lost-connection 503 retryable. | [optional] 

## Methods

### NewV2KvLedgerAppendRequest

`func NewV2KvLedgerAppendRequest(collection string, value map[string]interface{}, ) *V2KvLedgerAppendRequest`

NewV2KvLedgerAppendRequest instantiates a new V2KvLedgerAppendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerAppendRequestWithDefaults

`func NewV2KvLedgerAppendRequestWithDefaults() *V2KvLedgerAppendRequest`

NewV2KvLedgerAppendRequestWithDefaults instantiates a new V2KvLedgerAppendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *V2KvLedgerAppendRequest) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *V2KvLedgerAppendRequest) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *V2KvLedgerAppendRequest) SetCollection(v string)`

SetCollection sets Collection field to given value.


### GetValue

`func (o *V2KvLedgerAppendRequest) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V2KvLedgerAppendRequest) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V2KvLedgerAppendRequest) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.


### GetIdempotencyKey

`func (o *V2KvLedgerAppendRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *V2KvLedgerAppendRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *V2KvLedgerAppendRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *V2KvLedgerAppendRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


