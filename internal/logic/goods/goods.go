package goods

import (
	"context"
	"github.com/gogf/gf-demo-user/v2/internal/dao"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/model/do"
	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

type (
	sGoods struct{}
)

func New() *sGoods {
	return &sGoods{}
}

//建议大家在编写完第一个logic方法后：
//1. 就在init函数中实现服务的注册；
//2. 然后实现去查看一下logic.go文件是否添加了相关的依赖，没有的话也可以手动添加一下
func init() {
	service.RegisterGoods(New())
}

//添加商品
func (s *sGoods) AddGoods(ctx context.Context, in model.GoodsAddInput) (err error) {
	_, err = dao.Goods.Ctx(ctx).Data(do.Goods{
		Name:        in.Name,
		Description: in.Description,
	}).Insert()
	return err
}

//查询商品
func (s *sGoods) GoodsInfo(ctx context.Context, in model.GoodsInfoInput) (goods *entity.Goods) {
	//注意：scan中要传入的是指针类型，不能直接传值类型！！！
	err := dao.Goods.Ctx(ctx).WherePri(in.Id).Scan(&goods)
	if err != nil {
		return nil
	}
	return
}
