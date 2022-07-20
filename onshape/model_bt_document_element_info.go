/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.150.5616-564f6a8676f1
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTDocumentElementInfo struct for BTDocumentElementInfo
type BTDocumentElementInfo struct {
	AccelerationUnits    *string          `json:"accelerationUnits,omitempty"`
	AngleUnits           *string          `json:"angleUnits,omitempty"`
	AngularVelocityUnits *string          `json:"angularVelocityUnits,omitempty"`
	AreaUnits            *string          `json:"areaUnits,omitempty"`
	DataType             *string          `json:"dataType,omitempty"`
	Deleted              *bool            `json:"deleted,omitempty"`
	ElementType          *string          `json:"elementType,omitempty"`
	EnergyUnits          *string          `json:"energyUnits,omitempty"`
	Filename             *string          `json:"filename,omitempty"`
	ForceUnits           *string          `json:"forceUnits,omitempty"`
	ForeignDataId        *string          `json:"foreignDataId,omitempty"`
	Id                   *string          `json:"id,omitempty"`
	LengthUnits          *string          `json:"lengthUnits,omitempty"`
	MassUnits            *string          `json:"massUnits,omitempty"`
	MicroversionId       *string          `json:"microversionId,omitempty"`
	MomentUnits          *string          `json:"momentUnits,omitempty"`
	Name                 *string          `json:"name,omitempty"`
	PressureUnits        *string          `json:"pressureUnits,omitempty"`
	SpecifiedUnit        *string          `json:"specifiedUnit,omitempty"`
	ThumbnailInfo        *BTThumbnailInfo `json:"thumbnailInfo,omitempty"`
	Thumbnails           *string          `json:"thumbnails,omitempty"`
	TimeUnits            *string          `json:"timeUnits,omitempty"`
	Type                 *string          `json:"type,omitempty"`
	Unupdatable          *bool            `json:"unupdatable,omitempty"`
	VolumeUnits          *string          `json:"volumeUnits,omitempty"`
}

// NewBTDocumentElementInfo instantiates a new BTDocumentElementInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDocumentElementInfo() *BTDocumentElementInfo {
	this := BTDocumentElementInfo{}
	return &this
}

// NewBTDocumentElementInfoWithDefaults instantiates a new BTDocumentElementInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDocumentElementInfoWithDefaults() *BTDocumentElementInfo {
	this := BTDocumentElementInfo{}
	return &this
}

// GetAccelerationUnits returns the AccelerationUnits field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetAccelerationUnits() string {
	if o == nil || o.AccelerationUnits == nil {
		var ret string
		return ret
	}
	return *o.AccelerationUnits
}

// GetAccelerationUnitsOk returns a tuple with the AccelerationUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetAccelerationUnitsOk() (*string, bool) {
	if o == nil || o.AccelerationUnits == nil {
		return nil, false
	}
	return o.AccelerationUnits, true
}

// HasAccelerationUnits returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasAccelerationUnits() bool {
	if o != nil && o.AccelerationUnits != nil {
		return true
	}

	return false
}

// SetAccelerationUnits gets a reference to the given string and assigns it to the AccelerationUnits field.
func (o *BTDocumentElementInfo) SetAccelerationUnits(v string) {
	o.AccelerationUnits = &v
}

// GetAngleUnits returns the AngleUnits field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetAngleUnits() string {
	if o == nil || o.AngleUnits == nil {
		var ret string
		return ret
	}
	return *o.AngleUnits
}

// GetAngleUnitsOk returns a tuple with the AngleUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetAngleUnitsOk() (*string, bool) {
	if o == nil || o.AngleUnits == nil {
		return nil, false
	}
	return o.AngleUnits, true
}

// HasAngleUnits returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasAngleUnits() bool {
	if o != nil && o.AngleUnits != nil {
		return true
	}

	return false
}

// SetAngleUnits gets a reference to the given string and assigns it to the AngleUnits field.
func (o *BTDocumentElementInfo) SetAngleUnits(v string) {
	o.AngleUnits = &v
}

// GetAngularVelocityUnits returns the AngularVelocityUnits field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetAngularVelocityUnits() string {
	if o == nil || o.AngularVelocityUnits == nil {
		var ret string
		return ret
	}
	return *o.AngularVelocityUnits
}

