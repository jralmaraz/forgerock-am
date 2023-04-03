package accessmanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	accessmanagerclient "github.com/jralmaraz/forgerock-go-sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("XOpenAMUsername", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("XOpenAMPassword", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"realm": &schema.Resource{
				Create: resourceRealmCreate,
				Schema: map[string]*schema.Schema{
					"name": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
					"parent_path": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
					"active": &schema.Schema{
						Type:     schema.TypeBool,
						Required: true,
					},
					"aliases": &schema.Schema{
						Type:     schema.TypeList,
						Required: true,
					},
				},
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"accessmanager_realms": dataSourceRealms(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	//amadminSsotoken := d.Get("amadmin_ssotoken").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (username != "") && (password != "") { //&& (amadminSsotoken != "") {
		c, err := accessmanagerclient.NewClient(nil, &username, &password)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return c, diags
	}

	c, err := accessmanagerclient.NewClient(nil, nil, nil)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}
