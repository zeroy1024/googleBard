package translate

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Google struct {
	BaseURL string
	Client  string
	SrcLang string
	DstLang string
	Dt      string
}

func NewGoogle(srcLang, dstLang string) *Google {
	return &Google{
		BaseURL: "https://translate.google.com/translate_a/single",
		Client:  "gtx",
		SrcLang: srcLang,
		DstLang: dstLang,
		Dt:      "t",
	}
}

func (g *Google) handleResponse(response *http.Response) (string, error) {
	bodyBytes, _ := io.ReadAll(response.Body)
	var responseBody []interface{}
	err := json.Unmarshal(bodyBytes, &responseBody)
	if err != nil {
		return "", err
	}

	str := ""
	for _, v := range responseBody[0].([]interface{}) {
		str += v.([]interface{})[0].(string)
	}

	return str, nil
}

func (g *Google) Translate(keyword string) (string, error) {
	params := url.Values{}
	params.Add("client", g.Client)
	params.Add("sl", g.SrcLang)
	params.Add("tl", g.DstLang)
	params.Add("dt", g.Dt)
	params.Add("q", keyword)

	request, _ := http.NewRequest("GET", g.BaseURL, nil)
	request.URL.RawQuery = params.Encode()

	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	return g.handleResponse(response)
}
