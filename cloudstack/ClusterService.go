//
// Copyright 2014, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cloudstack

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type AddClusterParams struct {
	p map[string]interface{}
}

func (p *AddClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["clustername"]; found {
		u.Set("clustername", v.(string))
	}
	if v, found := p.p["clustertype"]; found {
		u.Set("clustertype", v.(string))
	}
	if v, found := p.p["guestvswitchname"]; found {
		u.Set("guestvswitchname", v.(string))
	}
	if v, found := p.p["guestvswitchtype"]; found {
		u.Set("guestvswitchtype", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["publicvswitchname"]; found {
		u.Set("publicvswitchname", v.(string))
	}
	if v, found := p.p["publicvswitchtype"]; found {
		u.Set("publicvswitchtype", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	if v, found := p.p["vsmipaddress"]; found {
		u.Set("vsmipaddress", v.(string))
	}
	if v, found := p.p["vsmpassword"]; found {
		u.Set("vsmpassword", v.(string))
	}
	if v, found := p.p["vsmusername"]; found {
		u.Set("vsmusername", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AddClusterParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
	return
}

func (p *AddClusterParams) SetClustername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustername"] = v
	return
}

func (p *AddClusterParams) SetClustertype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustertype"] = v
	return
}

func (p *AddClusterParams) SetGuestvswitchname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["guestvswitchname"] = v
	return
}

func (p *AddClusterParams) SetGuestvswitchtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["guestvswitchtype"] = v
	return
}

func (p *AddClusterParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
	return
}

func (p *AddClusterParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
	return
}

func (p *AddClusterParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
	return
}

func (p *AddClusterParams) SetPublicvswitchname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicvswitchname"] = v
	return
}

func (p *AddClusterParams) SetPublicvswitchtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicvswitchtype"] = v
	return
}

func (p *AddClusterParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
	return
}

func (p *AddClusterParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

func (p *AddClusterParams) SetVsmipaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vsmipaddress"] = v
	return
}

func (p *AddClusterParams) SetVsmpassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vsmpassword"] = v
	return
}

func (p *AddClusterParams) SetVsmusername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vsmusername"] = v
	return
}

