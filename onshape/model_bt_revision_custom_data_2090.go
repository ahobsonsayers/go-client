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

// BTRevisionCustomData2090 struct for BTRevisionCustomData2090
type BTRevisionCustomData2090 struct {
	BtType                 *string `json:"btType,omitempty"`
	PartNumber             *string `json:"partNumber,omitempty"`
	Revision               *string `json:"revision,omitempty"`
	ValidRevisionReference *bool   `json:"validRevisionReference,omitempty"`
}

// NewBTRevisionCustomData2090 instantiates a new BTRevisionCustomData2090 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTRevisionCustomData2090() *BTRevisionCustomData2090 {
	this := BTRevisionCustomData2090{}
	return &this
}

// NewBTRevisionCustomData2090WithDefaults instantiates a new BTRevisionCustomData2090 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTRevisionCustomData2090WithDefaults() *BTRevisionCustomData2090 {
	this := BTRevisionCustomData2090{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTRevisionCustomData2090) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionCustomData2090) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTRevisionCustomData2090) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTRevisionCustomData2090) SetBtType(v string) {
	o.BtType = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *BTRevisionCustomData2090) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionCustomData2090) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *BTRevisionCustomData2090) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *BTRevisionCustomData2090) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *BTRevisionCustomData2090) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionCustomData2090) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *BTRevisionCustomData2090) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *BTRevisionCustomData2090) SetRevision(v string) {
	o.Revision = &v
}

// GetValidRevisionReference returns the ValidRevisionReference field value if set, zero value otherwise.
func (o *BTRevisionCustomData2090) GetValidRevisionReference() bool {
	if o == nil || o.ValidRevisionReference == nil {
		var ret bool
		return ret
	}
	return *o.ValidRevisionReference
}

// GetValidRevisionReferenceOk returns a tuple with the ValidRevisionReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionCustomData2090) GetValidRevisionReferenceOk() (*bool, bool) {
	if o == nil || o.ValidRevisionReference == nil {
		return nil, false
	}
	return o.ValidRevisionReference, true
}

// HasValidRevisionReference returns a boolean if a field has been set.
func (o *BTRevisionCustomData2090) HasValidRevisionReference() bool {
	if o != nil && o.ValidRevisionReference != nil {
		return true
	}

	return false
}

// SetValidRevisionReference gets a reference to the given bool and assigns it to the ValidRevisionReference field.
func (o *BTRevisionCustomData2090) SetValidRevisionReference(v bool) {
	o.ValidRevisionReference = &v
}

func (o BTRevisionCustomData2090) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.PartNumber != nil {
		toSerialize["partNumber"] = o.PartNumber
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.ValidRevisionReference != nil {
		toSerialize["validRevisionReference"] = o.ValidRevisionReference
	}
	return json.Marshal(toSerialize)
}

type NullableBTRevisionCustomData2090 struct {
	value *BTRevisionCustomData2090
	isSet bool
}

func (v NullableBTRevisionCustomData2090) Get() *BTRevisionCustomData2090 {
	return v.value
}

func (v *NullableBTRevisionCustomData2090) Set(val *BTRevisionCustomData2090) {
	v.value = val
	v.isSet = true
}

func (v NullableBTRevisionCustomData2090) IsSet() bool {
	return v.isSet
}

func (v *NullableBTRevisionCustomData2090) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTRevisionCustomData2090(val *BTRevisionCustomData2090) *NullableBTRevisionCustomData2090 {
	return &NullableBTRevisionCustomData2090{value: val, isSet: true}
}

func (v NullableBTRevisionCustomData2090) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTRevisionCustomData2090) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
