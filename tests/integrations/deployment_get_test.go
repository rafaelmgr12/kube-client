package integration

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/rafaelmgr12/kube-client/deployment"
	"github.com/rafaelmgr12/kube-client/kube"
	ierrors "github.com/rafaelmgr12/kube-client/kube/errors"
)

func TestGetDeployment(t *testing.T) {
	c, err := kube.NewClient(
		kube.WithURL(url),
	)
	if err != nil {
		t.Errorf("should not fail to create client: %s", err)
		return
	}

	deploy := deployment.Deployment{
		ID:       uuid.New(),
		Image:    "nginx",
		Replicas: 1,
		Ports: []deployment.Port{
			{
				Name:   "http",
				Number: 80,
			},
		},
	}

	_, err = c.Deployment.CreateDeployment(context.Background(), deploy)
	if err != nil {
		t.Errorf("should not fail to create deployment: %s", err)
		return
	}
	defer c.Deployment.DeleteDeployment(context.Background(), deploy.ID.String())

	foundDeploy, err := c.Deployment.GetDeployment(context.Background(), deploy.ID.String())
	if err != nil {
		t.Errorf("should not fail to get deployment: %s", err)
		return
	}

	assertDeployment(t, &deploy, foundDeploy)
}

func TestGetNotFoundDeployment(t *testing.T) {
	c, err := kube.NewClient(
		kube.WithURL(url),
	)
	if err != nil {
		t.Errorf("should not fail to create client: %s", err)
		return
	}

	foundDeploy, err := c.Deployment.GetDeployment(context.Background(), uuid.NewString())
	if err == nil || foundDeploy != nil {
		t.Errorf("should fail to get non existent deployment")
		return
	}

	if _, ok := err.(ierrors.NotFoundResource); !ok {
		t.Errorf("should fail with ErrNotFound")
		return
	}
}
