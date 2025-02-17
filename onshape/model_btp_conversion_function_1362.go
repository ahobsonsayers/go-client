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

// BTPConversionFunction1362 struct for BTPConversionFunction1362
type BTPConversionFunction1362 struct {
	Atomic                *bool                       `json:"atomic,omitempty"`
	BtType                *string                     `json:"btType,omitempty"`
	DocumentationType     *GBTPDefinitionType         `json:"documentationType,omitempty"`
	EndSourceLocation     *int32                      `json:"endSourceLocation,omitempty"`
	NodeId                *string                     `json:"nodeId,omitempty"`
	ShortDescriptor       *string                     `json:"shortDescriptor,omitempty"`
	SpaceAfter            *BTPSpace10                 `json:"spaceAfter,omitempty"`
	SpaceBefore           *BTPSpace10                 `json:"spaceBefore,omitempty"`
	SpaceDefault          *bool                       `json:"spaceDefault,omitempty"`
	StartSourceLocation   *int32                      `json:"startSourceLocation,omitempty"`
	Annotation            *BTPAnnotation231           `json:"annotation,omitempty"`
	ArgumentsToDocument   []BTPArgumentDeclaration232 `json:"argumentsToDocument,omitempty"`
	Deprecated            *bool                       `json:"deprecated,omitempty"`
	DeprecatedExplanation *string                     `json:"deprecatedExplanation,omitempty"`
	ForExport             *bool                       `json:"forExport,omitempty"`
	SpaceAfterExport      *BTPSpace10                 `json:"spaceAfterExport,omitempty"`
	SymbolName            *BTPIdentifier8             `json:"symbolName,omitempty"`
	Arguments             []BTPArgumentDeclaration232 `json:"arguments,omitempty"`
	Body                  *BTPStatementBlock271       `json:"body,omitempty"`
	Precondition          *BTPStatement269            `json:"precondition,omitempty"`
	ReturnType            *BTPTypeName290             `json:"returnType,omitempty"`
	SpaceAfterArglist     *BTPSpace10                 `json:"spaceAfterArglist,omitempty"`
	SpaceInEmptyList      *BTPSpace10                 `json:"spaceInEmptyList,omitempty"`
	From                  *BTPLiteralNumber258        `json:"from,omitempty"`
	SpaceAfterType        *BTPSpace10                 `json:"spaceAfterType,omitempty"`
	To                    *BTPLiteralNumber258        `json:"to,omitempty"`
	TypeName              *BTPIdentifier8             `json:"typeName,omitempty"`
}

// NewBTPConversionFunction1362 instantiates a new BTPConversionFunction1362 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPConversionFunction1362() *BTPConversionFunction1362 {
	this := BTPConversionFunction1362{}
	return &this
}

// NewBTPConversionFunction1362WithDefaults instantiates a new BTPConversionFunction1362 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPConversionFunction1362WithDefaults() *BTPConversionFunction1362 {
	this := BTPConversionFunction1362{}
	return &this
}

// GetAtomic returns the Atomic field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetAtomic() bool {
	if o == nil || o.Atomic == nil {
		var ret bool
		return ret
	}
	return *o.Atomic
}

// GetAtomicOk returns a tuple with the Atomic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetAtomicOk() (*bool, bool) {
	if o == nil || o.Atomic == nil {
		return nil, false
	}
	return o.Atomic, true
}

// HasAtomic returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasAtomic() bool {
	if o != nil && o.Atomic != nil {
		return true
	}

	return false
}

// SetAtomic gets a reference to the given bool and assigns it to the Atomic field.
func (o *BTPConversionFunction1362) SetAtomic(v bool) {
	o.Atomic = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPConversionFunction1362) SetBtType(v string) {
	o.BtType = &v
}

// GetDocumentationType returns the DocumentationType field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetDocumentationType() GBTPDefinitionType {
	if o == nil || o.DocumentationType == nil {
		var ret GBTPDefinitionType
		return ret
	}
	return *o.DocumentationType
}

