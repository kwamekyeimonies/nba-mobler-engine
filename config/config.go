package config

import "log"

var WithConfigInjector *ConfigurationInjector

type ConfigurationInjector struct {
	KeystoreConfig *Keystore
}

func NewConfigurationInjector() *ConfigurationInjector {
	configKeys := &ConfigurationInjector{}

	if configKeys.KeystoreConfig = NewKeystoreConfig(); configKeys.KeystoreConfig == nil {
		log.Fatal("failed to load credentials")
		return nil
	}

	WithConfigInjector = configKeys

	return configKeys
}
