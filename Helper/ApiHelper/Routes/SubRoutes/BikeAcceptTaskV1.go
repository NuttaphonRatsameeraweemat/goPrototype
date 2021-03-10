package subroutes

import (
	acceptControl "redis-cache-api/Controllers/BikeControllers/V1"
	routeModel "redis-cache-api/Helper/ApiHelper/Routes/Models"
)

// GetAcceptTaskV1Routes func() []Route
func GetAcceptTaskV1Routes() []routeModel.Route {
	routes := []routeModel.Route{}
	routes = append(routes, routeModel.Route{
		Name:        "GetAcceptTask",
		Method:      "GET",
		Pattern:     "/api/v1/bike/acceptTask/{taskId}",
		HandlerFunc: acceptControl.GetAcceptTask,
	})
	routes = append(routes, routeModel.Route{
		Name:        "PostAcceptTask",
		Method:      "POST",
		Pattern:     "/api/v1/bike/acceptTask",
		HandlerFunc: acceptControl.PostAcceptTask,
	})
	routes = append(routes, routeModel.Route{
		Name:        "FetchData",
		Method:      "POST",
		Pattern:     "/api/v1/bike/acceptTask/fetchData",
		HandlerFunc: acceptControl.FetchData,
	})
	routes = append(routes, routeModel.Route{
		Name:        "VerifyTask",
		Method:      "GET",
		Pattern:     "/api/v1/bike/acceptTask/verifyTask/{taskId}",
		HandlerFunc: acceptControl.VerifyTask,
	})
	return routes
}
