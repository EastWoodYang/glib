package glib

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

/* ================================================================================
 * Http模块
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Http Get请求
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func HttpGet(url string, args ...string) (string, error) {
	requestUrl := url
	if len(args) == 1 {
		params := args[0]
		requestUrl = fmt.Sprintf("%s?%s", url, params)
	}

	resp, err := http.Get(requestUrl)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Http POST请求
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func HttpPost(url, params string, args ...string) (string, error) {
	cookie := ""
	if len(args) == 1 {
		cookie = args[0]
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(params))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if len(cookie) > 0 {
		req.Header.Set("Cookie", cookie)
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 上传文件
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func HttpPostFile(filename string, url string) (int, string, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	outFileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		return 0, "", err
	}

	// open file handle
	inFileData, err := os.Open(filename)
	if err != nil {
		return 0, "", err
	}

	// iocopy
	_, err = io.Copy(outFileWriter, inFileData)
	if err != nil {
		return 0, "", err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, "", err
	}

	statusCode, _ := strconv.Atoi(resp.Status)

	return statusCode, string(respBody), err
}
