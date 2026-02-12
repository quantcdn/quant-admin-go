# DeleteVectorDocumentsRequestMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** | Metadata field name (e.g., &#39;drupal_entity_id&#39;) | [optional] 
**Values** | Pointer to **[]string** | Values to match (OR logic) | [optional] 

## Methods

### NewDeleteVectorDocumentsRequestMetadata

`func NewDeleteVectorDocumentsRequestMetadata() *DeleteVectorDocumentsRequestMetadata`

NewDeleteVectorDocumentsRequestMetadata instantiates a new DeleteVectorDocumentsRequestMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteVectorDocumentsRequestMetadataWithDefaults

`func NewDeleteVectorDocumentsRequestMetadataWithDefaults() *DeleteVectorDocumentsRequestMetadata`

NewDeleteVectorDocumentsRequestMetadataWithDefaults instantiates a new DeleteVectorDocumentsRequestMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *DeleteVectorDocumentsRequestMetadata) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *DeleteVectorDocumentsRequestMetadata) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *DeleteVectorDocumentsRequestMetadata) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *DeleteVectorDocumentsRequestMetadata) HasField() bool`

HasField returns a boolean if a field has been set.

### GetValues

`func (o *DeleteVectorDocumentsRequestMetadata) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *DeleteVectorDocumentsRequestMetadata) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *DeleteVectorDocumentsRequestMetadata) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *DeleteVectorDocumentsRequestMetadata) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


