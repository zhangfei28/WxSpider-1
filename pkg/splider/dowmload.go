package splider

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	//"github.com/axgle/mahonia"
)

//加载数据
func HttpGet(urls, ip string) (io.Reader, int) {
	var client *http.Client

	proxy, err := url.Parse(ip)

	if ip != "local" {
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
			},
			Timeout: time.Second * 1,
		}
	} else {
		client = &http.Client{}
	}

	req, err := http.NewRequest("GET", urls, nil)
	if err != nil {
		log.Println("http.newRequest error:", err)
	}

	req.Header.Add("User-Agent", GetUserAgent())
	//req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	//req.Header.Add("Connection", "keep-alive")

	resp, err := client.Do(req)
	if err != nil {
		return nil, 500
	}

	return resp.Body, resp.StatusCode
}
