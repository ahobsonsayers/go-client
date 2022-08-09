/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5944-34bf93ccfb05
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTExportModelCoedge1342 struct for BTExportModelCoedge1342
type BTExportModelCoedge1342 struct {
	EdgeId      *string `json:"edgeId,omitempty"`
	Orientation *bool   `json:"orientation,omitempty"`
}

// NewBTExportModelCoedge1342 instantiates a new BTExportModelCoedge1342 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportModelCoedge1342() *BTExportModelCoedge1342 {
	this := BTExportModelCoedge1342{}
	return &this
}

// NewBTExportModelCoedge1342WithDefaults instantiates a new BTExportModelCoedge1342 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportModelCoedge1342WithDefaults() *BTExportModelCoedge1342 {
	this := BTExportModelCoedge1342{}
	return &this
}

// GetEdgeId returns the EdgeId field value if set, zero value otherwise.
func (o *BTExportModelCoedge1342) GetEdgeId() string {
	if o == nil || o.EdgeId == nil {
		var ret string
		return ret
	}
	return *o.EdgeId
}

// GetEdgeIdOk returns a tuple with the EdgeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelCoedge1342) GetEdgeIdOk() (*string, bool) {
	if o == nil || o.EdgeId == nil {
		return nil, false
	}
	return o.EdgeId, true
}

// HasEdgeId returns a boolean if a field has been set.
func (o *BTExportModelCoedge1342) HasEdgeId() bool {
	if o != nil && o.EdgeId != nil {
		return true
	}

	return false
}

// SetEdgeId gets a reference to the given string and assigns it to the EdgeId field.
func (o *BTExportModelCoedge1342) SetEdgeId(v string) {
	o.EdgeId = &v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *BTExportModelCoedge1342) GetOrientation() bool {
	if o == nil || o.Orientation == nil {
		var ret bool
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelCoedge1342) GetOrientationOk() (*bool, bool) {
	if o == nil || o.Orientation == nil {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *BTExportModelCoedge1342) HasOrientation() bool {
	if o != nil && o.Orientation != nil {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given bool and assigns it to the Orientation field.
func (o *BTExportModelCoedge1342) SetOrientation(v bool) {
	o.Orientation = &v
}

func (o BTExportModelCoedge1342) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EdgeId != nil {
		toSerialize["edgeId"] = o.EdgeId
	}
	if o.Orientation != nil {
		toSerialize["orientation"] = o.Orientation
	}
	return json.Marshal(toSerialize)
}

type NullableBTExportModelCoedge1342 struct {
	value *BTExportModelCoedge1342
	isSet bool
}

func (v NullableBTExportModelCoedge1342) Get() *BTExportModelCoedge1342 {
	return v.value
}

func (v *NullableBTExportModelCoedge1342) Set(val *BTExportModelCoedge1342) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportModelCoedge1342) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportModelCoedge1342) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportModelCoedge1342(val *BTExportModelCoedge1342) *NullableBTExportModelCoedge1342 {
	return &NullableBTExportModelCoedge1342{value: val, isSet: true}
}

func (v NullableBTExportModelCoedge1342) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportModelCoedge1342) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
