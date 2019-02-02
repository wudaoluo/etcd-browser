### webUI是etcd-browser
https://github.com/henszey/etcd-browser.git
#### api接口参考e3w
https://github.com/soyking/e3w.git
### etcd kv操作参考e3ch
https://github.com/soyking/e3ch.git

##### 最大亮点 同时支持etcdv2 v3版本 更多功能正在路上！！！

### version 0.1
    - 支持 etcd v2
    - 支持 etcd v3
    - 使用 go 代替了node
    - 支持添加多个etcd地址
    - 支持 配置文件动态更新
    - 支持 etcdv3 tls 加密
    - 支持 json toml 配置文件 

### version 0.2 (后悔药功能)
    - etcd v3记录每次操作key,value 和版本号
    - 添加leveldb (10万数据测试查询速度很快)
    - etcd v3 后悔药开发完成
    - etcd v3 备份功能开发完成
    - etcd v2 [后悔药功能开发]

### version 0.3
    - 支持认证
        - etcd  认证
        - 登录  认证
### 安装
    etcd v3 默认支持 开箱即用
    etcd v2 手动修改etcdbrowser.js ,5,6行 v3 改成v2
    配置文件修改成 etcd_version值改成 "v2"

### 生产tls证书
> https://www.cnblogs.com/Tempted/p/7737361.html

## Screen Shot
![etcd-browser Screen Shot](http://henszey.github.io/etcd-browser/images/etcdbrowser.png)