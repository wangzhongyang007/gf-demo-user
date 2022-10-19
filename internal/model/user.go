package model

//service中的传参是在这里定义的
//我们可以在model目录下新建文件，这里是可以编辑的；但是model目录下生成的do和entity目录是无法编辑的，要通过gf cli工具生成
type UserCreateInput struct {
	Passport string
	Password string
	Nickname string
}

type UserSignInInput struct {
	Passport string
	Password string
}
