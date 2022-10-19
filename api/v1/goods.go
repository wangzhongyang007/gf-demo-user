package v1

import (
	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

//查询单个商品
type GoodsInfoReq struct {
	//疑问：path是什么？路由吗？这么写路由不好管理呀？
	g.Meta `path:"/goods/info" method:"get" tags:"GoodsService" summary:"Get the info of current goods"`
	Id     int `v:"required"`
}
type GoodsInfoRes struct {
	*entity.Goods
}

//添加商品
type GoodsAddReq struct {
	g.Meta      `path:"/goods/add" method:"post" tags:"GoodsService" summary:"添加商品"`
	Name        string `v:"required|length:1,20"`
	Description string `v:"required|length:1,50"`
}
type GoodsAddRes struct{}
