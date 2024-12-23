package chatgptfree

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"strconv"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

const (
	urlChatGPTText = "https://main.gpt-chatbotru-4-o1.ru"
	uriChatGPTText = urlChatGPTText + "/api/openai/v1/chat/completions"
)

var (
	errResponseCodeIsNot200 = errors.New("response code is not 200")
	errEmptyRespChoices     = errors.New("response's choices are empty")
)

var client = fasthttp.Client{
	TLSConfig: &tls.Config{InsecureSkipVerify: true},
}

func GenerateText(ctx context.Context, prompt string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetRequestURI(uriChatGPTText)
	req.Header.Set(fasthttp.HeaderContentType, "application/json")
	req.Header.Set(fasthttp.HeaderOrigin, urlChatGPTText)
	req.SetBody(prepareRequestBody(prompt))
	bodyChan := make(chan []byte, 1)
	errChan := make(chan error, 1)
	go func() {
		if err := client.Do(req, resp); err != nil {
			errChan <- err
			return
		}
		if resp.StatusCode() != fasthttp.StatusOK {
			errChan <- errResponseCodeIsNot200
			return
		}
		var p fastjson.Parser
		v, err := p.ParseBytes(resp.Body())
		if err != nil {
			errChan <- err
			return
		}
		choices := v.GetArray("choices")
		if len(choices) == 0 {
			errChan <- errEmptyRespChoices
			return
		}
		bodyChan <- choices[0].Get("message").GetStringBytes("content")
	}()
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case body := <-bodyChan:
		return body, nil
	case err := <-errChan:
		return nil, err
	}
}

func prepareRequestBody(content string) []byte {
	var b bytes.Buffer
	b.WriteString(`{"messages":[{"role":"user","content":`)
	b.WriteString(strconv.Quote(content))
	b.WriteString(`}],"stream":false,"model":"gpt-4o","temperature":0.5,"presence_penalty":0,"frequency_penalty":0,"top_p":1,"chat_token":126,"captchaToken":"1"}`)
	return b.Bytes()
}
