# V2SecretStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
**Id** | **string** | Secret store ID | 
**Name** | **string** | Secret store name | 

## Methods

### NewV2SecretStore

`func NewV2SecretStore(message string, error_ bool, id string, name string, ) *V2SecretStore`

NewV2SecretStore instantiates a new V2SecretStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2SecretStoreWithDefaults

`func NewV2SecretStoreWithDefaults() *V2SecretStore`

NewV2SecretStoreWithDefaults instantiates a new V2SecretStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *V2SecretStore) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2SecretStore) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2SecretStore) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2SecretStore) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2SecretStore) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2SecretStore) SetError(v bool)`

SetError sets Error field to given value.


### GetId

`func (o *V2SecretStore) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2SecretStore) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2SecretStore) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *V2SecretStore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2SecretStore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2SecretStore) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


