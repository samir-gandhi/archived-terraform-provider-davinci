package davinci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConnection() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConnectionCreate,
		ReadContext:   resourceConnectionRead,
		UpdateContext: resourceConnectionUpdate,
		DeleteContext: resourceConnectionDelete,
		Schema: map[string]*schema.Schema{
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connections": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"connection_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"connector_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
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
										Computed: true,
									},
									"created_date": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"properties": &schema.Schema{
										Type:     schema.TypeMap,
										Computed: true,
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
					},
				},
			},
		},
	}
}

func resourceConnectionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m(*hc.Client)

	var diags diag.Diagnostics

	items := d.Get("items").([]interface{})
	cis := []dv.Connection{}

	for _, item := range items {
		i := item.(map[string]interface{})
		co := i["connection"]
	}

	return diags
}

func resourceConnectionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

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
