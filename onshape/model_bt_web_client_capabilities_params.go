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

// BTWebClientCapabilitiesParams struct for BTWebClientCapabilitiesParams
type BTWebClientCapabilitiesParams struct {
	AngleInstancedArrays        *bool    `json:"angleInstancedArrays,omitempty"`
	CompressedTextureS3tc       *bool    `json:"compressedTextureS3tc,omitempty"`
	DepthTexture                *bool    `json:"depthTexture,omitempty"`
	DevicePixelRatio            *float64 `json:"devicePixelRatio,omitempty"`
	DrawBuffers                 *bool    `json:"drawBuffers,omitempty"`
	ExtTextureFilterAnisotropic *bool    `json:"extTextureFilterAnisotropic,omitempty"`
	Has3dMouse                  *bool    `json:"has3dMouse,omitempty"`
	OesElementIndexUint         *bool    `json:"oesElementIndexUint,omitempty"`
	OesStandardDerivatives      *bool    `json:"oesStandardDerivatives,omitempty"`
	OesTextureFloat             *bool    `json:"oesTextureFloat,omitempty"`
	OesTextureFloatLinear       *bool    `json:"oesTextureFloatLinear,omitempty"`
	OesTextureHalfFloat         *bool    `json:"oesTextureHalfFloat,omitempty"`
	OesTextureHalfFloatLinear   *bool    `json:"oesTextureHalfFloatLinear,omitempty"`
	OesVertexArrayObject        *bool    `json:"oesVertexArrayObject,omitempty"`
	Renderer                    *string  `json:"renderer,omitempty"`
	ScreenHeight                *int32   `json:"screenHeight,omitempty"`
	ScreenWidth                 *int32   `json:"screenWidth,omitempty"`
	Vendor                      *string  `json:"vendor,omitempty"`
}

// NewBTWebClientCapabilitiesParams instantiates a new BTWebClientCapabilitiesParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTWebClientCapabilitiesParams() *BTWebClientCapabilitiesParams {
	this := BTWebClientCapabilitiesParams{}
	return &this
}

// NewBTWebClientCapabilitiesParamsWithDefaults instantiates a new BTWebClientCapabilitiesParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTWebClientCapabilitiesParamsWithDefaults() *BTWebClientCapabilitiesParams {
	this := BTWebClientCapabilitiesParams{}
	return &this
}

// GetAngleInstancedArrays returns the AngleInstancedArrays field value if set, zero value otherwise.
func (o *BTWebClientCapabilitiesParams) GetAngleInstancedArrays() bool {
	if o == nil || o.AngleInstancedArrays == nil {
		var ret bool
		return ret
	}
	return *o.AngleInstancedArrays
}

// GetAngleInstancedArraysOk returns a tuple with the AngleInstancedArrays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebClientCapabilitiesParams) GetAngleInstancedArraysOk() (*bool, bool) {
	if o == nil || o.AngleInstancedArrays == nil {
		return nil, false
	}
	return o.AngleInstancedArrays, true
}

// HasAngleInstancedArrays returns a boolean if a field has been set.
func (o *BTWebClientCapabilitiesParams) HasAngleInstancedArrays() bool {
	if o != nil && o.AngleInstancedArrays != nil {
		return true
	}

	return false
}

// SetAngleInstancedArrays gets a reference to the given bool and assigns it to the AngleInstancedArrays field.
func (o *BTWebClientCapabilitiesParams) SetAngleInstancedArrays(v bool) {
	o.AngleInstancedArrays = &v
}

// GetCompressedTextureS3tc returns the CompressedTextureS3tc field value if set, zero value otherwise.
func (o *BTWebClientCapabilitiesParams) GetCompressedTextureS3tc() bool {
	if o == nil || o.CompressedTextureS3tc == nil {
		var ret bool
		return ret
	}
	return *o.CompressedTextureS3tc
}

// GetCompressedTextureS3tcOk returns a tuple with the CompressedTextureS3tc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebClientCapabilitiesParams) GetCompressedTextureS3tcOk() (*bool, bool) {
	if o == nil || o.CompressedTextureS3tc == nil {
		return nil, false
	}
	return o.CompressedTextureS3tc, true
}

