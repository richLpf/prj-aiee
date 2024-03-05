## 待补充
- 全局go-cli配置
- docker-compose启动项目实例
- mysql脚本自动创建数据库表
- 字段规范，大小写、下划线
- swagger docs
  
## 使用过程中的问题
- 支持go-cli给多个项目生成api
- 字段类型支持longtext兼容问题
- 当我创建了多个表时，生成对应的API，将我之前的代码覆盖了，需要指定对应的表来生成
- 数据库连接方式可配置
- 支持不同的数据库
- 竞品调研，同类型工具
- 轻量，适合上手，企业业务开发
插入角色的时候，如何同时关联权限，在事务里面，此时角色的ID还没有生成

## 启动数据库
```
docker run -d \
  --name test-db \
  -e MYSQL_ROOT_PASSWORD=12345678 \
  -e MYSQL_DATABASE=aiee \
  -e MYSQL_USER=user \
  -e MYSQL_PASSWORD=12345678 \
  -p 3306:3306 \
  mariadb
```

### swag

GOPATH = go env
GOPATH = /Users/user/go

> export PATH=$PATH:<swag_install_path>

export PATH=$PATH:/Users/user/go/bin

source ~/.bashrc

swag init