# Hyperledger Fabric

[toc]

## Reference

```markdown
https://juejin.cn/user/1042770075321405
```

## Prepare

`FABRIC_CFG_PATH` 环境变量指向的路径 默认地址`/etc/hyperledger/fabric`

| Item | Path | Introduction |
| - | - | - |
| peer | $FABRIC_CFG_PATH/core.yaml | 指定peer节点运行时参数 |
| orderer | $FABRIC_CFG_PATH/orderer.yaml | 指定order二节点运行时参数 |
| fabric-ca | $FABRIC_CFG_PATH/fabric-ca-server-config.yaml | 指定CA节点运行时参数 |
| configtxgen | $FABRIC_CFG_PATH/configtx.yaml | 指定configtxgen命令运行时参数 |

### configtx.yaml

Transaction(Tx) 相关的配置，如应用通道、锚节点、排序服务等，均在configtx.yaml文件中配置。六个部分，其中前五个部分被第六个Profiles引用，Profiles引用其他部分时可能覆盖特定配置。

#### Organizations

Membership Service Provider (MSP) 是一个组织身份标识，在Fabric中用统一的MSPID标识。

```yaml
Organizations:
    - &councilMSP           # 定义一个组织引用，类似于变量，可在Profile部分被引用；所有带 & 符号的都是引用变量，使用 * 来引用
        Name: councilMSP    # 组织名称
        ID: councilMSP      # 组织ID
        MSPDir: ../orgs/council.ifantasy.net/msp    # 组织MSP文件夹的路径
        Policies:           # 组织策略
            Readers:
                Type: Signature
                Rule: "OR('councilMSP.member')"
            Writers:
                Type: Signature
                Rule: "OR('councilMSP.member')"
            Admins:
                Type: Signature
                Rule: "OR('councilMSP.admin')"
        # 此文件内的Orderer端口皆为容器内端口
        OrdererEndpoints:   # 定义排序节点（可多个），客户端和对等点可以分别连接到这些orderer以推送transactions和接收区块。
            - "orderer1.council.ifantasy.net:7051"
            - "orderer2.council.ifantasy.net:7054"
            - "orderer3.council.ifantasy.net:7057"
        AnchorPeers:    # 定义锚节点，锚节点对外代表本组织通信
            - Host: peer1.soft.ifantasy.net
              Port: 7251
```

#### Capabilities

定义了Fabric网络加入时必须支持的特性，定义通道能力，不满足要求的无法处理交易。

```yaml
Capabilities:
    # Channel配置同时针对通道上的Orderer节点和Peer节点(设置为ture表明要求节点具备该能力)；
    Channel: &ChannelCapabilities
        V2_0: true  # 要求Channel上的所有Orderer节点和Peer节点达到v2.0.0或更高版本
     # Orderer配置仅针对Orderer节点，不限制Peer节点
    Orderer: &OrdererCapabilities
        V2_0: true  # 要求所有Orderer节点升级到v2.0.0或更高版本
    # Application配置仅应用于对等网络，不需考虑排序节点的升级
    Application: &ApplicationCapabilities
        V2_0: true
```

#### Application

定义应用内访问控制策略和参与组织。