// HasCompressedTextureS3tc returns a boolean if a field has been set.
func (o *BTWebClientCapabilitiesParams) HasCompressedTextureS3tc() bool {
	if o != nil && o.CompressedTextureS3tc != nil {
		return true
	}

	return false
}

// SetCompressedTextureS3tc gets a reference to the given bool and assigns it to the CompressedTextureS3tc field.
func (o *BTWebClientCapabilitiesParams) SetCompressedTextureS3tc(v bool) {
	o.CompressedTextureS3tc = &v
}

// GetDepthTexture returns the DepthTexture field value if set, zero value otherwise.
func (o *BTWebClientCapabilitiesParams) GetDepthTexture() bool {
	if o == nil || o.DepthTexture == nil {
		var ret bool
		return ret
	}
	return *o.DepthTexture
}

// GetDepthTextureOk returns a tuple with the DepthTexture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebClientCapabilitiesParams) GetDepthTextureOk() (*bool, bool) {
	if o == nil || o.DepthTexture == nil {
		return nil, false
	}
	return o.DepthTexture, true
}

// HasDepthTexture returns a boolean if a field has been set.
func (o *BTWebClientCapabilitiesParams) HasDepthTexture() bool {
	if o != nil && o.DepthTexture != nil {
		return true
	}

	return false
}

// SetDepthTexture gets a reference to the given bool and assigns it to the DepthTexture field.
func (o *BTWebClientCapabilitiesParams) SetDepthTexture(v bool) {
	o.DepthTexture = &v
}

// GetDevicePixelRatio returns the DevicePixelRatio field value if set, zero value otherwise.
func (o *BTWebClientCapabilitiesParams) GetDevicePixelRatio() float64 {
	if o == nil || o.DevicePixelRatio == nil {
		var ret float64
		return ret
	}
	return *o.DevicePixelRatio
}

// GetDevicePixelRatioOk returns a tuple with the DevicePixelRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebClientCapabilitiesParams) GetDevicePixelRatioOk() (*float64, bool) {
	if o == nil || o.DevicePixelRatio == nil {
		return nil, false
	}
	return o.DevicePixelRatio, true
}

// HasDevicePixelRatio returns a boolean if a field has been set.
func (o *BTWebClientCapabilitiesParams) HasDevicePixelRatio() bool {
	if o != nil && o.DevicePixelRatio != nil {
		return true
	}

	return false
}

// SetDevicePixelRatio gets a reference to the given float64 and assigns it to the DevicePixelRatio field.
func (o *BTWebClientCapabilitiesParams) SetDevicePixelRatio(v float64) {
	o.DevicePixelRatio = &v
}

// GetDrawBuffers returns the DrawBuffers field value if set, zero value otherwise.
func (o *BTWebClientCapabilitiesParams) GetDrawBuffers() bool {
	if o == nil || o.DrawBuffers == nil {
		var ret bool
		return ret
	}
	return *o.DrawBuffers
}

// GetDrawBuffersOk returns a tuple with the DrawBuffers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebClientCapabilitiesParams) GetDrawBuffersOk() (*bool, bool) {
	if o == nil || o.DrawBuffers == nil {
		return nil, false
	}
	return o.DrawBuffers, true
}

// HasDrawBuffers returns a boolean if a field has been set.
func (o *BTWebClientCapabilitiesParams) HasDrawBuffers() bool {
	if o != nil && o.DrawBuffers != nil {
		return true
	}

	return false
}

// SetDrawBuffers gets a reference to the given bool and assigns it to the DrawBuffers field.
func (o *BTWebClientCapabilitiesParams) SetDrawBuffers(v bool) {
	o.DrawBuffers = &v
}

// GetExtTextureFilterAnisotropic returns the ExtTextureFilterAnisotropic field value if set, zero value otherwise.
func (o *BTWebClientCapabilitiesParams) GetExtTextureFilterAnisotropic() bool {
	if o == nil || o.ExtTextureFilterAnisotropic == nil {
		var ret bool
		return ret
	}
	return *o.ExtTextureFilterAnisotropic
}

