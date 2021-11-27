package config

import "github.com/maiaaraujo5/gostart/config"

const root = "gostart.app"

func versionValue() string {
	return config.GetStringValue(root + ".version")
}

func nameValue() string {
	return config.GetStringValue(root + ".name")
}

func DefaultAppFields() map[string]interface{} {
	fields := make(map[string]interface{})
	if versionValue() != "" {
		fields["version"] = versionValue()
	}

	if nameValue() != "" {
		fields["name"] = nameValue()
	}

	return fields
}
