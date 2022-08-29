package davinci

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

func dataSourceConnections() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConnectionsRead,
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
	}
}

func dataSourceConnectionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.Client)
	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Warning Message Summary",
		Detail:   "This is the detailed warning message from dataSourceConnectionRead",
	})

	// connectionID := d.Get("connection_id")
	// if connectionID == nil {
	// 	return diag.Errorf("error: connection_id is nil")
	// }
	fmt.Printf("COMPANY ID IS %v", c.CompanyID)
	resp, err := c.ReadConnections(&c.CompanyID, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	// orderItems := flattenConnItemsData(&order.Items)
	// if err := d.Set("items", orderItems); err != nil {
	// 	return diag.FromErr(err)
	// }
	// ois := make([]interface{}, len(*orderItems), len(*orderItems))

	conns := make([]interface{}, len(resp), len(resp))
	for i, connItem := range resp {
		thisResp := resp[i]
		thisResp.Properties = nil
		resp[i] = thisResp
		conn := make(map[string]interface{})
		conn["connection_id"] = connItem.ConnectionID
		conn["connector_id"] = connItem.ConnectorID
		conn["name"] = connItem.Name
		conn["created_date"] = connItem.CreatedDate
		conn["company_id"] = connItem.CompanyID
		//TODO implement properties
		// conn["properties"] = connItem.Properties
		conns[i] = conn
	}

	if err := d.Set("connections", conns); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
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
