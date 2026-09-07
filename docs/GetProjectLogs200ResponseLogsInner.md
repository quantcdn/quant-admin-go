# GetProjectLogs200ResponseLogsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Project** | Pointer to **NullableString** |  | [optional] 
**Domain** | Pointer to **NullableString** |  | [optional] 
**Method** | Pointer to **NullableString** |  | [optional] 
**Uri** | Pointer to **NullableString** |  | [optional] 
**StatusCode** | Pointer to **NullableInt32** |  | [optional] 
**CacheStatus** | Pointer to **NullableString** |  | [optional] 
**ClientIp** | Pointer to **NullableString** |  | [optional] 
**UserAgent** | Pointer to **NullableString** |  | [optional] 
**Bytes** | Pointer to **NullableInt32** |  | [optional] 
**TimeTaken** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewGetProjectLogs200ResponseLogsInner

`func NewGetProjectLogs200ResponseLogsInner() *GetProjectLogs200ResponseLogsInner`

NewGetProjectLogs200ResponseLogsInner instantiates a new GetProjectLogs200ResponseLogsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectLogs200ResponseLogsInnerWithDefaults

`func NewGetProjectLogs200ResponseLogsInnerWithDefaults() *GetProjectLogs200ResponseLogsInner`

NewGetProjectLogs200ResponseLogsInnerWithDefaults instantiates a new GetProjectLogs200ResponseLogsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *GetProjectLogs200ResponseLogsInner) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetProjectLogs200ResponseLogsInner) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetProjectLogs200ResponseLogsInner) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GetProjectLogs200ResponseLogsInner) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *GetProjectLogs200ResponseLogsInner) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *GetProjectLogs200ResponseLogsInner) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetProject

`func (o *GetProjectLogs200ResponseLogsInner) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GetProjectLogs200ResponseLogsInner) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GetProjectLogs200ResponseLogsInner) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *GetProjectLogs200ResponseLogsInner) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *GetProjectLogs200ResponseLogsInner) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *GetProjectLogs200ResponseLogsInner) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetDomain

`func (o *GetProjectLogs200ResponseLogsInner) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GetProjectLogs200ResponseLogsInner) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GetProjectLogs200ResponseLogsInner) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GetProjectLogs200ResponseLogsInner) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *GetProjectLogs200ResponseLogsInner) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *GetProjectLogs200ResponseLogsInner) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetMethod

`func (o *GetProjectLogs200ResponseLogsInner) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *GetProjectLogs200ResponseLogsInner) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *GetProjectLogs200ResponseLogsInner) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *GetProjectLogs200ResponseLogsInner) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *GetProjectLogs200ResponseLogsInner) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *GetProjectLogs200ResponseLogsInner) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil
### GetUri

`func (o *GetProjectLogs200ResponseLogsInner) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *GetProjectLogs200ResponseLogsInner) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *GetProjectLogs200ResponseLogsInner) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *GetProjectLogs200ResponseLogsInner) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *GetProjectLogs200ResponseLogsInner) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *GetProjectLogs200ResponseLogsInner) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetStatusCode

