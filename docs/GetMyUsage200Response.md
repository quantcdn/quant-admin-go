# GetMyUsage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** |  | [optional] 
**CurrentMonth** | Pointer to **string** |  | [optional] 
**Monthly** | Pointer to [**GetMyUsage200ResponseMonthly**](GetMyUsage200ResponseMonthly.md) |  | [optional] 
**Daily** | Pointer to [**GetMyUsage200ResponseDaily**](GetMyUsage200ResponseDaily.md) |  | [optional] 
**Quota** | Pointer to [**NullableGetMyUsage200ResponseQuota**](GetMyUsage200ResponseQuota.md) |  | [optional] 

## Methods

### NewGetMyUsage200Response

`func NewGetMyUsage200Response() *GetMyUsage200Response`

NewGetMyUsage200Response instantiates a new GetMyUsage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMyUsage200ResponseWithDefaults

`func NewGetMyUsage200ResponseWithDefaults() *GetMyUsage200Response`

NewGetMyUsage200ResponseWithDefaults instantiates a new GetMyUsage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *GetMyUsage200Response) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetMyUsage200Response) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetMyUsage200Response) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GetMyUsage200Response) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCurrentMonth

`func (o *GetMyUsage200Response) GetCurrentMonth() string`

GetCurrentMonth returns the CurrentMonth field if non-nil, zero value otherwise.

### GetCurrentMonthOk

`func (o *GetMyUsage200Response) GetCurrentMonthOk() (*string, bool)`

GetCurrentMonthOk returns a tuple with the CurrentMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMonth

`func (o *GetMyUsage200Response) SetCurrentMonth(v string)`

SetCurrentMonth sets CurrentMonth field to given value.

### HasCurrentMonth

`func (o *GetMyUsage200Response) HasCurrentMonth() bool`

HasCurrentMonth returns a boolean if a field has been set.

### GetMonthly

`func (o *GetMyUsage200Response) GetMonthly() GetMyUsage200ResponseMonthly`

GetMonthly returns the Monthly field if non-nil, zero value otherwise.

### GetMonthlyOk

`func (o *GetMyUsage200Response) GetMonthlyOk() (*GetMyUsage200ResponseMonthly, bool)`

GetMonthlyOk returns a tuple with the Monthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthly

`func (o *GetMyUsage200Response) SetMonthly(v GetMyUsage200ResponseMonthly)`

SetMonthly sets Monthly field to given value.

### HasMonthly

`func (o *GetMyUsage200Response) HasMonthly() bool`

HasMonthly returns a boolean if a field has been set.

### GetDaily

`func (o *GetMyUsage200Response) GetDaily() GetMyUsage200ResponseDaily`

GetDaily returns the Daily field if non-nil, zero value otherwise.

### GetDailyOk

`func (o *GetMyUsage200Response) GetDailyOk() (*GetMyUsage200ResponseDaily, bool)`

GetDailyOk returns a tuple with the Daily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaily

`func (o *GetMyUsage200Response) SetDaily(v GetMyUsage200ResponseDaily)`

SetDaily sets Daily field to given value.

### HasDaily

`func (o *GetMyUsage200Response) HasDaily() bool`

HasDaily returns a boolean if a field has been set.

### GetQuota

`func (o *GetMyUsage200Response) GetQuota() GetMyUsage200ResponseQuota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *GetMyUsage200Response) GetQuotaOk() (*GetMyUsage200ResponseQuota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *GetMyUsage200Response) SetQuota(v GetMyUsage200ResponseQuota)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *GetMyUsage200Response) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### SetQuotaNil

`func (o *GetMyUsage200Response) SetQuotaNil(b bool)`

 SetQuotaNil sets the value for Quota to be an explicit nil

### UnsetQuota
`func (o *GetMyUsage200Response) UnsetQuota()`

UnsetQuota ensures that no value is present for Quota, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


