### Donkey

一个简单的用go实现的容器引擎

### 使用环境
- OS: Ubuntu 14.04
- Kernel: 3.13.0-153-generic
- go:1.7


### 运行(root)
```bash
# 生成容器终端
donkey run -ti /bin/sh
{"level":"info","msg":"init come on","time":"2018-07-28T18:02:10Z"}
{"level":"info","msg":"command /bin/sh","time":"2018-07-28T18:02:10Z"}
{"level":"info","msg":"command /bin/sh","time":"2018-07-28T18:02:10Z"}
#
# echo "hello"
hello
#

# 执行命令
donkey run -ti /bin/ls
root@vagrant-ubuntu-trusty-64:/home/vagrant/works/donkey# donkey run -ti /bin/ls
{"level":"info","msg":"init come on","time":"2018-07-28T18:01:48Z"}
{"level":"info","msg":"command /bin/ls","time":"2018-07-28T18:01:48Z"}
{"level":"info","msg":"command /bin/ls","time":"2018-07-28T18:01:48Z"}
container  donkey  main_command.go  main.go  readme.md	run.go
```

