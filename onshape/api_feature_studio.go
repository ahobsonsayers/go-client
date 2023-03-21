/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.161.13200-ff216a970a02
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

// FeatureStudioApiService FeatureStudioApi service
type FeatureStudioApiService service

type ApiCreateFeatureStudioRequest struct {
	ctx                  context.Context
	ApiService           *FeatureStudioApiService
	did                  string
	wid                  string
	bTModelElementParams *BTModelElementParams
}

func (r ApiCreateFeatureStudioRequest) BTModelElementParams(bTModelElementParams BTModelElementParams) ApiCreateFeatureStudioRequest {
	r.bTModelElementParams = &bTModelElementParams
	return r
}

func (r ApiCreateFeatureStudioRequest) Execute() (*BTDocumentElementInfo, *http.Response, error) {
	return r.ApiService.CreateFeatureStudioExecute(r)
}

/*
CreateFeatureStudio Create Feature Studio by document ID and workspace ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @param wid
 @return ApiCreateFeatureStudioRequest
*/
func (a *FeatureStudioApiService) CreateFeatureStudio(ctx context.Context, did string, wid string) ApiCreateFeatureStudioRequest {
	return ApiCreateFeatureStudioRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wid:        wid,
	}
}

// Execute executes the request
//  @return BTDocumentElementInfo
func (a *FeatureStudioApiService) CreateFeatureStudioExecute(r ApiCreateFeatureStudioRequest) (*BTDocumentElementInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTDocumentElementInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureStudioApiService.CreateFeatureStudio")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/featurestudios/d/{did}/w/{wid}"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wid"+"}", url.PathEscape(parameterToString(r.wid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTModelElementParams == nil {
		return localVarReturnValue, nil, reportError("bTModelElementParams is required and must be specified")
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
	localVarPostBody = r.bTModelElementParams
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

type ApiGetFeatureStudioContentsRequest struct {
	ctx        context.Context
	ApiService *FeatureStudioApiService
	did        string
	wvm        string
	wvmid      string
	eid        string
}

func (r ApiGetFeatureStudioContentsRequest) Execute() (*BTFeatureStudioContents2239, *http.Response, error) {
	return r.ApiService.GetFeatureStudioContentsExecute(r)
}

/*
GetFeatureStudioContents Method for GetFeatureStudioContents

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @param wvm
 @param wvmid
 @param eid
 @return ApiGetFeatureStudioContentsRequest
*/
func (a *FeatureStudioApiService) GetFeatureStudioContents(ctx context.Context, did string, wvm string, wvmid string, eid string) ApiGetFeatureStudioContentsRequest {
	return ApiGetFeatureStudioContentsRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wvm:        wvm,
		wvmid:      wvmid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return BTFeatureStudioContents2239
func (a *FeatureStudioApiService) GetFeatureStudioContentsExecute(r ApiGetFeatureStudioContentsRequest) (*BTFeatureStudioContents2239, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTFeatureStudioContents2239
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureStudioApiService.GetFeatureStudioContents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/featurestudios/d/{did}/{wvm}/{wvmid}/e/{eid}"
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
		var v BTFeatureStudioContents2239
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

type ApiGetFeatureStudioSpecsRequest struct {
	ctx        context.Context
	ApiService *FeatureStudioApiService
	did        string
	wvm        string
	wvmid      string
	eid        string
}

func (r ApiGetFeatureStudioSpecsRequest) Execute() (*BTFeatureSpecsResponse664, *http.Response, error) {
	return r.ApiService.GetFeatureStudioSpecsExecute(r)
}

/*
GetFeatureStudioSpecs Retrieve Feature Studio specs by document ID, workspace or version or microversion ID, and tab ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @param wvm
 @param wvmid
 @param eid
 @return ApiGetFeatureStudioSpecsRequest
*/
func (a *FeatureStudioApiService) GetFeatureStudioSpecs(ctx context.Context, did string, wvm string, wvmid string, eid string) ApiGetFeatureStudioSpecsRequest {
	return ApiGetFeatureStudioSpecsRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wvm:        wvm,
		wvmid:      wvmid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return BTFeatureSpecsResponse664
func (a *FeatureStudioApiService) GetFeatureStudioSpecsExecute(r ApiGetFeatureStudioSpecsRequest) (*BTFeatureSpecsResponse664, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTFeatureSpecsResponse664
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureStudioApiService.GetFeatureStudioSpecs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/featurestudios/d/{did}/{wvm}/{wvmid}/e/{eid}/featurespecs"
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
		var v BTFeatureSpecsResponse664
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

type ApiUpdateFeatureStudioContentsRequest struct {
	ctx                         context.Context
	ApiService                  *FeatureStudioApiService
	did                         string
	wvm                         string
	wvmid                       string
	eid                         string
	bTFeatureStudioContents2239 *BTFeatureStudioContents2239
}

func (r ApiUpdateFeatureStudioContentsRequest) BTFeatureStudioContents2239(bTFeatureStudioContents2239 BTFeatureStudioContents2239) ApiUpdateFeatureStudioContentsRequest {
	r.bTFeatureStudioContents2239 = &bTFeatureStudioContents2239
	return r
}

func (r ApiUpdateFeatureStudioContentsRequest) Execute() (*BTFeatureStudioContents2239, *http.Response, error) {
	return r.ApiService.UpdateFeatureStudioContentsExecute(r)
}

/*
UpdateFeatureStudioContents Update Feature Studio contents by document ID, workspace or version or microversion ID, and tab ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @param wvm
 @param wvmid
 @param eid
 @return ApiUpdateFeatureStudioContentsRequest
*/
func (a *FeatureStudioApiService) UpdateFeatureStudioContents(ctx context.Context, did string, wvm string, wvmid string, eid string) ApiUpdateFeatureStudioContentsRequest {
	return ApiUpdateFeatureStudioContentsRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wvm:        wvm,
		wvmid:      wvmid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return BTFeatureStudioContents2239
func (a *FeatureStudioApiService) UpdateFeatureStudioContentsExecute(r ApiUpdateFeatureStudioContentsRequest) (*BTFeatureStudioContents2239, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTFeatureStudioContents2239
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureStudioApiService.UpdateFeatureStudioContents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/featurestudios/d/{did}/{wvm}/{wvmid}/e/{eid}"
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
	localVarPostBody = r.bTFeatureStudioContents2239
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
		var v BTFeatureStudioContents2239
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
