package pterodactyl

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"github.com/surdaft/pterodactyl-valheim-discord-egg/pterodactyl/responses"
)

type Client struct {
	Config Config
}

func (ptClient *Client) Get (endpoint string, responseObject interface{}) *Response {
	var (
		err error
		client *http.Client
		headers http.Header
		req *http.Request
		resp *http.Response
		data interface{}
	)

	client = &http.Client{}
	headers = ptClient.getHeaders()

	log.WithFields(log.Fields{
		"uri": endpoint,
		"method": "GET",
	}).Printf("Making request")

	req, err = http.NewRequest("GET", ptClient.Config.ApiUri + endpoint, nil)
	if err != nil {
		return handleErr(nil, err)
	}

	req.Header = headers

	resp, err = client.Do(req)
	if err != nil {
		return handleErr(resp, err)
	}

	data = bodyFromResponse(resp, responseObject)

	return &Response{
		StatusCode: resp.StatusCode,
		Data: data,
	}
}

func (ptClient *Client) Post (endpoint string, postBody *interface{}, responseObject *interface{}) *Response {
	var (
		err error
		client *http.Client
		headers http.Header
		req *http.Request
		resp *http.Response
		data interface{}
		postBodyReader io.Reader
		marshalledPostBody []byte
	)

	client = &http.Client{}
	headers = ptClient.getHeaders()

	marshalledPostBody, err = json.Marshal(postBody)
	postBodyReader = bytes.NewBuffer(marshalledPostBody)

	log.WithFields(log.Fields{
		"uri": endpoint,
		"method": "POST",
	}).Printf("Making request")

	req, err = http.NewRequest("POST", ptClient.Config.ApiUri + endpoint, postBodyReader)
	if err != nil {
		return handleErr(nil, err)
	}

	req.Header = headers

	resp, err = client.Do(req)
	if err != nil {
		return handleErr(resp, err)
	}

	data = bodyFromResponse(resp, responseObject)

	return &Response{
		StatusCode: resp.StatusCode,
		Data: data,
	}
}

func (ptClient *Client) getHeaders () http.Header {
	var headers *http.Header

	headers = &http.Header{}
	headers.Add("Content-Type", "application/json")
	headers.Add("Accept", "Application/vnd.pterodactyl.v1+json")
	headers.Add("Authorization", "Bearer " + ptClient.Config.ClientToken)

	return *headers
}

func handleErr(resp *http.Response, err error) *Response {
	var data *responses.ErrorResponse
	var responseData interface{}

	var statusCode int = 0
	if resp != nil {
		statusCode = resp.StatusCode
	}

	data = &responses.ErrorResponse{
		Data: responseData,
		Message: err.Error(),
	}

	data.Data = bodyFromResponse(resp, responseData)

	return &Response{
		StatusCode: statusCode,
		Data: data,
	}
}

func bodyFromResponse(resp *http.Response, responseObject interface{}) interface{} {
	var (
		err error
		byteData []byte
	)

	if resp == nil {
		return nil
	}

	byteData, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		log.Printf("error: %s", err.Error())
		return nil
	}

	err = json.Unmarshal(byteData, responseObject)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return nil
	}

	return responseObject
}