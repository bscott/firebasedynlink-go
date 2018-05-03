package firebasedynlink

type Client struct {
  // https://firebasedynamiclinks.googleapis.com/v1/shortLinks
  BaseURL *url.URL
  UserAgent string
  apiKey  string
  
  httpClient *http.Client
}

type Request struct {
	LongDynamicLink string `json:"longDynamicLink"`
	Suffix          struct {
		Option string `json:"option"`
	} `json:"suffix"`
}

func (c *Client) CreateLink(longURL string) {
  
}

