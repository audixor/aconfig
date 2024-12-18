//
// Copyright (c) 2024 Tenebris Technologies Inc.
// Please see the LICENSE file for details
//

// Code Windows Only
//go:build windows

package aconfig

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"strings"
)

// For *nix systems, simply read and store from the configuration structure

func (c *AConfig) set(key string, value string) error {
	// Check for key constraints
	if !c.checkKey(key) {
		return fmt.Errorf("invalid key: %s", key)
	}

	if c.WindowsRegistryKey == "" {
		return fmt.Errorf("windows registry key not set")
	}

	// Create path
	rPath := "Software\\" + c.WindowsRegistryKey

	// Open the registry key
	rkey, _, err := registry.CreateKey(registry.CURRENT_USER, rPath, registry.ALL_ACCESS)
	if err != nil {
		return fmt.Errorf("failed to open registry key %s: %v", rPath, err)
	}

	// Defer closing the key
	defer func(rkey registry.Key) {
		_ = rkey.Close()
	}(rkey)

	err = rkey.SetStringValue(strings.ToLower(key), value)
	if err != nil {
		return fmt.Errorf("failed to set %s registry value: %v", strings.ToLower(key), err)
	}
	return nil
}

func (c *AConfig) get(key string) (string, error) {
	// Check for key constraints
	if !c.checkKey(key) {
		return "", fmt.Errorf("invalid key: %s", key)
	}

	if c.WindowsRegistryKey == "" {
		return "", fmt.Errorf("windows registry key not set")
	}

	// Create path
	rPath := "Software\\" + c.WindowsRegistryKey

	// Open the registry key
	rkey, _, err := registry.CreateKey(registry.CURRENT_USER, rPath, registry.ALL_ACCESS)
	if err != nil {
		return "", fmt.Errorf("failed to open registry key %s: %v", rPath, err)
	}

	// Defer closing the key
	defer func(rkey registry.Key) {
		_ = rkey.Close()
	}(rkey)

	// Get the key
	value, _, err := rkey.GetStringValue(strings.ToLower(key))
	if err != nil {
		return "", nil
	}
	return value, nil
}