// GetExtTextureFilterAnisotropicOk returns a tuple with the ExtTextureFilterAnisotropic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebClientCapabilitiesParams) GetExtTextureFilterAnisotropicOk() (*bool, bool) {
	if o == nil || o.ExtTextureFilterAnisotropic == nil {
		return nil, false
	}
	return o.ExtTextureFilterAnisotropic, true
}

// HasExtTextureFilterAnisotropic returns a boolean if a field has been set.
func (o *BTWebClientCapabilitiesParams) HasExtTextureFilterAnisotropic() bool {
	if o != nil && o.ExtTextureFilterAnisotropic != nil {
		return true
	}

	return false
}

// SetExtTextureFilterAnisotropic gets a reference to the given bool and assigns it to the ExtTextureFilterAnisotropic field.
func (o *BTWebClientCapabilitiesParams) SetExtTextureFilterAnisotropic(v bool) {
	o.ExtTextureFilterAnisotropic = &v
}

// GetHas3dMouse returns the Has3dMouse field value if set, zero value otherwise.
func (o *BTWebClientCapabilitiesParams) GetHas3dMouse() bool {
	if o == nil || o.Has3dMouse == nil {
		var ret bool
		return ret
	}
	return *o.Has3dMouse
}

// GetHas3dMouseOk returns a tuple with the Has3dMouse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebClientCapabilitiesParams) GetHas3dMouseOk() (*bool, bool) {
	if o == nil || o.Has3dMouse == nil {
		return nil, false
	}
	return o.Has3dMouse, true
}

// HasHas3dMouse returns a boolean if a field has been set.
func (o *BTWebClientCapabilitiesParams) HasHas3dMouse() bool {
	if o != nil && o.Has3dMouse != nil {
		return true
	}

	return false
}

// SetHas3dMouse gets a reference to the given bool and assigns it to the Has3dMouse field.
func (o *BTWebClientCapabilitiesParams) SetHas3dMouse(v bool) {
	o.Has3dMouse = &v
}

// GetOesElementIndexUint returns the OesElementIndexUint field value if set, zero value otherwise.
func (o *BTWebClientCapabilitiesParams) GetOesElementIndexUint() bool {
	if o == nil || o.OesElementIndexUint == nil {
		var ret bool
		return ret
	}
	return *o.OesElementIndexUint
}

// GetOesElementIndexUintOk returns a tuple with the OesElementIndexUint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebClientCapabilitiesParams) GetOesElementIndexUintOk() (*bool, bool) {
	if o == nil || o.OesElementIndexUint == nil {
		return nil, false
	}
	return o.OesElementIndexUint, true
}

// HasOesElementIndexUint returns a boolean if a field has been set.
func (o *BTWebClientCapabilitiesParams) HasOesElementIndexUint() bool {
	if o != nil && o.OesElementIndexUint != nil {
		return true
	}

	return false
}

// SetOesElementIndexUint gets a reference to the given bool and assigns it to the OesElementIndexUint field.
func (o *BTWebClientCapabilitiesParams) SetOesElementIndexUint(v bool) {
	o.OesElementIndexUint = &v
}

// GetOesStandardDerivatives returns the OesStandardDerivatives field value if set, zero value otherwise.
func (o *BTWebClientCapabilitiesParams) GetOesStandardDerivatives() bool {
	if o == nil || o.OesStandardDerivatives == nil {
		var ret bool
		return ret
	}
	return *o.OesStandardDerivatives
}

// GetOesStandardDerivativesOk returns a tuple with the OesStandardDerivatives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebClientCapabilitiesParams) GetOesStandardDerivativesOk() (*bool, bool) {
	if o == nil || o.OesStandardDerivatives == nil {
		return nil, false
	}
	return o.OesStandardDerivatives, true
}

// HasOesStandardDerivatives returns a boolean if a field has been set.
func (o *BTWebClientCapabilitiesParams) HasOesStandardDerivatives() bool {
	if o != nil && o.OesStandardDerivatives != nil {
		return true
	}

	return false
}

