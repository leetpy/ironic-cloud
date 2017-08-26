package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Auth struct {
	TenantName          string            `json:"tenantName"`
	PasswordCredentials map[string]string `json:"passwordCredentials"`
}

type Token struct {
	KAuth Auth `json:"auth"`
}

func GetToken(username, password string, address IPAddr) (string, error) {

	var pwd = map[string]string{
		"username": username,
		"password": password,
	}
	var auth = Auth{TenantName: "admin", PasswordCredentials: pwd}
	var token = Token{KAuth: auth}

	client := &http.Client{}
	url := "http://10.43.210.22:5000/v2.0/tokens"
	data, err := json.Marshal(token)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	request, err := http.NewRequest("POST", url, bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err.Error())
	}

	response, _ := client.Do(request)
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var f interface{}
	if err = json.Unmarshal(body, &f); err != nil {
		fmt.Println(err.Error())
	}

	m := f.(map[string]interface{})
	access := m["access"].(map[string]interface{})
	tk := access["token"].(map[string]interface{})
	fmt.Println(tk["id"].(string))

}
