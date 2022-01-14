package aurweb_rpc

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

/*
Documentation:
    https://aur.archlinux.org/rpc/
    https://wiki.archlinux.org/title/Aurweb_RPC_interface

For the ReturnType search, ReturnData may contain the following fields:

    ID
    Name
    PackageBaseID
    PackageBase
    Version
    Description
    URL
    NumVotes
    Popularity
    OutOfDate
    Maintainer
    FirstSubmitted
    LastModified
    URLPath

    Depends
    MakeDepends
    OptDepends
    CheckDepends
    Conflicts
    Provides
    Replaces
    Groups
    License
    Keywords
*/

/*
{
    "version":5,
    "type":"multiinfo",
    "resultcount":1,
    "results":[{
        "ID":229417,
        "Name":"cower",
        "PackageBaseID":44921,
        "PackageBase":"cower",
        "Version":"14-2",
        "Description":"A simple AUR agent with a pretentious name",
        "URL":"http:\/\/github.com\/falconindy\/cower",
        "NumVotes":590,
        "Popularity":24.595536,
        "OutOfDate":null,
        "Maintainer":"falconindy",
        "FirstSubmitted":1293676237,
        "LastModified":1441804093,
        "URLPath":"\/cgit\/aur.git\/snapshot\/cower.tar.gz",
        "Depends":[
            "curl",
            "openssl",
            "pacman",
            "yajl"
        ],
        "MakeDepends":[
            "perl"
        ],
        "License":[
            "MIT"
        ],
        "Keywords":[]
    }]
 }

Response

{"version":5,"type":ReturnType,"resultcount":0,"results":ReturnData}
	// search_query := "https://aur.archlinux.org/rpc/?v=5&type=search&arg=%s"

Search By:
    name (search by package name only)
    name-desc (search by package name and description)
    maintainer (search by package maintainer)
    depends (search for packages that depend on keywords)
    makedepends (search for packages that makedepend on keywords)
    optdepends (search for packages that optdepend on keywords)
    checkdepends (search for packages that checkdepend on keywords)
*/

type Package struct {
	ID             int64
	Name           string
	PackageBaseID  int64
	PackageBase    string
	Version        string
	Description    string
	URL            string
	NumVotes       int64
	Popularity     float64
	OutOfDate      int64
	Maintainer     string
	FirstSubmitted int64
	LastModified   int64
	URLPath        string

	Depends      []string
	MakeDepends  []string
	OptDepends   []string
	CheckDepends []string
	Conflicts    []string
	Provides     []string
	Replaces     []string
	Groups       []string
	License      []string
	Keywords     []string
}

type Response struct {
	Version     int64     `json:version`
	Type        string    `json:type`
	ResultCount int64     `json:resultcount`
	Results     []Package `json:results`
}

type DataRequester func(u string, r *Response)

type Config struct {
	base_url  string
	requester DataRequester
	valid_by  []string
}

var byValues []string = []string{
	"name",
	"name-desc",
	"maintainer",
	"depends",
	"makedepends",
	"optdepends",
	"checkdepends",
}

var config = Config{
	base_url:  "https://aur.archlinux.org/rpc/",
	requester: BasicRequest,
	valid_by:  byValues,
}

func SetRequester(r DataRequester) {
	config.requester = r
}
func GetRequester() DataRequester { return config.requester }

func BasicRequest(u string, r *Response) {
	resp, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(r)

	log.Printf("Protocol Versions: %d; Type: %s; Amount of Objects: %d\n",
		r.Version, r.Type, r.ResultCount)
}

func ValidateBy(p string) bool {
	for _, v := range config.valid_by {
		if p == v {
			return true
		}
	}
	return false
}

func BuildSearchUrl(q, by string) (string, error) {
	u := config.base_url + "?v=5&type=search"
	if by != "" {
		if !ValidateBy(by) {
			return "", fmt.Errorf("Wrong value of parameter by: %s", by)
		}
		u += "&by=" + by
	}
	u += "&arg=" + url.QueryEscape(q)
	return u, nil
}

func GetSearch(query string, by string) (Response, error) {
	resp := Response{}
	rq, err := BuildSearchUrl(query, by)
	if err != nil {
		return resp, err
	}

	config.requester(rq, &resp)

	return resp, nil
}

func GetInfo(plist []string) (Response, error) {
	response := Response{}
	return response, nil
}
