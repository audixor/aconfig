// Copyright (c) 2021-2024 Tenebris Technologies Inc.
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package easyconfig

import (
	"fmt"
	"os"
)

//goland:noinspection GoUnusedExportedFunction
func WithLoad(fileName string) func(*EasyConfig) error {
	return func(c *EasyConfig) error {
		return c.load(fileName)
	}
}

//goland:noinspection GoUnusedExportedFunction
func WithLoadOrCreate(filename string) func(*EasyConfig) error {
	return func(c *EasyConfig) error {
		err := c.load(filename)
		if err != nil {
			c.Loaded = true
			return c.save(filename)
		}
		return nil
	}
}

//goland:noinspection GoUnusedExportedFunction
func WithFind(filenames []string) func(*EasyConfig) error {
	return func(c *EasyConfig) error {
		// Iterate through the possible configuration files
		for _, filename := range filenames {
			if _, err := os.Stat(filename); err == nil {
				return c.load(filename)
			}
		}
		return fmt.Errorf("no configuration file found")
	}
}

//goland:noinspection GoUnusedExportedFunction
func WithFindOrCreate(filenames []string) func(*EasyConfig) error {
	return func(c *EasyConfig) error {
		// Iterate through the possible configuration files
		for _, filename := range filenames {
			if _, err := os.Stat(filename); err == nil {
				return c.load(filename)
			}
		}

		// File was not found, so create it
		// Iterate through the possible configuration files
		for _, filename := range filenames {
			// Attempt to write
			file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
			if err == nil {
				// Success
				_ = file.Close()
				c.Loaded = true
				return c.save(filename)
			}
		}
		return fmt.Errorf("could not create configuration file")
	}
}

//goland:noinspection GoUnusedExportedFunction
func WithWindowsRegistry(key string) func(*EasyConfig) error {
	return func(c *EasyConfig) error {
		c.WindowsRegistry = true
		c.WindowsRegistryKey = key
		return nil
	}
}
