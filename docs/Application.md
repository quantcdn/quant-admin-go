# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | **string** | Application name | 
**Organisation** | **string** | Organisation machine name | 
**Database** | Pointer to [**NullableApplicationDatabase**](ApplicationDatabase.md) |  | [optional] 
**Filesystem** | Pointer to [**NullableApplicationFilesystem**](ApplicationFilesystem.md) |  | [optional] 
**ComposeDefinition** | Pointer to [**Compose**](Compose.md) |  | [optional] 
**Status** | Pointer to **NullableString** | Application status | [optional] 
**DeploymentInformation** | Pointer to [**[]ApplicationDeploymentInformationInner**](ApplicationDeploymentInformationInner.md) | Deployment history | [optional] 
**ImageReference** | Pointer to [**NullableApplicationImageReference**](ApplicationImageReference.md) |  | [optional] 
**ContainerNames** | Pointer to **[]string** | List of container names | [optional] 
**MinCapacity** | Pointer to **NullableInt32** | Minimum task count for auto-scaling | [optional] 
**MaxCapacity** | Pointer to **NullableInt32** | Maximum task count for auto-scaling | [optional] 
**DesiredCount** | Pointer to **NullableInt32** | Desired task count | [optional] 
**RunningCount** | Pointer to **NullableInt32** | Currently running task count | [optional] 
**Environments** | Pointer to [**[]ApplicationEnvironmentsInner**](ApplicationEnvironmentsInner.md) | List of environments | [optional] 

## Methods

### NewApplication

`func NewApplication(appName string, organisation string, ) *Application`

NewApplication instantiates a new Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithDefaults

`func NewApplicationWithDefaults() *Application`

NewApplicationWithDefaults instantiates a new Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *Application) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *Application) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *Application) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetOrganisation

`func (o *Application) GetOrganisation() string`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *Application) GetOrganisationOk() (*string, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *Application) SetOrganisation(v string)`

SetOrganisation sets Organisation field to given value.


### GetDatabase

`func (o *Application) GetDatabase() ApplicationDatabase`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *Application) GetDatabaseOk() (*ApplicationDatabase, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *Application) SetDatabase(v ApplicationDatabase)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *Application) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### SetDatabaseNil

`func (o *Application) SetDatabaseNil(b bool)`

 SetDatabaseNil sets the value for Database to be an explicit nil

### UnsetDatabase
`func (o *Application) UnsetDatabase()`

UnsetDatabase ensures that no value is present for Database, not even an explicit nil
### GetFilesystem

`func (o *Application) GetFilesystem() ApplicationFilesystem`

GetFilesystem returns the Filesystem field if non-nil, zero value otherwise.

### GetFilesystemOk

`func (o *Application) GetFilesystemOk() (*ApplicationFilesystem, bool)`

GetFilesystemOk returns a tuple with the Filesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystem

`func (o *Application) SetFilesystem(v ApplicationFilesystem)`

SetFilesystem sets Filesystem field to given value.

### HasFilesystem

`func (o *Application) HasFilesystem() bool`

HasFilesystem returns a boolean if a field has been set.

### SetFilesystemNil

`func (o *Application) SetFilesystemNil(b bool)`

 SetFilesystemNil sets the value for Filesystem to be an explicit nil

### UnsetFilesystem
`func (o *Application) UnsetFilesystem()`

UnsetFilesystem ensures that no value is present for Filesystem, not even an explicit nil
### GetComposeDefinition

`func (o *Application) GetComposeDefinition() Compose`

GetComposeDefinition returns the ComposeDefinition field if non-nil, zero value otherwise.

### GetComposeDefinitionOk

`func (o *Application) GetComposeDefinitionOk() (*Compose, bool)`

GetComposeDefinitionOk returns a tuple with the ComposeDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposeDefinition

`func (o *Application) SetComposeDefinition(v Compose)`

SetComposeDefinition sets ComposeDefinition field to given value.

### HasComposeDefinition

`func (o *Application) HasComposeDefinition() bool`

HasComposeDefinition returns a boolean if a field has been set.

### GetStatus

`func (o *Application) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Application) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Application) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Application) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Application) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Application) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDeploymentInformation

`func (o *Application) GetDeploymentInformation() []ApplicationDeploymentInformationInner`

GetDeploymentInformation returns the DeploymentInformation field if non-nil, zero value otherwise.

### GetDeploymentInformationOk

`func (o *Application) GetDeploymentInformationOk() (*[]ApplicationDeploymentInformationInner, bool)`

GetDeploymentInformationOk returns a tuple with the DeploymentInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentInformation

`func (o *Application) SetDeploymentInformation(v []ApplicationDeploymentInformationInner)`

SetDeploymentInformation sets DeploymentInformation field to given value.

### HasDeploymentInformation

`func (o *Application) HasDeploymentInformation() bool`

HasDeploymentInformation returns a boolean if a field has been set.

### SetDeploymentInformationNil

`func (o *Application) SetDeploymentInformationNil(b bool)`

 SetDeploymentInformationNil sets the value for DeploymentInformation to be an explicit nil

### UnsetDeploymentInformation
`func (o *Application) UnsetDeploymentInformation()`

UnsetDeploymentInformation ensures that no value is present for DeploymentInformation, not even an explicit nil
### GetImageReference

`func (o *Application) GetImageReference() ApplicationImageReference`

GetImageReference returns the ImageReference field if non-nil, zero value otherwise.

### GetImageReferenceOk

`func (o *Application) GetImageReferenceOk() (*ApplicationImageReference, bool)`

GetImageReferenceOk returns a tuple with the ImageReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageReference

`func (o *Application) SetImageReference(v ApplicationImageReference)`

