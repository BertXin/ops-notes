```yaml
helm install drone community-charts/drone \
  --set server.persistentVolume.storageClass=#存储 \
  --set server.env.DRONE_GITEA_SERVER=https://try.gitea.io \
  --set server.env.DRONE_GITEA_CLIENT_ID=#gitID \
  --set server.env.DRONE_GITEA_CLIENT_SECRET=#gitsecret \
  --set server.env.DRONE_RPC_SECRET=#密钥 \
  --set server.env.DRONE_SERVER_HOST=#域名\
  --set server.env.DRONE_SERVER_PROTO=http \
    -n gitops
 ```
```yaml
helm install my-release argo/argo-cd -n argocd --set server.service.type=LoadBalancer
```