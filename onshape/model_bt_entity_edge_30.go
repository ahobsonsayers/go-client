/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.157.9035-1a253bfe8c38
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTEntityEdge30 struct for BTEntityEdge30
type BTEntityEdge30 struct {
	BtType                      *string                `json:"btType,omitempty"`
	Compressed                  *bool                  `json:"compressed,omitempty"`
	Decompressed                *BTEntityGeometry35    `json:"decompressed,omitempty"`
	ErrorCode                   *int32                 `json:"errorCode,omitempty"`
	EstimatedMemoryUsageInBytes *int32                 `json:"estimatedMemoryUsageInBytes,omitempty"`
	Face                        *bool                  `json:"face,omitempty"`
	HasTessellationError        *bool                  `json:"hasTessellationError,omitempty"`
	SettingIndex                *int32                 `json:"settingIndex,omitempty"`
	CompressedPoints            *BTImmutableByteArray  `json:"compressedPoints,omitempty"`
	EdgeSmoothnessStatus        *string                `json:"edgeSmoothnessStatus,omitempty"`
	EdgeType                    *string                `json:"edgeType,omitempty"`
	IsClosed                    *bool                  `json:"isClosed,omitempty"`
	IsInternalEdge              *bool                  `json:"isInternalEdge,omitempty"`
	Points                      *BTImmutableFloatArray `json:"points,omitempty"`
}

// NewBTEntityEdge30 instantiates a new BTEntityEdge30 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTEntityEdge30() *BTEntityEdge30 {
	this := BTEntityEdge30{}
	return &this
}

// NewBTEntityEdge30WithDefaults instantiates a new BTEntityEdge30 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTEntityEdge30WithDefaults() *BTEntityEdge30 {
	this := BTEntityEdge30{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTEntityEdge30) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityEdge30) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTEntityEdge30) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTEntityEdge30) SetBtType(v string) {
	o.BtType = &v
}

// GetCompressed returns the Compressed field value if set, zero value otherwise.
func (o *BTEntityEdge30) GetCompressed() bool {
	if o == nil || o.Compressed == nil {
		var ret bool
		return ret
	}
	return *o.Compressed
}

// GetCompressedOk returns a tuple with the Compressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityEdge30) GetCompressedOk() (*bool, bool) {
	if o == nil || o.Compressed == nil {
		return nil, false
	}
	return o.Compressed, true
}

// HasCompressed returns a boolean if a field has been set.
func (o *BTEntityEdge30) HasCompressed() bool {
	if o != nil && o.Compressed != nil {
		return true
	}

	return false
}

// SetCompressed gets a reference to the given bool and assigns it to the Compressed field.
func (o *BTEntityEdge30) SetCompressed(v bool) {
	o.Compressed = &v
}

// GetDecompressed returns the Decompressed field value if set, zero value otherwise.
func (o *BTEntityEdge30) GetDecompressed() BTEntityGeometry35 {
	if o == nil || o.Decompressed == nil {
		var ret BTEntityGeometry35
		return ret
	}
	return *o.Decompressed
}

// GetDecompressedOk returns a tuple with the Decompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityEdge30) GetDecompressedOk() (*BTEntityGeometry35, bool) {
	if o == nil || o.Decompressed == nil {
		return nil, false
	}
	return o.Decompressed, true
}

// HasDecompressed returns a boolean if a field has been set.
func (o *BTEntityEdge30) HasDecompressed() bool {
	if o != nil && o.Decompressed != nil {
		return true
	}

	return false
}

// SetDecompressed gets a reference to the given BTEntityGeometry35 and assigns it to the Decompressed field.
func (o *BTEntityEdge30) SetDecompressed(v BTEntityGeometry35) {
	o.Decompressed = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *BTEntityEdge30) GetErrorCode() int32 {
	if o == nil || o.ErrorCode == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityEdge30) GetErrorCodeOk() (*int32, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *BTEntityEdge30) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *BTEntityEdge30) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetEstimatedMemoryUsageInBytes returns the EstimatedMemoryUsageInBytes field value if set, zero value otherwise.
