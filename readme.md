# openeuler/website-operator

#### 介绍

随着越来越多的开发者参与到社区开发，openEuler官网作为第一门户，如何方便大家在提交修改时，能快速和准确的看到修改变化尤为重要，本项目旨在提供临时网站渲染链接，支撑开发人员查看和修改。

#### 项目目标

1. 自动构建镜像，支持多实例(PR)渲染，有过期退出策略，资源自动回收。
2. 支持容器化部署


#### 安装教程

1.  build RepoScanner image & push image to your registry
2.  build operator

#### 已实现：

1.  基于k8s operator 实现website代码仓库自动扫描服务的自动部署与管理，
2.  website代码仓库自动扫描与镜像构建

#### 计划：

1.  web-ui临时镜像构建后的自动部署与服务地址管理，
2.  web-ui临时实例的生命周期管理，实现超时退出，资源释放回收，