// SetOesStandardDerivatives gets a reference to the given bool and assigns it to the OesStandardDerivatives field.
func (o *BTWebClientCapabilitiesParams) SetOesStandardDerivatives(v bool) {
	o.OesStandardDerivatives = &v
}

// GetOesTextureFloat returns the OesTextureFloat field value if set, zero value otherwise.
func (o *BTWebClientCapabilitiesParams) GetOesTextureFloat() bool {
	if o == nil || o.OesTextureFloat == nil {
		var ret bool
		return ret
	}
	return *o.OesTextureFloat
}

// GetOesTextureFloatOk returns a tuple with the OesTextureFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebClientCapabilitiesParams) GetOesTextureFloatOk() (*bool, bool) {
	if o == nil || o.OesTextureFloat == nil {
		return nil, false
	}
	return o.OesTextureFloat, true
}

// HasOesTextureFloat returns a boolean if a field has been set.
func (o *BTWebClientCapabilitiesParams) HasOesTextureFloat() bool {
	if o != nil && o.OesTextureFloat != nil {
		return true
	}

	return false
}

// SetOesTextureFloat gets a reference to the given bool and assigns it to the OesTextureFloat field.
func (o *BTWebClientCapabilitiesParams) SetOesTextureFloat(v bool) {
	o.OesTextureFloat = &v
}

// GetOesTextureFloatLinear returns the OesTextureFloatLinear field value if set, zero value otherwise.
func (o *BTWebClientCapabilitiesParams) GetOesTextureFloatLinear() bool {
	if o == nil || o.OesTextureFloatLinear == nil {
		var ret bool
		return ret
	}
	return *o.OesTextureFloatLinear
}

// GetOesTextureFloatLinearOk returns a tuple with the OesTextureFloatLinear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebClientCapabilitiesParams) GetOesTextureFloatLinearOk() (*bool, bool) {
	if o == nil || o.OesTextureFloatLinear == nil {
		return nil, false
	}
	return o.OesTextureFloatLinear, true
}

// HasOesTextureFloatLinear returns a boolean if a field has been set.
func (o *BTWebClientCapabilitiesParams) HasOesTextureFloatLinear() bool {
	if o != nil && o.OesTextureFloatLinear != nil {
		return true
	}

	return false
}

// SetOesTextureFloatLinear gets a reference to the given bool and assigns it to the OesTextureFloatLinear field.
func (o *BTWebClientCapabilitiesParams) SetOesTextureFloatLinear(v bool) {
	o.OesTextureFloatLinear = &v
}

// GetOesTextureHalfFloat returns the OesTextureHalfFloat field value if set, zero value otherwise.
func (o *BTWebClientCapabilitiesParams) GetOesTextureHalfFloat() bool {
	if o == nil || o.OesTextureHalfFloat == nil {
		var ret bool
		return ret
	}
	return *o.OesTextureHalfFloat
}

// GetOesTextureHalfFloatOk returns a tuple with the OesTextureHalfFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebClientCapabilitiesParams) GetOesTextureHalfFloatOk() (*bool, bool) {
	if o == nil || o.OesTextureHalfFloat == nil {
		return nil, false
	}
	return o.OesTextureHalfFloat, true
}

// HasOesTextureHalfFloat returns a boolean if a field has been set.
func (o *BTWebClientCapabilitiesParams) HasOesTextureHalfFloat() bool {
	if o != nil && o.OesTextureHalfFloat != nil {
		return true
	}

	return false
}

// SetOesTextureHalfFloat gets a reference to the given bool and assigns it to the OesTextureHalfFloat field.
func (o *BTWebClientCapabilitiesParams) SetOesTextureHalfFloat(v bool) {
	o.OesTextureHalfFloat = &v
}

// GetOesTextureHalfFloatLinear returns the OesTextureHalfFloatLinear field value if set, zero value otherwise.
func (o *BTWebClientCapabilitiesParams) GetOesTextureHalfFloatLinear() bool {
	if o == nil || o.OesTextureHalfFloatLinear == nil {
		var ret bool
		return ret
	}
	return *o.OesTextureHalfFloatLinear
}

