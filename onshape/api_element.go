/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6672-3324fec9d980
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// ElementApiService ElementApi service
type ElementApiService service

type ApiCopyElementFromSourceDocumentRequest struct {
	ctx                 context.Context
	ApiService          *ElementApiService
	did                 string
	wid                 string
	bTCopyElementParams *BTCopyElementParams
}

func (r ApiCopyElementFromSourceDocumentRequest) BTCopyElementParams(bTCopyElementParams BTCopyElementParams) ApiCopyElementFromSourceDocumentRequest {
	r.bTCopyElementParams = &bTCopyElementParams
	return r
}

func (r ApiCopyElementFromSourceDocumentRequest) Execute() (*BTDocumentElementInfo, *http.Response, error) {
	return r.ApiService.CopyElementFromSourceDocumentExecute(r)
}

/*
CopyElementFromSourceDocument Copy tab by document ID and workspace ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @param wid
 @return ApiCopyElementFromSourceDocumentRequest
*/
func (a *ElementApiService) CopyElementFromSourceDocument(ctx context.Context, did string, wid string) ApiCopyElementFromSourceDocumentRequest {
	return ApiCopyElementFromSourceDocumentRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wid:        wid,
	}
}

// Execute executes the request
//  @return BTDocumentElementInfo
func (a *ElementApiService) CopyElementFromSourceDocumentExecute(r ApiCopyElementFromSourceDocumentRequest) (*BTDocumentElementInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTDocumentElementInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ElementApiService.CopyElementFromSourceDocument")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/elements/copyelement/{did}/workspace/{wid}"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wid"+"}", url.PathEscape(parameterToString(r.wid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTCopyElementParams == nil {
		return localVarReturnValue, nil, reportError("bTCopyElementParams is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bTCopyElementParams
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTDocumentElementInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDecodeConfigurationRequest struct {
	ctx               context.Context
	ApiService        *ElementApiService
	did               string
	wvm               string
	wvmid             string
	eid               string
	cid               string
	linkDocumentId    *string
	includeDisplay    *bool
	configurationIsId *bool
}

func (r ApiDecodeConfigurationRequest) LinkDocumentId(linkDocumentId string) ApiDecodeConfigurationRequest {
	r.linkDocumentId = &linkDocumentId
	return r
}

func (r ApiDecodeConfigurationRequest) IncludeDisplay(includeDisplay bool) ApiDecodeConfigurationRequest {
	r.includeDisplay = &includeDisplay
	return r
}

func (r ApiDecodeConfigurationRequest) ConfigurationIsId(configurationIsId bool) ApiDecodeConfigurationRequest {
	r.configurationIsId = &configurationIsId
	return r
}

func (r ApiDecodeConfigurationRequest) Execute() (*BTConfigurationInfo, *http.Response, error) {
	return r.ApiService.DecodeConfigurationExecute(r)
}

/*
DecodeConfiguration Decode configuration string by documentation ID, workspace or version or microversion ID, and tab ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @param wvm
 @param wvmid
 @param eid
 @param cid
 @return ApiDecodeConfigurationRequest
*/
func (a *ElementApiService) DecodeConfiguration(ctx context.Context, did string, wvm string, wvmid string, eid string, cid string) ApiDecodeConfigurationRequest {
	return ApiDecodeConfigurationRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wvm:        wvm,
		wvmid:      wvmid,
		eid:        eid,
		cid:        cid,
	}
}

