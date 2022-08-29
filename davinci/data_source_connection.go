package davinci

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

func dataSourceConnection() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConnectionRead,
		Schema: map[string]*schema.Schema{
			"dv_connection": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connection_id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
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
	}
}

func dataSourceConnectionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.Client)

	var diags diag.Diagnostics

	connectionID := d.Get("connection_id").(string)

	conn, err := c.ReadConnection(&c.CompanyID, connectionID)
	if err != nil {
		return diag.FromErr(err)
	}

	// orderItems := flattenConnItemsData(&order.Items)
	// if err := d.Set("items", orderItems); err != nil {
	// 	return diag.FromErr(err)
	// }
	ois := make([]interface{}, len(*orderItems), len(*orderItems))

	if err := d.Set("conn", conn); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(conn.CreatedDate, 10))
	return diags
}

// func flattenOrderItemsData(orderItems *[]dv.Connection) []interface{} {
// 	if orderItems != nil {
// 		ois := make([]interface{}, len(*orderItems), len(*orderItems))

// 		for i, orderItem := range *orderItems {
// 			oi := make(map[string]interface{})

// 			oi["coffee_id"] = orderItem.Coffee.ID
// 			oi["coffee_name"] = orderItem.Coffee.Name
// 			oi["coffee_teaser"] = orderItem.Coffee.Teaser
// 			oi["coffee_description"] = orderItem.Coffee.Description
// 			oi["coffee_price"] = orderItem.Coffee.Price
// 			oi["coffee_image"] = orderItem.Coffee.Image
// 			oi["quantity"] = orderItem.Quantity

// 			ois[i] = oi
// 		}

// 		return ois
// 	}

// 	return make([]interface{}, 0)
// }
