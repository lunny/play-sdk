package play

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Result struct {
	Errors string                   `json:"Errors"`
	Events []map[string]interface{} `json:"Events"`
}

type FormatResult struct {
	Body  string `json:"Body"`
	Error string `json:"Error"`
}

func Format(codes string) (string, error) {
	client := &http.Client{}
	resp, err := client.PostForm("http://play.golang.org/fmt",
		url.Values{"body": {codes}})
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}

	res := new(FormatResult)
	err = json.Unmarshal(body, res)
	if err == nil {
		if res.Error == "" {
			return res.Body, nil
		}
		return "", errors.New(res.Error)
	}

	return "", err
}

func compileBytes(codes string) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.PostForm("http://play.golang.org/compile",
		url.Values{"version": {"2"}, "body": {codes}})
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, nil
}

func CompileCode(codes string) (string, error) {
	res, err := compileBytes(codes)
	return string(res), err
}

func Compile(codes string) (*Result, error) {
	res, err := compileBytes(codes)
	if err != nil {
		return nil, err
	}

	result := new(Result)
	err = json.Unmarshal(res, result)
	return result, err
}

func Share(codes string) (string, error) {
	client := &http.Client{}
	resp, err := client.Post("http://play.golang.org/share",
		"application/x-www-form-urlencoded", strings.NewReader(codes))
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("http://play.golang.org/p/%v", string(body)), nil
}
