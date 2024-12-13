# RuleCustomResponseAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomResponseBody** | **string** |  | 
**CustomResponseStatusCode** | **int32** |  | [default to 200]

## Methods

### NewRuleCustomResponseAction

`func NewRuleCustomResponseAction(customResponseBody string, customResponseStatusCode int32, ) *RuleCustomResponseAction`

NewRuleCustomResponseAction instantiates a new RuleCustomResponseAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleCustomResponseActionWithDefaults

`func NewRuleCustomResponseActionWithDefaults() *RuleCustomResponseAction`

NewRuleCustomResponseActionWithDefaults instantiates a new RuleCustomResponseAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomResponseBody

`func (o *RuleCustomResponseAction) GetCustomResponseBody() string`

GetCustomResponseBody returns the CustomResponseBody field if non-nil, zero value otherwise.

### GetCustomResponseBodyOk

`func (o *RuleCustomResponseAction) GetCustomResponseBodyOk() (*string, bool)`

GetCustomResponseBodyOk returns a tuple with the CustomResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomResponseBody

`func (o *RuleCustomResponseAction) SetCustomResponseBody(v string)`

SetCustomResponseBody sets CustomResponseBody field to given value.


### GetCustomResponseStatusCode

`func (o *RuleCustomResponseAction) GetCustomResponseStatusCode() int32`

GetCustomResponseStatusCode returns the CustomResponseStatusCode field if non-nil, zero value otherwise.

### GetCustomResponseStatusCodeOk

`func (o *RuleCustomResponseAction) GetCustomResponseStatusCodeOk() (*int32, bool)`

GetCustomResponseStatusCodeOk returns a tuple with the CustomResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomResponseStatusCode

`func (o *RuleCustomResponseAction) SetCustomResponseStatusCode(v int32)`

SetCustomResponseStatusCode sets CustomResponseStatusCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