func (o *BTEntityEdge30) GetEstimatedMemoryUsageInBytes() int32 {
	if o == nil || o.EstimatedMemoryUsageInBytes == nil {
		var ret int32
		return ret
	}
	return *o.EstimatedMemoryUsageInBytes
}

// GetEstimatedMemoryUsageInBytesOk returns a tuple with the EstimatedMemoryUsageInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityEdge30) GetEstimatedMemoryUsageInBytesOk() (*int32, bool) {
	if o == nil || o.EstimatedMemoryUsageInBytes == nil {
		return nil, false
	}
	return o.EstimatedMemoryUsageInBytes, true
}

// HasEstimatedMemoryUsageInBytes returns a boolean if a field has been set.
func (o *BTEntityEdge30) HasEstimatedMemoryUsageInBytes() bool {
	if o != nil && o.EstimatedMemoryUsageInBytes != nil {
		return true
	}

	return false
}

// SetEstimatedMemoryUsageInBytes gets a reference to the given int32 and assigns it to the EstimatedMemoryUsageInBytes field.
func (o *BTEntityEdge30) SetEstimatedMemoryUsageInBytes(v int32) {
	o.EstimatedMemoryUsageInBytes = &v
}

// GetFace returns the Face field value if set, zero value otherwise.
func (o *BTEntityEdge30) GetFace() bool {
	if o == nil || o.Face == nil {
		var ret bool
		return ret
	}
	return *o.Face
}

// GetFaceOk returns a tuple with the Face field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityEdge30) GetFaceOk() (*bool, bool) {
	if o == nil || o.Face == nil {
		return nil, false
	}
	return o.Face, true
}

// HasFace returns a boolean if a field has been set.
func (o *BTEntityEdge30) HasFace() bool {
	if o != nil && o.Face != nil {
		return true
	}

	return false
}

// SetFace gets a reference to the given bool and assigns it to the Face field.
func (o *BTEntityEdge30) SetFace(v bool) {
	o.Face = &v
}

// GetHasTessellationError returns the HasTessellationError field value if set, zero value otherwise.
func (o *BTEntityEdge30) GetHasTessellationError() bool {
	if o == nil || o.HasTessellationError == nil {
		var ret bool
		return ret
	}
	return *o.HasTessellationError
}

// GetHasTessellationErrorOk returns a tuple with the HasTessellationError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityEdge30) GetHasTessellationErrorOk() (*bool, bool) {
	if o == nil || o.HasTessellationError == nil {
		return nil, false
	}
	return o.HasTessellationError, true
}

// HasHasTessellationError returns a boolean if a field has been set.
func (o *BTEntityEdge30) HasHasTessellationError() bool {
	if o != nil && o.HasTessellationError != nil {
		return true
	}

	return false
}

// SetHasTessellationError gets a reference to the given bool and assigns it to the HasTessellationError field.
func (o *BTEntityEdge30) SetHasTessellationError(v bool) {
	o.HasTessellationError = &v
}

// GetSettingIndex returns the SettingIndex field value if set, zero value otherwise.
func (o *BTEntityEdge30) GetSettingIndex() int32 {
	if o == nil || o.SettingIndex == nil {
		var ret int32
		return ret
	}
	return *o.SettingIndex
}

// GetSettingIndexOk returns a tuple with the SettingIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityEdge30) GetSettingIndexOk() (*int32, bool) {
	if o == nil || o.SettingIndex == nil {
		return nil, false
	}
	return o.SettingIndex, true
}

// HasSettingIndex returns a boolean if a field has been set.
func (o *BTEntityEdge30) HasSettingIndex() bool {
	if o != nil && o.SettingIndex != nil {
		return true
	}

	return false
}

