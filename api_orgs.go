/*
Mist API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
Contact: api@mist.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mist_sdk

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)


// OrgsApiService OrgsApi service
type OrgsApiService service

type ApiGetMemberRequest struct {
	ctx context.Context
	ApiService *OrgsApiService
	org string
	member string
	only *string
}

// Only return these fields
func (r ApiGetMemberRequest) Only(only string) ApiGetMemberRequest {
	r.only = &only
	return r
}

func (r ApiGetMemberRequest) Execute() (*GetOrgMemberResponse, *http.Response, error) {
	return r.ApiService.GetMemberExecute(r)
}

/*
GetMember Get Org

Get details about target member

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param org
 @param member
 @return ApiGetMemberRequest
*/
func (a *OrgsApiService) GetMember(ctx context.Context, org string, member string) ApiGetMemberRequest {
	return ApiGetMemberRequest{
		ApiService: a,
		ctx: ctx,
		org: org,
		member: member,
	}
}

// Execute executes the request
//  @return GetOrgMemberResponse
func (a *OrgsApiService) GetMemberExecute(r ApiGetMemberRequest) (*GetOrgMemberResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetOrgMemberResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrgsApiService.GetMember")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/orgs/{org}/members/{member}"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", url.PathEscape(parameterToString(r.org, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"member"+"}", url.PathEscape(parameterToString(r.member, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetOrgRequest struct {
	ctx context.Context
	ApiService *OrgsApiService
	org string
	only *string
	deref *string
}

// Only return these fields
func (r ApiGetOrgRequest) Only(only string) ApiGetOrgRequest {
	r.only = &only
	return r
}

// Dereference foreign keys
func (r ApiGetOrgRequest) Deref(deref string) ApiGetOrgRequest {
	r.deref = &deref
	return r
}

func (r ApiGetOrgRequest) Execute() (*GetOrgResponse, *http.Response, error) {
	return r.ApiService.GetOrgExecute(r)
}

/*
GetOrg Get Org

Get details about target org

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param org
 @return ApiGetOrgRequest
*/
func (a *OrgsApiService) GetOrg(ctx context.Context, org string) ApiGetOrgRequest {
	return ApiGetOrgRequest{
		ApiService: a,
		ctx: ctx,
		org: org,
	}
}

// Execute executes the request
//  @return GetOrgResponse
func (a *OrgsApiService) GetOrgExecute(r ApiGetOrgRequest) (*GetOrgResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetOrgResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrgsApiService.GetOrg")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/orgs/{org}"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", url.PathEscape(parameterToString(r.org, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListOrgMembersRequest struct {
	ctx context.Context
	ApiService *OrgsApiService
	org string
	search *string
	sort *string
	start *string
	limit *int32
	only *string
	at *time.Time
}

// Only return results matching search filter
func (r ApiListOrgMembersRequest) Search(search string) ApiListOrgMembersRequest {
	r.search = &search
	return r
}

// Order results by
func (r ApiListOrgMembersRequest) Sort(sort string) ApiListOrgMembersRequest {
	r.sort = &sort
	return r
}

// Start results from index or id
func (r ApiListOrgMembersRequest) Start(start string) ApiListOrgMembersRequest {
	r.start = &start
	return r
}

// Limit number of results, 1000 max
func (r ApiListOrgMembersRequest) Limit(limit int32) ApiListOrgMembersRequest {
	r.limit = &limit
	return r
}

// Only return these fields
func (r ApiListOrgMembersRequest) Only(only string) ApiListOrgMembersRequest {
	r.only = &only
	return r
}

// Limit results by specific datetime. Return resources created before or at, or deleted after or at, given datetime.
func (r ApiListOrgMembersRequest) At(at time.Time) ApiListOrgMembersRequest {
	r.at = &at
	return r
}

func (r ApiListOrgMembersRequest) Execute() (*ListOrgMembersResponse, *http.Response, error) {
	return r.ApiService.ListOrgMembersExecute(r)
}

/*
ListOrgMembers List org members

List org members owned by the requester. The requester must be a member of the org.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param org
 @return ApiListOrgMembersRequest
*/
func (a *OrgsApiService) ListOrgMembers(ctx context.Context, org string) ApiListOrgMembersRequest {
	return ApiListOrgMembersRequest{
		ApiService: a,
		ctx: ctx,
		org: org,
	}
}

