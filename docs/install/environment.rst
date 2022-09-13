Hyperledger Fabric 基础环境安装
---------------------------------------

安装Hyperledger Fabric框架2.2版本运行环境包括如下内容

.. list-table:: Environment
    :widths: 15 15 30
    :header-rows: 1

    * - Software
      - Version
      - Description
    * - Docker
      - latest
      - 提供容器运行环境
    * - Docker-Compose
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
    * - Fabric
      - 2.2
      - Hyperledger Fabric
    * - Fabric-CA
      - 1.4.8
      - 完成容器运行
    * - Golang
      - go1.19 or higher
      - 编译Fabric文件、执行链码

.. list-table:: Image List
    :widths: 15 15
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
