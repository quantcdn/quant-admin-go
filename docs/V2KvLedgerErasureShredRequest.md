# V2KvLedgerErasureShredRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectId** | [**V2KvLedgerSubjectId**](V2KvLedgerSubjectId.md) |  | 
**Reason** | **string** | Why the subject is shredded. | 
**Authority** | **string** | The legal authority, for example Records authority 2023/00456 class 12. | 
**RequestedBy** | **string** | Who asked for the shred. | 

## Methods

### NewV2KvLedgerErasureShredRequest

`func NewV2KvLedgerErasureShredRequest(subjectId V2KvLedgerSubjectId, reason string, authority string, requestedBy string, ) *V2KvLedgerErasureShredRequest`

NewV2KvLedgerErasureShredRequest instantiates a new V2KvLedgerErasureShredRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerErasureShredRequestWithDefaults

`func NewV2KvLedgerErasureShredRequestWithDefaults() *V2KvLedgerErasureShredRequest`

NewV2KvLedgerErasureShredRequestWithDefaults instantiates a new V2KvLedgerErasureShredRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectId

`func (o *V2KvLedgerErasureShredRequest) GetSubjectId() V2KvLedgerSubjectId`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *V2KvLedgerErasureShredRequest) GetSubjectIdOk() (*V2KvLedgerSubjectId, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *V2KvLedgerErasureShredRequest) SetSubjectId(v V2KvLedgerSubjectId)`

SetSubjectId sets SubjectId field to given value.


### GetReason

`func (o *V2KvLedgerErasureShredRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V2KvLedgerErasureShredRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V2KvLedgerErasureShredRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetAuthority

`func (o *V2KvLedgerErasureShredRequest) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *V2KvLedgerErasureShredRequest) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *V2KvLedgerErasureShredRequest) SetAuthority(v string)`

SetAuthority sets Authority field to given value.


### GetRequestedBy

`func (o *V2KvLedgerErasureShredRequest) GetRequestedBy() string`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *V2KvLedgerErasureShredRequest) GetRequestedByOk() (*string, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *V2KvLedgerErasureShredRequest) SetRequestedBy(v string)`

SetRequestedBy sets RequestedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


