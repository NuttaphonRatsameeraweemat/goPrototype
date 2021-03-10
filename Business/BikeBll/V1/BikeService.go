package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	env "redis-cache-api/EnvironmentVariable"
)

// CallFetchDataAPI func(startDate Time, endDate Time)
func CallFetchDataAPI(startDate time.Time, endDate time.Time, target interface{}) {
	requestBody, _ := json.Marshal(map[string]string{
		"startDate": startDate.Format("2006-01-02"),
		"endDate":   endDate.Format("2006-01-02"),
	})
	client := http.Client{}
	request, _ := http.NewRequest("POST", env.GetBikeAPIURL()+"v2/intergration/GetTasksBroadcastList", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	json.NewDecoder(response.Body).Decode(target)
}