// GetOesTextureHalfFloatLinearOk returns a tuple with the OesTextureHalfFloatLinear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebClientCapabilitiesParams) GetOesTextureHalfFloatLinearOk() (*bool, bool) {
	if o == nil || o.OesTextureHalfFloatLinear == nil {
		return nil, false
	}
	return o.OesTextureHalfFloatLinear, true
}

// HasOesTextureHalfFloatLinear returns a boolean if a field has been set.
func (o *BTWebClientCapabilitiesParams) HasOesTextureHalfFloatLinear() bool {
	if o != nil && o.OesTextureHalfFloatLinear != nil {
		return true
	}

	return false
}

// SetOesTextureHalfFloatLinear gets a reference to the given bool and assigns it to the OesTextureHalfFloatLinear field.
func (o *BTWebClientCapabilitiesParams) SetOesTextureHalfFloatLinear(v bool) {
	o.OesTextureHalfFloatLinear = &v
}

// GetOesVertexArrayObject returns the OesVertexArrayObject field value if set, zero value otherwise.
func (o *BTWebClientCapabilitiesParams) GetOesVertexArrayObject() bool {
	if o == nil || o.OesVertexArrayObject == nil {
		var ret bool
		return ret
	}
	return *o.OesVertexArrayObject
}

// GetOesVertexArrayObjectOk returns a tuple with the OesVertexArrayObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebClientCapabilitiesParams) GetOesVertexArrayObjectOk() (*bool, bool) {
	if o == nil || o.OesVertexArrayObject == nil {
		return nil, false
	}
	return o.OesVertexArrayObject, true
}

// HasOesVertexArrayObject returns a boolean if a field has been set.
func (o *BTWebClientCapabilitiesParams) HasOesVertexArrayObject() bool {
	if o != nil && o.OesVertexArrayObject != nil {
		return true
	}

	return false
}

// SetOesVertexArrayObject gets a reference to the given bool and assigns it to the OesVertexArrayObject field.
func (o *BTWebClientCapabilitiesParams) SetOesVertexArrayObject(v bool) {
	o.OesVertexArrayObject = &v
}

// GetRenderer returns the Renderer field value if set, zero value otherwise.
func (o *BTWebClientCapabilitiesParams) GetRenderer() string {
	if o == nil || o.Renderer == nil {
		var ret string
		return ret
	}
	return *o.Renderer
}

// GetRendererOk returns a tuple with the Renderer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebClientCapabilitiesParams) GetRendererOk() (*string, bool) {
	if o == nil || o.Renderer == nil {
		return nil, false
	}
	return o.Renderer, true
}

// HasRenderer returns a boolean if a field has been set.
func (o *BTWebClientCapabilitiesParams) HasRenderer() bool {
	if o != nil && o.Renderer != nil {
		return true
	}

	return false
}

// SetRenderer gets a reference to the given string and assigns it to the Renderer field.
func (o *BTWebClientCapabilitiesParams) SetRenderer(v string) {
	o.Renderer = &v
}

// GetScreenHeight returns the ScreenHeight field value if set, zero value otherwise.
func (o *BTWebClientCapabilitiesParams) GetScreenHeight() int32 {
	if o == nil || o.ScreenHeight == nil {
		var ret int32
		return ret
	}
	return *o.ScreenHeight
}

// GetScreenHeightOk returns a tuple with the ScreenHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebClientCapabilitiesParams) GetScreenHeightOk() (*int32, bool) {
	if o == nil || o.ScreenHeight == nil {
		return nil, false
	}
	return o.ScreenHeight, true
}

// HasScreenHeight returns a boolean if a field has been set.
func (o *BTWebClientCapabilitiesParams) HasScreenHeight() bool {
	if o != nil && o.ScreenHeight != nil {
		return true
	}

	return false
}

// SetScreenHeight gets a reference to the given int32 and assigns it to the ScreenHeight field.
func (o *BTWebClientCapabilitiesParams) SetScreenHeight(v int32) {
	o.ScreenHeight = &v
}

// GetScreenWidth returns the ScreenWidth field value if set, zero value otherwise.
func (o *BTWebClientCapabilitiesParams) GetScreenWidth() int32 {
	if o == nil || o.ScreenWidth == nil {
		var ret int32
		return ret
	}
	return *o.ScreenWidth
}

