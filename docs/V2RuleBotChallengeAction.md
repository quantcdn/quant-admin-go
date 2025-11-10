# V2RuleBotChallengeAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
**RobotChallengeType** | **string** | Challenge type (invisible or checkbox) | 
**RobotChallengeVerificationTtl** | Pointer to **int32** | Verification TTL in seconds | [optional] [default to 10800]
**RobotChallengeChallengeTtl** | Pointer to **int32** | Challenge TTL in seconds | [optional] [default to 30]

## Methods

### NewV2RuleBotChallengeAction

`func NewV2RuleBotChallengeAction(message string, error_ bool, robotChallengeType string, ) *V2RuleBotChallengeAction`

NewV2RuleBotChallengeAction instantiates a new V2RuleBotChallengeAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleBotChallengeActionWithDefaults

`func NewV2RuleBotChallengeActionWithDefaults() *V2RuleBotChallengeAction`

NewV2RuleBotChallengeActionWithDefaults instantiates a new V2RuleBotChallengeAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *V2RuleBotChallengeAction) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2RuleBotChallengeAction) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2RuleBotChallengeAction) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2RuleBotChallengeAction) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2RuleBotChallengeAction) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2RuleBotChallengeAction) SetError(v bool)`

SetError sets Error field to given value.


### GetRobotChallengeType

`func (o *V2RuleBotChallengeAction) GetRobotChallengeType() string`

GetRobotChallengeType returns the RobotChallengeType field if non-nil, zero value otherwise.

### GetRobotChallengeTypeOk

`func (o *V2RuleBotChallengeAction) GetRobotChallengeTypeOk() (*string, bool)`

GetRobotChallengeTypeOk returns a tuple with the RobotChallengeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRobotChallengeType

`func (o *V2RuleBotChallengeAction) SetRobotChallengeType(v string)`

SetRobotChallengeType sets RobotChallengeType field to given value.


### GetRobotChallengeVerificationTtl

`func (o *V2RuleBotChallengeAction) GetRobotChallengeVerificationTtl() int32`

GetRobotChallengeVerificationTtl returns the RobotChallengeVerificationTtl field if non-nil, zero value otherwise.

### GetRobotChallengeVerificationTtlOk

`func (o *V2RuleBotChallengeAction) GetRobotChallengeVerificationTtlOk() (*int32, bool)`

GetRobotChallengeVerificationTtlOk returns a tuple with the RobotChallengeVerificationTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRobotChallengeVerificationTtl

`func (o *V2RuleBotChallengeAction) SetRobotChallengeVerificationTtl(v int32)`

SetRobotChallengeVerificationTtl sets RobotChallengeVerificationTtl field to given value.

### HasRobotChallengeVerificationTtl

`func (o *V2RuleBotChallengeAction) HasRobotChallengeVerificationTtl() bool`

HasRobotChallengeVerificationTtl returns a boolean if a field has been set.

### GetRobotChallengeChallengeTtl

`func (o *V2RuleBotChallengeAction) GetRobotChallengeChallengeTtl() int32`

GetRobotChallengeChallengeTtl returns the RobotChallengeChallengeTtl field if non-nil, zero value otherwise.

### GetRobotChallengeChallengeTtlOk

`func (o *V2RuleBotChallengeAction) GetRobotChallengeChallengeTtlOk() (*int32, bool)`

GetRobotChallengeChallengeTtlOk returns a tuple with the RobotChallengeChallengeTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRobotChallengeChallengeTtl

`func (o *V2RuleBotChallengeAction) SetRobotChallengeChallengeTtl(v int32)`

SetRobotChallengeChallengeTtl sets RobotChallengeChallengeTtl field to given value.

### HasRobotChallengeChallengeTtl

`func (o *V2RuleBotChallengeAction) HasRobotChallengeChallengeTtl() bool`

HasRobotChallengeChallengeTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


