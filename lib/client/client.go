package client

import (
	"encoding/json"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type NotionApi struct {
	apiOrigin string
	token     string
}

func New(token string) NotionApi {
	api := NotionApi{"https://www.notion.so/api/v3/", token}
	return api
}

func (api NotionApi) post(url string, body interface{}) (*resty.Response, error) {
	client := resty.New()
	return client.R().SetCookie(&http.Cookie{
		Name:  "token_v2",
		Value: api.token,
	}).SetBody(body).Post(api.apiOrigin + url)
}

func (api NotionApi) ExportSpace(spaceId, exportType string) (string, error) {
	requestBody := map[string]interface{}{
		"task": map[string]interface{}{
			"eventName": "exportSpace",
			"request": map[string]interface{}{
				"spaceId": spaceId,
				"exportOptions": map[string]interface{}{
					"exportType": exportType,
					"timeZone":   "Europe/Prague",
					"locale":     "en",
				},
			},
		},
	}
	response, _ := api.post("enqueueTask", requestBody)
	var responseData map[string]string
	json.Unmarshal(response.Body(), &responseData)
	return responseData["taskId"], nil
}

type GetTasksExportSpaceResponse struct {
	Results []GetTasksExportSpaceResponseResult
}

type GetTasksExportSpaceResponseResult struct {
	ID     string
	Status GetTasksExportSpaceResponseResultStatus
}

type GetTasksExportSpaceResponseResultStatus struct {
	Type          string
	PagesExported int
	ExportURL     string
}

func (api NotionApi) GetTasksExportSpace(taskID string) (bool, int, string, error) {
	requestBody := map[string]interface{}{
		"taskIds": []string{taskID},
	}
	response, _ := api.post("getTasks", requestBody)
	var responseData GetTasksExportSpaceResponse
	json.Unmarshal(response.Body(), &responseData)
	return responseData.Results[0].Status.Type == "complete",
		responseData.Results[0].Status.PagesExported,
		responseData.Results[0].Status.ExportURL,
		nil
}
