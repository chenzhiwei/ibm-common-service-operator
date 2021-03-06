//
// Copyright 2020 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package size

const Small = `
- name: ibm-cert-manager-operator
  spec:
    certManager:
      certManagerCAInjector:
        resources:
          limits:
            cpu: 35m
            memory: 290Mi
          requests:
            cpu: 30m
            memory: 230Mi
      certManagerController:
        resources:
          limits:
            cpu: 30m
            memory: 230Mi
          requests:
            cpu: 10m
            memory: 175Mi
      certManagerWebhook:
        resources:
          limits:
            cpu: 35m
            memory: 40Mi
          requests:
            cpu: 15m
            memory: 30Mi
      configMapWatcher:
        resources:
          limits:
            cpu: 10m
            memory: 15Mi
          requests:
            cpu: 10m
            memory: 15Mi
- name: ibm-mongodb-operator
  spec: {}
- name: ibm-iam-operator
  spec:
    authentication:
      replicas: 1
      auditService:
        resources:
          limits:
            cpu: 1m
            memory: 10Mi
          requests:
            cpu: 1m
            memory: 10Mi
      authService:
        resources:
          limits:
            cpu: 650m
            memory: 555Mi
          requests:
            cpu: 140m
            memory: 525Mi
      clientRegistration:
        resources:
          limits:
            cpu: 5m
            memory: 5Mi
          requests:
            cpu: 5m
            memory: 5Mi
      identityManager:
        resources:
          limits:
            cpu: 35m
            memory: 160Mi
          requests:
            cpu: 10m
            memory: 120Mi
      identityProvider:
        resources:
          limits:
            cpu: 160m
            memory: 195Mi
          requests:
            cpu: 80m
            memory: 130Mi
    oidcclientwatcher:
      replicas: 1
      resources:
        limits:
          cpu: 15m
          memory: 25Mi
        requests:
          cpu: 10m
          memory: 20Mi
    pap:
      auditService:
        resources:
          limits:
            cpu: 1m
            memory: 10Mi
          requests:
            cpu: 1m
            memory: 10Mi
      papService:
        resources:
          limits:
            cpu: 15m
            memory: 330Mi
          requests:
            cpu: 5m
            memory: 160Mi
      replicas: 1
    policycontroller:
      replicas: 1
      resources:
        limits:
          cpu: 20m
          memory: 30Mi
        requests:
          cpu: 20m
          memory: 20Mi
    policydecision:
      auditService:
        resources:
          limits:
            cpu: 1m
            memory: 10Mi
          requests:
            cpu: 1m
            memory: 10Mi
      resources:
        limits:
          cpu: 30m
          memory: 15Mi
        requests:
          cpu: 20m
          memory: 10Mi
      replicas: 1
    secretWatcher:
      resources:
        limits:
          cpu: 20m
          memory: 145Mi
        requests:
          cpu: 10m
          memory: 120Mi
      replicas: 1
    securityonboarding:
      replicas: 1
      resources:
        limits:
          cpu: 1m
          memory: 1Mi
        requests:
          cpu: 1m
          memory: 1Mi
      iamOnboarding:
        resources:
          limits:
            cpu: 20m
            memory: 1024Mi
          requests:
            cpu: 20m
            memory: 64M
- name: ibm-management-ingress-operator
  spec:
    managementIngress:
      replicas: 1
      resources:
        requests:
          cpu: 25m
          memory: 100Mi
        limits:
          cpu: 30m
          memory: 170Mi
- name: ibm-ingress-nginx-operator
  spec:
    nginxIngress:
      ingress:
        replicas: 1
        resources:
          requests:
            cpu: 10m
            memory: 140Mi
          limits:
            cpu: 10m
            memory: 225Mi
      defaultBackend:
        replicas: 1
        resources:
          requests:
            cpu: 1m
            memory: 10Mi
          limits:
            cpu: 1m
            memory: 10Mi
      kubectl:
        resources:
          requests:
            memory: "1Mi"
            cpu: "1m"
          limits:
            memory: "1Mi"
            cpu: "1m"
- name: ibm-metering-operator
  spec:
    metering:
      dataManager:
        dm:
          resources:
            limits:
              cpu: 450m
              memory: 850Mi
            requests:
              cpu: 200m
              memory: 140Mi
      reader:
        rdr:
          resources:
            limits:
              cpu: 30m
              memory: 200Mi
            requests:
              cpu: 25m
              memory: 175Mi
    meteringSender:
      replicas: 1
      sender:
        resources: {}
    meteringReportServer:
      reportServer:
        resources:
          limits:
            cpu: 50m
            memory: 50Mi
          requests:
            cpu: 50m
            memory: 50Mi
    meteringUI:
      replicas: 1
      ui:
        resources:
          limits:
            cpu: 100m
            memory: 256Mi
          requests:
            cpu: 50m
            memory: 100Mi
- name: ibm-licensing-operator
  spec:
    IBMLicensing:
      resources:
        requests:
          cpu: 10m
          memory: 220Mi
        limits:
          cpu: 35m
          memory: 250Mi
    IBMLicenseServiceReporter:
      databaseContainer:
        resources:
          requests:
            cpu: 200m
            memory: 256Mi
          limits:
            cpu: 300m
            memory: 300Mi
      receiverContainer:
        resources:
          requests:
            cpu: 200m
            memory: 256Mi
          limits:
            cpu: 300m
            memory: 300Mi
- name: ibm-commonui-operator
  spec:
    commonWebUI:
      replicas: 1
      resources:
        requests:
          memory: "105Mi"
          limits: "20m"
        limits:
          memory: "200Mi"
          limits: "105m"
- name: ibm-platform-api-operator
  spec:
    platformApi:
      auditService:
        resources:
          limits:
            cpu: 25m
            memory: 15Mi
          requests:
            cpu: 25m
            memory: 10Mi
      platformApi:
        resources:
          limits:
            cpu: 25m
            memory: 25Mi
          requests:
            cpu: 25m
            memory: 25Mi
      replicas: 1
- name: ibm-catalog-ui-operator
  spec:
    catalogUI:
      catalogui:
        resources:
          limits:
            cpu: 190m
            memory: 220Mi
          requests:
            cpu: 35m
            memory: 105Mi
      replicaCount: 1
- name: ibm-healthcheck-operator
  spec:
    healthService:
      memcached:
        replicas: 1
        resources:
          requests:
            memory: "10Mi"
            cpu: "1m"
          limits:
            memory: "15Mi"
            cpu: "1m"
      healthService:
        replicas: 1
        resources:
          requests:
            memory: "35Mi"
            cpu: "5m"
          limits:
            memory: "40Mi"
            cpu: "10m"
- name: ibm-auditlogging-operator
  spec:
    auditLogging:
      fluentd:
        resources:
          requests:
            cpu: 25m
            memory: 100Mi
          limits:
            cpu: 35m
            memory: 150Mi
- name: ibm-monitoring-exporters-operator
  spec:
    exporter:
      collectd:
        resources:
          requests:
            cpu: 1m
            memory: 15Mi
          limits:
            cpu: 1m
            memory: 15Mi
        routerResource:
          limits:
            cpu: 25m
            memory: 20Mi
          requests:
            cpu: 10m
            memory: 20Mi
      nodeExporter:
        resources:
          requests:
            cpu: 5m
            memory: 20Mi
          limits:
            cpu: 5m
            memory: 25Mi
        routerResource:
          requests:
            cpu: 50m
            memory: 128Mi
          limits:
            cpu: 100m
            memory: 256Mi
      kubeStateMetrics:
        resources:
          requests:
            cpu: 500m
            memory: 110Mi
          limits:
            cpu: 540m
            memory: 160Mi
        routerResource:
          limits:
            cpu: 25m
            memory: 20Mi
          requests:
            cpu: 10m
            memory: 20Mi
- name: ibm-monitoring-grafana-operator
  spec:
    grafana:
      resources:
        grafanaConfig:
          resources:
            requests:
              cpu: 20m
              memory: 65Mi
            limitis:
              cpu: 70m
              memory: 75Mi
        dashboardConfig:
          resources:
            requests:
              cpu: 5m
              memory: 45Mi
            limits:
              cpu: 10m
              memory: 60Mi
        routerConfig:
          resources:
            requests:
              cpu: 10m
              memory: 20Mi
            limits:
              cpu: 25m
              memory: 25Mi
- name: ibm-monitoring-prometheusext-operator
  spec:
    prometheusExt:
      prometheusConfig:
        resources:
          requests:
            cpu: 65m
            memory: 1920Mi
          limits:
            cpu: 130m
            memory: 2570Mi
      alertManagerConfig:
        resources:
          requests:
            cpu: 5m
            memory: 20Mi
          limits:
            cpu: 5m
            memory: 25Mi
      mcmMonitor:
        resources:
          requests:
            cpu: 1m
            memory: 15Mi
          limits:
            cpu: 1m
            memory: 15Mi
- name: ibm-elastic-stack-operator
  spec:
    elasticStack:
      curator:
        resources:
          limits:
            memory: 915Mi
          requests:
            memory: 2320Mi
        routerImage:
          resources:
            limits:
              memory: 256Mi
            requests:
              memory: 64Mi
      filebeat:
        resources:
          limits:
            memory: 80Mi
          requests:
            memory: 45Mi
      logstash:
        probe:
          resources:
            limits:
              memory: 1290Mi
            requests:
              memory: 810Mi
        replicas: 1
`
