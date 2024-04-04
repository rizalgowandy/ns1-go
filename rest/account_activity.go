package rest

import (
	"net/http"

	"gopkg.in/ns1/ns1-go.v2/rest/model/account"
)

// ActivityService handles 'account/activity' endpoint.
type ActivityService service

// List returns all activity in the account.
//
// NS1 API docs: https://developer.ibm.com/apis/catalog/ns1--ibm-ns1-connect-api/api/API--ns1--ibm-ns1-connect-api#getActivity
func (s *ActivityService) List() ([]*account.Activity, *http.Response, error) {
	// TODO: add support for url parameters to adjust endpoint behavior?
	req, err := s.client.NewRequest("GET", "account/activity", nil)
	if err != nil {
		return nil, nil, err
	}

	al := []*account.Activity{}
	resp, err := s.client.Do(req, &al)
	if err != nil {
		return nil, resp, err
	}

	return al, resp, nil
}
