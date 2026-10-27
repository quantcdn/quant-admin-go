# V2KvUpgradeState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **int32** |  | 
**State** | **string** | migrating while writes are frozen (an upgrade copy or a rollback). | 
**Table** | **string** | The table that serves the store. | 
**Upgrade** | [**NullableV2KvUpgradeStatus**](V2KvUpgradeStatus.md) | Null when the store was never upgraded. | 

## Methods

### NewV2KvUpgradeState

`func NewV2KvUpgradeState(format int32, state string, table string, upgrade NullableV2KvUpgradeStatus, ) *V2KvUpgradeState`

NewV2KvUpgradeState instantiates a new V2KvUpgradeState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvUpgradeStateWithDefaults

`func NewV2KvUpgradeStateWithDefaults() *V2KvUpgradeState`

NewV2KvUpgradeStateWithDefaults instantiates a new V2KvUpgradeState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *V2KvUpgradeState) GetFormat() int32`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *V2KvUpgradeState) GetFormatOk() (*int32, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *V2KvUpgradeState) SetFormat(v int32)`

SetFormat sets Format field to given value.


### GetState

`func (o *V2KvUpgradeState) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V2KvUpgradeState) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V2KvUpgradeState) SetState(v string)`

SetState sets State field to given value.


### GetTable

`func (o *V2KvUpgradeState) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *V2KvUpgradeState) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *V2KvUpgradeState) SetTable(v string)`

SetTable sets Table field to given value.


### GetUpgrade

`func (o *V2KvUpgradeState) GetUpgrade() V2KvUpgradeStatus`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *V2KvUpgradeState) GetUpgradeOk() (*V2KvUpgradeStatus, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *V2KvUpgradeState) SetUpgrade(v V2KvUpgradeStatus)`

SetUpgrade sets Upgrade field to given value.


### SetUpgradeNil

`func (o *V2KvUpgradeState) SetUpgradeNil(b bool)`

 SetUpgradeNil sets the value for Upgrade to be an explicit nil

### UnsetUpgrade
`func (o *V2KvUpgradeState) UnsetUpgrade()`

UnsetUpgrade ensures that no value is present for Upgrade, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


