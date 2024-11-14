// Copyright (c) 2021-2024 Tenebris Technologies Inc.
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package aconfig

import (
	"fmt"
	"os"
)

//goland:noinspection GoUnusedExportedFunction
func WithLoad(fileName string) func(*AConfig) error {
	return func(c *AConfig) error {
		return c.load(fileName)
	}
}

//goland:noinspection GoUnusedExportedFunction
func WithLoadOrCreate(filename string) func(*AConfig) error {
	return func(c *AConfig) error {
		err := c.load(filename)
		if err != nil {
			c.Loaded = true
			return c.save(filename)
		}
		return nil
	}
}

//goland:noinspection GoUnusedExportedFunction
func WithFind(filenames []string) func(*AConfig) error {
	return func(c *AConfig) error {
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
func WithFindOrCreate(filenames []string) func(*AConfig) error {
	return func(c *AConfig) error {
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
func WithWindowsRegistry(key string) func(*AConfig) error {
	return func(c *AConfig) error {
		c.WindowsRegistry = true
		c.WindowsRegistryKey = "Software\\" + key
		return nil
	}
}

//goland:noinspection GoUnusedExportedFunction
func WithKeyConstraints(keys []string) func(*AConfig) error {
	return func(c *AConfig) error {
		// Iterate over keys
		for _, key := range keys {
			c.KeyList = append(c.KeyList, key)
		}
		return nil
	}
}
