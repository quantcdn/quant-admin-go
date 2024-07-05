# Httpbl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpblEnabled** | **bool** |  | [default to false]
**ApiKey** | Pointer to **string** |  | [optional] 
**BlockSuspicious** | **bool** |  | [default to false]
**BlockHarvester** | **bool** |  | [default to false]
**BlockSpam** | **bool** |  | [default to false]
**BlockSearchEngine** | **bool** |  | [default to false]

## Methods

### NewHttpbl

`func NewHttpbl(httpblEnabled bool, blockSuspicious bool, blockHarvester bool, blockSpam bool, blockSearchEngine bool, ) *Httpbl`

NewHttpbl instantiates a new Httpbl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpblWithDefaults

`func NewHttpblWithDefaults() *Httpbl`

NewHttpblWithDefaults instantiates a new Httpbl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpblEnabled

`func (o *Httpbl) GetHttpblEnabled() bool`

GetHttpblEnabled returns the HttpblEnabled field if non-nil, zero value otherwise.

### GetHttpblEnabledOk

`func (o *Httpbl) GetHttpblEnabledOk() (*bool, bool)`

GetHttpblEnabledOk returns a tuple with the HttpblEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpblEnabled

`func (o *Httpbl) SetHttpblEnabled(v bool)`

SetHttpblEnabled sets HttpblEnabled field to given value.


### GetApiKey

`func (o *Httpbl) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *Httpbl) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *Httpbl) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *Httpbl) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetBlockSuspicious

`func (o *Httpbl) GetBlockSuspicious() bool`

GetBlockSuspicious returns the BlockSuspicious field if non-nil, zero value otherwise.

### GetBlockSuspiciousOk

`func (o *Httpbl) GetBlockSuspiciousOk() (*bool, bool)`

GetBlockSuspiciousOk returns a tuple with the BlockSuspicious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSuspicious

`func (o *Httpbl) SetBlockSuspicious(v bool)`

SetBlockSuspicious sets BlockSuspicious field to given value.


### GetBlockHarvester

`func (o *Httpbl) GetBlockHarvester() bool`

GetBlockHarvester returns the BlockHarvester field if non-nil, zero value otherwise.

### GetBlockHarvesterOk

`func (o *Httpbl) GetBlockHarvesterOk() (*bool, bool)`

GetBlockHarvesterOk returns a tuple with the BlockHarvester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHarvester

`func (o *Httpbl) SetBlockHarvester(v bool)`

SetBlockHarvester sets BlockHarvester field to given value.


### GetBlockSpam

`func (o *Httpbl) GetBlockSpam() bool`

GetBlockSpam returns the BlockSpam field if non-nil, zero value otherwise.

### GetBlockSpamOk

`func (o *Httpbl) GetBlockSpamOk() (*bool, bool)`

GetBlockSpamOk returns a tuple with the BlockSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSpam

`func (o *Httpbl) SetBlockSpam(v bool)`

SetBlockSpam sets BlockSpam field to given value.


### GetBlockSearchEngine

`func (o *Httpbl) GetBlockSearchEngine() bool`

GetBlockSearchEngine returns the BlockSearchEngine field if non-nil, zero value otherwise.

### GetBlockSearchEngineOk

`func (o *Httpbl) GetBlockSearchEngineOk() (*bool, bool)`

GetBlockSearchEngineOk returns a tuple with the BlockSearchEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSearchEngine

`func (o *Httpbl) SetBlockSearchEngine(v bool)`

SetBlockSearchEngine sets BlockSearchEngine field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