```yaml
Application: &ApplicationDefaults

    # 干预 创建链码的系统链码 的函数访问控制策略
    _lifecycle/CheckCommitReadiness: /Channel/Application/Writers       # CheckCommitReadiness 函数的访问策略
    _lifecycle/CommitChaincodeDefinition: /Channel/Application/Writers  # CommitChaincodeDefinition 函数的访问策略
    _lifecycle/QueryChaincodeDefinition: /Channel/Application/Writers   # QueryChaincodeDefinition 函数的访问策略
    _lifecycle/QueryChaincodeDefinitions: /Channel/Application/Writers  # QueryChaincodeDefinitions 函数的访问策略

    # 关于 生命周期系统链码（lscc） 的函数访问控制策略
    lscc/ChaincodeExists: /Channel/Application/Readers              # getid 函数的访问策略
    lscc/GetDeploymentSpec: /Channel/Application/Readers            # getdepspec 函数的访问策略
    lscc/GetChaincodeData: /Channel/Application/Readers             # getccdata 函数的访问策略
    lscc/GetInstantiatedChaincodes: /Channel/Application/Readers    # getchaincodes 函数的访问策略

    # 关于 查询系统链码（qscc） 的函数访问控制策略
    qscc/GetChainInfo: /Channel/Application/Readers         # GetChainInfo 函数的访问策略
    qscc/GetBlockByNumber: /Channel/Application/Readers     # GetBlockByNumber 函数的访问策略
    qscc/GetBlockByHash: /Channel/Application/Readers       # GetBlockByHash 函数的访问策略
    qscc/GetTransactionByID: /Channel/Application/Readers   # GetTransactionByID 函数的访问策略
    qscc/GetBlockByTxID: /Channel/Application/Readers       # GetBlockByTxID 函数的访问策略

    # 关于 配置系统链码（cscc） 的函数访问控制策略
    cscc/GetConfigBlock: /Channel/Application/Readers   # GetConfigBlock 函数的访问策略
    cscc/GetChannelConfig: /Channel/Application/Readers # GetChannelConfig 函数的访问策略
    
    # 关于 peer 节点的函数访问控制策略
    peer/Propose: /Channel/Application/Writers                  # Propose 函数的访问策略
    peer/ChaincodeToChaincode: /Channel/Application/Writers     # ChaincodeToChaincode 函数的访问策略

    # 关于事件资源的访问策略
    event/Block: /Channel/Application/Readers           # 发送区块事件的策略
    event/FilteredBlock: /Channel/Application/Readers   # 发送筛选区块事件的策略
    
    # 默认为空，在 Profiles 中定义
    Organizations:
    # 定义本层级的应用控制策略，路径为 /Channel/Application/<PolicyName>
    Policies:
        Readers:    # /Channel/Application/Readers
            Type: ImplicitMeta
            Rule: "ANY Readers"
        Writers:
            Type: ImplicitMeta
            Rule: "ANY Writers"
        Admins:
            Type: ImplicitMeta
            Rule: "MAJORITY Admins"
        LifecycleEndorsement:
            Type: ImplicitMeta
            Rule: "MAJORITY Endorsement"
        Endorsement:
            Type: ImplicitMeta
            Rule: "MAJORITY Endorsement"

    Capabilities:
        <<: *ApplicationCapabilities    # 引用上节 Capabilities 的 ApplicationCapabilities
```

#### Orderer

定义排序服务的参数，创建创世区块和交易。

```yaml
Orderer: &OrdererDefaults
  OrdererType: etcdraft   # 排序服务算法，目前可用：solo，kafka，etcdraft
  Addresses:              # 排序节点地址
    - orderer1.soft.ifantasy.net:7051
    - orderer2.web.ifantasy.net:7052
    - orderer3.hard.ifantasy.net:7053
  # 定义了 etcdRaft 排序类型被选择时的配置
  EtcdRaft:
    Consenters:         # 定义投票节点
      - Host: orderer1.council.ifantasy.net
        Port: 7051
        ClientTLSCert: ../orgs/council.ifantasy.net/registers/orderer1/tls-msp/signcerts/cert.pem # 节点的TLS签名证书
        ServerTLSCert: ../orgs/council.ifantasy.net/registers/orderer1/tls-msp/signcerts/cert.pem
      - Host: orderer2.council.ifantasy.net
        Port: 7054
        ClientTLSCert: ../orgs/council.ifantasy.net/registers/orderer2/tls-msp/signcerts/cert.pem
        ServerTLSCert: ../orgs/council.ifantasy.net/registers/orderer2/tls-msp/signcerts/cert.pem
      - Host: orderer3.council.ifantasy.net
        Port: 7057
        ClientTLSCert: ../orgs/council.ifantasy.net/registers/orderer3/tls-msp/signcerts/cert.pem
        ServerTLSCert: ../orgs/council.ifantasy.net/registers/orderer3/tls-msp/signcerts/cert.pem

  # 区块打包的最大超时时间 (到了该时间就打包区块)
  BatchTimeout: 2s
  # 区块链的单个区块配置（orderer端切分区块的参数）
  BatchSize:
    MaxMessageCount: 10         # 一个区块里最大的交易数
    AbsoluteMaxBytes: 99 MB     # 一个区块的最大字节数，任何时候都不能超过
    PreferredMaxBytes: 512 KB   # 一个区块的建议字节数，如果一个交易消息的大小超过了这个值, 就会被放入另外一个更大的区块中

  # 参与维护Orderer的组织，默认为空（通常在 Profiles 中再配置）
  Organizations:
  # 定义本层级的排序节点策略，其权威路径为 /Channel/Orderer/<PolicyName>
  Policies:
    Readers:    # /Channel/Orderer/Readers
      Type: ImplicitMeta
      Rule: "ANY Readers"
    Writers:
      Type: ImplicitMeta
      Rule: "ANY Writers"
    Admins:
      Type: ImplicitMeta
      Rule: "MAJORITY Admins"
    BlockValidation:    # 指定了哪些签名必须包含在区块中，以便peer节点进行验证
      Type: ImplicitMeta
      Rule: "ANY Writers"
  Capabilities:
    <<: *OrdererCapabilities    # 引用上节 Capabilities 的 OrdererCapabilities ```

