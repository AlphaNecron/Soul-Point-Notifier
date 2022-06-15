package noti

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func fetch() (bool, WorldList) {
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	req, _ := http.NewRequest("GET", "https://athena.wynntils.com/cache/get/serverList", nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "SPNoti/0.1")
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return false, WorldList{}
	}
	b, rErr := io.ReadAll(res.Body)
	if rErr != nil {
		log.Fatalln(err)
		return false, WorldList{}
	}
	var list WorldList
	jErr := json.Unmarshal(b, &list)
	if jErr != nil {
		log.Fatalln(err)
		return false, WorldList{}
	}
	return true, list
}
