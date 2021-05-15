/*
 * cityfalcon
 *
 * cityfalcon API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MetadataRegularGlobal struct for MetadataRegularGlobal
type MetadataRegularGlobal struct {
	TopicId                  *int64                           `json:"topicId,omitempty"`
	TopicName                *string                          `json:"topicName,omitempty"`
	TopicTickers             *[]string                        `json:"topicTickers,omitempty"`
	TopicLegalEntities       *string                          `json:"topicLegalEntities,omitempty"`
	OrderPriority            *int64                           `json:"orderPriority,omitempty"`
	Attributes               *MetadataRegularGlobalAttributes `json:"attributes,omitempty"`
	Groups                   *[]MetadataRegularGlobalGroup    `json:"groups,omitempty"`
	Tickers                  *[]MetadataRegularGlobalTicker   `json:"tickers,omitempty"`
	TradingViewTicker        *string                          `json:"tradingViewTicker,omitempty"`
	Figis                    *string                          `json:"figis,omitempty"`
	CompanyDomains           *string                          `json:"companyDomains,omitempty"`
	Score                    *string                          `json:"_score,omitempty"`
	Highlight                *string                          `json:"_highlight,omitempty"`
	Permalink                *string                          `json:"permalink,omitempty"`
	IsinAliases              *string                          `json:"isinAliases,omitempty"`
	AssetClassSearchKeywords *string                          `json:"assetClassSearchKeywords,omitempty"`
}

// NewMetadataRegularGlobal instantiates a new MetadataRegularGlobal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataRegularGlobal() *MetadataRegularGlobal {
	this := MetadataRegularGlobal{}
	return &this
}

// NewMetadataRegularGlobalWithDefaults instantiates a new MetadataRegularGlobal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataRegularGlobalWithDefaults() *MetadataRegularGlobal {
	this := MetadataRegularGlobal{}
	return &this
}

// GetTopicId returns the TopicId field value if set, zero value otherwise.
func (o *MetadataRegularGlobal) GetTopicId() int64 {
	if o == nil || o.TopicId == nil {
		var ret int64
		return ret
	}
	return *o.TopicId
}

// GetTopicIdOk returns a tuple with the TopicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataRegularGlobal) GetTopicIdOk() (*int64, bool) {
	if o == nil || o.TopicId == nil {
		return nil, false
	}
	return o.TopicId, true
}

// HasTopicId returns a boolean if a field has been set.
func (o *MetadataRegularGlobal) HasTopicId() bool {
	if o != nil && o.TopicId != nil {
		return true
	}

	return false
}

// SetTopicId gets a reference to the given int64 and assigns it to the TopicId field.
func (o *MetadataRegularGlobal) SetTopicId(v int64) {
	o.TopicId = &v
}

// GetTopicName returns the TopicName field value if set, zero value otherwise.
func (o *MetadataRegularGlobal) GetTopicName() string {
	if o == nil || o.TopicName == nil {
		var ret string
		return ret
	}
	return *o.TopicName
}

// GetTopicNameOk returns a tuple with the TopicName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataRegularGlobal) GetTopicNameOk() (*string, bool) {
	if o == nil || o.TopicName == nil {
		return nil, false
	}
	return o.TopicName, true
}

// HasTopicName returns a boolean if a field has been set.
func (o *MetadataRegularGlobal) HasTopicName() bool {
	if o != nil && o.TopicName != nil {
		return true
	}

	return false
}

// SetTopicName gets a reference to the given string and assigns it to the TopicName field.
func (o *MetadataRegularGlobal) SetTopicName(v string) {
	o.TopicName = &v
}

// GetTopicTickers returns the TopicTickers field value if set, zero value otherwise.
func (o *MetadataRegularGlobal) GetTopicTickers() []string {
	if o == nil || o.TopicTickers == nil {
		var ret []string
		return ret
	}
	return *o.TopicTickers
}

// GetTopicTickersOk returns a tuple with the TopicTickers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataRegularGlobal) GetTopicTickersOk() (*[]string, bool) {
	if o == nil || o.TopicTickers == nil {
		return nil, false
	}
	return o.TopicTickers, true
}

// HasTopicTickers returns a boolean if a field has been set.
func (o *MetadataRegularGlobal) HasTopicTickers() bool {
	if o != nil && o.TopicTickers != nil {
		return true
	}

	return false
}

// SetTopicTickers gets a reference to the given []string and assigns it to the TopicTickers field.
func (o *MetadataRegularGlobal) SetTopicTickers(v []string) {
	o.TopicTickers = &v
}

// GetTopicLegalEntities returns the TopicLegalEntities field value if set, zero value otherwise.
func (o *MetadataRegularGlobal) GetTopicLegalEntities() string {
	if o == nil || o.TopicLegalEntities == nil {
		var ret string
		return ret
	}
	return *o.TopicLegalEntities
}

// GetTopicLegalEntitiesOk returns a tuple with the TopicLegalEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataRegularGlobal) GetTopicLegalEntitiesOk() (*string, bool) {
	if o == nil || o.TopicLegalEntities == nil {
		return nil, false
	}
	return o.TopicLegalEntities, true
}

// HasTopicLegalEntities returns a boolean if a field has been set.
func (o *MetadataRegularGlobal) HasTopicLegalEntities() bool {
	if o != nil && o.TopicLegalEntities != nil {
		return true
	}

	return false
}

// SetTopicLegalEntities gets a reference to the given string and assigns it to the TopicLegalEntities field.
func (o *MetadataRegularGlobal) SetTopicLegalEntities(v string) {
	o.TopicLegalEntities = &v
}

// GetOrderPriority returns the OrderPriority field value if set, zero value otherwise.
func (o *MetadataRegularGlobal) GetOrderPriority() int64 {
	if o == nil || o.OrderPriority == nil {
		var ret int64
		return ret
	}
	return *o.OrderPriority
}

// GetOrderPriorityOk returns a tuple with the OrderPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataRegularGlobal) GetOrderPriorityOk() (*int64, bool) {
	if o == nil || o.OrderPriority == nil {
		return nil, false
	}
	return o.OrderPriority, true
}

// HasOrderPriority returns a boolean if a field has been set.
func (o *MetadataRegularGlobal) HasOrderPriority() bool {
	if o != nil && o.OrderPriority != nil {
		return true
	}

	return false
}

// SetOrderPriority gets a reference to the given int64 and assigns it to the OrderPriority field.
func (o *MetadataRegularGlobal) SetOrderPriority(v int64) {
	o.OrderPriority = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MetadataRegularGlobal) GetAttributes() MetadataRegularGlobalAttributes {
	if o == nil || o.Attributes == nil {
		var ret MetadataRegularGlobalAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataRegularGlobal) GetAttributesOk() (*MetadataRegularGlobalAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MetadataRegularGlobal) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given MetadataRegularGlobalAttributes and assigns it to the Attributes field.
func (o *MetadataRegularGlobal) SetAttributes(v MetadataRegularGlobalAttributes) {
	o.Attributes = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *MetadataRegularGlobal) GetGroups() []MetadataRegularGlobalGroup {
	if o == nil || o.Groups == nil {
		var ret []MetadataRegularGlobalGroup
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataRegularGlobal) GetGroupsOk() (*[]MetadataRegularGlobalGroup, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *MetadataRegularGlobal) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []MetadataRegularGlobalGroup and assigns it to the Groups field.
func (o *MetadataRegularGlobal) SetGroups(v []MetadataRegularGlobalGroup) {
	o.Groups = &v
}

// GetTickers returns the Tickers field value if set, zero value otherwise.
func (o *MetadataRegularGlobal) GetTickers() []MetadataRegularGlobalTicker {
	if o == nil || o.Tickers == nil {
		var ret []MetadataRegularGlobalTicker
		return ret
	}
	return *o.Tickers
}

// GetTickersOk returns a tuple with the Tickers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataRegularGlobal) GetTickersOk() (*[]MetadataRegularGlobalTicker, bool) {
	if o == nil || o.Tickers == nil {
		return nil, false
	}
	return o.Tickers, true
}

// HasTickers returns a boolean if a field has been set.
func (o *MetadataRegularGlobal) HasTickers() bool {
	if o != nil && o.Tickers != nil {
		return true
	}

	return false
}

// SetTickers gets a reference to the given []MetadataRegularGlobalTicker and assigns it to the Tickers field.
func (o *MetadataRegularGlobal) SetTickers(v []MetadataRegularGlobalTicker) {
	o.Tickers = &v
}

// GetTradingViewTicker returns the TradingViewTicker field value if set, zero value otherwise.
func (o *MetadataRegularGlobal) GetTradingViewTicker() string {
	if o == nil || o.TradingViewTicker == nil {
		var ret string
		return ret
	}
	return *o.TradingViewTicker
}

// GetTradingViewTickerOk returns a tuple with the TradingViewTicker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataRegularGlobal) GetTradingViewTickerOk() (*string, bool) {
	if o == nil || o.TradingViewTicker == nil {
		return nil, false
	}
	return o.TradingViewTicker, true
}

// HasTradingViewTicker returns a boolean if a field has been set.
func (o *MetadataRegularGlobal) HasTradingViewTicker() bool {
	if o != nil && o.TradingViewTicker != nil {
		return true
	}

	return false
}

// SetTradingViewTicker gets a reference to the given string and assigns it to the TradingViewTicker field.
func (o *MetadataRegularGlobal) SetTradingViewTicker(v string) {
	o.TradingViewTicker = &v
}

// GetFigis returns the Figis field value if set, zero value otherwise.
func (o *MetadataRegularGlobal) GetFigis() string {
	if o == nil || o.Figis == nil {
		var ret string
		return ret
	}
	return *o.Figis
}

// GetFigisOk returns a tuple with the Figis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataRegularGlobal) GetFigisOk() (*string, bool) {
	if o == nil || o.Figis == nil {
		return nil, false
	}
	return o.Figis, true
}

// HasFigis returns a boolean if a field has been set.
func (o *MetadataRegularGlobal) HasFigis() bool {
	if o != nil && o.Figis != nil {
		return true
	}

	return false
}

// SetFigis gets a reference to the given string and assigns it to the Figis field.
func (o *MetadataRegularGlobal) SetFigis(v string) {
	o.Figis = &v
}

// GetCompanyDomains returns the CompanyDomains field value if set, zero value otherwise.
func (o *MetadataRegularGlobal) GetCompanyDomains() string {
	if o == nil || o.CompanyDomains == nil {
		var ret string
		return ret
	}
	return *o.CompanyDomains
}

// GetCompanyDomainsOk returns a tuple with the CompanyDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataRegularGlobal) GetCompanyDomainsOk() (*string, bool) {
	if o == nil || o.CompanyDomains == nil {
		return nil, false
	}
	return o.CompanyDomains, true
}

// HasCompanyDomains returns a boolean if a field has been set.
func (o *MetadataRegularGlobal) HasCompanyDomains() bool {
	if o != nil && o.CompanyDomains != nil {
		return true
	}

	return false
}

// SetCompanyDomains gets a reference to the given string and assigns it to the CompanyDomains field.
func (o *MetadataRegularGlobal) SetCompanyDomains(v string) {
	o.CompanyDomains = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *MetadataRegularGlobal) GetScore() string {
	if o == nil || o.Score == nil {
		var ret string
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataRegularGlobal) GetScoreOk() (*string, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *MetadataRegularGlobal) HasScore() bool {
	if o != nil && o.Score != nil {
		return true
	}

	return false
}

// SetScore gets a reference to the given string and assigns it to the Score field.
func (o *MetadataRegularGlobal) SetScore(v string) {
	o.Score = &v
}

// GetHighlight returns the Highlight field value if set, zero value otherwise.
func (o *MetadataRegularGlobal) GetHighlight() string {
	if o == nil || o.Highlight == nil {
		var ret string
		return ret
	}
	return *o.Highlight
}

// GetHighlightOk returns a tuple with the Highlight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataRegularGlobal) GetHighlightOk() (*string, bool) {
	if o == nil || o.Highlight == nil {
		return nil, false
	}
	return o.Highlight, true
}

// HasHighlight returns a boolean if a field has been set.
func (o *MetadataRegularGlobal) HasHighlight() bool {
	if o != nil && o.Highlight != nil {
		return true
	}

	return false
}

// SetHighlight gets a reference to the given string and assigns it to the Highlight field.
func (o *MetadataRegularGlobal) SetHighlight(v string) {
	o.Highlight = &v
}

// GetPermalink returns the Permalink field value if set, zero value otherwise.
func (o *MetadataRegularGlobal) GetPermalink() string {
	if o == nil || o.Permalink == nil {
		var ret string
		return ret
	}
	return *o.Permalink
}

// GetPermalinkOk returns a tuple with the Permalink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataRegularGlobal) GetPermalinkOk() (*string, bool) {
	if o == nil || o.Permalink == nil {
		return nil, false
	}
	return o.Permalink, true
}

// HasPermalink returns a boolean if a field has been set.
func (o *MetadataRegularGlobal) HasPermalink() bool {
	if o != nil && o.Permalink != nil {
		return true
	}

	return false
}

// SetPermalink gets a reference to the given string and assigns it to the Permalink field.
func (o *MetadataRegularGlobal) SetPermalink(v string) {
	o.Permalink = &v
}

// GetIsinAliases returns the IsinAliases field value if set, zero value otherwise.
func (o *MetadataRegularGlobal) GetIsinAliases() string {
	if o == nil || o.IsinAliases == nil {
		var ret string
		return ret
	}
	return *o.IsinAliases
}

// GetIsinAliasesOk returns a tuple with the IsinAliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataRegularGlobal) GetIsinAliasesOk() (*string, bool) {
	if o == nil || o.IsinAliases == nil {
		return nil, false
	}
	return o.IsinAliases, true
}

// HasIsinAliases returns a boolean if a field has been set.
func (o *MetadataRegularGlobal) HasIsinAliases() bool {
	if o != nil && o.IsinAliases != nil {
		return true
	}

	return false
}

// SetIsinAliases gets a reference to the given string and assigns it to the IsinAliases field.
func (o *MetadataRegularGlobal) SetIsinAliases(v string) {
	o.IsinAliases = &v
}

// GetAssetClassSearchKeywords returns the AssetClassSearchKeywords field value if set, zero value otherwise.
func (o *MetadataRegularGlobal) GetAssetClassSearchKeywords() string {
	if o == nil || o.AssetClassSearchKeywords == nil {
		var ret string
		return ret
	}
	return *o.AssetClassSearchKeywords
}

// GetAssetClassSearchKeywordsOk returns a tuple with the AssetClassSearchKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataRegularGlobal) GetAssetClassSearchKeywordsOk() (*string, bool) {
	if o == nil || o.AssetClassSearchKeywords == nil {
		return nil, false
	}
	return o.AssetClassSearchKeywords, true
}

// HasAssetClassSearchKeywords returns a boolean if a field has been set.
func (o *MetadataRegularGlobal) HasAssetClassSearchKeywords() bool {
	if o != nil && o.AssetClassSearchKeywords != nil {
		return true
	}

	return false
}

// SetAssetClassSearchKeywords gets a reference to the given string and assigns it to the AssetClassSearchKeywords field.
func (o *MetadataRegularGlobal) SetAssetClassSearchKeywords(v string) {
	o.AssetClassSearchKeywords = &v
}

func (o MetadataRegularGlobal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TopicId != nil {
		toSerialize["topicId"] = o.TopicId
	}
	if o.TopicName != nil {
		toSerialize["topicName"] = o.TopicName
	}
	if o.TopicTickers != nil {
		toSerialize["topicTickers"] = o.TopicTickers
	}
	if o.TopicLegalEntities != nil {
		toSerialize["topicLegalEntities"] = o.TopicLegalEntities
	}
	if o.OrderPriority != nil {
		toSerialize["orderPriority"] = o.OrderPriority
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Tickers != nil {
		toSerialize["tickers"] = o.Tickers
	}
	if o.TradingViewTicker != nil {
		toSerialize["tradingViewTicker"] = o.TradingViewTicker
	}
	if o.Figis != nil {
		toSerialize["figis"] = o.Figis
	}
	if o.CompanyDomains != nil {
		toSerialize["companyDomains"] = o.CompanyDomains
	}
	if o.Score != nil {
		toSerialize["_score"] = o.Score
	}
	if o.Highlight != nil {
		toSerialize["_highlight"] = o.Highlight
	}
	if o.Permalink != nil {
		toSerialize["permalink"] = o.Permalink
	}
	if o.IsinAliases != nil {
		toSerialize["isinAliases"] = o.IsinAliases
	}
	if o.AssetClassSearchKeywords != nil {
		toSerialize["assetClassSearchKeywords"] = o.AssetClassSearchKeywords
	}
	return json.Marshal(toSerialize)
}

type NullableMetadataRegularGlobal struct {
	value *MetadataRegularGlobal
	isSet bool
}

func (v NullableMetadataRegularGlobal) Get() *MetadataRegularGlobal {
	return v.value
}

func (v *NullableMetadataRegularGlobal) Set(val *MetadataRegularGlobal) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataRegularGlobal) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataRegularGlobal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataRegularGlobal(val *MetadataRegularGlobal) *NullableMetadataRegularGlobal {
	return &NullableMetadataRegularGlobal{value: val, isSet: true}
}

func (v NullableMetadataRegularGlobal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataRegularGlobal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}