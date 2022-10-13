package rke

import (
	"reflect"
	"testing"

	rancher "github.com/rancher/rke/types"
)

var (
	testRKEClusterPrivateRegistriesConf      []rancher.PrivateRegistry
	testRKEClusterPrivateRegistriesInterface []interface{}
)

func init() {
	testRKEClusterPrivateRegistriesConf = []rancher.PrivateRegistry{
		{
			IsDefault: true,
			Password:  "XXXXXXXX",
			URL:       "url.terraform.test",
			User:      "user",
		},
		{
			IsDefault: true,
			URL:       "url.terraform.test",
			ECRCredentialPlugin: &rancher.ECRCredentialPlugin{
				AwsAccessKeyID:     "key",
				AwsSecretAccessKey: "secret",
			},
		},
		{
			IsDefault: true,
			URL:       "url.terraform.test",
			ECRCredentialPlugin: &rancher.ECRCredentialPlugin{
				AwsAccessKeyID:     "key",
				AwsSecretAccessKey: "secret",
				AwsSessionToken:    "token",
			},
		},
	}
	testRKEClusterPrivateRegistriesInterface = []interface{}{
		map[string]interface{}{
			"is_default": true,
			"password":   "XXXXXXXX",
			"url":        "url.terraform.test",
			"user":       "user",
		},
		map[string]interface{}{
			"is_default":            true,
			"url":                   "url.terraform.test",
			"aws_access_key_id":     "key",
			"aws_secret_access_key": "secret",
		},
		map[string]interface{}{
			"is_default":            true,
			"url":                   "url.terraform.test",
			"aws_access_key_id":     "key",
			"aws_secret_access_key": "secret",
			"aws_session_token":     "token",
		},
	}
}

func TestFlattenPrivateRegistries(t *testing.T) {

	cases := []struct {
		Input          []rancher.PrivateRegistry
		ExpectedOutput []interface{}
	}{
		{
			testRKEClusterPrivateRegistriesConf,
			testRKEClusterPrivateRegistriesInterface,
		},
	}

	for _, tc := range cases {
		output := flattenRKEClusterPrivateRegistries(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandPrivateRegistries(t *testing.T) {

	cases := []struct {
		Input          []interface{}
		ExpectedOutput []rancher.PrivateRegistry
	}{
		{
			testRKEClusterPrivateRegistriesInterface,
			testRKEClusterPrivateRegistriesConf,
		},
	}

	for _, tc := range cases {
		output := expandRKEClusterPrivateRegistries(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
