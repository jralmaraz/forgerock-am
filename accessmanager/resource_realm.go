package accessmanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	amclient "github.com/jralmaraz/forgerock-go-sdk"
)

func resourceRealmCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*amclient.Client)

	// Read the input values from the resource data
	name := d.Get("name").(string)
	parentPath := d.Get("parent_path").(string)

	// Call the Access Manager API to create the realm
	realm, err := c.CreateRealm(name, parentPath)
	if err != nil {
		return diag.FromErr(err)
	}

	// Save the new realm ID to the resource data
	d.SetId(realm.ID)

	// Return nil for success
	return nil
}
