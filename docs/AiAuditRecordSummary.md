# AiAuditRecordSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
**RecordType** | **string** | inference for a record without recordType; tool_denied, tool_proposal, tool_decision, tool_outcome, slack_bot_api, audit_access, ...; unknown for an unreadable object | 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**TokenId** | Pointer to **string** |  | [optional] 
**ModelId** | Pointer to **string** |  | [optional] 
**Interface** | Pointer to **string** |  | [optional] 
**Tools** | **[]string** |  | 
**Decision** | Pointer to **string** |  | [optional] 
**Outcome** | Pointer to **string** |  | [optional] 

## Methods

### NewAiAuditRecordSummary

`func NewAiAuditRecordSummary(message string, error_ bool, recordType string, tools []string, ) *AiAuditRecordSummary`

NewAiAuditRecordSummary instantiates a new AiAuditRecordSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAuditRecordSummaryWithDefaults

`func NewAiAuditRecordSummaryWithDefaults() *AiAuditRecordSummary`

NewAiAuditRecordSummaryWithDefaults instantiates a new AiAuditRecordSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AiAuditRecordSummary) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AiAuditRecordSummary) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AiAuditRecordSummary) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *AiAuditRecordSummary) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AiAuditRecordSummary) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AiAuditRecordSummary) SetError(v bool)`

SetError sets Error field to given value.


### GetRecordType

`func (o *AiAuditRecordSummary) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *AiAuditRecordSummary) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *AiAuditRecordSummary) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### GetTimestamp

`func (o *AiAuditRecordSummary) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AiAuditRecordSummary) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AiAuditRecordSummary) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AiAuditRecordSummary) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUserId

`func (o *AiAuditRecordSummary) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AiAuditRecordSummary) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AiAuditRecordSummary) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AiAuditRecordSummary) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetTokenId

`func (o *AiAuditRecordSummary) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *AiAuditRecordSummary) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *AiAuditRecordSummary) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *AiAuditRecordSummary) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetModelId

`func (o *AiAuditRecordSummary) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *AiAuditRecordSummary) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *AiAuditRecordSummary) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *AiAuditRecordSummary) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetInterface

`func (o *AiAuditRecordSummary) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *AiAuditRecordSummary) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *AiAuditRecordSummary) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *AiAuditRecordSummary) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetTools

`func (o *AiAuditRecordSummary) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AiAuditRecordSummary) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AiAuditRecordSummary) SetTools(v []string)`

SetTools sets Tools field to given value.


### GetDecision

`func (o *AiAuditRecordSummary) GetDecision() string`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *AiAuditRecordSummary) GetDecisionOk() (*string, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *AiAuditRecordSummary) SetDecision(v string)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *AiAuditRecordSummary) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### GetOutcome

`func (o *AiAuditRecordSummary) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *AiAuditRecordSummary) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *AiAuditRecordSummary) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *AiAuditRecordSummary) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


