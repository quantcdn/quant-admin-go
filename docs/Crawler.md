# Crawler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **string** |  | [optional] 
**UrlsList** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**DomainVerified** | Pointer to **int32** |  | [optional] 

## Methods

### NewCrawler

`func NewCrawler() *Crawler`

NewCrawler instantiates a new Crawler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrawlerWithDefaults

`func NewCrawlerWithDefaults() *Crawler`

NewCrawlerWithDefaults instantiates a new Crawler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Crawler) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Crawler) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Crawler) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Crawler) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectId

`func (o *Crawler) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Crawler) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Crawler) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Crawler) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetUuid

`func (o *Crawler) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Crawler) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Crawler) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Crawler) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetConfig

`func (o *Crawler) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Crawler) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Crawler) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Crawler) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetUrlsList

`func (o *Crawler) GetUrlsList() []string`

GetUrlsList returns the UrlsList field if non-nil, zero value otherwise.

### GetUrlsListOk

`func (o *Crawler) GetUrlsListOk() (*[]string, bool)`

GetUrlsListOk returns a tuple with the UrlsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlsList

`func (o *Crawler) SetUrlsList(v []string)`

SetUrlsList sets UrlsList field to given value.

### HasUrlsList

`func (o *Crawler) HasUrlsList() bool`

HasUrlsList returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Crawler) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Crawler) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Crawler) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Crawler) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Crawler) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Crawler) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Crawler) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Crawler) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Crawler) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Crawler) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Crawler) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Crawler) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDomain

`func (o *Crawler) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Crawler) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Crawler) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Crawler) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDomainVerified

`func (o *Crawler) GetDomainVerified() int32`

GetDomainVerified returns the DomainVerified field if non-nil, zero value otherwise.

### GetDomainVerifiedOk

`func (o *Crawler) GetDomainVerifiedOk() (*int32, bool)`

GetDomainVerifiedOk returns a tuple with the DomainVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainVerified

`func (o *Crawler) SetDomainVerified(v int32)`

SetDomainVerified sets DomainVerified field to given value.

### HasDomainVerified

`func (o *Crawler) HasDomainVerified() bool`

HasDomainVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


