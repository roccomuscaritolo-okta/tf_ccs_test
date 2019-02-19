package oktaccs

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/magiconair/properties"
)

func dataSourceCloudConfigServerSecrets() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCloudConfigServerSecretsRead,

		Schema: map[string]*schema.Schema{
			"profiles": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"properties": {
				Type:      schema.TypeMap,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func dataSourceCloudConfigServerSecretsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ccsClient)

	// declare empty StringSlice
	var profileSlice []string

	// get values from profiles attribute
	profileSet := d.Get("profiles").(*schema.Set)

	// append values from *schema.Set into slice
	for _, v := range profileSet.List() {
		profileSlice = append(profileSlice, v.(string))
	}

	// join profiles into a single comma separated string
	profiles := strings.Join(profileSlice, ",")

	// Example rendered url
	// url := "https://ct0.cloud-config.auw2l.internal/master/okta-monolith_ct1,monolith_ct2.properties"
	url := "https://" + client.Hostname + "/master/" + client.Application + "-" + profiles + ".properties"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth(client.Username, client.Password)

	resp, err := client.httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	s := string(bodyBytes)

	p := properties.MustLoadString(s)

	d.SetId(fmt.Sprintf("%s: %s", client.Hostname, profiles))
	d.Set("properties", p.Map())

	return nil
}
