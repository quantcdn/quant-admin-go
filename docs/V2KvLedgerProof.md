# V2KvLedgerProof

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Entry** | Pointer to [**V2KvLedgerProofEntry**](V2KvLedgerProofEntry.md) |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**V2KvLedgerProvenance**](V2KvLedgerProvenance.md) |  | [optional] 
**Period** | Pointer to **string** | Archive period label, YYYYMMDDTHHmmZ. | [optional] 
**ArchiveKey** | Pointer to **string** |  | [optional] 
**LeafIndex** | Pointer to **int32** |  | [optional] 
**TreeSize** | Pointer to **int32** |  | [optional] 
**MerkleRoot** | Pointer to **string** |  | [optional] 
**AuditPath** | Pointer to **[]string** |  | [optional] 
**Digest** | Pointer to **map[string]interface{}** |  | [optional] 
**DigestKey** | Pointer to **string** |  | [optional] 
**Signature** | Pointer to **string** | Base64 DER ECDSA P-256 over SHA-256 of the digest bytes. | [optional] 
**KeyId** | Pointer to **string** |  | [optional] 
**DigestSignedAt** | Pointer to **time.Time** |  | [optional] 
**Anchor** | Pointer to [**V2KvLedgerProofAnchor**](V2KvLedgerProofAnchor.md) |  | [optional] 
**AnchorUnavailable** | Pointer to **bool** |  | [optional] 

## Methods

### NewV2KvLedgerProof

`func NewV2KvLedgerProof(status string, ) *V2KvLedgerProof`

NewV2KvLedgerProof instantiates a new V2KvLedgerProof object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerProofWithDefaults

`func NewV2KvLedgerProofWithDefaults() *V2KvLedgerProof`

NewV2KvLedgerProofWithDefaults instantiates a new V2KvLedgerProof object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *V2KvLedgerProof) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V2KvLedgerProof) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V2KvLedgerProof) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetEntry

`func (o *V2KvLedgerProof) GetEntry() V2KvLedgerProofEntry`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *V2KvLedgerProof) GetEntryOk() (*V2KvLedgerProofEntry, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *V2KvLedgerProof) SetEntry(v V2KvLedgerProofEntry)`

SetEntry sets Entry field to given value.

### HasEntry

`func (o *V2KvLedgerProof) HasEntry() bool`

HasEntry returns a boolean if a field has been set.

### GetHash

`func (o *V2KvLedgerProof) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *V2KvLedgerProof) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *V2KvLedgerProof) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *V2KvLedgerProof) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetSource

`func (o *V2KvLedgerProof) GetSource() V2KvLedgerProvenance`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *V2KvLedgerProof) GetSourceOk() (*V2KvLedgerProvenance, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *V2KvLedgerProof) SetSource(v V2KvLedgerProvenance)`

SetSource sets Source field to given value.

### HasSource

`func (o *V2KvLedgerProof) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPeriod

`func (o *V2KvLedgerProof) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *V2KvLedgerProof) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *V2KvLedgerProof) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *V2KvLedgerProof) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetArchiveKey

`func (o *V2KvLedgerProof) GetArchiveKey() string`

GetArchiveKey returns the ArchiveKey field if non-nil, zero value otherwise.

### GetArchiveKeyOk

`func (o *V2KvLedgerProof) GetArchiveKeyOk() (*string, bool)`

GetArchiveKeyOk returns a tuple with the ArchiveKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveKey

`func (o *V2KvLedgerProof) SetArchiveKey(v string)`

SetArchiveKey sets ArchiveKey field to given value.

### HasArchiveKey

`func (o *V2KvLedgerProof) HasArchiveKey() bool`

HasArchiveKey returns a boolean if a field has been set.

### GetLeafIndex

`func (o *V2KvLedgerProof) GetLeafIndex() int32`

GetLeafIndex returns the LeafIndex field if non-nil, zero value otherwise.

### GetLeafIndexOk

`func (o *V2KvLedgerProof) GetLeafIndexOk() (*int32, bool)`

GetLeafIndexOk returns a tuple with the LeafIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafIndex

`func (o *V2KvLedgerProof) SetLeafIndex(v int32)`

SetLeafIndex sets LeafIndex field to given value.

### HasLeafIndex

`func (o *V2KvLedgerProof) HasLeafIndex() bool`

HasLeafIndex returns a boolean if a field has been set.

### GetTreeSize

`func (o *V2KvLedgerProof) GetTreeSize() int32`

GetTreeSize returns the TreeSize field if non-nil, zero value otherwise.

### GetTreeSizeOk

`func (o *V2KvLedgerProof) GetTreeSizeOk() (*int32, bool)`

GetTreeSizeOk returns a tuple with the TreeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreeSize

`func (o *V2KvLedgerProof) SetTreeSize(v int32)`

SetTreeSize sets TreeSize field to given value.

### HasTreeSize

`func (o *V2KvLedgerProof) HasTreeSize() bool`

HasTreeSize returns a boolean if a field has been set.

### GetMerkleRoot

`func (o *V2KvLedgerProof) GetMerkleRoot() string`

GetMerkleRoot returns the MerkleRoot field if non-nil, zero value otherwise.

### GetMerkleRootOk

`func (o *V2KvLedgerProof) GetMerkleRootOk() (*string, bool)`

GetMerkleRootOk returns a tuple with the MerkleRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleRoot

`func (o *V2KvLedgerProof) SetMerkleRoot(v string)`

SetMerkleRoot sets MerkleRoot field to given value.

### HasMerkleRoot

`func (o *V2KvLedgerProof) HasMerkleRoot() bool`

HasMerkleRoot returns a boolean if a field has been set.

### GetAuditPath

`func (o *V2KvLedgerProof) GetAuditPath() []string`

GetAuditPath returns the AuditPath field if non-nil, zero value otherwise.

### GetAuditPathOk

`func (o *V2KvLedgerProof) GetAuditPathOk() (*[]string, bool)`

GetAuditPathOk returns a tuple with the AuditPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditPath

`func (o *V2KvLedgerProof) SetAuditPath(v []string)`

SetAuditPath sets AuditPath field to given value.

### HasAuditPath

`func (o *V2KvLedgerProof) HasAuditPath() bool`

HasAuditPath returns a boolean if a field has been set.

### GetDigest

`func (o *V2KvLedgerProof) GetDigest() map[string]interface{}`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *V2KvLedgerProof) GetDigestOk() (*map[string]interface{}, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *V2KvLedgerProof) SetDigest(v map[string]interface{})`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *V2KvLedgerProof) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetDigestKey

