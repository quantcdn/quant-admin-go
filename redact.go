package quantadmingo

import (
	"fmt"
	"net/http"
)

// redactedDumpRequest renders an outbound request for debug logging as a single
// request line built only from non-credential fields (method, scheme, host,
// path/query, protocol). Header values and the body are intentionally omitted:
// request.Header is tainted by the Authorization / basic-auth credentials, so
// logging any header value — even from a redacted copy — trips CodeQL
// go/clear-text-logging. request.URL.User (basic-auth userinfo) is never read,
// so credentials cannot reach the log.
//
// This file is recreated verbatim by the SDK pipeline
// (portal .github/scripts/post-generate-sdk.sh), which also rewrites the single
// call site in client.go, so the redaction survives every regeneration. Keep
// this file and that script's heredoc identical.
func redactedDumpRequest(request *http.Request) ([]byte, error) {
	scheme := request.URL.Scheme
	if scheme == "" {
		scheme = "https"
	}
	line := fmt.Sprintf("%s %s://%s%s %s", request.Method, scheme, request.URL.Host, request.URL.RequestURI(), request.Proto)
	return []byte(line), nil
}
