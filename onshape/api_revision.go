/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5944-34bf93ccfb05
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

// RevisionApiService RevisionApi service
type RevisionApiService service

type ApiDeleteRevisionHistoryRequest struct {
	ctx                   context.Context
	ApiService            *RevisionApiService
	cid                   string
	pnum                  string
	et                    string
	ignoreLinkedDocuments *bool
}

func (r ApiDeleteRevisionHistoryRequest) IgnoreLinkedDocuments(ignoreLinkedDocuments bool) ApiDeleteRevisionHistoryRequest {
	r.ignoreLinkedDocuments = &ignoreLinkedDocuments
	return r
}

func (r ApiDeleteRevisionHistoryRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteRevisionHistoryExecute(r)
}

/*
DeleteRevisionHistory Method for DeleteRevisionHistory

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param cid
 @param pnum
 @param et
 @return ApiDeleteRevisionHistoryRequest
*/
func (a *RevisionApiService) DeleteRevisionHistory(ctx context.Context, cid string, pnum string, et string) ApiDeleteRevisionHistoryRequest {
	return ApiDeleteRevisionHistoryRequest{
		ApiService: a,
		ctx:        ctx,
		cid:        cid,
		pnum:       pnum,
		et:         et,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RevisionApiService) DeleteRevisionHistoryExecute(r ApiDeleteRevisionHistoryRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RevisionApiService.DeleteRevisionHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/revisions/companies/{cid}/partnumber/{pnum}/elementType/{et}"
	localVarPath = strings.Replace(localVarPath, "{"+"cid"+"}", url.PathEscape(parameterToString(r.cid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pnum"+"}", url.PathEscape(parameterToString(r.pnum, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"et"+"}", url.PathEscape(parameterToString(r.et, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.ignoreLinkedDocuments != nil {
		localVarQueryParams.Add("ignoreLinkedDocuments", parameterToString(*r.ignoreLinkedDocuments, ""))
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

type ApiEnumerateRevisionsRequest struct {
	ctx         context.Context
	ApiService  *RevisionApiService
	cid         string
	elementType *int32
	limit       *int32
	offset      *int32
	latestOnly  *bool
	after       *JSONTime
}

func (r ApiEnumerateRevisionsRequest) ElementType(elementType int32) ApiEnumerateRevisionsRequest {
	r.elementType = &elementType
	return r
}

func (r ApiEnumerateRevisionsRequest) Limit(limit int32) ApiEnumerateRevisionsRequest {
	r.limit = &limit
	return r
}

func (r ApiEnumerateRevisionsRequest) Offset(offset int32) ApiEnumerateRevisionsRequest {
	r.offset = &offset
	return r
}

func (r ApiEnumerateRevisionsRequest) LatestOnly(latestOnly bool) ApiEnumerateRevisionsRequest {
	r.latestOnly = &latestOnly
	return r
}

func (r ApiEnumerateRevisionsRequest) After(after JSONTime) ApiEnumerateRevisionsRequest {
	r.after = &after
	return r
}

func (r ApiEnumerateRevisionsRequest) Execute() (*BTListResponseBTRevisionInfo, *http.Response, error) {
	return r.ApiService.EnumerateRevisionsExecute(r)
}

/*
EnumerateRevisions Enumerate all revisions released in a company by company ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param cid
 @return ApiEnumerateRevisionsRequest
*/
func (a *RevisionApiService) EnumerateRevisions(ctx context.Context, cid string) ApiEnumerateRevisionsRequest {
	return ApiEnumerateRevisionsRequest{
		ApiService: a,
		ctx:        ctx,
		cid:        cid,
	}
}

// Execute executes the request
//  @return BTListResponseBTRevisionInfo
func (a *RevisionApiService) EnumerateRevisionsExecute(r ApiEnumerateRevisionsRequest) (*BTListResponseBTRevisionInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTListResponseBTRevisionInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RevisionApiService.EnumerateRevisions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/revisions/companies/{cid}"
	localVarPath = strings.Replace(localVarPath, "{"+"cid"+"}", url.PathEscape(parameterToString(r.cid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.elementType != nil {
		localVarQueryParams.Add("elementType", parameterToString(*r.elementType, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.latestOnly != nil {
		localVarQueryParams.Add("latestOnly", parameterToString(*r.latestOnly, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
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
		var v BTListResponseBTRevisionInfo
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

type ApiGetLatestInDocumentOrCompanyRequest struct {
	ctx        context.Context
	ApiService *RevisionApiService
	cd         string
	cdid       string
	pnum       string
	et         *string
}

// 0: Part Studio, 1: Assembly, 2: Drawing. 4: Blob
func (r ApiGetLatestInDocumentOrCompanyRequest) Et(et string) ApiGetLatestInDocumentOrCompanyRequest {
	r.et = &et
	return r
}

func (r ApiGetLatestInDocumentOrCompanyRequest) Execute() (*BTRevisionInfo, *http.Response, error) {
	return r.ApiService.GetLatestInDocumentOrCompanyExecute(r)
}

/*
GetLatestInDocumentOrCompany Retrieve latest revisions for a part number in a document or company by document ID, workspace or version or microversion ID, and tab ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param cd
 @param cdid
 @param pnum
 @return ApiGetLatestInDocumentOrCompanyRequest
*/
func (a *RevisionApiService) GetLatestInDocumentOrCompany(ctx context.Context, cd string, cdid string, pnum string) ApiGetLatestInDocumentOrCompanyRequest {
	return ApiGetLatestInDocumentOrCompanyRequest{
		ApiService: a,
		ctx:        ctx,
		cd:         cd,
		cdid:       cdid,
		pnum:       pnum,
	}
}

// Execute executes the request
//  @return BTRevisionInfo
func (a *RevisionApiService) GetLatestInDocumentOrCompanyExecute(r ApiGetLatestInDocumentOrCompanyRequest) (*BTRevisionInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTRevisionInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RevisionApiService.GetLatestInDocumentOrCompany")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/revisions/{cd}/{cdid}/p/{pnum}/latest"
	localVarPath = strings.Replace(localVarPath, "{"+"cd"+"}", url.PathEscape(parameterToString(r.cd, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cdid"+"}", url.PathEscape(parameterToString(r.cdid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pnum"+"}", url.PathEscape(parameterToString(r.pnum, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.et == nil {
		return localVarReturnValue, nil, reportError("et is required and must be specified")
	}

	localVarQueryParams.Add("et", parameterToString(*r.et, ""))
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
		var v BTRevisionInfo
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

type ApiGetRevisionByPartNumberRequest struct {
	ctx         context.Context
	ApiService  *RevisionApiService
	cid         string
	pnum        string
	revision    *string
	elementType *int32
}

// Revision
func (r ApiGetRevisionByPartNumberRequest) Revision(revision string) ApiGetRevisionByPartNumberRequest {
	r.revision = &revision
	return r
}

// 0: Part Studio, 1: Assembly, 2: Drawing. 4: Blob
func (r ApiGetRevisionByPartNumberRequest) ElementType(elementType int32) ApiGetRevisionByPartNumberRequest {
	r.elementType = &elementType
	return r
}

func (r ApiGetRevisionByPartNumberRequest) Execute() (*BTRevisionInfo, *http.Response, error) {
	return r.ApiService.GetRevisionByPartNumberExecute(r)
}

/*
GetRevisionByPartNumber Get Navigation URL

Get Navigation URL for part number and revision.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param cid Company id
 @param pnum Part Number
 @return ApiGetRevisionByPartNumberRequest
*/
func (a *RevisionApiService) GetRevisionByPartNumber(ctx context.Context, cid string, pnum string) ApiGetRevisionByPartNumberRequest {
	return ApiGetRevisionByPartNumberRequest{
		ApiService: a,
		ctx:        ctx,
		cid:        cid,
		pnum:       pnum,
	}
}

// Execute executes the request
//  @return BTRevisionInfo
func (a *RevisionApiService) GetRevisionByPartNumberExecute(r ApiGetRevisionByPartNumberRequest) (*BTRevisionInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTRevisionInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RevisionApiService.GetRevisionByPartNumber")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/revisions/c/{cid}/partnumber/{pnum}"
	localVarPath = strings.Replace(localVarPath, "{"+"cid"+"}", url.PathEscape(parameterToString(r.cid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pnum"+"}", url.PathEscape(parameterToString(r.pnum, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.revision != nil {
		localVarQueryParams.Add("revision", parameterToString(*r.revision, ""))
	}
	if r.elementType != nil {
		localVarQueryParams.Add("elementType", parameterToString(*r.elementType, ""))
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
		var v BTRevisionInfo
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

type ApiGetRevisionHistoryInCompanyByElementIdRequest struct {
	ctx                  context.Context
	ApiService           *RevisionApiService
	cid                  string
	did                  string
	wv                   string
	wvid                 string
	eid                  string
	elementType          *string
	configuration        *string
	linkDocumentId       *string
	fillApprovers        *bool
	fillExportPermission *bool
}

func (r ApiGetRevisionHistoryInCompanyByElementIdRequest) ElementType(elementType string) ApiGetRevisionHistoryInCompanyByElementIdRequest {
	r.elementType = &elementType
	return r
}

func (r ApiGetRevisionHistoryInCompanyByElementIdRequest) Configuration(configuration string) ApiGetRevisionHistoryInCompanyByElementIdRequest {
	r.configuration = &configuration
	return r
}

func (r ApiGetRevisionHistoryInCompanyByElementIdRequest) LinkDocumentId(linkDocumentId string) ApiGetRevisionHistoryInCompanyByElementIdRequest {
	r.linkDocumentId = &linkDocumentId
	return r
}

func (r ApiGetRevisionHistoryInCompanyByElementIdRequest) FillApprovers(fillApprovers bool) ApiGetRevisionHistoryInCompanyByElementIdRequest {
	r.fillApprovers = &fillApprovers
	return r
}

func (r ApiGetRevisionHistoryInCompanyByElementIdRequest) FillExportPermission(fillExportPermission bool) ApiGetRevisionHistoryInCompanyByElementIdRequest {
	r.fillExportPermission = &fillExportPermission
	return r
}

func (r ApiGetRevisionHistoryInCompanyByElementIdRequest) Execute() (*BTRevisionListResponse, *http.Response, error) {
	return r.ApiService.GetRevisionHistoryInCompanyByElementIdExecute(r)
}

/*
GetRevisionHistoryInCompanyByElementId Retrieve a list of all revisions for a part in a company by company ID, document ID, workspace or version ID, and tab ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param cid
 @param did
 @param wv
 @param wvid
 @param eid
 @return ApiGetRevisionHistoryInCompanyByElementIdRequest
*/
func (a *RevisionApiService) GetRevisionHistoryInCompanyByElementId(ctx context.Context, cid string, did string, wv string, wvid string, eid string) ApiGetRevisionHistoryInCompanyByElementIdRequest {
	return ApiGetRevisionHistoryInCompanyByElementIdRequest{
		ApiService: a,
		ctx:        ctx,
		cid:        cid,
		did:        did,
		wv:         wv,
		wvid:       wvid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return BTRevisionListResponse
func (a *RevisionApiService) GetRevisionHistoryInCompanyByElementIdExecute(r ApiGetRevisionHistoryInCompanyByElementIdRequest) (*BTRevisionListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTRevisionListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RevisionApiService.GetRevisionHistoryInCompanyByElementId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/revisions/companies/{cid}/d/{did}/{wv}/{wvid}/e/{eid}"
	localVarPath = strings.Replace(localVarPath, "{"+"cid"+"}", url.PathEscape(parameterToString(r.cid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wv"+"}", url.PathEscape(parameterToString(r.wv, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvid"+"}", url.PathEscape(parameterToString(r.wvid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.elementType == nil {
		return localVarReturnValue, nil, reportError("elementType is required and must be specified")
	}

	localVarQueryParams.Add("elementType", parameterToString(*r.elementType, ""))
	if r.configuration != nil {
		localVarQueryParams.Add("configuration", parameterToString(*r.configuration, ""))
	}
	if r.linkDocumentId != nil {
		localVarQueryParams.Add("linkDocumentId", parameterToString(*r.linkDocumentId, ""))
	}
	if r.fillApprovers != nil {
		localVarQueryParams.Add("fillApprovers", parameterToString(*r.fillApprovers, ""))
	}
	if r.fillExportPermission != nil {
		localVarQueryParams.Add("fillExportPermission", parameterToString(*r.fillExportPermission, ""))
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
		var v BTRevisionListResponse
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

type ApiGetRevisionHistoryInCompanyByPartIdRequest struct {
	ctx                  context.Context
	ApiService           *RevisionApiService
	cid                  string
	did                  string
	wv                   string
	wvid                 string
	eid                  string
	pid                  string
	configuration        *string
	linkDocumentId       *string
	fillApprovers        *bool
	fillExportPermission *bool
}

func (r ApiGetRevisionHistoryInCompanyByPartIdRequest) Configuration(configuration string) ApiGetRevisionHistoryInCompanyByPartIdRequest {
	r.configuration = &configuration
	return r
}

func (r ApiGetRevisionHistoryInCompanyByPartIdRequest) LinkDocumentId(linkDocumentId string) ApiGetRevisionHistoryInCompanyByPartIdRequest {
	r.linkDocumentId = &linkDocumentId
	return r
}

func (r ApiGetRevisionHistoryInCompanyByPartIdRequest) FillApprovers(fillApprovers bool) ApiGetRevisionHistoryInCompanyByPartIdRequest {
	r.fillApprovers = &fillApprovers
	return r
}

func (r ApiGetRevisionHistoryInCompanyByPartIdRequest) FillExportPermission(fillExportPermission bool) ApiGetRevisionHistoryInCompanyByPartIdRequest {
	r.fillExportPermission = &fillExportPermission
	return r
}

func (r ApiGetRevisionHistoryInCompanyByPartIdRequest) Execute() (*BTRevisionListResponse, *http.Response, error) {
	return r.ApiService.GetRevisionHistoryInCompanyByPartIdExecute(r)
}

/*
GetRevisionHistoryInCompanyByPartId Retrieve a list of all revisions for a part in a company by company ID, document ID, workspace or version ID, tab ID, and part ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param cid
 @param did
 @param wv
 @param wvid
 @param eid
 @param pid
 @return ApiGetRevisionHistoryInCompanyByPartIdRequest
*/
func (a *RevisionApiService) GetRevisionHistoryInCompanyByPartId(ctx context.Context, cid string, did string, wv string, wvid string, eid string, pid string) ApiGetRevisionHistoryInCompanyByPartIdRequest {
	return ApiGetRevisionHistoryInCompanyByPartIdRequest{
		ApiService: a,
		ctx:        ctx,
		cid:        cid,
		did:        did,
		wv:         wv,
		wvid:       wvid,
		eid:        eid,
		pid:        pid,
	}
}

// Execute executes the request
//  @return BTRevisionListResponse
func (a *RevisionApiService) GetRevisionHistoryInCompanyByPartIdExecute(r ApiGetRevisionHistoryInCompanyByPartIdRequest) (*BTRevisionListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTRevisionListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RevisionApiService.GetRevisionHistoryInCompanyByPartId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/revisions/companies/{cid}/d/{did}/{wv}/{wvid}/e/{eid}/p/{pid}"
	localVarPath = strings.Replace(localVarPath, "{"+"cid"+"}", url.PathEscape(parameterToString(r.cid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wv"+"}", url.PathEscape(parameterToString(r.wv, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvid"+"}", url.PathEscape(parameterToString(r.wvid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pid"+"}", url.PathEscape(parameterToString(r.pid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.configuration != nil {
		localVarQueryParams.Add("configuration", parameterToString(*r.configuration, ""))
	}
	if r.linkDocumentId != nil {
		localVarQueryParams.Add("linkDocumentId", parameterToString(*r.linkDocumentId, ""))
	}
	if r.fillApprovers != nil {
		localVarQueryParams.Add("fillApprovers", parameterToString(*r.fillApprovers, ""))
	}
	if r.fillExportPermission != nil {
		localVarQueryParams.Add("fillExportPermission", parameterToString(*r.fillExportPermission, ""))
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
		var v BTRevisionListResponse
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

type ApiGetRevisionHistoryInCompanyByPartNumberRequest struct {
	ctx                  context.Context
	ApiService           *RevisionApiService
	cid                  string
	pnum                 string
	elementType          *string
	fillApprovers        *bool
	fillExportPermission *bool
}

func (r ApiGetRevisionHistoryInCompanyByPartNumberRequest) ElementType(elementType string) ApiGetRevisionHistoryInCompanyByPartNumberRequest {
	r.elementType = &elementType
	return r
}

func (r ApiGetRevisionHistoryInCompanyByPartNumberRequest) FillApprovers(fillApprovers bool) ApiGetRevisionHistoryInCompanyByPartNumberRequest {
	r.fillApprovers = &fillApprovers
	return r
}

func (r ApiGetRevisionHistoryInCompanyByPartNumberRequest) FillExportPermission(fillExportPermission bool) ApiGetRevisionHistoryInCompanyByPartNumberRequest {
	r.fillExportPermission = &fillExportPermission
	return r
}

func (r ApiGetRevisionHistoryInCompanyByPartNumberRequest) Execute() (*BTRevisionListResponse, *http.Response, error) {
	return r.ApiService.GetRevisionHistoryInCompanyByPartNumberExecute(r)
}

/*
GetRevisionHistoryInCompanyByPartNumber Retrieve a list of all revisions for a part number in a company by company ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param cid
 @param pnum
 @return ApiGetRevisionHistoryInCompanyByPartNumberRequest
*/
func (a *RevisionApiService) GetRevisionHistoryInCompanyByPartNumber(ctx context.Context, cid string, pnum string) ApiGetRevisionHistoryInCompanyByPartNumberRequest {
	return ApiGetRevisionHistoryInCompanyByPartNumberRequest{
		ApiService: a,
		ctx:        ctx,
		cid:        cid,
		pnum:       pnum,
	}
}

// Execute executes the request
//  @return BTRevisionListResponse
func (a *RevisionApiService) GetRevisionHistoryInCompanyByPartNumberExecute(r ApiGetRevisionHistoryInCompanyByPartNumberRequest) (*BTRevisionListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTRevisionListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RevisionApiService.GetRevisionHistoryInCompanyByPartNumber")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/revisions/companies/{cid}/partnumber/{pnum}"
	localVarPath = strings.Replace(localVarPath, "{"+"cid"+"}", url.PathEscape(parameterToString(r.cid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pnum"+"}", url.PathEscape(parameterToString(r.pnum, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.elementType == nil {
		return localVarReturnValue, nil, reportError("elementType is required and must be specified")
	}

	localVarQueryParams.Add("elementType", parameterToString(*r.elementType, ""))
	if r.fillApprovers != nil {
		localVarQueryParams.Add("fillApprovers", parameterToString(*r.fillApprovers, ""))
	}
	if r.fillExportPermission != nil {
		localVarQueryParams.Add("fillExportPermission", parameterToString(*r.fillExportPermission, ""))
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
		var v BTRevisionListResponse
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
