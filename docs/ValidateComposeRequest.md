# ValidateComposeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compose** | **string** | The docker-compose.yml file content as a string | 
**ImageSuffix** | Pointer to **string** | Optional image tag suffix (query parameter takes precedence) | [optional] 
**Application** | Pointer to **string** | Optional application name for context | [optional] 

## Methods

### NewValidateComposeRequest

`func NewValidateComposeRequest(compose string, ) *ValidateComposeRequest`

NewValidateComposeRequest instantiates a new ValidateComposeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateComposeRequestWithDefaults

`func NewValidateComposeRequestWithDefaults() *ValidateComposeRequest`

NewValidateComposeRequestWithDefaults instantiates a new ValidateComposeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompose

`func (o *ValidateComposeRequest) GetCompose() string`

GetCompose returns the Compose field if non-nil, zero value otherwise.

### GetComposeOk

`func (o *ValidateComposeRequest) GetComposeOk() (*string, bool)`

GetComposeOk returns a tuple with the Compose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompose

`func (o *ValidateComposeRequest) SetCompose(v string)`

SetCompose sets Compose field to given value.


### GetImageSuffix

`func (o *ValidateComposeRequest) GetImageSuffix() string`

GetImageSuffix returns the ImageSuffix field if non-nil, zero value otherwise.

### GetImageSuffixOk

`func (o *ValidateComposeRequest) GetImageSuffixOk() (*string, bool)`

GetImageSuffixOk returns a tuple with the ImageSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSuffix

`func (o *ValidateComposeRequest) SetImageSuffix(v string)`

SetImageSuffix sets ImageSuffix field to given value.

### HasImageSuffix

`func (o *ValidateComposeRequest) HasImageSuffix() bool`

HasImageSuffix returns a boolean if a field has been set.

### GetApplication

`func (o *ValidateComposeRequest) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ValidateComposeRequest) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ValidateComposeRequest) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *ValidateComposeRequest) HasApplication() bool`

HasApplication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


