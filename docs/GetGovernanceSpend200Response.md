# GetGovernanceSpend200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgTotal** | Pointer to [**GetGovernanceSpend200ResponseOrgTotal**](GetGovernanceSpend200ResponseOrgTotal.md) |  | [optional] 
**TodayTotal** | Pointer to [**GetGovernanceSpend200ResponseTodayTotal**](GetGovernanceSpend200ResponseTodayTotal.md) |  | [optional] 
**Budget** | Pointer to [**NullableGetGovernanceSpend200ResponseBudget**](GetGovernanceSpend200ResponseBudget.md) |  | [optional] 
**DailyBudget** | Pointer to [**NullableGetGovernanceSpend200ResponseDailyBudget**](GetGovernanceSpend200ResponseDailyBudget.md) |  | [optional] 
**UserTotal** | Pointer to [**NullableGetGovernanceSpend200ResponseUserTotal**](GetGovernanceSpend200ResponseUserTotal.md) |  | [optional] 

## Methods

### NewGetGovernanceSpend200Response

`func NewGetGovernanceSpend200Response() *GetGovernanceSpend200Response`

NewGetGovernanceSpend200Response instantiates a new GetGovernanceSpend200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGovernanceSpend200ResponseWithDefaults

`func NewGetGovernanceSpend200ResponseWithDefaults() *GetGovernanceSpend200Response`

NewGetGovernanceSpend200ResponseWithDefaults instantiates a new GetGovernanceSpend200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgTotal

`func (o *GetGovernanceSpend200Response) GetOrgTotal() GetGovernanceSpend200ResponseOrgTotal`

GetOrgTotal returns the OrgTotal field if non-nil, zero value otherwise.

### GetOrgTotalOk

`func (o *GetGovernanceSpend200Response) GetOrgTotalOk() (*GetGovernanceSpend200ResponseOrgTotal, bool)`

GetOrgTotalOk returns a tuple with the OrgTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgTotal

`func (o *GetGovernanceSpend200Response) SetOrgTotal(v GetGovernanceSpend200ResponseOrgTotal)`

SetOrgTotal sets OrgTotal field to given value.

### HasOrgTotal

`func (o *GetGovernanceSpend200Response) HasOrgTotal() bool`

HasOrgTotal returns a boolean if a field has been set.

### GetTodayTotal

`func (o *GetGovernanceSpend200Response) GetTodayTotal() GetGovernanceSpend200ResponseTodayTotal`

GetTodayTotal returns the TodayTotal field if non-nil, zero value otherwise.

### GetTodayTotalOk

`func (o *GetGovernanceSpend200Response) GetTodayTotalOk() (*GetGovernanceSpend200ResponseTodayTotal, bool)`

GetTodayTotalOk returns a tuple with the TodayTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodayTotal

`func (o *GetGovernanceSpend200Response) SetTodayTotal(v GetGovernanceSpend200ResponseTodayTotal)`

SetTodayTotal sets TodayTotal field to given value.

### HasTodayTotal

`func (o *GetGovernanceSpend200Response) HasTodayTotal() bool`

HasTodayTotal returns a boolean if a field has been set.

### GetBudget

`func (o *GetGovernanceSpend200Response) GetBudget() GetGovernanceSpend200ResponseBudget`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *GetGovernanceSpend200Response) GetBudgetOk() (*GetGovernanceSpend200ResponseBudget, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *GetGovernanceSpend200Response) SetBudget(v GetGovernanceSpend200ResponseBudget)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *GetGovernanceSpend200Response) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### SetBudgetNil

`func (o *GetGovernanceSpend200Response) SetBudgetNil(b bool)`

 SetBudgetNil sets the value for Budget to be an explicit nil

### UnsetBudget
`func (o *GetGovernanceSpend200Response) UnsetBudget()`

UnsetBudget ensures that no value is present for Budget, not even an explicit nil
### GetDailyBudget

`func (o *GetGovernanceSpend200Response) GetDailyBudget() GetGovernanceSpend200ResponseDailyBudget`

GetDailyBudget returns the DailyBudget field if non-nil, zero value otherwise.

### GetDailyBudgetOk

`func (o *GetGovernanceSpend200Response) GetDailyBudgetOk() (*GetGovernanceSpend200ResponseDailyBudget, bool)`

GetDailyBudgetOk returns a tuple with the DailyBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyBudget

`func (o *GetGovernanceSpend200Response) SetDailyBudget(v GetGovernanceSpend200ResponseDailyBudget)`

SetDailyBudget sets DailyBudget field to given value.

### HasDailyBudget

`func (o *GetGovernanceSpend200Response) HasDailyBudget() bool`

HasDailyBudget returns a boolean if a field has been set.

### SetDailyBudgetNil

`func (o *GetGovernanceSpend200Response) SetDailyBudgetNil(b bool)`

 SetDailyBudgetNil sets the value for DailyBudget to be an explicit nil

### UnsetDailyBudget
`func (o *GetGovernanceSpend200Response) UnsetDailyBudget()`

UnsetDailyBudget ensures that no value is present for DailyBudget, not even an explicit nil
### GetUserTotal

`func (o *GetGovernanceSpend200Response) GetUserTotal() GetGovernanceSpend200ResponseUserTotal`

GetUserTotal returns the UserTotal field if non-nil, zero value otherwise.

### GetUserTotalOk

`func (o *GetGovernanceSpend200Response) GetUserTotalOk() (*GetGovernanceSpend200ResponseUserTotal, bool)`

GetUserTotalOk returns a tuple with the UserTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTotal

`func (o *GetGovernanceSpend200Response) SetUserTotal(v GetGovernanceSpend200ResponseUserTotal)`

SetUserTotal sets UserTotal field to given value.

### HasUserTotal

`func (o *GetGovernanceSpend200Response) HasUserTotal() bool`

HasUserTotal returns a boolean if a field has been set.

### SetUserTotalNil

`func (o *GetGovernanceSpend200Response) SetUserTotalNil(b bool)`

 SetUserTotalNil sets the value for UserTotal to be an explicit nil

### UnsetUserTotal
`func (o *GetGovernanceSpend200Response) UnsetUserTotal()`

UnsetUserTotal ensures that no value is present for UserTotal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


