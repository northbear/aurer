package aurweb_rpc

import (
	"testing"
	"net/url"
)

var ProperByValues []string = []string{
	"name",
	"name-desc",
	"maintainer",
	"depends",
	"makedepends",
	"optdepends",
	"checkdepends",
}

func TestByValidate(t *testing.T) {
	for _, v := range ProperByValues {
		if !ValidateBy(v) {
			t.Errorf("By value %s isn't validated properly", v)
		}
	}
}

func TestGetSearch(t *testing.T) {
	query := "keyword"
	_, err := GetSearch(query, "name-desc")
	if err != nil {
		t.Error("GetSearch return error: ", err) 
	}
}

func TestGetSearchWrongBy(t *testing.T) {
	query := "keyword"
	wrongby := "blabla"
	_, err := GetSearch(query, wrongby)
	if err == nil {
		t.Error("Wrong value of parameter by doesn't cause Error")
	}
}

func TestBuildSearchUrl(t *testing.T) {
	var (
		pattern string = "keyword"
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
