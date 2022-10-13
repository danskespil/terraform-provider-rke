package rke

import (
	rancher "github.com/rancher/rke/types"
)

// Flatteners

func flattenRKEClusterPrivateRegistries(p []rancher.PrivateRegistry) []interface{} {
	out := []interface{}{}

	for _, in := range p {
		obj := make(map[string]interface{})
		obj["is_default"] = in.IsDefault

		if len(in.Password) > 0 {
			obj["password"] = in.Password
		}

		if len(in.URL) > 0 {
			obj["url"] = in.URL
		}

		if len(in.User) > 0 {
			obj["user"] = in.User
		}

		if in.ECRCredentialPlugin != nil {
			if len(in.ECRCredentialPlugin.AwsAccessKeyID) > 0 {
				obj["aws_access_key_id"] = in.ECRCredentialPlugin.AwsAccessKeyID
			}

			if len(in.ECRCredentialPlugin.AwsSecretAccessKey) > 0 {
				obj["aws_secret_access_key"] = in.ECRCredentialPlugin.AwsSecretAccessKey
			}

			if len(in.ECRCredentialPlugin.AwsSessionToken) > 0 {
				obj["aws_session_token"] = in.ECRCredentialPlugin.AwsSessionToken
			}
		}

		out = append(out, obj)
	}

	return out
}

// Expanders

func expandRKEClusterPrivateRegistries(p []interface{}) []rancher.PrivateRegistry {
	out := []rancher.PrivateRegistry{}
	if len(p) == 0 || p[0] == nil {
		return out
	}

	for i := range p {
		in := p[i].(map[string]interface{})
		obj := rancher.PrivateRegistry{}

		if v, ok := in["is_default"].(bool); ok {
			obj.IsDefault = v
		}

		if v, ok := in["password"].(string); ok && len(v) > 0 {
			obj.Password = v
		}

		if v, ok := in["url"].(string); ok && len(v) > 0 {
			obj.URL = v
		}

		if v, ok := in["user"].(string); ok && len(v) > 0 {
			obj.User = v
		}

		if v, ok := in["aws_access_key_id"].(string); ok && len(v) > 0 {
			ecrCredentialPlugin := &rancher.ECRCredentialPlugin{}

			ecrCredentialPlugin.AwsAccessKeyID = v

			if v, ok := in["aws_secret_access_key"].(string); ok && len(v) > 0 {
				ecrCredentialPlugin.AwsSecretAccessKey = v

				if v, ok := in["aws_session_token"].(string); ok && len(v) > 0 {
					ecrCredentialPlugin.AwsSessionToken = v
				}
			}

			obj.ECRCredentialPlugin = ecrCredentialPlugin
		}

		out = append(out, obj)
	}

	return out
}
