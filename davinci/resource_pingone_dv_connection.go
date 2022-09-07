package davinci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

func resourceConnection() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConnectionCreate,
		ReadContext:   resourceConnectionRead,
		UpdateContext: resourceConnectionUpdate,
		DeleteContext: resourceConnectionDelete,
		Schema: map[string]*schema.Schema{
			"connection_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"connector_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"company_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"customer_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"created_date": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"properties": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			//TODO - implement properties
			// "properties": &schema.Schema{
			// 	Type:     schema.TypeMap,
			// 	Computed: true,
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"awsAccessKey": &schema.Schema{
			// 				Type:     schema.TypeMap,
			// 				Computed: true,
			// 				Elem:     &schema.Resource{
			// 					Schema: map[string]*schema.Schema{
			// 						"type":
			// 					},
			// 				},
			// 			},
			// 		},
			// 	},
			// },
		},
	}
}

func resourceConnectionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.Client)

	var diags diag.Diagnostics

	connection := dv.Connection{
		ConnectorID: d.Get("connector_id").(string),
		Name:        d.Get("name").(string),
	}
	res, err := c.CreateConnection(&c.CompanyID, &connection)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(res.ConnectionID)

	resourceConnectionRead(ctx, d, m)

	return diags
}

func resourceConnectionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.Client)

	var diags diag.Diagnostics

	id := d.Id()

	res, err := c.ReadConnection(&c.CompanyID, id)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", res.Name); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceConnectionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}

func resourceConnectionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}
