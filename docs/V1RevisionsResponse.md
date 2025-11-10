# V1RevisionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revisions** | Pointer to **map[string]interface{}** | Revision objects, mapped by revision number | [optional] 
**Url** | Pointer to **string** | The url of the asset | [optional] 
**Published** | Pointer to **bool** | Published state of the asset | [optional] 
**PublishedRevision** | Pointer to **int32** | Published revision number of the asset | [optional] 
**Transitions** | Pointer to [**[]V1Transition**](V1Transition.md) |  | [optional] 
**HighestRevisionNumber** | Pointer to **int32** | Last revision number | [optional] 
**TransitionRevision** | Pointer to **int32** | The transition number, if set | [optional] 

## Methods

### NewV1RevisionsResponse

`func NewV1RevisionsResponse() *V1RevisionsResponse`

NewV1RevisionsResponse instantiates a new V1RevisionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RevisionsResponseWithDefaults

`func NewV1RevisionsResponseWithDefaults() *V1RevisionsResponse`

NewV1RevisionsResponseWithDefaults instantiates a new V1RevisionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisions

`func (o *V1RevisionsResponse) GetRevisions() map[string]interface{}`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *V1RevisionsResponse) GetRevisionsOk() (*map[string]interface{}, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *V1RevisionsResponse) SetRevisions(v map[string]interface{})`

SetRevisions sets Revisions field to given value.

### HasRevisions

`func (o *V1RevisionsResponse) HasRevisions() bool`

HasRevisions returns a boolean if a field has been set.

### GetUrl

`func (o *V1RevisionsResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V1RevisionsResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V1RevisionsResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *V1RevisionsResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetPublished

`func (o *V1RevisionsResponse) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *V1RevisionsResponse) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *V1RevisionsResponse) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *V1RevisionsResponse) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetPublishedRevision

`func (o *V1RevisionsResponse) GetPublishedRevision() int32`

GetPublishedRevision returns the PublishedRevision field if non-nil, zero value otherwise.

### GetPublishedRevisionOk

`func (o *V1RevisionsResponse) GetPublishedRevisionOk() (*int32, bool)`

GetPublishedRevisionOk returns a tuple with the PublishedRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedRevision

`func (o *V1RevisionsResponse) SetPublishedRevision(v int32)`

SetPublishedRevision sets PublishedRevision field to given value.

### HasPublishedRevision

`func (o *V1RevisionsResponse) HasPublishedRevision() bool`

HasPublishedRevision returns a boolean if a field has been set.

### GetTransitions

`func (o *V1RevisionsResponse) GetTransitions() []V1Transition`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *V1RevisionsResponse) GetTransitionsOk() (*[]V1Transition, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *V1RevisionsResponse) SetTransitions(v []V1Transition)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *V1RevisionsResponse) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.

### GetHighestRevisionNumber

`func (o *V1RevisionsResponse) GetHighestRevisionNumber() int32`

GetHighestRevisionNumber returns the HighestRevisionNumber field if non-nil, zero value otherwise.

### GetHighestRevisionNumberOk

`func (o *V1RevisionsResponse) GetHighestRevisionNumberOk() (*int32, bool)`

GetHighestRevisionNumberOk returns a tuple with the HighestRevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestRevisionNumber

`func (o *V1RevisionsResponse) SetHighestRevisionNumber(v int32)`

SetHighestRevisionNumber sets HighestRevisionNumber field to given value.

### HasHighestRevisionNumber

`func (o *V1RevisionsResponse) HasHighestRevisionNumber() bool`

HasHighestRevisionNumber returns a boolean if a field has been set.

### GetTransitionRevision

`func (o *V1RevisionsResponse) GetTransitionRevision() int32`

GetTransitionRevision returns the TransitionRevision field if non-nil, zero value otherwise.

### GetTransitionRevisionOk

`func (o *V1RevisionsResponse) GetTransitionRevisionOk() (*int32, bool)`

GetTransitionRevisionOk returns a tuple with the TransitionRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionRevision

`func (o *V1RevisionsResponse) SetTransitionRevision(v int32)`

SetTransitionRevision sets TransitionRevision field to given value.

### HasTransitionRevision

`func (o *V1RevisionsResponse) HasTransitionRevision() bool`

HasTransitionRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


