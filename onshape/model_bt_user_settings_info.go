/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.166.18273-3025d52f81b7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTUserSettingsInfo struct for BTUserSettingsInfo
type BTUserSettingsInfo struct {
	CommonUnits                     *BTCommonUnitsInfo                     `json:"commonUnits,omitempty"`
	CustomColors                    []string                               `json:"customColors,omitempty"`
	DefaultUnits                    *BTDefaultUnitsInfo                    `json:"defaultUnits,omitempty"`
	DisplayAssemblyProperties       *bool                                  `json:"displayAssemblyProperties,omitempty"`
	DrawingBackgroundId             *int32                                 `json:"drawingBackgroundId,omitempty"`
	EnforceApplicationAcl           *bool                                  `json:"enforceApplicationAcl,omitempty"`
	ExportDrawingOptions            *string                                `json:"exportDrawingOptions,omitempty"`
	ExportSolidOptions              *string                                `json:"exportSolidOptions,omitempty"`
	Id                              *string                                `json:"id,omitempty"`
	ImportOptions                   *string                                `json:"importOptions,omitempty"`
	Locale                          *string                                `json:"locale,omitempty"`
	MaterialLibrarySettings         *BTMaterialLibrarySettingsInfo         `json:"materialLibrarySettings,omitempty"`
	MiniToolbarSettings             *string                                `json:"miniToolbarSettings,omitempty"`
	MouseActions                    *string                                `json:"mouseActions,omitempty"`
	PreviousSketchFont              *string                                `json:"previousSketchFont,omitempty"`
	ReverseScrollWheelZoomDirection *bool                                  `json:"reverseScrollWheelZoomDirection,omitempty"`
	SelectItemViewStateInfos        []BTSelectItemViewStateInfo            `json:"selectItemViewStateInfos,omitempty"`
	StartupPage                     *int32                                 `json:"startupPage,omitempty"`
	SubstituteApprovers             []BTSubstituteApproverInfo             `json:"substituteApprovers,omitempty"`
	UnitsDisplayPrecision           *map[string]int32                      `json:"unitsDisplayPrecision,omitempty"`
	UnitsMaximumDisplayPrecision    *BTUnitsMaximumDisplayPrecisionInfo    `json:"unitsMaximumDisplayPrecision,omitempty"`
	Use24HourTime                   *bool                                  `json:"use24HourTime,omitempty"`
	UseDecimalComma                 *bool                                  `json:"useDecimalComma,omitempty"`
	ViewManipulationMouseKeyMapping *BTViewManipulationMouseKeyMappingInfo `json:"viewManipulationMouseKeyMapping,omitempty"`
	ViewMappingId                   *int32                                 `json:"viewMappingId,omitempty"`
}

// NewBTUserSettingsInfo instantiates a new BTUserSettingsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUserSettingsInfo() *BTUserSettingsInfo {
	this := BTUserSettingsInfo{}
	return &this
}

// NewBTUserSettingsInfoWithDefaults instantiates a new BTUserSettingsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUserSettingsInfoWithDefaults() *BTUserSettingsInfo {
	this := BTUserSettingsInfo{}
	return &this
}

// GetCommonUnits returns the CommonUnits field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetCommonUnits() BTCommonUnitsInfo {
	if o == nil || o.CommonUnits == nil {
		var ret BTCommonUnitsInfo
		return ret
	}
	return *o.CommonUnits
}

// GetCommonUnitsOk returns a tuple with the CommonUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetCommonUnitsOk() (*BTCommonUnitsInfo, bool) {
	if o == nil || o.CommonUnits == nil {
		return nil, false
	}
	return o.CommonUnits, true
}

// HasCommonUnits returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasCommonUnits() bool {
	if o != nil && o.CommonUnits != nil {
		return true
	}

	return false
}

// SetCommonUnits gets a reference to the given BTCommonUnitsInfo and assigns it to the CommonUnits field.
func (o *BTUserSettingsInfo) SetCommonUnits(v BTCommonUnitsInfo) {
	o.CommonUnits = &v
}

// GetCustomColors returns the CustomColors field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetCustomColors() []string {
	if o == nil || o.CustomColors == nil {
		var ret []string
		return ret
	}
	return o.CustomColors
}