`func (o *V2KvLedgerProof) GetDigestKey() string`

GetDigestKey returns the DigestKey field if non-nil, zero value otherwise.

### GetDigestKeyOk

`func (o *V2KvLedgerProof) GetDigestKeyOk() (*string, bool)`

GetDigestKeyOk returns a tuple with the DigestKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestKey

`func (o *V2KvLedgerProof) SetDigestKey(v string)`

SetDigestKey sets DigestKey field to given value.

### HasDigestKey

`func (o *V2KvLedgerProof) HasDigestKey() bool`

HasDigestKey returns a boolean if a field has been set.

### GetSignature

`func (o *V2KvLedgerProof) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *V2KvLedgerProof) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *V2KvLedgerProof) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *V2KvLedgerProof) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetKeyId

`func (o *V2KvLedgerProof) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *V2KvLedgerProof) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *V2KvLedgerProof) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *V2KvLedgerProof) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetDigestSignedAt

`func (o *V2KvLedgerProof) GetDigestSignedAt() time.Time`

GetDigestSignedAt returns the DigestSignedAt field if non-nil, zero value otherwise.

### GetDigestSignedAtOk

`func (o *V2KvLedgerProof) GetDigestSignedAtOk() (*time.Time, bool)`

GetDigestSignedAtOk returns a tuple with the DigestSignedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestSignedAt

`func (o *V2KvLedgerProof) SetDigestSignedAt(v time.Time)`

SetDigestSignedAt sets DigestSignedAt field to given value.

### HasDigestSignedAt

`func (o *V2KvLedgerProof) HasDigestSignedAt() bool`

HasDigestSignedAt returns a boolean if a field has been set.

### GetAnchor

`func (o *V2KvLedgerProof) GetAnchor() V2KvLedgerProofAnchor`

GetAnchor returns the Anchor field if non-nil, zero value otherwise.

### GetAnchorOk

`func (o *V2KvLedgerProof) GetAnchorOk() (*V2KvLedgerProofAnchor, bool)`

GetAnchorOk returns a tuple with the Anchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchor

`func (o *V2KvLedgerProof) SetAnchor(v V2KvLedgerProofAnchor)`

SetAnchor sets Anchor field to given value.

### HasAnchor

`func (o *V2KvLedgerProof) HasAnchor() bool`

HasAnchor returns a boolean if a field has been set.

### GetAnchorUnavailable

`func (o *V2KvLedgerProof) GetAnchorUnavailable() bool`

GetAnchorUnavailable returns the AnchorUnavailable field if non-nil, zero value otherwise.

### GetAnchorUnavailableOk

`func (o *V2KvLedgerProof) GetAnchorUnavailableOk() (*bool, bool)`

GetAnchorUnavailableOk returns a tuple with the AnchorUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorUnavailable

`func (o *V2KvLedgerProof) SetAnchorUnavailable(v bool)`

SetAnchorUnavailable sets AnchorUnavailable field to given value.

### HasAnchorUnavailable

`func (o *V2KvLedgerProof) HasAnchorUnavailable() bool`

HasAnchorUnavailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


