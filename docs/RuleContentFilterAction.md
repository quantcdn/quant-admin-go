# RuleContentFilterAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentFilterBody** | **string** |  | 
**ContentFilterStatusCode** | **int32** |  | [default to 200]

## Methods

### NewRuleContentFilterAction

`func NewRuleContentFilterAction(contentFilterBody string, contentFilterStatusCode int32, ) *RuleContentFilterAction`

NewRuleContentFilterAction instantiates a new RuleContentFilterAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleContentFilterActionWithDefaults

`func NewRuleContentFilterActionWithDefaults() *RuleContentFilterAction`

NewRuleContentFilterActionWithDefaults instantiates a new RuleContentFilterAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentFilterBody

`func (o *RuleContentFilterAction) GetContentFilterBody() string`

GetContentFilterBody returns the ContentFilterBody field if non-nil, zero value otherwise.

### GetContentFilterBodyOk

`func (o *RuleContentFilterAction) GetContentFilterBodyOk() (*string, bool)`

GetContentFilterBodyOk returns a tuple with the ContentFilterBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFilterBody

`func (o *RuleContentFilterAction) SetContentFilterBody(v string)`

SetContentFilterBody sets ContentFilterBody field to given value.


### GetContentFilterStatusCode

`func (o *RuleContentFilterAction) GetContentFilterStatusCode() int32`

GetContentFilterStatusCode returns the ContentFilterStatusCode field if non-nil, zero value otherwise.

### GetContentFilterStatusCodeOk

`func (o *RuleContentFilterAction) GetContentFilterStatusCodeOk() (*int32, bool)`

GetContentFilterStatusCodeOk returns a tuple with the ContentFilterStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFilterStatusCode

`func (o *RuleContentFilterAction) SetContentFilterStatusCode(v int32)`

SetContentFilterStatusCode sets ContentFilterStatusCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