#### Channel

定义创世区块和配置交易的通道参数。

```yaml
Channel: &ChannelDefaults
    #   定义本层级的通道访问策略，其权威路径为 /Channel/<PolicyName>
    Policies:
        Readers:    # 定义谁可以调用 'Deliver' 接口
            Type: ImplicitMeta
            Rule: "ANY Readers"
        Writers:    # 定义谁可以调用 'Broadcast' 接口
            Type: ImplicitMeta
            Rule: "ANY Writers"
        # By default, who may modify elements at this config level
        Admins:     # 定义谁可以修改本层策略
            Type: ImplicitMeta
            Rule: "MAJORITY Admins"

    Capabilities:
        <<: *ChannelCapabilities        # 引用上节 Capabilities 的 ChannelCapabilities 
```

#### Profiles

用于configtxgen工具配置入口，定义配欸之模板，模块包括`Application`, `Capabilities`, `Consortium`, `Consortiums`, `Policies`, `Orderer`等配置字段。

```yaml
Profiles:
    # OrgsChannel用来生成channel配置信息，名字可以任意
    # 需要包含Consortium和Applicatioon两部分。
    OrgsChannel:
        Consortium: SampleConsortium    # 通道所关联的联盟名称
        <<: *ChannelDefaults
        Orderer:
            <<: *OrdererDefaults
            Organizations:
                - *councilMSP
            Capabilities: *OrdererCapabilities
        Application:
            <<: *ApplicationDefaults
            Organizations:
                - *softMSP
                - *webMSP
                - *hardMSP
            Capabilities:
                <<: *ApplicationCapabilities
```

### Peer 配置

按照优先级顺序，尝试命令行参数、环境变量和配置文件的配置信息。环境变量读取信息时出日志`FABRIC_LOGGING_SPEC`环境变量单独指定外，其余以`CORE_`开头，如配置文件`peer.id`配置为`CORE_PEER_ID`。

默认路径按照顺序读取`$FABRIC_CFG_PATH/core.yaml`, `./core.yaml`, `/etc/hyperledger/fabric/core.yaml`。

Peer配置文件由`peer`, `vm`, `chaincode`, `ledger`, `operations`, `metrics`六大部分组成。

docker容器配置文件

