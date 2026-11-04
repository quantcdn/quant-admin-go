# GetProjectEdgeEvents200ResponseEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **NullableString** |  | [optional] 
**FunctionUuid** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Truncated** | Pointer to **bool** |  | [optional] 
**Host** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**ClientIp** | Pointer to **NullableString** |  | [optional] 
**RuleId** | Pointer to **NullableString** |  | [optional] 
**WafMode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetProjectEdgeEvents200ResponseEventsInner

`func NewGetProjectEdgeEvents200ResponseEventsInner() *GetProjectEdgeEvents200ResponseEventsInner`

NewGetProjectEdgeEvents200ResponseEventsInner instantiates a new GetProjectEdgeEvents200ResponseEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectEdgeEvents200ResponseEventsInnerWithDefaults

`func NewGetProjectEdgeEvents200ResponseEventsInnerWithDefaults() *GetProjectEdgeEvents200ResponseEventsInner`

NewGetProjectEdgeEvents200ResponseEventsInnerWithDefaults instantiates a new GetProjectEdgeEvents200ResponseEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *GetProjectEdgeEvents200ResponseEventsInner) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetProjectEdgeEvents200ResponseEventsInner) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetProjectEdgeEvents200ResponseEventsInner) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GetProjectEdgeEvents200ResponseEventsInner) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *GetProjectEdgeEvents200ResponseEventsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetProjectEdgeEvents200ResponseEventsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetProjectEdgeEvents200ResponseEventsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetProjectEdgeEvents200ResponseEventsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLevel

`func (o *GetProjectEdgeEvents200ResponseEventsInner) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *GetProjectEdgeEvents200ResponseEventsInner) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *GetProjectEdgeEvents200ResponseEventsInner) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *GetProjectEdgeEvents200ResponseEventsInner) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### SetLevelNil

`func (o *GetProjectEdgeEvents200ResponseEventsInner) SetLevelNil(b bool)`

 SetLevelNil sets the value for Level to be an explicit nil

### UnsetLevel
`func (o *GetProjectEdgeEvents200ResponseEventsInner) UnsetLevel()`

UnsetLevel ensures that no value is present for Level, not even an explicit nil
### GetFunctionUuid

`func (o *GetProjectEdgeEvents200ResponseEventsInner) GetFunctionUuid() string`

GetFunctionUuid returns the FunctionUuid field if non-nil, zero value otherwise.

### GetFunctionUuidOk

`func (o *GetProjectEdgeEvents200ResponseEventsInner) GetFunctionUuidOk() (*string, bool)`

GetFunctionUuidOk returns a tuple with the FunctionUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionUuid

`func (o *GetProjectEdgeEvents200ResponseEventsInner) SetFunctionUuid(v string)`

SetFunctionUuid sets FunctionUuid field to given value.

### HasFunctionUuid

`func (o *GetProjectEdgeEvents200ResponseEventsInner) HasFunctionUuid() bool`

HasFunctionUuid returns a boolean if a field has been set.

### SetFunctionUuidNil

`func (o *GetProjectEdgeEvents200ResponseEventsInner) SetFunctionUuidNil(b bool)`

 SetFunctionUuidNil sets the value for FunctionUuid to be an explicit nil

### UnsetFunctionUuid
`func (o *GetProjectEdgeEvents200ResponseEventsInner) UnsetFunctionUuid()`

UnsetFunctionUuid ensures that no value is present for FunctionUuid, not even an explicit nil
### GetMessage

`func (o *GetProjectEdgeEvents200ResponseEventsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetProjectEdgeEvents200ResponseEventsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetProjectEdgeEvents200ResponseEventsInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetProjectEdgeEvents200ResponseEventsInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTruncated

`func (o *GetProjectEdgeEvents200ResponseEventsInner) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *GetProjectEdgeEvents200ResponseEventsInner) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *GetProjectEdgeEvents200ResponseEventsInner) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *GetProjectEdgeEvents200ResponseEventsInner) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### GetHost

`func (o *GetProjectEdgeEvents200ResponseEventsInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GetProjectEdgeEvents200ResponseEventsInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GetProjectEdgeEvents200ResponseEventsInner) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GetProjectEdgeEvents200ResponseEventsInner) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *GetProjectEdgeEvents200ResponseEventsInner) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *GetProjectEdgeEvents200ResponseEventsInner) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPath

`func (o *GetProjectEdgeEvents200ResponseEventsInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GetProjectEdgeEvents200ResponseEventsInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GetProjectEdgeEvents200ResponseEventsInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GetProjectEdgeEvents200ResponseEventsInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *GetProjectEdgeEvents200ResponseEventsInner) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *GetProjectEdgeEvents200ResponseEventsInner) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetClientIp

`func (o *GetProjectEdgeEvents200ResponseEventsInner) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *GetProjectEdgeEvents200ResponseEventsInner) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *GetProjectEdgeEvents200ResponseEventsInner) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *GetProjectEdgeEvents200ResponseEventsInner) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### SetClientIpNil

`func (o *GetProjectEdgeEvents200ResponseEventsInner) SetClientIpNil(b bool)`

 SetClientIpNil sets the value for ClientIp to be an explicit nil

### UnsetClientIp
`func (o *GetProjectEdgeEvents200ResponseEventsInner) UnsetClientIp()`

UnsetClientIp ensures that no value is present for ClientIp, not even an explicit nil
### GetRuleId

`func (o *GetProjectEdgeEvents200ResponseEventsInner) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *GetProjectEdgeEvents200ResponseEventsInner) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *GetProjectEdgeEvents200ResponseEventsInner) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *GetProjectEdgeEvents200ResponseEventsInner) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### SetRuleIdNil

`func (o *GetProjectEdgeEvents200ResponseEventsInner) SetRuleIdNil(b bool)`

 SetRuleIdNil sets the value for RuleId to be an explicit nil

### UnsetRuleId
`func (o *GetProjectEdgeEvents200ResponseEventsInner) UnsetRuleId()`

UnsetRuleId ensures that no value is present for RuleId, not even an explicit nil
### GetWafMode

`func (o *GetProjectEdgeEvents200ResponseEventsInner) GetWafMode() string`

GetWafMode returns the WafMode field if non-nil, zero value otherwise.

### GetWafModeOk

`func (o *GetProjectEdgeEvents200ResponseEventsInner) GetWafModeOk() (*string, bool)`

GetWafModeOk returns a tuple with the WafMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafMode

`func (o *GetProjectEdgeEvents200ResponseEventsInner) SetWafMode(v string)`

SetWafMode sets WafMode field to given value.

### HasWafMode

`func (o *GetProjectEdgeEvents200ResponseEventsInner) HasWafMode() bool`

HasWafMode returns a boolean if a field has been set.

### SetWafModeNil

`func (o *GetProjectEdgeEvents200ResponseEventsInner) SetWafModeNil(b bool)`

 SetWafModeNil sets the value for WafMode to be an explicit nil

### UnsetWafMode
`func (o *GetProjectEdgeEvents200ResponseEventsInner) UnsetWafMode()`

UnsetWafMode ensures that no value is present for WafMode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


