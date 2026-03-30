# UpsertAgentOverlayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelId** | Pointer to **string** | Override the base agent&#39;s model | [optional] 
**Temperature** | Pointer to **float32** | Override temperature | [optional] 
**MaxTokens** | Pointer to **int32** | Override max tokens | [optional] 
**DisabledSkills** | Pointer to **[]string** | Global skill IDs to exclude | [optional] 
**AdditionalSkills** | Pointer to **[]string** | Org-owned skill IDs to add | [optional] 
**AdditionalTools** | Pointer to **[]string** | Tool names to add | [optional] 
**DisabledTools** | Pointer to **[]string** | Tool names to remove | [optional] 
**SystemPromptAppend** | Pointer to **string** | Text appended to base system prompt | [optional] 
**AllowedCollections** | Pointer to **[]string** | Vector DB collections | [optional] 
**GuardrailPreset** | Pointer to **string** | Guardrail preset | [optional] 
**Version** | Pointer to **int32** | Current version for optimistic concurrency | [optional] 

## Methods

### NewUpsertAgentOverlayRequest

`func NewUpsertAgentOverlayRequest() *UpsertAgentOverlayRequest`

NewUpsertAgentOverlayRequest instantiates a new UpsertAgentOverlayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertAgentOverlayRequestWithDefaults

`func NewUpsertAgentOverlayRequestWithDefaults() *UpsertAgentOverlayRequest`

NewUpsertAgentOverlayRequestWithDefaults instantiates a new UpsertAgentOverlayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelId

`func (o *UpsertAgentOverlayRequest) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *UpsertAgentOverlayRequest) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *UpsertAgentOverlayRequest) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *UpsertAgentOverlayRequest) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetTemperature

`func (o *UpsertAgentOverlayRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *UpsertAgentOverlayRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *UpsertAgentOverlayRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *UpsertAgentOverlayRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetMaxTokens

`func (o *UpsertAgentOverlayRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *UpsertAgentOverlayRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *UpsertAgentOverlayRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *UpsertAgentOverlayRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetDisabledSkills

`func (o *UpsertAgentOverlayRequest) GetDisabledSkills() []string`

GetDisabledSkills returns the DisabledSkills field if non-nil, zero value otherwise.

### GetDisabledSkillsOk

`func (o *UpsertAgentOverlayRequest) GetDisabledSkillsOk() (*[]string, bool)`

GetDisabledSkillsOk returns a tuple with the DisabledSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledSkills

`func (o *UpsertAgentOverlayRequest) SetDisabledSkills(v []string)`

SetDisabledSkills sets DisabledSkills field to given value.

### HasDisabledSkills

`func (o *UpsertAgentOverlayRequest) HasDisabledSkills() bool`

HasDisabledSkills returns a boolean if a field has been set.

### GetAdditionalSkills

`func (o *UpsertAgentOverlayRequest) GetAdditionalSkills() []string`

GetAdditionalSkills returns the AdditionalSkills field if non-nil, zero value otherwise.

### GetAdditionalSkillsOk

`func (o *UpsertAgentOverlayRequest) GetAdditionalSkillsOk() (*[]string, bool)`

GetAdditionalSkillsOk returns a tuple with the AdditionalSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSkills

`func (o *UpsertAgentOverlayRequest) SetAdditionalSkills(v []string)`

SetAdditionalSkills sets AdditionalSkills field to given value.

### HasAdditionalSkills

`func (o *UpsertAgentOverlayRequest) HasAdditionalSkills() bool`

HasAdditionalSkills returns a boolean if a field has been set.

### GetAdditionalTools

`func (o *UpsertAgentOverlayRequest) GetAdditionalTools() []string`

GetAdditionalTools returns the AdditionalTools field if non-nil, zero value otherwise.

### GetAdditionalToolsOk

`func (o *UpsertAgentOverlayRequest) GetAdditionalToolsOk() (*[]string, bool)`

GetAdditionalToolsOk returns a tuple with the AdditionalTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalTools

`func (o *UpsertAgentOverlayRequest) SetAdditionalTools(v []string)`

SetAdditionalTools sets AdditionalTools field to given value.

### HasAdditionalTools

`func (o *UpsertAgentOverlayRequest) HasAdditionalTools() bool`

HasAdditionalTools returns a boolean if a field has been set.

### GetDisabledTools

`func (o *UpsertAgentOverlayRequest) GetDisabledTools() []string`

GetDisabledTools returns the DisabledTools field if non-nil, zero value otherwise.

### GetDisabledToolsOk

`func (o *UpsertAgentOverlayRequest) GetDisabledToolsOk() (*[]string, bool)`

GetDisabledToolsOk returns a tuple with the DisabledTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledTools

`func (o *UpsertAgentOverlayRequest) SetDisabledTools(v []string)`

SetDisabledTools sets DisabledTools field to given value.

### HasDisabledTools

`func (o *UpsertAgentOverlayRequest) HasDisabledTools() bool`

HasDisabledTools returns a boolean if a field has been set.

### GetSystemPromptAppend

`func (o *UpsertAgentOverlayRequest) GetSystemPromptAppend() string`

GetSystemPromptAppend returns the SystemPromptAppend field if non-nil, zero value otherwise.

### GetSystemPromptAppendOk

`func (o *UpsertAgentOverlayRequest) GetSystemPromptAppendOk() (*string, bool)`

GetSystemPromptAppendOk returns a tuple with the SystemPromptAppend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPromptAppend

`func (o *UpsertAgentOverlayRequest) SetSystemPromptAppend(v string)`

SetSystemPromptAppend sets SystemPromptAppend field to given value.

### HasSystemPromptAppend

`func (o *UpsertAgentOverlayRequest) HasSystemPromptAppend() bool`

HasSystemPromptAppend returns a boolean if a field has been set.

### GetAllowedCollections

`func (o *UpsertAgentOverlayRequest) GetAllowedCollections() []string`

GetAllowedCollections returns the AllowedCollections field if non-nil, zero value otherwise.

### GetAllowedCollectionsOk

`func (o *UpsertAgentOverlayRequest) GetAllowedCollectionsOk() (*[]string, bool)`

GetAllowedCollectionsOk returns a tuple with the AllowedCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCollections

`func (o *UpsertAgentOverlayRequest) SetAllowedCollections(v []string)`

SetAllowedCollections sets AllowedCollections field to given value.

### HasAllowedCollections

`func (o *UpsertAgentOverlayRequest) HasAllowedCollections() bool`

HasAllowedCollections returns a boolean if a field has been set.

### GetGuardrailPreset

`func (o *UpsertAgentOverlayRequest) GetGuardrailPreset() string`

GetGuardrailPreset returns the GuardrailPreset field if non-nil, zero value otherwise.

### GetGuardrailPresetOk

`func (o *UpsertAgentOverlayRequest) GetGuardrailPresetOk() (*string, bool)`

GetGuardrailPresetOk returns a tuple with the GuardrailPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrailPreset

`func (o *UpsertAgentOverlayRequest) SetGuardrailPreset(v string)`

SetGuardrailPreset sets GuardrailPreset field to given value.

### HasGuardrailPreset

`func (o *UpsertAgentOverlayRequest) HasGuardrailPreset() bool`

HasGuardrailPreset returns a boolean if a field has been set.

### GetVersion

`func (o *UpsertAgentOverlayRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpsertAgentOverlayRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpsertAgentOverlayRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpsertAgentOverlayRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


