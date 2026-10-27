# V2KvArchivePurgeReceipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | **string** | The canonical JSON text of the receipt. | 
**Signature** | **string** | Base64 DER ECDSA signature over the SHA-256 of the json bytes. | 
**KeyId** | **string** |  | 
**Algorithm** | **string** |  | 

## Methods

### NewV2KvArchivePurgeReceipt

`func NewV2KvArchivePurgeReceipt(json string, signature string, keyId string, algorithm string, ) *V2KvArchivePurgeReceipt`

NewV2KvArchivePurgeReceipt instantiates a new V2KvArchivePurgeReceipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvArchivePurgeReceiptWithDefaults

`func NewV2KvArchivePurgeReceiptWithDefaults() *V2KvArchivePurgeReceipt`

NewV2KvArchivePurgeReceiptWithDefaults instantiates a new V2KvArchivePurgeReceipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *V2KvArchivePurgeReceipt) GetJson() string`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *V2KvArchivePurgeReceipt) GetJsonOk() (*string, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *V2KvArchivePurgeReceipt) SetJson(v string)`

SetJson sets Json field to given value.


### GetSignature

`func (o *V2KvArchivePurgeReceipt) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *V2KvArchivePurgeReceipt) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *V2KvArchivePurgeReceipt) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetKeyId

`func (o *V2KvArchivePurgeReceipt) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *V2KvArchivePurgeReceipt) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *V2KvArchivePurgeReceipt) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetAlgorithm

`func (o *V2KvArchivePurgeReceipt) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *V2KvArchivePurgeReceipt) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *V2KvArchivePurgeReceipt) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


