package routing

import (
	"encoding/json"
	"os"

	"github.com/nturbo1/reverse-proxy/internal/configs"
	"github.com/nturbo1/reverse-proxy/internal/log"
)

func parseRoutes(appConfigs *configs.AppConfigs) ([]*Route, error) {
	routesMaster, err := parseRoutesMasterFile(appConfigs)
	if err != nil {
		return nil, err
	}

	var allRoutes []*Route
	for _, filename := range routesMaster.Files {
		routes, err := parseRoutesFile(filename)
		if err != nil {
			return nil, err
		}
		allRoutes = append(allRoutes, routes...)
	}

	return allRoutes, nil
}

func parseRoutesFile(filename string) ([]*Route, error) {
	routesBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Error("Failed to read from a routes file: %s", filename)
		return nil, err
	}

	var routes []*Route
	err = json.Unmarshal(routesBytes, &routes)
	if err != nil {
		log.Error("Failed to json unmarshal '%s' file content bytes due to: %s", filename, err)
		return nil, err
	}

	return routes, nil
}

func parseRoutesMasterFile(appConfigs *configs.AppConfigs) (*Routes, error) {
	routesMasterBytes, err := os.ReadFile(appConfigs.RoutesMasterFile)
	if err != nil {
		log.Error(
			"Failed to read from the routes master file '%s' due to %s",
			appConfigs.RoutesMasterFile,
			err,
		)
		return nil, err
	}

	var routesMaster Routes
	err = json.Unmarshal(routesMasterBytes, &routesMaster)
	if err != nil {
		log.Error("Failed to json unmarshal the routes master file content bytes due to: %s", err)
		return nil, err
	}

	return &routesMaster, nil
}
