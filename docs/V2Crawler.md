# V2Crawler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Crawler ID | 
**Name** | Pointer to **string** | Crawler name | [optional] 
**ProjectId** | **int32** | Project ID | 
**Uuid** | **string** | Crawler UUID | 
**Config** | **string** | Crawler configuration (YAML) | 
**Domain** | **string** | Crawler domain | 
**DomainVerified** | Pointer to **int32** | Domain verification status | [optional] [default to 0]
**UrlsList** | Pointer to **string** | URLs list (YAML) | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation timestamp | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Last update timestamp | [optional] 
**DeletedAt** | Pointer to **time.Time** | Deletion timestamp | [optional] 

## Methods

### NewV2Crawler

`func NewV2Crawler(id int32, projectId int32, uuid string, config string, domain string, ) *V2Crawler`

NewV2Crawler instantiates a new V2Crawler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2CrawlerWithDefaults

`func NewV2CrawlerWithDefaults() *V2Crawler`

NewV2CrawlerWithDefaults instantiates a new V2Crawler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2Crawler) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2Crawler) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2Crawler) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *V2Crawler) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2Crawler) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2Crawler) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2Crawler) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectId

`func (o *V2Crawler) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *V2Crawler) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *V2Crawler) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetUuid

`func (o *V2Crawler) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *V2Crawler) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *V2Crawler) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetConfig

`func (o *V2Crawler) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *V2Crawler) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *V2Crawler) SetConfig(v string)`

SetConfig sets Config field to given value.


### GetDomain

`func (o *V2Crawler) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V2Crawler) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V2Crawler) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetDomainVerified

`func (o *V2Crawler) GetDomainVerified() int32`

GetDomainVerified returns the DomainVerified field if non-nil, zero value otherwise.

### GetDomainVerifiedOk

`func (o *V2Crawler) GetDomainVerifiedOk() (*int32, bool)`

GetDomainVerifiedOk returns a tuple with the DomainVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainVerified

`func (o *V2Crawler) SetDomainVerified(v int32)`

SetDomainVerified sets DomainVerified field to given value.

### HasDomainVerified

`func (o *V2Crawler) HasDomainVerified() bool`

HasDomainVerified returns a boolean if a field has been set.

### GetUrlsList

`func (o *V2Crawler) GetUrlsList() string`

GetUrlsList returns the UrlsList field if non-nil, zero value otherwise.

### GetUrlsListOk

`func (o *V2Crawler) GetUrlsListOk() (*string, bool)`

GetUrlsListOk returns a tuple with the UrlsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlsList

`func (o *V2Crawler) SetUrlsList(v string)`

SetUrlsList sets UrlsList field to given value.

### HasUrlsList

`func (o *V2Crawler) HasUrlsList() bool`

HasUrlsList returns a boolean if a field has been set.

### GetCreatedAt

`func (o *V2Crawler) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2Crawler) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2Crawler) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V2Crawler) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *V2Crawler) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V2Crawler) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V2Crawler) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V2Crawler) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *V2Crawler) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *V2Crawler) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *V2Crawler) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *V2Crawler) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


