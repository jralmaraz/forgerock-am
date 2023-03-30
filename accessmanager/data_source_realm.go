package accessmanager

import (
	"context"
	"crypto/tls"
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
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alias": {
										Type:     schema.TypeString,
										Computed: true,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceRealmsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	type Response struct {
		Result []struct {
			ID         string      `json:"_id"`
			Rev        string      `json:"_rev"`
			ParentPath interface{} `json:"parentPath"`
			Active     bool        `json:"active"`
			Name       string      `json:"name"`
			Aliases    []string    `json:"aliases"`
		} `json:"result"`
		ResultCount             int         `json:"resultCount"`
		PagedResultsCookie      interface{} `json:"pagedResultsCookie"`
		TotalPagedResultsPolicy string      `json:"totalPagedResultsPolicy"`
		TotalPagedResults       int         `json:"totalPagedResults"`
		RemainingPagedResults   int         `json:"remainingPagedResults"`
	}

	client := m.(*amclient.Client)

	client.HTTPClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // test server certificate is not trusted.
		},
	}

	// client.Transport = logging.NewTransport("ForgeRock", client.Transport)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/am/json/global-config/realms?_queryFilter=true", "https://dev.example.com"), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	//Just to test without releasing a new client version
	req.Header.Set("Accept-API-Version", "resource=1.0, protocol=2.1")
	r, err := client.DoRequest(req)
	if err != nil {
		return diag.FromErr(err)
	}

	//realms := make([]map[string]interface{}, 0)
	realms := new(Response)

	if err := json.Unmarshal(r, &realms); err != nil {
		diag.FromErr(err)
		//log.Output(2, err)

	}

	for i := range realms.Result {
		fmt.Printf("%v\n", realms.Result[i])
	}

	//err = json.NewDecoder(r.Body).Decode(&realms)
	//if err != nil {
	//	return diag.FromErr(err)
	//}

	if err := d.Set("realms", realms); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
