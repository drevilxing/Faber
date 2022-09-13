基础环境安装
---------------------------------------

安装Hyperledger Fabric框架2.2版本运行环境包括如下内容

.. list-table:: Environment
    :widths: 15 15 30
    :header-rows: 1

    * - Software
      - Version
      - Description
    * - docker
      - latest
      - 提供容器运行环境
    * - docker-compose
      - latest
      - 完成容器运行 
    * - git
      - latest
      - 克隆Fabric代码仓库
    * - wget
      - latest
      - 下载压缩包
    * - build-essential
      - latest
      - 编译Fabric二进制文件时使用
    * - fabric
      - 2.2
      - Hyperledger Fabric
    * - fabric-ca
      - 1.4.8
      - 完成容器运行
    * - golang
      - go1.19 or higher
      - 编译Fabric文件、执行链码

其中docker、docker-compose、git、wget、build-essential使用apt完成安装，fabric、fabric-ca可以通过直接下载二进制文件解压或本地编译两种方式完成安装，golang采用直接下载二进制文件解压完成安装。

fabric、fabric-ca存储在fabGo/bin目录下，go存储在fabGo/go/bin目录下。

.. code-block:: sh
   :caption: Go Env
   :name: Go Env
   :linenos:

    go env -w GO111MODULE=on
    go env -w GOPROXY=https://goproxy.cn,direct

部分网络无法正常完成Go Modules下载，需要设置代理。

.. list-table:: Docker Image List
    :widths: 30 30
    :header-rows: 1

    * - Image
      - Version
    * - ca
      - 1.4.8
    * - baseos
      - 2.2
    * - ccenv
      - 2.2
    * - orderer
      - 2.2
    * - peer
      - 2.2
    * - tools
      - 2.2

进行网络部署前首先需要完成对上述配置的检查。
