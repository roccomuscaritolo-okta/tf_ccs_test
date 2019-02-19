package oktaccs

import (
	"net/http"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"application": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "okta",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"oktaccs_config_server_secret": dataSourceCloudConfigServerSecrets(),
		},

		ConfigureFunc: providerConfigure,
	}
}

// ccsClient provides Client credentials
type ccsClient struct {
	Hostname    string
	Username    string
	Password    string
	Application string
	httpClient  *http.Client
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	client := ccsClient{
		Hostname:    d.Get("hostname").(string),
		Username:    d.Get("username").(string),
		Password:    d.Get("password").(string),
		Application: d.Get("application").(string),
		httpClient:  &http.Client{Timeout: time.Second * 5},
	}

	return &client, nil
}