// SetSettingIndex gets a reference to the given int32 and assigns it to the SettingIndex field.
func (o *BTEntityEdge30) SetSettingIndex(v int32) {
	o.SettingIndex = &v
}

// GetCompressedPoints returns the CompressedPoints field value if set, zero value otherwise.
func (o *BTEntityEdge30) GetCompressedPoints() BTImmutableByteArray {
	if o == nil || o.CompressedPoints == nil {
		var ret BTImmutableByteArray
		return ret
	}
	return *o.CompressedPoints
}

// GetCompressedPointsOk returns a tuple with the CompressedPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityEdge30) GetCompressedPointsOk() (*BTImmutableByteArray, bool) {
	if o == nil || o.CompressedPoints == nil {
		return nil, false
	}
	return o.CompressedPoints, true
}

// HasCompressedPoints returns a boolean if a field has been set.
func (o *BTEntityEdge30) HasCompressedPoints() bool {
	if o != nil && o.CompressedPoints != nil {
		return true
	}

	return false
}

// SetCompressedPoints gets a reference to the given BTImmutableByteArray and assigns it to the CompressedPoints field.
func (o *BTEntityEdge30) SetCompressedPoints(v BTImmutableByteArray) {
	o.CompressedPoints = &v
}

// GetEdgeSmoothnessStatus returns the EdgeSmoothnessStatus field value if set, zero value otherwise.
func (o *BTEntityEdge30) GetEdgeSmoothnessStatus() string {
	if o == nil || o.EdgeSmoothnessStatus == nil {
		var ret string
		return ret
	}
	return *o.EdgeSmoothnessStatus
}

// GetEdgeSmoothnessStatusOk returns a tuple with the EdgeSmoothnessStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityEdge30) GetEdgeSmoothnessStatusOk() (*string, bool) {
	if o == nil || o.EdgeSmoothnessStatus == nil {
		return nil, false
	}
	return o.EdgeSmoothnessStatus, true
}

// HasEdgeSmoothnessStatus returns a boolean if a field has been set.
func (o *BTEntityEdge30) HasEdgeSmoothnessStatus() bool {
	if o != nil && o.EdgeSmoothnessStatus != nil {
		return true
	}

	return false
}

// SetEdgeSmoothnessStatus gets a reference to the given string and assigns it to the EdgeSmoothnessStatus field.
func (o *BTEntityEdge30) SetEdgeSmoothnessStatus(v string) {
	o.EdgeSmoothnessStatus = &v
}

// GetEdgeType returns the EdgeType field value if set, zero value otherwise.
func (o *BTEntityEdge30) GetEdgeType() string {
	if o == nil || o.EdgeType == nil {
		var ret string
		return ret
	}
	return *o.EdgeType
}

// GetEdgeTypeOk returns a tuple with the EdgeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityEdge30) GetEdgeTypeOk() (*string, bool) {
	if o == nil || o.EdgeType == nil {
		return nil, false
	}
	return o.EdgeType, true
}

// HasEdgeType returns a boolean if a field has been set.
func (o *BTEntityEdge30) HasEdgeType() bool {
	if o != nil && o.EdgeType != nil {
		return true
	}

	return false
}

// SetEdgeType gets a reference to the given string and assigns it to the EdgeType field.
func (o *BTEntityEdge30) SetEdgeType(v string) {
	o.EdgeType = &v
}

// GetIsClosed returns the IsClosed field value if set, zero value otherwise.
func (o *BTEntityEdge30) GetIsClosed() bool {
	if o == nil || o.IsClosed == nil {
		var ret bool
		return ret
	}
	return *o.IsClosed
}

// GetIsClosedOk returns a tuple with the IsClosed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityEdge30) GetIsClosedOk() (*bool, bool) {
	if o == nil || o.IsClosed == nil {
		return nil, false
	}
	return o.IsClosed, true
}

// HasIsClosed returns a boolean if a field has been set.
func (o *BTEntityEdge30) HasIsClosed() bool {
	if o != nil && o.IsClosed != nil {
		return true
	}

	return false
}