// GetAngularVelocityUnitsOk returns a tuple with the AngularVelocityUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetAngularVelocityUnitsOk() (*string, bool) {
	if o == nil || o.AngularVelocityUnits == nil {
		return nil, false
	}
	return o.AngularVelocityUnits, true
}

// HasAngularVelocityUnits returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasAngularVelocityUnits() bool {
	if o != nil && o.AngularVelocityUnits != nil {
		return true
	}

	return false
}

// SetAngularVelocityUnits gets a reference to the given string and assigns it to the AngularVelocityUnits field.
func (o *BTDocumentElementInfo) SetAngularVelocityUnits(v string) {
	o.AngularVelocityUnits = &v
}

// GetAreaUnits returns the AreaUnits field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetAreaUnits() string {
	if o == nil || o.AreaUnits == nil {
		var ret string
		return ret
	}
	return *o.AreaUnits
}

// GetAreaUnitsOk returns a tuple with the AreaUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetAreaUnitsOk() (*string, bool) {
	if o == nil || o.AreaUnits == nil {
		return nil, false
	}
	return o.AreaUnits, true
}

// HasAreaUnits returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasAreaUnits() bool {
	if o != nil && o.AreaUnits != nil {
		return true
	}

	return false
}

// SetAreaUnits gets a reference to the given string and assigns it to the AreaUnits field.
func (o *BTDocumentElementInfo) SetAreaUnits(v string) {
	o.AreaUnits = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetDataType() string {
	if o == nil || o.DataType == nil {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetDataTypeOk() (*string, bool) {
	if o == nil || o.DataType == nil {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasDataType() bool {
	if o != nil && o.DataType != nil {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *BTDocumentElementInfo) SetDataType(v string) {
	o.DataType = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *BTDocumentElementInfo) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetElementType returns the ElementType field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetElementType() string {
	if o == nil || o.ElementType == nil {
		var ret string
		return ret
	}
	return *o.ElementType
}

// GetElementTypeOk returns a tuple with the ElementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetElementTypeOk() (*string, bool) {
	if o == nil || o.ElementType == nil {
		return nil, false
	}
	return o.ElementType, true
}

// HasElementType returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasElementType() bool {
	if o != nil && o.ElementType != nil {
		return true
	}

	return false
}

// SetElementType gets a reference to the given string and assigns it to the ElementType field.
func (o *BTDocumentElementInfo) SetElementType(v string) {
	o.ElementType = &v
}

// GetEnergyUnits returns the EnergyUnits field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetEnergyUnits() string {
	if o == nil || o.EnergyUnits == nil {
		var ret string
		return ret
	}
	return *o.EnergyUnits
}

// GetEnergyUnitsOk returns a tuple with the EnergyUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetEnergyUnitsOk() (*string, bool) {
	if o == nil || o.EnergyUnits == nil {
		return nil, false
	}
	return o.EnergyUnits, true
}

// HasEnergyUnits returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasEnergyUnits() bool {
	if o != nil && o.EnergyUnits != nil {
		return true
	}

	return false
}

// SetEnergyUnits gets a reference to the given string and assigns it to the EnergyUnits field.
func (o *BTDocumentElementInfo) SetEnergyUnits(v string) {
	o.EnergyUnits = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *BTDocumentElementInfo) SetFilename(v string) {
	o.Filename = &v
}

// GetForceUnits returns the ForceUnits field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetForceUnits() string {
	if o == nil || o.ForceUnits == nil {
		var ret string
		return ret
	}
	return *o.ForceUnits
}

// GetForceUnitsOk returns a tuple with the ForceUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetForceUnitsOk() (*string, bool) {
	if o == nil || o.ForceUnits == nil {
		return nil, false
	}
	return o.ForceUnits, true
}

// HasForceUnits returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasForceUnits() bool {
	if o != nil && o.ForceUnits != nil {
		return true
	}

	return false
}

// SetForceUnits gets a reference to the given string and assigns it to the ForceUnits field.
func (o *BTDocumentElementInfo) SetForceUnits(v string) {
	o.ForceUnits = &v
}

