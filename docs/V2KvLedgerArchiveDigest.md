# V2KvLedgerArchiveDigest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | digests/&lt;org&gt;/&lt;store&gt;/&lt;period&gt;.json (a supplementary digest is &lt;period&gt;_&lt;n&gt;.json). | 
**Url** | **string** |  | 
**SigUrl** | Pointer to **string** | Presigned URL of &lt;key&gt;.sig; absent if the signature does not exist yet. | [optional] 
**Period** | **string** |  | 
**Name** | **string** |  | 

## Methods

### NewV2KvLedgerArchiveDigest

`func NewV2KvLedgerArchiveDigest(key string, url string, period string, name string, ) *V2KvLedgerArchiveDigest`

NewV2KvLedgerArchiveDigest instantiates a new V2KvLedgerArchiveDigest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerArchiveDigestWithDefaults

`func NewV2KvLedgerArchiveDigestWithDefaults() *V2KvLedgerArchiveDigest`

NewV2KvLedgerArchiveDigestWithDefaults instantiates a new V2KvLedgerArchiveDigest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *V2KvLedgerArchiveDigest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *V2KvLedgerArchiveDigest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *V2KvLedgerArchiveDigest) SetKey(v string)`

SetKey sets Key field to given value.


### GetUrl

`func (o *V2KvLedgerArchiveDigest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V2KvLedgerArchiveDigest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V2KvLedgerArchiveDigest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSigUrl

`func (o *V2KvLedgerArchiveDigest) GetSigUrl() string`

GetSigUrl returns the SigUrl field if non-nil, zero value otherwise.

### GetSigUrlOk

`func (o *V2KvLedgerArchiveDigest) GetSigUrlOk() (*string, bool)`

GetSigUrlOk returns a tuple with the SigUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigUrl

`func (o *V2KvLedgerArchiveDigest) SetSigUrl(v string)`

SetSigUrl sets SigUrl field to given value.

### HasSigUrl

`func (o *V2KvLedgerArchiveDigest) HasSigUrl() bool`

HasSigUrl returns a boolean if a field has been set.

### GetPeriod

`func (o *V2KvLedgerArchiveDigest) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *V2KvLedgerArchiveDigest) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *V2KvLedgerArchiveDigest) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### GetName

`func (o *V2KvLedgerArchiveDigest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2KvLedgerArchiveDigest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2KvLedgerArchiveDigest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