// GetCustomColorsOk returns a tuple with the CustomColors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetCustomColorsOk() ([]string, bool) {
	if o == nil || o.CustomColors == nil {
		return nil, false
	}
	return o.CustomColors, true
}

// HasCustomColors returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasCustomColors() bool {
	if o != nil && o.CustomColors != nil {
		return true
	}

	return false
}

// SetCustomColors gets a reference to the given []string and assigns it to the CustomColors field.
func (o *BTUserSettingsInfo) SetCustomColors(v []string) {
	o.CustomColors = v
}

// GetDefaultUnits returns the DefaultUnits field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetDefaultUnits() BTDefaultUnitsInfo {
	if o == nil || o.DefaultUnits == nil {
		var ret BTDefaultUnitsInfo
		return ret
	}
	return *o.DefaultUnits
}

// GetDefaultUnitsOk returns a tuple with the DefaultUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetDefaultUnitsOk() (*BTDefaultUnitsInfo, bool) {
	if o == nil || o.DefaultUnits == nil {
		return nil, false
	}
	return o.DefaultUnits, true
}

// HasDefaultUnits returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasDefaultUnits() bool {
	if o != nil && o.DefaultUnits != nil {
		return true
	}

	return false
}

// SetDefaultUnits gets a reference to the given BTDefaultUnitsInfo and assigns it to the DefaultUnits field.
func (o *BTUserSettingsInfo) SetDefaultUnits(v BTDefaultUnitsInfo) {
	o.DefaultUnits = &v
}

// GetDisplayAssemblyProperties returns the DisplayAssemblyProperties field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetDisplayAssemblyProperties() bool {
	if o == nil || o.DisplayAssemblyProperties == nil {
		var ret bool
		return ret
	}
	return *o.DisplayAssemblyProperties
}

// GetDisplayAssemblyPropertiesOk returns a tuple with the DisplayAssemblyProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetDisplayAssemblyPropertiesOk() (*bool, bool) {
	if o == nil || o.DisplayAssemblyProperties == nil {
		return nil, false
	}
	return o.DisplayAssemblyProperties, true
}

// HasDisplayAssemblyProperties returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasDisplayAssemblyProperties() bool {
	if o != nil && o.DisplayAssemblyProperties != nil {
		return true
	}

	return false
}

// SetDisplayAssemblyProperties gets a reference to the given bool and assigns it to the DisplayAssemblyProperties field.
func (o *BTUserSettingsInfo) SetDisplayAssemblyProperties(v bool) {
	o.DisplayAssemblyProperties = &v
}

// GetDrawingBackgroundId returns the DrawingBackgroundId field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetDrawingBackgroundId() int32 {
	if o == nil || o.DrawingBackgroundId == nil {
		var ret int32
		return ret
	}
	return *o.DrawingBackgroundId
}

// GetDrawingBackgroundIdOk returns a tuple with the DrawingBackgroundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetDrawingBackgroundIdOk() (*int32, bool) {
	if o == nil || o.DrawingBackgroundId == nil {
		return nil, false
	}
	return o.DrawingBackgroundId, true
}

// HasDrawingBackgroundId returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasDrawingBackgroundId() bool {
	if o != nil && o.DrawingBackgroundId != nil {
		return true
	}

	return false
}

// SetDrawingBackgroundId gets a reference to the given int32 and assigns it to the DrawingBackgroundId field.
func (o *BTUserSettingsInfo) SetDrawingBackgroundId(v int32) {
	o.DrawingBackgroundId = &v
}

// GetEnforceApplicationAcl returns the EnforceApplicationAcl field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetEnforceApplicationAcl() bool {
	if o == nil || o.EnforceApplicationAcl == nil {
		var ret bool
		return ret
	}
	return *o.EnforceApplicationAcl
}

// GetEnforceApplicationAclOk returns a tuple with the EnforceApplicationAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetEnforceApplicationAclOk() (*bool, bool) {
	if o == nil || o.EnforceApplicationAcl == nil {
		return nil, false
	}
	return o.EnforceApplicationAcl, true
}

// HasEnforceApplicationAcl returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasEnforceApplicationAcl() bool {
	if o != nil && o.EnforceApplicationAcl != nil {
		return true
	}

	return false
}

