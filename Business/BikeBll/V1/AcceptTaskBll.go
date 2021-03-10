package v1

import (
	"encoding/json"
	"log"
	acceptModel "redis-cache-api/Business/BikeBll/V1/Models"
	redis "redis-cache-api/RedisCache"
	"time"
)

// redisKey string
const redisKey = "bikeAcceptTask"

// GetTask func(taskID string) string
func GetTask(taskID string) bool {
	result := false
	value, err := redis.GetValue(redisKey)
	if err == nil {
		acceptList := []acceptModel.AcceptTaskRequest{}
		json.Unmarshal([]byte(value.(string)), &acceptList)
		taskIndex := findTask(acceptList, taskID)
		if taskIndex != -1 {
			result = true
			acceptList = removeTask(acceptList, taskIndex)
			redis.SetValue(redisKey, acceptList)
		}
	}
	return result
}

// AcceptTask func(acceptTask AcceptTaskRequest) bool
func AcceptTask(acceptTask acceptModel.AcceptTaskRequest) bool {
	result := false
	listAcceptTask := []acceptModel.AcceptTaskRequest{}
	clearTask(acceptTask.TaskID)
	value, err := redis.GetValue(redisKey)
	if err == nil {
		json.Unmarshal([]byte(value.(string)), &listAcceptTask)
	}
	taskIndex := findTask(listAcceptTask, acceptTask.TaskID)
	if taskIndex == -1 {
		result = true
		listAcceptTask = append(listAcceptTask, acceptTask)
		redis.SetValue(redisKey, listAcceptTask)
	}
	return result
}

// StampTask func(taskID string) bool
func StampTask(taskID string) (bool, bool) {
	result := false
	notFound := false
	_, err := redis.GetValue(taskID)
	if err != nil {
		listAcceptTask := []acceptModel.AcceptTaskRequest{}
		value, err := redis.GetValue(redisKey)
		if err == nil {
			json.Unmarshal([]byte(value.(string)), &listAcceptTask)
		}
		taskIndex := findTask(listAcceptTask, taskID)
		if taskIndex == -1 {
			notFound = true
		} else {
			result = true
			redis.SetValue(taskID, time.Now())
		}
	}
	return result, notFound
}

// FetchData func()
func FetchData() {
	t := time.Now()
	startDate, endDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC), time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.UTC)
	result := acceptModel.BikeResponseModel{}
	CallFetchDataAPI(startDate, endDate, &result)
	// Map data to redis cache
	listAcceptTask := []acceptModel.AcceptTaskRequest{}
	value, err := redis.GetValue(redisKey)
	if err == nil {
		json.Unmarshal([]byte(value.(string)), &listAcceptTask)
	}
	for _, element := range result.Result {
		log.Println(element.Task_id)
		taskIndex := findTask(listAcceptTask, element.Task_id)
		log.Println(taskIndex)
		if taskIndex == -1 {
			acceptTask := acceptModel.AcceptTaskRequest{}
			acceptTask.TaskID = element.Task_id
			listAcceptTask = append(listAcceptTask, acceptTask)
			log.Println(listAcceptTask)
			redis.SetValue(redisKey, listAcceptTask)
		}
	}
}

// VerifyTask func(taskID string) bool
func VerifyTask(taskID string) bool {
	result := false
	value, err := redis.GetValue(redisKey)
	if err == nil {
		acceptList := []acceptModel.AcceptTaskRequest{}
		json.Unmarshal([]byte(value.(string)), &acceptList)
		taskIndex := findTask(acceptList, taskID)
		if taskIndex != -1 {
			result = true
		}
	}
	return result
}

func clearTask(taskID string) {
	redis.ClearValue(taskID)
}

func findTask(acceptList []acceptModel.AcceptTaskRequest, taskID string) int {
	for i, acceptTask := range acceptList {
		if taskID == acceptTask.TaskID {
			return i
		}
	}
	return -1
}

func removeTask(acceptList []acceptModel.AcceptTaskRequest, index int) []acceptModel.AcceptTaskRequest {
	return append(acceptList[:index], acceptList[index+1:]...)
}
