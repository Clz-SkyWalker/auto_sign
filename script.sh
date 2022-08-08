#!/bin/bash

#获取输入的第一个参数
cmd=$1

programeName=auto_sign

#没有输入参数时提醒内容 $#参数的个数
if [ $# -eq 0 ]; then 
	echo "please input start|stop|restart|build|buildDocker"
	exit
fi

#获取需要杀死进程的id  xxx.jar 你需要拉起的服务名
vehicle_pid=`ps -ef | grep $programeName |grep -v 'grep'|awk '{print $2}'`

#杀死服务 注意 [] 和$vehicle_pid 存在空格
killpid(){
	if [ $vehicle_pid ];then
	for id in $vehicle_pid
	do
	kill -9 $id
	done
else
	echo "vehicle_pid is not exists"
fi
echo "stop success"
}

#启动服务
startup(){
	#服务所存放的目录，我这里是jar包存放的目录
	vehicle_repid=`ps -ef | grep $programeName |grep -v 'grep'|awk '{print $2}'`
	if [ $vehicle_repid ];then
		echo "服务已经启动"
else
  nohup ./auto_sign > auto_sign.log 2>&1 &
  vehicle_pid_temp=`ps -ef | grep $programeName |grep -v 'grep'|awk '{print $2}'`
	if [ $vehicle_pid_temp ];then
		echo  "服务启动成功，pid:${vehicle_pid_temp}"
	else
		echo "启动失败"
	fi
fi
echo "start success"
}

# 打包
build(){
  export GOPROXY=https://goproxy.io
  go mod tidy
  # 打印依赖，部署成功后查看版本依赖是否如预期
  cat ./go.mod
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/auto_sign ./cmd/main.go
  echo "build success"
}
# 部署
deploy(){
  docker-compose down
  docker-compose build
  docker-compose up -d
  echo "deploy success"
}

if [ $cmd == "build" ]
then
  build
fi

if [ $cmd == "buildDocker" ]
then
  build
  deploy
fi

if [ $cmd == "start" ]
then
  startup
fi

if [ $cmd == "stop" ];then
  killpid
fi

if [ $cmd == "restart" ];then
  killpid
  startup
  echo "restart success"
fi
