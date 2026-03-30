# GetAgentOverlay200ResponseOverlay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelId** | Pointer to **string** |  | [optional] 
**Temperature** | Pointer to **float32** |  | [optional] 
**MaxTokens** | Pointer to **int32** |  | [optional] 
**DisabledSkills** | Pointer to **[]string** |  | [optional] 
**AdditionalSkills** | Pointer to **[]string** |  | [optional] 
**AdditionalTools** | Pointer to **[]string** |  | [optional] 
**DisabledTools** | Pointer to **[]string** |  | [optional] 
**SystemPromptAppend** | Pointer to **string** |  | [optional] 
**AllowedCollections** | Pointer to **[]string** |  | [optional] 
**GuardrailPreset** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetAgentOverlay200ResponseOverlay

`func NewGetAgentOverlay200ResponseOverlay() *GetAgentOverlay200ResponseOverlay`

NewGetAgentOverlay200ResponseOverlay instantiates a new GetAgentOverlay200ResponseOverlay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAgentOverlay200ResponseOverlayWithDefaults

`func NewGetAgentOverlay200ResponseOverlayWithDefaults() *GetAgentOverlay200ResponseOverlay`

NewGetAgentOverlay200ResponseOverlayWithDefaults instantiates a new GetAgentOverlay200ResponseOverlay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelId

`func (o *GetAgentOverlay200ResponseOverlay) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *GetAgentOverlay200ResponseOverlay) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *GetAgentOverlay200ResponseOverlay) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *GetAgentOverlay200ResponseOverlay) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetTemperature

`func (o *GetAgentOverlay200ResponseOverlay) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *GetAgentOverlay200ResponseOverlay) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *GetAgentOverlay200ResponseOverlay) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *GetAgentOverlay200ResponseOverlay) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetMaxTokens

