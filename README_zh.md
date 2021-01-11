# CPChain DApps - 身份注册

身份注册 DApp 是 CPChain DApp 生态网络中的基础 DApp，此 DApp 用于 CPChain 用户注册他们的身份。身份是公私钥对，在 DApp 中，用户注册他们的公钥，因此，其它用户想要给此用户发送消息或文件时，通过 DApp 获取其公钥，加密后发送，只有此公钥对应的私钥才能打开。

用户也可使用自己的公钥进行签名，验证者通过 DApp 获取其公钥验证签名。

## 编译合约

```bash

docker run --rm -v `pwd`:/go/src/bitbucket.org/cpchain/chain/identity -it cpchain2018/abigen abigen --sol ./identity/identity.sol --pkg identity --out ./identity/identity.go

```

## 部署合约

```bash

build/main identity deploy --keystore ./dapps-admin/keystore/ --endpoint http://52.220.174.168:8501

```

合约地址：0x38ef6127a67C2d14FBa0a14cAEBe61Db093d3a4A