// Execute executes the request
//  @return ListOrgMembersResponse
func (a *OrgsApiService) ListOrgMembersExecute(r ApiListOrgMembersRequest) (*ListOrgMembersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListOrgMembersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrgsApiService.ListOrgMembers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/orgs/{org}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", url.PathEscape(parameterToString(r.org, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.at != nil {
		localVarQueryParams.Add("at", parameterToString(*r.at, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListOrgTeamsRequest struct {
	ctx context.Context
	ApiService *OrgsApiService
	org string
	search *string
	sort *string
	start *string
	limit *int32
	only *string
	deref *string
	at *time.Time
}

// Only return results matching search filter
func (r ApiListOrgTeamsRequest) Search(search string) ApiListOrgTeamsRequest {
	r.search = &search
	return r
}

// Order results by
func (r ApiListOrgTeamsRequest) Sort(sort string) ApiListOrgTeamsRequest {
	r.sort = &sort
	return r
}

// Start results from index or id
func (r ApiListOrgTeamsRequest) Start(start string) ApiListOrgTeamsRequest {
	r.start = &start
	return r
}

// Limit number of results, 1000 max
func (r ApiListOrgTeamsRequest) Limit(limit int32) ApiListOrgTeamsRequest {
	r.limit = &limit
	return r
}

// Only return these fields
func (r ApiListOrgTeamsRequest) Only(only string) ApiListOrgTeamsRequest {
	r.only = &only
	return r
}

// Dereference foreign keys
func (r ApiListOrgTeamsRequest) Deref(deref string) ApiListOrgTeamsRequest {
	r.deref = &deref
	return r
}

// Limit results by specific datetime. Return resources created before or at, or deleted after or at, given datetime.
func (r ApiListOrgTeamsRequest) At(at time.Time) ApiListOrgTeamsRequest {
	r.at = &at
	return r
}

func (r ApiListOrgTeamsRequest) Execute() (*ListOrgTeamsResponse, *http.Response, error) {
	return r.ApiService.ListOrgTeamsExecute(r)
}

/*
ListOrgTeams List org teams

List teams in org.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param org Organization id
 @return ApiListOrgTeamsRequest
*/
func (a *OrgsApiService) ListOrgTeams(ctx context.Context, org string) ApiListOrgTeamsRequest {
	return ApiListOrgTeamsRequest{
		ApiService: a,
		ctx: ctx,
		org: org,
	}
}

// Execute executes the request
//  @return ListOrgTeamsResponse
func (a *OrgsApiService) ListOrgTeamsExecute(r ApiListOrgTeamsRequest) (*ListOrgTeamsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListOrgTeamsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrgsApiService.ListOrgTeams")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/orgs/{org}/teams"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", url.PathEscape(parameterToString(r.org, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.at != nil {
		localVarQueryParams.Add("at", parameterToString(*r.at, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListOrgsRequest struct {
	ctx context.Context
	ApiService *OrgsApiService
	allorgs *string
	search *string
	sort *string
	start *string
	limit *int32
	only *string
	deref *string
	at *time.Time
}

// Return all existing organizations
func (r ApiListOrgsRequest) Allorgs(allorgs string) ApiListOrgsRequest {
	r.allorgs = &allorgs
	return r
}

// Only return results matching search filter
func (r ApiListOrgsRequest) Search(search string) ApiListOrgsRequest {
	r.search = &search
	return r
}

// Order results by
func (r ApiListOrgsRequest) Sort(sort string) ApiListOrgsRequest {
	r.sort = &sort
	return r
}

// Start results from index or id
func (r ApiListOrgsRequest) Start(start string) ApiListOrgsRequest {
	r.start = &start
	return r
}

// Limit number of results, 100 max
func (r ApiListOrgsRequest) Limit(limit int32) ApiListOrgsRequest {
	r.limit = &limit
	return r
}

// Only return these fields
func (r ApiListOrgsRequest) Only(only string) ApiListOrgsRequest {
	r.only = &only
	return r
}

// Dereference foreign keys
func (r ApiListOrgsRequest) Deref(deref string) ApiListOrgsRequest {
	r.deref = &deref
	return r
}

// Limit results by specific datetime. Return resources created before or at, or deleted after or at, given datetime.
func (r ApiListOrgsRequest) At(at time.Time) ApiListOrgsRequest {
	r.at = &at
	return r
}

func (r ApiListOrgsRequest) Execute() (*ListOrgsResponse, *http.Response, error) {
	return r.ApiService.ListOrgsExecute(r)
}

/*
ListOrgs List orgs

List orgs owned by the requester. If parameter allorgs is true and requester is an admin then all orgs will be listed.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListOrgsRequest
*/
func (a *OrgsApiService) ListOrgs(ctx context.Context) ApiListOrgsRequest {
	return ApiListOrgsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListOrgsResponse
func (a *OrgsApiService) ListOrgsExecute(r ApiListOrgsRequest) (*ListOrgsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListOrgsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrgsApiService.ListOrgs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/orgs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.at != nil {
		localVarQueryParams.Add("at", parameterToString(*r.at, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
