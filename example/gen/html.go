package gen

import (
	"html/template"
	"net/http"
	"net/url"
)

var page = template.Must(template.New("graphiql").Parse(`<!DOCTYPE html>
<html>
  <head>
    <title>{{.title}}</title>
    <link
		rel="stylesheet"
		href="https://cdnjs.webstatic.cn/ajax/libs/graphiql/3.7.1/graphiql.min.css"
	/>
  </head>
  <body style="margin: 0;">
    <div id="graphiql" style="height: 100vh;"></div>

	<script
		src="https://cdnjs.webstatic.cn/ajax/libs/react/18.3.1/umd/react.production.min.js"
	></script>
	<script
		src="https://cdnjs.webstatic.cn/ajax/libs/react-dom/18.3.1/umd/react-dom.production.min.js"
	></script>
	<script
		src="https://cdnjs.webstatic.cn/ajax/libs/graphiql/3.7.1/graphiql.js"
	></script>

    <script>
{{- if .endpointIsAbsolute}}
      const url = {{.endpoint}};
      const subscriptionUrl = {{.subscriptionEndpoint}};
{{- else}}
      const url = location.protocol + '//' + location.host + {{.endpoint}};
      const wsProto = location.protocol == 'https:' ? 'wss:' : 'ws:';
      const subscriptionUrl = wsProto + '//' + location.host + {{.endpoint}};
{{- end}}

      const fetcher = GraphiQL.createFetcher({ url, subscriptionUrl });
      ReactDOM.render(
        React.createElement(GraphiQL, {
          fetcher: fetcher,
          tabs: true,
          headerEditorEnabled: true,
          shouldPersistHeaders: true
        }),
        document.getElementById('graphiql'),
      );
    </script>
  </body>
</html>
`))

// Handler responsible for setting up the playground
func HandlerHtml(title string, endpoint string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		err := page.Execute(w, map[string]interface{}{
			"title":                title,
			"endpoint":             endpoint,
			"endpointIsAbsolute":   endpointHasScheme(endpoint),
			"subscriptionEndpoint": getSubscriptionEndpoint(endpoint),
			"version":              "1.8.2",
			"cssSRI":               "sha256-CDHiHbYkDSUc3+DS2TU89I9e2W3sJRUOqSmp7JC+LBw=",
			"jsSRI":                "sha256-X8vqrqZ6Rvvoq4tvRVM3LoMZCQH8jwW92tnX0iPiHPc=",
			"reactSRI":             "sha256-Ipu/TQ50iCCVZBUsZyNJfxrDk0E2yhaEIz0vqI+kFG8=",
			"reactDOMSRI":          "sha256-nbMykgB6tsOFJ7OdVmPpdqMFVk4ZsqWocT6issAPUF0=",
		})
		if err != nil {
			panic(err)
		}
	}
}

// endpointHasScheme checks if the endpoint has a scheme.
func endpointHasScheme(endpoint string) bool {
	u, err := url.Parse(endpoint)
	return err == nil && u.Scheme != ""
}

// getSubscriptionEndpoint returns the subscription endpoint for the given
// endpoint if it is parsable as a URL, or an empty string.
func getSubscriptionEndpoint(endpoint string) string {
	u, err := url.Parse(endpoint)
	if err != nil {
		return ""
	}

	switch u.Scheme {
	case "https":
		u.Scheme = "wss"
	default:
		u.Scheme = "ws"
	}

	return u.String()
}
