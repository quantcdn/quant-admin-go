# ChatInference200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**ChatInference200ResponseResponse**](ChatInference200ResponseResponse.md) |  | [optional] 
**Model** | Pointer to **string** | Model used for generation | [optional] 
**RequestId** | Pointer to **string** | Unique request identifier | [optional] 
**FinishReason** | Pointer to **string** | Why the model stopped generating | [optional] 
**Usage** | Pointer to [**ChatInference200ResponseUsage**](ChatInference200ResponseUsage.md) |  | [optional] 

## Methods

### NewChatInference200Response

`func NewChatInference200Response() *ChatInference200Response`

NewChatInference200Response instantiates a new ChatInference200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatInference200ResponseWithDefaults

`func NewChatInference200ResponseWithDefaults() *ChatInference200Response`

NewChatInference200ResponseWithDefaults instantiates a new ChatInference200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *ChatInference200Response) GetResponse() ChatInference200ResponseResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ChatInference200Response) GetResponseOk() (*ChatInference200ResponseResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ChatInference200Response) SetResponse(v ChatInference200ResponseResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ChatInference200Response) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetModel

`func (o *ChatInference200Response) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatInference200Response) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatInference200Response) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ChatInference200Response) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRequestId

`func (o *ChatInference200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ChatInference200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ChatInference200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ChatInference200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetFinishReason

`func (o *ChatInference200Response) GetFinishReason() string`

GetFinishReason returns the FinishReason field if non-nil, zero value otherwise.

### GetFinishReasonOk

`func (o *ChatInference200Response) GetFinishReasonOk() (*string, bool)`

GetFinishReasonOk returns a tuple with the FinishReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishReason

`func (o *ChatInference200Response) SetFinishReason(v string)`

SetFinishReason sets FinishReason field to given value.

### HasFinishReason

`func (o *ChatInference200Response) HasFinishReason() bool`

HasFinishReason returns a boolean if a field has been set.

### GetUsage

`func (o *ChatInference200Response) GetUsage() ChatInference200ResponseUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ChatInference200Response) GetUsageOk() (*ChatInference200ResponseUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ChatInference200Response) SetUsage(v ChatInference200ResponseUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ChatInference200Response) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