// GetForeignDataId returns the ForeignDataId field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetForeignDataId() string {
	if o == nil || o.ForeignDataId == nil {
		var ret string
		return ret
	}
	return *o.ForeignDataId
}

// GetForeignDataIdOk returns a tuple with the ForeignDataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetForeignDataIdOk() (*string, bool) {
	if o == nil || o.ForeignDataId == nil {
		return nil, false
	}
	return o.ForeignDataId, true
}

// HasForeignDataId returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasForeignDataId() bool {
	if o != nil && o.ForeignDataId != nil {
		return true
	}

	return false
}

// SetForeignDataId gets a reference to the given string and assigns it to the ForeignDataId field.
func (o *BTDocumentElementInfo) SetForeignDataId(v string) {
	o.ForeignDataId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTDocumentElementInfo) SetId(v string) {
	o.Id = &v
}

// GetLengthUnits returns the LengthUnits field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetLengthUnits() string {
	if o == nil || o.LengthUnits == nil {
		var ret string
		return ret
	}
	return *o.LengthUnits
}

// GetLengthUnitsOk returns a tuple with the LengthUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetLengthUnitsOk() (*string, bool) {
	if o == nil || o.LengthUnits == nil {
		return nil, false
	}
	return o.LengthUnits, true
}

// HasLengthUnits returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasLengthUnits() bool {
	if o != nil && o.LengthUnits != nil {
		return true
	}

	return false
}

// SetLengthUnits gets a reference to the given string and assigns it to the LengthUnits field.
func (o *BTDocumentElementInfo) SetLengthUnits(v string) {
	o.LengthUnits = &v
}

// GetMassUnits returns the MassUnits field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetMassUnits() string {
	if o == nil || o.MassUnits == nil {
		var ret string
		return ret
	}
	return *o.MassUnits
}

// GetMassUnitsOk returns a tuple with the MassUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetMassUnitsOk() (*string, bool) {
	if o == nil || o.MassUnits == nil {
		return nil, false
	}
	return o.MassUnits, true
}

// HasMassUnits returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasMassUnits() bool {
	if o != nil && o.MassUnits != nil {
		return true
	}

	return false
}

// SetMassUnits gets a reference to the given string and assigns it to the MassUnits field.
func (o *BTDocumentElementInfo) SetMassUnits(v string) {
	o.MassUnits = &v
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetMicroversionId() string {
	if o == nil || o.MicroversionId == nil {
		var ret string
		return ret
	}
	return *o.MicroversionId
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetMicroversionIdOk() (*string, bool) {
	if o == nil || o.MicroversionId == nil {
		return nil, false
	}
	return o.MicroversionId, true
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasMicroversionId() bool {
	if o != nil && o.MicroversionId != nil {
		return true
	}

	return false
}

// SetMicroversionId gets a reference to the given string and assigns it to the MicroversionId field.
func (o *BTDocumentElementInfo) SetMicroversionId(v string) {
	o.MicroversionId = &v
}

// GetMomentUnits returns the MomentUnits field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetMomentUnits() string {
	if o == nil || o.MomentUnits == nil {
		var ret string
		return ret
	}
	return *o.MomentUnits
}

// GetMomentUnitsOk returns a tuple with the MomentUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetMomentUnitsOk() (*string, bool) {
	if o == nil || o.MomentUnits == nil {
		return nil, false
	}
	return o.MomentUnits, true
}

// HasMomentUnits returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasMomentUnits() bool {
	if o != nil && o.MomentUnits != nil {
		return true
	}

	return false
}

// SetMomentUnits gets a reference to the given string and assigns it to the MomentUnits field.
func (o *BTDocumentElementInfo) SetMomentUnits(v string) {
	o.MomentUnits = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTDocumentElementInfo) SetName(v string) {
	o.Name = &v
}

// GetPressureUnits returns the PressureUnits field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetPressureUnits() string {
	if o == nil || o.PressureUnits == nil {
		var ret string
		return ret
	}
	return *o.PressureUnits
}

// GetPressureUnitsOk returns a tuple with the PressureUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetPressureUnitsOk() (*string, bool) {
	if o == nil || o.PressureUnits == nil {
		return nil, false
	}
	return o.PressureUnits, true
}

