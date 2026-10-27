# V2KvLedgerAnchor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**BucketArn** | Pointer to **string** | customer-bucket only (required on declare). | [optional] 
**Prefix** | Pointer to **string** | customer-bucket only, optional. Anchor keys are &lt;prefix&gt;&lt;org&gt;/&lt;store&gt;/…. | [optional] [default to ""]
**AccountId** | Pointer to **string** | customer-bucket only (required on declare): the AWS account that owns the bucket. Every anchor write sends it as ExpectedBucketOwner. | [optional] 
**TsaUrl** | Pointer to **string** | rfc3161 only. | [optional] 

## Methods

### NewV2KvLedgerAnchor

`func NewV2KvLedgerAnchor(type_ string, ) *V2KvLedgerAnchor`

NewV2KvLedgerAnchor instantiates a new V2KvLedgerAnchor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerAnchorWithDefaults

`func NewV2KvLedgerAnchorWithDefaults() *V2KvLedgerAnchor`

NewV2KvLedgerAnchorWithDefaults instantiates a new V2KvLedgerAnchor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *V2KvLedgerAnchor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V2KvLedgerAnchor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V2KvLedgerAnchor) SetType(v string)`

SetType sets Type field to given value.


### GetBucketArn

`func (o *V2KvLedgerAnchor) GetBucketArn() string`

GetBucketArn returns the BucketArn field if non-nil, zero value otherwise.

### GetBucketArnOk

`func (o *V2KvLedgerAnchor) GetBucketArnOk() (*string, bool)`

GetBucketArnOk returns a tuple with the BucketArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketArn

`func (o *V2KvLedgerAnchor) SetBucketArn(v string)`

SetBucketArn sets BucketArn field to given value.

### HasBucketArn

`func (o *V2KvLedgerAnchor) HasBucketArn() bool`

HasBucketArn returns a boolean if a field has been set.

### GetPrefix

`func (o *V2KvLedgerAnchor) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *V2KvLedgerAnchor) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *V2KvLedgerAnchor) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *V2KvLedgerAnchor) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetAccountId

`func (o *V2KvLedgerAnchor) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *V2KvLedgerAnchor) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *V2KvLedgerAnchor) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *V2KvLedgerAnchor) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetTsaUrl

`func (o *V2KvLedgerAnchor) GetTsaUrl() string`

GetTsaUrl returns the TsaUrl field if non-nil, zero value otherwise.

### GetTsaUrlOk

`func (o *V2KvLedgerAnchor) GetTsaUrlOk() (*string, bool)`

GetTsaUrlOk returns a tuple with the TsaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsaUrl

`func (o *V2KvLedgerAnchor) SetTsaUrl(v string)`

SetTsaUrl sets TsaUrl field to given value.

### HasTsaUrl

`func (o *V2KvLedgerAnchor) HasTsaUrl() bool`

HasTsaUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


