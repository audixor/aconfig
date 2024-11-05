// Copyright (c) 2021-2024 Tenebris Technologies Inc.
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

// Code for operating systems other than windows
//go:build !windows

package easyconfig

import (
	"fmt"
	"strings"
)

// For *nix systems, simply read and store from the configuration structure

func (c *EasyConfig) set(key string, value string) error {
	c.Data[strings.ToLower(key)] = value
	return nil
}

func (c *EasyConfig) get(key string) (string, error) {
	value, ok := c.Data[strings.ToLower(key)]
	if ok == false {
		return "", fmt.Errorf("key %s does not exist", key)
	}
	return value, nil
}
