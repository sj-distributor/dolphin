package templates

import (
	"strings"
	"testing"
)

func TestHTTPHandlerTemplateDoesNotGenerateAuthenticationOrSigningSecrets(t *testing.T) {
	for _, forbidden := range []string{
		"MySecret",
		"MakeToken",
		"ParseJWT",
		"SigningMethodHS256",
		`"{{.Config.Package}}/auth"`,
	} {
		if strings.Contains(HttpHandler, forbidden) {
			t.Errorf("HTTP handler template must not contain %q", forbidden)
		}
	}
}
