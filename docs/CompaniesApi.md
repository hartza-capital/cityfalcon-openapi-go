# \CompaniesApi

All URIs are relative to *https://www.cityfalcon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMetadata**](CompaniesApi.md#ListMetadata) | **Get** /webapi/v1/search/metadata/items | 
[**ListStories**](CompaniesApi.md#ListStories) | **Get** /webapi/v1/stories | 



## ListMetadata

> Metadata ListMetadata(ctx).Query(query).Subsection(subsection).Limit(limit).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    query := "query_example" // string | string query (name or id) of the metadata
    subsection := "subsection_example" // string | string subsection (name or id) of the metadata (optional)
    limit := "limit_example" // string | string limit (name or id) of the metadata (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompaniesApi.ListMetadata(context.Background()).Query(query).Subsection(subsection).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.ListMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMetadata`: Metadata
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.ListMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | string query (name or id) of the metadata | 
 **subsection** | **string** | string subsection (name or id) of the metadata | 
 **limit** | **string** | string limit (name or id) of the metadata | 

### Return type

[**Metadata**](Metadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStories

> Stories ListStories(ctx).OrderBy(orderBy).Categories(categories).Languages(languages).MinScore(minScore).FoldSimilarStories(foldSimilarStories).TimeFilter(timeFilter).AssetClasses(assetClasses).Limit(limit).Query(query).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orderBy := "orderBy_example" // string | string order_by (name or id) of the stories
    categories := "categories_example" // string | string categories (name or id) of the stories (optional)
    languages := "languages_example" // string | string languages (name or id) of the stories (optional)
    minScore := "minScore_example" // string | string min_score (name or id) of the stories (optional)
    foldSimilarStories := "foldSimilarStories_example" // string | string fold_similar_stories (name or id) of the stories (optional)
    timeFilter := "timeFilter_example" // string | string time_filter (name or id) of the stories (optional)
    assetClasses := "assetClasses_example" // string | string asset_classes (name or id) of the stories (optional)
    limit := "limit_example" // string | string limit (name or id) of the stories (optional)
    query := "query_example" // string | string query (name or id) of the stories (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompaniesApi.ListStories(context.Background()).OrderBy(orderBy).Categories(categories).Languages(languages).MinScore(minScore).FoldSimilarStories(foldSimilarStories).TimeFilter(timeFilter).AssetClasses(assetClasses).Limit(limit).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.ListStories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStories`: Stories
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.ListStories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **string** | string order_by (name or id) of the stories | 
 **categories** | **string** | string categories (name or id) of the stories | 
 **languages** | **string** | string languages (name or id) of the stories | 
 **minScore** | **string** | string min_score (name or id) of the stories | 
 **foldSimilarStories** | **string** | string fold_similar_stories (name or id) of the stories | 
 **timeFilter** | **string** | string time_filter (name or id) of the stories | 
 **assetClasses** | **string** | string asset_classes (name or id) of the stories | 
 **limit** | **string** | string limit (name or id) of the stories | 
 **query** | **string** | string query (name or id) of the stories | 

### Return type

[**Stories**](Stories.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

