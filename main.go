package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	for {
		for _, proxyString := range proxyList {
			doAPI(proxyString)
		}
	}
}

// ? กำหนด URL ของ Proxy Server
var proxyList = []string{
	// IPV4 With username and password
	// "http://username:password@ip:port",
}

func doAPI(proxyString string) {
	proxyURL, err := url.Parse(proxyString)
	if err != nil {
		fmt.Println("Error parsing proxy URL:", err)
		return
	}

	// กำหนด Client ที่ใช้งาน Proxy
	transport := &http.Transport{}
	if proxyString != "" {
		transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	// กำหนด Client และ Transport ที่ใช้งาน Proxy
	client := &http.Client{
		Transport: transport,
	}

	// ส่ง Request ผ่าน Proxy และรับ Response
	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/todos", nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	// อ่าน Response Body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading HTTP response body:", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("HTTP status code:", resp.StatusCode)
		return
	}
	// แสดงผลลัพธ์
	fmt.Println(string(body))
}
