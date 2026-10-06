// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied. See the License for the
// specific language governing permissions and limitations
// under the License.
package cloudstack

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestMoldAPIResponseFailures(t *testing.T) {
	cases := []struct {
		name, body, wantError string
		status                int
	}{
		{"HTTP200APIError", `{"errorresponse":{"errorcode":530,"errortext":"backend error"}}`, "backend error", 200},
		{"HTTP401", `{"errorresponse":{"errorcode":401,"errortext":"denied"}}`, "denied", 401},
		{"HTTP503", `{"errorresponse":{"errorcode":503,"errortext":"unavailable"}}`, "unavailable", 503},
		{"CommandError", `{"listvirtualmachinesresponse":{"errorcode":530,"errortext":"backend error"}}`, "backend error", 200},
		{"NonJSON", "<html>unavailable</html>", "invalid JSON", 503},
		{"NullEnvelope", "null", "empty response envelope", 200},
		{"EmptyEnvelope", `{}`, "empty response envelope", 200},
		{"WrongCommand", `{"listvolumesresponse":{"count":0}}`, "unexpected response envelope", 200},
		{"AmbiguousEnvelope", `{"listvirtualmachinesresponse":{},"extraresponse":{}}`, "unexpected response envelope", 200},
		{"NullPayload", `{"listvirtualmachinesresponse":null}`, "payload must be an object", 200},
		{"ArrayPayload", `{"listvirtualmachinesresponse":[]}`, "payload must be an object", 200},
		{"EmptyListCount", `{"listvirtualmachinesresponse":{"count":0}}`, "", 200},
		{"EmptyListOmittedCount", `{"listvirtualmachinesresponse":{}}`, "", 200},
		{"ValidVM", `{"listvirtualmachinesresponse":{"count":1,"virtualmachine":[{"id":"fixture-vm","state":"Running"}]}}`, "", 200},
	}
	for _, post := range []bool{false, true} {
		for _, tc := range cases {
			method := map[bool]string{false: "GET", true: "POST"}[post]
			t.Run(method+"/"+tc.name, func(t *testing.T) {
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if r.Method != method {
						t.Errorf("method=%s", r.Method)
					}
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(tc.status)
					_, _ = w.Write([]byte(tc.body))
				}))
				defer server.Close()
				client := newClient(server.URL, "fixture-key", "fixture-secret", false, true)
				_, err := client.newRawRequest("listVirtualMachines", post, url.Values{})
				if tc.wantError == "" {
					if err != nil {
						t.Fatal(err)
					}
					return
				}
				if err == nil || !strings.Contains(err.Error(), tc.wantError) {
					t.Fatalf("err=%v, want %q", err, tc.wantError)
				}
			})
		}
	}
}

func TestLegacyAsyncResponseWrapper(t *testing.T) {
	_, err := decodeAPIResponse("deleteVirtualMachine", http.StatusOK, []byte(`{"jobresult":{"success":true}}`))
	if err != nil {
		t.Fatal(err)
	}
}
