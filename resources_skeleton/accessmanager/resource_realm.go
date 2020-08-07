package accessmanager

import "github.com/hashicorp/terraform/helper/schema"

func resourceRealm() *schema.Resource {
	return &schema.Resource{
		Create: resourceRealmCreate,
		Read:   resourceRealmRead,
		Update: resourceRealmUpdate,
		Delete: resourceRealmDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"active": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"parentPath": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"aliases": &schema.Schema{
				Type:     schema.TypeList,
				Required: false,
			},
		},
	}
}

func resourceRealmCreate(d *schema.ResourceData, m interface{}) error {
	realm := d.Get("realm").(string)
	d.SetId(realm)
	return resourceRealmRead(d, m)
}

func resourceRealmRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*MyClient)

	// Attempt to read from an upstream API
	obj, ok := client.Get(d.Id())

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if !ok {
		d.SetId("")
		return nil
	}

	d.Set("realm", obj.Address)
	return nil

}

func resourceRealmUpdate(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)

	if d.HasChange("realm") {
		// Try updating the address
		if err := updateRealm(d, m); err != nil {
			return err
		}

		d.SetPartial("realm")
	}

	// If we were to return here, before disabling partial mode below,
	// then only the "address" field would be saved.

	// We succeeded, disable partial mode. This causes Terraform to save
	// all fields again.
	d.Partial(false)

	return resourceRealmRead(d, m)
}

func resourceRealmDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
