package octopusdeploy

import (
	"fmt"

	"github.com/MattHodge/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataMachinePolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataMachinePolicyReadByName,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"isdefault": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataMachinePolicyReadByName(d *schema.ResourceData, m interface{}) error {
	client := m.(*octopusdeploy.Client)

	policyName := d.Get("name").(string)
	policies, err := client.MachinePolicy.GetAll()
	if err == octopusdeploy.ErrItemNotFound {
		return nil
	}
	if err != nil {
		return fmt.Errorf("error reading machine policy with name %s: %s", policyName, err.Error())
	}

	for _, p := range *policies {
		if p.Name == policyName {
			d.SetId(p.ID)
			d.Set("description", p.Description)
			d.Set("isdefault", p.IsDefault)
		}
	}

	return nil
}
