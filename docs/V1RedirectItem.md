# V1RedirectItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | * @OA\\Schema(   schema&#x3D;\&quot;V1RedirectItem\&quot;,   type&#x3D;\&quot;object\&quot;,   required&#x3D;{\&quot;url\&quot;,\&quot;redirect_url\&quot;,\&quot;redirect_http_code\&quot;,\&quot;date_timestamp\&quot;,\&quot;published\&quot;,\&quot;revision_count\&quot;}, | 
**RedirectUrl** | **string** |  | 
**RedirectHttpCode** | **int32** |  | 
**DateTimestamp** | **string** |  | 
**Published** | **string** |  | 
**RevisionCount** | **int32** |  | 

## Methods

### NewV1RedirectItem

`func NewV1RedirectItem(url string, redirectUrl string, redirectHttpCode int32, dateTimestamp string, published string, revisionCount int32, ) *V1RedirectItem`

NewV1RedirectItem instantiates a new V1RedirectItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RedirectItemWithDefaults

`func NewV1RedirectItemWithDefaults() *V1RedirectItem`

NewV1RedirectItemWithDefaults instantiates a new V1RedirectItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *V1RedirectItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V1RedirectItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V1RedirectItem) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRedirectUrl

`func (o *V1RedirectItem) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *V1RedirectItem) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *V1RedirectItem) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.


### GetRedirectHttpCode

`func (o *V1RedirectItem) GetRedirectHttpCode() int32`

GetRedirectHttpCode returns the RedirectHttpCode field if non-nil, zero value otherwise.

### GetRedirectHttpCodeOk

`func (o *V1RedirectItem) GetRedirectHttpCodeOk() (*int32, bool)`

GetRedirectHttpCodeOk returns a tuple with the RedirectHttpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectHttpCode

`func (o *V1RedirectItem) SetRedirectHttpCode(v int32)`

SetRedirectHttpCode sets RedirectHttpCode field to given value.


### GetDateTimestamp

`func (o *V1RedirectItem) GetDateTimestamp() string`

GetDateTimestamp returns the DateTimestamp field if non-nil, zero value otherwise.

### GetDateTimestampOk

`func (o *V1RedirectItem) GetDateTimestampOk() (*string, bool)`

GetDateTimestampOk returns a tuple with the DateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimestamp

`func (o *V1RedirectItem) SetDateTimestamp(v string)`

SetDateTimestamp sets DateTimestamp field to given value.


### GetPublished

`func (o *V1RedirectItem) GetPublished() string`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *V1RedirectItem) GetPublishedOk() (*string, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *V1RedirectItem) SetPublished(v string)`

SetPublished sets Published field to given value.


### GetRevisionCount

`func (o *V1RedirectItem) GetRevisionCount() int32`

GetRevisionCount returns the RevisionCount field if non-nil, zero value otherwise.

### GetRevisionCountOk

`func (o *V1RedirectItem) GetRevisionCountOk() (*int32, bool)`

GetRevisionCountOk returns a tuple with the RevisionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionCount

`func (o *V1RedirectItem) SetRevisionCount(v int32)`

SetRevisionCount sets RevisionCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


