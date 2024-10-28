package mockns1

import (
	"fmt"
	"net/http"

	"gopkg.in/ns1/ns1-go.v2/rest/model/alerting"
)

// Should be identical to rest.alertListResponse
type mockAlertListResponse struct {
	Limit        *int64            `json:"limit,omitempty"`
	Next         *string           `json:"next,omitempty"`
	Results      []*alerting.Alert `json:"results"`
	TotalResults *int64            `json:"total_results,omitempty"`
}

const alertPath = "../alerting/v1beta1/alerts"

// AddAlertListTestCase sets up a test case for the api.Client.Alert.List()
// function
func (s *Service) AddAlertListTestCase(
	params string, requestHeaders, responseHeaders http.Header,
	response []*alerting.Alert,
) error {
	length := int64(len(response))
	next := ""
	if length > 0 {
		next = *response[length-1].Name
	}
	listResponse := &mockAlertListResponse{

		Next:         &next,
		Results:      response,
		Limit:        &length,
		TotalResults: &length,
	}
	uri := alertPath
	if params != "" {
		uri = fmt.Sprintf("%s?%s", uri, params)
	}
	return s.AddTestCase(
		http.MethodGet, uri, http.StatusOK, requestHeaders,
		responseHeaders, "", listResponse,
	)
}

// AddAlertGetTestCase sets up a test case for the api.Client.Alerts.Get()
// function
func (s *Service) AddAlertGetTestCase(
	id string,
	requestHeaders, responseHeaders http.Header,
	response *alerting.Alert,
) error {
	return s.AddTestCase(
		http.MethodGet, fmt.Sprintf("%s/%s", alertPath, id), http.StatusOK, requestHeaders,
		responseHeaders, "", response,
	)
}

// AddAlertCreateTestCase sets up a test case for the api.Client.Alerts.Update()
// function
func (s *Service) AddAlertCreateTestCase(
	requestHeaders, responseHeaders http.Header,
	request alerting.Alert,
	response alerting.Alert,
) error {
	return s.AddTestCase(
		http.MethodPost, alertPath, http.StatusOK, requestHeaders,
		responseHeaders, request, response,
	)
}

// AddAlertUpdateTestCase sets up a test case for the api.Client.Alerts.Update()
// function
func (s *Service) AddAlertUpdateTestCase(
	requestHeaders, responseHeaders http.Header,
	request alerting.Alert,
	response alerting.Alert,
) error {
	return s.AddTestCase(
		http.MethodPatch, fmt.Sprintf("%s/%s", alertPath, *request.ID), http.StatusOK, requestHeaders,
		responseHeaders, request, response,
	)
}

// AddAlertReplaceTestCase sets up a test case for the api.Client.Alerts.Update()
// function
func (s *Service) AddAlertReplaceTestCase(
	requestHeaders, responseHeaders http.Header,
	request alerting.Alert,
	response alerting.Alert,
) error {
	return s.AddTestCase(
		http.MethodPut, fmt.Sprintf("%s/%s", alertPath, *request.ID), http.StatusOK, requestHeaders,
		responseHeaders, request, response,
	)
}

// AddAlertDeleteTestCase sets up a test case for the api.Client.Alerts.Delete()
// function
func (s *Service) AddAlertDeleteTestCase(
	id string,
	requestHeaders, responseHeaders http.Header,
) error {
	return s.AddTestCase(
		http.MethodDelete, fmt.Sprintf("%s/%s", alertPath, id), http.StatusNoContent, requestHeaders,
		responseHeaders, "", nil,
	)
}

// AddAlertTestPostTestCase sets up a test case for the api.Client.Alerts.Test()
// function
func (s *Service) AddAlertTestPostTestCase(
	id string,
	requestHeaders, responseHeaders http.Header,
) error {
	return s.AddTestCase(
		http.MethodPost, fmt.Sprintf("%s/%s/test", alertPath, id), http.StatusNoContent, requestHeaders,
		responseHeaders, "", nil,
	)
}

// AddAlertFailTestCase sets up a test case for the api.Client.Alerts.*()
// functions that fails.
func (s *Service) AddAlertFailTestCase(
	method string, id string, returnStatus int,
	requestHeaders, responseHeaders http.Header,
	responseBody string,
) error {
	path := alertPath
	if id != "" {
		path = fmt.Sprintf("%s/%s", alertPath, id)
	}
	return s.AddTestCase(
		method, path, returnStatus,
		nil, nil, "", responseBody)
}

func (s *Service) AddAlertFailTestCaseWithReqBody(
	method string, id string, returnStatus int,
	requestHeaders, responseHeaders http.Header,
	requestBody interface{},
	responseBody string,
) error {
	path := alertPath
	if id != "" {
		path = fmt.Sprintf("%s/%s", alertPath, id)
	}
	return s.AddTestCase(
		method, path, returnStatus,
		nil, nil, requestBody, responseBody)
}
