/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.6926-35d1d6168263
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTExportTessellatedFacesFace1192 struct for BTExportTessellatedFacesFace1192
type BTExportTessellatedFacesFace1192 struct {
	Appearance           *BTGraphicsAppearance1152           `json:"appearance,omitempty"`
	AppearanceSourceId   *string                             `json:"appearanceSourceId,omitempty"`
	AppearanceSourceName *string                             `json:"appearanceSourceName,omitempty"`
	BtType               *string                             `json:"btType,omitempty"`
	ErrorMessage         *string                             `json:"errorMessage,omitempty"`
	Facets               []BTExportTessellatedFacesFacet1417 `json:"facets,omitempty"`
	Id                   *string                             `json:"id,omitempty"`
}

// NewBTExportTessellatedFacesFace1192 instantiates a new BTExportTessellatedFacesFace1192 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportTessellatedFacesFace1192() *BTExportTessellatedFacesFace1192 {
	this := BTExportTessellatedFacesFace1192{}
	return &this
}

// NewBTExportTessellatedFacesFace1192WithDefaults instantiates a new BTExportTessellatedFacesFace1192 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportTessellatedFacesFace1192WithDefaults() *BTExportTessellatedFacesFace1192 {
	this := BTExportTessellatedFacesFace1192{}
	return &this
}

// GetAppearance returns the Appearance field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesFace1192) GetAppearance() BTGraphicsAppearance1152 {
	if o == nil || o.Appearance == nil {
		var ret BTGraphicsAppearance1152
		return ret
	}
	return *o.Appearance
}

// GetAppearanceOk returns a tuple with the Appearance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesFace1192) GetAppearanceOk() (*BTGraphicsAppearance1152, bool) {
	if o == nil || o.Appearance == nil {
		return nil, false
	}
	return o.Appearance, true
}

// HasAppearance returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesFace1192) HasAppearance() bool {
	if o != nil && o.Appearance != nil {
		return true
	}

	return false
}

// SetAppearance gets a reference to the given BTGraphicsAppearance1152 and assigns it to the Appearance field.
func (o *BTExportTessellatedFacesFace1192) SetAppearance(v BTGraphicsAppearance1152) {
	o.Appearance = &v
}

// GetAppearanceSourceId returns the AppearanceSourceId field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesFace1192) GetAppearanceSourceId() string {
	if o == nil || o.AppearanceSourceId == nil {
		var ret string
		return ret
	}
	return *o.AppearanceSourceId
}

// GetAppearanceSourceIdOk returns a tuple with the AppearanceSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesFace1192) GetAppearanceSourceIdOk() (*string, bool) {
	if o == nil || o.AppearanceSourceId == nil {
		return nil, false
	}
	return o.AppearanceSourceId, true
}

// HasAppearanceSourceId returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesFace1192) HasAppearanceSourceId() bool {
	if o != nil && o.AppearanceSourceId != nil {
		return true
	}

	return false
}

// SetAppearanceSourceId gets a reference to the given string and assigns it to the AppearanceSourceId field.
func (o *BTExportTessellatedFacesFace1192) SetAppearanceSourceId(v string) {
	o.AppearanceSourceId = &v
}

// GetAppearanceSourceName returns the AppearanceSourceName field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesFace1192) GetAppearanceSourceName() string {
	if o == nil || o.AppearanceSourceName == nil {
		var ret string
		return ret
	}
	return *o.AppearanceSourceName
}

// GetAppearanceSourceNameOk returns a tuple with the AppearanceSourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesFace1192) GetAppearanceSourceNameOk() (*string, bool) {
	if o == nil || o.AppearanceSourceName == nil {
		return nil, false
	}
	return o.AppearanceSourceName, true
}

// HasAppearanceSourceName returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesFace1192) HasAppearanceSourceName() bool {
	if o != nil && o.AppearanceSourceName != nil {
		return true
	}

	return false
}

// SetAppearanceSourceName gets a reference to the given string and assigns it to the AppearanceSourceName field.
func (o *BTExportTessellatedFacesFace1192) SetAppearanceSourceName(v string) {
	o.AppearanceSourceName = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesFace1192) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesFace1192) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesFace1192) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTExportTessellatedFacesFace1192) SetBtType(v string) {
	o.BtType = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesFace1192) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesFace1192) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesFace1192) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *BTExportTessellatedFacesFace1192) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetFacets returns the Facets field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesFace1192) GetFacets() []BTExportTessellatedFacesFacet1417 {
	if o == nil || o.Facets == nil {
		var ret []BTExportTessellatedFacesFacet1417
		return ret
	}
	return o.Facets
}

// GetFacetsOk returns a tuple with the Facets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesFace1192) GetFacetsOk() ([]BTExportTessellatedFacesFacet1417, bool) {
	if o == nil || o.Facets == nil {
		return nil, false
	}
	return o.Facets, true
}

// HasFacets returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesFace1192) HasFacets() bool {
	if o != nil && o.Facets != nil {
		return true
	}

	return false
}

// SetFacets gets a reference to the given []BTExportTessellatedFacesFacet1417 and assigns it to the Facets field.
func (o *BTExportTessellatedFacesFace1192) SetFacets(v []BTExportTessellatedFacesFacet1417) {
	o.Facets = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesFace1192) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesFace1192) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesFace1192) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTExportTessellatedFacesFace1192) SetId(v string) {
	o.Id = &v
}

func (o BTExportTessellatedFacesFace1192) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Appearance != nil {
		toSerialize["appearance"] = o.Appearance
	}
	if o.AppearanceSourceId != nil {
		toSerialize["appearanceSourceId"] = o.AppearanceSourceId
	}
	if o.AppearanceSourceName != nil {
		toSerialize["appearanceSourceName"] = o.AppearanceSourceName
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.Facets != nil {
		toSerialize["facets"] = o.Facets
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableBTExportTessellatedFacesFace1192 struct {
	value *BTExportTessellatedFacesFace1192
	isSet bool
}

func (v NullableBTExportTessellatedFacesFace1192) Get() *BTExportTessellatedFacesFace1192 {
	return v.value
}

func (v *NullableBTExportTessellatedFacesFace1192) Set(val *BTExportTessellatedFacesFace1192) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportTessellatedFacesFace1192) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportTessellatedFacesFace1192) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportTessellatedFacesFace1192(val *BTExportTessellatedFacesFace1192) *NullableBTExportTessellatedFacesFace1192 {
	return &NullableBTExportTessellatedFacesFace1192{value: val, isSet: true}
}

func (v NullableBTExportTessellatedFacesFace1192) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportTessellatedFacesFace1192) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
