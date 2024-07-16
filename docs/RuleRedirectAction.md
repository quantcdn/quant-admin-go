# RuleRedirectAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **string** |  | 
**StatusCode** | **string** |  | [default to "301"]

## Methods

### NewRuleRedirectAction

`func NewRuleRedirectAction(to string, statusCode string, ) *RuleRedirectAction`

NewRuleRedirectAction instantiates a new RuleRedirectAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleRedirectActionWithDefaults

`func NewRuleRedirectActionWithDefaults() *RuleRedirectAction`

NewRuleRedirectActionWithDefaults instantiates a new RuleRedirectAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *RuleRedirectAction) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *RuleRedirectAction) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *RuleRedirectAction) SetTo(v string)`

SetTo sets To field to given value.


### GetStatusCode

`func (o *RuleRedirectAction) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *RuleRedirectAction) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *RuleRedirectAction) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


