package tiltify

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	AuthKey string
	BaseURL string
}

func (c *Client) SetURL(url string) {
	c.BaseURL = url
}

func (c *Client) SetAuthKey(key string) {
	c.AuthKey = key
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	ctx, cancelFunc := context.WithTimeout(context.TODO(), time.Millisecond*1000)
	defer cancelFunc()
	req.Header.Add("Authorization", "Bearer "+c.AuthKey)
	req = req.WithContext(ctx)
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
