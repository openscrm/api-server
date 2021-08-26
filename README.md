<p style="text-align: center">
  <img alt="logo" height="48" src="https://openscrm.oss-cn-hangzhou.aliyuncs.com/public/openscrm_logo.svg">
</p>

<h3 style="text-align: center">
安全，强大，易开发的企业微信SCRM
</h3>

[文档](https://docs.openscrm.cn/) |
[截图](#项目截图) |
[演示](#联系作者) |
[安装](https://docs.openscrm.cn/an-zhuang-jiao-cheng) 

### 项目简介

> OpenSCRM是一套基于**Go**和**React**的**超高质量**企业微信私域流量管理系统

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

### 对于PHPer
我想大多数PHPer看到本项目一定有强烈的熟悉感，没错，此项目吸收了大量Laravel的宝贵经验，代码清晰易读，
几乎没有额外门槛。对于没有接触过Go的PHPer，只需学习掌握Go的管道协程即可驾驭本项目。
> 与其学Swoole，何不用Go

Swoole和Go的学习成本是差不多的，最核心的都是协程，管道。额外的只需学习Go生态的Gin和GORM两个库即可。
Go基础库不熟，可以使用goframe过渡。

### 项目截图

![](https://openscrm.oss-cn-hangzhou.aliyuncs.com/public/screenshots/%E5%90%8E%E5%8F%B0%E9%A6%96%E9%A1%B5.png)
![](https://openscrm.oss-cn-hangzhou.aliyuncs.com/public/screenshots/%E4%BF%AE%E6%94%B9%E6%B8%A0%E9%81%93%E6%B4%BB%E7%A0%81.png)
![](https://openscrm.oss-cn-hangzhou.aliyuncs.com/public/screenshots/%E4%BC%9A%E8%AF%9D%E5%AD%98%E6%A1%A3.png)
![](https://openscrm.oss-cn-hangzhou.aliyuncs.com/public/screenshots/%E4%BF%AE%E6%94%B9%E6%B8%A0%E9%81%93%E6%B4%BB%E7%A0%812.png)
![](https://openscrm.oss-cn-hangzhou.aliyuncs.com/public/screenshots/%E5%AE%A2%E6%88%B7%E6%A0%87%E7%AD%BE%E7%AE%A1%E7%90%86.png)
![](https://openscrm.oss-cn-hangzhou.aliyuncs.com/public/screenshots/%E4%BF%AE%E6%94%B9%E7%BE%A4%E5%8F%91.png)
![](https://openscrm.oss-cn-hangzhou.aliyuncs.com/public/screenshots/%E4%BF%AE%E6%94%B9%E6%AC%A2%E8%BF%8E%E8%AF%AD.png)

### 联系作者加入交流群

<img src="https://openscrm.oss-cn-hangzhou.aliyuncs.com/public/screenshots/qrcode.png" width="200" />

扫码可加入交流群

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
https://docs.openscrm.cn/an-zhuang-jiao-cheng

### 演示地址
请联系作者获取演示地址

### 联系作者

<img src="https://openscrm.oss-cn-hangzhou.aliyuncs.com/public/screenshots/qrcode.png" width="200" />

扫码可加入交流群/获取演示地址

### 版权声明

OpenSCRM是开源软件，但仅用于学习和研究，商用请联系我们购买授权。

<img src="https://openscrm.oss-cn-hangzhou.aliyuncs.com/public/screenshots/OpenSCRM%E7%A7%81%E5%9F%9F%E6%B5%81%E9%87%8F%E7%AE%A1%E7%90%86%E7%B3%BB%E7%BB%9F-%E8%AF%81%E4%B9%A6.jpg" width="400" />