```yaml
  peer-base:
    image: hyperledger/fabric-peer:${FABRIC_BASE_VERSION}
    environment:
      - FABRIC_LOGGING_SPEC=info
      - CORE_PEER_ID=peer1.soft.ifantasy.net
      - CORE_PEER_LISTENADDRESS=0.0.0.0:7251
      - CORE_PEER_ADDRESS=peer1.soft.ifantasy.net:7251
      - CORE_PEER_LOCALMSPID=softMSP
      - CORE_PEER_MSPCONFIGPATH=${DOCKER_CA_PATH}/peer/msp
      - CORE_PEER_TLS_ENABLED=true
      - CORE_PEER_TLS_CERT_FILE=${DOCKER_CA_PATH}/peer/tls-msp/signcerts/cert.pem
      - CORE_PEER_TLS_KEY_FILE=${DOCKER_CA_PATH}/peer/tls-msp/keystore/key.pem
      - CORE_PEER_TLS_ROOTCERT_FILE=${DOCKER_CA_PATH}/peer/tls-msp/tlscacerts/tls-council-ifantasy-net-7050.pem
      - CORE_PEER_GOSSIP_USELEADERELECTION=true
      - CORE_PEER_GOSSIP_ORGLEADER=false
      - CORE_PEER_GOSSIP_EXTERNALENDPOINT=peer1.soft.ifantasy.net:7251
      - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock
      - CORE_VM_DOCKER_HOSTCONFIG_NETWORKMODE=${COMPOSE_PROJECT_NAME}_${DOCKER_NETWORKS}

    working_dir: ${DOCKER_CA_PATH}/peer
    volumes:
      - /var/run:/host/var/run
    networks:
      - ${DOCKER_NETWORKS}

#FABRIC_LOGGING_SPEC ：指定日志级别
#CORE_PEER_ID ： Peer 在网络中的 ID 信息，用于辨识不同的节点
#CORE_PEER_LISTENADDRESS ：服务监听的本地地址，本地有多个网络接口时可指定仅监听某个接口
#CORE_PEER_ADDRESS ：对同组织内其他节点的监听连接地址。当服务在NAT设备上运行时，该配置可以指定服务对外宣称的可访问地址。如果是客户端，则作为其连接的 Peer 服务地址
#CORE_PEER_LOCALMSPID ：Peer 所关联的 MSPID ，一般为所属组织名称，需要与通道配置内名称一致
#CORE_PEER_MSPCONFIGPATH ：MSP 目录所在的路径，可以为绝对路径，或相对配置目录的路径
#CORE_PEER_TLS_ENABLED ：是否开启 server 端 TLS 检查
#CORE_PEER_TLS_CERT_FILE ：server 端使用的 TLS 证书路径
#CORE_PEER_TLS_KEY_FILE ：server 端使用的 TLS 私钥路径
#CORE_PEER_TLS_ROOTCERT_FILE ：server 端使用的根CA的证书，签发服务端的 TLS证书
#CORE_PEER_GOSSIP_USELEADERELECTION ：是否允许节点之间动态进行组织的代表（leader）节点选举，通常情况下推荐开启
#CORE_PEER_GOSSIP_ORGLEADER ：本节点是否指定为组织的代表节点，与useLeaderElection不能同时指定为true
#CORE_PEER_GOSSIP_EXTERNALENDPOINT ：节点向组织外节点公开的服务地址，默认为空，代表不被其他组织所感知
#CORE_VM_ENDPOINT ：docker daemon 的地址
#CORE_VM_DOCKER_HOSTCONFIG_NETWORKMODE ：运行链码容器的网络
#
```

### Orderer 配置

按照优先级顺序，尝试命令行参数、环境变量和配置文件的配置信息。以`ORDERER_`开头。

默认路径按照顺序读取`$FABRIC_CFG_PATH/orderer.yaml`, `./orderer.yaml`, `/etc/hyperledger/fabric/orderer.yaml`。

Peer配置文件由`General`, `FileLedger`, `RAMLedger`, `Kafka`, `Debug`, `Operations`, `Metrics`, `Consensus`八大部分组成。

docker容器配置文件

