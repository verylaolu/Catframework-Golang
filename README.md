# golang framework jwt


### 项目简介
    golang Catframework ( jwt   mysql   redis )
    GOLANG开发的超简洁框架
    包含了JWT TOKEN ,MYSQL ,REDIS的验证规则示例
    可以作为DEMO学习， 也可以作为简单的框架进行基础开发
    不包含模板引擎相关

### 安装方法
    自己装golang环境
    下载就一个main.go
### 使用方法
    go run main.go
    http://127.0.0.1:8888

### 用户信息验证代码(※※必须按照自己项目修改※※)   
        
### 验证权限规则
    
  

### 测试代码
   

### 模块路由
        r.HandleFunc("/", Middleware.Auth(Controller.Home))
	r.HandleFunc("/jwt/encode",Middleware.Auth(Controller.Encodetoken))
	r.HandleFunc("/jwt/decode", Middleware.Auth(Controller.Decodetoken))

### 数据库
    

### 返回格式


数据返回格式(直接返回数据内容)：
  
    {
    	"code": 200,
        "msg": "success",
        "data": {
            "id": 2
        }
    }



### httpcode
    - 200 GET 操作成功
    - 201 POST 创建成功
    - 205 PUT 修改成功
    - 200 DELETE 删除成功
    - 204 DELETE 删除失败
    - 400 参数校验错误
    - 401 未授权
    - 500 服务器处理错误

