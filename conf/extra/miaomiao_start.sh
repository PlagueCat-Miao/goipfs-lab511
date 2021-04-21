#!/bin/bash

die() {
    echo "ERROR: $*. Aborting." >&2
    exit 1
}

opts=false
opte=false

while getopts ":a:seh" opt #第一个冒号表示忽略错误；字符后面的冒号表示该选项必须有自己的参数。
do
    case $opt in
        a)
          ipfs bootstrap add $OPTARG
          ;;
        s)
          $opte && die "Cannot specify option a after specifying option : start"
          opts=true
          ;;
        e)
          $opts && die "Cannot specify option a after specifying option : end"
          opte=true
          kill -9 $(lsof -i tcp:4001 -t)
          ;;
        h)
          echo "usage: `basename $0` [-a <ipfs Address>|-n] [-h]"
          echo "        -a            ipfs 加入节点，需要重启ipfs daemon"
          echo "                      like: `basename $0` -a /ip4/<IP Address>/tcp/4001/ipfs/<Node Hash>"
          echo "        -s            运行环境启动"
          echo "        -e            关闭ipfs守护进程"
          exit 1;;
    esac
done
if [ $opts == true ];then
    nginx
    service mysql start
    ipfs daemon &
    echo 运行环境已启动
fi

if [ $# == 0 ];then
  echo `basename $0` -h 显示帮助
  exit
fi








