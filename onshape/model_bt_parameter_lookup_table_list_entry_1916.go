/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.6979-8ce9514d51cf
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTParameterLookupTableListEntry1916 struct for BTParameterLookupTableListEntry1916
type BTParameterLookupTableListEntry1916 struct {
	AdditionalLocalizedStrings *int32                            `json:"additionalLocalizedStrings,omitempty"`
	BtType                     *string                           `json:"btType,omitempty"`
	DefaultIndex               *int32                            `json:"defaultIndex,omitempty"`
	DisplayName                *string                           `json:"displayName,omitempty"`
	Entries                    []BTParameterLookupTableEntry1667 `json:"entries,omitempty"`
	Label                      *string                           `json:"label,omitempty"`
	LocalizableName            *string                           `json:"localizableName,omitempty"`
	LocalizedLabel             *string                           `json:"localizedLabel,omitempty"`
	LocalizedName              *string                           `json:"localizedName,omitempty"`
	Name                       *string                           `json:"name,omitempty"`
	StringsToLocalize          []string                          `json:"stringsToLocalize,omitempty"`
}

// NewBTParameterLookupTableListEntry1916 instantiates a new BTParameterLookupTableListEntry1916 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterLookupTableListEntry1916() *BTParameterLookupTableListEntry1916 {
	this := BTParameterLookupTableListEntry1916{}
	return &this
}

// NewBTParameterLookupTableListEntry1916WithDefaults instantiates a new BTParameterLookupTableListEntry1916 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterLookupTableListEntry1916WithDefaults() *BTParameterLookupTableListEntry1916 {
	this := BTParameterLookupTableListEntry1916{}
	return &this
}

// GetAdditionalLocalizedStrings returns the AdditionalLocalizedStrings field value if set, zero value otherwise.
func (o *BTParameterLookupTableListEntry1916) GetAdditionalLocalizedStrings() int32 {
	if o == nil || o.AdditionalLocalizedStrings == nil {
		var ret int32
		return ret
	}
	return *o.AdditionalLocalizedStrings
}

// GetAdditionalLocalizedStringsOk returns a tuple with the AdditionalLocalizedStrings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableListEntry1916) GetAdditionalLocalizedStringsOk() (*int32, bool) {
	if o == nil || o.AdditionalLocalizedStrings == nil {
		return nil, false
	}
	return o.AdditionalLocalizedStrings, true
}

// HasAdditionalLocalizedStrings returns a boolean if a field has been set.
func (o *BTParameterLookupTableListEntry1916) HasAdditionalLocalizedStrings() bool {
	if o != nil && o.AdditionalLocalizedStrings != nil {
		return true
	}

	return false
}

// SetAdditionalLocalizedStrings gets a reference to the given int32 and assigns it to the AdditionalLocalizedStrings field.
func (o *BTParameterLookupTableListEntry1916) SetAdditionalLocalizedStrings(v int32) {
	o.AdditionalLocalizedStrings = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterLookupTableListEntry1916) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableListEntry1916) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTParameterLookupTableListEntry1916) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTParameterLookupTableListEntry1916) SetBtType(v string) {
	o.BtType = &v
}

// GetDefaultIndex returns the DefaultIndex field value if set, zero value otherwise.
func (o *BTParameterLookupTableListEntry1916) GetDefaultIndex() int32 {
	if o == nil || o.DefaultIndex == nil {
		var ret int32
		return ret
	}
	return *o.DefaultIndex
}

// GetDefaultIndexOk returns a tuple with the DefaultIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableListEntry1916) GetDefaultIndexOk() (*int32, bool) {
	if o == nil || o.DefaultIndex == nil {
		return nil, false
	}
	return o.DefaultIndex, true
}

// HasDefaultIndex returns a boolean if a field has been set.
func (o *BTParameterLookupTableListEntry1916) HasDefaultIndex() bool {
	if o != nil && o.DefaultIndex != nil {
		return true
	}

	return false
}

// SetDefaultIndex gets a reference to the given int32 and assigns it to the DefaultIndex field.
func (o *BTParameterLookupTableListEntry1916) SetDefaultIndex(v int32) {
	o.DefaultIndex = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *BTParameterLookupTableListEntry1916) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableListEntry1916) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *BTParameterLookupTableListEntry1916) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *BTParameterLookupTableListEntry1916) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *BTParameterLookupTableListEntry1916) GetEntries() []BTParameterLookupTableEntry1667 {
	if o == nil || o.Entries == nil {
		var ret []BTParameterLookupTableEntry1667
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableListEntry1916) GetEntriesOk() ([]BTParameterLookupTableEntry1667, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *BTParameterLookupTableListEntry1916) HasEntries() bool {
	if o != nil && o.Entries != nil {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []BTParameterLookupTableEntry1667 and assigns it to the Entries field.
func (o *BTParameterLookupTableListEntry1916) SetEntries(v []BTParameterLookupTableEntry1667) {
	o.Entries = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *BTParameterLookupTableListEntry1916) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableListEntry1916) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *BTParameterLookupTableListEntry1916) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *BTParameterLookupTableListEntry1916) SetLabel(v string) {
	o.Label = &v
}

// GetLocalizableName returns the LocalizableName field value if set, zero value otherwise.
func (o *BTParameterLookupTableListEntry1916) GetLocalizableName() string {
	if o == nil || o.LocalizableName == nil {
		var ret string
		return ret
	}
	return *o.LocalizableName
}

