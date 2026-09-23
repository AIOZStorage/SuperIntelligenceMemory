package superintelligencememory

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestGoldenSigV4Signature(t *testing.T) {
	body := []byte(`{"vectorBucketName":"my-bucket","statement":"Paris is the capital of France."}`)

	req, err := http.NewRequest(http.MethodPost, "https://example.com/Record", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("new request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	const wantPayloadHash = "51eb3503a79c3820bc6bfa6e0757ef303b0f627b3ce05c23abceb540ec2e4eb4"
	req.Header.Set("X-Amz-Content-Sha256", wantPayloadHash)

	signingTime := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)

	client := &ApiClient{
		accessKeyID:     "AKIDEXAMPLE",
		secretAccessKey: "wJalrXUtnFEMI/K7MDENG+bPxRfiCYEXAMPLEKEY",
		region:          "us-east-1",
	}

	if err := signWithFixedTime(context.Background(), client, req, wantPayloadHash, signingTime); err != nil {
		t.Fatalf("sign: %v", err)
	}

	const wantAuth = "AWS4-HMAC-SHA256 Credential=AKIDEXAMPLE/20260101/us-east-1/s3vectors/aws4_request, SignedHeaders=content-length;content-type;host;x-amz-content-sha256;x-amz-date, Signature=257dee0ecf0805e037c7130077aac364ccf3f58dc2584f83b74733ef68cb8785"

	if got := req.Header.Get("Authorization"); got != wantAuth {
		t.Fatalf("Authorization =\n  %q\nwant\n  %q", got, wantAuth)
	}
	if got := req.Header.Get("X-Amz-Date"); got != "20260101T000000Z" {
		t.Fatalf("X-Amz-Date = %q", got)
	}
}

// TestPrepareRequest_PreservesBaseURLPathPrefix is a regression test for a
// base-URL path-prefix bug: prepareRequest used to resolve the request's
// absolute path (e.g. "/Record") against c.BaseURL via url.URL.Parse,
// which per RFC 3986 reference resolution replaces the base's path entirely
// -- silently dropping any existing path component in the base URL (e.g.
// "https://host/api" + "/Record" became "https://host/Record" instead
// of "https://host/api/Record"). prepareRequest now joins paths via
// path.Join so an existing base path segment is preserved.
func TestPrepareRequest_PreservesBaseURLPathPrefix(t *testing.T) {
	base, err := url.Parse("https://host.example.com/gateway")
	if err != nil {
		t.Fatalf("parse base url: %v", err)
	}

	client := &ApiClient{
		BaseURL:         base,
		accessKeyID:     "AKIDEXAMPLE",
		secretAccessKey: "wJalrXUtnFEMI/K7MDENG+bPxRfiCYEXAMPLEKEY",
		region:          "us-east-1",
	}

	req, err := client.prepareRequest(context.Background(), http.MethodPost, "/Record", nil, nil, "", false)
	if err != nil {
		t.Fatalf("prepareRequest: %v", err)
	}

	const want = "https://host.example.com/gateway/Record"
	if got := req.URL.String(); got != want {
		t.Fatalf("URL = %q, want %q", got, want)
	}
}
