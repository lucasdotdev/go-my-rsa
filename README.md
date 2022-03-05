<h1 align="center">🔐 Go My RSA</h1>

<p align="center"><img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/lucasdotdev/go-my-rsa">
  <a href="blob/master/LICENSE"><img alt="GitHub" src="https://img.shields.io/github/license/lucasdotdev/go-my-rsa"></a></p>

This is a very simple Go implementation of the RSA algorithm that I've made while studying this awesome language. Feel free to check the code and do suggestions if you find something that can be improved. ❤️

## 🏗️ Building

You can build it as a regular Go program, there are no dependencies:

```shell
go build
```

## ⚙️ Usage

There are basically three commands to check here: `generate-profile`, `encrypt` and `decrypt`.

### 🔑 Generating a Profile

You have to start creating a profile for encryption:

```
./my-rsa generate-profile --size 4
```

It will print a key with 6 pieces used to encrypt and decrypt messages.

### 🔒 Encrypting a Message

With a profile in hands, you can encrypt messages by calling the command `encrypt`:

```
./my-rsa encrypt "A simple message" --profile xxx-xxx-xxx-xxx-xxx-xxx
```

The result will be an encrypted message that can only be decrypted with the profile string.

### 🔓 Decrypting a Message

The decryption process is identical to the previous, so you just have to pass the encrypted message instead of the plain text:

```
./my-rsa encrypt "d;´129u /sas/g34" --profile xxx-xxx-xxx-xxx-xxx-xxx
```

It will print your message back!

## 📃 License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.
