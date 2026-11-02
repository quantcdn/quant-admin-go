# RulesIndexList200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Position** | Pointer to **int32** | 0-based index in the stored order. | [optional] 
**Managed** | Pointer to **bool** | True when a template deploy owns the rule in this project. | [optional] 

## Methods

### NewRulesIndexList200ResponseInner

`func NewRulesIndexList200ResponseInner() *RulesIndexList200ResponseInner`

NewRulesIndexList200ResponseInner instantiates a new RulesIndexList200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesIndexList200ResponseInnerWithDefaults

`func NewRulesIndexList200ResponseInnerWithDefaults() *RulesIndexList200ResponseInner`

NewRulesIndexList200ResponseInnerWithDefaults instantiates a new RulesIndexList200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosition

`func (o *RulesIndexList200ResponseInner) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *RulesIndexList200ResponseInner) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *RulesIndexList200ResponseInner) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *RulesIndexList200ResponseInner) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetManaged

`func (o *RulesIndexList200ResponseInner) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *RulesIndexList200ResponseInner) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *RulesIndexList200ResponseInner) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *RulesIndexList200ResponseInner) HasManaged() bool`

HasManaged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


