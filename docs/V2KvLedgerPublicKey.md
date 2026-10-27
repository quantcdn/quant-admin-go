# V2KvLedgerPublicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyId** | **string** | The KMS key ARN. | 
**Algorithm** | **string** |  | 
**Curve** | **string** |  | 
**PublicKeyPem** | **string** | SubjectPublicKeyInfo, PEM. | 
**Keys** | Pointer to [**[]V2KvLedgerPublicKeyKeysInner**](V2KvLedgerPublicKeyKeysInner.md) | Current and historical verification keys. Select by the signature keyId. | [optional] 

## Methods

### NewV2KvLedgerPublicKey

`func NewV2KvLedgerPublicKey(keyId string, algorithm string, curve string, publicKeyPem string, ) *V2KvLedgerPublicKey`

NewV2KvLedgerPublicKey instantiates a new V2KvLedgerPublicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerPublicKeyWithDefaults

`func NewV2KvLedgerPublicKeyWithDefaults() *V2KvLedgerPublicKey`

NewV2KvLedgerPublicKeyWithDefaults instantiates a new V2KvLedgerPublicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyId

`func (o *V2KvLedgerPublicKey) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *V2KvLedgerPublicKey) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *V2KvLedgerPublicKey) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetAlgorithm

`func (o *V2KvLedgerPublicKey) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *V2KvLedgerPublicKey) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *V2KvLedgerPublicKey) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetCurve

`func (o *V2KvLedgerPublicKey) GetCurve() string`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *V2KvLedgerPublicKey) GetCurveOk() (*string, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *V2KvLedgerPublicKey) SetCurve(v string)`

SetCurve sets Curve field to given value.


### GetPublicKeyPem

`func (o *V2KvLedgerPublicKey) GetPublicKeyPem() string`

GetPublicKeyPem returns the PublicKeyPem field if non-nil, zero value otherwise.

### GetPublicKeyPemOk

`func (o *V2KvLedgerPublicKey) GetPublicKeyPemOk() (*string, bool)`

GetPublicKeyPemOk returns a tuple with the PublicKeyPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyPem

`func (o *V2KvLedgerPublicKey) SetPublicKeyPem(v string)`

SetPublicKeyPem sets PublicKeyPem field to given value.


### GetKeys

`func (o *V2KvLedgerPublicKey) GetKeys() []V2KvLedgerPublicKeyKeysInner`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *V2KvLedgerPublicKey) GetKeysOk() (*[]V2KvLedgerPublicKeyKeysInner, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *V2KvLedgerPublicKey) SetKeys(v []V2KvLedgerPublicKeyKeysInner)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *V2KvLedgerPublicKey) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


