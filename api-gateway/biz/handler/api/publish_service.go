// Code generated by hertz generator.

package api

import (
	"context"

	api "api-gateway/biz/model/api"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// PublishAction .
// @router /douyin/publish/action/ [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PublishActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.PublishActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetPublishList .
// @router /douyin/publish/list/ [GET]
func GetPublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.PublishListResponse)

	c.JSON(consts.StatusOK, resp)
}
