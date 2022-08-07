#!/bin/bash

#获取输入的第一个参数
cmd=$1

programeName=auto_sign

#没有输入参数时提醒内容 $#参数的个数
if [ $# -eq 0 ]; then 
	echo "please input start|stop|restart|build"
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
}

if [ $cmd == "build" ]
then
  go build -o ./bin/auto_sign ./cmd/main.go
  echo "build success"
fi

if [ $cmd == "start" ]
then
  startup
  echo "start success"
fi

if [ $cmd == "stop" ];then
  killpid
  echo "stop success"
fi

if [ $cmd == "restart" ];then
  killpid
  startup
  echo "restart success"
fi
