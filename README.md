# gin-mvc
### 1. 项目结构
```
gin-mvc
├── app                   主要负责业务逻辑方面
│   ├── controller        控制层
│   ├── service           逻辑层
│   └── model             实体层
├── common
│   ├── auth              鉴权
│   ├── config            配置
│   ├── db                XORM引擎
│   ├── errors            自定义错误
│   ├── logger            日志组件
│   └── validator         参数校验
├── conf                  配置文件
├── router                路由信息
│   ├── middleware        路由中间件
│   └── router.go
├── script                与设置应用程序相关的脚本，例如，数据库脚本
├── test                  测试文件
├── go.mod                包管理
├── go.sum
├── main.go               这里是应用程序的起点，这个是所有程序的起点。
├── README.md
└── LICENSE
```

### 2. 通用组件

#### 2.1 logger 日志组件

- 日志组件，使用logrus作为基本日志库，对日志进行配置和封装；封装指定的日志格式，配置了日志级别，并根据日志级别将日志存入对应文件中，控制日志的切割和保留时间。
- TODO

#### 2.2 errors 自定义错误

- 自定义错误负责统一错误信息，封装了错误码和指定错误展示格式，每个错误码对应了指定的httpCode进行返回。
- TODO

#### 2.3 config 项目配置

- 配置组件，负责加载配置文件中的相关配置。通过toml配置文件=>项目定义的结构体。
```
conf
├── dev  开发环境配置-默认：
│   └── app.conf.toml
├── pre  预发环境配置：预发环境配置与线上环境配置相同但无流量打入
│   └── app.conf.toml
├── pro  线上环境配置：
│   └── app.conf.toml
└── test 测试环境配置：
   └── app.conf.toml
```

#### 2.4 auth 鉴权

- 使用jwt鉴权，负责生成和解析鉴权的token。加密了username和userid作为鉴权信息，返回对应的token，并控制了token的生效时间。

#### 2.5 validator 参数校验

- 对gin框架的参数绑定进行扩充，在参数绑定时支持了validate相关的校验规则。

#### 2.6 middleware 路由中间件
```
router
├── middleware
│   ├── auth.go
│   ├── bindKey.go
│   ├── log.go
│   ├── max_allowed.go
│   └── recovery.go
└── router.go
```

- router.go文件主要负责了url到处理函数的映射，通过router指定了每个url到具体业务的映射并对接口设置了全局中间件和局部中间件。
- middleware中间件中暂时包含了五种中间件。
  - 鉴权中间件，负责用户登录鉴权。
  - bindkey中间件，负责注入后续所需要的数据，现在包含了requestId、数据库会话的注入。
  - 日志中间件，负责指定gin框架的日志格式和日志文件，主要负责access_log日志部分。
  - max_allowed中间件，负责限制最大的请求数目。
  - recover中间件，负责捕获API异常，打印异常日志，并统一对外返回数据，对未知的错误进行打印堆栈的操作。



### 3. 业务逻辑

#### 3.1 controller 控制层

- base.go实现了通用方法，控制返回值格式统一、获取注入数据【登录信息、请求ID、数据库session等】。
- 其余go文件都是用作参数绑定，控制参数输入、获取注入数据、调用API指定业务 处理的service。

#### 3.2 service 逻辑层

- 实现了具体的业务逻辑，是用户请求接口后最终执行的逻辑部分。

#### 3.3 model 实体层

- 实现了数据库相关封装；对应的封装内容应该是针对业务实体的操作，而非单纯对某个表的增删改查。



### 4. 项目测试

#### 4.1 路由测试

- boot.go文件初始化了部分使用组件。运行路由测试文件时需要指定添加boot.go，此时无需再开启web服务器即可通过路由测试文件直接使用http方法请求指定接口进行测试。
- 路由测试文件编写时，如果相关路由包含数据库相关逻辑，需要在测试文件中调用model层保证数据库环境符合预期后再使用http方法请求指定接口进行测试。

#### 4.2 单元测试

- TODO

### 5. 打包部署

- TODO
