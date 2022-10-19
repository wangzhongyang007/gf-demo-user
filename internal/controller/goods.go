package controller

import (
	"context"
	v1 "github.com/gogf/gf-demo-user/v2/api/v1"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

var Goods = cGoods{}

type cGoods struct{}

//添加商品
func (c *cGoods) AddGoods(ctx context.Context, req *v1.GoodsAddReq) (res *v1.GoodsAddRes, err error) {
	//和internal内部，和model层同级的有service层。
	err = service.Goods().AddGoods(ctx, model.GoodsAddInput{
		Name:        req.Name,
		Description: req.Description,
	})
	return
}

//查询商品
func (c *cGoods) GoodsInfo(ctx context.Context, req *v1.GoodsInfoReq) (res *v1.GoodsInfoRes, err error) {
	res = &v1.GoodsInfoRes{
		Goods: service.Goods().GoodsInfo(ctx, model.GoodsInfoInput{
			Id: req.Id,
		}),
	}
	g.Dump(res)
	return
}
