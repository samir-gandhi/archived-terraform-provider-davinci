package davinci

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/samir-gandhi/davinci-client-go/davinci"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DAVINCI_USERNAME", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("DAVINCI_PASSWORD", nil),
			},
			"company_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("DAVINCI_COMPANY_ID", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"davinci_connections": dataSourceConnections(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	company_id := d.Get("company_id").(string)

	var diags diag.Diagnostics
	fmt.Printf("company_id is: %v\n", company_id)
	// fmt.Printf("c.CompanyID is: %v\n", c.CompanyID)
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Provider Info",
		Detail:   "This is the detailed warning message from providerConfigure",
	})

	if (username != "") && (password != "") {
		// fmt.Printf("username is: %s", username)
		c, err := davinci.NewClient(nil, &username, &password)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create Davinci client",
				Detail:   "Unable to auth user",
			})
			return nil, diags
		}

		return c, diags
	}
	c, err := davinci.NewClient(nil, nil, nil)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	fmt.Printf("company_id is: %v\n", company_id)
	fmt.Printf("c.CompanyID is: %v\n", c.CompanyID)

	if company_id != "" {
		c.CompanyID = company_id
	}

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "FooBar",
		Detail:   "This is the detailed warning message from providerConfigure",
	})

	return c, diags
}
