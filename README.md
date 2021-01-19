# CPChain DApp - Identity

Identity DApp is the fundamental DApp of the CPChain DApp Ecosystem. This DApp is used for registering user's public-key and other customize information.

A user generates a pair named private-key and public-key. Then register public-key in the DApp. When someone wants to send a message or file to him, they can get the public-key in Identity DApp, then encrypt the message or file. Finally, send an encrypted message or file to him. The user decrypts a message or file with his private key.

A user also can sign a message or a file with the private-key, then others validate it with the public key got from DApp.

## Compile Contract

```bash

docker run --rm -v `pwd`:/go/src/bitbucket.org/cpchain/chain/identity -it cpchain2018/abigen abigen --sol ./identity/identity.sol --pkg identity --out ./identity/identity.go

```

## Deploy

```bash

make build

build/main identity deploy --keystore ./dapps-admin/keystore/ --endpoint http://52.220.174.168:8501

```

Contract address： 0xC53367856164DA3De57784E0c96710088DA77e20

## Methods

+ `register(identity: string)` : register identity of the sender（register again when override the original）
+ `remove()` : remove the identity of sender
+ `count()` : count of identies
+ `get(address: address)` : get identity by the address

Identity need be encoded in JSON, required fields as below:

```json

{
    "pub_key": "", // marshal with Base64
    "name": "", // username
    "version": "1.0"
}

```

+ *version is for the parser on the client*
