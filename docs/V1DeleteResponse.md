# V1DeleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **bool** | Indicates at least one error occurred | [optional] 
**Errors** | Pointer to **[]map[string]interface{}** | List of any errors | [optional] 
**Meta** | Pointer to [**[]V1Revision**](V1Revision.md) |  | [optional] 

## Methods

### NewV1DeleteResponse

`func NewV1DeleteResponse() *V1DeleteResponse`

NewV1DeleteResponse instantiates a new V1DeleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeleteResponseWithDefaults

`func NewV1DeleteResponseWithDefaults() *V1DeleteResponse`

NewV1DeleteResponseWithDefaults instantiates a new V1DeleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *V1DeleteResponse) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V1DeleteResponse) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V1DeleteResponse) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *V1DeleteResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrors

`func (o *V1DeleteResponse) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V1DeleteResponse) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V1DeleteResponse) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V1DeleteResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetMeta

`func (o *V1DeleteResponse) GetMeta() []V1Revision`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V1DeleteResponse) GetMetaOk() (*[]V1Revision, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V1DeleteResponse) SetMeta(v []V1Revision)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V1DeleteResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


