package helper

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/QuantumNous/new-api/dto"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	relayconstant "github.com/QuantumNous/new-api/relay/constant"
	"github.com/QuantumNous/new-api/setting/ratio_setting"
	"github.com/gin-gonic/gin"
)

func newModelMappedTestContext() *gin.Context {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/v1/chat/completions", nil)
	return c
}

func TestModelMappedHelperNormalizeGeminiFlowImageModel(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "gemini-2.5-flash-image-flow",
			input: "gemini-2.5-flash-image-Flow",
			want:  "gemini-2.5-flash-image",
		},
		{
			name:  "gemini-3.0-pro-image-flow",
			input: "gemini-3.0-pro-image-Flow",
			want:  "gemini-3.0-pro-image",
		},
		{
			name:  "gemini-3.1-flash-image-flow",
			input: "gemini-3.1-flash-image-Flow",
			want:  "gemini-3.1-flash-image",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			c := newModelMappedTestContext()
			info := &relaycommon.RelayInfo{
				OriginModelName: tt.input,
			}
			req := &dto.GeneralOpenAIRequest{Model: info.OriginModelName}

			if err := ModelMappedHelper(c, info, req); err != nil {
				t.Fatalf("ModelMappedHelper returned error: %v", err)
			}

			if info.UpstreamModelName != tt.want {
				t.Fatalf("unexpected upstream model, got %q want %q", info.UpstreamModelName, tt.want)
			}
			if info.OriginModelName != tt.want {
				t.Fatalf("unexpected origin model, got %q want %q", info.OriginModelName, tt.want)
			}
			if req.Model != tt.want {
				t.Fatalf("unexpected request model, got %q want %q", req.Model, tt.want)
			}
		})
	}
}

func TestModelMappedHelperNormalizeGeminiFlowImageModelWithCompactSuffix(t *testing.T) {
	c := newModelMappedTestContext()
	info := &relaycommon.RelayInfo{
		OriginModelName: ratio_setting.WithCompactModelSuffix("gemini-2.5-flash-image-Flow"),
		RelayMode:       relayconstant.RelayModeResponsesCompact,
	}
	req := &dto.GeneralOpenAIRequest{Model: info.OriginModelName}

	if err := ModelMappedHelper(c, info, req); err != nil {
		t.Fatalf("ModelMappedHelper returned error: %v", err)
	}

	const expectedUpstream = "gemini-2.5-flash-image"
	expectedOrigin := ratio_setting.WithCompactModelSuffix(expectedUpstream)

	if info.UpstreamModelName != expectedUpstream {
		t.Fatalf("unexpected upstream model, got %q want %q", info.UpstreamModelName, expectedUpstream)
	}
	if info.OriginModelName != expectedOrigin {
		t.Fatalf("unexpected origin model, got %q want %q", info.OriginModelName, expectedOrigin)
	}
	if req.Model != expectedUpstream {
		t.Fatalf("unexpected request model, got %q want %q", req.Model, expectedUpstream)
	}
}

func TestModelMappedHelperKeepsNonGeminiFlowSuffix(t *testing.T) {
	c := newModelMappedTestContext()
	info := &relaycommon.RelayInfo{
		OriginModelName: "gpt-4o-Flow",
	}
	req := &dto.GeneralOpenAIRequest{Model: info.OriginModelName}

	if err := ModelMappedHelper(c, info, req); err != nil {
		t.Fatalf("ModelMappedHelper returned error: %v", err)
	}

	const expected = "gpt-4o-Flow"
	if info.UpstreamModelName != expected {
		t.Fatalf("unexpected upstream model, got %q want %q", info.UpstreamModelName, expected)
	}
	if info.OriginModelName != expected {
		t.Fatalf("unexpected origin model, got %q want %q", info.OriginModelName, expected)
	}
	if req.Model != expected {
		t.Fatalf("unexpected request model, got %q want %q", req.Model, expected)
	}
}