// HasPressureUnits returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasPressureUnits() bool {
	if o != nil && o.PressureUnits != nil {
		return true
	}

	return false
}

// SetPressureUnits gets a reference to the given string and assigns it to the PressureUnits field.
func (o *BTDocumentElementInfo) SetPressureUnits(v string) {
	o.PressureUnits = &v
}

// GetSpecifiedUnit returns the SpecifiedUnit field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetSpecifiedUnit() string {
	if o == nil || o.SpecifiedUnit == nil {
		var ret string
		return ret
	}
	return *o.SpecifiedUnit
}

// GetSpecifiedUnitOk returns a tuple with the SpecifiedUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetSpecifiedUnitOk() (*string, bool) {
	if o == nil || o.SpecifiedUnit == nil {
		return nil, false
	}
	return o.SpecifiedUnit, true
}

// HasSpecifiedUnit returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasSpecifiedUnit() bool {
	if o != nil && o.SpecifiedUnit != nil {
		return true
	}

	return false
}

// SetSpecifiedUnit gets a reference to the given string and assigns it to the SpecifiedUnit field.
func (o *BTDocumentElementInfo) SetSpecifiedUnit(v string) {
	o.SpecifiedUnit = &v
}

// GetThumbnailInfo returns the ThumbnailInfo field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetThumbnailInfo() BTThumbnailInfo {
	if o == nil || o.ThumbnailInfo == nil {
		var ret BTThumbnailInfo
		return ret
	}
	return *o.ThumbnailInfo
}

// GetThumbnailInfoOk returns a tuple with the ThumbnailInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetThumbnailInfoOk() (*BTThumbnailInfo, bool) {
	if o == nil || o.ThumbnailInfo == nil {
		return nil, false
	}
	return o.ThumbnailInfo, true
}

// HasThumbnailInfo returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasThumbnailInfo() bool {
	if o != nil && o.ThumbnailInfo != nil {
		return true
	}

	return false
}

// SetThumbnailInfo gets a reference to the given BTThumbnailInfo and assigns it to the ThumbnailInfo field.
func (o *BTDocumentElementInfo) SetThumbnailInfo(v BTThumbnailInfo) {
	o.ThumbnailInfo = &v
}

// GetThumbnails returns the Thumbnails field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetThumbnails() string {
	if o == nil || o.Thumbnails == nil {
		var ret string
		return ret
	}
	return *o.Thumbnails
}

// GetThumbnailsOk returns a tuple with the Thumbnails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetThumbnailsOk() (*string, bool) {
	if o == nil || o.Thumbnails == nil {
		return nil, false
	}
	return o.Thumbnails, true
}

// HasThumbnails returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasThumbnails() bool {
	if o != nil && o.Thumbnails != nil {
		return true
	}

	return false
}

// SetThumbnails gets a reference to the given string and assigns it to the Thumbnails field.
func (o *BTDocumentElementInfo) SetThumbnails(v string) {
	o.Thumbnails = &v
}

// GetTimeUnits returns the TimeUnits field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetTimeUnits() string {
	if o == nil || o.TimeUnits == nil {
		var ret string
		return ret
	}
	return *o.TimeUnits
}

// GetTimeUnitsOk returns a tuple with the TimeUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetTimeUnitsOk() (*string, bool) {
	if o == nil || o.TimeUnits == nil {
		return nil, false
	}
	return o.TimeUnits, true
}

// HasTimeUnits returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasTimeUnits() bool {
	if o != nil && o.TimeUnits != nil {
		return true
	}

	return false
}

// SetTimeUnits gets a reference to the given string and assigns it to the TimeUnits field.
func (o *BTDocumentElementInfo) SetTimeUnits(v string) {
	o.TimeUnits = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTDocumentElementInfo) SetType(v string) {
	o.Type = &v
}

// GetUnupdatable returns the Unupdatable field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetUnupdatable() bool {
	if o == nil || o.Unupdatable == nil {
		var ret bool
		return ret
	}
	return *o.Unupdatable
}

// GetUnupdatableOk returns a tuple with the Unupdatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetUnupdatableOk() (*bool, bool) {
	if o == nil || o.Unupdatable == nil {
		return nil, false
	}
	return o.Unupdatable, true
}

