# SecretWriteResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The secret name | 
**Warnings** | Pointer to **[]string** | Notices from Quant Cloud about the change. Present only when there is at least one. | [optional] 

## Methods

### NewSecretWriteResult

`func NewSecretWriteResult(key string, ) *SecretWriteResult`

NewSecretWriteResult instantiates a new SecretWriteResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretWriteResultWithDefaults

`func NewSecretWriteResultWithDefaults() *SecretWriteResult`

NewSecretWriteResultWithDefaults instantiates a new SecretWriteResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SecretWriteResult) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SecretWriteResult) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SecretWriteResult) SetKey(v string)`

SetKey sets Key field to given value.


### GetWarnings

`func (o *SecretWriteResult) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SecretWriteResult) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SecretWriteResult) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SecretWriteResult) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


