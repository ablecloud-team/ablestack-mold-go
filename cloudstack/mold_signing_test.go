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
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

// These vectors use Java URLEncoder's value-only encoding, independently
// generated from Mold ApiServer's contract. The keys and credentials are dummy.
func TestMoldSigningVector(t *testing.T) {
	var v struct {
		Secret, Canonical, Signature string
		Params                       map[string]string
	}
	b, err := os.ReadFile("testdata/mold-signing.json")
	if err != nil {
		t.Fatal(err)
	}
	if err = json.Unmarshal(b, &v); err != nil {
		t.Fatal(err)
	}
	p := url.Values{}
	for k, x := range v.Params {
		p.Set(k, x)
	}
	canonical := strings.ToLower(EncodeValues(p))
	if canonical != v.Canonical {
		t.Fatalf("canonical = %q, want %q", canonical, v.Canonical)
	}
	mac := hmac.New(sha256.New, []byte(v.Secret))
	mac.Write([]byte(canonical))
	if got := base64.StdEncoding.EncodeToString(mac.Sum(nil)); got != v.Signature {
		t.Fatalf("signature = %s", got)
	}
}

func TestMoldSignedGETAndPOST(t *testing.T) {
	for _, post := range []bool{false, true} {
		t.Run(map[bool]string{false: "GET", true: "POST"}[post], func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if err := r.ParseForm(); err != nil {
					t.Error(err)
					w.WriteHeader(400)
					return
				}
				if r.Method != map[bool]string{false: "GET", true: "POST"}[post] {
					t.Errorf("method %s", r.Method)
				}
				p := r.Form
				sig := p.Get("signature")
				p.Del("signature")
				if p.Get("name") != "a b+c%*~한글" || p.Get("signatureversion") != "3" || p.Get("expires") == "" {
					t.Error("parameters or expiration changed in transport")
				}
				// Use Java's exact safe alphabet rather than the client's encoder.
				keys := []string{"apiKey", "command", "expires", "name", "response", "signatureversion"}
				parts := []string{}
				for _, k := range keys {
					var out strings.Builder
					for _, b := range []byte(p.Get(k)) {
						if strings.ContainsRune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.-_*", rune(b)) {
							out.WriteByte(b)
						} else {
							const hex = "0123456789ABCDEF"
							out.WriteByte('%')
							out.WriteByte(hex[b>>4])
							out.WriteByte(hex[b&15])
						}
					}
					parts = append(parts, k+"="+out.String())
				}
				mac := hmac.New(sha256.New, []byte("fixture-secret"))
				mac.Write([]byte(strings.ToLower(strings.Join(parts, "&"))))
				if !hmac.Equal([]byte(sig), []byte(base64.StdEncoding.EncodeToString(mac.Sum(nil)))) {
					t.Error("Mold rejected signature")
					w.WriteHeader(401)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"listkubernetesclustersresponse":{"count":0}}`))
			}))
			defer server.Close()
			cs := newClient(server.URL, "Mold+Key ~*한글", "fixture-secret", false, true)
			if _, err := cs.newRawRequest("listKubernetesClusters", post, url.Values{"name": {"a b+c%*~한글"}}); err != nil {
				t.Fatal(err)
			}
		})
	}
}
