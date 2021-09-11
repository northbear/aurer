package main

import (
	"fmt"
	"io"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

func Usage() string {
	return "Usage: aurer <options> <pattern>|<package-name>"
}


type Action interface{
	Name() string
}

type ActionSearch struct{
	pattern string
	by string
}

// func New(pattern, by string) *ActionSearch { return &ActionSearch{ pattern: pattern, by: by } }
func (a ActionSearch) Name() string { return "Search" } 
func (a ActionSearch) Pattern() string { return a.pattern }
func (a ActionSearch) ByField() string { return a.by }

type ActionGetInfo struct{
	pkg_name string
}

// func New(pkg string) ActionGetInfo { return &ActionGetInfo{ pkg_name: pkg } }
func (a ActionGetInfo) Name() string { return "GetInfo" }
func (a ActionGetInfo) Pattern() string { return a.pkg_name }


func GetAction(argv []string, cf string) (Action, error) {
	action := ActionSearch{ pattern: "zoom" }
	
	return action, nil
}

// type PackageContent map[string]string

type PackageDescr struct {
	ID string
	Name string
	Description string
	Version string
	URL string
	//	data PackageContent
}

type AurResponse struct{
	Version int64 `json:version`
	Type string `json:type`
	ResultCount int64 `json:resultcount`
	Results []PackageDescr `json:results`
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

type Config struct{
	Args *[]string
}

func NewConfig() *Config { return &Config{} }
func (c Config) Init(args *[]string) error {
	return nil
}

func (c Config) GetAction() (Action, error) {
	action := ActionSearch{ pattern: "zoom" }
	
	return action, nil
}

func main() {	
	config := NewConfig()
	if err := config.Init(&os.Args); err != nil {
		log.Fatal(err)
	}
	
	action, err := GetAction(os.Args, "")
	if err != nil {
		fmt.Println(Usage())
		os.Exit(0)
	}
	switch action.(type) {
	case ActionSearch:
		as := action.(ActionSearch)
		pdl := Search(as.Pattern(), as.ByField())
		for c, pd := range pdl {
			fmt.Printf("Package name: %s\n", pd.Name)
			fmt.Printf("Descr: %s\n", pd.Description)
			fmt.Printf("Version: %s\n", pd.Version)
			fmt.Printf("URL: %s\n", pd.URL)
			if c + 1 < len(pdl) {
				fmt.Println("")
			}
		}
		fmt.Println("+++ End +++")
	case ActionGetInfo:
		ai := action.(ActionGetInfo)
		info_list := GetInfo(ai.Pattern())
		for _, pd := range info_list {
			fmt.Println("Name: ", pd.Name)
		}
	}
	
	fmt.Printf("Action executed: %s\n", action.Name())
	fmt.Printf("CLI params: %v\n", os.Args)
	// fmt.Printf("Version: %s", aurs.GetVersion())
	// fmt.Println("Hello, world!", cliopt.SearchPattern, "...")
}
