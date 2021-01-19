# CPChain DApps - 身份注册

身份注册 DApp 是 CPChain DApp 生态网络中的基础 DApp，此 DApp 用于 CPChain 用户注册他们的身份。身份是公私钥对，在 DApp 中，用户注册他们的公钥，因此，其它用户想要给此用户发送消息或文件时，通过 DApp 获取其公钥，加密后发送，只有此公钥对应的私钥才能打开。

用户也可使用自己的公钥进行签名，验证者通过 DApp 获取其公钥验证签名。

## 编译合约

```bash

docker run --rm -v `pwd`:/go/src/bitbucket.org/cpchain/chain/identity -it cpchain2018/abigen abigen --sol ./identity/identity.sol --pkg identity --out ./identity/identity.go

```

## 部署合约

```bash

make build

build/main identity deploy --keystore ./dapps-admin/keystore/ --endpoint http://52.220.174.168:8501

```

合约地址： 0xC53367856164DA3De57784E0c96710088DA77e20

## 方法简介

+ `register(identity: string)` : 注册身份（重新注册将覆盖之前的）
+ `remove()` : 移除身份
+ `count()` : 当前身份个数
+ `get(address: address)` : 根据钱包地址获取用户身份

身份需采用 JSON 进行序列化，字段为：

```json

{
    "pub_key": "", // 使用 Base64 进行序列化
    "name": "", // 用户名称
    "version": "1.0"
}

```

+ *version 用于客户端消息处理方式的版本控制*

## FAQ

### 为什么不使用用户自己的钱包公私钥

为了安全和可替换，尽量让钱包的私钥仅用于发送交易。
