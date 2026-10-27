# V2KvLedgerVerificationReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Report** | [**V2KvLedgerVerificationReportReport**](V2KvLedgerVerificationReportReport.md) |  | 
**Json** | **string** | The exact signed text. | 
**Signature** | **string** |  | 
**KeyId** | **string** |  | 
**Algorithm** | **string** |  | 

## Methods

### NewV2KvLedgerVerificationReport

`func NewV2KvLedgerVerificationReport(report V2KvLedgerVerificationReportReport, json string, signature string, keyId string, algorithm string, ) *V2KvLedgerVerificationReport`

NewV2KvLedgerVerificationReport instantiates a new V2KvLedgerVerificationReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerVerificationReportWithDefaults

`func NewV2KvLedgerVerificationReportWithDefaults() *V2KvLedgerVerificationReport`

NewV2KvLedgerVerificationReportWithDefaults instantiates a new V2KvLedgerVerificationReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReport

`func (o *V2KvLedgerVerificationReport) GetReport() V2KvLedgerVerificationReportReport`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *V2KvLedgerVerificationReport) GetReportOk() (*V2KvLedgerVerificationReportReport, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *V2KvLedgerVerificationReport) SetReport(v V2KvLedgerVerificationReportReport)`

SetReport sets Report field to given value.


### GetJson

`func (o *V2KvLedgerVerificationReport) GetJson() string`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *V2KvLedgerVerificationReport) GetJsonOk() (*string, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *V2KvLedgerVerificationReport) SetJson(v string)`

SetJson sets Json field to given value.


### GetSignature

`func (o *V2KvLedgerVerificationReport) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *V2KvLedgerVerificationReport) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *V2KvLedgerVerificationReport) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetKeyId

`func (o *V2KvLedgerVerificationReport) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *V2KvLedgerVerificationReport) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *V2KvLedgerVerificationReport) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetAlgorithm

`func (o *V2KvLedgerVerificationReport) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *V2KvLedgerVerificationReport) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *V2KvLedgerVerificationReport) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


