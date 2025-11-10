# V2RuleRedirectAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **string** | Redirect destination URL | 
**StatusCode** | Pointer to **string** | HTTP status code for redirect | [optional] [default to "301"]

## Methods

### NewV2RuleRedirectAction

`func NewV2RuleRedirectAction(to string, ) *V2RuleRedirectAction`

NewV2RuleRedirectAction instantiates a new V2RuleRedirectAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleRedirectActionWithDefaults

`func NewV2RuleRedirectActionWithDefaults() *V2RuleRedirectAction`

NewV2RuleRedirectActionWithDefaults instantiates a new V2RuleRedirectAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *V2RuleRedirectAction) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *V2RuleRedirectAction) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *V2RuleRedirectAction) SetTo(v string)`

SetTo sets To field to given value.


### GetStatusCode

`func (o *V2RuleRedirectAction) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *V2RuleRedirectAction) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *V2RuleRedirectAction) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *V2RuleRedirectAction) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


