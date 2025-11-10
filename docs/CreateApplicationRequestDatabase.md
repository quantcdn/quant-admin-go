# CreateApplicationRequestDatabase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engine** | Pointer to **string** | Database engine type (MySQL 8.4, Postgres) | [optional] [default to "mysql"]
**InstanceClass** | Pointer to **string** | RDS instance class | [optional] [default to "db.t4g.micro"]
**StorageGb** | Pointer to **int32** | Allocated storage in GiB | [optional] [default to 20]
**MultiAz** | Pointer to **bool** | Enable Multi-AZ deployment (higher availability and cost) | [optional] [default to false]

## Methods

### NewCreateApplicationRequestDatabase

`func NewCreateApplicationRequestDatabase() *CreateApplicationRequestDatabase`

NewCreateApplicationRequestDatabase instantiates a new CreateApplicationRequestDatabase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApplicationRequestDatabaseWithDefaults

`func NewCreateApplicationRequestDatabaseWithDefaults() *CreateApplicationRequestDatabase`

NewCreateApplicationRequestDatabaseWithDefaults instantiates a new CreateApplicationRequestDatabase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngine

`func (o *CreateApplicationRequestDatabase) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *CreateApplicationRequestDatabase) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *CreateApplicationRequestDatabase) SetEngine(v string)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *CreateApplicationRequestDatabase) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetInstanceClass

`func (o *CreateApplicationRequestDatabase) GetInstanceClass() string`

GetInstanceClass returns the InstanceClass field if non-nil, zero value otherwise.

### GetInstanceClassOk

`func (o *CreateApplicationRequestDatabase) GetInstanceClassOk() (*string, bool)`

GetInstanceClassOk returns a tuple with the InstanceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceClass

`func (o *CreateApplicationRequestDatabase) SetInstanceClass(v string)`

SetInstanceClass sets InstanceClass field to given value.

### HasInstanceClass

`func (o *CreateApplicationRequestDatabase) HasInstanceClass() bool`

HasInstanceClass returns a boolean if a field has been set.

### GetStorageGb

`func (o *CreateApplicationRequestDatabase) GetStorageGb() int32`

GetStorageGb returns the StorageGb field if non-nil, zero value otherwise.

### GetStorageGbOk

`func (o *CreateApplicationRequestDatabase) GetStorageGbOk() (*int32, bool)`

GetStorageGbOk returns a tuple with the StorageGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGb

`func (o *CreateApplicationRequestDatabase) SetStorageGb(v int32)`

SetStorageGb sets StorageGb field to given value.

### HasStorageGb

`func (o *CreateApplicationRequestDatabase) HasStorageGb() bool`

HasStorageGb returns a boolean if a field has been set.

### GetMultiAz

`func (o *CreateApplicationRequestDatabase) GetMultiAz() bool`

GetMultiAz returns the MultiAz field if non-nil, zero value otherwise.

### GetMultiAzOk

`func (o *CreateApplicationRequestDatabase) GetMultiAzOk() (*bool, bool)`

GetMultiAzOk returns a tuple with the MultiAz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAz

`func (o *CreateApplicationRequestDatabase) SetMultiAz(v bool)`

SetMultiAz sets MultiAz field to given value.

### HasMultiAz

`func (o *CreateApplicationRequestDatabase) HasMultiAz() bool`

HasMultiAz returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


