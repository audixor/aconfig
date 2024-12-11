//
// Copyright (c) 2024 Tenebris Technologies Inc.
// Please see the LICENSE file for details
//

// Code for operating systems other than windows
//go:build !windows

package aconfig

import (
	"fmt"
	"strings"
)

// For *nix systems, simply read and store from the configuration structure

func (c *AConfig) set(key string, value string) error {
	// Check for key constraints
	if !c.checkKey(key) {
		return fmt.Errorf("invalid key: %s", key)
	}

	c.Data[strings.ToLower(key)] = value
	return nil
}

func (c *AConfig) get(key string) (string, error) {
	// Check for key constraints
	if !c.checkKey(key) {
		return "", fmt.Errorf("invalid key: %s", key)
	}

	value, ok := c.Data[strings.ToLower(key)]
	if ok == false {
		return "", nil
	}
	return value, nil
}
