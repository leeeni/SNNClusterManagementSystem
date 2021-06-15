package service

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// 发送post请求
func HttpPost(posturl string, postdata url.Values, header ...string) (map[string]interface{}, error) {

	// 响应数据接收
	var respondStruct map[string]interface{}

	//把post表单发送给目标服务器
	// 建立客户端
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	//提交请求
	request, err := http.NewRequest("POST", posturl, strings.NewReader(postdata.Encode()))

	if err != nil {
		return respondStruct, err
	}

	// 设置发送格式，添加文件头
	if len(header) == 1 {
		request.Header.Add("Content-Type", header[0])
	} else if len(header) == 2 {
		request.Header.Add("Content-Type", header[0])
		request.Header.Add("Authorization-Type", header[1])
	}

	//处理返回结果
	response, err := client.Do(request)
	if err != nil {
		return respondStruct, err
	}

	defer response.Body.Close()

	// 获取返回结信息
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return respondStruct, err
	}

	// 将返回信息结构化
	_ = json.Unmarshal(body, &respondStruct)

	return respondStruct, err
}

// 发送Get请求
func HttpGet(posturl string, header ...string) (map[string]interface{}, error) {
	var respondStruct map[string]interface{}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	//提交请求
	request, err := http.NewRequest("GET", posturl, nil)

	if err != nil {
		return respondStruct, err
	}

	//增加header选项
	request.Header.Add("Authorization", header[0])

	//处理返回结果
	response, err := client.Do(request)
	if err != nil {
		return respondStruct, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return respondStruct, err
	}

	// 将返回信息结构化
	_ = json.Unmarshal(body, &respondStruct)

	return respondStruct, err
}
