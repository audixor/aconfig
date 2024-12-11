//
// Copyright (c) 2024 Tenebris Technologies Inc.
// Please see the LICENSE file for details
//

package aconfig

func (c *AConfig) checkKey(key string) bool {
	// If no key constraints, any are valid
	if len(c.KeyList) == 0 {
		return true
	}

	// Iterate over allowed keys
	for _, k := range c.KeyList {
		if k == key {
			return true
		}
	}
	return false
}
