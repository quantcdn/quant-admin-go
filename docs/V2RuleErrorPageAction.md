# V2RuleErrorPageAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorPagePath** | **string** | Published path of the page to serve as the error page | 
**StatusCodes** | **[]string** | Status codes this page is served for | 

## Methods

### NewV2RuleErrorPageAction

`func NewV2RuleErrorPageAction(errorPagePath string, statusCodes []string, ) *V2RuleErrorPageAction`

NewV2RuleErrorPageAction instantiates a new V2RuleErrorPageAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleErrorPageActionWithDefaults

`func NewV2RuleErrorPageActionWithDefaults() *V2RuleErrorPageAction`

NewV2RuleErrorPageActionWithDefaults instantiates a new V2RuleErrorPageAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorPagePath

`func (o *V2RuleErrorPageAction) GetErrorPagePath() string`

GetErrorPagePath returns the ErrorPagePath field if non-nil, zero value otherwise.

### GetErrorPagePathOk

`func (o *V2RuleErrorPageAction) GetErrorPagePathOk() (*string, bool)`

GetErrorPagePathOk returns a tuple with the ErrorPagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPagePath

`func (o *V2RuleErrorPageAction) SetErrorPagePath(v string)`

SetErrorPagePath sets ErrorPagePath field to given value.


### GetStatusCodes

`func (o *V2RuleErrorPageAction) GetStatusCodes() []string`

GetStatusCodes returns the StatusCodes field if non-nil, zero value otherwise.

### GetStatusCodesOk

`func (o *V2RuleErrorPageAction) GetStatusCodesOk() (*[]string, bool)`

GetStatusCodesOk returns a tuple with the StatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodes

`func (o *V2RuleErrorPageAction) SetStatusCodes(v []string)`

SetStatusCodes sets StatusCodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


