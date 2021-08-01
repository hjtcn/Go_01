#### 作业命题：按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。
参考Kratos：
```
|- api                  # api目录是对外保留的proto文件及生成的pb.go文件
|- cmd		        # 项目主干，main所在
|   |-- myapp
|      |--- main.go
|- configs		 # configs 为配置文件目录
| --db.toml					
|- internal              # 项目内部包
|   |--dao               # dao层，用户数据库、cache、MQ等资源访问
|   |--di	         # 依赖注入层 采用wire静态分析依赖
|      |--- wire.go      # wire 声明
|      |--- wire_gen.go  # go generate 生成的代码
|   |--model		 # model 层，用于声明业务结构体
|   |--server            # server层，用于初始化grpc和http serverå
|   |--service           # service层，用于业务逻辑处理
|- test                  # 测试资源层
```