// SetEnforceApplicationAcl gets a reference to the given bool and assigns it to the EnforceApplicationAcl field.
func (o *BTUserSettingsInfo) SetEnforceApplicationAcl(v bool) {
	o.EnforceApplicationAcl = &v
}

// GetExportDrawingOptions returns the ExportDrawingOptions field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetExportDrawingOptions() string {
	if o == nil || o.ExportDrawingOptions == nil {
		var ret string
		return ret
	}
	return *o.ExportDrawingOptions
}

// GetExportDrawingOptionsOk returns a tuple with the ExportDrawingOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetExportDrawingOptionsOk() (*string, bool) {
	if o == nil || o.ExportDrawingOptions == nil {
		return nil, false
	}
	return o.ExportDrawingOptions, true
}

// HasExportDrawingOptions returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasExportDrawingOptions() bool {
	if o != nil && o.ExportDrawingOptions != nil {
		return true
	}

	return false
}

// SetExportDrawingOptions gets a reference to the given string and assigns it to the ExportDrawingOptions field.
func (o *BTUserSettingsInfo) SetExportDrawingOptions(v string) {
	o.ExportDrawingOptions = &v
}

// GetExportSolidOptions returns the ExportSolidOptions field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetExportSolidOptions() string {
	if o == nil || o.ExportSolidOptions == nil {
		var ret string
		return ret
	}
	return *o.ExportSolidOptions
}

// GetExportSolidOptionsOk returns a tuple with the ExportSolidOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetExportSolidOptionsOk() (*string, bool) {
	if o == nil || o.ExportSolidOptions == nil {
		return nil, false
	}
	return o.ExportSolidOptions, true
}

// HasExportSolidOptions returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasExportSolidOptions() bool {
	if o != nil && o.ExportSolidOptions != nil {
		return true
	}

	return false
}

// SetExportSolidOptions gets a reference to the given string and assigns it to the ExportSolidOptions field.
func (o *BTUserSettingsInfo) SetExportSolidOptions(v string) {
	o.ExportSolidOptions = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTUserSettingsInfo) SetId(v string) {
	o.Id = &v
}

// GetImportOptions returns the ImportOptions field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetImportOptions() string {
	if o == nil || o.ImportOptions == nil {
		var ret string
		return ret
	}
	return *o.ImportOptions
}

// GetImportOptionsOk returns a tuple with the ImportOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetImportOptionsOk() (*string, bool) {
	if o == nil || o.ImportOptions == nil {
		return nil, false
	}
	return o.ImportOptions, true
}

// HasImportOptions returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasImportOptions() bool {
	if o != nil && o.ImportOptions != nil {
		return true
	}

	return false
}

// SetImportOptions gets a reference to the given string and assigns it to the ImportOptions field.
func (o *BTUserSettingsInfo) SetImportOptions(v string) {
	o.ImportOptions = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetLocale() string {
	if o == nil || o.Locale == nil {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetLocaleOk() (*string, bool) {
	if o == nil || o.Locale == nil {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasLocale() bool {
	if o != nil && o.Locale != nil {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *BTUserSettingsInfo) SetLocale(v string) {
	o.Locale = &v
}

// GetMaterialLibrarySettings returns the MaterialLibrarySettings field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetMaterialLibrarySettings() BTMaterialLibrarySettingsInfo {
	if o == nil || o.MaterialLibrarySettings == nil {
		var ret BTMaterialLibrarySettingsInfo
		return ret
	}
	return *o.MaterialLibrarySettings
}

// GetMaterialLibrarySettingsOk returns a tuple with the MaterialLibrarySettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetMaterialLibrarySettingsOk() (*BTMaterialLibrarySettingsInfo, bool) {
	if o == nil || o.MaterialLibrarySettings == nil {
		return nil, false
	}
	return o.MaterialLibrarySettings, true
}

// HasMaterialLibrarySettings returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasMaterialLibrarySettings() bool {
	if o != nil && o.MaterialLibrarySettings != nil {
		return true
	}

	return false
}

// SetMaterialLibrarySettings gets a reference to the given BTMaterialLibrarySettingsInfo and assigns it to the MaterialLibrarySettings field.
func (o *BTUserSettingsInfo) SetMaterialLibrarySettings(v BTMaterialLibrarySettingsInfo) {
	o.MaterialLibrarySettings = &v
}

// GetMiniToolbarSettings returns the MiniToolbarSettings field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetMiniToolbarSettings() string {
	if o == nil || o.MiniToolbarSettings == nil {
		var ret string
		return ret
	}
	return *o.MiniToolbarSettings
}

// GetMiniToolbarSettingsOk returns a tuple with the MiniToolbarSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetMiniToolbarSettingsOk() (*string, bool) {
	if o == nil || o.MiniToolbarSettings == nil {
		return nil, false
	}
	return o.MiniToolbarSettings, true
}

// HasMiniToolbarSettings returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasMiniToolbarSettings() bool {
	if o != nil && o.MiniToolbarSettings != nil {
		return true
	}

	return false
}

