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

package cloudstack

import (
	"encoding/json"
	"testing"
)

func TestMoldIPAllocationReceiptAndReleasePrecondition(t *testing.T) {
	var ip PublicIpAddress
	if err := json.Unmarshal([]byte(`{"id":"ip-1","allocated":"2026-10-06T19:00:00+0000","allocationgeneration":"generation-1"}`), &ip); err != nil {
		t.Fatal(err)
	}
	if ip.Allocationgeneration != "generation-1" {
		t.Fatal("allocation receipt lost")
	}
	p := (&AddressService{}).NewDisassociateIpAddressParams(ip.Id)
	p.SetExpectedallocationgeneration(ip.Allocationgeneration)
	if p.toURLValues().Get("expectedallocationgeneration") != "generation-1" {
		t.Fatal("release precondition omitted")
	}
	p.ResetExpectedallocationgeneration()
	if _, found := p.GetExpectedallocationgeneration(); found || p.toURLValues().Has("expectedallocationgeneration") {
		t.Fatal("release precondition reset failed")
	}
}
