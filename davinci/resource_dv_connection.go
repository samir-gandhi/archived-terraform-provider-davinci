package davinci

import (
	"context"
	"fmt"
	"strconv"

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
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
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

	connection.Properties = *makeProperties(d)

	res, err := c.CreateInitializedConnection(&c.CompanyID, &connection)
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

	connId := d.Id()

	res, err := c.ReadConnection(&c.CompanyID, connId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", res.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("connection_id", res.ConnectionID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("connector_id", res.ConnectorID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_date", res.CreatedDate); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("company_id", res.CompanyID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("customer_id", res.CustomerID); err != nil {
		return diag.FromErr(err)
	}
	props, err := flattenConnectionProperties(&res.Properties)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("properties", props); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceConnectionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.Client)
	connId := d.Id()
	if d.HasChange("properties") {
		connection := dv.Connection{
			ConnectorID:  d.Get("connector_id").(string),
			Name:         d.Get("name").(string),
			ConnectionID: connId,
		}

		connection.Properties = *makeProperties(d)
		_, err := c.UpdateConnection(&c.CompanyID, &connection)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceConnectionRead(ctx, d, m)
}

func resourceConnectionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.Client)
	var diags diag.Diagnostics
	connId := d.Id()

	_, err := c.DeleteConnection(&c.CompanyID, connId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func flattenConnectionProperties(connectionProperties *dv.Properties) ([]map[string]interface{}, error) {
	if connectionProperties == nil {
		return nil, fmt.Errorf("no properties")
	}
	connProps := []map[string]interface{}{}
	for propName, propVal := range *connectionProperties {
		pMap := propVal.(map[string]interface{})
		if pMap == nil {
			return nil, fmt.Errorf("Unable to assert property values for %v\n", propName)
		}
		thisProp := map[string]interface{}{
			"name":  propName,
			"value": "",
		}
		if propType, ok := pMap["type"].(string); ok {
			// log.Printf("pType is: %v", pType)
			thisProp["type"] = propType
			switch propType {
			case "string", "":
				if _, ok := pMap["value"].(string); ok {
					thisProp["value"] = pMap["value"].(string)
				}
			case "boolean":
				if pValue, ok := pMap["value"].(bool); ok {
					thisProp["value"] = strconv.FormatBool(pValue)
				}
			default:
				return nil, fmt.Errorf("For Property '%v': unable to identify value type, only string or boolean is currently supported", thisProp["name"])
			}
		} else {
			switch pMap["value"].(type) {
			case string:
				if _, ok := pMap["value"].(string); ok {
					thisProp["value"] = pMap["value"].(string)
				}
			case bool:
				if pValue, ok := pMap["value"].(bool); ok {
					thisProp["value"] = strconv.FormatBool(pValue)
				}
			default:
				return nil, fmt.Errorf("For Property '%v': unable to identify value type, only string or boolean is currently supported", thisProp["name"])
			}
		}
		connProps = append(connProps, thisProp)
	}
	return connProps, nil
}

func makeProperties(d *schema.ResourceData) *dv.Properties {
	connProps := dv.Properties{}
	props := d.Get("properties").(*schema.Set).List()
	// fmt.Printf(props)
	for _, raw := range props {
		fmt.Printf("\nThis prop is: %v\n", raw)
		prop := raw.(map[string]interface{})
		connProps[prop["name"].(string)] = map[string]interface{}{
			"value": prop["value"].(string),
		}
	}
	return &connProps
}