// GetLocalizableNameOk returns a tuple with the LocalizableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableListEntry1916) GetLocalizableNameOk() (*string, bool) {
	if o == nil || o.LocalizableName == nil {
		return nil, false
	}
	return o.LocalizableName, true
}

// HasLocalizableName returns a boolean if a field has been set.
func (o *BTParameterLookupTableListEntry1916) HasLocalizableName() bool {
	if o != nil && o.LocalizableName != nil {
		return true
	}

	return false
}

// SetLocalizableName gets a reference to the given string and assigns it to the LocalizableName field.
func (o *BTParameterLookupTableListEntry1916) SetLocalizableName(v string) {
	o.LocalizableName = &v
}

// GetLocalizedLabel returns the LocalizedLabel field value if set, zero value otherwise.
func (o *BTParameterLookupTableListEntry1916) GetLocalizedLabel() string {
	if o == nil || o.LocalizedLabel == nil {
		var ret string
		return ret
	}
	return *o.LocalizedLabel
}

// GetLocalizedLabelOk returns a tuple with the LocalizedLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableListEntry1916) GetLocalizedLabelOk() (*string, bool) {
	if o == nil || o.LocalizedLabel == nil {
		return nil, false
	}
	return o.LocalizedLabel, true
}

// HasLocalizedLabel returns a boolean if a field has been set.
func (o *BTParameterLookupTableListEntry1916) HasLocalizedLabel() bool {
	if o != nil && o.LocalizedLabel != nil {
		return true
	}

	return false
}

// SetLocalizedLabel gets a reference to the given string and assigns it to the LocalizedLabel field.
func (o *BTParameterLookupTableListEntry1916) SetLocalizedLabel(v string) {
	o.LocalizedLabel = &v
}

// GetLocalizedName returns the LocalizedName field value if set, zero value otherwise.
func (o *BTParameterLookupTableListEntry1916) GetLocalizedName() string {
	if o == nil || o.LocalizedName == nil {
		var ret string
		return ret
	}
	return *o.LocalizedName
}

// GetLocalizedNameOk returns a tuple with the LocalizedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableListEntry1916) GetLocalizedNameOk() (*string, bool) {
	if o == nil || o.LocalizedName == nil {
		return nil, false
	}
	return o.LocalizedName, true
}

// HasLocalizedName returns a boolean if a field has been set.
func (o *BTParameterLookupTableListEntry1916) HasLocalizedName() bool {
	if o != nil && o.LocalizedName != nil {
		return true
	}

	return false
}

// SetLocalizedName gets a reference to the given string and assigns it to the LocalizedName field.
func (o *BTParameterLookupTableListEntry1916) SetLocalizedName(v string) {
	o.LocalizedName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTParameterLookupTableListEntry1916) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableListEntry1916) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTParameterLookupTableListEntry1916) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTParameterLookupTableListEntry1916) SetName(v string) {
	o.Name = &v
}

// GetStringsToLocalize returns the StringsToLocalize field value if set, zero value otherwise.
func (o *BTParameterLookupTableListEntry1916) GetStringsToLocalize() []string {
	if o == nil || o.StringsToLocalize == nil {
		var ret []string
		return ret
	}
	return o.StringsToLocalize
}

// GetStringsToLocalizeOk returns a tuple with the StringsToLocalize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableListEntry1916) GetStringsToLocalizeOk() ([]string, bool) {
	if o == nil || o.StringsToLocalize == nil {
		return nil, false
	}
	return o.StringsToLocalize, true
}

// HasStringsToLocalize returns a boolean if a field has been set.
func (o *BTParameterLookupTableListEntry1916) HasStringsToLocalize() bool {
	if o != nil && o.StringsToLocalize != nil {
		return true
	}

	return false
}

// SetStringsToLocalize gets a reference to the given []string and assigns it to the StringsToLocalize field.
func (o *BTParameterLookupTableListEntry1916) SetStringsToLocalize(v []string) {
	o.StringsToLocalize = v
}

func (o BTParameterLookupTableListEntry1916) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalLocalizedStrings != nil {
		toSerialize["additionalLocalizedStrings"] = o.AdditionalLocalizedStrings
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DefaultIndex != nil {
		toSerialize["defaultIndex"] = o.DefaultIndex
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.LocalizableName != nil {
		toSerialize["localizableName"] = o.LocalizableName
	}
	if o.LocalizedLabel != nil {
		toSerialize["localizedLabel"] = o.LocalizedLabel
	}
	if o.LocalizedName != nil {
		toSerialize["localizedName"] = o.LocalizedName
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.StringsToLocalize != nil {
		toSerialize["stringsToLocalize"] = o.StringsToLocalize
	}
	return json.Marshal(toSerialize)
}

type NullableBTParameterLookupTableListEntry1916 struct {
	value *BTParameterLookupTableListEntry1916
	isSet bool
}

func (v NullableBTParameterLookupTableListEntry1916) Get() *BTParameterLookupTableListEntry1916 {
	return v.value
}

func (v *NullableBTParameterLookupTableListEntry1916) Set(val *BTParameterLookupTableListEntry1916) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterLookupTableListEntry1916) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterLookupTableListEntry1916) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterLookupTableListEntry1916(val *BTParameterLookupTableListEntry1916) *NullableBTParameterLookupTableListEntry1916 {
	return &NullableBTParameterLookupTableListEntry1916{value: val, isSet: true}
}

func (v NullableBTParameterLookupTableListEntry1916) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterLookupTableListEntry1916) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
