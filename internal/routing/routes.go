package routing

import (
	"fmt"
	"io"
	"net/http"

	"github.com/nturbo1/reverse-proxy/internal/log"
	"github.com/nturbo1/reverse-proxy/internal/configs"
)

// Routes struct represents the configurations parsed from the master file of the routes configurations.
type Routes struct {
	Files []string `json:"files"` // filenames of the files that contain routing configurations
}

// Route struct represents a route configuration information parsed from the routes configuration
// files.
// Fields:
//   - Host: a host of the underlying backend service that will be called by the api gateway.
//   - Port: a port of the underlying backend service that will be called by the api gateway.
//   - Endpoint: an endpoint of the api gateway.
//   - BackendEndpoint: the backend service endpoint that is mapped to the api gateway endpoint.
//   - Proto: the application protocol used for communication between the api gateway and the internal
//     backend servers.
//
// Any substring value that is preceded by '$', contains only one ore more alphanumeric and '_'
// (underscore) characters inside string values of the Route struct fields are considered to be an
// environment variable name.
type Route struct {
	Host            string `json:"host"`
	Port            string `json:"port"`
	Endpoint        string `json:"endpoint"`
	BackendEndpoint string `json:"backendEndpoint"`
	Proto           string `json:"proto"`
	Method          string `json:"method"`
}
func (r Route) String() string {
	return fmt.Sprintf(
		"{host: %s, port: %s, endpoint: %s, backendEndpoint: %s, proto: %s, method: %s}",
		r.Host,
		r.Port,
		r.Endpoint,
		r.BackendEndpoint,
		r.Proto,
		r.Method,
	)
}

func GetRoutes(appConfigs *configs.AppConfigs) ([]*Route, error) {
	routes, err := parseRoutes(appConfigs)
	if err != nil {
		log.Error("Failed to parse routes configuration :(")
		return nil, err
	}

	return routes, nil
}

var httpClient = http.Client{} // TODO: SET UP THE CLIENT PROPERLY!!!

func SetUpRouteHandlers(
	appConfigs *configs.AppConfigs, env *configs.Environment, mux *http.ServeMux,
) error {
	routes, err := GetRoutes(appConfigs)
	if err != nil {
		log.Error("Failed to get routes due to: %s", err)
		return err
	}
	log.Debug("Routes: %s", routes)

	for _, route := range routes {
		log.Debug("Replacing variables with their values in a route: %s", route)
		err := configs.ReplaceEnvVarsInConfigs(route, env.Variables)
		if err != nil {
			return err
		}
		log.Debug("Replaced route: %s", route)
	}

	for _, route := range routes {
		log.Debug("Adding a handler for route: %s", route)
		addRouteHandler(route, mux)
	}

	return nil
}

func getBackendUrl(r *Route) string {
	return fmt.Sprintf("%s://%s:%s%s", r.Proto, r.Host, r.Port, r.BackendEndpoint)
}

func addRouteHandler(route *Route, mux *http.ServeMux) {
	mux.HandleFunc(
		route.Endpoint,
		func(w http.ResponseWriter, r *http.Request) {
			log.Fixme("IMPLEMENT REQUEST PREPOCESSING!!!")
			req, err := http.NewRequest(route.Method, getBackendUrl(route), r.Body)
			if err != nil {
				log.Error("Failed to create a request for the http client for the route: %s", route)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			// Copy headers
			for key, values := range req.Header {
				for _, value := range values {
					req.Header.Add(key, value)
				}
			}

			resp, err := httpClient.Do(req)
			if err != nil {
				log.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
			defer resp.Body.Close()

			log.Fixme("IMPLEMENT RESPONSE PREPROCESSING!!!")

			// Copy headers
			for key, values := range resp.Header {
				for _, value := range values {
					w.Header().Add(key, value)
				}
			}

			w.WriteHeader(resp.StatusCode)
			io.Copy(w, resp.Body)
		},
	)
}
