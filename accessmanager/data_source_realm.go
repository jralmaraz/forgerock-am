package accessmanager

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/helper/logging"
)

func dataSourceRealms() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRealmsRead,
		Schema: map[string]*schema.Schema{
			"realms": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"_rev": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"parentpath": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"active": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"aliases": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alias": &schema.Schema{
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

	//client := &http.Client{Timeout: 10 * time.Second},
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // test server certificate is not trusted.
			},
		},
	}
	client.Transport = logging.NewTransport("ForgeRock", client.Transport)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/am/json/global-config/realms?_queryFilter=true", "https://forgerock.iam.dtt-iam.xyz"), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	//realms := make([]map[string]interface{}, 0)
	realms := new(Response)

	if err := json.NewDecoder(r.Body).Decode(realms); err != nil {
		log.Fatal(err)
		log.Output(err)
	}
	for i := range realms.Result {
		fmt.Printf("%v\n", realms.Result[i])
	}

	//err = json.NewDecoder(r.Body).Decode(&realms)
	//if err != nil {
	//	return diag.FromErr(err)
	//}

	//if err := d.Set("realms", realms); err != nil {
	//	return diag.FromErr(err)
	//}

	// always run
	//d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
