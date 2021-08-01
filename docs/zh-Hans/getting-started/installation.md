# 安装

## 国际化

[English](../../installation.md) | [简体中文](installation.md)

## 介绍

Tiga CLI 是整个 Tiga 生态的命令行工具，包含了对前端和 Go 服务端开发的支持，用户可能仅仅是前端开发者或者 Go 开发者，所以我们对两者分别提供了不同的安装方法，您可以自行选择适合您的方式。

## Go 安装

从 [Go 下载页面]()下载 Go 语言开发工具。

确保您遵守官方的 [Go 安装说明]()。还需要确保您的`PATH`环境变量包含 `~/go/bin`目录路径。重启终端并执行以下命令检查：

- 检查 Go 是否正确安装：`go version`
- 检查“~/go/bin”路径是否在 PATH 变量内：`echo $PATH | grep go/bin`
- 检查 “GO111MODULE” 环境变量是否设置为 ”on“： `echo ${GO111MODULE}`

运行安装命令：

```shell
go install github.com/tigateam/tiga-cli/cmd/tiga

```

## NPM 安装

确保您的环境中有 Node 和 NPM：

- 检查 npm 是否正确安装：`npm -v`

运行安装命令：

```shell
npm i -g @tigateam/cli

```

## 检查是否安装成功

运行命令：

```shell
tiga version

```

如果看到输出版本号，则表示您已经正确安装了 Tiga CLI。
