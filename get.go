// Copyright (c) 2021-2024 Tenebris Technologies Inc.
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package easyconfig

import (
	"errors"
	"strconv"
	"strings"
)

// GetStr retrieves a string value or returns an error
func (c *EasyConfig) GetStr(key string) (string, error) {
	value, err := c.get(key)
	if err != nil {
		return "", err
	}
	return value, nil
}

// GetStrDefault retrieves a string value or returns the specified default
func (c *EasyConfig) GetStrDefault(key string, defaultValue string) string {
	value, err := c.GetStr(key)
	if err != nil {
		return defaultValue
	}
	return value
}

// GetInt retrieves an integer value or returns an error
func (c *EasyConfig) GetInt(key string) (int, error) {
	value, err := c.get(key)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.New("error converting value to integer")
	}
	return i, nil
}

// GetIntDefault retrieves an integer value or returns the specified default
func (c *EasyConfig) GetIntDefault(key string, defaultValue int) int {
	value, err := c.GetInt(key)
	if err != nil {
		return defaultValue
	}
	return value
}

// GetInt64 retrieves an int64 value or returns an error
func (c *EasyConfig) GetInt64(key string) (int64, error) {
	value, err := c.get(key)
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, errors.New("error converting value to integer")
	}

	return i, nil
}

// GetInt64Default retrieves an int64 value or returns the specified default
func (c *EasyConfig) GetInt64Default(key string, defaultValue int64) int64 {
	value, err := c.GetInt64(key)
	if err != nil {
		return defaultValue
	}
	return value
}

// GetBool retrieves a boolean value or returns an error
func (c *EasyConfig) GetBool(key string) (bool, error) {
	value, err := c.get(key)
	if err != nil {
		return false, err
	}

	if strings.EqualFold(value, "true") {
		return true, nil
	}

	if strings.EqualFold(value, "yes") {
		return true, nil
	}

	if strings.EqualFold(value, "1") {
		return true, nil
	}

	if strings.EqualFold(value, "false") {
		return false, nil
	}

	if strings.EqualFold(value, "no") {
		return false, nil
	}

	if strings.EqualFold(value, "0") {
		return false, nil
	}

	return false, errors.New("error converting value to boolean")
}

// GetBoolDefault retrieves a boolean value or the specified default
func (c *EasyConfig) GetBoolDefault(key string, defaultValue bool) bool {
	value, err := c.GetBool(key)
	if err != nil {
		return defaultValue
	}
	return value
}
