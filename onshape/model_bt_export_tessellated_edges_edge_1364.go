/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13200-ff216a970a02
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTExportTessellatedEdgesEdge1364 struct for BTExportTessellatedEdgesEdge1364
type BTExportTessellatedEdgesEdge1364 struct {
	BtType   *string         `json:"btType,omitempty"`
	Id       *string         `json:"id,omitempty"`
	Vertices []BTVector3d389 `json:"vertices,omitempty"`
}

// NewBTExportTessellatedEdgesEdge1364 instantiates a new BTExportTessellatedEdgesEdge1364 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportTessellatedEdgesEdge1364() *BTExportTessellatedEdgesEdge1364 {
	this := BTExportTessellatedEdgesEdge1364{}
	return &this
}

// NewBTExportTessellatedEdgesEdge1364WithDefaults instantiates a new BTExportTessellatedEdgesEdge1364 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportTessellatedEdgesEdge1364WithDefaults() *BTExportTessellatedEdgesEdge1364 {
	this := BTExportTessellatedEdgesEdge1364{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTExportTessellatedEdgesEdge1364) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedEdgesEdge1364) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTExportTessellatedEdgesEdge1364) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTExportTessellatedEdgesEdge1364) SetBtType(v string) {
	o.BtType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTExportTessellatedEdgesEdge1364) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedEdgesEdge1364) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTExportTessellatedEdgesEdge1364) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTExportTessellatedEdgesEdge1364) SetId(v string) {
	o.Id = &v
}

// GetVertices returns the Vertices field value if set, zero value otherwise.
func (o *BTExportTessellatedEdgesEdge1364) GetVertices() []BTVector3d389 {
	if o == nil || o.Vertices == nil {
		var ret []BTVector3d389
		return ret
	}
	return o.Vertices
}

// GetVerticesOk returns a tuple with the Vertices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedEdgesEdge1364) GetVerticesOk() ([]BTVector3d389, bool) {
	if o == nil || o.Vertices == nil {
		return nil, false
	}
	return o.Vertices, true
}

// HasVertices returns a boolean if a field has been set.
func (o *BTExportTessellatedEdgesEdge1364) HasVertices() bool {
	if o != nil && o.Vertices != nil {
		return true
	}

	return false
}

// SetVertices gets a reference to the given []BTVector3d389 and assigns it to the Vertices field.
func (o *BTExportTessellatedEdgesEdge1364) SetVertices(v []BTVector3d389) {
	o.Vertices = v
}

func (o BTExportTessellatedEdgesEdge1364) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Vertices != nil {
		toSerialize["vertices"] = o.Vertices
	}
	return json.Marshal(toSerialize)
}

type NullableBTExportTessellatedEdgesEdge1364 struct {
	value *BTExportTessellatedEdgesEdge1364
	isSet bool
}

func (v NullableBTExportTessellatedEdgesEdge1364) Get() *BTExportTessellatedEdgesEdge1364 {
	return v.value
}

func (v *NullableBTExportTessellatedEdgesEdge1364) Set(val *BTExportTessellatedEdgesEdge1364) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportTessellatedEdgesEdge1364) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportTessellatedEdgesEdge1364) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportTessellatedEdgesEdge1364(val *BTExportTessellatedEdgesEdge1364) *NullableBTExportTessellatedEdgesEdge1364 {
	return &NullableBTExportTessellatedEdgesEdge1364{value: val, isSet: true}
}

func (v NullableBTExportTessellatedEdgesEdge1364) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportTessellatedEdgesEdge1364) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
