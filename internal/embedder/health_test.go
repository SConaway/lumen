// Copyright 2026 Aeneas Rekkas
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package embedder

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ory/lumen/internal/config"
)

func TestProbeServer_OpenAI_TrailingV1InHostNotDoubled(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/models" {
			t.Errorf("expected path /v1/models, got %q (host's /v1 must not be duplicated)", r.URL.Path)
		}
		_, _ = w.Write([]byte(`{"data":[{"id":"test-model"}]}`))
	}))
	defer server.Close()

	srv := config.ServerConfig{Backend: config.BackendOpenAI, Host: server.URL + "/v1", Model: "test-model"}
	if err := ProbeServer(context.Background(), srv); err != nil {
		t.Fatalf("ProbeServer: %v", err)
	}
}