// SetMiniToolbarSettings gets a reference to the given string and assigns it to the MiniToolbarSettings field.
func (o *BTUserSettingsInfo) SetMiniToolbarSettings(v string) {
	o.MiniToolbarSettings = &v
}

// GetMouseActions returns the MouseActions field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetMouseActions() string {
	if o == nil || o.MouseActions == nil {
		var ret string
		return ret
	}
	return *o.MouseActions
}

// GetMouseActionsOk returns a tuple with the MouseActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetMouseActionsOk() (*string, bool) {
	if o == nil || o.MouseActions == nil {
		return nil, false
	}
	return o.MouseActions, true
}

// HasMouseActions returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasMouseActions() bool {
	if o != nil && o.MouseActions != nil {
		return true
	}

	return false
}

// SetMouseActions gets a reference to the given string and assigns it to the MouseActions field.
func (o *BTUserSettingsInfo) SetMouseActions(v string) {
	o.MouseActions = &v
}

// GetPreviousSketchFont returns the PreviousSketchFont field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetPreviousSketchFont() string {
	if o == nil || o.PreviousSketchFont == nil {
		var ret string
		return ret
	}
	return *o.PreviousSketchFont
}

// GetPreviousSketchFontOk returns a tuple with the PreviousSketchFont field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetPreviousSketchFontOk() (*string, bool) {
	if o == nil || o.PreviousSketchFont == nil {
		return nil, false
	}
	return o.PreviousSketchFont, true
}

// HasPreviousSketchFont returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasPreviousSketchFont() bool {
	if o != nil && o.PreviousSketchFont != nil {
		return true
	}

	return false
}

// SetPreviousSketchFont gets a reference to the given string and assigns it to the PreviousSketchFont field.
func (o *BTUserSettingsInfo) SetPreviousSketchFont(v string) {
	o.PreviousSketchFont = &v
}

// GetReverseScrollWheelZoomDirection returns the ReverseScrollWheelZoomDirection field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetReverseScrollWheelZoomDirection() bool {
	if o == nil || o.ReverseScrollWheelZoomDirection == nil {
		var ret bool
		return ret
	}
	return *o.ReverseScrollWheelZoomDirection
}

// GetReverseScrollWheelZoomDirectionOk returns a tuple with the ReverseScrollWheelZoomDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetReverseScrollWheelZoomDirectionOk() (*bool, bool) {
	if o == nil || o.ReverseScrollWheelZoomDirection == nil {
		return nil, false
	}
	return o.ReverseScrollWheelZoomDirection, true
}

// HasReverseScrollWheelZoomDirection returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasReverseScrollWheelZoomDirection() bool {
	if o != nil && o.ReverseScrollWheelZoomDirection != nil {
		return true
	}

	return false
}

// SetReverseScrollWheelZoomDirection gets a reference to the given bool and assigns it to the ReverseScrollWheelZoomDirection field.
func (o *BTUserSettingsInfo) SetReverseScrollWheelZoomDirection(v bool) {
	o.ReverseScrollWheelZoomDirection = &v
}

// GetSelectItemViewStateInfos returns the SelectItemViewStateInfos field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetSelectItemViewStateInfos() []BTSelectItemViewStateInfo {
	if o == nil || o.SelectItemViewStateInfos == nil {
		var ret []BTSelectItemViewStateInfo
		return ret
	}
	return o.SelectItemViewStateInfos
}

// GetSelectItemViewStateInfosOk returns a tuple with the SelectItemViewStateInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetSelectItemViewStateInfosOk() ([]BTSelectItemViewStateInfo, bool) {
	if o == nil || o.SelectItemViewStateInfos == nil {
		return nil, false
	}
	return o.SelectItemViewStateInfos, true
}

