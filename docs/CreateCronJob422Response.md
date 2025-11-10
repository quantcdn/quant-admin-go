# CreateCronJob422Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateCronJob422Response

`func NewCreateCronJob422Response() *CreateCronJob422Response`

NewCreateCronJob422Response instantiates a new CreateCronJob422Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCronJob422ResponseWithDefaults

`func NewCreateCronJob422ResponseWithDefaults() *CreateCronJob422Response`

NewCreateCronJob422ResponseWithDefaults instantiates a new CreateCronJob422Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CreateCronJob422Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateCronJob422Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateCronJob422Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateCronJob422Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrors

`func (o *CreateCronJob422Response) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateCronJob422Response) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateCronJob422Response) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreateCronJob422Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


