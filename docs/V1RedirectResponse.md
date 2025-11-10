# V1RedirectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Md5** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ByteLength** | Pointer to **int32** |  | [optional] 
**RevisionNumber** | Pointer to **int32** |  | [optional] 
**DateTimestamp** | Pointer to **int32** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DeletedTimestamp** | Pointer to **int32** |  | [optional] 
**Transitions** | Pointer to [**[]V1Transition**](V1Transition.md) |  | [optional] 
**Info** | Pointer to [**V1Info**](V1Info.md) |  | [optional] 
**RedirectHttpCode** | Pointer to **int32** | The HTTP code | [optional] 
**HighestRevisionNumber** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | The redirect from URL | [optional] 

## Methods

### NewV1RedirectResponse

`func NewV1RedirectResponse() *V1RedirectResponse`

NewV1RedirectResponse instantiates a new V1RedirectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RedirectResponseWithDefaults

`func NewV1RedirectResponseWithDefaults() *V1RedirectResponse`

NewV1RedirectResponseWithDefaults instantiates a new V1RedirectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMd5

`func (o *V1RedirectResponse) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *V1RedirectResponse) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *V1RedirectResponse) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *V1RedirectResponse) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetType

`func (o *V1RedirectResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1RedirectResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1RedirectResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1RedirectResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetByteLength

`func (o *V1RedirectResponse) GetByteLength() int32`

GetByteLength returns the ByteLength field if non-nil, zero value otherwise.

### GetByteLengthOk

`func (o *V1RedirectResponse) GetByteLengthOk() (*int32, bool)`

GetByteLengthOk returns a tuple with the ByteLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteLength

`func (o *V1RedirectResponse) SetByteLength(v int32)`

SetByteLength sets ByteLength field to given value.

### HasByteLength

`func (o *V1RedirectResponse) HasByteLength() bool`

HasByteLength returns a boolean if a field has been set.

### GetRevisionNumber

`func (o *V1RedirectResponse) GetRevisionNumber() int32`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *V1RedirectResponse) GetRevisionNumberOk() (*int32, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *V1RedirectResponse) SetRevisionNumber(v int32)`

SetRevisionNumber sets RevisionNumber field to given value.

### HasRevisionNumber

`func (o *V1RedirectResponse) HasRevisionNumber() bool`

HasRevisionNumber returns a boolean if a field has been set.

### GetDateTimestamp

`func (o *V1RedirectResponse) GetDateTimestamp() int32`

GetDateTimestamp returns the DateTimestamp field if non-nil, zero value otherwise.

### GetDateTimestampOk

`func (o *V1RedirectResponse) GetDateTimestampOk() (*int32, bool)`

GetDateTimestampOk returns a tuple with the DateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimestamp

`func (o *V1RedirectResponse) SetDateTimestamp(v int32)`

SetDateTimestamp sets DateTimestamp field to given value.

### HasDateTimestamp

`func (o *V1RedirectResponse) HasDateTimestamp() bool`

HasDateTimestamp returns a boolean if a field has been set.

### GetDeleted

`func (o *V1RedirectResponse) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *V1RedirectResponse) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *V1RedirectResponse) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *V1RedirectResponse) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDeletedTimestamp

`func (o *V1RedirectResponse) GetDeletedTimestamp() int32`

GetDeletedTimestamp returns the DeletedTimestamp field if non-nil, zero value otherwise.

### GetDeletedTimestampOk

`func (o *V1RedirectResponse) GetDeletedTimestampOk() (*int32, bool)`

GetDeletedTimestampOk returns a tuple with the DeletedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedTimestamp

`func (o *V1RedirectResponse) SetDeletedTimestamp(v int32)`

SetDeletedTimestamp sets DeletedTimestamp field to given value.

### HasDeletedTimestamp

`func (o *V1RedirectResponse) HasDeletedTimestamp() bool`

HasDeletedTimestamp returns a boolean if a field has been set.

### GetTransitions

`func (o *V1RedirectResponse) GetTransitions() []V1Transition`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *V1RedirectResponse) GetTransitionsOk() (*[]V1Transition, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *V1RedirectResponse) SetTransitions(v []V1Transition)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *V1RedirectResponse) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.

### GetInfo

`func (o *V1RedirectResponse) GetInfo() V1Info`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *V1RedirectResponse) GetInfoOk() (*V1Info, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *V1RedirectResponse) SetInfo(v V1Info)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *V1RedirectResponse) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetRedirectHttpCode

`func (o *V1RedirectResponse) GetRedirectHttpCode() int32`

GetRedirectHttpCode returns the RedirectHttpCode field if non-nil, zero value otherwise.

### GetRedirectHttpCodeOk

`func (o *V1RedirectResponse) GetRedirectHttpCodeOk() (*int32, bool)`

GetRedirectHttpCodeOk returns a tuple with the RedirectHttpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectHttpCode

`func (o *V1RedirectResponse) SetRedirectHttpCode(v int32)`

SetRedirectHttpCode sets RedirectHttpCode field to given value.

### HasRedirectHttpCode

`func (o *V1RedirectResponse) HasRedirectHttpCode() bool`

HasRedirectHttpCode returns a boolean if a field has been set.

### GetHighestRevisionNumber

`func (o *V1RedirectResponse) GetHighestRevisionNumber() string`

GetHighestRevisionNumber returns the HighestRevisionNumber field if non-nil, zero value otherwise.

### GetHighestRevisionNumberOk

`func (o *V1RedirectResponse) GetHighestRevisionNumberOk() (*string, bool)`

GetHighestRevisionNumberOk returns a tuple with the HighestRevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestRevisionNumber

`func (o *V1RedirectResponse) SetHighestRevisionNumber(v string)`

SetHighestRevisionNumber sets HighestRevisionNumber field to given value.

### HasHighestRevisionNumber

`func (o *V1RedirectResponse) HasHighestRevisionNumber() bool`

HasHighestRevisionNumber returns a boolean if a field has been set.

### GetUrl

`func (o *V1RedirectResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V1RedirectResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V1RedirectResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *V1RedirectResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


