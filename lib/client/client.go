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

func (api NotionApi) postWithCookie(url string, body interface{}, cookie *http.Cookie) (*resty.Response, error) {
	client := resty.New()
	return client.R().SetCookie(&http.Cookie{
		Name:  "token_v2",
		Value: api.token,
	}).SetBody(body).SetCookie(cookie).Post(api.apiOrigin + url)
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

type SendTemporaryPasswordResult struct {
	CsrfState string
}

func (api NotionApi) SendTemporaryPassword(email string) (string, string, error) {
	requestBody := map[string]interface{}{
		"email":            email,
		"disableLoginLink": false,
		"native":           false,
		"isSignup":         false,
	}
	response, _ := api.post("sendTemporaryPassword", requestBody)
	var responseData SendTemporaryPasswordResult
	var csrf string
	json.Unmarshal(response.Body(), &responseData)
	for _, cookie := range response.Cookies() {
		if cookie.Name == "csrf" {
			csrf = cookie.Value
		}
	}
	return responseData.CsrfState, csrf, nil
}

func (api NotionApi) LoginWithEmail(state, csrf, password string) (string, error) {
	requestBody := map[string]interface{}{
		"state":    state,
		"password": password,
	}
	cookie := http.Cookie{
		Name:  "csrf",
		Value: csrf,
	}
	response, _ := api.postWithCookie("loginWithEmail", requestBody, &cookie)
	var token_v2 string
	for _, cookie := range response.Cookies() {
		if cookie.Name == "token_v2" {
			token_v2 = cookie.Value
		}
	}
	return token_v2, nil
}

type GetSpacesResult struct {
	RecordMap struct {
		Space map[string]struct {
			Value struct {
				Name string
			}
		} `json:"space"`
	} `json:"recordMap"`
}

func (api NotionApi) GetSpaces() (map[string]string, error) {
	response, _ := api.post("loadUserContent", nil)
	var responseData GetSpacesResult
	json.Unmarshal(response.Body(), &responseData)
	spaces := make(map[string]string, len(responseData.RecordMap.Space))
	for key, val := range responseData.RecordMap.Space {
		spaces[key] = val.Value.Name
	}
	return spaces, nil
}
