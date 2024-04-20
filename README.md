<div align="center">
<h1>glactl</h1>

[![Auth](https://img.shields.io/badge/Auth-eryajf-ff69b4)](https://github.com/eryajf)
[![GitHub Pull Requests](https://img.shields.io/github/stars/eryajf/glactl)](https://github.com/eryajf/glactl/stargazers)
[![HitCount](https://views.whatilearened.today/views/github/eryajf/glactl.svg)](https://github.com/eryajf/glactl)
[![GitHub license](https://img.shields.io/github/license/eryajf/glactl)](https://github.com/eryajf/glactl/blob/main/LICENSE)
[![](https://img.shields.io/badge/Awesome-MyStarList-c780fa?logo=Awesome-Lists)](https://github.com/eryajf/awesome-stars-eryajf#readme)

<p> 🌉 go-ldap-admin 项目对应的IM测试工具 🌉</p>

<img src="https://cdn.jsdelivr.net/gh/eryajf/tu@main/img/image_20240420_214408.gif" width="800"  height="3">

</div>

运维也可以如此优雅！快用这个框架打造一个专属于你的工具箱吧！

通过这个框架，你可以快速上手，直接构建你想要的运维工具，而不必再考虑配置，框架设计等内容。

## 如何使用


也可以编译成二进制，然后通过如下方式查看帮助信息：

```
# 编译
$ make build

#运行测试
$ ./glactl  -h
通过命令行获取配置信息

Available Commands:
  completion  generate the autocompletion script for the specified shell
  dingding    用于测试钉钉的数据获取是否正常
  feishu      用于测试钉钉的数据获取是否正常
  help        Help about any command
  wecom       用于测试企微的数据获取是否正常
```

也可以直接在release中下载二进制。或者直接使用docker来执行：

```
docker run -it registry.cn-hangzhou.aliyuncs.com/eryajf/glactl /app/glactl -h
```

## 感谢开源

- [eryajfctl](https://github.com/eryajf/eryajfctl)

如果觉得项目不错，请别忘了一键三连，给个 star。