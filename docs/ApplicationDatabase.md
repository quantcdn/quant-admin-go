# ApplicationDatabase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RdsInstanceIdentifier** | Pointer to **string** | RDS instance identifier | [optional] 
**RdsInstanceEndpoint** | Pointer to **string** | RDS instance endpoint address | [optional] 
**RdsInstanceEngine** | Pointer to **string** | Database engine | [optional] 
**RdsInstanceStatus** | Pointer to **string** | RDS instance status | [optional] 

## Methods

### NewApplicationDatabase

`func NewApplicationDatabase() *ApplicationDatabase`

NewApplicationDatabase instantiates a new ApplicationDatabase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationDatabaseWithDefaults

`func NewApplicationDatabaseWithDefaults() *ApplicationDatabase`

NewApplicationDatabaseWithDefaults instantiates a new ApplicationDatabase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRdsInstanceIdentifier

`func (o *ApplicationDatabase) GetRdsInstanceIdentifier() string`

GetRdsInstanceIdentifier returns the RdsInstanceIdentifier field if non-nil, zero value otherwise.

### GetRdsInstanceIdentifierOk

`func (o *ApplicationDatabase) GetRdsInstanceIdentifierOk() (*string, bool)`

GetRdsInstanceIdentifierOk returns a tuple with the RdsInstanceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsInstanceIdentifier

`func (o *ApplicationDatabase) SetRdsInstanceIdentifier(v string)`

SetRdsInstanceIdentifier sets RdsInstanceIdentifier field to given value.

### HasRdsInstanceIdentifier

`func (o *ApplicationDatabase) HasRdsInstanceIdentifier() bool`

HasRdsInstanceIdentifier returns a boolean if a field has been set.

### GetRdsInstanceEndpoint

`func (o *ApplicationDatabase) GetRdsInstanceEndpoint() string`

GetRdsInstanceEndpoint returns the RdsInstanceEndpoint field if non-nil, zero value otherwise.

### GetRdsInstanceEndpointOk

`func (o *ApplicationDatabase) GetRdsInstanceEndpointOk() (*string, bool)`

GetRdsInstanceEndpointOk returns a tuple with the RdsInstanceEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsInstanceEndpoint

`func (o *ApplicationDatabase) SetRdsInstanceEndpoint(v string)`

SetRdsInstanceEndpoint sets RdsInstanceEndpoint field to given value.

### HasRdsInstanceEndpoint

`func (o *ApplicationDatabase) HasRdsInstanceEndpoint() bool`

HasRdsInstanceEndpoint returns a boolean if a field has been set.

### GetRdsInstanceEngine

`func (o *ApplicationDatabase) GetRdsInstanceEngine() string`

GetRdsInstanceEngine returns the RdsInstanceEngine field if non-nil, zero value otherwise.

### GetRdsInstanceEngineOk

`func (o *ApplicationDatabase) GetRdsInstanceEngineOk() (*string, bool)`

GetRdsInstanceEngineOk returns a tuple with the RdsInstanceEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsInstanceEngine

`func (o *ApplicationDatabase) SetRdsInstanceEngine(v string)`

SetRdsInstanceEngine sets RdsInstanceEngine field to given value.

### HasRdsInstanceEngine

`func (o *ApplicationDatabase) HasRdsInstanceEngine() bool`

HasRdsInstanceEngine returns a boolean if a field has been set.

### GetRdsInstanceStatus

`func (o *ApplicationDatabase) GetRdsInstanceStatus() string`

GetRdsInstanceStatus returns the RdsInstanceStatus field if non-nil, zero value otherwise.

### GetRdsInstanceStatusOk

`func (o *ApplicationDatabase) GetRdsInstanceStatusOk() (*string, bool)`

GetRdsInstanceStatusOk returns a tuple with the RdsInstanceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsInstanceStatus

`func (o *ApplicationDatabase) SetRdsInstanceStatus(v string)`

SetRdsInstanceStatus sets RdsInstanceStatus field to given value.

### HasRdsInstanceStatus

`func (o *ApplicationDatabase) HasRdsInstanceStatus() bool`

HasRdsInstanceStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


