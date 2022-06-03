package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	aur "northbear/aurer/aurweb_rpc"
	"os"
)

/*
  search
  info
  download
  build
*/

func GetAction(argv []string, cf string) (Action, error) {
	action := ActionSearch{pattern: "zoom"}

	return action, nil
}

// type PackageContent map[string]string

type PackageDescr struct {
	ID          string
	Name        string
	Description string
	Version     string
	URL         string
	//	data PackageContent
}

type AurResponse struct {
	Version     int64          `json:version`
	Type        string         `json:type`
	ResultCount int64          `json:resultcount`
	Results     []PackageDescr `json:results`
	//	Results []map[string]interface{} `json:results`
}

func GetRequest(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return fmt.Sprintf("%s", body), err
}

func Search(pattern string, by string) []PackageDescr {
	var ar AurResponse

	base_url := "https://aur.archlinux.org/rpc/?v=5"
	u, err := url.Parse(base_url)
	if err != nil {
		log.Printf("Incorrect base url: %s", base_url)
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("type", "search")
	q.Set("arg", pattern)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	// search_query := "https://aur.archlinux.org/rpc/?v=5&type=search&arg=%s"
	// resp, err := http.Get(fmt.Sprintf(search_query, url.QueryEscape(pattern)))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&ar)

	log.Printf("Protocol Versions: %d; Type: %s; Amount of Objects: %d\n",
		ar.Version, ar.Type, ar.ResultCount)

	return ar.Results
}

func GetInfo(pattern string) []PackageDescr {
	var ar AurResponse
	return ar.Results
}

func main() {
	config := NewConfig()
	if err := config.Init(os.Args); err != nil {
		log.Fatal(err)
	}

	action, err := config.GetAction()
	if err != nil {
		log.Fatal("Error: ", err)
	}

	switch act := action.(type) {
	case ActionUsage:
		usage := act
		fmt.Println(usage.message)
		os.Exit(0)
	case ActionSearch:
		search := act
		response, err := aur.GetSearch(search.Pattern(), search.ByField())
		if err != nil {
			log.Fatal("Search cannot get response: ", err)
		}
		for c, pkg := range response.Results {
			// fmt.Printf("Package name: %s\n", pkg.Name)
			// fmt.Printf("Descr: %s\n", pkg.Description)
			// fmt.Printf("Version: %s\n", pkg.Version)
			fmt.Printf("%s %s\n", pkg.Name, pkg.Version)
			fmt.Printf("    %s\n", pkg.Description)
			fmt.Printf("    https://aur.archlinux.org%s\n", pkg.URLPath)
			if int64(c+1) < response.ResultCount {
				fmt.Println("")
			}
		}
		fmt.Println("+++ End +++")
	case ActionGetInfo:
		ai := act
		response, err := aur.GetInfo(ai.Pattern())
		if err != nil {
			log.Fatal("GetInfo cannot get response: ", err)
		}
		for _, pkg := range response.Results {
			fmt.Println("Name: ", pkg.Name)
		}
	}

	fmt.Printf("Action executed: %s\n", action.Name())
	fmt.Printf("CLI params: %v\n", os.Args)
	// fmt.Printf("Version: %s", aurs.GetVersion())
	// fmt.Println("Hello, world!", cliopt.SearchPattern, "...")
}

// wget https://aur.archlinux.org/cgit/aur.git/snapshot/taisei.tar.gz -O - | tar xzC sources/
