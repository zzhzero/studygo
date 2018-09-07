package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetcher(urluser string) ([]byte, error) {
	req, _ := http.NewRequest(
		http.MethodGet, urluser, nil)
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.6,en;q=0.4")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3493.0 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)

	//resp, err := http.Get(urluser)
	if err != nil {
		//log.Print("Err status not ok:", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Print("Err status not ok:", resp.StatusCode)
		return nil, fmt.Errorf("err status not ok:%d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)

}
