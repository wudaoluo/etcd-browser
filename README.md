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


### 安装
    etcd v3 默认支持 开箱即用

    etcd v2 执行下面语句在linux中
    sed -i "s/v3/v2/g" etcdbrowser.js
    或者手动修改etcdbrowser.js ,5,6行 v3 改成v2

## Screen Shot
![etcd-browser Screen Shot](http://henszey.github.io/etcd-browser/images/etcdbrowser.png)