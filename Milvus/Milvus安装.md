###Helm方式安装
```shell
添加chart仓库

helm repo add milvus https://milvus-io.github.io/milvus-helm/
helm repo update
```

```shell
集群方式部署

helm install my-release milvus/milvus \
  --set metrics.serviceMonitor.enabled=true \
  --set log.format=json \
  --set log.persistence.enabled=true \
  --set attu.enabled=true \
  --set log.persistence.persistentVolumeClaim.storageClass=#storageClassname \ 
  --set standalone.persistence.persistentVolumeClaim.storageClass=#storageClassname 
```

```shell
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

