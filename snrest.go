package snrest

import (
	"net/http"

	"github.com/dghubble/sling"
)

type GetParams struct {
	Query                string `url:"sysparm_query,omitempty"`
	DisplayValue         string `url:"sysparm_display_value,omitempty"`
	Fields               string `url:"sysparm_fields,omitempty"`
	View                 string `url:"sysparm_view,omitempty"`
	Limit                int    `url:"sysparm_limit,omitempty"`
	Offset               int    `url:"sysparm_offset,omitempty"`
	ExcludeReferenceLink bool   `url:"sysparm_exclude_reference_link,omitempty"`
}

type PostParams struct {
	DisplayValue         string `url:"sysparm_display_value,omitempty"`
	Fields               string `url:"sysparm_fields,omitempty"`
	View                 string `url:"sysparm_view,omitempty"`
	ExcludeReferenceLink bool   `url:"sysparm_exclude_reference_link,omitempty"`
	InputDisplayValue    bool   `url:"sysparm_input_display_value,omitempty"`
}

type PutParams struct {
	DisplayValue         string `url:"sysparm_display_value,omitempty"`
	Fields               string `url:"sysparm_fields,omitempty"`
	View                 string `url:"sysparm_view,omitempty"`
	ExcludeReferenceLink bool   `url:"sysparm_exclude_reference_link,omitempty"`
	InputDisplayValue    bool   `url:"sysparm_input_display_value,omitempty"`
}

type JsonV2Params struct {
	Action           string `url:"sysparm_action,omitempty"`
	SysID            string `url:"sysparm_sys_id,omitempty"`
	Query            string `url:"sysparm_query,omitempty"`
	View             string `url:"sysparm_view,omitempty"`
	RecordCount      int    `url:"sysparm_record_count,omitempty"`
	DisplayValue     string `url:"displayvalue,omitempty"`
	DisplayVariables bool   `url:"displayvariables,omitempty"`
}

type RestClient struct {
	sling *sling.Sling
}

func New(instance string, user string, pass string) *RestClient {
	baseUrl := "https://" + instance + ".service-now.com/"
	return &RestClient{
		sling: sling.New().Client(&http.Client{}).Base(baseUrl).Set("Accept", "application/json").SetBasicAuth(user, pass),
	}
}

func (c *RestClient) GetKeys(table string, params JsonV2Params) ([]string, error) {
	type GetKeysResult struct {
		Records []string `json:"records"`
	}
	params.Action = "getKeys"
	result := GetKeysResult{}
	_, err := c.sling.Get(table + ".do?JSONv2").QueryStruct(params).ReceiveSuccess(&result)
	return result.Records, err
}

func (c *RestClient) Get(table string, sysID string, params GetParams, successV interface{}) (*http.Response, error) {
	return c.sling.Get("/api/now/table/" + table + "/" + sysID).QueryStruct(params).ReceiveSuccess(successV)
}

func (c *RestClient) GetMultiple(table string, params GetParams, successV interface{}) (*http.Response, error) {
	return c.sling.Get("/api/now/table/" + table).QueryStruct(params).ReceiveSuccess(successV)
}

func (c *RestClient) Insert(table string, params PostParams, data interface{}, successV interface{}) (*http.Response, error) {
	return c.sling.Post("/api/now/table/" + table).QueryStruct(params).BodyJSON(data).ReceiveSuccess(successV)
}

func (c *RestClient) InsertMultiple(table string, params JsonV2Params, data interface{}, successV interface{}) (*http.Response, error) {
	return c.sling.Post("/" + table + ".do?JSONv2").QueryStruct(params).BodyJSON(data).ReceiveSuccess(successV)
}

func (c *RestClient) Update(table string, params PutParams, data interface{}, successV interface{}) (*http.Response, error) {
	return c.sling.Put("/api/now/table/" + table).QueryStruct(params).BodyJSON(data).ReceiveSuccess(successV)
}

func (c *RestClient) Delete(table string, sysID string) (*http.Response, error) {
	type Dummy struct{}
	d := Dummy{}
	return c.sling.Delete("/api/now/table/" + table + "/" + sysID).ReceiveSuccess(d)
}

func (c *RestClient) DeleteMultiple(table string, params JsonV2Params) (int, error) {
	type DeleteMultipleResult struct {
		Records []struct {
			Count int `json:"count"`
		} `json:"records"`
	}
	params.Action = "deleteMultiple"
	result := DeleteMultipleResult{}
	_, err := c.sling.Get("/" + table + ".do?JSONv2").QueryStruct(params).ReceiveSuccess(&result)
	return result.Records[0].Count, err
}
