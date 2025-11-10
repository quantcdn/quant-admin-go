# Cron

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 
**Command** | Pointer to **string** |  | [optional] 

## Methods

### NewCron

`func NewCron() *Cron`

NewCron instantiates a new Cron object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronWithDefaults

`func NewCronWithDefaults() *Cron`

NewCronWithDefaults instantiates a new Cron object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Cron) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cron) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cron) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Cron) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchedule

`func (o *Cron) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Cron) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Cron) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *Cron) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetCommand

`func (o *Cron) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *Cron) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *Cron) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *Cron) HasCommand() bool`

HasCommand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


