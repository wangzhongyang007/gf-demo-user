# A simple user API service

## Installation

### 1. You need a go development environment setup before everything starts taking off.

### 2. Use `git clone` to clone the repo to your local folder.

```
https://github.com/wangzhongyang007/gf-demo-user
```

### 3. Import `manifest/sql/create.sql` to your database.

Import `manifest/sql/goods.sql` to your database.

注意：记得修改config.yaml中的数据库密码

### 4. Update `manifest/config/config.yaml` according to your local configurations if necessary.

### 5. Run command `go run main.go`, and you'll see something as follows if success:

```
  ADDRESS | METHOD |        ROUTE         |                                  HANDLER                                   |         MIDDLEWARE           
----------|--------|----------------------|----------------------------------------------------------------------------|------------------------------
  :8199   | ALL    | /*                   | github.com/gogf/gf/v2/net/ghttp.internalMiddlewareServerTracing            | GLOBAL MIDDLEWARE            
----------|--------|----------------------|----------------------------------------------------------------------------|------------------------------
  :8199   | ALL    | /*                   | github.com/gogf/gf/v2/net/ghttp.MiddlewareHandlerResponse                  | GLOBAL MIDDLEWARE            
----------|--------|----------------------|----------------------------------------------------------------------------|------------------------------
  :8199   | ALL    | /api.json            | github.com/gogf/gf/v2/net/ghttp.(*Server).openapiSpec                      |                              
----------|--------|----------------------|----------------------------------------------------------------------------|------------------------------
  :8199   | ALL    | /swagger/*           | github.com/gogf/gf/v2/net/ghttp.(*Server).swaggerUI                        | HOOK_BEFORE_SERVE            
----------|--------|----------------------|----------------------------------------------------------------------------|------------------------------
  :8199   | POST   | /user/check-passport | github.com/gogf/gf-demo-user/v2/internal/controller.(*cUser).CheckPassport | service.(*sMiddleware).Ctx   
          |        |                      |                                                                            | service.(*sMiddleware).CORS  
----------|--------|----------------------|----------------------------------------------------------------------------|------------------------------
  :8199   | POST   | /user/is-signed-in   | github.com/gogf/gf-demo-user/v2/internal/controller.(*cUser).IsSignedIn    | service.(*sMiddleware).Ctx   
          |        |                      |                                                                            | service.(*sMiddleware).CORS  
----------|--------|----------------------|----------------------------------------------------------------------------|------------------------------
  :8199   | GET    | /user/profile        | github.com/gogf/gf-demo-user/v2/internal/controller.(*cUser).Profile       | service.(*sMiddleware).Ctx   
          |        |                      |                                                                            | service.(*sMiddleware).CORS  
          |        |                      |                                                                            | service.(*sMiddleware).Auth  
----------|--------|----------------------|----------------------------------------------------------------------------|------------------------------
  :8199   | POST   | /user/sign-in        | github.com/gogf/gf-demo-user/v2/internal/controller.(*cUser).SignIn        | service.(*sMiddleware).Ctx   
          |        |                      |                                                                            | service.(*sMiddleware).CORS  
----------|--------|----------------------|----------------------------------------------------------------------------|------------------------------
  :8199   | POST   | /user/sign-out       | github.com/gogf/gf-demo-user/v2/internal/controller.(*cUser).SignOut       | service.(*sMiddleware).Ctx   
          |        |                      |                                                                            | service.(*sMiddleware).CORS  
----------|--------|----------------------|----------------------------------------------------------------------------|------------------------------
  :8199   | POST   | /user/sign-up        | github.com/gogf/gf-demo-user/v2/internal/controller.(*cUser).SignUp        | service.(*sMiddleware).Ctx   
          |        |                      |                                                                            | service.(*sMiddleware).CORS  
----------|--------|----------------------|----------------------------------------------------------------------------|------------------------------
```

# GoFrame Sites

### GoFrame Repo

* [https://github.com/gogf/gf](https://github.com/gogf/gf)

### GoFrame Home

* [https://goframe.org](https://goframe.org) (中文)

# 联系我

## 微信：wangzhongyang1993

## 公众号：程序员升级打怪之旅

## 我学习GoFrame的专栏：

https://juejin.cn/column/7146133427706658847