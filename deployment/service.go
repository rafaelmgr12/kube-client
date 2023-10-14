package deployment

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/rafaelmgr12/kube-client/errors"
)

type Service struct {
	client  *http.Client
	urlBase string
}

func NewService(client *http.Client, urlBase string) *Service {
	return &Service{
		client:  client,
		urlBase: urlBase,
	}
}

func (s *Service) CreateDeployment(deployment *Deployment) (*Deployment, error) {
	body, err := json.Marshal(deployment)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Post(s.urlBase+"/deployments", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.FromHTTPResponse(resp)
	}

	var createdDeployment Deployment
	if err := json.NewDecoder(resp.Body).Decode(&createdDeployment); err != nil {
		return nil, err
	}

	return &createdDeployment, nil
}

func (s *Service) DeleteDeployment(id string) error {
	req, err := http.NewRequest(http.MethodDelete, s.urlBase+"/deployments/"+id, nil)
	if err != nil {
		return err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return errors.FromHTTPResponse(resp)
	}

	return nil
}

func (s *Service) GetDeployment(id string) (*Deployment, error) {
	resp, err := s.client.Get(s.urlBase + "/deployments/" + id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.FromHTTPResponse(resp)
	}

	var deployment Deployment
	if err := json.NewDecoder(resp.Body).Decode(&deployment); err != nil {
		return nil, err
	}

	return &deployment, nil
}
