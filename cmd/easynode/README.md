集成 collect,blockchain,task,task_api,store等服务到一个application, 该app拥有easynode全部功能模块

## Prerequisites

- go version >=1.20

## Building the source

(以linux系统为例)

- mkdir supernode & cd supernode
- git clone https://github.com/sunjiangjun/supernode.git
- cd supernode/cmd/supernode
- CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o supernode app.go
  (mac下编译linux程序为例，其他交叉编译的命令请自行搜索)
- ./supernode

## config.json 详解

- blockchain_config.json

  请参考[blockchain](https://github.com/sunjiangjun/supernode/blob/main/cmd/blockchain/README.md)


- collect_config.json

  请参考[collect](https://github.com/sunjiangjun/supernode/blob/main/cmd/collect/README.md)


- store_config.json

  请参考[store](https://github.com/sunjiangjun/supernode/blob/main/cmd/store/README.md)


- task_config.json

  请参考[task](https://github.com/sunjiangjun/supernode/blob/main/cmd/task/README.md)


- taskapi_config.json

  请参考[taskapi](https://github.com/sunjiangjun/supernode/blob/main/cmd/taskapi/README.md)

## run command

./supernode [OPTIONS]

``````
-collect_config  string  config path of collect server

-task_config  string  config path of task server

-blockchain_config  string  config path of blockchain server

-taskapi_config  string  config path of taskapi server

-taskapi_config  string  config path of taskapi server

-store_config  string  config path of store server

``````

## usages

- blockchain 服务的使用

  请参考 [blockchain](https://github.com/sunjiangjun/supernode/blob/main/cmd/blockchain/README.md)


- collect 服务的使用

  请参考 [collect](https://github.com/sunjiangjun/supernode/blob/main/cmd/collect/README.md)


- task 服务的使用

  请参考 [task](https://github.com/sunjiangjun/supernode/blob/main/cmd/task/README.md)


- taskapi 服务的使用

  请参考 [taskapi](https://github.com/sunjiangjun/supernode/blob/main/cmd/taskapi/README.md)


- store 服务的使用

  请参考 [store](https://github.com/sunjiangjun/supernode/blob/main/cmd/store/README.md)