// GetScreenWidthOk returns a tuple with the ScreenWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebClientCapabilitiesParams) GetScreenWidthOk() (*int32, bool) {
	if o == nil || o.ScreenWidth == nil {
		return nil, false
	}
	return o.ScreenWidth, true
}

// HasScreenWidth returns a boolean if a field has been set.
func (o *BTWebClientCapabilitiesParams) HasScreenWidth() bool {
	if o != nil && o.ScreenWidth != nil {
		return true
	}

	return false
}

// SetScreenWidth gets a reference to the given int32 and assigns it to the ScreenWidth field.
func (o *BTWebClientCapabilitiesParams) SetScreenWidth(v int32) {
	o.ScreenWidth = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *BTWebClientCapabilitiesParams) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebClientCapabilitiesParams) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *BTWebClientCapabilitiesParams) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *BTWebClientCapabilitiesParams) SetVendor(v string) {
	o.Vendor = &v
}

func (o BTWebClientCapabilitiesParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AngleInstancedArrays != nil {
		toSerialize["angleInstancedArrays"] = o.AngleInstancedArrays
	}
	if o.CompressedTextureS3tc != nil {
		toSerialize["compressedTextureS3tc"] = o.CompressedTextureS3tc
	}
	if o.DepthTexture != nil {
		toSerialize["depthTexture"] = o.DepthTexture
	}
	if o.DevicePixelRatio != nil {
		toSerialize["devicePixelRatio"] = o.DevicePixelRatio
	}
	if o.DrawBuffers != nil {
		toSerialize["drawBuffers"] = o.DrawBuffers
	}
	if o.ExtTextureFilterAnisotropic != nil {
		toSerialize["extTextureFilterAnisotropic"] = o.ExtTextureFilterAnisotropic
	}
	if o.Has3dMouse != nil {
		toSerialize["has3dMouse"] = o.Has3dMouse
	}
	if o.OesElementIndexUint != nil {
		toSerialize["oesElementIndexUint"] = o.OesElementIndexUint
	}
	if o.OesStandardDerivatives != nil {
		toSerialize["oesStandardDerivatives"] = o.OesStandardDerivatives
	}
	if o.OesTextureFloat != nil {
		toSerialize["oesTextureFloat"] = o.OesTextureFloat
	}
	if o.OesTextureFloatLinear != nil {
		toSerialize["oesTextureFloatLinear"] = o.OesTextureFloatLinear
	}
	if o.OesTextureHalfFloat != nil {
		toSerialize["oesTextureHalfFloat"] = o.OesTextureHalfFloat
	}
	if o.OesTextureHalfFloatLinear != nil {
		toSerialize["oesTextureHalfFloatLinear"] = o.OesTextureHalfFloatLinear
	}
	if o.OesVertexArrayObject != nil {
		toSerialize["oesVertexArrayObject"] = o.OesVertexArrayObject
	}
	if o.Renderer != nil {
		toSerialize["renderer"] = o.Renderer
	}
	if o.ScreenHeight != nil {
		toSerialize["screenHeight"] = o.ScreenHeight
	}
	if o.ScreenWidth != nil {
		toSerialize["screenWidth"] = o.ScreenWidth
	}
	if o.Vendor != nil {
		toSerialize["vendor"] = o.Vendor
	}
	return json.Marshal(toSerialize)
}

type NullableBTWebClientCapabilitiesParams struct {
	value *BTWebClientCapabilitiesParams
	isSet bool
}

func (v NullableBTWebClientCapabilitiesParams) Get() *BTWebClientCapabilitiesParams {
	return v.value
}

func (v *NullableBTWebClientCapabilitiesParams) Set(val *BTWebClientCapabilitiesParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWebClientCapabilitiesParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWebClientCapabilitiesParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWebClientCapabilitiesParams(val *BTWebClientCapabilitiesParams) *NullableBTWebClientCapabilitiesParams {
	return &NullableBTWebClientCapabilitiesParams{value: val, isSet: true}
}

func (v NullableBTWebClientCapabilitiesParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWebClientCapabilitiesParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
