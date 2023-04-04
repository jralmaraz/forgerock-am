package accessmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	amclient "github.com/jralmaraz/forgerock-go-sdk"
)

func dataSourceRealms() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRealmsRead,
		Schema: map[string]*schema.Schema{
			"realms": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"_rev": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parentpath": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"active": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"aliases": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceRealmsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*amclient.Client)

	// client.Transport = logging.NewTransport("ForgeRock", client.Transport)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/am/json/global-config/realms?_queryFilter=true", client.HostURL), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	//Just to test without releasing a new client version
	req.Header.Set("Accept-API-Version", "resource=1.0, protocol=2.1")
	r, err := client.DoRequest(req)
	if err != nil {
		return diag.FromErr(err)
	}

	realms := new(amclient.Response)

	if err := json.Unmarshal(r, &realms); err != nil {
		return diag.FromErr(err)
	}

	if len(realms.Result) == 0 {
		// Set an empty slice to the 'realms' key
		if err := d.Set("realms", []interface{}{}); err != nil {
			return diag.FromErr(err)
		}
	} else {
		if err := d.Set("realms", mapResponseResult(&realms.Result)); err != nil {
			return diag.FromErr(err)
		}
	}
	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func mapResponseResult(result *[]amclient.Realm) []interface{} {
	if result != nil {
		realms := make([]interface{}, len(*result), len(*result))

		for i, realm := range *result {
			realms[i] = mapRealm(realm)
		}

		return realms
	}

	return make([]interface{}, 0)
}

func mapRealm(realm amclient.Realm) map[string]interface{} {
	r := make(map[string]interface{})
	r["_id"] = realm.ID
	r["_rev"] = realm.Rev
	r["parentpath"] = realm.ParentPath
	r["active"] = realm.Active
	r["name"] = realm.Name
	r["aliases"] = realm.Aliases

	return r
}
