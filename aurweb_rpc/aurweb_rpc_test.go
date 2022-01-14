package aurweb_rpc

import (
	"net/url"
	"strings"
	"testing"
)

type RequesterStub struct {
	q        string
	r        *Response
	executed bool
}

func (rq *RequesterStub) Set() DataRequester {
	rq.executed = false
	return func(q string, r *Response) {
		rq.executed = true
		rq.q = q
		rq.r = r
	}
}
func (rq *RequesterStub) isCalled() bool { return rq.executed }

var (
	ProperByValues []string = []string{
		"name",
		"name-desc",
		"maintainer",
		"depends",
		"makedepends",
		"optdepends",
		"checkdepends",
	}
)

func TestByValidate(t *testing.T) {
	for _, v := range ProperByValues {
		if !ValidateBy(v) {
			t.Errorf("By value %s isn't validated properly", v)
		}
	}
}

func TestSetDataRequester(t *testing.T) {
	q, rs := "teststring", &Response{}

	rqs := RequesterStub{}
	SetRequester(rqs.Set())

	GetRequester()(q, rs)
	if q != rqs.q {
		t.Error("Improper Set Requester. The string param isn't provided")
	}
	if rs != rqs.r {
		t.Error("Improper Set Requester. The response param isn't provided")
	}
}

func TestGetSearch(t *testing.T) {
	query := "keyword"

	rqs := RequesterStub{}
	SetRequester(rqs.Set())

	_, err := GetSearch(query, "name-desc")

	if !strings.Contains(rqs.q, query) {
		t.Error("The query string isn't not requested properly")
	}
	if err != nil {
		t.Error("GetSearch return error: ", err)
	}
}

func TestGetSearchWrongBy(t *testing.T) {
	query := "keyword"
	wrongby := "blabla"
	_, err := GetSearch(query, wrongby)
	if err == nil {
		t.Error("Wrong value of parameter by doesn't cause error")
	}
}

func TestBuildSearchUrl(t *testing.T) {
	var (
		pattern   string = "keyword"
		search_by string = "name-desc"
	)
	base_url, err := BuildSearchUrl(pattern, search_by)
	// "https://aur.archlinux.org/rpc/?v=5"
	if err != nil {
		t.Error("Error on build URL", err)
	}

	u, err := url.Parse(base_url)
	if err != nil {
		t.Error("Error on Parsing URL", err)
	}

	if u.Host != "aur.archlinux.org" {
		t.Errorf("Wrong host in parsed URL: %s", u.Host)
	}
	if u.Path != "/rpc/" {
		t.Errorf("Wrong path in parsed URL: %s", u.Path)
	}

	q := u.Query()
	if q["v"][0] != "5" {
		t.Errorf("Wrong v value in parsed URL: %s", q["v"])
	}
	if q["type"][0] != "search" {
		t.Errorf("Wrong search value in parsed URL: %s", q["type"])
	}
	if q["by"][0] != search_by {
		t.Errorf("Wrong by value in parsed URL: %s", q["by"])
	}
	if q["arg"][0] != pattern {
		t.Errorf("Wrong arg value in parsed URL: %s", q["arg"])
	}
}

// func TestGetSearchWithBasicRequest(t *testing.T) {
// 	SetRequester(BasicRequest)

// 	r, err := GetSearch("docker", "name-desc")
// 	if err != nil {
// 		t.Error("Fails with: ", err) 
// 	}
// 	println("amount items recieved: ", r.ResultCount)
// }
