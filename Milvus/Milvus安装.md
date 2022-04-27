###Helm方式安装
```yaml
添加chart仓库

helm repo add milvus https://milvus-io.github.io/milvus-helm/
helm repo update
```

```yaml
集群方式部署  
helm install milvus milvus/milvus \
--set attu.enabled=true \
--set ingress.enabled=true \
--set metrics.serviceMonitor.enabled=true \
--set log.format=json \
--set log.persistence.storageClass=#storageClassname \
--set etcd.persistence.storageClass=#storageClassname \
--set kafka.persistence.storageClass=#storageClassname \
--set minio.persistence.storageClass=#storageClassname \
--set pulsar.zookeeper.volumes.data.storageClassName=#storageClassname \
--set pulsar.bookkeeper.volumes.journal.storageClassName=#storageClassname \
--set pulsar.bookkeeper.volumes.ledgers.storageClassName=#storageClassname \
--set pulsar.bookkeeper.volumes.common.storageClassName=#storageClassname \
--set pulsar.prometheus.volumes.data.storageClassName=#storageClassname \
```

```yaml
单机方式部署

helm install my-milvus milvus/milvus --version 1.1.6 -n milvus-1-1 
    --set cache.insertBufferSize=5G \
    --set cache.cacheSize=45 \
    --set logs.logToStdOut=true \
    --set persistence.enabled=true \
    --set attu.enabled=true \
    --set metrics.serviceMonitor.enabled=true \
    --set persistence.persistentVolumeClaim.storageClass=#storageClassname milvus存储 \
    --set metrics.enabled=true \
    --set mysql.persistence.storageClass=#storageClassname数据库存储
```


>引用文档
>>https://artifacthub.io/packages/helm/milvus/milvus
>>https://milvus.io/cn/docs/v2.0.x/attu_install-helm.md