SetImageReference sets ImageReference field to given value.

### HasImageReference

`func (o *Application) HasImageReference() bool`

HasImageReference returns a boolean if a field has been set.

### SetImageReferenceNil

`func (o *Application) SetImageReferenceNil(b bool)`

 SetImageReferenceNil sets the value for ImageReference to be an explicit nil

### UnsetImageReference
`func (o *Application) UnsetImageReference()`

UnsetImageReference ensures that no value is present for ImageReference, not even an explicit nil
### GetContainerNames

`func (o *Application) GetContainerNames() []string`

GetContainerNames returns the ContainerNames field if non-nil, zero value otherwise.

### GetContainerNamesOk

`func (o *Application) GetContainerNamesOk() (*[]string, bool)`

GetContainerNamesOk returns a tuple with the ContainerNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerNames

`func (o *Application) SetContainerNames(v []string)`

SetContainerNames sets ContainerNames field to given value.

### HasContainerNames

`func (o *Application) HasContainerNames() bool`

HasContainerNames returns a boolean if a field has been set.

### SetContainerNamesNil

`func (o *Application) SetContainerNamesNil(b bool)`

 SetContainerNamesNil sets the value for ContainerNames to be an explicit nil

### UnsetContainerNames
`func (o *Application) UnsetContainerNames()`

UnsetContainerNames ensures that no value is present for ContainerNames, not even an explicit nil
### GetMinCapacity

`func (o *Application) GetMinCapacity() int32`

GetMinCapacity returns the MinCapacity field if non-nil, zero value otherwise.

### GetMinCapacityOk

`func (o *Application) GetMinCapacityOk() (*int32, bool)`

GetMinCapacityOk returns a tuple with the MinCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCapacity

`func (o *Application) SetMinCapacity(v int32)`

SetMinCapacity sets MinCapacity field to given value.

### HasMinCapacity

`func (o *Application) HasMinCapacity() bool`

HasMinCapacity returns a boolean if a field has been set.

### SetMinCapacityNil

`func (o *Application) SetMinCapacityNil(b bool)`

 SetMinCapacityNil sets the value for MinCapacity to be an explicit nil

### UnsetMinCapacity
`func (o *Application) UnsetMinCapacity()`

UnsetMinCapacity ensures that no value is present for MinCapacity, not even an explicit nil
### GetMaxCapacity

`func (o *Application) GetMaxCapacity() int32`

GetMaxCapacity returns the MaxCapacity field if non-nil, zero value otherwise.

### GetMaxCapacityOk

`func (o *Application) GetMaxCapacityOk() (*int32, bool)`

GetMaxCapacityOk returns a tuple with the MaxCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCapacity

`func (o *Application) SetMaxCapacity(v int32)`

SetMaxCapacity sets MaxCapacity field to given value.

### HasMaxCapacity

`func (o *Application) HasMaxCapacity() bool`

HasMaxCapacity returns a boolean if a field has been set.

### SetMaxCapacityNil

`func (o *Application) SetMaxCapacityNil(b bool)`

 SetMaxCapacityNil sets the value for MaxCapacity to be an explicit nil

### UnsetMaxCapacity
`func (o *Application) UnsetMaxCapacity()`

UnsetMaxCapacity ensures that no value is present for MaxCapacity, not even an explicit nil
### GetDesiredCount

`func (o *Application) GetDesiredCount() int32`

GetDesiredCount returns the DesiredCount field if non-nil, zero value otherwise.

### GetDesiredCountOk

`func (o *Application) GetDesiredCountOk() (*int32, bool)`

GetDesiredCountOk returns a tuple with the DesiredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredCount

`func (o *Application) SetDesiredCount(v int32)`

SetDesiredCount sets DesiredCount field to given value.

### HasDesiredCount

`func (o *Application) HasDesiredCount() bool`

HasDesiredCount returns a boolean if a field has been set.

### SetDesiredCountNil

`func (o *Application) SetDesiredCountNil(b bool)`

 SetDesiredCountNil sets the value for DesiredCount to be an explicit nil

### UnsetDesiredCount
`func (o *Application) UnsetDesiredCount()`

UnsetDesiredCount ensures that no value is present for DesiredCount, not even an explicit nil
### GetRunningCount

`func (o *Application) GetRunningCount() int32`

GetRunningCount returns the RunningCount field if non-nil, zero value otherwise.

### GetRunningCountOk

`func (o *Application) GetRunningCountOk() (*int32, bool)`

GetRunningCountOk returns a tuple with the RunningCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningCount

`func (o *Application) SetRunningCount(v int32)`

SetRunningCount sets RunningCount field to given value.

### HasRunningCount

`func (o *Application) HasRunningCount() bool`

HasRunningCount returns a boolean if a field has been set.

### SetRunningCountNil

`func (o *Application) SetRunningCountNil(b bool)`

 SetRunningCountNil sets the value for RunningCount to be an explicit nil

### UnsetRunningCount
`func (o *Application) UnsetRunningCount()`

UnsetRunningCount ensures that no value is present for RunningCount, not even an explicit nil
### GetEnvironments

`func (o *Application) GetEnvironments() []ApplicationEnvironmentsInner`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *Application) GetEnvironmentsOk() (*[]ApplicationEnvironmentsInner, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *Application) SetEnvironments(v []ApplicationEnvironmentsInner)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *Application) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### SetEnvironmentsNil

`func (o *Application) SetEnvironmentsNil(b bool)`

 SetEnvironmentsNil sets the value for Environments to be an explicit nil

### UnsetEnvironments
`func (o *Application) UnsetEnvironments()`

UnsetEnvironments ensures that no value is present for Environments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