// SetIsClosed gets a reference to the given bool and assigns it to the IsClosed field.
func (o *BTEntityEdge30) SetIsClosed(v bool) {
	o.IsClosed = &v
}

// GetIsInternalEdge returns the IsInternalEdge field value if set, zero value otherwise.
func (o *BTEntityEdge30) GetIsInternalEdge() bool {
	if o == nil || o.IsInternalEdge == nil {
		var ret bool
		return ret
	}
	return *o.IsInternalEdge
}

// GetIsInternalEdgeOk returns a tuple with the IsInternalEdge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityEdge30) GetIsInternalEdgeOk() (*bool, bool) {
	if o == nil || o.IsInternalEdge == nil {
		return nil, false
	}
	return o.IsInternalEdge, true
}

// HasIsInternalEdge returns a boolean if a field has been set.
func (o *BTEntityEdge30) HasIsInternalEdge() bool {
	if o != nil && o.IsInternalEdge != nil {
		return true
	}

	return false
}

// SetIsInternalEdge gets a reference to the given bool and assigns it to the IsInternalEdge field.
func (o *BTEntityEdge30) SetIsInternalEdge(v bool) {
	o.IsInternalEdge = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *BTEntityEdge30) GetPoints() BTImmutableFloatArray {
	if o == nil || o.Points == nil {
		var ret BTImmutableFloatArray
		return ret
	}
	return *o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityEdge30) GetPointsOk() (*BTImmutableFloatArray, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *BTEntityEdge30) HasPoints() bool {
	if o != nil && o.Points != nil {
		return true
	}

	return false
}

// SetPoints gets a reference to the given BTImmutableFloatArray and assigns it to the Points field.
func (o *BTEntityEdge30) SetPoints(v BTImmutableFloatArray) {
	o.Points = &v
}

func (o BTEntityEdge30) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Compressed != nil {
		toSerialize["compressed"] = o.Compressed
	}
	if o.Decompressed != nil {
		toSerialize["decompressed"] = o.Decompressed
	}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.EstimatedMemoryUsageInBytes != nil {
		toSerialize["estimatedMemoryUsageInBytes"] = o.EstimatedMemoryUsageInBytes
	}
	if o.Face != nil {
		toSerialize["face"] = o.Face
	}
	if o.HasTessellationError != nil {
		toSerialize["hasTessellationError"] = o.HasTessellationError
	}
	if o.SettingIndex != nil {
		toSerialize["settingIndex"] = o.SettingIndex
	}
	if o.CompressedPoints != nil {
		toSerialize["compressedPoints"] = o.CompressedPoints
	}
	if o.EdgeSmoothnessStatus != nil {
		toSerialize["edgeSmoothnessStatus"] = o.EdgeSmoothnessStatus
	}
	if o.EdgeType != nil {
		toSerialize["edgeType"] = o.EdgeType
	}
	if o.IsClosed != nil {
		toSerialize["isClosed"] = o.IsClosed
	}
	if o.IsInternalEdge != nil {
		toSerialize["isInternalEdge"] = o.IsInternalEdge
	}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}
	return json.Marshal(toSerialize)
}

type NullableBTEntityEdge30 struct {
	value *BTEntityEdge30
	isSet bool
}

func (v NullableBTEntityEdge30) Get() *BTEntityEdge30 {
	return v.value
}

func (v *NullableBTEntityEdge30) Set(val *BTEntityEdge30) {
	v.value = val
	v.isSet = true
}

func (v NullableBTEntityEdge30) IsSet() bool {
	return v.isSet
}

func (v *NullableBTEntityEdge30) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTEntityEdge30(val *BTEntityEdge30) *NullableBTEntityEdge30 {
	return &NullableBTEntityEdge30{value: val, isSet: true}
}

func (v NullableBTEntityEdge30) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTEntityEdge30) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
