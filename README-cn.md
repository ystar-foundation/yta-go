用 Go 语言与 YTA.IO 交互的 API 库
=========================

[![GoDoc](https://godoc.org/github.com/ystar-foundation/yta-go?status.svg)](https://godoc.org/github.com/ystar-foundation/yta-go)

该库提供对数据架构（二进制打包和JSON接口）的简单访问，
以及对远程或本地运行的EOS.IO RPC服务器的API调用。 
它提供钱包功能（KeyBag），或者可以通过 `keosd` 钱包签署交易。 
它还明白端口9876上的P2P协议。

截至6月的发布之前，这个库不断的在变化。 先不要期望稳定性，
因为我们要追着主网 `eosio` 代码库的脚步，而它的变化又那么快。

该库主网启动编排工具是 `eosio` 的基础，网址：
https://github.com/ystar-foundation/yta-bios


基本用法
-----------

```go
api := yta.New("http://testnet1.yta.io")

infoResp, _ := api.GetInfo()
accountResp, _ := api.GetAccount("initn")
fmt.Println("Permission for initn:", accountResp.Permissions[0].RequiredAuth.Keys)
```

`eosio.system` 和 `eosio.token` 的 _Actions_ 合约分别在:
* https://github.com/ystar-foundation/yta-go/tree/master/system ([godocs](https://godoc.org/github.com/ystar-foundation/yta-go/system))
* https://github.com/ystar-foundation/yta-go/tree/master/token ([godocs](https://godoc.org/github.com/ystar-foundation/yta-go/token))

范例
-------

看看库的用法的例子：

* https://github.com/ystar-foundation/yta-bios/blob/master/bios/bios.go
* https://github.com/ystar-foundation/yta-bios/blob/master/bios/ops.go
* `cmd/` 下还有一些其他的 `main` 工具包。


召集开源贡献者
------------

我们欢迎所有的开源贡献，直接用 GitHub-fu来提议、帮我们改进吧。


证书
-------

MIT
