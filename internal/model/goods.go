package model

import "github.com/gogf/gf-demo-user/v2/internal/model/entity"

//service中的传参是在这里定义的
//我们可以在model目录下新建文件，这里是可以编辑的；
//但是model目录下生成的do和entity目录是无法编辑的，要通过gf cli工具生成
type GoodsInfoInput struct {
	Id int
}

//为什么在api中定义一遍，在model中还要重复定义一遍呢？？？
//原因是把业务模块和数据模块解耦 这个问题在复杂项目的后期是个很头疼的问题
//我们可以简单理解为，业务模块按需处理业务，对应工程目录中的controller和service
//而数据模块只处理数据，可以不关心业务是什么，按照自己处理数据最优的原则去处理数据就ok了。
type GoodsAddInput struct {
	Name        string
	Description string
}

type GoodsInfoOutput struct {
	Info *entity.Goods
}

/**
为什么要定义两遍数据结构呢?在api层定义了一遍？在model层有定义了一遍？
这个设计真的深得我心，我一下就get到了goframe团队的点。
我们之前在在开发商品中心统一入库是就遇到了难以维护的问题，原因就是业务逻辑和数据处理逻辑耦合在一起。
随着业务的复杂度越来越高，项目维护成本越来越高，甚至达到了难以维护的程度。
我们是如何解决的呢？
答案和goframe的api和model分别定义数据结构，进行业务和数据的分层，底层思想是一样的。
我们把复杂的逻辑进行了拆分：定义了业务模块（也就是api层）和数据处理模块（也就是model）。
业务模块：去处理负责的传参，并不需要关心如何入库和取值，按需去对接业务侧的需求就可以了。
数据模块：不需要关心业务测的具体实现，统一定义了入参标准，要求业务模块按照数据模块的要求，统一传入数据；数据模块考虑的重点是如何高效的批量插入数据，如何高效的按需取值，并不需要关心多变的业务侧需求。
*/
