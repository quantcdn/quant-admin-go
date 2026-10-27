# V2Store

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Store ID | 
**Name** | **string** | Store name | 
**Status** | Pointer to **string** | KV_create and KV_show only. ready: the store accepts data requests. provisioning: the table is still being created; data requests are 503 UNAVAILABLE retryable until it is ready. | [optional] 

## Methods

### NewV2Store

`func NewV2Store(id string, name string, ) *V2Store`

NewV2Store instantiates a new V2Store object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2StoreWithDefaults

`func NewV2StoreWithDefaults() *V2Store`

NewV2StoreWithDefaults instantiates a new V2Store object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2Store) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2Store) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2Store) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *V2Store) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2Store) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2Store) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *V2Store) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V2Store) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V2Store) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V2Store) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


