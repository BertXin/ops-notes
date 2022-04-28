```yaml
helm install drone community-charts/drone \
  --set server.persistentVolume.storageClass=gitops-sc \
  --set server.env.DRONE_GITEA_SERVER=https://try.gitea.io \
  --set server.env.DRONE_GITEA_CLIENT_ID=aa69a818-e8f7-4def-aa69-983f783b594f \
  --set server.env.DRONE_GITEA_CLIENT_SECRET=gto_wjf6ypehmdmzmfhqcfywvdudjotvygmiheoves2iqsr2kdfybn3a \
  --set server.env.DRONE_RPC_SECRET=bea26a2221fd8090ea38720fc445eca6 \
  --set server.env.DRONE_SERVER_HOST=drone.qingteng-inc.com \
  --set server.env.DRONE_SERVER_PROTO=http \
    -n gitops
 ```
