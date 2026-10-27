# V2KvLedgerStatusAnchor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**LastDay** | Pointer to **string** |  | [optional] 
**FailingSince** | Pointer to **time.Time** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**BucketPolicy** | Pointer to **string** | customer-bucket only: the bucket policy (JSON) that lets Quant write the witness copies. | [optional] 

## Methods

### NewV2KvLedgerStatusAnchor

`func NewV2KvLedgerStatusAnchor(type_ string, ) *V2KvLedgerStatusAnchor`

NewV2KvLedgerStatusAnchor instantiates a new V2KvLedgerStatusAnchor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerStatusAnchorWithDefaults

`func NewV2KvLedgerStatusAnchorWithDefaults() *V2KvLedgerStatusAnchor`

NewV2KvLedgerStatusAnchorWithDefaults instantiates a new V2KvLedgerStatusAnchor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *V2KvLedgerStatusAnchor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V2KvLedgerStatusAnchor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V2KvLedgerStatusAnchor) SetType(v string)`

SetType sets Type field to given value.


### GetLastDay

`func (o *V2KvLedgerStatusAnchor) GetLastDay() string`

GetLastDay returns the LastDay field if non-nil, zero value otherwise.

### GetLastDayOk

`func (o *V2KvLedgerStatusAnchor) GetLastDayOk() (*string, bool)`

GetLastDayOk returns a tuple with the LastDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDay

`func (o *V2KvLedgerStatusAnchor) SetLastDay(v string)`

SetLastDay sets LastDay field to given value.

### HasLastDay

`func (o *V2KvLedgerStatusAnchor) HasLastDay() bool`

HasLastDay returns a boolean if a field has been set.

### GetFailingSince

`func (o *V2KvLedgerStatusAnchor) GetFailingSince() time.Time`

GetFailingSince returns the FailingSince field if non-nil, zero value otherwise.

### GetFailingSinceOk

`func (o *V2KvLedgerStatusAnchor) GetFailingSinceOk() (*time.Time, bool)`

GetFailingSinceOk returns a tuple with the FailingSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailingSince

`func (o *V2KvLedgerStatusAnchor) SetFailingSince(v time.Time)`

SetFailingSince sets FailingSince field to given value.

### HasFailingSince

`func (o *V2KvLedgerStatusAnchor) HasFailingSince() bool`

HasFailingSince returns a boolean if a field has been set.

### GetError

`func (o *V2KvLedgerStatusAnchor) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2KvLedgerStatusAnchor) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2KvLedgerStatusAnchor) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V2KvLedgerStatusAnchor) HasError() bool`

HasError returns a boolean if a field has been set.

### GetBucketPolicy

`func (o *V2KvLedgerStatusAnchor) GetBucketPolicy() string`

GetBucketPolicy returns the BucketPolicy field if non-nil, zero value otherwise.

### GetBucketPolicyOk

`func (o *V2KvLedgerStatusAnchor) GetBucketPolicyOk() (*string, bool)`

GetBucketPolicyOk returns a tuple with the BucketPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketPolicy

`func (o *V2KvLedgerStatusAnchor) SetBucketPolicy(v string)`

SetBucketPolicy sets BucketPolicy field to given value.

### HasBucketPolicy

`func (o *V2KvLedgerStatusAnchor) HasBucketPolicy() bool`

HasBucketPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


