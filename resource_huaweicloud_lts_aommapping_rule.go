package lts

// 改完已跑通
import (
	"context"
	"encoding/json"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/entity"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/httpclient_go"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils/fmtp"
	"io/ioutil"
	"strings"
)

func ResourceAomMappingRule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAomMappingRuleCreate,
		ReadContext:   resourceAomMappingRuleRead,
		DeleteContext:   resourceAomMappingRuleDelete,
		UpdateContext:   resourceAomMappingRuleUpdate,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_batch": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rule_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cluster_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name_space": {
				Type:     schema.TypeString,
				Required: true,
			},
			"container_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"deployments": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"files": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema:map[string]*schema.Schema{
						"file_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"log_stream_info": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema:map[string]*schema.Schema{
									"target_log_group_id": {
										Type: schema.TypeString,
										Required:true,
									},
									"target_log_group_name": {
										Type: schema.TypeString,
										Required:true,
									},
									"target_log_stream_id": {
										Type: schema.TypeString,
										Required:true,
									},
									"target_log_stream_name": {
										Type: schema.TypeString,
										Required:true,
									},
								}
							}
						}
					}
				}
			}
		}
	}
}
func buildLogStreamOpts(rawRule []interface{}) entity.AomMappingLogStreamInfo {
	s := rawRule[0].(map[string]interface{})
	rst := entity.AomMappingLogStreamInfo{
		TargetLogGroupId:    s["target_log_group_id"].(string),
		TargetLogGroupName:  s["target_log_group_name"].(string),
		TargetLogStreamId:   s["target_log_stream_id"].(string),
		TargetLogStreamName: s["target_log_stream_name"].(string),
	}
	return rst
}

func buildFileOpts(rawRules []interface{}) []entity.AomMappingfilesInfo {
	file := make([]entity.AomMappingfilesInfo,len(rawRules))
	for i, v := range rawRules {
		rawRule := v.(map[string]interface{})
		file[i].FileName = rawRule["file_name"].(string)
		file[i].LogStreamInfo = buildLogStreamOpts(rawRule["log_stream_info"].([]interface{}))
	}
	return file
}

func resourceAomMappingRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	client, diaErr := httpclient_go.NewHttpclientGo(config)
	if diaErr != nil {
		return diaErr
	}
	url := strings.Replace(config.Endpoints["lts"], "https//", "https://", -1) + "v2/" + config.HwClient.ProjectID + 
	       "/lts/aom-mapping" + "?isBatch" + d.Get("is_batch").(string)
	aomMappingRequestInfo := entity.AomMappingRequestInfo{
		ProjectId: config.HwClient.ProjectID,
		RuleName: d.Get("rule_name").(string),
		RuleInfo: entity.AomMappingRuleInfo{
			ClusterId:   d.Get("cluster_id").(string),
			ClusterName: d.Get("cluster_name").(string),
			Namespace:   d.Get("name_space").(string),
			Deloyments:  utils.ExpandToStringList(d.Get("deployments").([]interface{})),
			Files:       buildFileOpts(d.Get("files").([]interface{}))
		},
	}
	client.WidthMethod(httpclient_go.MethodPost).WithUrl(url).WithBody(aomMappingRequestInfo)
	response, err := client.Do()
	if err != nil {
		return diag.Errorf("error creating AomMappingRule fields %s: %s", aomMappingRequestInfo, err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return diag.Errorf("error convert data %s , %s", string(body), err)
	}
	if response.StatusCode == 201 {
		rlt := make([]entity.AomMappingRuleResp, 0)
		err = json.Unmarshal(body, &rlt)
		if err != nil {
			return diag.Errorf("error convert data %s , %s", string(body), err)
		}
		return resourceAomMappingRuleRead(ctx, d, meta)
	}
	d.SetId(rlt[0].Ruled)
	return diag.Errorf("error AomMappingRule Response %s : %s", aomMappingRequestInfo, string(body))
}

func resourceAomMappingRuleRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	client, diaErr := httpclient_go.NewHttpclientGo(config)
	if diaErr != nil {
		return diaErr
	}
	url := strings.Replace(config.Endpoints["lts"], "https//", "https://", -1) + "v2/" + 
	config.HwClient.ProjectID + "/lts/aom-mapping" + d.Id()
	header := make(map[string]string)
	header["content-type"] = "application/json;charset=UTF8"
	client.WidthMethod(httpclient_go.MethodGet).WithUrl(url).WithHeader(header)
	response, err := client.Do()
	body, diag := client.CheckDeletedDiag(d, err, response, "error AomMappingRule")
	if body == nil {
		return diag
	}
	rlt := make([]entity.AomMappingRequestInfo, 0)
	err = json.Unmarshal(body, &rlt)
	mErr := multierror.Append(nil,
		d.Set("rule_name", rlt[0].RuleName),
		d.Set("cluster_id",rlt[0].RuleInfo.ClusterId),
		d.Set("cluster_name",rlt[0].RuleInfo.ClusterName),
		d.Set("container_name",rlt[0].RuleInfo.ContainerName),
		d.Set("deployments",rlt[0].RuleInfo.Deployments),
	)
	if err := mErr.ErrorOrNil(); err != nil {
		return fmtp.DiagErrorf("error AomMappingRule fields: %w", err)
	}
	return nil
}

func resourceAomMappingRuleDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	client, diaErr := httpclient_go.NewHttpclientGo(config)
	if diaErr != nil {
		return diaErr
	}
	url := strings.Replace(config.Endpoints["lts"], "https//", "https://", -1) + "v2/" + config.HwClient.ProjectID + 
	       "/lts/aom-mapping?id=" + d.Id()
	header := make(map[string]string)
	header["content-type"] = "application/json;charset=UTF8"
	client.WidthMethod(httpclient_go.MethodDelete).WithUrl(url).WithHeader(header)
	response, err := client.Do()
	if err != nil {
		return diag.Errorf("error delete AomMappingRule %s: %s", d.Id(), err)
	}
	if response.StatusCode == 200 {
		return nil
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return diag.Errorf("error delete AomMappingRule %s: %s", d.Id(), err)
	}
	return diag.Errorf("error delete AomMappingRule %s:  %s", d.Id(), string(body))
}

func resourceAomMappingRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	client, diaErr := httpclient_go.NewHttpclientGo(config)
	if diaErr != nil {
		return diaErr
	}
	url := strings.Replace(config.Endpoints["lts"], "https//", "https://", -1) + "v2/" + config.HwClient.ProjectID + 
	       "/lts/aom-mapping"
	createOpts := entity.AomMappingRequestInfo{
		ProjectId: config.HwClient.ProjectID,
		RuleId:    d.Id(),
		RuleName:  d.Get("rule_name").(string),
		RuleInfo:  entity.AomMappingRuleInfo{
			ClusterId:   d.Get("cluster_id").(string),
			ClusterName: d.Get("cluster_name").(string),
			Namespace:   d.Get("name_space").(string),
			Deloyments:  utils.ExpandToStringList(d.Get("deployments").([]interface{})),
			Files:       buildFileOpts(d.Get("files").([]interface{}))
		},
	}
	client.WidthMethod(httpclient_go.MethodPut).WithUrl(url).WithBody(createOpts)
	response, err := client.Do()
	if err != nil {
		return diag.Errorf("error update AomMappingRule fields %s: %s", createOpts.RuleName, err)
	}
	d.SetId(createOpts.RuleId)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return diag.Errorf("error update AomMappingRule %s: %s", string(body), err)
	}
	return diag.Errorf("error update AomMappingRule %s:  %s", createOpts, err)
}
