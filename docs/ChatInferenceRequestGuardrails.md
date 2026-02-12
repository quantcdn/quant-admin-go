# ChatInferenceRequestGuardrails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuardrailIdentifier** | Pointer to **string** | Guardrail identifier from AWS Bedrock | [optional] 
**GuardrailVersion** | Pointer to **string** | Guardrail version | [optional] 
**Trace** | Pointer to **string** | Enable guardrail trace output | [optional] 

## Methods

### NewChatInferenceRequestGuardrails

`func NewChatInferenceRequestGuardrails() *ChatInferenceRequestGuardrails`

NewChatInferenceRequestGuardrails instantiates a new ChatInferenceRequestGuardrails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatInferenceRequestGuardrailsWithDefaults

`func NewChatInferenceRequestGuardrailsWithDefaults() *ChatInferenceRequestGuardrails`

NewChatInferenceRequestGuardrailsWithDefaults instantiates a new ChatInferenceRequestGuardrails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuardrailIdentifier

`func (o *ChatInferenceRequestGuardrails) GetGuardrailIdentifier() string`

GetGuardrailIdentifier returns the GuardrailIdentifier field if non-nil, zero value otherwise.

### GetGuardrailIdentifierOk

`func (o *ChatInferenceRequestGuardrails) GetGuardrailIdentifierOk() (*string, bool)`

GetGuardrailIdentifierOk returns a tuple with the GuardrailIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrailIdentifier

`func (o *ChatInferenceRequestGuardrails) SetGuardrailIdentifier(v string)`

SetGuardrailIdentifier sets GuardrailIdentifier field to given value.

### HasGuardrailIdentifier

`func (o *ChatInferenceRequestGuardrails) HasGuardrailIdentifier() bool`

HasGuardrailIdentifier returns a boolean if a field has been set.

### GetGuardrailVersion

`func (o *ChatInferenceRequestGuardrails) GetGuardrailVersion() string`

GetGuardrailVersion returns the GuardrailVersion field if non-nil, zero value otherwise.

### GetGuardrailVersionOk

`func (o *ChatInferenceRequestGuardrails) GetGuardrailVersionOk() (*string, bool)`

GetGuardrailVersionOk returns a tuple with the GuardrailVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrailVersion

`func (o *ChatInferenceRequestGuardrails) SetGuardrailVersion(v string)`

SetGuardrailVersion sets GuardrailVersion field to given value.

### HasGuardrailVersion

`func (o *ChatInferenceRequestGuardrails) HasGuardrailVersion() bool`

HasGuardrailVersion returns a boolean if a field has been set.

### GetTrace

`func (o *ChatInferenceRequestGuardrails) GetTrace() string`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *ChatInferenceRequestGuardrails) GetTraceOk() (*string, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *ChatInferenceRequestGuardrails) SetTrace(v string)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *ChatInferenceRequestGuardrails) HasTrace() bool`

HasTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