```yaml
  orderer-base:
    image: hyperledger/fabric-orderer:${FABRIC_BASE_VERSION}
    environment:
      - ORDERER_HOME=${DOCKER_CA_PATH}/orderer
      - ORDERER_HOST=orderer1.council.ifantasy.net
      - ORDERER_GENERAL_LOCALMSPID=councilMSP
      - ORDERER_GENERAL_LISTENPORT=7051
      - ORDERER_GENERAL_LISTENADDRESS=0.0.0.0
      - ORDERER_GENERAL_BOOTSTRAPMETHOD=none
      - ORDERER_CHANNELPARTICIPATION_ENABLED=true
      # - ORDERER_GENERAL_GENESISMETHOD=file
      # - ORDERER_GENERAL_GENESISFILE=${DOCKER_CA_PATH}/orderer/genesis.block
      - ORDERER_GENERAL_LOCALMSPDIR=${DOCKER_CA_PATH}/orderer/msp
      - ORDERER_GENERAL_LOGLEVEL=debug
      - ORDERER_GENERAL_TLS_ENABLED=true
      - ORDERER_GENERAL_TLS_CERTIFICATE=${DOCKER_CA_PATH}/orderer/tls-msp/signcerts/cert.pem
      - ORDERER_GENERAL_TLS_PRIVATEKEY=${DOCKER_CA_PATH}/orderer/tls-msp/keystore/key.pem
      - ORDERER_GENERAL_TLS_ROOTCAS=[${DOCKER_CA_PATH}/orderer/tls-msp/tlscacerts/tls-council-ifantasy-net-7050.pem]
      - ORDERER_GENERAL_CLUSTER_CLIENTCERTIFICATE=${DOCKER_CA_PATH}/orderer/tls-msp/signcerts/cert.pem
      - ORDERER_GENERAL_CLUSTER_CLIENTPRIVATEKEY=${DOCKER_CA_PATH}/orderer/tls-msp/keystore/key.pem
      - ORDERER_GENERAL_CLUSTER_ROOTCAS=[${DOCKER_CA_PATH}/orderer/tls-msp/tlscacerts/tls-council-ifantasy-net-7050.pem]
      - ORDERER_ADMIN_TLS_ENABLED=true
      - ORDERER_ADMIN_TLS_CERTIFICATE=${DOCKER_CA_PATH}/orderer/tls-msp/signcerts/cert.pem
      - ORDERER_ADMIN_TLS_PRIVATEKEY=${DOCKER_CA_PATH}/orderer/tls-msp/keystore/key.pem
      - ORDERER_ADMIN_TLS_ROOTCAS=[${DOCKER_CA_PATH}/orderer/tls-msp/tlscacerts/tls-council-ifantasy-net-7050.pem]
      - ORDERER_ADMIN_TLS_CLIENTROOTCAS=[${DOCKER_CA_PATH}/orderer/tls-msp/tlscacerts/tls-council-ifantasy-net-7050.pem]
      - ORDERER_ADMIN_LISTENADDRESS=0.0.0.0:8888
      - ORDERER_METRICS_PROVIDER=prometheus
      - ORDERER_OPERATIONS_LISTENADDRESS=0.0.0.0:9999
      - ORDERER_DEBUG_BROADCASTTRACEDIR=data/logs
    networks:
      - ${DOCKER_NETWORKS}
#ORDERER_HOME ：orderer 运行的根目录
#ORDERER_HOST ：orderer 运行的主机
#ORDERER_GENERAL_LOCALMSPID ： orderer 所关联的 MSPID ，一般为所属组织名称，需要与通道配置内名称一致
#ORDERER_GENERAL_LISTENPORT ：服务绑定的监听端口
#ORDERER_GENERAL_LISTENADDRESS ：服务绑定的监听地址，一般需要指定为所服务的特定网络接口的地址或全网（0.0.0.0）
#ORDERER_GENERAL_BOOTSTRAPMETHOD ：获取引导块的方法，2.x版本中仅支持file或none
#ORDERER_CHANNELPARTICIPATION_ENABLED ：是否提供参与通道的 API
#ORDERER_GENERAL_GENESISMETHOD ：当 ORDERER_GENERAL_BOOTSTRAPMETHOD 为 file 时启用，指定创世区块类型
#ORDERER_GENERAL_GENESISFILE ：指定创世区块位置
#ORDERER_GENERAL_LOCALMSPDIR ：本地 MSP 文件路径
#ORDERER_GENERAL_LOGLEVEL ：日志级别
#ORDERER_GENERAL_TLS_ENABLED ：启用TLS时的相关配置
#ORDERER_GENERAL_TLS_CERTIFICATE ：Orderer 身份证书
#ORDERER_GENERAL_TLS_PRIVATEKEY ：Orderer 签名私钥
#ORDERER_GENERAL_TLS_ROOTCAS ：信任的根证书
#ORDERER_GENERAL_CLUSTER_CLIENTCERTIFICATE ：双向TLS认证时，作为客户端证书的文件路径，如果没设置会使用 TLS.Certificate
#ORDERER_GENERAL_CLUSTER_CLIENTPRIVATEKEY ：双向TLS认证时，作为客户端私钥的文件路径，如果没设置会使用 TLS.PrivateKey
#ORDERER_GENERAL_CLUSTER_ROOTCAS ：信任的根证书
#ORDERER_ADMIN_TLS_ENABLED ：是否启用 orderer 的管理服务面板
#ORDERER_ADMIN_TLS_CERTIFICATE ：管理服务的证书
#ORDERER_ADMIN_TLS_PRIVATEKEY ：管理服务的私钥
#ORDERER_ADMIN_TLS_ROOTCAS ：管理服务的可信根证书
#ORDERER_ADMIN_TLS_CLIENTROOTCAS ：管理服务客户端的可信根证书
#ORDERER_ADMIN_LISTENADDRESS ：管理服务监听地址
#ORDERER_METRICS_PROVIDER ：统计服务类型，可以为statsd(推送模式)，prometheus(拉取模式)，disabled
#ORDERER_OPERATIONS_LISTENADDRESS ：RESTful 管理服务的监听地址
#ORDERER_DEBUG_BROADCASTTRACEDIR ：广播请求的追踪路径

```
