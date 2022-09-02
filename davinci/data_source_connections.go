package davinci

import (
	"context"
	"fmt"
	"log"
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
						"properties": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prop_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"value_string": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"value_bool": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
							// Elem: &schema.Schema{
							// 	Type:     schema.TypeSet,
							// 	MaxItems: 1,
							// 	Optional: true,
							// 	Elem: &schema.Resource{
							// 		Schema: map[string]*schema.Schema{
							// 			"value_string": {
							// 				//TODO: Value is not always string.
							// 				//In SDK it's defined at runtime.
							// 				Type:     schema.TypeString,
							// 				Optional: true,
							// 			},
							// 		},
							// 	},
							// },

							// Elem: &schema.Resource{
							// 	Schema: map[string]*schema.Schema{
							// 		"prop_name": {
							// 			Type:     schema.TypeSet,
							// 			Optional: true,
							// 			MaxItems: 1,
							// 			//TODO: implement all possibilities here
							// 			Elem: &schema.Resource{
							// 				Schema: map[string]*schema.Schema{
							// "display_name": {
							// 	Type:     schema.TypeString,
							// 	Optional: true,
							// },
							// "info": {
							// 	Type:     schema.TypeString,
							// 	Optional: true,
							// },
							// "description": {
							// 	Type:     schema.TypeString,
							// 	Optional: true,
							// },
							// "preferred_control_type": {
							// 	Type:     schema.TypeList,
							// 	Optional: true,
							// 	MaxItems: 1,
							// 	Elem: &schema.Schema{
							// 		Type: schema.TypeString,
							// 	},
							// },
							// "enableParameters": {
							// 	Type:     schema.TypeBool,
							// 	Optional: true,
							// },
							// // implementing only value_string for now.

							// "placeholder": {
							// 	Type:     schema.TypeString,
							// 	Optional: true,
							// },
							// "placeholderAdd": {
							// 	Type:     schema.TypeString,
							// 	Optional: true,
							// },
							// "userViewPreferredControlType": {
							// 	Type:     schema.TypeString,
							// 	Optional: true,
							// },
							// "constructType": {
							// 	Type:     schema.TypeList,
							// 	Optional: true,
							// 	MaxItems: 1,
							// 	Elem: &schema.Schema{
							// 		Type: schema.TypeString,
							// 	},
							// },

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
	log.Printf("COMPANY ID IS %v", c.CompanyID)
	resp, err := c.ReadConnections(&c.CompanyID, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	// orderItems := flattenConnItemsData(&order.Items)
	// if err := d.Set("items", orderItems); err != nil {
	// 	return diag.FromErr(err)
	// }
	// ois := make([]interface{}, len(*orderItems), len(*orderItems))

	// log.Printf("len(resp) is: %v\n", len(resp))
	conns := make([]interface{}, len(resp), len(resp))
	for i, connItem := range resp {
		conn := make(map[string]interface{})
		conn = map[string]interface{}{
			"connection_id": connItem.ConnectionID,
			"connector_id":  connItem.ConnectorID,
			"name":          connItem.Name,
			"created_date":  connItem.CreatedDate,
			"company_id":    connItem.CompanyID,
		}
		// log.Println("Im here 1")

		if connItem.Properties != nil {
			connProps := []map[string]interface{}{}
			// fmt.Printf("connItem.Properties is: %v \n", connItem.Properties)
			for propi, propv := range connItem.Properties {
				// log.Printf("v is %v\n", propv)
				// fmt.Printf("checking propi: %v \n and propv: %q \n", propi, propv)
				pMap := propv.(map[string]interface{})
				// log.Printf("here at pMap")
				if pMap == nil {
					return diag.Errorf("Unable to assert Property to map interface")
				}
				// log.Printf("here at thisProp")
				thisProp := map[string]interface{}{
					"prop_name": propi,
				}
				// log.Printf("conns is: %q \n", conns)
				// fmt.Printf("checking propi: %v \n and pMap[value]: %v \n", propi, pMap["value"])
				// log.Printf("here at pValue\n")
				if pValue, ok := pMap["value"]; ok {
					if pType, ok := pMap["type"]; ok {
						// log.Printf("here at pValue: %v\n", pType)
						// log.Printf("here at pMap: %q\n", pMap)
						// pType = pMap["type"].(string)
						switch {
						case pType == "string":
							thisProp["value_string"] = pMap["value"].(string)
						case pType == "boolean":
							thisProp["value_bool"] = pMap["value"].(bool)
						default:
							return diag.Errorf("Unable to identify type of value in %v", thisProp["prop_name"])
						}
					} else {
						switch pValue.(type) {
						case string:
							thisProp["value_string"] = pMap["value"].(string)
						case bool:
							thisProp["value_bool"] = pMap["value"].(bool)
						default:
							return diag.Errorf("Unable to identify type of value in %v", thisProp["prop_name"])
						}
					}
				} else {
					if pType, ok := pMap["type"].(string); ok {
						switch {
						case pType == "string":
							thisProp["value_string"] = ""
						case pType == "boolean":
							thisProp["value_bool"] = false
						default:
							return diag.Errorf("Unable to identify type of value in %v", thisProp["prop_name"])
						}
					} else {
						switch thisProp["value_string"].(type) {
						case string:
							thisProp["value_string"] = ""
						case bool:
							thisProp["value_bool"] = false
						default:
							return diag.Errorf("Unable to identify type of value in %v", thisProp["prop_name"])
						}
					}
				}

				connProps = append(connProps, thisProp)
			}
			conn["properties"] = connProps
			// conn["properties"] = []map[string]interface{}{
			// 	{"prop_name": "aws_secret"}}
		}
		conns[i] = conn
		fmt.Printf("conns[%v] issuccessful \n", i)
		fmt.Printf("conns[i] is: %q \n", conns[i])
		fmt.Printf("conn is: %q \n", conn)
	}

	// fmt.Printf("TRYING SETTINGCONN: %v \n", conns)
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