// Execute executes the request
//  @return BTConfigurationInfo
func (a *ElementApiService) DecodeConfigurationExecute(r ApiDecodeConfigurationRequest) (*BTConfigurationInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTConfigurationInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ElementApiService.DecodeConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/elements/d/{did}/{wvm}/{wvmid}/e/{eid}/configurationencodings/{cid}"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvm"+"}", url.PathEscape(parameterToString(r.wvm, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvmid"+"}", url.PathEscape(parameterToString(r.wvmid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cid"+"}", url.PathEscape(parameterToString(r.cid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.linkDocumentId != nil {
		localVarQueryParams.Add("linkDocumentId", parameterToString(*r.linkDocumentId, ""))
	}
	if r.includeDisplay != nil {
		localVarQueryParams.Add("includeDisplay", parameterToString(*r.includeDisplay, ""))
	}
	if r.configurationIsId != nil {
		localVarQueryParams.Add("configurationIsId", parameterToString(*r.configurationIsId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTConfigurationInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteElementRequest struct {
	ctx        context.Context
	ApiService *ElementApiService
	did        string
	wid        string
	eid        string
}

func (r ApiDeleteElementRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteElementExecute(r)
}

/*
DeleteElement Method for DeleteElement

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @param wid
 @param eid
 @return ApiDeleteElementRequest
*/
func (a *ElementApiService) DeleteElement(ctx context.Context, did string, wid string, eid string) ApiDeleteElementRequest {
	return ApiDeleteElementRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wid:        wid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ElementApiService) DeleteElementExecute(r ApiDeleteElementRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ElementApiService.DeleteElement")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/elements/d/{did}/w/{wid}/e/{eid}"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wid"+"}", url.PathEscape(parameterToString(r.wid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v map[string]interface{}
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiEncodeConfigurationMapRequest struct {
	ctx                   context.Context
	ApiService            *ElementApiService
	did                   string
	eid                   string
	bTConfigurationParams *BTConfigurationParams
	versionId             *string
	linkDocumentId        *string
}

func (r ApiEncodeConfigurationMapRequest) BTConfigurationParams(bTConfigurationParams BTConfigurationParams) ApiEncodeConfigurationMapRequest {
	r.bTConfigurationParams = &bTConfigurationParams
	return r
}

func (r ApiEncodeConfigurationMapRequest) VersionId(versionId string) ApiEncodeConfigurationMapRequest {
	r.versionId = &versionId
	return r
}

func (r ApiEncodeConfigurationMapRequest) LinkDocumentId(linkDocumentId string) ApiEncodeConfigurationMapRequest {
	r.linkDocumentId = &linkDocumentId
	return r
}

func (r ApiEncodeConfigurationMapRequest) Execute() (*BTEncodedConfigurationInfo, *http.Response, error) {
	return r.ApiService.EncodeConfigurationMapExecute(r)
}

/*
EncodeConfigurationMap Encode configuration by documentation ID and tab ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @param eid
 @return ApiEncodeConfigurationMapRequest
*/
func (a *ElementApiService) EncodeConfigurationMap(ctx context.Context, did string, eid string) ApiEncodeConfigurationMapRequest {
	return ApiEncodeConfigurationMapRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		eid:        eid,
	}
}

// Execute executes the request
//  @return BTEncodedConfigurationInfo
func (a *ElementApiService) EncodeConfigurationMapExecute(r ApiEncodeConfigurationMapRequest) (*BTEncodedConfigurationInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTEncodedConfigurationInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ElementApiService.EncodeConfigurationMap")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/elements/d/{did}/e/{eid}/configurationencodings"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTConfigurationParams == nil {
		return localVarReturnValue, nil, reportError("bTConfigurationParams is required and must be specified")
	}

	if r.versionId != nil {
		localVarQueryParams.Add("versionId", parameterToString(*r.versionId, ""))
	}
	if r.linkDocumentId != nil {
		localVarQueryParams.Add("linkDocumentId", parameterToString(*r.linkDocumentId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bTConfigurationParams
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTEncodedConfigurationInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetConfigurationRequest struct {
	ctx        context.Context
	ApiService *ElementApiService
	did        string
	wvm        string
	wvmid      string
	eid        string
}

func (r ApiGetConfigurationRequest) Execute() (*BTConfigurationResponse2019, *http.Response, error) {
	return r.ApiService.GetConfigurationExecute(r)
}

/*
GetConfiguration Retrieve configuration by document ID, workspace or version or microversion ID, and tab ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @param wvm
 @param wvmid
 @param eid
 @return ApiGetConfigurationRequest
*/
func (a *ElementApiService) GetConfiguration(ctx context.Context, did string, wvm string, wvmid string, eid string) ApiGetConfigurationRequest {
	return ApiGetConfigurationRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wvm:        wvm,
		wvmid:      wvmid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return BTConfigurationResponse2019
func (a *ElementApiService) GetConfigurationExecute(r ApiGetConfigurationRequest) (*BTConfigurationResponse2019, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTConfigurationResponse2019
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ElementApiService.GetConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/elements/d/{did}/{wvm}/{wvmid}/e/{eid}/configuration"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvm"+"}", url.PathEscape(parameterToString(r.wvm, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvmid"+"}", url.PathEscape(parameterToString(r.wvmid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTConfigurationResponse2019
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetElementTranslatorFormatsByVersionOrWorkspaceRequest struct {
	ctx            context.Context
	ApiService     *ElementApiService
	did            string
	wv             string
	wvid           string
	eid            string
	linkDocumentId *string
	checkContent   *bool
	configuration  *string
}

// The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both.
func (r ApiGetElementTranslatorFormatsByVersionOrWorkspaceRequest) LinkDocumentId(linkDocumentId string) ApiGetElementTranslatorFormatsByVersionOrWorkspaceRequest {
	r.linkDocumentId = &linkDocumentId
	return r
}

func (r ApiGetElementTranslatorFormatsByVersionOrWorkspaceRequest) CheckContent(checkContent bool) ApiGetElementTranslatorFormatsByVersionOrWorkspaceRequest {
	r.checkContent = &checkContent
	return r
}

func (r ApiGetElementTranslatorFormatsByVersionOrWorkspaceRequest) Configuration(configuration string) ApiGetElementTranslatorFormatsByVersionOrWorkspaceRequest {
	r.configuration = &configuration
	return r
}

func (r ApiGetElementTranslatorFormatsByVersionOrWorkspaceRequest) Execute() ([]BTModelFormatInfo, *http.Response, error) {
	return r.ApiService.GetElementTranslatorFormatsByVersionOrWorkspaceExecute(r)
}

/*
GetElementTranslatorFormatsByVersionOrWorkspace Method for GetElementTranslatorFormatsByVersionOrWorkspace

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did The id of the document in which to perform the operation.
 @param wv
 @param wvid
 @param eid The id of the element in which to perform the operation.
 @return ApiGetElementTranslatorFormatsByVersionOrWorkspaceRequest
*/
func (a *ElementApiService) GetElementTranslatorFormatsByVersionOrWorkspace(ctx context.Context, did string, wv string, wvid string, eid string) ApiGetElementTranslatorFormatsByVersionOrWorkspaceRequest {
	return ApiGetElementTranslatorFormatsByVersionOrWorkspaceRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wv:         wv,
		wvid:       wvid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return []BTModelFormatInfo
func (a *ElementApiService) GetElementTranslatorFormatsByVersionOrWorkspaceExecute(r ApiGetElementTranslatorFormatsByVersionOrWorkspaceRequest) ([]BTModelFormatInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []BTModelFormatInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ElementApiService.GetElementTranslatorFormatsByVersionOrWorkspace")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/elements/translatorFormats/{did}/{wv}/{wvid}/{eid}"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wv"+"}", url.PathEscape(parameterToString(r.wv, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvid"+"}", url.PathEscape(parameterToString(r.wvid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.linkDocumentId != nil {
		localVarQueryParams.Add("linkDocumentId", parameterToString(*r.linkDocumentId, ""))
	}
	if r.checkContent != nil {
		localVarQueryParams.Add("checkContent", parameterToString(*r.checkContent, ""))
	}
	if r.configuration != nil {
		localVarQueryParams.Add("configuration", parameterToString(*r.configuration, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v []BTModelFormatInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateConfigurationRequest struct {
	ctx                           context.Context
	ApiService                    *ElementApiService
	did                           string
	wvm                           string
	wvmid                         string
	eid                           string
	bTConfigurationUpdateCall2933 *BTConfigurationUpdateCall2933
}

func (r ApiUpdateConfigurationRequest) BTConfigurationUpdateCall2933(bTConfigurationUpdateCall2933 BTConfigurationUpdateCall2933) ApiUpdateConfigurationRequest {
	r.bTConfigurationUpdateCall2933 = &bTConfigurationUpdateCall2933
	return r
}

func (r ApiUpdateConfigurationRequest) Execute() (*BTConfigurationResponse2019, *http.Response, error) {
	return r.ApiService.UpdateConfigurationExecute(r)
}

/*
UpdateConfiguration Update configuration by document ID, workspace or microversion ID, and tab ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @param wvm
 @param wvmid
 @param eid
 @return ApiUpdateConfigurationRequest
*/
func (a *ElementApiService) UpdateConfiguration(ctx context.Context, did string, wvm string, wvmid string, eid string) ApiUpdateConfigurationRequest {
	return ApiUpdateConfigurationRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wvm:        wvm,
		wvmid:      wvmid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return BTConfigurationResponse2019
func (a *ElementApiService) UpdateConfigurationExecute(r ApiUpdateConfigurationRequest) (*BTConfigurationResponse2019, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTConfigurationResponse2019
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ElementApiService.UpdateConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/elements/d/{did}/{wvm}/{wvmid}/e/{eid}/configuration"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvm"+"}", url.PathEscape(parameterToString(r.wvm, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvmid"+"}", url.PathEscape(parameterToString(r.wvmid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bTConfigurationUpdateCall2933
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTConfigurationResponse2019
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateReferencesRequest struct {
	ctx                     context.Context
	ApiService              *ElementApiService
	did                     string
	wid                     string
	eid                     string
	bTUpdateReferenceParams *BTUpdateReferenceParams
}

func (r ApiUpdateReferencesRequest) BTUpdateReferenceParams(bTUpdateReferenceParams BTUpdateReferenceParams) ApiUpdateReferencesRequest {
	r.bTUpdateReferenceParams = &bTUpdateReferenceParams
	return r
}

func (r ApiUpdateReferencesRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateReferencesExecute(r)
}

/*
UpdateReferences Update or replace node references by document ID, workspace ID, and tab ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @param wid
 @param eid
 @return ApiUpdateReferencesRequest
*/
func (a *ElementApiService) UpdateReferences(ctx context.Context, did string, wid string, eid string) ApiUpdateReferencesRequest {
	return ApiUpdateReferencesRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wid:        wid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ElementApiService) UpdateReferencesExecute(r ApiUpdateReferencesRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ElementApiService.UpdateReferences")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/elements/d/{did}/w/{wid}/e/{eid}/updatereferences"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wid"+"}", url.PathEscape(parameterToString(r.wid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTUpdateReferenceParams == nil {
		return localVarReturnValue, nil, reportError("bTUpdateReferenceParams is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bTUpdateReferenceParams
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v map[string]interface{}
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
