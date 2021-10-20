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
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServiceOfferingService_CreateServiceOffering(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiName := "createServiceOffering"
		response, err := ReadData(apiName, "ServiceOfferingService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(w, response[apiName])
	}))
	defer server.Close()

	client := NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	p := client.ServiceOffering.NewCreateServiceOfferingParams("testServiceOffering", "testServiceOffering")
	resp, err := client.ServiceOffering.CreateServiceOffering(p)
	if err != nil {
		t.Errorf("Failed to create service offering due to %v", err)
		return
	}
	if resp == nil || resp.Name != "testServiceOffering" {
		t.Errorf("Failed to create service offering")
	}
}

func TestServiceOfferingService_UpdateServiceOffering(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiName := "updateServiceOffering"
		response, err := ReadData(apiName, "ServiceOfferingService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(w, response[apiName])
	}))
	defer server.Close()

	client := NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	p := client.ServiceOffering.NewUpdateServiceOfferingParams("efaeeab0-4b09-4729-8b6f-62645db41b37")
	p.SetName("testServiceOfferingUpdated")
	resp, err := client.ServiceOffering.UpdateServiceOffering(p)
	if err != nil {
		t.Errorf("Failed to update service offering due to %v", err)
		return
	}

	if resp == nil || resp.Name != "testServiceOfferingUpdated" {
		t.Errorf("Failed to create service offering name")
	}
}

func TestServiceOfferingService_DeleteServiceOffering(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiName := "deleteServiceOffering"
		response, err := ReadData(apiName, "ServiceOfferingService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(w, response[apiName])
	}))
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	p := client.ServiceOffering.NewDeleteServiceOfferingParams("efaeeab0-4b09-4729-8b6f-62645db41b37")
	resp, err := client.ServiceOffering.DeleteServiceOffering(p)
	if err != nil {
		t.Errorf("Failed to delete service offering due to %v", err)
		return
	}
	if resp == nil || !resp.Success {
		t.Errorf("Failed to delete service offering")
	}
}
