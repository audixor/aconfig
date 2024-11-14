// Copyright (c) 2021-2024 Tenebris Technologies Inc.
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package aconfig

import "fmt"

// dump the configuration to stdout
func (c *AConfig) dump() {
	fmt.Printf("Current configuration:\n")
	for n, v := range c.Data {
		fmt.Printf("\t%s = %s\n", n, v)
	}
}