// GetDocumentationTypeOk returns a tuple with the DocumentationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetDocumentationTypeOk() (*GBTPDefinitionType, bool) {
	if o == nil || o.DocumentationType == nil {
		return nil, false
	}
	return o.DocumentationType, true
}

// HasDocumentationType returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasDocumentationType() bool {
	if o != nil && o.DocumentationType != nil {
		return true
	}

	return false
}

// SetDocumentationType gets a reference to the given GBTPDefinitionType and assigns it to the DocumentationType field.
func (o *BTPConversionFunction1362) SetDocumentationType(v GBTPDefinitionType) {
	o.DocumentationType = &v
}

// GetEndSourceLocation returns the EndSourceLocation field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetEndSourceLocation() int32 {
	if o == nil || o.EndSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.EndSourceLocation
}

// GetEndSourceLocationOk returns a tuple with the EndSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetEndSourceLocationOk() (*int32, bool) {
	if o == nil || o.EndSourceLocation == nil {
		return nil, false
	}
	return o.EndSourceLocation, true
}

// HasEndSourceLocation returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasEndSourceLocation() bool {
	if o != nil && o.EndSourceLocation != nil {
		return true
	}

	return false
}

// SetEndSourceLocation gets a reference to the given int32 and assigns it to the EndSourceLocation field.
func (o *BTPConversionFunction1362) SetEndSourceLocation(v int32) {
	o.EndSourceLocation = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTPConversionFunction1362) SetNodeId(v string) {
	o.NodeId = &v
}

// GetShortDescriptor returns the ShortDescriptor field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetShortDescriptor() string {
	if o == nil || o.ShortDescriptor == nil {
		var ret string
		return ret
	}
	return *o.ShortDescriptor
}

// GetShortDescriptorOk returns a tuple with the ShortDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetShortDescriptorOk() (*string, bool) {
	if o == nil || o.ShortDescriptor == nil {
		return nil, false
	}
	return o.ShortDescriptor, true
}

// HasShortDescriptor returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasShortDescriptor() bool {
	if o != nil && o.ShortDescriptor != nil {
		return true
	}

	return false
}

// SetShortDescriptor gets a reference to the given string and assigns it to the ShortDescriptor field.
func (o *BTPConversionFunction1362) SetShortDescriptor(v string) {
	o.ShortDescriptor = &v
}

// GetSpaceAfter returns the SpaceAfter field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetSpaceAfter() BTPSpace10 {
	if o == nil || o.SpaceAfter == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfter
}

// GetSpaceAfterOk returns a tuple with the SpaceAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetSpaceAfterOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfter == nil {
		return nil, false
	}
	return o.SpaceAfter, true
}

// HasSpaceAfter returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasSpaceAfter() bool {
	if o != nil && o.SpaceAfter != nil {
		return true
	}

	return false
}

// SetSpaceAfter gets a reference to the given BTPSpace10 and assigns it to the SpaceAfter field.
func (o *BTPConversionFunction1362) SetSpaceAfter(v BTPSpace10) {
	o.SpaceAfter = &v
}

// GetSpaceBefore returns the SpaceBefore field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetSpaceBefore() BTPSpace10 {
	if o == nil || o.SpaceBefore == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBefore
}

// GetSpaceBeforeOk returns a tuple with the SpaceBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetSpaceBeforeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBefore == nil {
		return nil, false
	}
	return o.SpaceBefore, true
}

// HasSpaceBefore returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasSpaceBefore() bool {
	if o != nil && o.SpaceBefore != nil {
		return true
	}

	return false
}

// SetSpaceBefore gets a reference to the given BTPSpace10 and assigns it to the SpaceBefore field.
func (o *BTPConversionFunction1362) SetSpaceBefore(v BTPSpace10) {
	o.SpaceBefore = &v
}

// GetSpaceDefault returns the SpaceDefault field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetSpaceDefault() bool {
	if o == nil || o.SpaceDefault == nil {
		var ret bool
		return ret
	}
	return *o.SpaceDefault
}

