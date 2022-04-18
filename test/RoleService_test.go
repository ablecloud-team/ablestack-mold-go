//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package test

import (
	"testing"

	"github.com/apache/cloudstack-go/v2/cloudstack"
)

func TestRoleService(t *testing.T) {
	service := "RoleService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testcreateRole := func(t *testing.T) {
		if _, ok := response["createRole"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Role.NewCreateRoleParams("name")
		_, err := client.Role.CreateRole(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateRole", testcreateRole)

	testcreateRolePermission := func(t *testing.T) {
		if _, ok := response["createRolePermission"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Role.NewCreateRolePermissionParams("permission", "roleid", "rule")
		_, err := client.Role.CreateRolePermission(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateRolePermission", testcreateRolePermission)

	testdeleteRole := func(t *testing.T) {
		if _, ok := response["deleteRole"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Role.NewDeleteRoleParams("id")
		_, err := client.Role.DeleteRole(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteRole", testdeleteRole)

	testdeleteRolePermission := func(t *testing.T) {
		if _, ok := response["deleteRolePermission"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Role.NewDeleteRolePermissionParams("id")
		_, err := client.Role.DeleteRolePermission(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteRolePermission", testdeleteRolePermission)

	testlistRolePermissions := func(t *testing.T) {
		if _, ok := response["listRolePermissions"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Role.NewListRolePermissionsParams()
		_, err := client.Role.ListRolePermissions(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListRolePermissions", testlistRolePermissions)

	testlistRoles := func(t *testing.T) {
		if _, ok := response["listRoles"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Role.NewListRolesParams()
		_, err := client.Role.ListRoles(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListRoles", testlistRoles)

	testupdateRole := func(t *testing.T) {
		if _, ok := response["updateRole"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Role.NewUpdateRoleParams("id")
		_, err := client.Role.UpdateRole(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpdateRole", testupdateRole)

	testupdateRolePermission := func(t *testing.T) {
		if _, ok := response["updateRolePermission"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Role.NewUpdateRolePermissionParams("roleid")
		_, err := client.Role.UpdateRolePermission(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpdateRolePermission", testupdateRolePermission)

}
