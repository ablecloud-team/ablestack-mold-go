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
	"net/url"
	"strconv"
)

type QueryAsyncJobResultParams struct {
	p map[string]interface{}
}

func (p *QueryAsyncJobResultParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["jobid"]; found {
		u.Set("jobid", v.(string))
	}
	return u
}

func (p *QueryAsyncJobResultParams) SetJobid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["jobid"] = v
	return
}

// You should always use this function to get a new QueryAsyncJobResultParams instance,
// as then you are sure you have configured all required params
func (s *AsyncjobService) NewQueryAsyncJobResultParams(jobid string) *QueryAsyncJobResultParams {
	p := &QueryAsyncJobResultParams{}
	p.p = make(map[string]interface{})
	p.p["jobid"] = jobid
	return p
}

// Retrieves the current status of asynchronous job.
func (s *AsyncjobService) QueryAsyncJobResult(p *QueryAsyncJobResultParams) (*QueryAsyncJobResultResponse, error) {
	resp, err := s.cs.newRequest("queryAsyncJobResult", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r QueryAsyncJobResultResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type QueryAsyncJobResultResponse struct {
	Jobresulttype   string          `json:"jobresulttype,omitempty"`
	Created         string          `json:"created,omitempty"`
	Accountid       string          `json:"accountid,omitempty"`
	Jobresult       json.RawMessage `json:"jobresult,omitempty"`
	Jobinstancetype string          `json:"jobinstancetype,omitempty"`
	Jobinstanceid   string          `json:"jobinstanceid,omitempty"`
	Jobstatus       int             `json:"jobstatus,omitempty"`
	Cmd             string          `json:"cmd,omitempty"`
	Jobresultcode   int             `json:"jobresultcode,omitempty"`
	Jobprocstatus   int             `json:"jobprocstatus,omitempty"`
	Userid          string          `json:"userid,omitempty"`
}

type ListAsyncJobsParams struct {
	p map[string]interface{}
}

func (p *ListAsyncJobsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	if v, found := p.p["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	return u
}

func (p *ListAsyncJobsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListAsyncJobsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListAsyncJobsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListAsyncJobsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListAsyncJobsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListAsyncJobsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListAsyncJobsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListAsyncJobsParams) SetStartdate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startdate"] = v
	return
}

// You should always use this function to get a new ListAsyncJobsParams instance,
// as then you are sure you have configured all required params
func (s *AsyncjobService) NewListAsyncJobsParams() *ListAsyncJobsParams {
	p := &ListAsyncJobsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists all pending asynchronous jobs for the account.
func (s *AsyncjobService) ListAsyncJobs(p *ListAsyncJobsParams) (*ListAsyncJobsResponse, error) {
	resp, err := s.cs.newRequest("listAsyncJobs", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListAsyncJobsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListAsyncJobsResponse struct {
	Count     int         `json:"count"`
	AsyncJobs []*AsyncJob `json:"asyncjob"`
}

type AsyncJob struct {
	Accountid       string          `json:"accountid,omitempty"`
	Created         string          `json:"created,omitempty"`
	Jobstatus       int             `json:"jobstatus,omitempty"`
	Jobinstanceid   string          `json:"jobinstanceid,omitempty"`
	Jobprocstatus   int             `json:"jobprocstatus,omitempty"`
	Jobresult       json.RawMessage `json:"jobresult,omitempty"`
	Jobresulttype   string          `json:"jobresulttype,omitempty"`
	Jobresultcode   int             `json:"jobresultcode,omitempty"`
	Jobinstancetype string          `json:"jobinstancetype,omitempty"`
	Cmd             string          `json:"cmd,omitempty"`
	Userid          string          `json:"userid,omitempty"`
}
