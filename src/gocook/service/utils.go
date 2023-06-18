package service

import (
	"context"
	"gocook/authorization"
)

func PrepareDecsisionRequest(ctx context.Context, ressource interface{}) (*authorization.DecisionRequest, error) {
	user, err := GetUserByID(ctx)
	if err != nil {
		return nil, err
	}
	method := ctx.Value("method")
	path := ctx.Value("path")
	decisionRequest := authorization.DecisionRequest{
		Method:    method.(string),
		Path:      path.([]string),
		User:      *user,
		Ressource: ressource,
	}
	return &decisionRequest, nil
}