// HasSelectItemViewStateInfos returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasSelectItemViewStateInfos() bool {
	if o != nil && o.SelectItemViewStateInfos != nil {
		return true
	}

	return false
}

// SetSelectItemViewStateInfos gets a reference to the given []BTSelectItemViewStateInfo and assigns it to the SelectItemViewStateInfos field.
func (o *BTUserSettingsInfo) SetSelectItemViewStateInfos(v []BTSelectItemViewStateInfo) {
	o.SelectItemViewStateInfos = v
}

// GetStartupPage returns the StartupPage field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetStartupPage() int32 {
	if o == nil || o.StartupPage == nil {
		var ret int32
		return ret
	}
	return *o.StartupPage
}

// GetStartupPageOk returns a tuple with the StartupPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetStartupPageOk() (*int32, bool) {
	if o == nil || o.StartupPage == nil {
		return nil, false
	}
	return o.StartupPage, true
}

// HasStartupPage returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasStartupPage() bool {
	if o != nil && o.StartupPage != nil {
		return true
	}

	return false
}

// SetStartupPage gets a reference to the given int32 and assigns it to the StartupPage field.
func (o *BTUserSettingsInfo) SetStartupPage(v int32) {
	o.StartupPage = &v
}

// GetSubstituteApprovers returns the SubstituteApprovers field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetSubstituteApprovers() []BTSubstituteApproverInfo {
	if o == nil || o.SubstituteApprovers == nil {
		var ret []BTSubstituteApproverInfo
		return ret
	}
	return o.SubstituteApprovers
}

// GetSubstituteApproversOk returns a tuple with the SubstituteApprovers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetSubstituteApproversOk() ([]BTSubstituteApproverInfo, bool) {
	if o == nil || o.SubstituteApprovers == nil {
		return nil, false
	}
	return o.SubstituteApprovers, true
}

// HasSubstituteApprovers returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasSubstituteApprovers() bool {
	if o != nil && o.SubstituteApprovers != nil {
		return true
	}

	return false
}

// SetSubstituteApprovers gets a reference to the given []BTSubstituteApproverInfo and assigns it to the SubstituteApprovers field.
func (o *BTUserSettingsInfo) SetSubstituteApprovers(v []BTSubstituteApproverInfo) {
	o.SubstituteApprovers = v
}

// GetUnitsDisplayPrecision returns the UnitsDisplayPrecision field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetUnitsDisplayPrecision() map[string]int32 {
	if o == nil || o.UnitsDisplayPrecision == nil {
		var ret map[string]int32
		return ret
	}
	return *o.UnitsDisplayPrecision
}

// GetUnitsDisplayPrecisionOk returns a tuple with the UnitsDisplayPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetUnitsDisplayPrecisionOk() (*map[string]int32, bool) {
	if o == nil || o.UnitsDisplayPrecision == nil {
		return nil, false
	}
	return o.UnitsDisplayPrecision, true
}

// HasUnitsDisplayPrecision returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasUnitsDisplayPrecision() bool {
	if o != nil && o.UnitsDisplayPrecision != nil {
		return true
	}

	return false
}

// SetUnitsDisplayPrecision gets a reference to the given map[string]int32 and assigns it to the UnitsDisplayPrecision field.
func (o *BTUserSettingsInfo) SetUnitsDisplayPrecision(v map[string]int32) {
	o.UnitsDisplayPrecision = &v
}

// GetUnitsMaximumDisplayPrecision returns the UnitsMaximumDisplayPrecision field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetUnitsMaximumDisplayPrecision() BTUnitsMaximumDisplayPrecisionInfo {
	if o == nil || o.UnitsMaximumDisplayPrecision == nil {
		var ret BTUnitsMaximumDisplayPrecisionInfo
		return ret
	}
	return *o.UnitsMaximumDisplayPrecision
}

// GetUnitsMaximumDisplayPrecisionOk returns a tuple with the UnitsMaximumDisplayPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetUnitsMaximumDisplayPrecisionOk() (*BTUnitsMaximumDisplayPrecisionInfo, bool) {
	if o == nil || o.UnitsMaximumDisplayPrecision == nil {
		return nil, false
	}
	return o.UnitsMaximumDisplayPrecision, true
}