// GetSpaceDefaultOk returns a tuple with the SpaceDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetSpaceDefaultOk() (*bool, bool) {
	if o == nil || o.SpaceDefault == nil {
		return nil, false
	}
	return o.SpaceDefault, true
}

// HasSpaceDefault returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasSpaceDefault() bool {
	if o != nil && o.SpaceDefault != nil {
		return true
	}

	return false
}

// SetSpaceDefault gets a reference to the given bool and assigns it to the SpaceDefault field.
func (o *BTPConversionFunction1362) SetSpaceDefault(v bool) {
	o.SpaceDefault = &v
}

// GetStartSourceLocation returns the StartSourceLocation field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetStartSourceLocation() int32 {
	if o == nil || o.StartSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.StartSourceLocation
}

// GetStartSourceLocationOk returns a tuple with the StartSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetStartSourceLocationOk() (*int32, bool) {
	if o == nil || o.StartSourceLocation == nil {
		return nil, false
	}
	return o.StartSourceLocation, true
}

// HasStartSourceLocation returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasStartSourceLocation() bool {
	if o != nil && o.StartSourceLocation != nil {
		return true
	}

	return false
}

// SetStartSourceLocation gets a reference to the given int32 and assigns it to the StartSourceLocation field.
func (o *BTPConversionFunction1362) SetStartSourceLocation(v int32) {
	o.StartSourceLocation = &v
}

// GetAnnotation returns the Annotation field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetAnnotation() BTPAnnotation231 {
	if o == nil || o.Annotation == nil {
		var ret BTPAnnotation231
		return ret
	}
	return *o.Annotation
}

// GetAnnotationOk returns a tuple with the Annotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetAnnotationOk() (*BTPAnnotation231, bool) {
	if o == nil || o.Annotation == nil {
		return nil, false
	}
	return o.Annotation, true
}

// HasAnnotation returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasAnnotation() bool {
	if o != nil && o.Annotation != nil {
		return true
	}

	return false
}

// SetAnnotation gets a reference to the given BTPAnnotation231 and assigns it to the Annotation field.
func (o *BTPConversionFunction1362) SetAnnotation(v BTPAnnotation231) {
	o.Annotation = &v
}

// GetArgumentsToDocument returns the ArgumentsToDocument field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetArgumentsToDocument() []BTPArgumentDeclaration232 {
	if o == nil || o.ArgumentsToDocument == nil {
		var ret []BTPArgumentDeclaration232
		return ret
	}
	return o.ArgumentsToDocument
}

// GetArgumentsToDocumentOk returns a tuple with the ArgumentsToDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetArgumentsToDocumentOk() ([]BTPArgumentDeclaration232, bool) {
	if o == nil || o.ArgumentsToDocument == nil {
		return nil, false
	}
	return o.ArgumentsToDocument, true
}

// HasArgumentsToDocument returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasArgumentsToDocument() bool {
	if o != nil && o.ArgumentsToDocument != nil {
		return true
	}

	return false
}

