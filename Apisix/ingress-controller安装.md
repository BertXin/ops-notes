###Heml方式安装

```shell
添加chart仓库

helm repo add apisix https://charts.apiseven.com
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
```

```shell
执行安装命令

helm install apisix apisix/apisix \
  --set gateway.type=LoadBalancer \
  --set ingress-controller.enabled=true \
  --set serviceMonitor.enabled=true \
  --set etcd.persistence.storageClass="apisix-etcd-new-sc" \
  --set etcd.persistence.size="20Gi" \
  --set etcd.startFromSnapshot.enabled=true \
  --set etcd.startFromSnapshot.existingClaim=etcd-snapshot-new-pvc \
  --set etcd.startFromSnapshot.snapshotFilename=my.db \
  --namespace apisix-new \
  --set ingress-controller.config.apisix.serviceNamespace=apisix-new
```
###Etcd数据迁移
```shell
helm install apisix-etcd bitnami/etcd \
 --set persistence.storageClass="apisix-etcd-new-sc" \
 --set persistence.size="20Gi" \
 --set 'replicaCount=3' \
 --set startFromSnapshot.enabled=true \
 --set startFromSnapshot.existingClaim=etcd-snapshot-new-pvc \
 --set startFromSnapshot.snapshotFilename=my.db \
 --namespace apisix-new
```

###创建路由模版
```yaml 
apiVersion: apisix.apache.org/v2beta2
kind: ApisixRoute
metadata:
  name: #名称
  namespace: ml
spec:
  http:
  - name: #名称
    match:
      hosts:
      - #域名
      paths:
      - /* #匹配规则
    plugins:
       - name: proxy-rewrite #插件
         enable: true
         config:
           regex_uri: ["^/(.*)","/$1"] #插件规则
    backends:
       - serviceName: #后端svc
         servicePort: #后端svc端口
```

>引用文档
>>https://apisix.apache.org/zh/docs/ingress-controller/deployments/azure/
>>https://apisix.apache.org/zh/docs/ingress-controller/concepts/apisix_route
>>https://artifacthub.io/packages/helm/bitnami/etcd
