# ApplicationDeploymentInformationInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | Pointer to **string** | Deployment identifier | [optional] 
**TaskDefinitionArn** | Pointer to **string** | Task definition ARN used | [optional] 
**CreatedAt** | Pointer to **time.Time** | Deployment creation timestamp | [optional] 
**Status** | Pointer to **string** | Deployment status | [optional] 
**ImageTag** | Pointer to **string** | Image tag deployed | [optional] 

## Methods

### NewApplicationDeploymentInformationInner

`func NewApplicationDeploymentInformationInner() *ApplicationDeploymentInformationInner`

NewApplicationDeploymentInformationInner instantiates a new ApplicationDeploymentInformationInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationDeploymentInformationInnerWithDefaults

`func NewApplicationDeploymentInformationInnerWithDefaults() *ApplicationDeploymentInformationInner`

NewApplicationDeploymentInformationInnerWithDefaults instantiates a new ApplicationDeploymentInformationInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *ApplicationDeploymentInformationInner) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *ApplicationDeploymentInformationInner) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *ApplicationDeploymentInformationInner) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *ApplicationDeploymentInformationInner) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetTaskDefinitionArn

`func (o *ApplicationDeploymentInformationInner) GetTaskDefinitionArn() string`

GetTaskDefinitionArn returns the TaskDefinitionArn field if non-nil, zero value otherwise.

### GetTaskDefinitionArnOk

`func (o *ApplicationDeploymentInformationInner) GetTaskDefinitionArnOk() (*string, bool)`

GetTaskDefinitionArnOk returns a tuple with the TaskDefinitionArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinitionArn

`func (o *ApplicationDeploymentInformationInner) SetTaskDefinitionArn(v string)`

SetTaskDefinitionArn sets TaskDefinitionArn field to given value.

### HasTaskDefinitionArn

`func (o *ApplicationDeploymentInformationInner) HasTaskDefinitionArn() bool`

HasTaskDefinitionArn returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApplicationDeploymentInformationInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationDeploymentInformationInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationDeploymentInformationInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApplicationDeploymentInformationInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *ApplicationDeploymentInformationInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplicationDeploymentInformationInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplicationDeploymentInformationInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplicationDeploymentInformationInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetImageTag

`func (o *ApplicationDeploymentInformationInner) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *ApplicationDeploymentInformationInner) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *ApplicationDeploymentInformationInner) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.

### HasImageTag

`func (o *ApplicationDeploymentInformationInner) HasImageTag() bool`

HasImageTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


