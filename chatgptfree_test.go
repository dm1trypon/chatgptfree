package chatgptfree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_prepareRequestBody(t *testing.T) {
	tests := []struct {
		name           string
		style          string
		prompt         string
		negativePrompt string
		width          int
		height         int
		exp            string
	}{
		{
			name:   "Successful assembly of the request body",
			prompt: "test prompt",
			exp:    `{"messages":[{"role":"user","content":"test prompt"}],"stream":false,"model":"gpt-4o","temperature":0.5,"presence_penalty":0,"frequency_penalty":0,"top_p":1,"chat_token":126,"captchaToken":"1"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.exp, string(prepareRequestBody(tt.prompt)))
		})
	}
}
