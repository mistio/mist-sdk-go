/*
 * Mist API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 * Contact: api@mist.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mist_sdk

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// OrgsApiService OrgsApi service
type OrgsApiService service

type ApiGetMemberRequest struct {
	ctx _context.Context
	ApiService *OrgsApiService
	org string
	member string
	only *string
}

func (r ApiGetMemberRequest) Only(only string) ApiGetMemberRequest {
	r.only = &only
	return r
}

func (r ApiGetMemberRequest) Execute() (GetOrgMemberResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMemberExecute(r)
}

/*
 * GetMember Get Org
 * Get details about target member
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param org
 * @param member
 * @return ApiGetMemberRequest
 */
func (a *OrgsApiService) GetMember(ctx _context.Context, org string, member string) ApiGetMemberRequest {
	return ApiGetMemberRequest{
		ApiService: a,
		ctx: ctx,
		org: org,
		member: member,
	}
}

/*
 * Execute executes the request
 * @return GetOrgMemberResponse
 */
func (a *OrgsApiService) GetMemberExecute(r ApiGetMemberRequest) (GetOrgMemberResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetOrgMemberResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrgsApiService.GetMember")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/orgs/{org}/members/{member}"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", _neturl.PathEscape(parameterToString(r.org, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"member"+"}", _neturl.PathEscape(parameterToString(r.member, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.only != nil {
		localVarQueryParams.Add("only", parameterToString(*r.only, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetOrgRequest struct {
	ctx _context.Context
	ApiService *OrgsApiService
	org string
	only *string
	deref *string
}

func (r ApiGetOrgRequest) Only(only string) ApiGetOrgRequest {
	r.only = &only
	return r
}
func (r ApiGetOrgRequest) Deref(deref string) ApiGetOrgRequest {
	r.deref = &deref
	return r
}

func (r ApiGetOrgRequest) Execute() (GetOrgResponse, *_nethttp.Response, error) {
	return r.ApiService.GetOrgExecute(r)
}

/*
 * GetOrg Get Org
 * Get details about target org
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param org
 * @return ApiGetOrgRequest
 */
func (a *OrgsApiService) GetOrg(ctx _context.Context, org string) ApiGetOrgRequest {
	return ApiGetOrgRequest{
		ApiService: a,
		ctx: ctx,
		org: org,
	}
}

/*
 * Execute executes the request
 * @return GetOrgResponse
 */
func (a *OrgsApiService) GetOrgExecute(r ApiGetOrgRequest) (GetOrgResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetOrgResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrgsApiService.GetOrg")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/orgs/{org}"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", _neturl.PathEscape(parameterToString(r.org, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.only != nil {
		localVarQueryParams.Add("only", parameterToString(*r.only, ""))
	}
	if r.deref != nil {
		localVarQueryParams.Add("deref", parameterToString(*r.deref, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListOrgMembersRequest struct {
	ctx _context.Context
	ApiService *OrgsApiService
	org string
	search *string
	sort *string
	start *string
	limit *int32
	only *string
}

func (r ApiListOrgMembersRequest) Search(search string) ApiListOrgMembersRequest {
	r.search = &search
	return r
}
func (r ApiListOrgMembersRequest) Sort(sort string) ApiListOrgMembersRequest {
	r.sort = &sort
	return r
}
func (r ApiListOrgMembersRequest) Start(start string) ApiListOrgMembersRequest {
	r.start = &start
	return r
}
func (r ApiListOrgMembersRequest) Limit(limit int32) ApiListOrgMembersRequest {
	r.limit = &limit
	return r
}
func (r ApiListOrgMembersRequest) Only(only string) ApiListOrgMembersRequest {
	r.only = &only
	return r
}

func (r ApiListOrgMembersRequest) Execute() (ListOrgMembersResponse, *_nethttp.Response, error) {
	return r.ApiService.ListOrgMembersExecute(r)
}

/*
 * ListOrgMembers List org members
 * List org members owned by the requester. The requester must be a member of the org.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param org
 * @return ApiListOrgMembersRequest
 */
func (a *OrgsApiService) ListOrgMembers(ctx _context.Context, org string) ApiListOrgMembersRequest {
	return ApiListOrgMembersRequest{
		ApiService: a,
		ctx: ctx,
		org: org,
	}
}

/*
 * Execute executes the request
 * @return ListOrgMembersResponse
 */
func (a *OrgsApiService) ListOrgMembersExecute(r ApiListOrgMembersRequest) (ListOrgMembersResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListOrgMembersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrgsApiService.ListOrgMembers")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/orgs/{org}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", _neturl.PathEscape(parameterToString(r.org, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.start != nil {
		localVarQueryParams.Add("start", parameterToString(*r.start, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.only != nil {
		localVarQueryParams.Add("only", parameterToString(*r.only, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListOrgTeamsRequest struct {
	ctx _context.Context
	ApiService *OrgsApiService
	org string
	search *string
	sort *string
	start *string
	limit *int32
	only *string
	deref *string
}

func (r ApiListOrgTeamsRequest) Search(search string) ApiListOrgTeamsRequest {
	r.search = &search
	return r
}
func (r ApiListOrgTeamsRequest) Sort(sort string) ApiListOrgTeamsRequest {
	r.sort = &sort
	return r
}
func (r ApiListOrgTeamsRequest) Start(start string) ApiListOrgTeamsRequest {
	r.start = &start
	return r
}
func (r ApiListOrgTeamsRequest) Limit(limit int32) ApiListOrgTeamsRequest {
	r.limit = &limit
	return r
}
func (r ApiListOrgTeamsRequest) Only(only string) ApiListOrgTeamsRequest {
	r.only = &only
	return r
}
func (r ApiListOrgTeamsRequest) Deref(deref string) ApiListOrgTeamsRequest {
	r.deref = &deref
	return r
}

func (r ApiListOrgTeamsRequest) Execute() (ListOrgTeamsResponse, *_nethttp.Response, error) {
	return r.ApiService.ListOrgTeamsExecute(r)
}

/*
 * ListOrgTeams List org teams
 * List teams in org.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param org Organization id
 * @return ApiListOrgTeamsRequest
 */
func (a *OrgsApiService) ListOrgTeams(ctx _context.Context, org string) ApiListOrgTeamsRequest {
	return ApiListOrgTeamsRequest{
		ApiService: a,
		ctx: ctx,
		org: org,
	}
}

/*
 * Execute executes the request
 * @return ListOrgTeamsResponse
 */
func (a *OrgsApiService) ListOrgTeamsExecute(r ApiListOrgTeamsRequest) (ListOrgTeamsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListOrgTeamsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrgsApiService.ListOrgTeams")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/orgs/{org}/teams"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", _neturl.PathEscape(parameterToString(r.org, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.start != nil {
		localVarQueryParams.Add("start", parameterToString(*r.start, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.only != nil {
		localVarQueryParams.Add("only", parameterToString(*r.only, ""))
	}
	if r.deref != nil {
		localVarQueryParams.Add("deref", parameterToString(*r.deref, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListOrgsRequest struct {
	ctx _context.Context
	ApiService *OrgsApiService
	allorgs *string
	search *string
	sort *string
	start *string
	limit *int32
	only *string
	deref *string
}

func (r ApiListOrgsRequest) Allorgs(allorgs string) ApiListOrgsRequest {
	r.allorgs = &allorgs
	return r
}
func (r ApiListOrgsRequest) Search(search string) ApiListOrgsRequest {
	r.search = &search
	return r
}
func (r ApiListOrgsRequest) Sort(sort string) ApiListOrgsRequest {
	r.sort = &sort
	return r
}
func (r ApiListOrgsRequest) Start(start string) ApiListOrgsRequest {
	r.start = &start
	return r
}
func (r ApiListOrgsRequest) Limit(limit int32) ApiListOrgsRequest {
	r.limit = &limit
	return r
}
func (r ApiListOrgsRequest) Only(only string) ApiListOrgsRequest {
	r.only = &only
	return r
}
func (r ApiListOrgsRequest) Deref(deref string) ApiListOrgsRequest {
	r.deref = &deref
	return r
}

func (r ApiListOrgsRequest) Execute() (ListOrgsResponse, *_nethttp.Response, error) {
	return r.ApiService.ListOrgsExecute(r)
}

/*
 * ListOrgs List orgs
 * List orgs owned by the requester. If parameter allorgs is true and requester is an admin then all orgs will be listed.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListOrgsRequest
 */
func (a *OrgsApiService) ListOrgs(ctx _context.Context) ApiListOrgsRequest {
	return ApiListOrgsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ListOrgsResponse
 */
func (a *OrgsApiService) ListOrgsExecute(r ApiListOrgsRequest) (ListOrgsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListOrgsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrgsApiService.ListOrgs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/orgs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.allorgs != nil {
		localVarQueryParams.Add("allorgs", parameterToString(*r.allorgs, ""))
	}
	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.start != nil {
		localVarQueryParams.Add("start", parameterToString(*r.start, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.only != nil {
		localVarQueryParams.Add("only", parameterToString(*r.only, ""))
	}
	if r.deref != nil {
		localVarQueryParams.Add("deref", parameterToString(*r.deref, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}