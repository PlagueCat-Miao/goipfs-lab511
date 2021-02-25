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

###附录
   