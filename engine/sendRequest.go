package engine

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"gocrud/models"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// SendRequest is the final step method for prepare crud request based on files and selected variables and
// demonstrate result of request
func (e *Engine) SendRequest() {
	var req models.Request
	fullContent := e.TranslatedContent()
	fullContentByted := []byte(fullContent)
	err := json.Unmarshal(fullContentByted, &req)
	if err != nil {
		//log.Fatal("Error in unmarshal content ", err.Error())
		fmt.Println("ERROR: JSON format is wrong, ", e.SelectedDir, "/", e.SelectedFile)
		fmt.Println("Raw content: ", e.Content)
		return

	}
	payloadJSON, err := json.Marshal(req.Payload)
	if err != nil {
		log.Fatal("Error in marshal content ", err.Error())
	}
	payload := bytes.NewReader(payloadJSON)

	if req.Url == "" || req.Method == "" {
		fmt.Println("Error: Invalid format for json file")
		return
	}

	url := req.Url
	stopSignal := make(chan bool, 1)

	go func(stop chan bool) {
		for {
			select {
			case <-stop:
				return
			case <-time.After(800 * time.Millisecond):
				fmt.Print(".")

			}
		}
	}(stopSignal)

	result, _ := http.NewRequest(strings.ToUpper(req.Method), url, payload)

	result.Header.Add("Content-Type", e.Header.ContentType)
	result.Header.Add("Authorization", req.Authorization)
	result.Header.Add("User-Agent", e.Header.UserAgent)
	result.Header.Add("Accept", e.Header.Accept)
	result.Header.Add("Cache-Control", e.Header.CacheControl)
	result.Header.Add("Connection", e.Header.Connection)

	// Ignore TLS certificate verification
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}
	res, err := client.Do(result)


	//res, err := http.DefaultClient.Do(result)
	stopSignal <- true
	if err != nil {
		fmt.Println("\nError: ", err.Error())
	} else {

		defer res.Body.Close()
		resBody, _ := ioutil.ReadAll(res.Body)
		fmt.Printf("Status: %+v\n", res.Status)

		var resInterface interface{}
		err = json.Unmarshal(resBody, &resInterface)
		if err != nil {
			fmt.Println("Response is not JSON")
			fmt.Println(string(resBody))
		} else {
			finalJSON, err := json.MarshalIndent(resInterface, "", "\t")
			if err != nil {
				fmt.Println("Error in marshal resInterface ", err.Error())
			}

			fmt.Println(string(finalJSON))
		}
	}

}
