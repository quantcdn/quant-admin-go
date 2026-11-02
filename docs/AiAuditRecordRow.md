# AiAuditRecordRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
**Key** | **string** |  | 
**Date** | **string** |  | 
**Summary** | [**AiAuditRecordSummary**](AiAuditRecordSummary.md) |  | 
**Record** | Pointer to **map[string]interface{}** | The stored record. prompt and response are present only when the record has promptStored or responseStored true. | [optional] 
**Omitted** | Pointer to **string** | too-large: read the record with aiAuditRecordGet. | [optional] 

## Methods

### NewAiAuditRecordRow

`func NewAiAuditRecordRow(message string, error_ bool, key string, date string, summary AiAuditRecordSummary, ) *AiAuditRecordRow`

NewAiAuditRecordRow instantiates a new AiAuditRecordRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAuditRecordRowWithDefaults

`func NewAiAuditRecordRowWithDefaults() *AiAuditRecordRow`

NewAiAuditRecordRowWithDefaults instantiates a new AiAuditRecordRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AiAuditRecordRow) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AiAuditRecordRow) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AiAuditRecordRow) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *AiAuditRecordRow) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AiAuditRecordRow) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AiAuditRecordRow) SetError(v bool)`

SetError sets Error field to given value.


### GetKey

`func (o *AiAuditRecordRow) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AiAuditRecordRow) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AiAuditRecordRow) SetKey(v string)`

SetKey sets Key field to given value.


### GetDate

`func (o *AiAuditRecordRow) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AiAuditRecordRow) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AiAuditRecordRow) SetDate(v string)`

SetDate sets Date field to given value.


### GetSummary

`func (o *AiAuditRecordRow) GetSummary() AiAuditRecordSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AiAuditRecordRow) GetSummaryOk() (*AiAuditRecordSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AiAuditRecordRow) SetSummary(v AiAuditRecordSummary)`

SetSummary sets Summary field to given value.


### GetRecord

`func (o *AiAuditRecordRow) GetRecord() map[string]interface{}`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *AiAuditRecordRow) GetRecordOk() (*map[string]interface{}, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *AiAuditRecordRow) SetRecord(v map[string]interface{})`

SetRecord sets Record field to given value.

### HasRecord

`func (o *AiAuditRecordRow) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetOmitted

`func (o *AiAuditRecordRow) GetOmitted() string`

GetOmitted returns the Omitted field if non-nil, zero value otherwise.

### GetOmittedOk

`func (o *AiAuditRecordRow) GetOmittedOk() (*string, bool)`

GetOmittedOk returns a tuple with the Omitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitted

`func (o *AiAuditRecordRow) SetOmitted(v string)`

SetOmitted sets Omitted field to given value.

### HasOmitted

`func (o *AiAuditRecordRow) HasOmitted() bool`

HasOmitted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


