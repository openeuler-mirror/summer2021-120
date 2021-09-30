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

#### 方案描述

###### 技术体系
1. Docker
2. Kubernetes
3. golang
4. python

###### 总体方案
1. 基于kubernetes技术实现对这个解决方案中程序实例的容器化部署和编排
2. 基于kubebuilder技术开发框架,实现kubernetes的自定义资源建模设计以及实例管理
3. 实现对website源码仓库的监控,可抽象为对任意源码仓库的监控,实现可配置的对指定分支的扫描,实现自动构建,持续集成
4. 基于docker技术实现CI插件化管理,使系统可以方便的对接任意其他CI平台.
5. 通过定义kubernetes的自定义资源,以及实现对CDR的动态创建,实现对website持续集成构建版本的自动化部署,服务映射，以及对website持续集成构建版本部署实例的声明周期管理,自动销毁

###### 计划与进度
1.  基于k8s operator 实现website代码仓库自动扫描服务的自动部署与管理.  ## 已完成
2.  website代码仓库自动扫描与镜像构建.  ## 已完成
3.  web-ui临时镜像构建后的自动部署与服务地址管理.  ## 已完成
4.  web-ui临时实例的生命周期管理，实现超时退出，资源释放回收.  ## 已完成

###### 困难与心得
1. 关于k8s-operator与kubebuilder的取舍：
    
    k8s-operator与kubebuilder同为k8s生态全开发自定义资源管理控制服务的底层框架,k8s-operator诞生的比较早,开源案例多
    kubebuilder为官方推荐框架,功能全面,开源贡献工程师组织实力更强大,但相对k8s-operator可参考的案例较少。综合考虑还是选择kubebuilder,
    我们认为kubebuilder正在逐步取代k8s-operator。
    
2. 关于web-ui临时实例的生命周期管理,先后尝试了cronjob方案, 外部控制job方案, 均已失败告终, 
最终还是选择通过RepoScanner模块调有kubectl来创建和删除WebUI用户自定义资源,由kubebuilder实现的webui-controller
来实现对应的pod，service，ingress资源的创建和销毁。

