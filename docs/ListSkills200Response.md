# ListSkills200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skills** | Pointer to [**[]ListSkills200ResponseSkillsInner**](ListSkills200ResponseSkillsInner.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 

## Methods

### NewListSkills200Response

`func NewListSkills200Response() *ListSkills200Response`

NewListSkills200Response instantiates a new ListSkills200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSkills200ResponseWithDefaults

`func NewListSkills200ResponseWithDefaults() *ListSkills200Response`

NewListSkills200ResponseWithDefaults instantiates a new ListSkills200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkills

`func (o *ListSkills200Response) GetSkills() []ListSkills200ResponseSkillsInner`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *ListSkills200Response) GetSkillsOk() (*[]ListSkills200ResponseSkillsInner, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *ListSkills200Response) SetSkills(v []ListSkills200ResponseSkillsInner)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *ListSkills200Response) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### GetCount

`func (o *ListSkills200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListSkills200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListSkills200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListSkills200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


