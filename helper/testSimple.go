package helper

/*j
import (
	"encoding/json"
	"fmt"
	"go_crud/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
func SimplePost() {
	f, err := os.Open("funds/post.json")
	if err != nil {
		log.Fatal("can't open post.json")
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	basic := models.Basic{}
	err = decoder.Decode(&basic)
	if err != nil {
		log.Fatal("can't decode basic.json: ", err)
	}
	rr := basic.Payload.(map[string]interface{})
	kk := KeysString(rr)
	url := "http://127.0.0.1:8080/m/funds"

	//payload := strings.NewReader("{\n    \"type\": \"in\",\n    \"invoice_id\": 1,\n    \"peer_id\": null,\n    \"amount\": 1000\n}")
	payload := strings.NewReader(kk)
	log.Println(">>>>>>>>>>>     ", kk)

	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFwcHVzZXIiLCJleHAiOjE1NzQ2ODU3Njh9.5xS1e9CdSoYWiOqwUbrvRGDdku03ObpF-zP8F20z1Go")
	req.Header.Add("User-Agent", "PostmanRuntime/7.17.1")
	req.Header.Add("Accept", "* /*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "9949e72f-4f37-4834-b0fa-e334472c585c,0cd3c3d6-52aa-49e5-b8d2-81812a3146f8")
	req.Header.Add("Host", "127.0.0.1:8080")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Content-Length", "82")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
*/