// SetArgumentsToDocument gets a reference to the given []BTPArgumentDeclaration232 and assigns it to the ArgumentsToDocument field.
func (o *BTPConversionFunction1362) SetArgumentsToDocument(v []BTPArgumentDeclaration232) {
	o.ArgumentsToDocument = v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetDeprecated() bool {
	if o == nil || o.Deprecated == nil {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetDeprecatedOk() (*bool, bool) {
	if o == nil || o.Deprecated == nil {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasDeprecated() bool {
	if o != nil && o.Deprecated != nil {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *BTPConversionFunction1362) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetDeprecatedExplanation returns the DeprecatedExplanation field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetDeprecatedExplanation() string {
	if o == nil || o.DeprecatedExplanation == nil {
		var ret string
		return ret
	}
	return *o.DeprecatedExplanation
}

// GetDeprecatedExplanationOk returns a tuple with the DeprecatedExplanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetDeprecatedExplanationOk() (*string, bool) {
	if o == nil || o.DeprecatedExplanation == nil {
		return nil, false
	}
	return o.DeprecatedExplanation, true
}

// HasDeprecatedExplanation returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasDeprecatedExplanation() bool {
	if o != nil && o.DeprecatedExplanation != nil {
		return true
	}

	return false
}

// SetDeprecatedExplanation gets a reference to the given string and assigns it to the DeprecatedExplanation field.
func (o *BTPConversionFunction1362) SetDeprecatedExplanation(v string) {
	o.DeprecatedExplanation = &v
}

// GetForExport returns the ForExport field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetForExport() bool {
	if o == nil || o.ForExport == nil {
		var ret bool
		return ret
	}
	return *o.ForExport
}

// GetForExportOk returns a tuple with the ForExport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetForExportOk() (*bool, bool) {
	if o == nil || o.ForExport == nil {
		return nil, false
	}
	return o.ForExport, true
}

// HasForExport returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasForExport() bool {
	if o != nil && o.ForExport != nil {
		return true
	}

	return false
}

// SetForExport gets a reference to the given bool and assigns it to the ForExport field.
func (o *BTPConversionFunction1362) SetForExport(v bool) {
	o.ForExport = &v
}

// GetSpaceAfterExport returns the SpaceAfterExport field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetSpaceAfterExport() BTPSpace10 {
	if o == nil || o.SpaceAfterExport == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterExport
}

// GetSpaceAfterExportOk returns a tuple with the SpaceAfterExport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetSpaceAfterExportOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterExport == nil {
		return nil, false
	}
	return o.SpaceAfterExport, true
}

// HasSpaceAfterExport returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasSpaceAfterExport() bool {
	if o != nil && o.SpaceAfterExport != nil {
		return true
	}

	return false
}

// SetSpaceAfterExport gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterExport field.
func (o *BTPConversionFunction1362) SetSpaceAfterExport(v BTPSpace10) {
	o.SpaceAfterExport = &v
}

// GetSymbolName returns the SymbolName field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetSymbolName() BTPIdentifier8 {
	if o == nil || o.SymbolName == nil {
		var ret BTPIdentifier8
		return ret
	}
	return *o.SymbolName
}

// GetSymbolNameOk returns a tuple with the SymbolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetSymbolNameOk() (*BTPIdentifier8, bool) {
	if o == nil || o.SymbolName == nil {
		return nil, false
	}
	return o.SymbolName, true
}

// HasSymbolName returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasSymbolName() bool {
	if o != nil && o.SymbolName != nil {
		return true
	}

	return false
}