func (p *AddClusterParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new AddClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewAddClusterParams(clustername string, clustertype string, hypervisor string, podid string, zoneid string) *AddClusterParams {
	p := &AddClusterParams{}
	p.p = make(map[string]interface{})
	p.p["clustername"] = clustername
	p.p["clustertype"] = clustertype
	p.p["hypervisor"] = hypervisor
	p.p["podid"] = podid
	p.p["zoneid"] = zoneid
	return p
}

// Adds a new cluster
func (s *ClusterService) AddCluster(p *AddClusterParams) (*AddClusterResponse, error) {
	resp, err := s.cs.newRequest("addCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type AddClusterResponse struct {
	Podid           string `json:"podid,omitempty"`
	Clustertype     string `json:"clustertype,omitempty"`
	Zonename        string `json:"zonename,omitempty"`
	Managedstate    string `json:"managedstate,omitempty"`
	Allocationstate string `json:"allocationstate,omitempty"`
	Capacity        []struct {
		Podname       string `json:"podname,omitempty"`
		Type          int    `json:"type,omitempty"`
		Zoneid        string `json:"zoneid,omitempty"`
		Capacityused  int    `json:"capacityused,omitempty"`
		Percentused   string `json:"percentused,omitempty"`
		Capacitytotal int    `json:"capacitytotal,omitempty"`
		Clusterid     string `json:"clusterid,omitempty"`
		Clustername   string `json:"clustername,omitempty"`
		Podid         string `json:"podid,omitempty"`
		Zonename      string `json:"zonename,omitempty"`
	} `json:"capacity,omitempty"`
	Zoneid                string `json:"zoneid,omitempty"`
	Cpuovercommitratio    string `json:"cpuovercommitratio,omitempty"`
	Id                    string `json:"id,omitempty"`
	Podname               string `json:"podname,omitempty"`
	Name                  string `json:"name,omitempty"`
	Memoryovercommitratio string `json:"memoryovercommitratio,omitempty"`
	Hypervisortype        string `json:"hypervisortype,omitempty"`
}

type DeleteClusterParams struct {
	p map[string]interface{}
}

func (p *DeleteClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteClusterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewDeleteClusterParams(id string) *DeleteClusterParams {
	p := &DeleteClusterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a cluster.
func (s *ClusterService) DeleteCluster(p *DeleteClusterParams) (*DeleteClusterResponse, error) {
	resp, err := s.cs.newRequest("deleteCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type DeleteClusterResponse struct {
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type UpdateClusterParams struct {
	p map[string]interface{}
}

func (p *UpdateClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["clustername"]; found {
		u.Set("clustername", v.(string))
	}
	if v, found := p.p["clustertype"]; found {
		u.Set("clustertype", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["managedstate"]; found {
		u.Set("managedstate", v.(string))
	}
	return u
}

func (p *UpdateClusterParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
	return
}

func (p *UpdateClusterParams) SetClustername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustername"] = v
	return
}

func (p *UpdateClusterParams) SetClustertype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustertype"] = v
	return
}

func (p *UpdateClusterParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
	return
}

func (p *UpdateClusterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateClusterParams) SetManagedstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["managedstate"] = v
	return
}

// You should always use this function to get a new UpdateClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewUpdateClusterParams(id string) *UpdateClusterParams {
	p := &UpdateClusterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates an existing cluster
func (s *ClusterService) UpdateCluster(p *UpdateClusterParams) (*UpdateClusterResponse, error) {
	resp, err := s.cs.newRequest("updateCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type UpdateClusterResponse struct {
	Managedstate          string `json:"managedstate,omitempty"`
	Name                  string `json:"name,omitempty"`
	Clustertype           string `json:"clustertype,omitempty"`
	Cpuovercommitratio    string `json:"cpuovercommitratio,omitempty"`
	Podid                 string `json:"podid,omitempty"`
	Zonename              string `json:"zonename,omitempty"`
	Hypervisortype        string `json:"hypervisortype,omitempty"`
	Allocationstate       string `json:"allocationstate,omitempty"`
	Id                    string `json:"id,omitempty"`
	Memoryovercommitratio string `json:"memoryovercommitratio,omitempty"`
	Capacity              []struct {
		Percentused   string `json:"percentused,omitempty"`
		Podid         string `json:"podid,omitempty"`
		Clustername   string `json:"clustername,omitempty"`
		Podname       string `json:"podname,omitempty"`
		Zonename      string `json:"zonename,omitempty"`
		Type          int    `json:"type,omitempty"`
		Zoneid        string `json:"zoneid,omitempty"`
		Capacityused  int    `json:"capacityused,omitempty"`
		Clusterid     string `json:"clusterid,omitempty"`
		Capacitytotal int    `json:"capacitytotal,omitempty"`
	} `json:"capacity,omitempty"`
	Podname string `json:"podname,omitempty"`
	Zoneid  string `json:"zoneid,omitempty"`
}

type ListClustersParams struct {
	p map[string]interface{}
}

func (p *ListClustersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["clustertype"]; found {
		u.Set("clustertype", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["managedstate"]; found {
		u.Set("managedstate", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["showcapacities"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showcapacities", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListClustersParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
	return
}

func (p *ListClustersParams) SetClustertype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustertype"] = v
	return
}

func (p *ListClustersParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
	return
}

func (p *ListClustersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListClustersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListClustersParams) SetManagedstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["managedstate"] = v
	return
}

func (p *ListClustersParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListClustersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListClustersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListClustersParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
	return
}

func (p *ListClustersParams) SetShowcapacities(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showcapacities"] = v
	return
}

func (p *ListClustersParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListClustersParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewListClustersParams() *ListClustersParams {
	p := &ListClustersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ClusterService) GetClusterID(name string) (string, error) {
	p := &ListClustersParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListClusters(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, name, l)
	}
	return l.Clusters[0].Id, nil
}

// Lists clusters.
func (s *ClusterService) ListClusters(p *ListClustersParams) (*ListClustersResponse, error) {
	resp, err := s.cs.newRequest("listClusters", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListClustersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListClustersResponse struct {
	Count    int        `json:"count"`
	Clusters []*Cluster `json:"cluster"`
}

type Cluster struct {
	Managedstate          string `json:"managedstate,omitempty"`
	Podname               string `json:"podname,omitempty"`
	Allocationstate       string `json:"allocationstate,omitempty"`
	Memoryovercommitratio string `json:"memoryovercommitratio,omitempty"`
	Clustertype           string `json:"clustertype,omitempty"`
	Id                    string `json:"id,omitempty"`
	Cpuovercommitratio    string `json:"cpuovercommitratio,omitempty"`
	Capacity              []struct {
		Capacityused  int    `json:"capacityused,omitempty"`
		Podid         string `json:"podid,omitempty"`
		Percentused   string `json:"percentused,omitempty"`
		Capacitytotal int    `json:"capacitytotal,omitempty"`
		Podname       string `json:"podname,omitempty"`
		Clusterid     string `json:"clusterid,omitempty"`
		Type          int    `json:"type,omitempty"`
		Clustername   string `json:"clustername,omitempty"`
		Zonename      string `json:"zonename,omitempty"`
		Zoneid        string `json:"zoneid,omitempty"`
	} `json:"capacity,omitempty"`
	Zoneid         string `json:"zoneid,omitempty"`
	Name           string `json:"name,omitempty"`
	Zonename       string `json:"zonename,omitempty"`
	Podid          string `json:"podid,omitempty"`
	Hypervisortype string `json:"hypervisortype,omitempty"`
}

type DedicateClusterParams struct {
	p map[string]interface{}
}

func (p *DedicateClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	return u
}

func (p *DedicateClusterParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *DedicateClusterParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
	return
}

func (p *DedicateClusterParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

// You should always use this function to get a new DedicateClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewDedicateClusterParams(clusterid string, domainid string) *DedicateClusterParams {
	p := &DedicateClusterParams{}
	p.p = make(map[string]interface{})
	p.p["clusterid"] = clusterid
	p.p["domainid"] = domainid
	return p
}

// Dedicate an existing cluster
func (s *ClusterService) DedicateCluster(p *DedicateClusterParams) (*DedicateClusterResponse, error) {
	resp, err := s.cs.newRequest("dedicateCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DedicateClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
		}

		var r DedicateClusterResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DedicateClusterResponse struct {
	JobID           string `json:"jobid,omitempty"`
	Id              string `json:"id,omitempty"`
	Domainid        string `json:"domainid,omitempty"`
	Clustername     string `json:"clustername,omitempty"`
	Accountid       string `json:"accountid,omitempty"`
	Affinitygroupid string `json:"affinitygroupid,omitempty"`
	Clusterid       string `json:"clusterid,omitempty"`
}

type ReleaseDedicatedClusterParams struct {
	p map[string]interface{}
}

func (p *ReleaseDedicatedClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	return u
}

func (p *ReleaseDedicatedClusterParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
	return
}

// You should always use this function to get a new ReleaseDedicatedClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewReleaseDedicatedClusterParams(clusterid string) *ReleaseDedicatedClusterParams {
	p := &ReleaseDedicatedClusterParams{}
	p.p = make(map[string]interface{})
	p.p["clusterid"] = clusterid
	return p
}

// Release the dedication for cluster
func (s *ClusterService) ReleaseDedicatedCluster(p *ReleaseDedicatedClusterParams) (*ReleaseDedicatedClusterResponse, error) {
	resp, err := s.cs.newRequest("releaseDedicatedCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleaseDedicatedClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
		}

		var r ReleaseDedicatedClusterResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type ReleaseDedicatedClusterResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type ListDedicatedClustersParams struct {
	p map[string]interface{}
}

func (p *ListDedicatedClustersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["affinitygroupid"]; found {
		u.Set("affinitygroupid", v.(string))
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListDedicatedClustersParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListDedicatedClustersParams) SetAffinitygroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupid"] = v
	return
}

func (p *ListDedicatedClustersParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
	return
}

func (p *ListDedicatedClustersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListDedicatedClustersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListDedicatedClustersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListDedicatedClustersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListDedicatedClustersParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewListDedicatedClustersParams() *ListDedicatedClustersParams {
	p := &ListDedicatedClustersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ClusterService) GetDedicatedClusterID(keyword string) (string, error) {
	p := &ListDedicatedClustersParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	l, err := s.ListDedicatedClusters(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.DedicatedClusters[0].Id, nil
}

// Lists dedicated clusters.
func (s *ClusterService) ListDedicatedClusters(p *ListDedicatedClustersParams) (*ListDedicatedClustersResponse, error) {
	resp, err := s.cs.newRequest("listDedicatedClusters", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDedicatedClustersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListDedicatedClustersResponse struct {
	Count             int                 `json:"count"`
	DedicatedClusters []*DedicatedCluster `json:"dedicatedcluster"`
}

type DedicatedCluster struct {
	Domainid        string `json:"domainid,omitempty"`
	Id              string `json:"id,omitempty"`
	Clusterid       string `json:"clusterid,omitempty"`
	Clustername     string `json:"clustername,omitempty"`
	Accountid       string `json:"accountid,omitempty"`
	Affinitygroupid string `json:"affinitygroupid,omitempty"`
}
