package accessmanager

// Reference: https://backstage.forgerock.com/knowledge/kb/article/a60067323
import "github.com/hashicorp/terraform/helper/schema"

func resourceIGRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceIGRouteCreate,
		Read:   resourceIGRouteRead,
		Update: resourceIGRouteUpdate,
		Delete: resourceIGRouteDelete,

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

func resourceIGRouteCreate(d *schema.ResourceData, m interface{}) error {
	realm := d.Get("realm").(string)
	d.SetId(realm)
	return resourceIGRouteRead(d, m)
}

func resourceIGRouteRead(d *schema.ResourceData, m interface{}) error {
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

func resourceIGRouteUpdate(d *schema.ResourceData, m interface{}) error {
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

	return resourceIGRouteRead(d, m)
}

func resourceIGRouteDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
