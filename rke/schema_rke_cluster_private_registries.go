package rke

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

//Schemas

func rkeClusterPrivateRegistriesFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"url": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Registry URL",
		},
		"is_default": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Set as default registry",
		},
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Registry password",
		},
		"user": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Registry user",
		},
		"ecr_credential_plugin": {
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Sensitive:   true,
			Description: "ECR credential plugin",
			Elem: &schema.Resource{
				Schema: rkeClusterPrivateRegistriesEcrCredentialPluginFields(),
			},
		},
	}
	return s
}

func rkeClusterPrivateRegistriesEcrCredentialPluginFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"aws_access_key_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "AWS Access Key ID for ECR access",
		},
		"aws_secret_access_key": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "AWS Secret Access Key for ECR access",
		},
	}
	return s
}
