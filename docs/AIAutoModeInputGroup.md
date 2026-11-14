# AIAutoModeInputGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | Pointer to **int32** | Revision from the last read. A stale value returns 409. | [optional] 
**Roles** | [**AIAutoModeConfigGroupRoles**](AIAutoModeConfigGroupRoles.md) |  | 

## Methods

### NewAIAutoModeInputGroup

`func NewAIAutoModeInputGroup(roles AIAutoModeConfigGroupRoles, ) *AIAutoModeInputGroup`

NewAIAutoModeInputGroup instantiates a new AIAutoModeInputGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIAutoModeInputGroupWithDefaults

`func NewAIAutoModeInputGroupWithDefaults() *AIAutoModeInputGroup`

NewAIAutoModeInputGroupWithDefaults instantiates a new AIAutoModeInputGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *AIAutoModeInputGroup) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *AIAutoModeInputGroup) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *AIAutoModeInputGroup) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *AIAutoModeInputGroup) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRoles

`func (o *AIAutoModeInputGroup) GetRoles() AIAutoModeConfigGroupRoles`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AIAutoModeInputGroup) GetRolesOk() (*AIAutoModeConfigGroupRoles, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AIAutoModeInputGroup) SetRoles(v AIAutoModeConfigGroupRoles)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


