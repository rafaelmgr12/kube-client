package kube

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
	ierrors "github.com/rafaelmgr12/kube-client/kube/errors"
)

type BaseService struct {
	client  *http.Client
	urlBase string
}

func NewBaseService(client *http.Client, urlBase string) *BaseService {
	return &BaseService{
		client:  client,
		urlBase: urlBase,
	}
}

func (s *BaseService) DoRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, s.urlBase+endpoint, body)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func HandleErrors(resp *http.Response) error {
	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusNoContent:
		return nil
	case http.StatusBadRequest:
		return ierrors.FromBadRequest(resp)
	case http.StatusNotFound:
		var notFoundBody ierrors.NotFoundResponseBody
		body, _ := io.ReadAll(resp.Body)
		if err := json.Unmarshal(body, &notFoundBody); err != nil {
			// Se a deserialização falhar, retorne um erro genérico
			return ierrors.NewNotFound(uuid.Nil, "Resource")
		}
		return ierrors.NewNotFound(notFoundBody.ID, notFoundBody.Resource)
	default:
		return ierrors.FromHTTPResponse(resp)
	}
}
