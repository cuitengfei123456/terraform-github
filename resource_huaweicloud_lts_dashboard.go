package lts

// 改完已跑通 
import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/entity"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/httpclient_go"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils/fmtp"
	"io/ioutil"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceLtsDashboard() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLtsDashboardCreate,
		ReadContext:   resourceLtsDashboardRead,
		DeleteContext: resourceLtsDashboardDelete,
		UpdateContext: resourceDashboardUpdate,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_delete_charts": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"title": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"log_group_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"log_stream_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"detail": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_stream_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"template_title": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString}
			},
			"filters": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}
			},
			"template_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}
			},
			"last_update_time": {
				Type:     schema.TypeInt,
				Optional: true,
			},	
		}
	}
}

func resourceLtsDashBoardCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	client, diaErr := httpclient_go.NewHttpclientGo(config)
	if diaErr != nil {
		return diaErr
	}
	url := strings.Replace(config.Endpoints["lts"], "https//", "https://", -1) + "v2/" + 
	       config.HwClient.ProjectID + "/lts/template-dashboard"
	header := make(map[string]string)
	header["content-type"] = "application/json;charset=UTF8"
	dashBoardRequest := entity.DashBoardRequest{
		LogGroupId:    d.Get("log_group_id").(string),
		LogGroupName:  d.Get("log_group_name").(string),
		LogStreamId:   d.Get("log_stream_id").(string),
		LogStreamName: d.Get("log_stream_name").(string),
		TemplateTitle: utils.ExpandToStringList(d.Get("template_title").([]interface{})),
		TemplateType:  utils.ExpandToStringList(d.Get("template_type").([]interface{})),
		GroupName:     d.Get("group_name").(string),
	}
	client.WidthMethod(httpclient_go.MethodPost).WithUrl(url).WithHeader(header).WithBody(dashBoardRequest)
	response, err := client.Do()
	if err != nil {
		return diag.Errorf("error creating LtsDashBoard fields %s: %s", dashBoardRequest, err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return diag.Errorf("error convert data %s, %s", string(body), err)
	}
	if response.StatusCode == 201 {
		rlt := make([]entity.DashBoard, 0)
		err = json.Unmarshal(body, &rlt)
		if err != nil {
			return diag.Errorf("error convert data %s , %s", string(body), err)
		}
		d.SetId(rlt[0].Id)
		return  resourceLtsDashboardRead(ctx, d, meta)
	}
	return diag.Errorf("error creating LtsDashBoard Response %s: %s", dashBoardRequest, string(body))
}

func resourceLtsDashBoardRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	client, diaErr := httpclient_go.NewHttpclientGo(config)
	if diaErr != nil {
		return diaErr
	}
	url := strings.Replace(config.Endpoints["lts"], "https//", "https://", -1) + "v2/" + 
	    config.HwClient.ProjectID + "/dashboard?id=" + d.Id()
	header := make(map[string]string)
	header["content-type"] = "application/json;charset=UTF8"
	client.WidthMethod(httpclient_go.MethodGet).WithUrl(url).WithHeader(header)
	response, err := client.Do()
	body, diag := client.CheckDeletedDiag(d, err, response, "error LtsDahBoard")
	if body == nil {
		return diag
	}
	rlt := entity.ReadDashBoardResp{}
	err = json.Unmarshal(body, &rlt)
	mErr := multierror.Append(nil,
		d.Set("region", config.GetRegion(d)),
		d.Set("title",rlt.Results[0].Title),
	)
	if err := mErr.ErrorOrNil(); err != nil {
		return fmtp.DiagErrorf("error setting Lts dashboard fields: %s", err)
	}
	return nil
}

func resourceLtsDashBoardDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	client, diaErr := httpclient_go.NewHttpclientGo(config)
	if diaErr != nil {
		return diaErr
	}
	url := strings.Replace(config.Endpoints["lts"], "https//", "https://", -1) + "v2/" + 
	    config.HwClient.ProjectID + "/dashboard?id=" + d.Id() + "is_delete_charts=" + d.Get("is_delete_charts").(string)
	header := make(map[string]string)
	header["content-type"] = "application/json;charset=UTF8"
	client.WidthMethod(httpclient_go.MethodDelete).WithUrl(url).WithHeader(header)
	response, err := client.Do()
	if err != nil {
		return diag.Errorf("error delete LtsDashBoard %s: %s", d.Id(), err)
	}
	if response.StatusCode == 200 {
		return nil
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return diag.Errorf("error delete LtsDashBoard %s: %s", d.Id(), err)
	}
	return diag.Errorf("error delete LtsDashBoard %s:  %s", d.Id(), string(body))
}

// 伪代码
func resourceDashBoardUpdate(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	client, diaErr := httpclient_go.NewHttpclientGo(config)
	if diaErr != nil {
		return diaErr
	}
	url := strings.Replace(config.Endpoints["lts"], "https//", "https://", -1) + "v2/" + 
	    config.HwClient.ProjectID + "/dashboard?id=" + d.Id()
	header := make(map[string]string)
	header["content-type"] = "application/json;charset=UTF8"
	ChartsOpt := entity.DashboardChars{
		ChartId: "3347146c-5809-488a-a1f0-89ccaad0c0b8",
		Height:  4,
		Width:   2,
		XPos:    0,
		YPos:    0,
		Chart: map[string]interface{}{
			"config": map[string]interface{}{
				"backgroundColor":   "dark",
				"fontSize":          20,
				"gridBottomPadding": 0,
				"gridLeftPadding":   0,
				"gridRightPadding":  0,
				"gridTopPadding":    0,
				"heightReseted":     true,
				"hideLegend":        true,
				"isCurve":           false,
				"showSymbol":        false,
				"showXAxis":         false,
				"showYAxis":         false,
				"unit":              "K",
				"unitFontsize":      15,
				"xData":             "xData",
				"YData":             "PV",
			},
			"id":              "3347146c-5809-488a-a1f0-89ccaad0c0b8",
			"log_group_id":    "5d574e6a-87da-42bc-bfd4-ff61a1b336a4",
			"log_group_name":  "hzx",
			"log_stream_id":   "98de5d5a-9f54-4d01-9882-eca7bec99d09",
			"log_stream_name": "WAF",
		},
	}
	Opt:= entity.DashBoard{
		LastUpdateTime:    1660447346801,
		Id:                "a29ea063-36a7-4fc0-bd79-ee3826d3c334",
		Title:             "cui",
		UseSystemTemplate: true,
		Charts:            []entity.DashboardChars{ChartsOpt},
	}
	client.WidthMethod(httpclient_go.MethodPut).WithUrl(url).WithHeader(header).WithBody(Opt)
	response, err := client.Do()
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Request error", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read Body error", err)
	}
	res := make([]entity.DashBoard, 0)
	// var res *AomMappingRequestInfo
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("Unmarshal body error", err)
	}
	fmt.Println(res)
	return nil
}
