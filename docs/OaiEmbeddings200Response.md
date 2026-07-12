# OaiEmbeddings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]OaiEmbeddings200ResponseDataInner**](OaiEmbeddings200ResponseDataInner.md) |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Usage** | Pointer to [**OaiEmbeddings200ResponseUsage**](OaiEmbeddings200ResponseUsage.md) |  | [optional] 

## Methods

### NewOaiEmbeddings200Response

`func NewOaiEmbeddings200Response() *OaiEmbeddings200Response`

NewOaiEmbeddings200Response instantiates a new OaiEmbeddings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOaiEmbeddings200ResponseWithDefaults

`func NewOaiEmbeddings200ResponseWithDefaults() *OaiEmbeddings200Response`

NewOaiEmbeddings200ResponseWithDefaults instantiates a new OaiEmbeddings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *OaiEmbeddings200Response) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OaiEmbeddings200Response) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OaiEmbeddings200Response) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *OaiEmbeddings200Response) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetData

`func (o *OaiEmbeddings200Response) GetData() []OaiEmbeddings200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OaiEmbeddings200Response) GetDataOk() (*[]OaiEmbeddings200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OaiEmbeddings200Response) SetData(v []OaiEmbeddings200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *OaiEmbeddings200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetModel

`func (o *OaiEmbeddings200Response) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OaiEmbeddings200Response) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OaiEmbeddings200Response) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OaiEmbeddings200Response) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetUsage

`func (o *OaiEmbeddings200Response) GetUsage() OaiEmbeddings200ResponseUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *OaiEmbeddings200Response) GetUsageOk() (*OaiEmbeddings200ResponseUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *OaiEmbeddings200Response) SetUsage(v OaiEmbeddings200ResponseUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *OaiEmbeddings200Response) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


