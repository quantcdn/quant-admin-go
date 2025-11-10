# WafConfigHttpbl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpblEnabled** | Pointer to **bool** | Enable HTTP:BL | [optional] [default to false]
**BlockSuspicious** | Pointer to **bool** | Block suspicious IPs | [optional] [default to false]
**BlockHarvester** | Pointer to **bool** | Block email harvesters | [optional] [default to false]
**BlockSpam** | Pointer to **bool** | Block spam sources | [optional] [default to false]
**BlockSearchEngine** | Pointer to **bool** | Block search engines | [optional] [default to false]
**HttpblKey** | Pointer to **string** | HTTP:BL API key | [optional] 

## Methods

### NewWafConfigHttpbl

`func NewWafConfigHttpbl() *WafConfigHttpbl`

NewWafConfigHttpbl instantiates a new WafConfigHttpbl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafConfigHttpblWithDefaults

`func NewWafConfigHttpblWithDefaults() *WafConfigHttpbl`

NewWafConfigHttpblWithDefaults instantiates a new WafConfigHttpbl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpblEnabled

`func (o *WafConfigHttpbl) GetHttpblEnabled() bool`

GetHttpblEnabled returns the HttpblEnabled field if non-nil, zero value otherwise.

### GetHttpblEnabledOk

`func (o *WafConfigHttpbl) GetHttpblEnabledOk() (*bool, bool)`

GetHttpblEnabledOk returns a tuple with the HttpblEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpblEnabled

`func (o *WafConfigHttpbl) SetHttpblEnabled(v bool)`

SetHttpblEnabled sets HttpblEnabled field to given value.

### HasHttpblEnabled

`func (o *WafConfigHttpbl) HasHttpblEnabled() bool`

HasHttpblEnabled returns a boolean if a field has been set.

### GetBlockSuspicious

`func (o *WafConfigHttpbl) GetBlockSuspicious() bool`

GetBlockSuspicious returns the BlockSuspicious field if non-nil, zero value otherwise.

### GetBlockSuspiciousOk

`func (o *WafConfigHttpbl) GetBlockSuspiciousOk() (*bool, bool)`

GetBlockSuspiciousOk returns a tuple with the BlockSuspicious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSuspicious

`func (o *WafConfigHttpbl) SetBlockSuspicious(v bool)`

SetBlockSuspicious sets BlockSuspicious field to given value.

### HasBlockSuspicious

`func (o *WafConfigHttpbl) HasBlockSuspicious() bool`

HasBlockSuspicious returns a boolean if a field has been set.

### GetBlockHarvester

`func (o *WafConfigHttpbl) GetBlockHarvester() bool`

GetBlockHarvester returns the BlockHarvester field if non-nil, zero value otherwise.

### GetBlockHarvesterOk

`func (o *WafConfigHttpbl) GetBlockHarvesterOk() (*bool, bool)`

GetBlockHarvesterOk returns a tuple with the BlockHarvester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHarvester

`func (o *WafConfigHttpbl) SetBlockHarvester(v bool)`

SetBlockHarvester sets BlockHarvester field to given value.

### HasBlockHarvester

`func (o *WafConfigHttpbl) HasBlockHarvester() bool`

HasBlockHarvester returns a boolean if a field has been set.

### GetBlockSpam

`func (o *WafConfigHttpbl) GetBlockSpam() bool`

GetBlockSpam returns the BlockSpam field if non-nil, zero value otherwise.

### GetBlockSpamOk

`func (o *WafConfigHttpbl) GetBlockSpamOk() (*bool, bool)`

GetBlockSpamOk returns a tuple with the BlockSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSpam

`func (o *WafConfigHttpbl) SetBlockSpam(v bool)`

SetBlockSpam sets BlockSpam field to given value.

### HasBlockSpam

`func (o *WafConfigHttpbl) HasBlockSpam() bool`

HasBlockSpam returns a boolean if a field has been set.

### GetBlockSearchEngine

`func (o *WafConfigHttpbl) GetBlockSearchEngine() bool`

GetBlockSearchEngine returns the BlockSearchEngine field if non-nil, zero value otherwise.

### GetBlockSearchEngineOk

`func (o *WafConfigHttpbl) GetBlockSearchEngineOk() (*bool, bool)`

GetBlockSearchEngineOk returns a tuple with the BlockSearchEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSearchEngine

`func (o *WafConfigHttpbl) SetBlockSearchEngine(v bool)`

SetBlockSearchEngine sets BlockSearchEngine field to given value.

### HasBlockSearchEngine

`func (o *WafConfigHttpbl) HasBlockSearchEngine() bool`

HasBlockSearchEngine returns a boolean if a field has been set.

### GetHttpblKey

`func (o *WafConfigHttpbl) GetHttpblKey() string`

GetHttpblKey returns the HttpblKey field if non-nil, zero value otherwise.

### GetHttpblKeyOk

`func (o *WafConfigHttpbl) GetHttpblKeyOk() (*string, bool)`

GetHttpblKeyOk returns a tuple with the HttpblKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpblKey

`func (o *WafConfigHttpbl) SetHttpblKey(v string)`

SetHttpblKey sets HttpblKey field to given value.

### HasHttpblKey

`func (o *WafConfigHttpbl) HasHttpblKey() bool`

HasHttpblKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


