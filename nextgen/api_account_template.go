
/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the Access Control Service. This is under active development. Beware of the breaking change with respect to the generated code stub.
 *
 * API version: 1.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type AccountTemplateApiService service
/*
AccountTemplateApiService Create Template
Creates a Template in the Account scope.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *AccountTemplateApiCreateTemplatesAccOpts - Optional Parameters:
     * @param "Body" (optional.Interface of TemplateCreateRequestBody) -  Templates Create Request Body
     * @param "HarnessAccount" (optional.String) -  Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped.
@return TemplateResponse
*/

type AccountTemplateApiCreateTemplatesAccOpts struct {
    Body optional.Interface
    HarnessAccount optional.String
}

func (a *AccountTemplateApiService) CreateTemplatesAcc(ctx context.Context, localVarOptionals *AccountTemplateApiCreateTemplatesAccOpts) (TemplateResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TemplateResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/templates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/yaml"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.HarnessAccount.IsSet() {
		localVarHeaderParams["Harness-Account"] = parameterToString(localVarOptionals.HarnessAccount.Value(), "")
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		
		localVarOptionalBody:= localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 201 {
			var v TemplateResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
AccountTemplateApiService Delete Template
Deletes particular version of Template at Account scope.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param template Template Identifier
 * @param version Version Label for Template
 * @param optional nil or *AccountTemplateApiDeleteTemplateAccOpts - Optional Parameters:
     * @param "HarnessAccount" (optional.String) -  Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped.
     * @param "Comments" (optional.String) -  Specify comment with respect to changes  

*/

type AccountTemplateApiDeleteTemplateAccOpts struct {
    HarnessAccount optional.String
    Comments optional.String
}

func (a *AccountTemplateApiService) DeleteTemplateAcc(ctx context.Context, template string, version string, localVarOptionals *AccountTemplateApiDeleteTemplateAccOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/templates/{template}/versions/{version}"
	localVarPath = strings.Replace(localVarPath, "{"+"template"+"}", fmt.Sprintf("%v", template), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", fmt.Sprintf("%v", version), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Comments.IsSet() {
		localVarQueryParams.Add("comments", parameterToString(localVarOptionals.Comments.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.HarnessAccount.IsSet() {
		localVarHeaderParams["Harness-Account"] = parameterToString(localVarOptionals.HarnessAccount.Value(), "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
AccountTemplateApiService Retrieve a Template
Retrieves particular version of Template at Account scope.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param template Template Identifier
 * @param version Version Label for Template
 * @param optional nil or *AccountTemplateApiGetTemplateAccOpts - Optional Parameters:
     * @param "HarnessAccount" (optional.String) -  Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped.
     * @param "IncludeYaml" (optional.Bool) -  Use it to get Template along with Input Set YAML
     * @param "BranchName" (optional.String) -  Name of the branch
     * @param "ParentConnectorRef" (optional.String) -  Connector ref of parent template if its remote
     * @param "ParentRepoName" (optional.String) -  Repo name of parent template if its remote
     * @param "ParentAccountId" (optional.String) -  Account name of parent template if its remote
     * @param "ParentOrgId" (optional.String) -  Organization name of parent template if its remote
     * @param "ParentProjectId" (optional.String) -  Project name of parent entity if its remote
@return TemplateWithInputsResponse
*/

type AccountTemplateApiGetTemplateAccOpts struct {
    HarnessAccount optional.String
    IncludeYaml optional.Bool
    BranchName optional.String
    ParentConnectorRef optional.String
    ParentRepoName optional.String
    ParentAccountId optional.String
    ParentOrgId optional.String
    ParentProjectId optional.String
}

func (a *AccountTemplateApiService) GetTemplateAcc(ctx context.Context, template string, version string, localVarOptionals *AccountTemplateApiGetTemplateAccOpts) (TemplateWithInputsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TemplateWithInputsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/templates/{template}/versions/{version}"
	localVarPath = strings.Replace(localVarPath, "{"+"template"+"}", fmt.Sprintf("%v", template), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", fmt.Sprintf("%v", version), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.IncludeYaml.IsSet() {
		localVarQueryParams.Add("include_yaml", parameterToString(localVarOptionals.IncludeYaml.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BranchName.IsSet() {
		localVarQueryParams.Add("branch_name", parameterToString(localVarOptionals.BranchName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ParentConnectorRef.IsSet() {
		localVarQueryParams.Add("parent_connector_ref", parameterToString(localVarOptionals.ParentConnectorRef.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ParentRepoName.IsSet() {
		localVarQueryParams.Add("parent_repo_name", parameterToString(localVarOptionals.ParentRepoName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ParentAccountId.IsSet() {
		localVarQueryParams.Add("parent_account_id", parameterToString(localVarOptionals.ParentAccountId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ParentOrgId.IsSet() {
		localVarQueryParams.Add("parent_org_id", parameterToString(localVarOptionals.ParentOrgId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ParentProjectId.IsSet() {
		localVarQueryParams.Add("parent_project_id", parameterToString(localVarOptionals.ParentProjectId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.HarnessAccount.IsSet() {
		localVarHeaderParams["Harness-Account"] = parameterToString(localVarOptionals.HarnessAccount.Value(), "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v TemplateWithInputsResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
AccountTemplateApiService Get Stable Template
Retrieves stable version of Template at Account scope.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param template Template Identifier
 * @param optional nil or *AccountTemplateApiGetTemplateStableAccOpts - Optional Parameters:
     * @param "HarnessAccount" (optional.String) -  Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped.
     * @param "IncludeYaml" (optional.Bool) -  Use it to get Template along with Input Set YAML
     * @param "BranchName" (optional.String) -  Name of the branch
     * @param "ParentConnectorRef" (optional.String) -  Connector ref of parent template if its remote
     * @param "ParentRepoName" (optional.String) -  Repo name of parent template if its remote
     * @param "ParentAccountId" (optional.String) -  Account name of parent template if its remote
     * @param "ParentOrgId" (optional.String) -  Organization name of parent template if its remote
     * @param "ParentProjectId" (optional.String) -  Project name of parent entity if its remote
@return TemplateWithInputsResponse
*/

type AccountTemplateApiGetTemplateStableAccOpts struct {
    HarnessAccount optional.String
    IncludeYaml optional.Bool
    BranchName optional.String
    ParentConnectorRef optional.String
    ParentRepoName optional.String
    ParentAccountId optional.String
    ParentOrgId optional.String
    ParentProjectId optional.String
}

func (a *AccountTemplateApiService) GetTemplateStableAcc(ctx context.Context, template string, localVarOptionals *AccountTemplateApiGetTemplateStableAccOpts) (TemplateWithInputsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TemplateWithInputsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/templates/{template}"
	localVarPath = strings.Replace(localVarPath, "{"+"template"+"}", fmt.Sprintf("%v", template), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.IncludeYaml.IsSet() {
		localVarQueryParams.Add("include_yaml", parameterToString(localVarOptionals.IncludeYaml.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BranchName.IsSet() {
		localVarQueryParams.Add("branch_name", parameterToString(localVarOptionals.BranchName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ParentConnectorRef.IsSet() {
		localVarQueryParams.Add("parent_connector_ref", parameterToString(localVarOptionals.ParentConnectorRef.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ParentRepoName.IsSet() {
		localVarQueryParams.Add("parent_repo_name", parameterToString(localVarOptionals.ParentRepoName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ParentAccountId.IsSet() {
		localVarQueryParams.Add("parent_account_id", parameterToString(localVarOptionals.ParentAccountId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ParentOrgId.IsSet() {
		localVarQueryParams.Add("parent_org_id", parameterToString(localVarOptionals.ParentOrgId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ParentProjectId.IsSet() {
		localVarQueryParams.Add("parent_project_id", parameterToString(localVarOptionals.ParentProjectId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.HarnessAccount.IsSet() {
		localVarHeaderParams["Harness-Account"] = parameterToString(localVarOptionals.HarnessAccount.Value(), "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v TemplateWithInputsResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
AccountTemplateApiService Get Templates List
Retrieves list of Template with meta-data at Account scope.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *AccountTemplateApiGetTemplatesListAccOpts - Optional Parameters:
     * @param "HarnessAccount" (optional.String) -  Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped.
     * @param "Page" (optional.Int32) -  Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page 
     * @param "Limit" (optional.Int32) -  Pagination: Number of items to return
     * @param "Sort" (optional.String) -  Parameter on the basis of which sorting is done.
     * @param "Order" (optional.String) -  Order on the basis of which sorting is done.
     * @param "SearchTerm" (optional.String) -  This would be used to filter resources having attributes matching with search term.
     * @param "Type_" (optional.String) -  Template List Type
     * @param "Recursive" (optional.Bool) -  Specify true if all accessible Templates are to be included
     * @param "Names" (optional.Interface of []string) -  Template names for filtering
     * @param "Identifiers" (optional.Interface of []string) -  Template Ids for Filtering
     * @param "Description" (optional.String) -  Filter properties description
     * @param "EntityTypes" (optional.Interface of []string) -  Type of Template
     * @param "ChildTypes" (optional.Interface of []string) -  Template Child Types for filtering
@return []TemplateMetadataSummaryResponse
*/

type AccountTemplateApiGetTemplatesListAccOpts struct {
    HarnessAccount optional.String
    Page optional.Int32
    Limit optional.Int32
    Sort optional.String
    Order optional.String
    SearchTerm optional.String
    Type_ optional.String
    Recursive optional.Bool
    Names optional.Interface
    Identifiers optional.Interface
    Description optional.String
    EntityTypes optional.Interface
    ChildTypes optional.Interface
}

func (a *AccountTemplateApiService) GetTemplatesListAcc(ctx context.Context, localVarOptionals *AccountTemplateApiGetTemplatesListAccOpts) ([]TemplateMetadataSummaryResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []TemplateMetadataSummaryResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/templates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		localVarQueryParams.Add("sort", parameterToString(localVarOptionals.Sort.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Order.IsSet() {
		localVarQueryParams.Add("order", parameterToString(localVarOptionals.Order.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchTerm.IsSet() {
		localVarQueryParams.Add("search_term", parameterToString(localVarOptionals.SearchTerm.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Recursive.IsSet() {
		localVarQueryParams.Add("recursive", parameterToString(localVarOptionals.Recursive.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Names.IsSet() {
		localVarQueryParams.Add("names", parameterToString(localVarOptionals.Names.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.Identifiers.IsSet() {
		localVarQueryParams.Add("identifiers", parameterToString(localVarOptionals.Identifiers.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.Description.IsSet() {
		localVarQueryParams.Add("description", parameterToString(localVarOptionals.Description.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EntityTypes.IsSet() {
		localVarQueryParams.Add("entity_types", parameterToString(localVarOptionals.EntityTypes.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.ChildTypes.IsSet() {
		localVarQueryParams.Add("child_types", parameterToString(localVarOptionals.ChildTypes.Value(), "multi"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.HarnessAccount.IsSet() {
		localVarHeaderParams["Harness-Account"] = parameterToString(localVarOptionals.HarnessAccount.Value(), "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []TemplateMetadataSummaryResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
AccountTemplateApiService Update Template
Updates particular version of Template at Account scope.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param template Template Identifier
 * @param version Version Label for Template
 * @param optional nil or *AccountTemplateApiUpdateTemplateAccOpts - Optional Parameters:
     * @param "Body" (optional.Interface of TemplateUpdateRequestBody) -  Templates Update Request Body
     * @param "HarnessAccount" (optional.String) -  Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped.
@return TemplateResponse
*/

type AccountTemplateApiUpdateTemplateAccOpts struct {
    Body optional.Interface
    HarnessAccount optional.String
}

func (a *AccountTemplateApiService) UpdateTemplateAcc(ctx context.Context, template string, version string, localVarOptionals *AccountTemplateApiUpdateTemplateAccOpts) (TemplateResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TemplateResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/templates/{template}/versions/{version}"
	localVarPath = strings.Replace(localVarPath, "{"+"template"+"}", fmt.Sprintf("%v", template), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", fmt.Sprintf("%v", version), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/yaml"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.HarnessAccount.IsSet() {
		localVarHeaderParams["Harness-Account"] = parameterToString(localVarOptionals.HarnessAccount.Value(), "")
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		
		localVarOptionalBody:= localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v TemplateResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
AccountTemplateApiService Update Stable Template
Updates the stable version of Template at Account scope.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param template Template Identifier
 * @param version Version Label for Template
 * @param optional nil or *AccountTemplateApiUpdateTemplateStableAccOpts - Optional Parameters:
     * @param "Body" (optional.Interface of GitFindDetails) -  Templates Fetch Request Body
     * @param "HarnessAccount" (optional.String) -  Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped.
@return TemplateUpdateStableResponse
*/

type AccountTemplateApiUpdateTemplateStableAccOpts struct {
    Body optional.Interface
    HarnessAccount optional.String
}

func (a *AccountTemplateApiService) UpdateTemplateStableAcc(ctx context.Context, template string, version string, localVarOptionals *AccountTemplateApiUpdateTemplateStableAccOpts) (TemplateUpdateStableResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TemplateUpdateStableResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/templates/{template}/versions/{version}/stable"
	localVarPath = strings.Replace(localVarPath, "{"+"template"+"}", fmt.Sprintf("%v", template), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", fmt.Sprintf("%v", version), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/yaml"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.HarnessAccount.IsSet() {
		localVarHeaderParams["Harness-Account"] = parameterToString(localVarOptionals.HarnessAccount.Value(), "")
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		
		localVarOptionalBody:= localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v TemplateUpdateStableResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
