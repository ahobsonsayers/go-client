/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5768-0013f50d23d2
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTElementLocationParams The location at which the new element should be inserted.
type BTElementLocationParams struct {
	// The id of an element which provides context for the position value specified.
	ElementId *string `json:"elementId,omitempty"`
	// An indicator for the relative placement of the new element. If elementId is specified, a negative number indicates insertion prior to the element and a non-negative number indicates insertion following the element. If no elementId is specified, a negative value indicates insertion at the end of the element list and a non-negative number indicates insertion at the start of the element list.
	Position *int32 `json:"position,omitempty"`
}

// NewBTElementLocationParams instantiates a new BTElementLocationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTElementLocationParams() *BTElementLocationParams {
	this := BTElementLocationParams{}
	return &this
}

// NewBTElementLocationParamsWithDefaults instantiates a new BTElementLocationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTElementLocationParamsWithDefaults() *BTElementLocationParams {
	this := BTElementLocationParams{}
	return &this
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTElementLocationParams) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementLocationParams) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTElementLocationParams) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTElementLocationParams) SetElementId(v string) {
	o.ElementId = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *BTElementLocationParams) GetPosition() int32 {
	if o == nil || o.Position == nil {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementLocationParams) GetPositionOk() (*int32, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *BTElementLocationParams) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *BTElementLocationParams) SetPosition(v int32) {
	o.Position = &v
}

func (o BTElementLocationParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	return json.Marshal(toSerialize)
}

type NullableBTElementLocationParams struct {
	value *BTElementLocationParams
	isSet bool
}

func (v NullableBTElementLocationParams) Get() *BTElementLocationParams {
	return v.value
}

func (v *NullableBTElementLocationParams) Set(val *BTElementLocationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTElementLocationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTElementLocationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTElementLocationParams(val *BTElementLocationParams) *NullableBTElementLocationParams {
	return &NullableBTElementLocationParams{value: val, isSet: true}
}

func (v NullableBTElementLocationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTElementLocationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
