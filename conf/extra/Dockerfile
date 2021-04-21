FROM ubuntu:16.04

MAINTAINER PlagueCat-Miao

# 创建文件目录
RUN mkdir -p /root/program/pkg
RUN mkdir -p /root/goworkspace/bin
RUN mkdir -p /root/goworkspace/pkg
RUN mkdir -p /root/goworkspace/src/github.com/PlagueCat-Miao

# 环境工具安装
RUN apt-get update
RUN apt-get install -y wget
RUN apt-get install -y git
RUN apt-get install -y lsof
RUN apt-get install -y zip
RUN apt-get install -y make

# 代码及conf文件下载
RUN git clone https://github.com/PlagueCat-Miao/goipfs-lab511.git /root/goworkspace/src/github.com/PlagueCat-Miao/goipfs-lab511
COPY miaomiao_start.sh /root/

# go 环境
RUN wget -c https://dl.google.com/go/go1.13.14.linux-amd64.tar.gz -P /root/program/pkg \
    && tar -C .. -xzf go1.13.14.linux-amd64.tar.gz
RUN export GOROOT=/root/program/go  && export GOPATH=/root/goworkspace &&  export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
RUN export GOPROXY=http://goproxy.cn,direct && export GO111MODULE="on"

# go-ipfs 环境
RUN wget -c https://github.com/ipfs/go-ipfs/releases/download/v0.8.0/go-ipfs_v0.8.0_linux-amd64.tar.gz -P /root/program/pkg \
    && tar -C .. -xzf go-ipfs_v0.8.0_linux-amd64.tar.gz
RUN export PATH=/root/program/go-ipfs:$PATH && ipfs init && ipfs bootstrap rm --all
COPY swarm.key /root/.ipfs/

# mysql 环境 静默安装
RUN { echo mysql-server-5.7 mysql-server/root_password password 12345678; echo mysql-server-5.7 mysql-server/root_password_again password 12345678 ; } \
| debconf-set-selections && apt-get update && apt-get install -y mysql-server
COPY Dump20201004.sql /root/program/pkg/
RUN service mysql start && echo "CREATE DATABASE ipfs_fileinfo;" >>  /root/program/pkg/CreateIpfs_fileinfo.sql && mysql -uroot -p12345678  <  /root/program/pkg/CreateIpfs_fileinfo.sql
RUN service mysql start && mysql -uroot -p12345678 ipfs_fileinfo <  /root/program/pkg/Dump20201004.sql

# nginx 环境安装 https://www.cnblogs.com/yanjieli/archive/2019/03/28/10615638.html
RUN apt-get install -y openssl libssl-dev
RUN apt-get install -y libpcre3 libpcre3-dev

RUN wget -c https://github.com/arut/nginx-rtmp-module/archive/master.zip -P /root/program/pkg \
    && unzip master.zip -d ..
RUN wget -c http://nginx.org/download/nginx-1.12.2.tar.gz -P /root/program/pkg \
    && tar -C .. -xzf nginx-1.12.2.tar.gz
RUN cd /root/program/nginx-1.12.2 \
    && ./configure --prefix=/opt/nginx-1.12.2 --add-module=/root/program/nginx-rtmp-module-master --with-http_ssl_module  \
    && make \
    && make install
RUN export PATH=/opt/nginx-1.12.2/sbin:$PATH

COPY nginx.conf /root/program/pkg/
RUN nginx -t -c /root/program/pkg/nginx.conf
RUN cp -f /root/program/pkg/nginx.conf  /opt/nginx-1.12.2/conf/nginx.conf && nginx -s reload

CMD cd /root/







