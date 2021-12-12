# 模块10作业

## 作业要求

- [x] 为 HTTPServer 添加 0-2 秒的随机延时
- [x] 为 HTTPServer 项目添加延时 Metric
- [x] 将 HTTPServer 部署至测试集群，并完成 Prometheus 配置
- [x] 从 Promethus 界面中查询延时指标数据
- [x] （可选）创建一个 Grafana Dashboard 展现延时分配情况

## 扩展httpserver增加metric

参见: https://github.com/evanli18/cncf/tree/main/src/github.com/evanli18/cncf/homework/httpserver

## 部署loki

from: https://github.com/cncamp/101/tree/master/module10/loki-stack

## 更新deploy，yaml如下

my deploy:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: app
  template:
    metadata:
      annotations:
        prometheus.io/port: "8080"
        prometheus.io/scrape: "true"
      labels:
        app: app
    spec:
      containers:
      - envFrom:
        - configMapRef:
            name: app-config
        image: evanchn/httpserver:v1.1
        imagePullPolicy: IfNotPresent
        name: app
        readinessProbe:
          failureThreshold: 3
          httpGet:
            path: /healthz
            port: 8080
            scheme: HTTP
          initialDelaySeconds: 30
          periodSeconds: 5
          successThreshold: 2
          timeoutSeconds: 1
        resources:
          limits:
            cpu: 500m
            memory: 200Mi
          requests:
            cpu: 200m
            memory: 100Mi
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      terminationGracePeriodSeconds: 30
```

更新版本为 1.1
增加 prometheus 的 annotations(prometheus.io/port, prometheus.io/scrape)

## 查询一些指标信息

QPS Rate: sum(rate(http_requests_total{}[5m]))

Latency: histogram_quantile(0.9, sum(rate(http_request_duration_seconds_bucket{}[5m])) by (path, le)) * 1000
    分为 0.75, 0.5

样图
![image](./dashboard.png)

模板可导入文件在当前目录下的json文件。
