package lim

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type LIM struct {
	endpoint string
	client   *http.Client
}

func NewLIM(endpoint string, client *http.Client) (*LIM, error) {
	if client == nil {
		client = &http.Client{}
	}

	return &LIM{
		endpoint: endpoint,
		client:   client,
	}, nil
}

func (c *LIM) GetAuthSign(ctx context.Context, req *AuthSignReqModel) (*AuthSignResModel, int, error) {
	res, err := c.request(ctx, http.MethodPost, URLPathAuthSign, req)
	if err != nil {
		return nil, 0, err
	}

	if res.Code != 0 {
		return nil, res.Code, nil
	}

	var ret AuthSignResModel
	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return nil, res.Code, nil
	}

	return &ret, 0, nil
}

func (c *LIM) CreateUser(ctx context.Context, req *UserCreateReqModel) (*UserCreateResModel, int, error) {
	res, err := c.request(ctx, http.MethodPost, URLPathUserCreate, req)
	if err != nil {
		return nil, 0, err
	}

	if res.Code != 0 {
		return nil, res.Code, nil
	}

	var ret UserCreateResModel
	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return nil, res.Code, nil
	}

	return &ret, 0, nil
}
func (c *LIM) DeleteUser(ctx context.Context, req *UserDeleteReqModel) (*UserDeleteResModel, int, error) {
	res, err := c.request(ctx, http.MethodPost, URLPathUserDelete, req)
	if err != nil {
		return nil, 0, err
	}

	if res.Code != 0 {
		return nil, res.Code, nil
	}

	var ret UserDeleteResModel
	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return nil, res.Code, nil
	}

	return &ret, 0, nil
}
func (c *LIM) DetailUser(ctx context.Context, req *UserProfileReqModel) (*UserProfileResModel, int, error) {
	res, err := c.request(ctx, http.MethodPost, URLPathUserProfile, req)
	if err != nil {
		return nil, 0, err
	}

	if res.Code != 0 {
		return nil, res.Code, nil
	}

	var ret UserProfileResModel
	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return nil, res.Code, nil
	}

	return &ret, 0, nil
}

func (c *LIM) UpdateUser(ctx context.Context, req *UserUpdateReqModel) (*UserUpdateResModel, int, error) {
	res, err := c.request(ctx, http.MethodPost, URLPathUserProfileUpdate, req)
	if err != nil {
		return nil, 0, err
	}

	if res.Code != 0 {
		return nil, res.Code, nil
	}

	var ret UserUpdateResModel
	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return nil, res.Code, nil
	}

	return &ret, 0, nil
}

func (c *LIM) SendMessage(ctx context.Context, req *MessageSendReqModel) (*MessageSendResModel, int, error) {
	res, err := c.request(ctx, http.MethodPost, URLPathMessageSend, req)
	if err != nil {
		return nil, 0, err
	}

	if res.Code != 0 {
		return nil, res.Code, nil
	}

	var ret MessageSendResModel
	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return nil, res.Code, nil
	}

	return &ret, 0, nil
}

func (c *LIM) request(ctx context.Context, method string, url string, body any) (*ResponseModel, error) {
	v, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	url = c.endpoint + url
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(v))
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result ResponseModel
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
