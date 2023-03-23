/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13361-b3ca458f6e4e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAppElementParams struct for BTAppElementParams
type BTAppElementParams struct {
	// The label that will appear in the document's edit history for this operation. If blank, a value will be auto-generated.
	Description *string `json:"description,omitempty"`
	// The data type of the application. This string allows an application to distinguish their elements from elements of another application.
	FormatId string `json:"formatId"`
	// Initialization data for the new element's json tree.
	JsonTree *map[string]interface{}  `json:"jsonTree,omitempty"`
	Location *BTElementLocationParams `json:"location,omitempty"`
	// The name of the element being created. If blank, a name will be auto-generated.
	Name *string `json:"name,omitempty"`
	// Initialization data for the new element's subelements.
	Subelements []BTAppElementChangeParams `json:"subelements,omitempty"`
}

// NewBTAppElementParams instantiates a new BTAppElementParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppElementParams(formatId string) *BTAppElementParams {
	this := BTAppElementParams{}
	this.FormatId = formatId
	return &this
}

// NewBTAppElementParamsWithDefaults instantiates a new BTAppElementParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppElementParamsWithDefaults() *BTAppElementParams {
	this := BTAppElementParams{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTAppElementParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTAppElementParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTAppElementParams) SetDescription(v string) {
	o.Description = &v
}

// GetFormatId returns the FormatId field value
func (o *BTAppElementParams) GetFormatId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FormatId
}

// GetFormatIdOk returns a tuple with the FormatId field value
// and a boolean to check if the value has been set.
func (o *BTAppElementParams) GetFormatIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormatId, true
}

// SetFormatId sets field value
func (o *BTAppElementParams) SetFormatId(v string) {
	o.FormatId = v
}

// GetJsonTree returns the JsonTree field value if set, zero value otherwise.
func (o *BTAppElementParams) GetJsonTree() map[string]interface{} {
	if o == nil || o.JsonTree == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.JsonTree
}

// GetJsonTreeOk returns a tuple with the JsonTree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementParams) GetJsonTreeOk() (*map[string]interface{}, bool) {
	if o == nil || o.JsonTree == nil {
		return nil, false
	}
	return o.JsonTree, true
}

// HasJsonTree returns a boolean if a field has been set.
func (o *BTAppElementParams) HasJsonTree() bool {
	if o != nil && o.JsonTree != nil {
		return true
	}

	return false
}

// SetJsonTree gets a reference to the given map[string]interface{} and assigns it to the JsonTree field.
func (o *BTAppElementParams) SetJsonTree(v map[string]interface{}) {
	o.JsonTree = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *BTAppElementParams) GetLocation() BTElementLocationParams {
	if o == nil || o.Location == nil {
		var ret BTElementLocationParams
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementParams) GetLocationOk() (*BTElementLocationParams, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *BTAppElementParams) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given BTElementLocationParams and assigns it to the Location field.
func (o *BTAppElementParams) SetLocation(v BTElementLocationParams) {
	o.Location = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTAppElementParams) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementParams) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTAppElementParams) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTAppElementParams) SetName(v string) {
	o.Name = &v
}

// GetSubelements returns the Subelements field value if set, zero value otherwise.
func (o *BTAppElementParams) GetSubelements() []BTAppElementChangeParams {
	if o == nil || o.Subelements == nil {
		var ret []BTAppElementChangeParams
		return ret
	}
	return o.Subelements
}

// GetSubelementsOk returns a tuple with the Subelements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementParams) GetSubelementsOk() ([]BTAppElementChangeParams, bool) {
	if o == nil || o.Subelements == nil {
		return nil, false
	}
	return o.Subelements, true
}

// HasSubelements returns a boolean if a field has been set.
func (o *BTAppElementParams) HasSubelements() bool {
	if o != nil && o.Subelements != nil {
		return true
	}

	return false
}

// SetSubelements gets a reference to the given []BTAppElementChangeParams and assigns it to the Subelements field.
func (o *BTAppElementParams) SetSubelements(v []BTAppElementChangeParams) {
	o.Subelements = v
}

func (o BTAppElementParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["formatId"] = o.FormatId
	}
	if o.JsonTree != nil {
		toSerialize["jsonTree"] = o.JsonTree
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Subelements != nil {
		toSerialize["subelements"] = o.Subelements
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppElementParams struct {
	value *BTAppElementParams
	isSet bool
}

func (v NullableBTAppElementParams) Get() *BTAppElementParams {
	return v.value
}

func (v *NullableBTAppElementParams) Set(val *BTAppElementParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementParams(val *BTAppElementParams) *NullableBTAppElementParams {
	return &NullableBTAppElementParams{value: val, isSet: true}
}

func (v NullableBTAppElementParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