// HasUnitsMaximumDisplayPrecision returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasUnitsMaximumDisplayPrecision() bool {
	if o != nil && o.UnitsMaximumDisplayPrecision != nil {
		return true
	}

	return false
}

// SetUnitsMaximumDisplayPrecision gets a reference to the given BTUnitsMaximumDisplayPrecisionInfo and assigns it to the UnitsMaximumDisplayPrecision field.
func (o *BTUserSettingsInfo) SetUnitsMaximumDisplayPrecision(v BTUnitsMaximumDisplayPrecisionInfo) {
	o.UnitsMaximumDisplayPrecision = &v
}

// GetUse24HourTime returns the Use24HourTime field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetUse24HourTime() bool {
	if o == nil || o.Use24HourTime == nil {
		var ret bool
		return ret
	}
	return *o.Use24HourTime
}

// GetUse24HourTimeOk returns a tuple with the Use24HourTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetUse24HourTimeOk() (*bool, bool) {
	if o == nil || o.Use24HourTime == nil {
		return nil, false
	}
	return o.Use24HourTime, true
}

// HasUse24HourTime returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasUse24HourTime() bool {
	if o != nil && o.Use24HourTime != nil {
		return true
	}

	return false
}

// SetUse24HourTime gets a reference to the given bool and assigns it to the Use24HourTime field.
func (o *BTUserSettingsInfo) SetUse24HourTime(v bool) {
	o.Use24HourTime = &v
}

// GetUseDecimalComma returns the UseDecimalComma field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetUseDecimalComma() bool {
	if o == nil || o.UseDecimalComma == nil {
		var ret bool
		return ret
	}
	return *o.UseDecimalComma
}

// GetUseDecimalCommaOk returns a tuple with the UseDecimalComma field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetUseDecimalCommaOk() (*bool, bool) {
	if o == nil || o.UseDecimalComma == nil {
		return nil, false
	}
	return o.UseDecimalComma, true
}

// HasUseDecimalComma returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasUseDecimalComma() bool {
	if o != nil && o.UseDecimalComma != nil {
		return true
	}

	return false
}

// SetUseDecimalComma gets a reference to the given bool and assigns it to the UseDecimalComma field.
func (o *BTUserSettingsInfo) SetUseDecimalComma(v bool) {
	o.UseDecimalComma = &v
}

// GetViewManipulationMouseKeyMapping returns the ViewManipulationMouseKeyMapping field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetViewManipulationMouseKeyMapping() BTViewManipulationMouseKeyMappingInfo {
	if o == nil || o.ViewManipulationMouseKeyMapping == nil {
		var ret BTViewManipulationMouseKeyMappingInfo
		return ret
	}
	return *o.ViewManipulationMouseKeyMapping
}

// GetViewManipulationMouseKeyMappingOk returns a tuple with the ViewManipulationMouseKeyMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetViewManipulationMouseKeyMappingOk() (*BTViewManipulationMouseKeyMappingInfo, bool) {
	if o == nil || o.ViewManipulationMouseKeyMapping == nil {
		return nil, false
	}
	return o.ViewManipulationMouseKeyMapping, true
}

// HasViewManipulationMouseKeyMapping returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasViewManipulationMouseKeyMapping() bool {
	if o != nil && o.ViewManipulationMouseKeyMapping != nil {
		return true
	}

	return false
}

// SetViewManipulationMouseKeyMapping gets a reference to the given BTViewManipulationMouseKeyMappingInfo and assigns it to the ViewManipulationMouseKeyMapping field.
func (o *BTUserSettingsInfo) SetViewManipulationMouseKeyMapping(v BTViewManipulationMouseKeyMappingInfo) {
	o.ViewManipulationMouseKeyMapping = &v
}

// GetViewMappingId returns the ViewMappingId field value if set, zero value otherwise.
func (o *BTUserSettingsInfo) GetViewMappingId() int32 {
	if o == nil || o.ViewMappingId == nil {
		var ret int32
		return ret
	}
	return *o.ViewMappingId
}

// GetViewMappingIdOk returns a tuple with the ViewMappingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSettingsInfo) GetViewMappingIdOk() (*int32, bool) {
	if o == nil || o.ViewMappingId == nil {
		return nil, false
	}
	return o.ViewMappingId, true
}

