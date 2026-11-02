# AiAuditRecordGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Date** | **string** |  | 
**Summary** | [**AiAuditRecordSummary**](AiAuditRecordSummary.md) |  | 
**Record** | **map[string]interface{}** |  | 

## Methods

### NewAiAuditRecordGet200Response

`func NewAiAuditRecordGet200Response(key string, date string, summary AiAuditRecordSummary, record map[string]interface{}, ) *AiAuditRecordGet200Response`

NewAiAuditRecordGet200Response instantiates a new AiAuditRecordGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAuditRecordGet200ResponseWithDefaults

`func NewAiAuditRecordGet200ResponseWithDefaults() *AiAuditRecordGet200Response`

NewAiAuditRecordGet200ResponseWithDefaults instantiates a new AiAuditRecordGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AiAuditRecordGet200Response) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AiAuditRecordGet200Response) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AiAuditRecordGet200Response) SetKey(v string)`

SetKey sets Key field to given value.


### GetDate

`func (o *AiAuditRecordGet200Response) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AiAuditRecordGet200Response) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AiAuditRecordGet200Response) SetDate(v string)`

SetDate sets Date field to given value.


### GetSummary

`func (o *AiAuditRecordGet200Response) GetSummary() AiAuditRecordSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AiAuditRecordGet200Response) GetSummaryOk() (*AiAuditRecordSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AiAuditRecordGet200Response) SetSummary(v AiAuditRecordSummary)`

SetSummary sets Summary field to given value.


### GetRecord

`func (o *AiAuditRecordGet200Response) GetRecord() map[string]interface{}`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *AiAuditRecordGet200Response) GetRecordOk() (*map[string]interface{}, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *AiAuditRecordGet200Response) SetRecord(v map[string]interface{})`

SetRecord sets Record field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


