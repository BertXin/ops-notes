###Gitlab.CI

```yaml
# 定义全局变量,镜像名称，命名空间，镜像拉取的密码的变量名
variables:
  IMAGE_HUB: "仓库名称"
  REPOSITORY: $REGISTRY_DOMAIN/$IMAGE_HUB/$CI_PROJECT_NAME:$CI_PIPELINE_ID-$CI_COMMIT_REF_SLUG-$CI_COMMIT_SHORT_SHA
  KUBECONFIG: /etc/deploy/config
# stages顺序运行, 同一个stage的所有job并行
stages:
  - BuildImage
  - Deploy-Master_K8s
  - Deploy-test_K8s
  
# 这里打印一些变量，仅仅为了展示，看下这些gitlab预设的变量值
before_script:
  - echo "开始构建 build"
  - echo $CI_PIPELINE_ID
  - echo $REPOSITORY
  - echo $CI_COMMIT_REF_SLUG
  - echo $CI_COMMIT_SHORT_SHA
  - echo $KUBE_CONFIG
  - echo $CI_COMMIT_AUTHOR
build_image_job:
  # 构建镜像名
  image: dengtangxiaochou/shell:0.0.2
  # 任务名称
  stage: BuildImage
  script:
    - mkdir -p /kaniko/.docker
    - echo "{\"auths\":{\"$REGISTRY_DOMAIN\":{\"username\":\"$REGISTRY_USERNAME\",\"password\":\"$REGISTRY_PASSWORD\"}}}" > /kaniko/.docker/config.json
    - executor --context $CI_PROJECT_DIR --dockerfile Dockerfile --destination $REPOSITORY --cleanup
deploy_master_k8s_job:
  #部署代码
  image: registry.cn-hangzhou.aliyuncs.com/haoshuwei24/kubectl:1.16.6
  #任务名称
  stage: Deploy-Master_K8s
  script:
    - mkdir -p /etc/deploy
    - cat $KUBE_CONFIG  > $KUBECONFIG
    - kubectl get ns
    - export
    - "curl --location --request POST '钉钉地址' \
--header 'Content-Type: application/json' \
--data-raw '{
     \"msgtype\": \"markdown\",
     \"markdown\": {
         \"title\":\"构建详情\",
         \"text\": \"### 构建任务  \n> **构建项目**：['$CI_PROJECT_NAME']('$CI_PROJECT_URL')  \n> **构建任务ID**：'$CI_PIPELINE_ID'  \n> **CommitID**：'$CI_COMMIT_SHORT_SHA'  \n> **构建分支**：'$CI_COMMIT_REF_SLUG'  \n> **构建状态**：'$CI_JOB_STATUS'  \n> **镜像TAG**： '$REPOSITORY'  \n>**构建时间**：'$CI_PIPELINE_CREATED_AT'  \n> **提交人**：'$GITLAB_USER_EMAIL'\"
     }
 }'"
  only:
    refs:
      - /^master.*$/     
deploy_test_k8s_job:
  #部署代码
  image: registry.cn-hangzhou.aliyuncs.com/haoshuwei24/kubectl:1.16.6
  #任务名称
  stage: Deploy-test_K8s
  script:
    - mkdir -p /etc/deploy
    - cat $KUBE_CONFIG  > $KUBECONFIG
    - kubectl get node
  only:
    refs:
      - /^devops.*$/
```