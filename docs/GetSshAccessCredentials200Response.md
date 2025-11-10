# GetSshAccessCredentials200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Credentials** | Pointer to [**GetSshAccessCredentials200ResponseCredentials**](GetSshAccessCredentials200ResponseCredentials.md) |  | [optional] 
**ClusterName** | Pointer to **string** |  | [optional] 
**TaskArn** | Pointer to **string** |  | [optional] 
**TaskId** | Pointer to **string** |  | [optional] 
**ContainerNames** | Pointer to **[]string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**ExpiresIn** | Pointer to **int32** |  | [optional] 
**OrganizationScope** | Pointer to **string** |  | [optional] 

## Methods

### NewGetSshAccessCredentials200Response

`func NewGetSshAccessCredentials200Response() *GetSshAccessCredentials200Response`

NewGetSshAccessCredentials200Response instantiates a new GetSshAccessCredentials200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSshAccessCredentials200ResponseWithDefaults

`func NewGetSshAccessCredentials200ResponseWithDefaults() *GetSshAccessCredentials200Response`

NewGetSshAccessCredentials200ResponseWithDefaults instantiates a new GetSshAccessCredentials200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetSshAccessCredentials200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetSshAccessCredentials200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetSshAccessCredentials200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetSshAccessCredentials200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetCredentials

`func (o *GetSshAccessCredentials200Response) GetCredentials() GetSshAccessCredentials200ResponseCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GetSshAccessCredentials200Response) GetCredentialsOk() (*GetSshAccessCredentials200ResponseCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GetSshAccessCredentials200Response) SetCredentials(v GetSshAccessCredentials200ResponseCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *GetSshAccessCredentials200Response) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetClusterName

`func (o *GetSshAccessCredentials200Response) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *GetSshAccessCredentials200Response) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *GetSshAccessCredentials200Response) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *GetSshAccessCredentials200Response) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetTaskArn

`func (o *GetSshAccessCredentials200Response) GetTaskArn() string`

GetTaskArn returns the TaskArn field if non-nil, zero value otherwise.

### GetTaskArnOk

`func (o *GetSshAccessCredentials200Response) GetTaskArnOk() (*string, bool)`

GetTaskArnOk returns a tuple with the TaskArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskArn

`func (o *GetSshAccessCredentials200Response) SetTaskArn(v string)`

SetTaskArn sets TaskArn field to given value.

### HasTaskArn

`func (o *GetSshAccessCredentials200Response) HasTaskArn() bool`

HasTaskArn returns a boolean if a field has been set.

### GetTaskId

`func (o *GetSshAccessCredentials200Response) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *GetSshAccessCredentials200Response) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *GetSshAccessCredentials200Response) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *GetSshAccessCredentials200Response) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetContainerNames

`func (o *GetSshAccessCredentials200Response) GetContainerNames() []string`

GetContainerNames returns the ContainerNames field if non-nil, zero value otherwise.

### GetContainerNamesOk

`func (o *GetSshAccessCredentials200Response) GetContainerNamesOk() (*[]string, bool)`

GetContainerNamesOk returns a tuple with the ContainerNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerNames

`func (o *GetSshAccessCredentials200Response) SetContainerNames(v []string)`

SetContainerNames sets ContainerNames field to given value.

### HasContainerNames

`func (o *GetSshAccessCredentials200Response) HasContainerNames() bool`

HasContainerNames returns a boolean if a field has been set.

### GetRegion

`func (o *GetSshAccessCredentials200Response) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetSshAccessCredentials200Response) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetSshAccessCredentials200Response) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetSshAccessCredentials200Response) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetExpiresIn

`func (o *GetSshAccessCredentials200Response) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *GetSshAccessCredentials200Response) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *GetSshAccessCredentials200Response) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *GetSshAccessCredentials200Response) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetOrganizationScope

`func (o *GetSshAccessCredentials200Response) GetOrganizationScope() string`

GetOrganizationScope returns the OrganizationScope field if non-nil, zero value otherwise.

### GetOrganizationScopeOk

`func (o *GetSshAccessCredentials200Response) GetOrganizationScopeOk() (*string, bool)`

GetOrganizationScopeOk returns a tuple with the OrganizationScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationScope

`func (o *GetSshAccessCredentials200Response) SetOrganizationScope(v string)`

SetOrganizationScope sets OrganizationScope field to given value.

### HasOrganizationScope

`func (o *GetSshAccessCredentials200Response) HasOrganizationScope() bool`

HasOrganizationScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