// HasViewMappingId returns a boolean if a field has been set.
func (o *BTUserSettingsInfo) HasViewMappingId() bool {
	if o != nil && o.ViewMappingId != nil {
		return true
	}

	return false
}

// SetViewMappingId gets a reference to the given int32 and assigns it to the ViewMappingId field.
func (o *BTUserSettingsInfo) SetViewMappingId(v int32) {
	o.ViewMappingId = &v
}

func (o BTUserSettingsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CommonUnits != nil {
		toSerialize["commonUnits"] = o.CommonUnits
	}
	if o.CustomColors != nil {
		toSerialize["customColors"] = o.CustomColors
	}
	if o.DefaultUnits != nil {
		toSerialize["defaultUnits"] = o.DefaultUnits
	}
	if o.DisplayAssemblyProperties != nil {
		toSerialize["displayAssemblyProperties"] = o.DisplayAssemblyProperties
	}
	if o.DrawingBackgroundId != nil {
		toSerialize["drawingBackgroundId"] = o.DrawingBackgroundId
	}
	if o.EnforceApplicationAcl != nil {
		toSerialize["enforceApplicationAcl"] = o.EnforceApplicationAcl
	}
	if o.ExportDrawingOptions != nil {
		toSerialize["exportDrawingOptions"] = o.ExportDrawingOptions
	}
	if o.ExportSolidOptions != nil {
		toSerialize["exportSolidOptions"] = o.ExportSolidOptions
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ImportOptions != nil {
		toSerialize["importOptions"] = o.ImportOptions
	}
	if o.Locale != nil {
		toSerialize["locale"] = o.Locale
	}
	if o.MaterialLibrarySettings != nil {
		toSerialize["materialLibrarySettings"] = o.MaterialLibrarySettings
	}
	if o.MiniToolbarSettings != nil {
		toSerialize["miniToolbarSettings"] = o.MiniToolbarSettings
	}
	if o.MouseActions != nil {
		toSerialize["mouseActions"] = o.MouseActions
	}
	if o.PreviousSketchFont != nil {
		toSerialize["previousSketchFont"] = o.PreviousSketchFont
	}
	if o.ReverseScrollWheelZoomDirection != nil {
		toSerialize["reverseScrollWheelZoomDirection"] = o.ReverseScrollWheelZoomDirection
	}
	if o.SelectItemViewStateInfos != nil {
		toSerialize["selectItemViewStateInfos"] = o.SelectItemViewStateInfos
	}
	if o.StartupPage != nil {
		toSerialize["startupPage"] = o.StartupPage
	}
	if o.SubstituteApprovers != nil {
		toSerialize["substituteApprovers"] = o.SubstituteApprovers
	}
	if o.UnitsDisplayPrecision != nil {
		toSerialize["unitsDisplayPrecision"] = o.UnitsDisplayPrecision
	}
	if o.UnitsMaximumDisplayPrecision != nil {
		toSerialize["unitsMaximumDisplayPrecision"] = o.UnitsMaximumDisplayPrecision
	}
	if o.Use24HourTime != nil {
		toSerialize["use24HourTime"] = o.Use24HourTime
	}
	if o.UseDecimalComma != nil {
		toSerialize["useDecimalComma"] = o.UseDecimalComma
	}
	if o.ViewManipulationMouseKeyMapping != nil {
		toSerialize["viewManipulationMouseKeyMapping"] = o.ViewManipulationMouseKeyMapping
	}
	if o.ViewMappingId != nil {
		toSerialize["viewMappingId"] = o.ViewMappingId
	}
	return json.Marshal(toSerialize)
}

type NullableBTUserSettingsInfo struct {
	value *BTUserSettingsInfo
	isSet bool
}

func (v NullableBTUserSettingsInfo) Get() *BTUserSettingsInfo {
	return v.value
}

func (v *NullableBTUserSettingsInfo) Set(val *BTUserSettingsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUserSettingsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUserSettingsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUserSettingsInfo(val *BTUserSettingsInfo) *NullableBTUserSettingsInfo {
	return &NullableBTUserSettingsInfo{value: val, isSet: true}
}

func (v NullableBTUserSettingsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUserSettingsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
