package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_crud/models"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func (e *Engine) SendRequest() {
	var req models.Request
	fullContent := e.TranslatedContent()
	fullContentByted := []byte(fullContent)
	err := json.Unmarshal(fullContentByted, &req)
	if err != nil {
		log.Fatal("Error in unmarshal content", err.Error())
	}
	payloadJSON, err := json.Marshal(req.Payload)
	if err != nil {
		log.Fatal("Error in marshal content", err.Error())
	}
	payload := bytes.NewReader(payloadJSON)

	if req.Url == "" || req.Method == "" {
		fmt.Println("Invalid format for json file")
		return
	}

	url := req.Url

	result, _ := http.NewRequest(strings.ToUpper(req.Method), url, payload)

	result.Header.Add("Content-Type", e.Header.ContentType)
	result.Header.Add("Authorization", req.Authorization)
	result.Header.Add("User-Agent", e.Header.UserAgent)
	result.Header.Add("Accept", e.Header.Accept)
	result.Header.Add("Cache-Control", e.Header.CacheControl)
	result.Header.Add("Connection", e.Header.Connection)

	res, _ := http.DefaultClient.Do(result)

	defer res.Body.Close()
	resBody, _ := ioutil.ReadAll(res.Body)

	var resInterface interface{}
	err = json.Unmarshal(resBody, &resInterface)
	if err != nil {
		log.Println("Error in unmarshal resBody ", err.Error())
		//fmt.Println("Response: ", string(resBody))
	}

	finalJSON, err := json.MarshalIndent(resInterface, "", "\t")
	if err != nil {
		log.Println("Error in marshal resInterface ", err.Error())
	}

	fmt.Printf("Response: %+v\n", res.Status)
	fmt.Println(string(finalJSON))

}
