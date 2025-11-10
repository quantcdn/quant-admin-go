# V1ContentItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**DateTimestamp** | **string** |  | 
**Published** | **string** |  | 
**RevisionCount** | **int32** |  | 
**Desc** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewV1ContentItem

`func NewV1ContentItem(url string, dateTimestamp string, published string, revisionCount int32, ) *V1ContentItem`

NewV1ContentItem instantiates a new V1ContentItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ContentItemWithDefaults

`func NewV1ContentItemWithDefaults() *V1ContentItem`

NewV1ContentItemWithDefaults instantiates a new V1ContentItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *V1ContentItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V1ContentItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V1ContentItem) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDateTimestamp

`func (o *V1ContentItem) GetDateTimestamp() string`

GetDateTimestamp returns the DateTimestamp field if non-nil, zero value otherwise.

### GetDateTimestampOk

`func (o *V1ContentItem) GetDateTimestampOk() (*string, bool)`

GetDateTimestampOk returns a tuple with the DateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimestamp

`func (o *V1ContentItem) SetDateTimestamp(v string)`

SetDateTimestamp sets DateTimestamp field to given value.


### GetPublished

`func (o *V1ContentItem) GetPublished() string`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *V1ContentItem) GetPublishedOk() (*string, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *V1ContentItem) SetPublished(v string)`

SetPublished sets Published field to given value.


### GetRevisionCount

`func (o *V1ContentItem) GetRevisionCount() int32`

GetRevisionCount returns the RevisionCount field if non-nil, zero value otherwise.

### GetRevisionCountOk

`func (o *V1ContentItem) GetRevisionCountOk() (*int32, bool)`

GetRevisionCountOk returns a tuple with the RevisionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionCount

`func (o *V1ContentItem) SetRevisionCount(v int32)`

SetRevisionCount sets RevisionCount field to given value.


### GetDesc

`func (o *V1ContentItem) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *V1ContentItem) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *V1ContentItem) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *V1ContentItem) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetUuid

`func (o *V1ContentItem) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *V1ContentItem) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *V1ContentItem) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *V1ContentItem) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


