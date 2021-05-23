/*
 * osrsbox-api
 *
 * An open, free, complete and up-to-date RESTful API for Old School RuneScape (OSRS) items, monsters and prayers.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
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

// MonsterApiService MonsterApi service
type MonsterApiService service

type ApiDeleteMonsterItemRequest struct {
	ctx _context.Context
	ApiService *MonsterApiService
	monsterId string
	ifMatch *string
}

func (r ApiDeleteMonsterItemRequest) IfMatch(ifMatch string) ApiDeleteMonsterItemRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiDeleteMonsterItemRequest) Execute() (*_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.DeleteMonsterItemExecute(r)
}

/*
 * DeleteMonsterItem Deletes a Monster document
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param monsterId Unique OSRS monster ID number.
 * @return ApiDeleteMonsterItemRequest
 */
func (a *MonsterApiService) DeleteMonsterItem(ctx _context.Context, monsterId string) ApiDeleteMonsterItemRequest {
	return ApiDeleteMonsterItemRequest{
		ApiService: a,
		ctx: ctx,
		monsterId: monsterId,
	}
}

/*
 * Execute executes the request
 */
func (a *MonsterApiService) DeleteMonsterItemExecute(r ApiDeleteMonsterItemRequest) (*_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonsterApiService.DeleteMonsterItem")
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarPath := localBasePath + "/monsters/{monsterId}"
	localVarPath = strings.Replace(localVarPath, "{"+"monsterId"+"}", _neturl.PathEscape(parameterToString(r.monsterId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.ifMatch == nil {
		executionError.error = "ifMatch is required and must be specified"
		return nil, executionError
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
	localVarHeaderParams["If-Match"] = parameterToString(*r.ifMatch, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, executionError
}

type ApiDeletemonstersRequest struct {
	ctx _context.Context
	ApiService *MonsterApiService
}


func (r ApiDeletemonstersRequest) Execute() (*_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.DeletemonstersExecute(r)
}

/*
 * Deletemonsters Deletes all monsters
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiDeletemonstersRequest
 */
func (a *MonsterApiService) Deletemonsters(ctx _context.Context) ApiDeletemonstersRequest {
	return ApiDeletemonstersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *MonsterApiService) DeletemonstersExecute(r ApiDeletemonstersRequest) (*_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonsterApiService.Deletemonsters")
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarPath := localBasePath + "/monsters"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, executionError
}

type ApiGetMonsterItemRequest struct {
	ctx _context.Context
	ApiService *MonsterApiService
	monsterId string
}


func (r ApiGetMonsterItemRequest) Execute() (Monster, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.GetMonsterItemExecute(r)
}

/*
 * GetMonsterItem Retrieves a Monster document
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param monsterId Unique OSRS monster ID number.
 * @return ApiGetMonsterItemRequest
 */
func (a *MonsterApiService) GetMonsterItem(ctx _context.Context, monsterId string) ApiGetMonsterItemRequest {
	return ApiGetMonsterItemRequest{
		ApiService: a,
		ctx: ctx,
		monsterId: monsterId,
	}
}

/*
 * Execute executes the request
 * @return Monster
 */
func (a *MonsterApiService) GetMonsterItemExecute(r ApiGetMonsterItemRequest) (Monster, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  Monster
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonsterApiService.GetMonsterItem")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/monsters/{monsterId}"
	localVarPath = strings.Replace(localVarPath, "{"+"monsterId"+"}", _neturl.PathEscape(parameterToString(r.monsterId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

	return localVarReturnValue, localVarHTTPResponse, executionError
}

type ApiGetmonstersRequest struct {
	ctx _context.Context
	ApiService *MonsterApiService
	where *string
	projection *string
	sort *string
	page *int32
	maxResults *int32
}

func (r ApiGetmonstersRequest) Where(where string) ApiGetmonstersRequest {
	r.where = &where
	return r
}
func (r ApiGetmonstersRequest) Projection(projection string) ApiGetmonstersRequest {
	r.projection = &projection
	return r
}
func (r ApiGetmonstersRequest) Sort(sort string) ApiGetmonstersRequest {
	r.sort = &sort
	return r
}
func (r ApiGetmonstersRequest) Page(page int32) ApiGetmonstersRequest {
	r.page = &page
	return r
}
func (r ApiGetmonstersRequest) MaxResults(maxResults int32) ApiGetmonstersRequest {
	r.maxResults = &maxResults
	return r
}

func (r ApiGetmonstersRequest) Execute() (InlineResponse2003, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.GetmonstersExecute(r)
}

/*
 * Getmonsters Retrieves one or more monsters
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetmonstersRequest
 */
func (a *MonsterApiService) Getmonsters(ctx _context.Context) ApiGetmonstersRequest {
	return ApiGetmonstersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse2003
 */
func (a *MonsterApiService) GetmonstersExecute(r ApiGetmonstersRequest) (InlineResponse2003, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  InlineResponse2003
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonsterApiService.Getmonsters")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/monsters"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, ""))
	}
	if r.projection != nil {
		localVarQueryParams.Add("projection", parameterToString(*r.projection, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.maxResults != nil {
		localVarQueryParams.Add("max_results", parameterToString(*r.maxResults, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

	return localVarReturnValue, localVarHTTPResponse, executionError
}

type ApiPostmonstersRequest struct {
	ctx _context.Context
	ApiService *MonsterApiService
	monster *Monster
}

func (r ApiPostmonstersRequest) Monster(monster Monster) ApiPostmonstersRequest {
	r.monster = &monster
	return r
}

func (r ApiPostmonstersRequest) Execute() (*_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.PostmonstersExecute(r)
}

/*
 * Postmonsters Stores one or more monsters.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPostmonstersRequest
 */
func (a *MonsterApiService) Postmonsters(ctx _context.Context) ApiPostmonstersRequest {
	return ApiPostmonstersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *MonsterApiService) PostmonstersExecute(r ApiPostmonstersRequest) (*_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonsterApiService.Postmonsters")
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarPath := localBasePath + "/monsters"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.monster == nil {
		executionError.error = "monster is required and must be specified"
		return nil, executionError
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.monster
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, executionError
}

type ApiPutMonsterItemRequest struct {
	ctx _context.Context
	ApiService *MonsterApiService
	monsterId string
	ifMatch *string
	monster *Monster
}

func (r ApiPutMonsterItemRequest) IfMatch(ifMatch string) ApiPutMonsterItemRequest {
	r.ifMatch = &ifMatch
	return r
}
func (r ApiPutMonsterItemRequest) Monster(monster Monster) ApiPutMonsterItemRequest {
	r.monster = &monster
	return r
}

func (r ApiPutMonsterItemRequest) Execute() (*_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.PutMonsterItemExecute(r)
}

/*
 * PutMonsterItem Replaces a Monster document
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param monsterId Unique OSRS monster ID number.
 * @return ApiPutMonsterItemRequest
 */
func (a *MonsterApiService) PutMonsterItem(ctx _context.Context, monsterId string) ApiPutMonsterItemRequest {
	return ApiPutMonsterItemRequest{
		ApiService: a,
		ctx: ctx,
		monsterId: monsterId,
	}
}

/*
 * Execute executes the request
 */
func (a *MonsterApiService) PutMonsterItemExecute(r ApiPutMonsterItemRequest) (*_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonsterApiService.PutMonsterItem")
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarPath := localBasePath + "/monsters/{monsterId}"
	localVarPath = strings.Replace(localVarPath, "{"+"monsterId"+"}", _neturl.PathEscape(parameterToString(r.monsterId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.ifMatch == nil {
		executionError.error = "ifMatch is required and must be specified"
		return nil, executionError
	}
	if r.monster == nil {
		executionError.error = "monster is required and must be specified"
		return nil, executionError
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarHeaderParams["If-Match"] = parameterToString(*r.ifMatch, "")
	// body params
	localVarPostBody = r.monster
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, executionError
}