`func (o *GetProjectLogs200ResponseLogsInner) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *GetProjectLogs200ResponseLogsInner) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *GetProjectLogs200ResponseLogsInner) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *GetProjectLogs200ResponseLogsInner) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### SetStatusCodeNil

`func (o *GetProjectLogs200ResponseLogsInner) SetStatusCodeNil(b bool)`

 SetStatusCodeNil sets the value for StatusCode to be an explicit nil

### UnsetStatusCode
`func (o *GetProjectLogs200ResponseLogsInner) UnsetStatusCode()`

UnsetStatusCode ensures that no value is present for StatusCode, not even an explicit nil
### GetCacheStatus

`func (o *GetProjectLogs200ResponseLogsInner) GetCacheStatus() string`

GetCacheStatus returns the CacheStatus field if non-nil, zero value otherwise.

### GetCacheStatusOk

`func (o *GetProjectLogs200ResponseLogsInner) GetCacheStatusOk() (*string, bool)`

GetCacheStatusOk returns a tuple with the CacheStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheStatus

`func (o *GetProjectLogs200ResponseLogsInner) SetCacheStatus(v string)`

SetCacheStatus sets CacheStatus field to given value.

### HasCacheStatus

`func (o *GetProjectLogs200ResponseLogsInner) HasCacheStatus() bool`

HasCacheStatus returns a boolean if a field has been set.

### SetCacheStatusNil

`func (o *GetProjectLogs200ResponseLogsInner) SetCacheStatusNil(b bool)`

 SetCacheStatusNil sets the value for CacheStatus to be an explicit nil

### UnsetCacheStatus
`func (o *GetProjectLogs200ResponseLogsInner) UnsetCacheStatus()`

UnsetCacheStatus ensures that no value is present for CacheStatus, not even an explicit nil
### GetClientIp

`func (o *GetProjectLogs200ResponseLogsInner) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *GetProjectLogs200ResponseLogsInner) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *GetProjectLogs200ResponseLogsInner) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *GetProjectLogs200ResponseLogsInner) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### SetClientIpNil

`func (o *GetProjectLogs200ResponseLogsInner) SetClientIpNil(b bool)`

 SetClientIpNil sets the value for ClientIp to be an explicit nil

### UnsetClientIp
`func (o *GetProjectLogs200ResponseLogsInner) UnsetClientIp()`

UnsetClientIp ensures that no value is present for ClientIp, not even an explicit nil
### GetUserAgent

`func (o *GetProjectLogs200ResponseLogsInner) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *GetProjectLogs200ResponseLogsInner) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *GetProjectLogs200ResponseLogsInner) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *GetProjectLogs200ResponseLogsInner) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### SetUserAgentNil

`func (o *GetProjectLogs200ResponseLogsInner) SetUserAgentNil(b bool)`

 SetUserAgentNil sets the value for UserAgent to be an explicit nil

### UnsetUserAgent
`func (o *GetProjectLogs200ResponseLogsInner) UnsetUserAgent()`

UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil
### GetBytes

`func (o *GetProjectLogs200ResponseLogsInner) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *GetProjectLogs200ResponseLogsInner) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *GetProjectLogs200ResponseLogsInner) SetBytes(v int32)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *GetProjectLogs200ResponseLogsInner) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### SetBytesNil

`func (o *GetProjectLogs200ResponseLogsInner) SetBytesNil(b bool)`

 SetBytesNil sets the value for Bytes to be an explicit nil

### UnsetBytes
`func (o *GetProjectLogs200ResponseLogsInner) UnsetBytes()`

UnsetBytes ensures that no value is present for Bytes, not even an explicit nil
### GetTimeTaken

`func (o *GetProjectLogs200ResponseLogsInner) GetTimeTaken() float32`

GetTimeTaken returns the TimeTaken field if non-nil, zero value otherwise.

### GetTimeTakenOk

`func (o *GetProjectLogs200ResponseLogsInner) GetTimeTakenOk() (*float32, bool)`

GetTimeTakenOk returns a tuple with the TimeTaken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTaken

`func (o *GetProjectLogs200ResponseLogsInner) SetTimeTaken(v float32)`

SetTimeTaken sets TimeTaken field to given value.

### HasTimeTaken

`func (o *GetProjectLogs200ResponseLogsInner) HasTimeTaken() bool`

HasTimeTaken returns a boolean if a field has been set.

### SetTimeTakenNil

`func (o *GetProjectLogs200ResponseLogsInner) SetTimeTakenNil(b bool)`

 SetTimeTakenNil sets the value for TimeTaken to be an explicit nil

### UnsetTimeTaken
`func (o *GetProjectLogs200ResponseLogsInner) UnsetTimeTaken()`

UnsetTimeTaken ensures that no value is present for TimeTaken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


