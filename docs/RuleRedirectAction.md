# RuleRedirectAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectTo** | **string** |  | 
**RedirectCode** | **string** |  | [default to "301"]

## Methods

### NewRuleRedirectAction

`func NewRuleRedirectAction(redirectTo string, redirectCode string, ) *RuleRedirectAction`

NewRuleRedirectAction instantiates a new RuleRedirectAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleRedirectActionWithDefaults

`func NewRuleRedirectActionWithDefaults() *RuleRedirectAction`

NewRuleRedirectActionWithDefaults instantiates a new RuleRedirectAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectTo

`func (o *RuleRedirectAction) GetRedirectTo() string`

GetRedirectTo returns the RedirectTo field if non-nil, zero value otherwise.

### GetRedirectToOk

`func (o *RuleRedirectAction) GetRedirectToOk() (*string, bool)`

GetRedirectToOk returns a tuple with the RedirectTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTo

`func (o *RuleRedirectAction) SetRedirectTo(v string)`

SetRedirectTo sets RedirectTo field to given value.


### GetRedirectCode

`func (o *RuleRedirectAction) GetRedirectCode() string`

GetRedirectCode returns the RedirectCode field if non-nil, zero value otherwise.

### GetRedirectCodeOk

`func (o *RuleRedirectAction) GetRedirectCodeOk() (*string, bool)`

GetRedirectCodeOk returns a tuple with the RedirectCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectCode

`func (o *RuleRedirectAction) SetRedirectCode(v string)`

SetRedirectCode sets RedirectCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


