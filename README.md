# 登录
http://47.111.252.117:8080/server/login?account=admin&password=123456

# 登出
http://47.111.252.117:8080/server/logout?account=admin&password=123456

# 增加表单
http://47.111.252.117:8080/server/entryform?company=琴牌牛奶&jobs=收奶&working=张三&leader=李四&date=20201213&problem=是&type=标准不符&filename=a.png&score=11

# 修改表单
http://47.111.252.117:8080/server/entryform?id=5fc88010ffaff854e8d3779f&company=琴牌牛奶&jobs=收奶&working=张三&leader=李四&date=20201213&problem=是&type=标准不符&filename=a.png&score=11

# 查询表单
http://47.111.252.117:8080/server/queryform

# 删除表单
http://47.111.252.117:8080/server/deleteform?id=5fc88010ffaff854e8d3779f

# 检索
http://47.111.252.117:8080/server/search?leader=李四&date=202011

# 增加条款
http://47.111.252.117:8080/server/entryterms?terms="标准不符"&score=11

# 修改条款
http://47.111.252.117:8080/server/entryterms?id=5fc88010ffaff854e8d3779f&terms="标准不符"&score=12

# 查询条款
http://47.111.252.117:8080/server/queryterms

# 删除条款
http://47.111.252.117:8080/server/deleteterms?id=5fc88010ffaff854e8d3779f

# 注册账户
http://47.111.252.117:8080/server/register?account=admin&password=123456&level=1&group=admin

# 删除账户
http://47.111.252.117:8080/server/unregister?id=21232f297a57a5a743894a0e4a801fc3

# 密码修改
http://47.111.252.117:8080/server/modifypwd?account=admin&oldpassword=123123&newpassword=123456

#上传文件
curl -F "upload=@./a.png" -X POST http://47.111.252.117:8080/server/upload

#下载文件
http://47.111.252.117:8080/server/download?filename=a.png