// HasUnupdatable returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasUnupdatable() bool {
	if o != nil && o.Unupdatable != nil {
		return true
	}

	return false
}

// SetUnupdatable gets a reference to the given bool and assigns it to the Unupdatable field.
func (o *BTDocumentElementInfo) SetUnupdatable(v bool) {
	o.Unupdatable = &v
}

// GetVolumeUnits returns the VolumeUnits field value if set, zero value otherwise.
func (o *BTDocumentElementInfo) GetVolumeUnits() string {
	if o == nil || o.VolumeUnits == nil {
		var ret string
		return ret
	}
	return *o.VolumeUnits
}

// GetVolumeUnitsOk returns a tuple with the VolumeUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentElementInfo) GetVolumeUnitsOk() (*string, bool) {
	if o == nil || o.VolumeUnits == nil {
		return nil, false
	}
	return o.VolumeUnits, true
}

// HasVolumeUnits returns a boolean if a field has been set.
func (o *BTDocumentElementInfo) HasVolumeUnits() bool {
	if o != nil && o.VolumeUnits != nil {
		return true
	}

	return false
}

// SetVolumeUnits gets a reference to the given string and assigns it to the VolumeUnits field.
func (o *BTDocumentElementInfo) SetVolumeUnits(v string) {
	o.VolumeUnits = &v
}

func (o BTDocumentElementInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccelerationUnits != nil {
		toSerialize["accelerationUnits"] = o.AccelerationUnits
	}
	if o.AngleUnits != nil {
		toSerialize["angleUnits"] = o.AngleUnits
	}
	if o.AngularVelocityUnits != nil {
		toSerialize["angularVelocityUnits"] = o.AngularVelocityUnits
	}
	if o.AreaUnits != nil {
		toSerialize["areaUnits"] = o.AreaUnits
	}
	if o.DataType != nil {
		toSerialize["dataType"] = o.DataType
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.ElementType != nil {
		toSerialize["elementType"] = o.ElementType
	}
	if o.EnergyUnits != nil {
		toSerialize["energyUnits"] = o.EnergyUnits
	}
	if o.Filename != nil {
		toSerialize["filename"] = o.Filename
	}
	if o.ForceUnits != nil {
		toSerialize["forceUnits"] = o.ForceUnits
	}
	if o.ForeignDataId != nil {
		toSerialize["foreignDataId"] = o.ForeignDataId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LengthUnits != nil {
		toSerialize["lengthUnits"] = o.LengthUnits
	}
	if o.MassUnits != nil {
		toSerialize["massUnits"] = o.MassUnits
	}
	if o.MicroversionId != nil {
		toSerialize["microversionId"] = o.MicroversionId
	}
	if o.MomentUnits != nil {
		toSerialize["momentUnits"] = o.MomentUnits
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PressureUnits != nil {
		toSerialize["pressureUnits"] = o.PressureUnits
	}
	if o.SpecifiedUnit != nil {
		toSerialize["specifiedUnit"] = o.SpecifiedUnit
	}
	if o.ThumbnailInfo != nil {
		toSerialize["thumbnailInfo"] = o.ThumbnailInfo
	}
	if o.Thumbnails != nil {
		toSerialize["thumbnails"] = o.Thumbnails
	}
	if o.TimeUnits != nil {
		toSerialize["timeUnits"] = o.TimeUnits
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Unupdatable != nil {
		toSerialize["unupdatable"] = o.Unupdatable
	}
	if o.VolumeUnits != nil {
		toSerialize["volumeUnits"] = o.VolumeUnits
	}
	return json.Marshal(toSerialize)
}

type NullableBTDocumentElementInfo struct {
	value *BTDocumentElementInfo
	isSet bool
}

func (v NullableBTDocumentElementInfo) Get() *BTDocumentElementInfo {
	return v.value
}

func (v *NullableBTDocumentElementInfo) Set(val *BTDocumentElementInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDocumentElementInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDocumentElementInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDocumentElementInfo(val *BTDocumentElementInfo) *NullableBTDocumentElementInfo {
	return &NullableBTDocumentElementInfo{value: val, isSet: true}
}

func (v NullableBTDocumentElementInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDocumentElementInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
