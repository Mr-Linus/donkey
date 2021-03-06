### Donkey——大笨驴

#### 一个简单的用go实现的容器引擎

![donkey](./img/logo.png)

> 本容器引擎在[xianlubird/mydocker](https://github.com/xianlubird/mydocker)基础上改进，实现了大部分基础功能，仍存在些许bug，不建议生产环境使用。


### 使用环境
- OS: Ubuntu 14.04
- Kernel: 3.13.0-153-generic
- go:1.7


### 运行示例(root)
```bash
# 执行容器终端
donkey run -ti busybox sh
{"level":"info","msg":"init come on","time":"2018-07-28T18:02:10Z"}
{"level":"info","msg":"command /bin/sh","time":"2018-07-28T18:02:10Z"}
{"level":"info","msg":"command /bin/sh","time":"2018-07-28T18:02:10Z"}
#
# echo "hello"
hello
#ps
  PID TTY          TIME CMD
    1 pts/0    00:00:00 sh
    5 pts/0    00:00:00 ps

# 执行命令
donkey run -ti busybox ls
root@vagrant-ubuntu-trusty-64:/home/vagrant/works/donkey# donkey run -ti /bin/ls
{"level":"info","msg":"init come on","time":"2018-07-28T18:01:48Z"}
{"level":"info","msg":"command /bin/ls","time":"2018-07-28T18:01:48Z"}
{"level":"info","msg":"command /bin/ls","time":"2018-07-28T18:01:48Z"}
container  donkey  main_command.go  main.go  readme.md	run.go

#资源限制
donkey  run -ti -m 100m -cpushare 512 busybox sh
{"level":"info","msg":"command all is /bin/sh","time":"2018-07-29T15:24:47Z"}
{"level":"info","msg":"init come on","time":"2018-07-29T15:24:47Z"}
{"level":"info","msg":"Find path /bin/sh","time":"2018-07-29T15:24:47Z"}
#

# 运行脚本
 donkey  run -ti -m 100m  busybox bash test.sh
{"level":"info","msg":"command all is bash test.sh","time":"2018-07-31T16:08:55Z"}
{"level":"info","msg":"init come on","time":"2018-07-31T16:08:55Z"}
{"level":"info","msg":"Find path /bin/bash","time":"2018-07-31T16:08:55Z"}
stress: info: [4] dispatching hogs: 0 cpu, 0 io, 1 vm, 0 hdd

# 使用镜像 busybox
donkey run -ti busybox sh
{"level":"info","msg":"command all is sh","time":"2018-08-04T10:02:24Z"}
{"level":"info","msg":"init come on","time":"2018-08-04T10:02:24Z"}
{"level":"info","msg":"Current location is /home/vagrant/works/donkey/images/busybox","time":"2018-08-04T10:02:24Z"}
{"level":"info","msg":"Find path /bin/sh","time":"2018-08-04T10:02:24Z"}
/ # busybox
BusyBox v1.29.2 (2018-07-31 20:19:16 UTC) multi-call binary.
BusyBox is copyrighted by many authors between 1998-2015.
Licensed under GPLv2. See source distribution for detailed
copyright notices.


# 映射目录
donkey run -ti -v ./cgroups/:/cgroups busybox sh
{"level":"error","msg":"Mkdir dir ./images/rw/ error. mkdir ./images/rw/: file exists","time":"2018-08-09T10:21:18Z"}
{"level":"info","msg":"Mkdir parent dir ./cgroups/ error. mkdir ./cgroups/: file exists","time":"2018-08-09T10:21:18Z"}
{"level":"info","msg":"[\"./cgroups/\" \"/cgroups\"]","time":"2018-08-09T10:21:18Z"}
{"level":"info","msg":"command all is sh","time":"2018-08-09T10:21:18Z"}
{"level":"info","msg":"init come on","time":"2018-08-09T10:21:18Z"}
{"level":"info","msg":"Current location is /root/donkey/images/mnt","time":"2018-08-09T10:21:18Z"}
{"level":"info","msg":"Find path /bin/sh","time":"2018-08-09T10:21:18Z"}
/ # ls
bin      cgroups  dev      etc      home     proc     root     sys      tmp      usr      var

# 镜像打包
donkey commit testcontainer test
test.tar

# 后台运行
donkey run -d busybox echo hello
{"level":"info","msg":"createTty false","time":"2018-08-14T09:30:36Z"}
{"level":"info","msg":"command all is echo hello","time":"2018-08-14T09:30:36Z"}


# 容器状态显示
donkey ps
ID           NAME         PID         STATUS      COMMAND     CREATED
6575021233   6575021233   2731        running     top         2018-08-17 14:18:46


#容器日志显示
donkey run --name test -d echo hello
{"level":"info","msg":"createTty false","time":"2018-08-17T15:17:29Z"}
{"level":"info","msg":"command all is echo hello","time":"2018-08-17T15:17:29Z"}
{"level":"warning","msg":"remove cgroup fail remove /sys/fs/cgroup/memory/donkey-cgroup/memory.kmem.tcp.max_usage_in_bytes: operation not permitted","time":"2018-08-17T15:17:29Z"}
{"level":"warning","msg":"remove cgroup fail remove /sys/fs/cgroup/cpu/donkey-cgroup/cpu.stat: operation not permitted","time":"2018-08-17T15:17:29Z"}
donkey logs test
{"level":"info","msg":"init come on","time":"2018-08-17T15:17:29Z"}
{"level":"info","msg":"Current location is /root/busybox","time":"2018-08-17T15:17:29Z"}
{"level":"info","msg":"Find path /bin/echo","time":"2018-08-17T15:17:29Z"}
hello


#进入容器namespace
donkey exec test sh
{"level":"info","msg":"container pid 2668","time":"2018-08-18T16:59:00Z"}
{"level":"info","msg":"command sh","time":"2018-08-18T16:59:00Z"}
/ # ps
PID   USER     TIME  COMMAND
    1 root      0:00 top
    4 root      0:00 sh
    5 root      0:00 ps
/ #

# 传入环境变量
donkey run -ti -e hello=test busybox sh
{"level":"info","msg":"createTty true","time":"2018-08-19T14:07:20Z"}
{"level":"info","msg":"command all is sh","time":"2018-08-19T14:07:20Z"}
{"level":"info","msg":"init come on","time":"2018-08-19T14:07:20Z"}
{"level":"info","msg":"Current location is /root/mnt/0263972906","time":"2018-08-19T14:07:20Z"}
{"level":"info","msg":"Find path /bin/sh","time":"2018-08-19T14:07:20Z"}
/ # echo $hello
test
```


### 开发日志
- V2.0.0
Date: 2018.6.27
单一程序实现简单容器构建,实现Namespace隔离
- V3.0.0 
Date: 2018.7.28
构建程序大体架构,参考Docker实现命令行构建容器
- V3.1.0
Date: 2018.7.29
增加CPU,Memory资源限制功能
- V3.2.0
Date: 2018.8.1
增加运行脚本功能                                                                                              
- V3.2.1
Date: 2018.8.4
增加Busybox镜像功能
- V3.3.0 
Date: 2018.8.7 
为容器运行增加写层
- V3.4.0
Date: 2018.8.9
增加容器映射目录功能
- V3.5.0
Date: 2018.8.14
增加镜像打包功能
- V3.5.1
Date: 2018.8.14
增加镜像打包功能， 增加容器后台运行功能
- V4.0.0-Alpha 
Date: 2018.8.17
增加容器状态显示功能,删除Volume功能,它将在4.0版本会进行重构
- V4.0.1-Alpha
Date: 2018.8.17
增加日志查看功能
- V4.1.1
Date: 2018.8.19
增加容器暂停、删除、Exec功能，重写Volume功能
- V4.2.1
Date: 2018.8.20
增加容器环境变量传入功能

![img-source-from-https://github.com/docker/dockercraft](https://github.com/docker/dockercraft/raw/master/docs/img/contribute.png?raw=true)