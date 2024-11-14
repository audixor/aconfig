// Copyright (c) 2021-2024 Tenebris Technologies Inc.
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package aconfig

import "fmt"

// AConfig holds all configuration data
type AConfig struct {
	Loaded             bool              // Configuration loaded
	WindowsRegistry    bool              // Use the Windows registry (ignored on non-Windows systems)
	WindowsRegistryKey string            // Windows registry key (ignored on non-Windows systems)
	ConfigFile         string            // Path to configuration file
	KeyList            []string          // List of keys
	Data               map[string]string // Configuration data
}

// New returns an AConfig instance
//
//goland:noinspection GoUnusedExportedFunction
func New(options ...func(*AConfig) error) (*AConfig, error) {
	c := &AConfig{
		Loaded:             false,
		WindowsRegistry:    false,
		WindowsRegistryKey: "",
		ConfigFile:         "",
		Data:               make(map[string]string),
	}

	// Process options (see options.go)
	for _, op := range options {
		err := op(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

// Init initializes the configuration data
func (c *AConfig) Init() {
	// Delete everything in the existing config
	for k := range c.Data {
		delete(c.Data, k)
	}
	c.Loaded = false
}

// Save the configuration to the specified file
func (c *AConfig) Save(filename string) error {
	if filename == "" {
		if c.ConfigFile != "" {
			filename = c.ConfigFile
		} else {
			return fmt.Errorf("no filename specified")
		}
	}
	return c.save(filename)
}

// Checkpoint saves the configuration to the last loaded file
func (c *AConfig) Checkpoint() error {
	if c.ConfigFile == "" {
		return fmt.Errorf("checkpoint requires a loaded configuration")
	}
	return c.save(c.ConfigFile)
}

// Dump configuration to stdout for diagnostics
func (c *AConfig) Dump() {
	c.dump()
}

// Exists checks if the specified configuration key exits
func (c *AConfig) Exists(key string) bool {
	_, ok := c.Data[key]
	return ok
}
