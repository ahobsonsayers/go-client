/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13685-0fb99d06bde5
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTGraphicsSectionPlaneData1429 struct for BTGraphicsSectionPlaneData1429
type BTGraphicsSectionPlaneData1429 struct {
	Center  *BTVector3d389 `json:"center,omitempty"`
	Normal  *BTVector3d389 `json:"normal,omitempty"`
	Tangent *BTVector3d389 `json:"tangent,omitempty"`
}

// NewBTGraphicsSectionPlaneData1429 instantiates a new BTGraphicsSectionPlaneData1429 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTGraphicsSectionPlaneData1429() *BTGraphicsSectionPlaneData1429 {
	this := BTGraphicsSectionPlaneData1429{}
	return &this
}

// NewBTGraphicsSectionPlaneData1429WithDefaults instantiates a new BTGraphicsSectionPlaneData1429 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTGraphicsSectionPlaneData1429WithDefaults() *BTGraphicsSectionPlaneData1429 {
	this := BTGraphicsSectionPlaneData1429{}
	return &this
}

// GetCenter returns the Center field value if set, zero value otherwise.
func (o *BTGraphicsSectionPlaneData1429) GetCenter() BTVector3d389 {
	if o == nil || o.Center == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Center
}

// GetCenterOk returns a tuple with the Center field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGraphicsSectionPlaneData1429) GetCenterOk() (*BTVector3d389, bool) {
	if o == nil || o.Center == nil {
		return nil, false
	}
	return o.Center, true
}

// HasCenter returns a boolean if a field has been set.
func (o *BTGraphicsSectionPlaneData1429) HasCenter() bool {
	if o != nil && o.Center != nil {
		return true
	}

	return false
}

// SetCenter gets a reference to the given BTVector3d389 and assigns it to the Center field.
func (o *BTGraphicsSectionPlaneData1429) SetCenter(v BTVector3d389) {
	o.Center = &v
}

// GetNormal returns the Normal field value if set, zero value otherwise.
func (o *BTGraphicsSectionPlaneData1429) GetNormal() BTVector3d389 {
	if o == nil || o.Normal == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Normal
}

// GetNormalOk returns a tuple with the Normal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGraphicsSectionPlaneData1429) GetNormalOk() (*BTVector3d389, bool) {
	if o == nil || o.Normal == nil {
		return nil, false
	}
	return o.Normal, true
}

// HasNormal returns a boolean if a field has been set.
func (o *BTGraphicsSectionPlaneData1429) HasNormal() bool {
	if o != nil && o.Normal != nil {
		return true
	}

	return false
}

// SetNormal gets a reference to the given BTVector3d389 and assigns it to the Normal field.
func (o *BTGraphicsSectionPlaneData1429) SetNormal(v BTVector3d389) {
	o.Normal = &v
}

// GetTangent returns the Tangent field value if set, zero value otherwise.
func (o *BTGraphicsSectionPlaneData1429) GetTangent() BTVector3d389 {
	if o == nil || o.Tangent == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Tangent
}

// GetTangentOk returns a tuple with the Tangent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGraphicsSectionPlaneData1429) GetTangentOk() (*BTVector3d389, bool) {
	if o == nil || o.Tangent == nil {
		return nil, false
	}
	return o.Tangent, true
}

// HasTangent returns a boolean if a field has been set.
func (o *BTGraphicsSectionPlaneData1429) HasTangent() bool {
	if o != nil && o.Tangent != nil {
		return true
	}

	return false
}

// SetTangent gets a reference to the given BTVector3d389 and assigns it to the Tangent field.
func (o *BTGraphicsSectionPlaneData1429) SetTangent(v BTVector3d389) {
	o.Tangent = &v
}

func (o BTGraphicsSectionPlaneData1429) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Center != nil {
		toSerialize["center"] = o.Center
	}
	if o.Normal != nil {
		toSerialize["normal"] = o.Normal
	}
	if o.Tangent != nil {
		toSerialize["tangent"] = o.Tangent
	}
	return json.Marshal(toSerialize)
}

type NullableBTGraphicsSectionPlaneData1429 struct {
	value *BTGraphicsSectionPlaneData1429
	isSet bool
}

func (v NullableBTGraphicsSectionPlaneData1429) Get() *BTGraphicsSectionPlaneData1429 {
	return v.value
}

func (v *NullableBTGraphicsSectionPlaneData1429) Set(val *BTGraphicsSectionPlaneData1429) {
	v.value = val
	v.isSet = true
}

func (v NullableBTGraphicsSectionPlaneData1429) IsSet() bool {
	return v.isSet
}

func (v *NullableBTGraphicsSectionPlaneData1429) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTGraphicsSectionPlaneData1429(val *BTGraphicsSectionPlaneData1429) *NullableBTGraphicsSectionPlaneData1429 {
	return &NullableBTGraphicsSectionPlaneData1429{value: val, isSet: true}
}

func (v NullableBTGraphicsSectionPlaneData1429) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTGraphicsSectionPlaneData1429) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
