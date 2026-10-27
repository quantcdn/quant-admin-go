# V2KvLedgerErasureShredReceiptDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | **string** | The RFC 8785 (JCS) text of the receipt that was signed. | 
**Signature** | **string** | Base64 DER ECDSA P-256 signature over the SHA-256 of the UTF-8 bytes of json. | 
**KeyId** | **string** | The KMS key ARN of the ledger key. | 
**Algorithm** | **string** |  | 
**Receipt** | [**V2KvLedgerErasureShredReceiptBody**](V2KvLedgerErasureShredReceiptBody.md) |  | 

## Methods

### NewV2KvLedgerErasureShredReceiptDocument

`func NewV2KvLedgerErasureShredReceiptDocument(json string, signature string, keyId string, algorithm string, receipt V2KvLedgerErasureShredReceiptBody, ) *V2KvLedgerErasureShredReceiptDocument`

NewV2KvLedgerErasureShredReceiptDocument instantiates a new V2KvLedgerErasureShredReceiptDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerErasureShredReceiptDocumentWithDefaults

`func NewV2KvLedgerErasureShredReceiptDocumentWithDefaults() *V2KvLedgerErasureShredReceiptDocument`

NewV2KvLedgerErasureShredReceiptDocumentWithDefaults instantiates a new V2KvLedgerErasureShredReceiptDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *V2KvLedgerErasureShredReceiptDocument) GetJson() string`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *V2KvLedgerErasureShredReceiptDocument) GetJsonOk() (*string, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *V2KvLedgerErasureShredReceiptDocument) SetJson(v string)`

SetJson sets Json field to given value.


### GetSignature

`func (o *V2KvLedgerErasureShredReceiptDocument) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *V2KvLedgerErasureShredReceiptDocument) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *V2KvLedgerErasureShredReceiptDocument) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetKeyId

`func (o *V2KvLedgerErasureShredReceiptDocument) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *V2KvLedgerErasureShredReceiptDocument) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *V2KvLedgerErasureShredReceiptDocument) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetAlgorithm

`func (o *V2KvLedgerErasureShredReceiptDocument) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *V2KvLedgerErasureShredReceiptDocument) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *V2KvLedgerErasureShredReceiptDocument) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetReceipt

`func (o *V2KvLedgerErasureShredReceiptDocument) GetReceipt() V2KvLedgerErasureShredReceiptBody`

GetReceipt returns the Receipt field if non-nil, zero value otherwise.

### GetReceiptOk

`func (o *V2KvLedgerErasureShredReceiptDocument) GetReceiptOk() (*V2KvLedgerErasureShredReceiptBody, bool)`

GetReceiptOk returns a tuple with the Receipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceipt

`func (o *V2KvLedgerErasureShredReceiptDocument) SetReceipt(v V2KvLedgerErasureShredReceiptBody)`

SetReceipt sets Receipt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


