package snov

import (
	"fmt"
	"getEmail/lib"
	"github.com/tidwall/gjson"
	"net/url"
	"regexp"
)

// GetToken 获取token
func GetToken() string{
	data := make(url.Values)
	url := "https://api.snov.io/v1/oauth/access_token"
	data["grant_type"] = []string{"client_credentials"}
	data["client_id"] = []string{"d47803b976c4a49c450ef673f6594a17"}
	data["client_secret"] = []string{"90dbd0018dd18eae5a3e1c3e9c3c8c64"}

	a := lib.Post(url,data)

	re1 := regexp.MustCompile(`access_token":".+`)
	if re1 == nil {
		fmt.Println("regexp err")
		return "mistake"
	}
	tokenSlice := re1.FindAllStringSubmatch(a, -1)
	token := tokenSlice[0][0]
	token= token[15:len(token)-42]
	return token
}

// GetEmail 获取Email
func GetEmail (domin string , lastId string) string {
	dm := domin
	tk := GetToken()
	tp := "all"
	limit := "100"
	lId := lastId
	url := "https://api.snov.io/v2/domain-emails-with-info/?"
	value := "domain="+dm+"&type="+tp+"&limit="+limit+"&lastId="+lId+"&access_token="+tk
	url = url +value
	html := lib.Get(url)
	return html
}

// JsonToSlice JsonToMap 把json数据解析然后拿出来
func JsonToSlice(jsonValue string ){
	jV := jsonValue
	value := gjson.Get(jV,"emails.#.email")
	lastId := gjson.Get(jV,"lastId")
	result := value.Array()
	for _,v := range result {
		fmt.Println(v)
	}
	fmt.Println("lastId",lastId)

}

