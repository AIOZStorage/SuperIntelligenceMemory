package superintelligencememory

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecall_NotFoundIsASuccessfulValue_NotAnError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(RecallResponse{Found: PtrBool(false)})
	}))
	defer srv.Close()

	opts := DefaultApiClientOptions()
	opts.BaseURL(srv.URL)
	opts.Credentials("AKIDEXAMPLE", "wJalrXUtnFEMI/K7MDENG+bPxRfiCYEXAMPLEKEY")
	client, err := NewApiClient(context.Background(), opts)
	if err != nil {
		t.Fatalf("NewApiClient: %v", err)
	}

	res, err := client.Memory.RecallWithContext(context.Background(), RecallRequest{
		VectorBucketName: PtrString("my-bucket"),
		Question:         PtrString("What is France's capital?"),
	})
	if err != nil {
		t.Fatalf("Recall returned an error for a not-found result: %v", err)
	}
	if res.GetFound() {
		t.Fatalf("Found = true, want false")
	}
}

func TestRecord_ErrorResponseDecodesTypedError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(AgentMemoryAPIError{Type: "ValidationException", Message: "statement is required"})
	}))
	defer srv.Close()

	opts := DefaultApiClientOptions()
	opts.BaseURL(srv.URL)
	opts.Credentials("AKIDEXAMPLE", "wJalrXUtnFEMI/K7MDENG+bPxRfiCYEXAMPLEKEY")
	client, err := NewApiClient(context.Background(), opts)
	if err != nil {
		t.Fatalf("NewApiClient: %v", err)
	}

	_, err = client.Memory.RecordWithContext(context.Background(), RecordRequest{
		VectorBucketName: PtrString("my-bucket"),
		Statement:        PtrString(""),
	})
	if err == nil {
		t.Fatalf("expected an error for a 400 response, got nil")
	}
	apiErr, ok := err.(*AgentMemoryAPIError)
	if !ok {
		t.Fatalf("error is %T, want *AgentMemoryAPIError", err)
	}
	if apiErr.Type != "ValidationException" || apiErr.Message != "statement is required" {
		t.Fatalf("apiErr = %+v", apiErr)
	}
}
