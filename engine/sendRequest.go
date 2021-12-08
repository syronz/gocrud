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
	"os"
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

	var reqs []models.Request
	err2 := json.Unmarshal(fullContentByted, &reqs)

	if err != nil && err2 != nil {
		fmt.Println("ERROR: JSON format is wrong, ", e.SelectedDir, "/", e.SelectedFile)
		fmt.Println("Raw content: ", e.Content)
		return
	}

	if len(reqs) == 0 {
		reqs = append(reqs, req)
	}

	for i := 0; i < len(reqs); i++ {
		req = reqs[i]
		time.Sleep(time.Duration(e.Config.Delay) * time.Millisecond)
		fmt.Printf("...\n...\n...\n")

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
		stopDot := make(chan bool, 1)

		go func(stop chan bool) {
			for {
				select {
				case <-stop:
					return
				case <-time.After(800 * time.Millisecond):
					fmt.Print(".")
				}

			}

		}(stopDot)

		go func(stop chan bool) {
			for {
				select {
				case <-stop:
					stopDot <- true
					return
				case <-time.After(time.Duration(e.Config.Timeout) * time.Second):
					fmt.Println("timeout exceeded!")
					stopDot <- true
					os.Exit(3)
				}
			}
		}(stopSignal)

		result, _ := http.NewRequest(strings.ToUpper(req.Method), url, payload)

		for _, v := range req.Headers {
			result.Header.Add(v.Name, v.Value)
		}

		result.Header.Add("Authorization", req.Authorization)
		result.Header.Add("Content-Type", e.Header.ContentType)
		result.Header.Add("User-Agent", e.Header.UserAgent)
		result.Header.Add("Accept", e.Header.Accept)
		result.Header.Add("Cache-Control", e.Header.CacheControl)
		result.Header.Add("Connection", e.Header.Connection)

		// Ignore TLS certificate verification (https)
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: transport}
		res, err := client.Do(result)

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

	} // end of for

}
