package service

import (
	"GoCook/authorization"
	"GoCook/model"
	"context"
)

func PrepareDecsisionRequest(ctx context.Context, recipe *model.Recipe) (*authorization.DecisionRequest, error) {
	user, err := GetUserByID(ctx)
	if err != nil {
		return nil, err
	}
	method := ctx.Value("method")
	path := ctx.Value("path")
	decisionRequest := authorization.DecisionRequest{
		Method: method.(string),
		Path:   path.([]string),
		User:   *user,
		Recipe: *recipe,
	}
	return &decisionRequest, nil
}