// SetSymbolName gets a reference to the given BTPIdentifier8 and assigns it to the SymbolName field.
func (o *BTPConversionFunction1362) SetSymbolName(v BTPIdentifier8) {
	o.SymbolName = &v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetArguments() []BTPArgumentDeclaration232 {
	if o == nil || o.Arguments == nil {
		var ret []BTPArgumentDeclaration232
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetArgumentsOk() ([]BTPArgumentDeclaration232, bool) {
	if o == nil || o.Arguments == nil {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasArguments() bool {
	if o != nil && o.Arguments != nil {
		return true
	}

	return false
}

// SetArguments gets a reference to the given []BTPArgumentDeclaration232 and assigns it to the Arguments field.
func (o *BTPConversionFunction1362) SetArguments(v []BTPArgumentDeclaration232) {
	o.Arguments = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetBody() BTPStatementBlock271 {
	if o == nil || o.Body == nil {
		var ret BTPStatementBlock271
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetBodyOk() (*BTPStatementBlock271, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given BTPStatementBlock271 and assigns it to the Body field.
func (o *BTPConversionFunction1362) SetBody(v BTPStatementBlock271) {
	o.Body = &v
}

// GetPrecondition returns the Precondition field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetPrecondition() BTPStatement269 {
	if o == nil || o.Precondition == nil {
		var ret BTPStatement269
		return ret
	}
	return *o.Precondition
}

// GetPreconditionOk returns a tuple with the Precondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetPreconditionOk() (*BTPStatement269, bool) {
	if o == nil || o.Precondition == nil {
		return nil, false
	}
	return o.Precondition, true
}

// HasPrecondition returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasPrecondition() bool {
	if o != nil && o.Precondition != nil {
		return true
	}

	return false
}

// SetPrecondition gets a reference to the given BTPStatement269 and assigns it to the Precondition field.
func (o *BTPConversionFunction1362) SetPrecondition(v BTPStatement269) {
	o.Precondition = &v
}

// GetReturnType returns the ReturnType field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetReturnType() BTPTypeName290 {
	if o == nil || o.ReturnType == nil {
		var ret BTPTypeName290
		return ret
	}
	return *o.ReturnType
}

// GetReturnTypeOk returns a tuple with the ReturnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetReturnTypeOk() (*BTPTypeName290, bool) {
	if o == nil || o.ReturnType == nil {
		return nil, false
	}
	return o.ReturnType, true
}

// HasReturnType returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasReturnType() bool {
	if o != nil && o.ReturnType != nil {
		return true
	}

	return false
}

// SetReturnType gets a reference to the given BTPTypeName290 and assigns it to the ReturnType field.
func (o *BTPConversionFunction1362) SetReturnType(v BTPTypeName290) {
	o.ReturnType = &v
}

// GetSpaceAfterArglist returns the SpaceAfterArglist field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetSpaceAfterArglist() BTPSpace10 {
	if o == nil || o.SpaceAfterArglist == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterArglist
}

// GetSpaceAfterArglistOk returns a tuple with the SpaceAfterArglist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetSpaceAfterArglistOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterArglist == nil {
		return nil, false
	}
	return o.SpaceAfterArglist, true
}

// HasSpaceAfterArglist returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasSpaceAfterArglist() bool {
	if o != nil && o.SpaceAfterArglist != nil {
		return true
	}

	return false
}

// SetSpaceAfterArglist gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterArglist field.
func (o *BTPConversionFunction1362) SetSpaceAfterArglist(v BTPSpace10) {
	o.SpaceAfterArglist = &v
}

// GetSpaceInEmptyList returns the SpaceInEmptyList field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetSpaceInEmptyList() BTPSpace10 {
	if o == nil || o.SpaceInEmptyList == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceInEmptyList
}

// GetSpaceInEmptyListOk returns a tuple with the SpaceInEmptyList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetSpaceInEmptyListOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceInEmptyList == nil {
		return nil, false
	}
	return o.SpaceInEmptyList, true
}

// HasSpaceInEmptyList returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasSpaceInEmptyList() bool {
	if o != nil && o.SpaceInEmptyList != nil {
		return true
	}

	return false
}

// SetSpaceInEmptyList gets a reference to the given BTPSpace10 and assigns it to the SpaceInEmptyList field.
func (o *BTPConversionFunction1362) SetSpaceInEmptyList(v BTPSpace10) {
	o.SpaceInEmptyList = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetFrom() BTPLiteralNumber258 {
	if o == nil || o.From == nil {
		var ret BTPLiteralNumber258
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetFromOk() (*BTPLiteralNumber258, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given BTPLiteralNumber258 and assigns it to the From field.
func (o *BTPConversionFunction1362) SetFrom(v BTPLiteralNumber258) {
	o.From = &v
}

// GetSpaceAfterType returns the SpaceAfterType field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetSpaceAfterType() BTPSpace10 {
	if o == nil || o.SpaceAfterType == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterType
}

// GetSpaceAfterTypeOk returns a tuple with the SpaceAfterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetSpaceAfterTypeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterType == nil {
		return nil, false
	}
	return o.SpaceAfterType, true
}

// HasSpaceAfterType returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasSpaceAfterType() bool {
	if o != nil && o.SpaceAfterType != nil {
		return true
	}

	return false
}

