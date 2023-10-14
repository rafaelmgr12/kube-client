package deployment

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/rafaelmgr12/kube-client/errors"
)

type Service struct {
	client  *http.Client
	urlBase string
}

func NewService(client *http.Client, url string) Service {
	return Service{
		client:  client,
		urlBase: url,
	}
}

func (s *Service) CreateDeployment(ctx context.Context, deployment Deployment) (*Deployment, error) {
	body, err := json.Marshal(deployment)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.urlBase+"/deployments", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := s.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return nil, errors.FromBadRequest(resp)
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.FromHTTPResponse(resp)
	}

	var createdDeployment Deployment
	if err := json.NewDecoder(resp.Body).Decode(&createdDeployment); err != nil {
		return nil, err
	}

	return &createdDeployment, nil
}

func (s *Service) DeleteDeployment(ctx context.Context, id string) error {

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, s.urlBase+"/deployments/"+id, nil)
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

func (s *Service) GetDeployment(ctx context.Context, id string) (*Deployment, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.urlBase+"/deployments/"+id, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.NewNotFound(uuid.MustParse(id), "deployment")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.FromHTTPResponse(resp)
	}

	deploy := Deployment{}
	if err := json.NewDecoder(resp.Body).Decode(&deploy); err != nil {
		return nil, err
	}

	return &deploy, nil

}
