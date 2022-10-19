package controller

import (
	"context"
	v1 "github.com/gogf/gf-demo-user/v2/api/v1"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/service"
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
	//这么写比较麻烦 todo 优化
	var goods = model.GoodsInfoOutput{}
	err, goods = service.Goods().GoodsInfo(ctx, model.GoodsInfoInput{
		Id: req.Id,
	})
	res.Goods = goods.Info
	if err != nil {
		return nil, err
	}
	return
}
