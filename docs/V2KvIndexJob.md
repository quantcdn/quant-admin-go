# V2KvIndexJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | settling: about 75 seconds before the backfill starts; scanning: the backfill reads existing items; cleaning: a drop removes index data; done: no work is left. | 
**Scanned** | **int32** | Items the job has examined. If phase and scanned do not change for more than 960 seconds, call KV_indexes_resume. | 

## Methods

### NewV2KvIndexJob

`func NewV2KvIndexJob(phase string, scanned int32, ) *V2KvIndexJob`

NewV2KvIndexJob instantiates a new V2KvIndexJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvIndexJobWithDefaults

`func NewV2KvIndexJobWithDefaults() *V2KvIndexJob`

NewV2KvIndexJobWithDefaults instantiates a new V2KvIndexJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *V2KvIndexJob) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *V2KvIndexJob) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *V2KvIndexJob) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetScanned

`func (o *V2KvIndexJob) GetScanned() int32`

GetScanned returns the Scanned field if non-nil, zero value otherwise.

### GetScannedOk

`func (o *V2KvIndexJob) GetScannedOk() (*int32, bool)`

GetScannedOk returns a tuple with the Scanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanned

`func (o *V2KvIndexJob) SetScanned(v int32)`

SetScanned sets Scanned field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


