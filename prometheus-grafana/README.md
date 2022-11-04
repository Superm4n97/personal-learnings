* For configuring and learning basics about prometheus in local machine: https://prometheus.io/docs/prometheus/latest/getting_started/
* Prometheus and Grafana in Kubernetes:
  * Get Helm chart: https://github.com/prometheus-community/helm-charts/tree/main/charts/prometheus
  * Learn about Prometheus and Grafana and other stuff in kubernetes: https://github.com/helm/charts/tree/master/stable/prometheus-operator
  * Learn about grafana: https://github.com/helm/charts/tree/master/stable/grafana
    * Default username: `admin`
    * Default password: `prom-operator`
  * **Youtube:**
    * Fundamentals about Prometheus and Grafana: https://www.youtube.com/watch?v=h4Sl21AKiDg
    * Setup Prometheus and Grafana in Kubernetes cluster using Helm: https://www.youtube.com/watch?v=QoDqxm7ybLc
    * How to monitor third party application (MongoDB) in Kubernetes using Prometheus: https://www.youtube.com/watch?v=mLPg49b33sA

## Notes
* Prometheus is like a server which listen to specific target ports.
* Prometheus tells about the states of different components of an application.
* Prometheus basically consists of three parts:
  1. **Storage:** It is a time series database stores the metric data.
  2. **Data Retrieval Worker:** It is a worker which pulls data from application and stores these data into storage.   
  3. **HTTP Server:** It accepts queries for the stored data, for querying it uses PromQL.
* Prometheus pulls metrics from application, but it pushes alert. Because pushing into Prometheus may cause high traffic load. But for rarely used jobs it uses pull operation.
* Prometheus' configuration file is used to tell the server which target and at which interval it should scrape data.
* You can port-forward or can use ingress for Grafana

### Service Monitor
Service Monitor is a custom kubernetes component. It is used to define an application you wish to scrape metrics from within Kubernetes, the controller will action the ServiceMonitors we define and automatically build the required Prometheus configuration.
In Service Monitor there is a field `label.release = prometheus` which basically tells that for which application it is discoverable by prometheus.

### Exporter
Exporter is basically a translator between apps data to Prometheus understandable metrics. It fetches data from application, converts and exposes /metric path to prometheus. 
