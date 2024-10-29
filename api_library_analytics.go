/*
Figma API

This is the OpenAPI specification for the [Figma REST API](https://www.figma.com/developers/api).  Note: we are releasing the OpenAPI specification as a beta given the large surface area and complexity of the REST API. If you notice any inaccuracies with the specification, please [file an issue](https://github.com/figma/rest-api-spec/issues).

API version: 0.20.0
Contact: support@figma.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// LibraryAnalyticsAPIService LibraryAnalyticsAPI service
type LibraryAnalyticsAPIService service

type ApiGetLibraryAnalyticsActionsRequest struct {
	ctx context.Context
	ApiService *LibraryAnalyticsAPIService
	fileKey string
	groupBy *string
	cursor *string
	startDate *string
	endDate *string
	order *string
}

// A dimension to group returned analytics data by. Accepts \&quot;component\&quot; or \&quot;team\&quot;.
func (r ApiGetLibraryAnalyticsActionsRequest) GroupBy(groupBy string) ApiGetLibraryAnalyticsActionsRequest {
	r.groupBy = &groupBy
	return r
}

// Cursor indicating what page of data to fetch. Obtained from prior API call.
func (r ApiGetLibraryAnalyticsActionsRequest) Cursor(cursor string) ApiGetLibraryAnalyticsActionsRequest {
	r.cursor = &cursor
	return r
}

// ISO 8601 date string (YYYY-MM-DD) of the earliest week to include. Dates are rounded back to the nearest start of a week. Defaults to one year prior.
func (r ApiGetLibraryAnalyticsActionsRequest) StartDate(startDate string) ApiGetLibraryAnalyticsActionsRequest {
	r.startDate = &startDate
	return r
}

// ISO 8601 date string (YYYY-MM-DD) of the latest week to include. Dates are rounded forward to the nearest end of a week. Defaults to the latest computed week.
func (r ApiGetLibraryAnalyticsActionsRequest) EndDate(endDate string) ApiGetLibraryAnalyticsActionsRequest {
	r.endDate = &endDate
	return r
}

// How to order response rows by week. This param can be either \&quot;asc\&quot; or \&quot;desc\&quot; (default).
func (r ApiGetLibraryAnalyticsActionsRequest) Order(order string) ApiGetLibraryAnalyticsActionsRequest {
	r.order = &order
	return r
}

func (r ApiGetLibraryAnalyticsActionsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetLibraryAnalyticsActionsExecute(r)
}

/*
GetLibraryAnalyticsActions Get library analytics action data.

Returns a list of library analytics actions data broken down by the requested dimension.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileKey File key of the library to fetch analytics data for.
 @return ApiGetLibraryAnalyticsActionsRequest
*/
func (a *LibraryAnalyticsAPIService) GetLibraryAnalyticsActions(ctx context.Context, fileKey string) ApiGetLibraryAnalyticsActionsRequest {
	return ApiGetLibraryAnalyticsActionsRequest{
		ApiService: a,
		ctx: ctx,
		fileKey: fileKey,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *LibraryAnalyticsAPIService) GetLibraryAnalyticsActionsExecute(r ApiGetLibraryAnalyticsActionsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LibraryAnalyticsAPIService.GetLibraryAnalyticsActions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analytics/libraries/{file_key}/actions"
	localVarPath = strings.Replace(localVarPath, "{"+"file_key"+"}", url.PathEscape(parameterValueToString(r.fileKey, "fileKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupBy == nil {
		return localVarReturnValue, nil, reportError("groupBy is required and must be specified")
	}

	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "group_by", r.groupBy, "form", "")
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "form", "")
	} else {
		var defaultValue string = "desc"
		r.order = &defaultValue
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
			if apiKey, ok := auth["PersonalAccessToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Figma-Token"] = key
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponsePayloadWithErrorBoolean
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponsePayloadWithErrorBoolean
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponsePayloadWithErrorBoolean
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponsePayloadWithErrorBoolean
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorResponsePayloadWithErrorBoolean
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetLibraryAnalyticsUsagesRequest struct {
	ctx context.Context
	ApiService *LibraryAnalyticsAPIService
	fileKey string
	groupBy *string
	cursor *string
	order *string
}

// A dimension to group returned analytics data by. Accepts \&quot;component\&quot; or \&quot;file\&quot;.
func (r ApiGetLibraryAnalyticsUsagesRequest) GroupBy(groupBy string) ApiGetLibraryAnalyticsUsagesRequest {
	r.groupBy = &groupBy
	return r
}

// Cursor indicating what page of data to fetch. Obtained from prior API call.
func (r ApiGetLibraryAnalyticsUsagesRequest) Cursor(cursor string) ApiGetLibraryAnalyticsUsagesRequest {
	r.cursor = &cursor
	return r
}

// How to order response rows by number of instances. This param can be either \&quot;asc\&quot; or \&quot;desc\&quot; (default).
func (r ApiGetLibraryAnalyticsUsagesRequest) Order(order string) ApiGetLibraryAnalyticsUsagesRequest {
	r.order = &order
	return r
}

func (r ApiGetLibraryAnalyticsUsagesRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetLibraryAnalyticsUsagesExecute(r)
}

/*
GetLibraryAnalyticsUsages Get library analytics usage data.

Returns a list of library analytics usage data broken down by the requested dimension.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileKey File key of the library to fetch analytics data for.
 @return ApiGetLibraryAnalyticsUsagesRequest
*/
func (a *LibraryAnalyticsAPIService) GetLibraryAnalyticsUsages(ctx context.Context, fileKey string) ApiGetLibraryAnalyticsUsagesRequest {
	return ApiGetLibraryAnalyticsUsagesRequest{
		ApiService: a,
		ctx: ctx,
		fileKey: fileKey,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *LibraryAnalyticsAPIService) GetLibraryAnalyticsUsagesExecute(r ApiGetLibraryAnalyticsUsagesRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LibraryAnalyticsAPIService.GetLibraryAnalyticsUsages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analytics/libraries/{file_key}/usages"
	localVarPath = strings.Replace(localVarPath, "{"+"file_key"+"}", url.PathEscape(parameterValueToString(r.fileKey, "fileKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupBy == nil {
		return localVarReturnValue, nil, reportError("groupBy is required and must be specified")
	}

	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "group_by", r.groupBy, "form", "")
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "form", "")
	} else {
		var defaultValue string = "desc"
		r.order = &defaultValue
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
			if apiKey, ok := auth["PersonalAccessToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Figma-Token"] = key
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponsePayloadWithErrorBoolean
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponsePayloadWithErrorBoolean
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponsePayloadWithErrorBoolean
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponsePayloadWithErrorBoolean
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorResponsePayloadWithErrorBoolean
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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
