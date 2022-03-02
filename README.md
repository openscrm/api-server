<p align="center">
  <img alt="logo" height="48" src="https://openscrm.oss-cn-hangzhou.aliyuncs.com/public/openscrm_logo.svg">
</p>

<h3 align="center">
安全，强大，易开发的企业微信SCRM
</h3>

<h4 align="center">

<a href="https://github.com/openscrm/api-server/wiki" target="_blank">文档</a> |
[截图](#项目截图) |
<a href="http://dashboard.demo.openscrm.cn:8000/" target="_blank">演示</a> |
<a href="https://github.com/openscrm/api-server/wiki/%E5%AE%89%E8%A3%85%E6%95%99%E7%A8%8B" target="_blank">安装</a> 

</h4>


### 项目简介

> OpenSCRM是一套基于**Go**和**React**的**超高质量**企业微信私域流量管理系统

### 在线演示

[http://dashboard.demo.openscrm.cn:8000/](http://dashboard.demo.openscrm.cn:8000/)

### 项目特点

* **安全性高**：企业微信控制了企业所有员工和客户的敏感数据，如电话号码，职位，客户标签，联系方式等，如果发生泄露，
  对企业的打击是非常致命的。我们团队有来自360的资深安全开发工程师cover此问题。


* **高性能，高稳定性**：得益于Go出色的工程能力，简单有效的并发控制能力，OpenSCRM具备比肩头部Saas厂商的高性能和高稳定


* **代码可读性优先**：我们深刻认同Google对于代码管理的看法，项目开发完成只是项目的开始，更多的工作在于维护和迭代，
  唯有易读的代码才能保证后期迭代的高质量，高效率，这也是Go语言的设计目标。我们有非常完善的代码注释，所有代码力求清晰易读。


* **易开发**：作为开源项目，我们为了让更多的人可以受益于此项目，我们做了大量工作力求项目简单易上手。
  我们只做必要抽象（MVC），避免引入新慨念。我们坚持尽量少的中间件依赖，仅依赖Mysql和Redis，
  比如延迟队列我们基于Redis实现， 没有引入Kafka；比如全文检索基于Mysql8全文检索实现，没有引入ES。

> Python，PHP，NodeJS开发者可以放心使用本项目，本项目做了大量工作力求简单，非常容易上手。

### 项目截图

![](https://openscrm.oss-cn-hangzhou.aliyuncs.com/public/screenshots/%E5%90%8E%E5%8F%B0%E9%A6%96%E9%A1%B5.png)
![](https://openscrm.oss-cn-hangzhou.aliyuncs.com/public/screenshots/%E4%BF%AE%E6%94%B9%E6%B8%A0%E9%81%93%E6%B4%BB%E7%A0%81.png)
![](https://openscrm.oss-cn-hangzhou.aliyuncs.com/public/screenshots/%E4%BC%9A%E8%AF%9D%E5%AD%98%E6%A1%A3.png)
![](https://openscrm.oss-cn-hangzhou.aliyuncs.com/public/screenshots/%E4%BF%AE%E6%94%B9%E6%B8%A0%E9%81%93%E6%B4%BB%E7%A0%812.png)
![](https://openscrm.oss-cn-hangzhou.aliyuncs.com/public/screenshots/%E5%AE%A2%E6%88%B7%E6%A0%87%E7%AD%BE%E7%AE%A1%E7%90%86.png)
![](https://openscrm.oss-cn-hangzhou.aliyuncs.com/public/screenshots/%E4%BF%AE%E6%94%B9%E7%BE%A4%E5%8F%91.png)
![](https://openscrm.oss-cn-hangzhou.aliyuncs.com/public/screenshots/%E4%BF%AE%E6%94%B9%E6%AC%A2%E8%BF%8E%E8%AF%AD.png)

### 技术栈
#### 后端技术栈
* [Go](https://learnku.com/docs/the-way-to-go)
* [Gin](https://learnku.com/docs/gin-gonic/2019)
* [GORM](https://gorm.io/zh_CN/docs/)
* [Goframe](https://goframe.org/pages/viewpage.action?pageId=1114411)
* Redis
* Mysql >= 5.7

#### 前端技术栈
* [React](https://zh-hans.reactjs.org/)
* [TypeScript](https://www.tslang.cn/docs/handbook/typescript-in-5-minutes.html)
* [Ant Design](https://ant.design/components/overview-cn/)
* [Ant Design Pro](https://pro.ant.design/zh-CN/docs/overview)
* [Pro Components](https://procomponents.ant.design/components)


### 目录结构

```
├─app
│  ├─callback 企业微信事件回调处理
│  │  ├─customer_event
│  │  ├─department_event
│  │  ├─group_chat_event
│  │  ├─msg_arch_event
│  │  ├─staff_event
│  │  └─tag_event
│  ├─constants 常量定义
│  ├─consumers 队列消费
│  ├─controller 控制器
│  ├─entities 消息实体，主要定义参数，请求，响应结构体
│  ├─middleware gin请求中间件
│  ├─models 数据库模型
│  ├─requests 请求定义
│  ├─responses 响应定义
│  ├─services 服务
│  ├─tasks 定时任务
├─bin 二进制文件
├─common 共同库
│  ├─app 基于Gin封装的常用请求响应处理函数
│  ├─delay_queue 基于Redis延迟队列
│  ├─ecode 错误码
│  ├─id_generator uuid生成
│  ├─log 日志
│  ├─redis redis操作库
│  ├─session session会话
│  ├─storage 存储
│  ├─util 常用工具函数
│  └─validator 请求验证
├─conf 配置文件
├─docker
│  ├─data
│  │  ├─dashboard
│  │  │  └─dist 管理后台构建的前端静态文件
│  │  ├─mysql
│  │  │  ├─conf mysql容器配置文件
│  │  │  └─db mysql容器数据文件
│  │  ├─nginx
│  │  │  ├─conf nginx容器配置文件
│  │  │  │  └─conf.d 
│  │  │  └─logs
│  │  ├─redis 
│  │  │  └─db redis容器数据文件
│  │  └─sidebar
│  │      └─dist 侧边栏构建的前端静态文件
│  └─lib 企业微信提供的会话存档动态链库
├─docs 文档
├─pkg 三方库
│  └─easywework 企业微信Api调用库
│      ├─errcodes 企业微信Api错误码
├─routers Gin路由
├─scripts 脚本
└─test 测试代码

```

### 安装教程
https://github.com/openscrm/api-server/wiki/%E5%AE%89%E8%A3%85%E6%95%99%E7%A8%8B

### Api调试
`docs目录包含postman导出文件，可方便调试api`

### 版权声明

OpenSCRM遵循Apache2.0协议，可免费商用
