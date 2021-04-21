# goipfs-lab511
喵喵毕设，一个基于IPFS的视频存储系统，（虽然一个文件系统上做存储系统怪怪的）。
欢迎交流与引用喵

## 介绍
系统主要建立一个中心对IPFS的swap存储策略进行了增强。尽可能避免IPFS的一些问题。
- 冷热数据存储分化。
- 处理不了RTSP视频流。
- 存储负载不均衡。
## quick-start
- ipfs初始化
  -  ```  shell
     ipfs init
     cp swarm.key ~/.ipfs/swarm.key
     ipfs bootstrap rm --all
     ipfs bootstrap add /ip4/<IP>/tcp/4001/ipfs/<node hash>
     ```
- 启动
   - `ipfs daemon`
   - `make runEdge` or  `make runGateway` or `make runCloud`

- 其他
   - 占用查询 `du -h --max-depth=1 ~/.ipfs/block`
   - 播放视频 `ffplay -fflags nobuffer -analyzeduration 500000 -i rtmp://127.0.0.1:1935/live`
   - 视频推流 `ffmpeg -r 30 -i /dev/video0 -vcodec h264 -max_delay 100 -f flv -g 5 -b 700000 rtmp://127.0.0.1:1935/live -map 0:0 -map 0:2`         
   - 性能展示 访问页面`http://127.0.0.1:xxxx/debug/pprof/` [pprof教程](https://segmentfault.com/a/1190000016412013)
   - 美化代码 `gofmt -l -w .` `go mod tidy`
## 使用docker的lab511小伙伴注意
   - 如果你是拿到了镜像压缩包
        - 加载镜像 `docker load -i xxxxx.tar` (goipfs.tar)
        - 运行镜像 `docker run -it <image id> -p <port:port> ` (aee0defcd78a,8434)
        - (进入docker系统) 
        - 运行初始化脚本 `/home/hellcat/my_init.sh` 
        - 进入代码工作区 `cd /home/hellcat/goworkspace/src/github.com/PlagueCat-Miao/goipfs-lab511`
        - 启动 `make runEdge` or  `make runGateway` or `make runCloud`
   - 如果你使用Dockerfile ( *建议* )
        - 放置附件 进入Dockerfile所在文件夹，并将以下文件拷贝至同文件夹下
            - smarm.key 将私钥文件
            - Dump20201004.sql 建库sql文件
            - nginx.conf nginx配置文件
            - miaomiao_start.sh 快捷交互脚本
        - 构建镜像 `docker build -t miao:v34 .`
            - 如果网不好可以多次尝试
            - 若还不行则自行下载包并对Dockerfile的对应模块进行COPY替换
        - 进入镜像 `docker run -it -p 8434:8434 miao:v34`
        - 添加邻居节点 `./miaomiao_start.sh -a <ipfs_id>`
        - 启动环境 `./miaomiao_start.sh -s` 
        - 进入代码工作区 `cd /root/goworkspace/src/github.com/PlagueCat-Miao/goipfs-lab511`
        - 启动 `make runEdge` or  `make runGateway` or `make runCloud`
        
            
       
## 文件目录说明
    .
    ├── cmd                             //脚本 - 仅边缘层节点使用
    ├── conf                            //配置文件   
    ├── constdef                        //常量配置文件
    ├── dal                             //底层数据交互方法
    │   ├── db                              //gorm 处理Mysql
    │   ├── httppack                        //http数据报文格式
    │   ├── ipfs                            //IPFS接口交互
    │   └── rtmpffmpeg                      //视频流处理
    ├── model                          //数据模板定义
    ├── nodes                          //定义各身份节点的特有方法
    ├── operate                        //定义服务所需struct及其操作函数
    ├── service                        //提供服务接口
    ├── util                           //系统通用处理函数
    ├── main.go                        
    ├── Makefile                       //定义启动函数
    └── README.md                          
    
    ~/.ipfs
    ├── block                           //ipfs数据存放
    └── swarm.key                       //系统通信密钥

### 附录

1. 怎么解决跨域？ 代理？
    本科生帖子
   
2. Create time 溢出了

3.  