package main

import (
	"github.com/gin-gonic/gin"
	routers "MA.Content.Services.OrgMapper/src/routers"
)

func main() {
	r := gin.Default()
	orgMapper := r.Group("orgmapper")
	routers.InitRoutes(orgMapper);
	r.Run(":9092")
}

// func basicAuth(username, password string) string {
// 	auth := `apiuser:Eag13^mdc`
// 	return base64.StdEncoding.EncodeToString([]byte(auth))
// }

// func redirectPolicyFunc(req *http.Request, via []*http.Request) error{
// 	req.Header.Add("Authorization","Basic " + basicAuth("apiuser","Eag13^mdc"))
// 	return nil
// }

// func resolveSolr(ch chan DocsInfo) {
// 	url := `http://ftc-lbeagfus302:8764/api/apollo/query-pipelines/eai_search_entity/collections/eai_search/select?start=0&rows=100&q=541000&fq=(org_has_bfm:(Y)%20OR%20org_has_cfm:(Y)%20OR%20org_has_edf:(Y)%20OR%20org_has_mir:(Y)%20OR%20org_has_rds:(Y)%20OR%20org_has_unrated:(Y))`
// 		req, _ := http.NewRequest("GET", url, nil)
// 		req.Header.Add("Authorization","Basic " + basicAuth("apiuser","Eag13^mdc"))
// 		client := &http.Client{}	
// 		resp, _ := client.Do(req)	
// 		data, _ := ioutil.ReadAll(resp.Body)
// 		var om OrgMapping
// 		err := json.Unmarshal(data, &om)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println("solre")
// 		ch <- om.Grouped.OdsRefIdGroup.Groups[0].Doclist.Docs[0]
		
// }

// func formatJson(val []byte, or interface{}) {
// 	err := json.Unmarshal(val, &or)
// 	if err != nil {
// 		panic(err)
// 	}
// }
