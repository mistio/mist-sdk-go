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


// ImagesApiService ImagesApi service
type ImagesApiService service

type ApiGetImageRequest struct {
	ctx context.Context
	ApiService *ImagesApiService
	image string
	only *string
	deref *string
}

// Only return these fields
func (r ApiGetImageRequest) Only(only string) ApiGetImageRequest {
	r.only = &only
	return r
}

// Dereference foreign keys
func (r ApiGetImageRequest) Deref(deref string) ApiGetImageRequest {
	r.deref = &deref
	return r
}

func (r ApiGetImageRequest) Execute() (*GetImageResponse, *http.Response, error) {
	return r.ApiService.GetImageExecute(r)
}

/*
GetImage Get image

Get details about target image

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param image
 @return ApiGetImageRequest
*/
func (a *ImagesApiService) GetImage(ctx context.Context, image string) ApiGetImageRequest {
	return ApiGetImageRequest{
		ApiService: a,
		ctx: ctx,
		image: image,
	}
}

// Execute executes the request
//  @return GetImageResponse
func (a *ImagesApiService) GetImageExecute(r ApiGetImageRequest) (*GetImageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetImageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImagesApiService.GetImage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/images/{image}"
	localVarPath = strings.Replace(localVarPath, "{"+"image"+"}", url.PathEscape(parameterToString(r.image, "")), -1)

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

type ApiListImagesRequest struct {
	ctx context.Context
	ApiService *ImagesApiService
	cloud *string
	search *string
	sort *string
	start *string
	limit *int32
	only *string
	deref *string
	at *time.Time
}

func (r ApiListImagesRequest) Cloud(cloud string) ApiListImagesRequest {
	r.cloud = &cloud
	return r
}

// Only return results matching search filter
func (r ApiListImagesRequest) Search(search string) ApiListImagesRequest {
	r.search = &search
	return r
}

// Order results by
func (r ApiListImagesRequest) Sort(sort string) ApiListImagesRequest {
	r.sort = &sort
	return r
}

// Start results from index or id
func (r ApiListImagesRequest) Start(start string) ApiListImagesRequest {
	r.start = &start
	return r
}

// Limit number of results, 1000 max
func (r ApiListImagesRequest) Limit(limit int32) ApiListImagesRequest {
	r.limit = &limit
	return r
}

// Only return these fields
func (r ApiListImagesRequest) Only(only string) ApiListImagesRequest {
	r.only = &only
	return r
}

// Dereference foreign keys
func (r ApiListImagesRequest) Deref(deref string) ApiListImagesRequest {
	r.deref = &deref
	return r
}

// Limit results by specific datetime. Return resources created before or at, or deleted after or at, given datetime.
func (r ApiListImagesRequest) At(at time.Time) ApiListImagesRequest {
	r.at = &at
	return r
}

func (r ApiListImagesRequest) Execute() (*ListImagesResponse, *http.Response, error) {
	return r.ApiService.ListImagesExecute(r)
}

/*
ListImages List images

List images owned by the active org. READ permission required on image & cloud.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListImagesRequest
*/
func (a *ImagesApiService) ListImages(ctx context.Context) ApiListImagesRequest {
	return ApiListImagesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListImagesResponse
func (a *ImagesApiService) ListImagesExecute(r ApiListImagesRequest) (*ListImagesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListImagesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImagesApiService.ListImages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/images"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
