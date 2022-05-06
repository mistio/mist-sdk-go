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

// SizesApiService SizesApi service
type SizesApiService service

type ApiGetSizeRequest struct {
	ctx _context.Context
	ApiService *SizesApiService
	size string
	only *string
	deref *string
}

func (r ApiGetSizeRequest) Only(only string) ApiGetSizeRequest {
	r.only = &only
	return r
}
func (r ApiGetSizeRequest) Deref(deref string) ApiGetSizeRequest {
	r.deref = &deref
	return r
}

func (r ApiGetSizeRequest) Execute() (GetSizeResponse, *_nethttp.Response, error) {
	return r.ApiService.GetSizeExecute(r)
}

/*
 * GetSize Get size
 * Get details about target size
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param size
 * @return ApiGetSizeRequest
 */
func (a *SizesApiService) GetSize(ctx _context.Context, size string) ApiGetSizeRequest {
	return ApiGetSizeRequest{
		ApiService: a,
		ctx: ctx,
		size: size,
	}
}

/*
 * Execute executes the request
 * @return GetSizeResponse
 */
func (a *SizesApiService) GetSizeExecute(r ApiGetSizeRequest) (GetSizeResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetSizeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SizesApiService.GetSize")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sizes/{size}"
	localVarPath = strings.Replace(localVarPath, "{"+"size"+"}", _neturl.PathEscape(parameterToString(r.size, "")), -1)

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

type ApiListSizesRequest struct {
	ctx _context.Context
	ApiService *SizesApiService
	cloud *string
	search *string
	sort *string
	start *string
	limit *int32
	only *string
	deref *string
}

func (r ApiListSizesRequest) Cloud(cloud string) ApiListSizesRequest {
	r.cloud = &cloud
	return r
}
func (r ApiListSizesRequest) Search(search string) ApiListSizesRequest {
	r.search = &search
	return r
}
func (r ApiListSizesRequest) Sort(sort string) ApiListSizesRequest {
	r.sort = &sort
	return r
}
func (r ApiListSizesRequest) Start(start string) ApiListSizesRequest {
	r.start = &start
	return r
}
func (r ApiListSizesRequest) Limit(limit int32) ApiListSizesRequest {
	r.limit = &limit
	return r
}
func (r ApiListSizesRequest) Only(only string) ApiListSizesRequest {
	r.only = &only
	return r
}
func (r ApiListSizesRequest) Deref(deref string) ApiListSizesRequest {
	r.deref = &deref
	return r
}

func (r ApiListSizesRequest) Execute() (ListSizesResponse, *_nethttp.Response, error) {
	return r.ApiService.ListSizesExecute(r)
}

/*
 * ListSizes List sizes
 * List sizes owned by the active org. READ permission required on size & cloud.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListSizesRequest
 */
func (a *SizesApiService) ListSizes(ctx _context.Context) ApiListSizesRequest {
	return ApiListSizesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ListSizesResponse
 */
func (a *SizesApiService) ListSizesExecute(r ApiListSizesRequest) (ListSizesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListSizesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SizesApiService.ListSizes")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sizes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.cloud != nil {
		localVarQueryParams.Add("cloud", parameterToString(*r.cloud, ""))
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
