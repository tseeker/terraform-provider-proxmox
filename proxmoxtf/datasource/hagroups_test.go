/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package datasource

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/bpg/terraform-provider-proxmox/proxmoxtf/test"
)

// TestHAGroupsInstantiation tests whether HAGroups can be instantiated.
func TestHAGroupsInstantiation(t *testing.T) {
	t.Parallel()

	s := HAGroups()
	if s == nil {
		t.Fatalf("Cannot instantiate HAGroups")
	}
}

// TestHAGroupsSchema tests the HAGroups schema.
func TestHAGroupsSchema(t *testing.T) {
	t.Parallel()

	s := HAGroups()

	test.AssertComputedAttributes(t, s, []string{
		mkDataSourceVirtualEnvironmentHAGroupsGroupIDs,
	})

	test.AssertValueTypes(t, s, map[string]schema.ValueType{
		mkDataSourceVirtualEnvironmentHAGroupsGroupIDs: schema.TypeList,
	})
}
