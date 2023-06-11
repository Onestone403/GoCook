package authorization

import (
	"GoCook/model"
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

type decisionResult struct {
	Allow bool `json:"allow"`
}

type decision struct {
	Result     *decisionResult `json:"result"`
	DecisionId string          `json:"decision_id"`
}

type DecisionRequest struct {
	Method string       `json:"method"`
	Path   []string     `json:"path"`
	User   model.User   `json:"user"`
	Recipe model.Recipe `json:"recipe"`
}

type decisionReqInternal struct {
	Input *DecisionRequest
}

type Client interface {
	IsAllowed(dreq *DecisionRequest) (bool, error)
}

type client struct {
	restClient *resty.Client
	endpoint   string
}

type config struct {
	restClient *resty.Client
	endpoint   string
}

func New() Client {
	return &client{
		restClient: resty.New(),
		endpoint:   "http://127.0.0.1:8181/v1/data/authz",
	}
}

func (c *client) IsAllowed(dreq *DecisionRequest) (bool, error) {
	dreqStr, err := json.Marshal(&decisionReqInternal{Input: dreq})
	if err != nil {
		return false, err
	}
	resp, err := c.restClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(dreqStr).
		SetResult(&decision{}).
		Post(c.endpoint)
	if err != nil || resp.IsError() {
		return false, err
	}
	return resp.Result().(*decision).Result.Allow, nil
}
