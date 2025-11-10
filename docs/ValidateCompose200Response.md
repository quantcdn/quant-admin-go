# ValidateCompose200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**TranslatedComposeDefinition** | **map[string]interface{}** | The translated internal compose definition format | 
**TranslationWarnings** | Pointer to **[]string** | Optional warnings encountered during translation | [optional] 

## Methods

### NewValidateCompose200Response

`func NewValidateCompose200Response(message string, translatedComposeDefinition map[string]interface{}, ) *ValidateCompose200Response`

NewValidateCompose200Response instantiates a new ValidateCompose200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateCompose200ResponseWithDefaults

`func NewValidateCompose200ResponseWithDefaults() *ValidateCompose200Response`

NewValidateCompose200ResponseWithDefaults instantiates a new ValidateCompose200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ValidateCompose200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidateCompose200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidateCompose200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTranslatedComposeDefinition

`func (o *ValidateCompose200Response) GetTranslatedComposeDefinition() map[string]interface{}`

GetTranslatedComposeDefinition returns the TranslatedComposeDefinition field if non-nil, zero value otherwise.

### GetTranslatedComposeDefinitionOk

`func (o *ValidateCompose200Response) GetTranslatedComposeDefinitionOk() (*map[string]interface{}, bool)`

GetTranslatedComposeDefinitionOk returns a tuple with the TranslatedComposeDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedComposeDefinition

`func (o *ValidateCompose200Response) SetTranslatedComposeDefinition(v map[string]interface{})`

SetTranslatedComposeDefinition sets TranslatedComposeDefinition field to given value.


### GetTranslationWarnings

`func (o *ValidateCompose200Response) GetTranslationWarnings() []string`

GetTranslationWarnings returns the TranslationWarnings field if non-nil, zero value otherwise.

### GetTranslationWarningsOk

`func (o *ValidateCompose200Response) GetTranslationWarningsOk() (*[]string, bool)`

GetTranslationWarningsOk returns a tuple with the TranslationWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationWarnings

`func (o *ValidateCompose200Response) SetTranslationWarnings(v []string)`

SetTranslationWarnings sets TranslationWarnings field to given value.

### HasTranslationWarnings

`func (o *ValidateCompose200Response) HasTranslationWarnings() bool`

HasTranslationWarnings returns a boolean if a field has been set.

### SetTranslationWarningsNil

`func (o *ValidateCompose200Response) SetTranslationWarningsNil(b bool)`

 SetTranslationWarningsNil sets the value for TranslationWarnings to be an explicit nil

### UnsetTranslationWarnings
`func (o *ValidateCompose200Response) UnsetTranslationWarnings()`

UnsetTranslationWarnings ensures that no value is present for TranslationWarnings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


