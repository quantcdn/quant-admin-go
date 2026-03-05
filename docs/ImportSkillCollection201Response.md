# ImportSkillCollection201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to **string** |  | [optional] 
**Imported** | Pointer to **int32** |  | [optional] 
**Failed** | Pointer to **int32** |  | [optional] 
**Skills** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Errors** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewImportSkillCollection201Response

`func NewImportSkillCollection201Response() *ImportSkillCollection201Response`

NewImportSkillCollection201Response instantiates a new ImportSkillCollection201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportSkillCollection201ResponseWithDefaults

`func NewImportSkillCollection201ResponseWithDefaults() *ImportSkillCollection201Response`

NewImportSkillCollection201ResponseWithDefaults instantiates a new ImportSkillCollection201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *ImportSkillCollection201Response) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ImportSkillCollection201Response) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ImportSkillCollection201Response) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ImportSkillCollection201Response) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetImported

`func (o *ImportSkillCollection201Response) GetImported() int32`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *ImportSkillCollection201Response) GetImportedOk() (*int32, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *ImportSkillCollection201Response) SetImported(v int32)`

SetImported sets Imported field to given value.

### HasImported

`func (o *ImportSkillCollection201Response) HasImported() bool`

HasImported returns a boolean if a field has been set.

### GetFailed

`func (o *ImportSkillCollection201Response) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ImportSkillCollection201Response) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ImportSkillCollection201Response) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ImportSkillCollection201Response) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetSkills

`func (o *ImportSkillCollection201Response) GetSkills() []map[string]interface{}`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *ImportSkillCollection201Response) GetSkillsOk() (*[]map[string]interface{}, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *ImportSkillCollection201Response) SetSkills(v []map[string]interface{})`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *ImportSkillCollection201Response) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### GetErrors

`func (o *ImportSkillCollection201Response) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ImportSkillCollection201Response) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ImportSkillCollection201Response) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ImportSkillCollection201Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


