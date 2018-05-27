package firebasedynlink

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// Client Struct
type Client struct {
	// https://firebasedynamiclinks.googleapis.com/v1/shortLinks
	BaseURL   *url.URL
	UserAgent string
	apiKey    string

	httpClient *http.Client
}

// Response type
type Response struct {
	LongDynamicLink string `json:"longDynamicLink"`
	Suffix          struct {
		Option string `json:"option"`
	} `json:"suffix"`
}

// CreateLink function method
func (c *Client) CreateLink(longURL string) (Response, error) {
	req, err := c.newRequest("POST", "/v1/shortlinks?key=c.api_key", nil)

	if err != nil {
		var r Response
		return r, err
	}

	var res Response
	_, err = c.do(req, &res)
	return res, nil
}

// Creates a new request
func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
