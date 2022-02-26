# 一、安装go语言的环境

下载地址：https://dl.google.com/go/go1.17.7.linux-amd64.tar.gz

安装步骤（可以去官网查看 https://golang.google.cn/doc/install）：

```sh
#将下载的文件传到 linux 服务器之后删除以前可能存在的go语言的环境并解压到 /usr/local 目录
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.7.linux-amd64.tar.gz
#设置 go 的环境变量
export PATH=$PATH:/usr/local/go/bin
#验证 go 语言环境是否安装成功
go version
```

配置代理环境（参考地址：https://goproxy.io/zh/），方便后面编译的时候下载插件包

```sh
# 配置 GOPROXY 环境变量
export GOPROXY=https://goproxy.io,direct
# 还可以设置不走 proxy 的私有仓库或组，多个用逗号相隔（可选）
# export GOPRIVATE=git.mycompany.com,github.com/my/private
```

# 二、安装 docker 环境

如果已经有了docker环境可以直接跳过此步骤

```sh
sudo yum install -y yum-utils device-mapper-persistent-data lvm2
# Step 2: 添加软件源信息
sudo yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
# Step 3: 更新并安装 Docker-CE
sudo yum makecache fast
sudo yum -y install docker-ce
# Step 4: 开启Docker服务
sudo service docker start

# 注意：其他注意事项在下面的注释中
# 官方软件源默认启用了最新的软件，您可以通过编辑软件源的方式获取各个版本的软件包。例如官方并没有将测试版本的软件源置为可用，你可以通过以下方式开启。同理可以开启各种测试版本等。
# vim /etc/yum.repos.d/docker-ce.repo
#   将 [docker-ce-test] 下方的 enabled=0 修改为 enabled=1
#
# 安装指定版本的Docker-CE:
# Step 1: 查找Docker-CE的版本:
# yum list docker-ce.x86_64 --showduplicates | sort -r
#   Loading mirror speeds from cached hostfile
#   Loaded plugins: branch, fastestmirror, langpacks
#   docker-ce.x86_64            17.03.1.ce-1.el7.centos            docker-ce-stable
#   docker-ce.x86_64            17.03.1.ce-1.el7.centos            @docker-ce-stable
#   docker-ce.x86_64            17.03.0.ce-1.el7.centos            docker-ce-stable
#   Available Packages
# Step2 : 安装指定版本的Docker-CE: (VERSION 例如上面的 17.03.0.ce.1-1.el7.centos)
# sudo yum -y install docker-ce-[VERSION]
# 注意：在某些版本之后，docker-ce安装出现了其他依赖包，如果安装失败的话请关注错误信息。例如 docker-ce 17.03 之后，需要先安装 docker-ce-selinux。
# yum list docker-ce-selinux- --showduplicates | sort -r
# sudo yum -y install docker-ce-selinux-[VERSION]

# 通过经典网络、VPC网络内网安装时，用以下命令替换Step 2中的命令
# 经典网络：
# sudo yum-config-manager --add-repo http://mirrors.aliyuncs.com/docker-ce/linux/centos/docker-ce.repo
# VPC网络：
# sudo yum-config-manager --add-repo http://mirrors.could.aliyuncs.com/docker-ce/linux/centos/docker-ce.repo
```

配置阿里镜像源

```sh
 sudo mkdir -p /etc/docker
 echo '{"registry-mirrors": ["https://thu8zyqr.mirror.aliyuncs.com"]}' >> /etc/docker/daemon.json
 sudo systemctl daemon-reload
 sudo systemctl restart docker
```

安装 docker-compose

```sh
#下载 docker-compose
curl -L https://github.com/docker/compose/releases/download/1.29.2/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
#由于github不太稳定，可以使用以下方式下载
curl -L https://get.daocloud.io/docker/compose/releases/download/1.29.2/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose

#授予可执行权限
chmod +x /usr/local/bin/docker-compose
#测试是否安装成功
docker compose version
```

# 三、下载源码并进行编译

下载地址：https://github.com/openscrm/api-server

此处下载zip压缩包

```sh
#上传之后解压
unzip api-server-master.zip
#修改名称
mv api-server-master api-server
#进入 api-server 文件夹
cd /home/api-server
#修改配置文件
cp conf/config.example.yaml conf/config.yaml
#编辑 config.yaml 文件，修改里面带 * 的配置
#以下如果使用 Docker for Mac 可能不需要修改（不确定，没用过 Mac 和 host.docker.internal 这个配置，去官网看好像只有 Mac 版本支持）
#将 Server.MsgArchSrvHost（会话存档服务访问地址）中的 host.docker.internal 改为 api-server_msg-archive-server_1
#将 DB.Host 中的 host.docker.internal 改为 api-server_mysql_1
#将 Redis.Host 中的 host.docker.internal 改为 api-server_redis_1
vim conf/config.yaml

#修改 nginx 中的配置
cd /home/api-server/docker/data/nginx/conf/conf.d
#修改 dashboard.conf，将 location 中的代理地址从 127.0.0.1 改为 api-server_api-server_1
vim dashboard.conf
#修改 sidebar.conf，同样将 location 中的代理地址从 127.0.0.1 改为 api-server_api-server_1
vim sidebar.conf

#编译
#进入 api-server 文件夹(如果目前就在这个文件夹，忽略此命令) 
cd /home/api-server
#下载各种插件
go mod vendor
#编译项目
go build -o bin/api-server
#修改 docker-compose.yaml 配置文件，找到编排文件中 api-server 服务，修改启动后执行的命令
#找到命令中以下一行
/srv/api-server
#将其改为
/srv/api-server/openscrm
```

# 四、使用docker-compose启动容器

```sh
#启动容器
docker-compose up -d
#注意：容器启动过程中可能会存在 msg-archive-server 与 api-server 启动失败的情况，不要急，这可能是 MySQL 数据库还没有初始化完成，等待一会，然后重新启动

#在等待的过程中 查看 redis 挂载的数据目录是否创建成功，当创建成功后需要授予其写入权限，用来保存持久化的rdb快照，否则 docker 会一直刷无法保存 redis 快照的错误日志，直到将服务器的磁盘撑爆
chomd 777 /home/api-server/docker/data/redis/db

#其它相关命令
#查看日志
docker-compose logs
#关闭容器
docker-compose down
```

# 五、错误问题清单

## 1、空员工数据问题

![微信图片_20220226213504](https://gitee.com/yezhiqiu521/picture/raw/master/picture/%E5%BE%AE%E4%BF%A1%E5%9B%BE%E7%89%87_20220226213504.png)

如果 api-server_api-server_1 服务没有起来，并且查看日志发现了上面截图的问题，此时需要在企业微信中邀请一个用户成为你的客户就可以了。
