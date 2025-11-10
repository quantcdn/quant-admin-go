# V1RedirectListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Draw** | **int32** |  | 
**ITotalRecords** | **int32** |  | 
**ITotalDisplayRecords** | **int32** |  | 
**AaData** | [**[]V1RedirectItem**](V1RedirectItem.md) |  | 

## Methods

### NewV1RedirectListResponse

`func NewV1RedirectListResponse(draw int32, iTotalRecords int32, iTotalDisplayRecords int32, aaData []V1RedirectItem, ) *V1RedirectListResponse`

NewV1RedirectListResponse instantiates a new V1RedirectListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RedirectListResponseWithDefaults

`func NewV1RedirectListResponseWithDefaults() *V1RedirectListResponse`

NewV1RedirectListResponseWithDefaults instantiates a new V1RedirectListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDraw

`func (o *V1RedirectListResponse) GetDraw() int32`

GetDraw returns the Draw field if non-nil, zero value otherwise.

### GetDrawOk

`func (o *V1RedirectListResponse) GetDrawOk() (*int32, bool)`

GetDrawOk returns a tuple with the Draw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraw

`func (o *V1RedirectListResponse) SetDraw(v int32)`

SetDraw sets Draw field to given value.


### GetITotalRecords

`func (o *V1RedirectListResponse) GetITotalRecords() int32`

GetITotalRecords returns the ITotalRecords field if non-nil, zero value otherwise.

### GetITotalRecordsOk

`func (o *V1RedirectListResponse) GetITotalRecordsOk() (*int32, bool)`

GetITotalRecordsOk returns a tuple with the ITotalRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetITotalRecords

`func (o *V1RedirectListResponse) SetITotalRecords(v int32)`

SetITotalRecords sets ITotalRecords field to given value.


### GetITotalDisplayRecords

`func (o *V1RedirectListResponse) GetITotalDisplayRecords() int32`

GetITotalDisplayRecords returns the ITotalDisplayRecords field if non-nil, zero value otherwise.

### GetITotalDisplayRecordsOk

`func (o *V1RedirectListResponse) GetITotalDisplayRecordsOk() (*int32, bool)`

GetITotalDisplayRecordsOk returns a tuple with the ITotalDisplayRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetITotalDisplayRecords

`func (o *V1RedirectListResponse) SetITotalDisplayRecords(v int32)`

SetITotalDisplayRecords sets ITotalDisplayRecords field to given value.


### GetAaData

`func (o *V1RedirectListResponse) GetAaData() []V1RedirectItem`

GetAaData returns the AaData field if non-nil, zero value otherwise.

### GetAaDataOk

`func (o *V1RedirectListResponse) GetAaDataOk() (*[]V1RedirectItem, bool)`

GetAaDataOk returns a tuple with the AaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaData

`func (o *V1RedirectListResponse) SetAaData(v []V1RedirectItem)`

SetAaData sets AaData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