`func (o *GetAgentOverlay200ResponseOverlay) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *GetAgentOverlay200ResponseOverlay) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *GetAgentOverlay200ResponseOverlay) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *GetAgentOverlay200ResponseOverlay) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetDisabledSkills

`func (o *GetAgentOverlay200ResponseOverlay) GetDisabledSkills() []string`

GetDisabledSkills returns the DisabledSkills field if non-nil, zero value otherwise.

### GetDisabledSkillsOk

`func (o *GetAgentOverlay200ResponseOverlay) GetDisabledSkillsOk() (*[]string, bool)`

GetDisabledSkillsOk returns a tuple with the DisabledSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledSkills

`func (o *GetAgentOverlay200ResponseOverlay) SetDisabledSkills(v []string)`

SetDisabledSkills sets DisabledSkills field to given value.

### HasDisabledSkills

`func (o *GetAgentOverlay200ResponseOverlay) HasDisabledSkills() bool`

HasDisabledSkills returns a boolean if a field has been set.

### GetAdditionalSkills

`func (o *GetAgentOverlay200ResponseOverlay) GetAdditionalSkills() []string`

GetAdditionalSkills returns the AdditionalSkills field if non-nil, zero value otherwise.

### GetAdditionalSkillsOk

`func (o *GetAgentOverlay200ResponseOverlay) GetAdditionalSkillsOk() (*[]string, bool)`

GetAdditionalSkillsOk returns a tuple with the AdditionalSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSkills

`func (o *GetAgentOverlay200ResponseOverlay) SetAdditionalSkills(v []string)`

SetAdditionalSkills sets AdditionalSkills field to given value.

### HasAdditionalSkills

`func (o *GetAgentOverlay200ResponseOverlay) HasAdditionalSkills() bool`

HasAdditionalSkills returns a boolean if a field has been set.

### GetAdditionalTools

`func (o *GetAgentOverlay200ResponseOverlay) GetAdditionalTools() []string`

GetAdditionalTools returns the AdditionalTools field if non-nil, zero value otherwise.

### GetAdditionalToolsOk

`func (o *GetAgentOverlay200ResponseOverlay) GetAdditionalToolsOk() (*[]string, bool)`

GetAdditionalToolsOk returns a tuple with the AdditionalTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalTools

`func (o *GetAgentOverlay200ResponseOverlay) SetAdditionalTools(v []string)`

SetAdditionalTools sets AdditionalTools field to given value.

### HasAdditionalTools

`func (o *GetAgentOverlay200ResponseOverlay) HasAdditionalTools() bool`

HasAdditionalTools returns a boolean if a field has been set.

### GetDisabledTools

`func (o *GetAgentOverlay200ResponseOverlay) GetDisabledTools() []string`

GetDisabledTools returns the DisabledTools field if non-nil, zero value otherwise.

### GetDisabledToolsOk

`func (o *GetAgentOverlay200ResponseOverlay) GetDisabledToolsOk() (*[]string, bool)`

GetDisabledToolsOk returns a tuple with the DisabledTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledTools

`func (o *GetAgentOverlay200ResponseOverlay) SetDisabledTools(v []string)`

SetDisabledTools sets DisabledTools field to given value.

### HasDisabledTools

`func (o *GetAgentOverlay200ResponseOverlay) HasDisabledTools() bool`

HasDisabledTools returns a boolean if a field has been set.

### GetSystemPromptAppend

`func (o *GetAgentOverlay200ResponseOverlay) GetSystemPromptAppend() string`

GetSystemPromptAppend returns the SystemPromptAppend field if non-nil, zero value otherwise.

### GetSystemPromptAppendOk

`func (o *GetAgentOverlay200ResponseOverlay) GetSystemPromptAppendOk() (*string, bool)`

GetSystemPromptAppendOk returns a tuple with the SystemPromptAppend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPromptAppend

`func (o *GetAgentOverlay200ResponseOverlay) SetSystemPromptAppend(v string)`

SetSystemPromptAppend sets SystemPromptAppend field to given value.

### HasSystemPromptAppend

`func (o *GetAgentOverlay200ResponseOverlay) HasSystemPromptAppend() bool`

HasSystemPromptAppend returns a boolean if a field has been set.

### GetAllowedCollections

`func (o *GetAgentOverlay200ResponseOverlay) GetAllowedCollections() []string`

GetAllowedCollections returns the AllowedCollections field if non-nil, zero value otherwise.

### GetAllowedCollectionsOk

`func (o *GetAgentOverlay200ResponseOverlay) GetAllowedCollectionsOk() (*[]string, bool)`

GetAllowedCollectionsOk returns a tuple with the AllowedCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCollections

`func (o *GetAgentOverlay200ResponseOverlay) SetAllowedCollections(v []string)`

SetAllowedCollections sets AllowedCollections field to given value.

### HasAllowedCollections

`func (o *GetAgentOverlay200ResponseOverlay) HasAllowedCollections() bool`

HasAllowedCollections returns a boolean if a field has been set.

### GetGuardrailPreset

`func (o *GetAgentOverlay200ResponseOverlay) GetGuardrailPreset() string`

GetGuardrailPreset returns the GuardrailPreset field if non-nil, zero value otherwise.

### GetGuardrailPresetOk

`func (o *GetAgentOverlay200ResponseOverlay) GetGuardrailPresetOk() (*string, bool)`

GetGuardrailPresetOk returns a tuple with the GuardrailPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrailPreset

`func (o *GetAgentOverlay200ResponseOverlay) SetGuardrailPreset(v string)`

SetGuardrailPreset sets GuardrailPreset field to given value.

### HasGuardrailPreset

`func (o *GetAgentOverlay200ResponseOverlay) HasGuardrailPreset() bool`

HasGuardrailPreset returns a boolean if a field has been set.

### GetVersion

`func (o *GetAgentOverlay200ResponseOverlay) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetAgentOverlay200ResponseOverlay) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetAgentOverlay200ResponseOverlay) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetAgentOverlay200ResponseOverlay) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


