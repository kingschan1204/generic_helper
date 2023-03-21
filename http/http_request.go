package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HttpGet send a http request to get data
func HttpGet(url string, httpHead map[string]string) (string, error) {
	// Create a custom request object
	request, err := http.NewRequest("GET", url, nil)
	// 检查错误
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// set http head attribute
	for k, v := range httpHead {
		request.Header.Set(k, v)
	}
	// Use the Do method to make a request and get a response
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// Close the connection after reading the response
	defer response.Body.Close()
	// read response content
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// Print response content and status code
	result := string(body)
	fmt.Println(result)
	fmt.Println(response.StatusCode)
	return result, nil
}

// HttpPost send a http request to get data
// requestBody : `{"name":"xxx","message":"xxx"}`
func HttpPost(url string, requestBody string, httpHead map[string]string) (string, error) {
	// Declare a JSON data variable
	data := []byte(requestBody)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// set http head attribute
	for k, v := range httpHead {
		request.Header.Set(k, v)
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// Print response content and status code
	result := string(body)
	fmt.Println(result)
	fmt.Println(response.StatusCode)
	return result, nil
}
