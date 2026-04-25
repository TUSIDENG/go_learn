# 引入外部本地包
mod文件关键配置：
```
# 这里需要用包全名
replace go_learn/package/custom_package/01 => ../01

# 然后执行go mod tidy出现这个，引入外部本地包
require go_learn/package/custom_package/01 v0.0.0-00010101000000-000000000000
```