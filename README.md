# GoLang Project

<img alt="GoLang logo" src="https://github.com/templ-project/go/blob/master/Go-Logo_Blue.png?raw=true" width="20%" align="right" />

> **go** is a template project, designed by [Templ Project](http://templ-project.github.io).
>
> **go** includes instructions for initializing a new [Go Lang](https://golang.org) project, and configuring it for
> development, unit testing as well as code linting and analysis.
>
> **go** implements:
>
> - [gocyclo](https://github.com/fzipp/gocyclo), [go-critic](https://github.com/go-critic/go-critic) for code analisys
> - [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports), [gofmt](https://golang.org/cmd/gofmt) for code formatting
> - [golangci-lint](https://github.com/golangci/golangci-lint), [golint](https://github.com/golang/lint) for linting
>
> **To use the template**, delete the content above, and fill in the following sections of this readme.

<!-- > TODO: -->

Set of shield/badges explaining where to find more information about the project (i.e. Where to look for unit test reports, where to see code coverage and code scans, etc.). You can find a lot of them on https://shields.io/)

[![Npm Version](https://img.shields.io/github/go-mod/go-version/templ-project/go)](https://img.shields.io/github/go-mod/go-version/templ-project/go)
[![TravisCI](https://travis-ci.org/templ-project/go.svg?branch=master)](https://travis-ci.org/templ-project/go)
[![Contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/templ-project/go/issues)

<!-- [![CircleCI](https://circleci.com/gh/templ-project/go.svg?style=shield)](https://circleci.com/gh/templ-project/go) -->

<!-- [![Sonarcloud Status](https://sonarcloud.io/api/project_badges/measure?project=templ-project_go&metric=alert_status)](https://sonarcloud.io/dashboard?id=templ-project_go)
[![SonarCloud Coverage](https://sonarcloud.io/api/project_badges/measure?project=templ-project_go&metric=coverage)](https://sonarcloud.io/component_measures/metric/coverage/list?id=templ-project_go)
[![SonarCloud Bugs](https://sonarcloud.io/api/project_badges/measure?project=templ-project_go&metric=bugs)](https://sonarcloud.io/component_measures/metric/reliability_rating/list?id=templ-project_go)
[![SonarCloud Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=templ-project_go&metric=vulnerabilities)](https://sonarcloud.io/component_measures/metric/security_rating/list?id=templ-project_go) -->

<!--
[![Donate to this project using Patreon](https://img.shields.io/badge/patreon-donate-yellow.svg)](https://patreon.com/dragoscirjan)
[![Donate to this project using Paypal](https://img.shields.io/badge/paypal-donate-yellow.svg)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=UMMN8JPLVAUR4&source=url)
[![Donate to this project using Flattr](https://img.shields.io/badge/flattr-donate-yellow.svg)](https://flattr.com/profile/balupton)
[![Donate to this project using Liberapay](https://img.shields.io/badge/liberapay-donate-yellow.svg)](https://liberapay.com/dragoscirjan)
[![Donate to this project using Thanks App](https://img.shields.io/badge/thanksapp-donate-yellow.svg)](https://givethanks.app/donate/npm/badges)
[![Donate to this project using Boost Lab](https://img.shields.io/badge/boostlab-donate-yellow.svg)](https://boost-lab.app/dragoscirjan/badges)
[![Donate to this project using Buy Me A Coffee](https://img.shields.io/badge/buy%20me%20a%20coffee-donate-yellow.svg)](https://buymeacoffee.com/balupton)
[![Donate to this project using Open Collective](https://img.shields.io/badge/open%20collective-donate-yellow.svg)](https://opencollective.com/dragoscirjan)
[![Donate to this project using Cryptocurrency](https://img.shields.io/badge/crypto-donate-yellow.svg)](https://dragoscirjan.me/crypto)
[![Donate to this project using Paypal](https://img.shields.io/badge/paypal-donate-yellow.svg)](https://dragoscirjan.me/paypal)
[![Buy an item on our wishlist for us](https://img.shields.io/badge/wishlist-donate-yellow.svg)](https://dragoscirjan.me/wishlist)
-->

One Paragraph of project description goes here

<!--
Insert Table of Contents Here
This can be done using [AlanWalk.markdown-toc](https://marketplace.visualstudio.com/items?itemName=AlanWalk.markdown-toc) plugin,
which is also included in
[itmcdev.generic-extension-pack](https://marketplace.visualstudio.com/items?itemName=itmcdev.generic-extension-pack) extension pack.
-->
<!-- TOC -->

- [GoLang Project](#golang-project)
  - [Getting Started](#getting-started)
    - [Prereqiusites / Dependencies](#prereqiusites--dependencies)
      - [For Windows](#for-windows)
      - [For Linux](#for-linux)
      - [Known Issues / Troubleshooting](#known-issues--troubleshooting)
    - [Installation](#installation)
      - [Say what the step will be](#say-what-the-step-will-be)
      - [And repeat](#and-repeat)
    - [Development](#development)
      - [Requirements](#requirements)
        - [For Windows](#for-windows-1)
        - [For Linux/Unix/OSX](#for-linuxunixosx)
    - [Testing](#testing)
      - [Single Tests](#single-tests)
    - [Deployment](#deployment)
  - [Authors](#authors)
  - [Issues / Support](#issues--support)
  - [License](#license)
  - [Changelog](#changelog)

<!-- /TOC -->

## Getting Started

### Prereqiusites / Dependencies

#### For Windows

```powershell
# Give Examples
```

#### For Linux

```bash
# Give Examples

apt-get install build-essential mono
npm install -y node-gyp
```

#### Known Issues / Troubleshooting

- When developing, [golangci-lint](https://github.com/golangci/golangci-lint) will not be available on Windows.

### Installation

A step by step series of examples that tell you how to get a development env running

#### Say what the step will be

```
Give the example
```

#### And repeat

```
until finished
```

End with an example of getting some data out of the system or using it for a little demo

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

#### Single Tests

Run single unit tests file, by calling `make test-single TEST_PATH=./path/to/file/...`

```bash
make test-single TEST_PATH=./src/greet/...
```

### Deployment

Add additional notes about how to deploy this on a live system

## Authors

- [Dragos Cirjan](mailto:dragos.cirjan@gmail.com) - Initial work - [Go Template](/templ-project/go)

See also the list of contributors who participated in this project.

## Issues / Support

Add a set of links to the [issues](/templ-project/go/issues) page/website, so people can know where to add issues/bugs or ask for support.

## License

(If the package is public, add licence)
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Changelog

Small changelog history. The rest should be added to [CHANGELOG.md](CHANGELOG.md).

See here a template for changelogs: https://keepachangelog.com/en/1.0.0/

Also see this tool for automatically generating them: https://www.npmjs.com/package/changelog
