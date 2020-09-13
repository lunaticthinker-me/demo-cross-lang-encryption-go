# Encrypt/Decrypt Example

Encrypt/Decrypt Example for the Article [Cross Programming Language Encryption – CSharp vs Go, Part 2](https://lunaticthinker.me/index.php/cross-programming-language-encryption-–-csharp-vs-go,-part-2).

[![Go Version](https://img.shields.io/github/go-mod/go-version/lunaticthinker-me/demo-cross-lang-encryption-go)](https://img.shields.io/github/go-mod/go-version/lunaticthinker-me/demo-cross-lang-encryption-go)
[![TravisCI](https://travis-ci.org/lunaticthinker-me/demo-cross-lang-encryption-go.svg?branch=master)](https://travis-ci.org/lunaticthinker-me/demo-cross-lang-encryption-go)
[![Contributions welcome](https://img.shields.io/github/contributors/lunaticthinker-me/demo-cross-lang-encryption-go)](https://img.shields.io/github/contributors/lunaticthinker-me/demo-cross-lang-encryption-go)

[![Sonarcloud Status](https://sonarcloud.io/api/project_badges/measure?project=lunaticthinker-me_demo-cross-lang-encryption-go&metric=alert_status)](https://sonarcloud.io/dashboard?id=lunaticthinker-me_demo-cross-lang-encryption-go)
[![SonarCloud Coverage](https://sonarcloud.io/api/project_badges/measure?project=lunaticthinker-me_demo-cross-lang-encryption-go&metric=coverage)](https://sonarcloud.io/component_measures/metric/coverage/list?id=lunaticthinker-me_demo-cross-lang-encryption-go)
[![SonarCloud Bugs](https://sonarcloud.io/api/project_badges/measure?project=lunaticthinker-me_demo-cross-lang-encryption-go&metric=bugs)](https://sonarcloud.io/component_measures/metric/reliability_rating/list?id=lunaticthinker-me_demo-cross-lang-encryption-go)
[![SonarCloud Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=lunaticthinker-me_demo-cross-lang-encryption-go&metric=vulnerabilities)](https://sonarcloud.io/component_measures/metric/security_rating/list?id=lunaticthinker-me_demo-cross-lang-encryption-go)


[![Donate to this project using Patreon](https://img.shields.io/badge/patreon-donate-yellow.svg)](https://patreon.com/dragoscirjan)
[![Donate to this project using Paypal](https://img.shields.io/badge/paypal-donate-yellow.svg)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=QBP6DEBJDEMV2&source=url)

<!--[![Donate to this project using Flattr](https://img.shields.io/badge/flattr-donate-yellow.svg)](https://flattr.com/profile/balupton)
[![Donate to this project using Liberapay](https://img.shields.io/badge/liberapay-donate-yellow.svg)](https://liberapay.com/dragoscirjan)
[![Donate to this project using Thanks App](https://img.shields.io/badge/thanksapp-donate-yellow.svg)](https://givethanks.app/donate/npm/badges)
[![Donate to this project using Boost Lab](https://img.shields.io/badge/boostlab-donate-yellow.svg)](https://boost-lab.app/dragoscirjan/badges)
[![Donate to this project using Buy Me A Coffee](https://img.shields.io/badge/buy%20me%20a%20coffee-donate-yellow.svg)](https://buymeacoffee.com/balupton)
[![Donate to this project using Open Collective](https://img.shields.io/badge/open%20collective-donate-yellow.svg)](https://opencollective.com/dragoscirjan)
[![Donate to this project using Cryptocurrency](https://img.shields.io/badge/crypto-donate-yellow.svg)](https://dragoscirjan.me/crypto)
[![Donate to this project using Paypal](https://img.shields.io/badge/paypal-donate-yellow.svg)](https://dragoscirjan.me/paypal)
[![Buy an item on our wishlist for us](https://img.shields.io/badge/wishlist-donate-yellow.svg)](https://dragoscirjan.me/wishlist)
-->

- [Encrypt/Decrypt Example](#encryptdecrypt-example)
  - [Compatibility](#compatibility)
  - [Getting Started](#getting-started)
    - [Prereqiusites / Dependencies](#prereqiusites--dependencies)
    - [Installation](#installation)
    - [Development](#development)
      - [Requirements](#requirements)
        - [For Windows](#for-windows)
        - [For Linux/Unix/OSX](#for-linuxunixosx)
    - [Testing](#testing)
    - [Running](#running)
  - [Authors](#authors)
  - [Issues / Support](#issues--support)

<!-- /TOC -->

## Compatibility

| Algorithm / Language (Go) | C# | Go | Js | Py |
|---|---|---|---|---|
| AES/CFB | ? | ✓ | ✓ | ✕ |
| AES/CFB8 | ✕ | ✕ | ✕ | ✕ |
| AES/CBC | ✓ | ✓ | ✓ | ✓ |
| RSA/OAEP | ✕ | ✓ | ✕ | ✕ |
| RSA/PCKS1V15 | ✓ | ✓ | ✓ | ✓ |

## Getting Started

### Prereqiusites / Dependencies

Have [openssl]() installed.

### Installation

```bash
git clone https://github.com/lunaticthinker-me/demo-cross-lang-encryption-go
```

### Development

#### Requirements

- Please install [GoLang](https://golang.org/dl/). We support version 1.13.x and 1.14.x
- Please instal a GoLang IDE
  - [Visual Studio Code](https://code.visualstudio.com/) with [ITMCDev Go Extension Pack](https://marketplace.visualstudio.com/items?itemName=itmcdev.go-extension-pack)
  - [Jetbrains GoLand](https://www.jetbrains.com/go/)
  - [Vim](https://www.vim.org/) with [vim-go extension](https://github.com/fatih/vim-go)

##### For Windows

- Please install [git-scm](https://git-scm.com/download/win) tool.
- Please install a form of make/cmake
  - Install [Make for Windows](http://gnuwin32.sourceforge.net/packages/make.htm)
  - Install [make](https://sourceforge.net/projects/ezwinports/files/) from [ezwinports](https://sourceforge.net/projects/ezwinports/files/)
  - Install [chocolatey](https://chocolatey.org/), run `choco install make`
  <!-- - Install [Visual Studio Community](https://visualstudio.microsoft.com/vs/community/)
    - You will find it under `C:\Program Files (x86)\Microsoft Visual Studio\2019\Community\VC\Tools\MSVC\14.25.28610\bin\Hostx64` -->

##### For Linux/Unix/OSX

- Please install `git` and `make`

```bash
sudo apt-get install git make -y
```

### Testing

Run unit tests using `make test`.

### Running

Please run `make run`

Demo output:

```
AES Encrypted Values:
CFB 128 => 5WPWeBKWEafSfZCAscojoXjpr6AG78cC7Sqx52X9/fo=
CFB 192 => lkyhuJGvKOwOT5cKYJz9mmO6ND2PGo/XOM5mv5OIvYM=
CFB 256 => bmcX3+xKhz3Xml4/mQTL9qILe7SEIOfocERs4ZcqD74=
CBC 128 => 6u9RmbQs5XQQEIug+lP1+zRssBPfkQ5e0Y78TUbCtUE=
CBC 192 => hCNI0Yb90jKAhds4x9c4G0c5CwtRxMtCfe4As3JIq8A=
CBC 256 => dXP9pSWf6cgAegouT5UDTfkDE+t7A3j9khe7N/vNB00=

RSA Encrypted Values:
FxGi+JNXalIIL3Y+poyP4F3j9Mp4yR75Rbe7yx8yI3MNix95OI3LY6jBYpGD5nhXoaYKgX2NrmZcaAeNg7uzIH3m95ULMrboa0Br3IPmEw2aMwW8uxDEL/I4x7Uvlux1QCHnv3rnYNX/Hyipg3DMeKKppmcAYZ1zpfatH6qXMD0vGttpX1KksUe/3TN/oz8swPecAePFg6I/MPcndCxIeVfTXLqUCpQbxvmN7GYQpWbxXGB7S6rQpxNkZLcssH6XHwM/6LRQ3laQ+U+o3kL/bCUUrSB27B6yAB65I0SsLyhFY+bxDjugxOND0MPaVxVpa7MM5lileUL8uqG5U58sBg==

X509 Encrypted Values:
leFEibNhepKTLi2Er/BtavLejoKJ895LnkNgLEcolG4FOak7n/dKa1cYAZNrHqA/gu/Sd2HBdRirNI2OyJ766Lm5I5iiMQzLGobFI4+jyeLGDCjo2RjySLFBVjPKKbjo3RxcxpS6C/V3qvet3Px2VwdzCEfB3Ffpimsk6RblaZgLPl6YzRQsV1qkYtxxdsk3mdlR3eyXxSjfSUlx1bSZvv/BD2sCJtq+SiHOP9QfvQ8iIG5IbZtNdT95oaQ92bpGxuWA76DYqbP4C0s/Iv4w/GvV0mcgxcK1ePuY3wzCeHy6a80l36OaQGXf6xEJffvQ/QRz7BaEG2V0Fz/Ezx8L9g==
```

## Authors

- [Dragos Cirjan](mailto:dragos.cirjan@gmail.com) - Initial work - [Go Template](/lunaticthinker-me/demo-cross-lang-encryption-go)

See also the list of contributors who participated in this project.

## Issues / Support

Add a set of links to the [issues](/lunaticthinker-me/demo-cross-lang-encryption-go/issues) page/website, so people can know where to add issues/bugs or ask for support.

<!-- ## Changelog

Small changelog history. The rest should be added to [CHANGELOG.md](CHANGELOG.md).

See here a template for changelogs: https://keepachangelog.com/en/1.0.0/

Also see this tool for automatically generating them: https://www.npmjs.com/package/changelog -->
