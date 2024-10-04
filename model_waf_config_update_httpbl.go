/*
QuantCDN Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the WAFConfigUpdateHttpbl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WAFConfigUpdateHttpbl{}

// WAFConfigUpdateHttpbl struct for WAFConfigUpdateHttpbl
type WAFConfigUpdateHttpbl struct {
	HttpblEnabled *bool `json:"httpbl_enabled,omitempty"`
	ApiKey *string `json:"api_key,omitempty"`
	BlockSuspicious *bool `json:"block_suspicious,omitempty"`
	BlockHarvester *bool `json:"block_harvester,omitempty"`
	BlockSpam *bool `json:"block_spam,omitempty"`
	BlockSearchEngine *bool `json:"block_search_engine,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WAFConfigUpdateHttpbl WAFConfigUpdateHttpbl

// NewWAFConfigUpdateHttpbl instantiates a new WAFConfigUpdateHttpbl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWAFConfigUpdateHttpbl() *WAFConfigUpdateHttpbl {
	this := WAFConfigUpdateHttpbl{}
	var httpblEnabled bool = false
	this.HttpblEnabled = &httpblEnabled
	var blockSuspicious bool = false
	this.BlockSuspicious = &blockSuspicious
	var blockHarvester bool = false
	this.BlockHarvester = &blockHarvester
	var blockSpam bool = false
	this.BlockSpam = &blockSpam
	var blockSearchEngine bool = false
	this.BlockSearchEngine = &blockSearchEngine
	return &this
}

// NewWAFConfigUpdateHttpblWithDefaults instantiates a new WAFConfigUpdateHttpbl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWAFConfigUpdateHttpblWithDefaults() *WAFConfigUpdateHttpbl {
	this := WAFConfigUpdateHttpbl{}
	var httpblEnabled bool = false
	this.HttpblEnabled = &httpblEnabled
	var blockSuspicious bool = false
	this.BlockSuspicious = &blockSuspicious
	var blockHarvester bool = false
	this.BlockHarvester = &blockHarvester
	var blockSpam bool = false
	this.BlockSpam = &blockSpam
	var blockSearchEngine bool = false
	this.BlockSearchEngine = &blockSearchEngine
	return &this
}

// GetHttpblEnabled returns the HttpblEnabled field value if set, zero value otherwise.
func (o *WAFConfigUpdateHttpbl) GetHttpblEnabled() bool {
	if o == nil || IsNil(o.HttpblEnabled) {
		var ret bool
		return ret
	}
	return *o.HttpblEnabled
}

// GetHttpblEnabledOk returns a tuple with the HttpblEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdateHttpbl) GetHttpblEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.HttpblEnabled) {
		return nil, false
	}
	return o.HttpblEnabled, true
}

// HasHttpblEnabled returns a boolean if a field has been set.
func (o *WAFConfigUpdateHttpbl) HasHttpblEnabled() bool {
	if o != nil && !IsNil(o.HttpblEnabled) {
		return true
	}

	return false
}

// SetHttpblEnabled gets a reference to the given bool and assigns it to the HttpblEnabled field.
func (o *WAFConfigUpdateHttpbl) SetHttpblEnabled(v bool) {
	o.HttpblEnabled = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *WAFConfigUpdateHttpbl) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdateHttpbl) GetApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *WAFConfigUpdateHttpbl) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *WAFConfigUpdateHttpbl) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetBlockSuspicious returns the BlockSuspicious field value if set, zero value otherwise.
func (o *WAFConfigUpdateHttpbl) GetBlockSuspicious() bool {
	if o == nil || IsNil(o.BlockSuspicious) {
		var ret bool
		return ret
	}
	return *o.BlockSuspicious
}

// GetBlockSuspiciousOk returns a tuple with the BlockSuspicious field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdateHttpbl) GetBlockSuspiciousOk() (*bool, bool) {
	if o == nil || IsNil(o.BlockSuspicious) {
		return nil, false
	}
	return o.BlockSuspicious, true
}

// HasBlockSuspicious returns a boolean if a field has been set.
func (o *WAFConfigUpdateHttpbl) HasBlockSuspicious() bool {
	if o != nil && !IsNil(o.BlockSuspicious) {
		return true
	}

	return false
}

// SetBlockSuspicious gets a reference to the given bool and assigns it to the BlockSuspicious field.
func (o *WAFConfigUpdateHttpbl) SetBlockSuspicious(v bool) {
	o.BlockSuspicious = &v
}

// GetBlockHarvester returns the BlockHarvester field value if set, zero value otherwise.
func (o *WAFConfigUpdateHttpbl) GetBlockHarvester() bool {
	if o == nil || IsNil(o.BlockHarvester) {
		var ret bool
		return ret
	}
	return *o.BlockHarvester
}

// GetBlockHarvesterOk returns a tuple with the BlockHarvester field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdateHttpbl) GetBlockHarvesterOk() (*bool, bool) {
	if o == nil || IsNil(o.BlockHarvester) {
		return nil, false
	}
	return o.BlockHarvester, true
}

// HasBlockHarvester returns a boolean if a field has been set.
func (o *WAFConfigUpdateHttpbl) HasBlockHarvester() bool {
	if o != nil && !IsNil(o.BlockHarvester) {
		return true
	}

	return false
}

// SetBlockHarvester gets a reference to the given bool and assigns it to the BlockHarvester field.
func (o *WAFConfigUpdateHttpbl) SetBlockHarvester(v bool) {
	o.BlockHarvester = &v
}

// GetBlockSpam returns the BlockSpam field value if set, zero value otherwise.
func (o *WAFConfigUpdateHttpbl) GetBlockSpam() bool {
	if o == nil || IsNil(o.BlockSpam) {
		var ret bool
		return ret
	}
	return *o.BlockSpam
}

// GetBlockSpamOk returns a tuple with the BlockSpam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdateHttpbl) GetBlockSpamOk() (*bool, bool) {
	if o == nil || IsNil(o.BlockSpam) {
		return nil, false
	}
	return o.BlockSpam, true
}

// HasBlockSpam returns a boolean if a field has been set.
func (o *WAFConfigUpdateHttpbl) HasBlockSpam() bool {
	if o != nil && !IsNil(o.BlockSpam) {
		return true
	}

	return false
}

// SetBlockSpam gets a reference to the given bool and assigns it to the BlockSpam field.
func (o *WAFConfigUpdateHttpbl) SetBlockSpam(v bool) {
	o.BlockSpam = &v
}

// GetBlockSearchEngine returns the BlockSearchEngine field value if set, zero value otherwise.
func (o *WAFConfigUpdateHttpbl) GetBlockSearchEngine() bool {
	if o == nil || IsNil(o.BlockSearchEngine) {
		var ret bool
		return ret
	}
	return *o.BlockSearchEngine
}

// GetBlockSearchEngineOk returns a tuple with the BlockSearchEngine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdateHttpbl) GetBlockSearchEngineOk() (*bool, bool) {
	if o == nil || IsNil(o.BlockSearchEngine) {
		return nil, false
	}
	return o.BlockSearchEngine, true
}

// HasBlockSearchEngine returns a boolean if a field has been set.
func (o *WAFConfigUpdateHttpbl) HasBlockSearchEngine() bool {
	if o != nil && !IsNil(o.BlockSearchEngine) {
		return true
	}

	return false
}

// SetBlockSearchEngine gets a reference to the given bool and assigns it to the BlockSearchEngine field.
func (o *WAFConfigUpdateHttpbl) SetBlockSearchEngine(v bool) {
	o.BlockSearchEngine = &v
}

func (o WAFConfigUpdateHttpbl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WAFConfigUpdateHttpbl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HttpblEnabled) {
		toSerialize["httpbl_enabled"] = o.HttpblEnabled
	}
	if !IsNil(o.ApiKey) {
		toSerialize["api_key"] = o.ApiKey
	}
	if !IsNil(o.BlockSuspicious) {
		toSerialize["block_suspicious"] = o.BlockSuspicious
	}
	if !IsNil(o.BlockHarvester) {
		toSerialize["block_harvester"] = o.BlockHarvester
	}
	if !IsNil(o.BlockSpam) {
		toSerialize["block_spam"] = o.BlockSpam
	}
	if !IsNil(o.BlockSearchEngine) {
		toSerialize["block_search_engine"] = o.BlockSearchEngine
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WAFConfigUpdateHttpbl) UnmarshalJSON(data []byte) (err error) {
	varWAFConfigUpdateHttpbl := _WAFConfigUpdateHttpbl{}

	err = json.Unmarshal(data, &varWAFConfigUpdateHttpbl)

	if err != nil {
		return err
	}

	*o = WAFConfigUpdateHttpbl(varWAFConfigUpdateHttpbl)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "httpbl_enabled")
		delete(additionalProperties, "api_key")
		delete(additionalProperties, "block_suspicious")
		delete(additionalProperties, "block_harvester")
		delete(additionalProperties, "block_spam")
		delete(additionalProperties, "block_search_engine")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWAFConfigUpdateHttpbl struct {
	value *WAFConfigUpdateHttpbl
	isSet bool
}

func (v NullableWAFConfigUpdateHttpbl) Get() *WAFConfigUpdateHttpbl {
	return v.value
}

func (v *NullableWAFConfigUpdateHttpbl) Set(val *WAFConfigUpdateHttpbl) {
	v.value = val
	v.isSet = true
}

func (v NullableWAFConfigUpdateHttpbl) IsSet() bool {
	return v.isSet
}

func (v *NullableWAFConfigUpdateHttpbl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWAFConfigUpdateHttpbl(val *WAFConfigUpdateHttpbl) *NullableWAFConfigUpdateHttpbl {
	return &NullableWAFConfigUpdateHttpbl{value: val, isSet: true}
}

func (v NullableWAFConfigUpdateHttpbl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWAFConfigUpdateHttpbl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


