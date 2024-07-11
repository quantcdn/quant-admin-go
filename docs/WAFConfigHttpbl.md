# WAFConfigHttpbl

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

### NewWAFConfigHttpbl

`func NewWAFConfigHttpbl(httpblEnabled bool, blockSuspicious bool, blockHarvester bool, blockSpam bool, blockSearchEngine bool, ) *WAFConfigHttpbl`

NewWAFConfigHttpbl instantiates a new WAFConfigHttpbl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWAFConfigHttpblWithDefaults

`func NewWAFConfigHttpblWithDefaults() *WAFConfigHttpbl`

NewWAFConfigHttpblWithDefaults instantiates a new WAFConfigHttpbl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpblEnabled

`func (o *WAFConfigHttpbl) GetHttpblEnabled() bool`

GetHttpblEnabled returns the HttpblEnabled field if non-nil, zero value otherwise.

### GetHttpblEnabledOk

`func (o *WAFConfigHttpbl) GetHttpblEnabledOk() (*bool, bool)`

GetHttpblEnabledOk returns a tuple with the HttpblEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpblEnabled

`func (o *WAFConfigHttpbl) SetHttpblEnabled(v bool)`

SetHttpblEnabled sets HttpblEnabled field to given value.


### GetApiKey

`func (o *WAFConfigHttpbl) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *WAFConfigHttpbl) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *WAFConfigHttpbl) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *WAFConfigHttpbl) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetBlockSuspicious

`func (o *WAFConfigHttpbl) GetBlockSuspicious() bool`

GetBlockSuspicious returns the BlockSuspicious field if non-nil, zero value otherwise.

### GetBlockSuspiciousOk

`func (o *WAFConfigHttpbl) GetBlockSuspiciousOk() (*bool, bool)`

GetBlockSuspiciousOk returns a tuple with the BlockSuspicious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSuspicious

`func (o *WAFConfigHttpbl) SetBlockSuspicious(v bool)`

SetBlockSuspicious sets BlockSuspicious field to given value.


### GetBlockHarvester

`func (o *WAFConfigHttpbl) GetBlockHarvester() bool`

GetBlockHarvester returns the BlockHarvester field if non-nil, zero value otherwise.

### GetBlockHarvesterOk

`func (o *WAFConfigHttpbl) GetBlockHarvesterOk() (*bool, bool)`

GetBlockHarvesterOk returns a tuple with the BlockHarvester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHarvester

`func (o *WAFConfigHttpbl) SetBlockHarvester(v bool)`

SetBlockHarvester sets BlockHarvester field to given value.


### GetBlockSpam

`func (o *WAFConfigHttpbl) GetBlockSpam() bool`

GetBlockSpam returns the BlockSpam field if non-nil, zero value otherwise.

### GetBlockSpamOk

`func (o *WAFConfigHttpbl) GetBlockSpamOk() (*bool, bool)`

GetBlockSpamOk returns a tuple with the BlockSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSpam

`func (o *WAFConfigHttpbl) SetBlockSpam(v bool)`

SetBlockSpam sets BlockSpam field to given value.


### GetBlockSearchEngine

`func (o *WAFConfigHttpbl) GetBlockSearchEngine() bool`

GetBlockSearchEngine returns the BlockSearchEngine field if non-nil, zero value otherwise.

### GetBlockSearchEngineOk

`func (o *WAFConfigHttpbl) GetBlockSearchEngineOk() (*bool, bool)`

GetBlockSearchEngineOk returns a tuple with the BlockSearchEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSearchEngine

`func (o *WAFConfigHttpbl) SetBlockSearchEngine(v bool)`

SetBlockSearchEngine sets BlockSearchEngine field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


