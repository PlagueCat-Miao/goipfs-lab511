# goipfs-lab511
ipfs

## 网关主要实现功能
 - login登记
 - 强制存储功能
 - 存储记录登记
## quick-start
 - ipfs安装
    - ``` shell
        ipfs init
        cp swarm.key ~/.ipfs/swarm.key
        ipfs bootstrap rm --all
        ## ipfs bootstrap add /ip4/104.236.76.40/tcp/4001/ipfs/QmSoLV4Bbm51jM9C4gDYZQ9Cy3U6aXMJDAbzgu2fzaDs64
      ``` 
 - ipfs启动
    - `ipfs daemon`
 - 各个节点启动
    - `make runEdge` 边缘节点-脚本
    - `make runGateway` 网关-服务-mysql
    - `make runCloud` 云节点-服务-存储

## ipfs 系统管理
   - ipfs数据存放在 ./ipfs/block 中，通过`du -h --max-depth=1 ~/.ipfs` 可以查询

