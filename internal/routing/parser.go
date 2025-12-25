package routing

import (
	"encoding/json"
	"os"
	"fmt"

	"github.com/nturbo1/reverse-proxy/internal/configs"
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
		fmt.Printf("Failed to read from a routes file: %s\n", filename)
		return nil, err
	}

	var routes []*Route
	err = json.Unmarshal(routesBytes, &routes)
	if err != nil {
		fmt.Printf("Failed to json unmarshal '%s' file content bytes due to: %s\n", filename, err)
		return nil, err
	}

	return routes, nil
}

func parseRoutesMasterFile(appConfigs *configs.AppConfigs) (*Routes, error) {
	routesMasterBytes, err := os.ReadFile(appConfigs.RoutesMasterFile)
	if err != nil {
		fmt.Printf(
			"Failed to read from the routes master file '%s' due to %s\n",
			appConfigs.RoutesMasterFile,
			err,
		)
		return nil, err
	}

	var routesMaster Routes
	err = json.Unmarshal(routesMasterBytes, &routesMaster)
	if err != nil {
		fmt.Printf("Failed to json unmarshal the routes master file content bytes due to: %s\n", err)
		return nil, err
	}

	return &routesMaster, nil
}
