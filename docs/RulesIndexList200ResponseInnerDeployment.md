# RulesIndexList200ResponseInnerDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **string** | The &#x60;namespace&#x60; of the edge-rules file that wrote the rule. | 
**Id** | **string** | The &#x60;id&#x60; of the rule in that file. | 
**Hash** | Pointer to **NullableString** | Hash of the rule as the deploy last wrote it, used to detect changes by other paths. | [optional] 

## Methods

### NewRulesIndexList200ResponseInnerDeployment

`func NewRulesIndexList200ResponseInnerDeployment(namespace string, id string, ) *RulesIndexList200ResponseInnerDeployment`

NewRulesIndexList200ResponseInnerDeployment instantiates a new RulesIndexList200ResponseInnerDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesIndexList200ResponseInnerDeploymentWithDefaults

`func NewRulesIndexList200ResponseInnerDeploymentWithDefaults() *RulesIndexList200ResponseInnerDeployment`

NewRulesIndexList200ResponseInnerDeploymentWithDefaults instantiates a new RulesIndexList200ResponseInnerDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *RulesIndexList200ResponseInnerDeployment) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *RulesIndexList200ResponseInnerDeployment) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *RulesIndexList200ResponseInnerDeployment) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetId

`func (o *RulesIndexList200ResponseInnerDeployment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RulesIndexList200ResponseInnerDeployment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RulesIndexList200ResponseInnerDeployment) SetId(v string)`

SetId sets Id field to given value.


### GetHash

`func (o *RulesIndexList200ResponseInnerDeployment) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *RulesIndexList200ResponseInnerDeployment) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *RulesIndexList200ResponseInnerDeployment) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *RulesIndexList200ResponseInnerDeployment) HasHash() bool`

HasHash returns a boolean if a field has been set.

### SetHashNil

`func (o *RulesIndexList200ResponseInnerDeployment) SetHashNil(b bool)`

 SetHashNil sets the value for Hash to be an explicit nil

### UnsetHash
`func (o *RulesIndexList200ResponseInnerDeployment) UnsetHash()`

UnsetHash ensures that no value is present for Hash, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