// SetSpaceAfterType gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterType field.
func (o *BTPConversionFunction1362) SetSpaceAfterType(v BTPSpace10) {
	o.SpaceAfterType = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetTo() BTPLiteralNumber258 {
	if o == nil || o.To == nil {
		var ret BTPLiteralNumber258
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetToOk() (*BTPLiteralNumber258, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given BTPLiteralNumber258 and assigns it to the To field.
func (o *BTPConversionFunction1362) SetTo(v BTPLiteralNumber258) {
	o.To = &v
}

// GetTypeName returns the TypeName field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetTypeName() BTPIdentifier8 {
	if o == nil || o.TypeName == nil {
		var ret BTPIdentifier8
		return ret
	}
	return *o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetTypeNameOk() (*BTPIdentifier8, bool) {
	if o == nil || o.TypeName == nil {
		return nil, false
	}
	return o.TypeName, true
}

// HasTypeName returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasTypeName() bool {
	if o != nil && o.TypeName != nil {
		return true
	}

	return false
}

// SetTypeName gets a reference to the given BTPIdentifier8 and assigns it to the TypeName field.
func (o *BTPConversionFunction1362) SetTypeName(v BTPIdentifier8) {
	o.TypeName = &v
}

func (o BTPConversionFunction1362) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Atomic != nil {
		toSerialize["atomic"] = o.Atomic
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DocumentationType != nil {
		toSerialize["documentationType"] = o.DocumentationType
	}
	if o.EndSourceLocation != nil {
		toSerialize["endSourceLocation"] = o.EndSourceLocation
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ShortDescriptor != nil {
		toSerialize["shortDescriptor"] = o.ShortDescriptor
	}
	if o.SpaceAfter != nil {
		toSerialize["spaceAfter"] = o.SpaceAfter
	}
	if o.SpaceBefore != nil {
		toSerialize["spaceBefore"] = o.SpaceBefore
	}
	if o.SpaceDefault != nil {
		toSerialize["spaceDefault"] = o.SpaceDefault
	}
	if o.StartSourceLocation != nil {
		toSerialize["startSourceLocation"] = o.StartSourceLocation
	}
	if o.Annotation != nil {
		toSerialize["annotation"] = o.Annotation
	}
	if o.ArgumentsToDocument != nil {
		toSerialize["argumentsToDocument"] = o.ArgumentsToDocument
	}
	if o.Deprecated != nil {
		toSerialize["deprecated"] = o.Deprecated
	}
	if o.DeprecatedExplanation != nil {
		toSerialize["deprecatedExplanation"] = o.DeprecatedExplanation
	}
	if o.ForExport != nil {
		toSerialize["forExport"] = o.ForExport
	}
	if o.SpaceAfterExport != nil {
		toSerialize["spaceAfterExport"] = o.SpaceAfterExport
	}
	if o.SymbolName != nil {
		toSerialize["symbolName"] = o.SymbolName
	}
	if o.Arguments != nil {
		toSerialize["arguments"] = o.Arguments
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.Precondition != nil {
		toSerialize["precondition"] = o.Precondition
	}
	if o.ReturnType != nil {
		toSerialize["returnType"] = o.ReturnType
	}
	if o.SpaceAfterArglist != nil {
		toSerialize["spaceAfterArglist"] = o.SpaceAfterArglist
	}
	if o.SpaceInEmptyList != nil {
		toSerialize["spaceInEmptyList"] = o.SpaceInEmptyList
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.SpaceAfterType != nil {
		toSerialize["spaceAfterType"] = o.SpaceAfterType
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	if o.TypeName != nil {
		toSerialize["typeName"] = o.TypeName
	}
	return json.Marshal(toSerialize)
}

type NullableBTPConversionFunction1362 struct {
	value *BTPConversionFunction1362
	isSet bool
}

func (v NullableBTPConversionFunction1362) Get() *BTPConversionFunction1362 {
	return v.value
}

func (v *NullableBTPConversionFunction1362) Set(val *BTPConversionFunction1362) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPConversionFunction1362) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPConversionFunction1362) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPConversionFunction1362(val *BTPConversionFunction1362) *NullableBTPConversionFunction1362 {
	return &NullableBTPConversionFunction1362{value: val, isSet: true}
}

func (v NullableBTPConversionFunction1362) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPConversionFunction1362) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
