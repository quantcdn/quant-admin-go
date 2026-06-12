package quantadmingo

import (
	"net/http"
	"strings"
	"testing"
)

// TestRedactedDumpRequest guards the debug-logging redaction (CodeQL
// go/clear-text-logging): credentials in the URL userinfo and the
// Authorization, Proxy-Authorization and Cookie headers must never appear in
// the bytes handed to the logger, while non-sensitive fields are preserved.
func TestRedactedDumpRequest(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "https://user:supersecretpw@api.example.com/v1/things?q=1", strings.NewReader(`{"k":"v"}`))
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}
	req.Header.Set("Authorization", "Bearer topsecrettoken")
	req.Header.Set("Proxy-Authorization", "Basic cHJveHk6c2VjcmV0")
	req.Header.Set("Cookie", "session=secretcookievalue")

	dump, err := redactedDumpRequest(req)
	if err != nil {
		t.Fatalf("redactedDumpRequest: %v", err)
	}
	out := string(dump)

	// No credential from URL userinfo, headers, or body may reach the log line.
	for _, secret := range []string{"supersecretpw", "topsecrettoken", "cHJveHk6c2VjcmV0", "secretcookievalue", `{"k":"v"}`} {
		if strings.Contains(out, secret) {
			t.Errorf("redacted log leaked secret %q:\n%s", secret, out)
		}
	}
	// The non-credential request line must survive so the log stays useful.
	for _, want := range []string{http.MethodPost, "api.example.com", "/v1/things"} {
		if !strings.Contains(out, want) {
			t.Errorf("expected %q in redacted log, got:\n%s", want, out)
		}
	}
}
