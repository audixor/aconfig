// Copyright (c) 2021-2024 Tenebris Technologies Inc.
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package easyconfig

import (
	"encoding/json"
	"fmt"
	"os"
)

// Load configuration from specified file
func (c *EasyConfig) load(filename string) error {

	// Loading not allowed when using the Windows registry
	if c.WindowsRegistry {
		return nil
	}

	// Have we previously loaded a config file?
	if c.Loaded {
		c.Init()
	}

	// Open the config file
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file %s: %v\n", filename, err)
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	// Decode the JSON data into the struct
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c.Data)
	if err != nil {
		return fmt.Errorf("error decoding JSON: %v\n", err)
	}

	// Turn on configured flag and save filename
	c.Loaded = true
	c.ConfigFile = filename
	return nil
}
