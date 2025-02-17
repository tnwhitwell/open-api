package porcelain

import (
	"github.com/tnwhitwell/open-api/go/models"
	"github.com/tnwhitwell/open-api/go/plumbing/operations"
	"github.com/tnwhitwell/open-api/go/porcelain/context"
)

func (n *Netlify) CreateDeployKey(ctx context.Context) (*models.DeployKey, error) {
	authInfo := context.GetAuthInfo(ctx)
	params := operations.NewCreateDeployKeyParams()
	resp, err := n.Netlify.Operations.CreateDeployKey(params, authInfo)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}
