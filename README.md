# 🧐 API Management Solutions / Platforms
Introduction to the `process of API Management` a crucial topic for every organization. Plus, analysis of existing `Open-Source APIM Solutions and Platforms`.

🦦 Table of contents
--------------------
| Number | Topic  |
| :-----: | :-: |
| 1 | What is API Management |
| 2 | What features to expect from an APIM solution |
| 3 | How to Pick the Right APIM Solution |
| 4 | Popular OS-APIM Soltions |
| 5 | Analyzing Current APIM Solutions |

## What is API Management?

<details>
<summary>Click me</summary>
<br/>

```Console
API management is the process of:
$ Building
$ Publishing
$ Securing
$ Versioning
$ Observing
$ Scaling
$ Cataloging
$ Retiring APIs
```

</details>

## What features to expect from an APIM solution?

<details>
<summary>Click me</summary>
<br/>

An API management solution is a SaaS or PaaS product at the end of the day, aiming at making APIM process easier and more efficient. But, every Open-Source APIM Solution or Platform tool out there has specific requirements and implementation limitations. Here we will list some of these popular features.

```Console
$ API Gateway
- [1] Handles the routing and management for all your APIs through policies that can be applied to individual APIs 
or a set of APIs.
- [2] rate-limiting: general-purpose capability to protect one’s information system from DDoS attacks.
- [3] Complex rate-limiting: e.g. for Billing
When a business is to sell data, they might sell them based on volume consumption. A central-access-point 
aka API Gateway suits complex distributed architectures that rely on several services to serve the required data 
and can reliably measure and charge usage.
```

```Console
$ API Portal
- User interfaces models for API "Management / Publishing / Documenting" which serves as a self-service hub for 
developers to gain access / introduced to avaliable APIs and other resources needed for consuming them.
```

```Console
$ API Dashboard
- Where you can observe API usage, view key API metrics, and manage a variety of API analytics.
```

```Console
$ API Catalog
- Organizing, cataloging, indexing, and presenting all the public and private APIs that are used in an 
organization.
```

```Console
$ API Builder
- Enable the use of low-code / no-code tools to design, build, integrate, and deploy APIs.
```

</details>

## How to Pick the Right Open-Source API Management Solution?

<details>
<summary>Click me</summary>
<br/>

```Console
$ [1] Features: 
The foremost factor to consider is the features provided from the APIMP.
Every OS-APIMP tool has specific requirements and implementation limitations.
So, performing a comparison in terms of your technical and non-technical requirements can be really helpful.
```

```Console
$ [2] Deployment Complexity:
Some gateways require:
- single-node while others run on multiple-nodes.
- many databases to operate, which can be a complicated thing to manage.
So, deployment-complexity is our second factor for finding the right OS-APIM.
```

```Console
$ [3] On-Premise vs Cloud hosted:
- On-Premise: can add more time to the deployment and maintenance process.
- Cloud-hosted: can introduce more latency due to the extra-hop and even lower down the availability
of your service when the vendor goes down.
Analyzing the pros and cons of both the options to determine which open-source API management tool to embrace
can also be a profitable idea.
```

```Console
$ [4] Community Support:
Having an idea of the type of community support each API management open source tool needs also helps in 
choosing the right one. The larger and active a community is, the higher are the chances of receiving 
updates and solutions to your queries; ultimately resulting in a better experience.
```

</details>

## Popular OS APIM Soltions

<details>
<summary>Click me</summary>
<br/>

```Console
$ Apache APISIX
$ Kong Enterprise
$ Gravitee.io
$ Tyk.io
$ KrakenD
$ Ambassador
$ Ocelet
```

```Console
I checked the following and they don't seem good candidates:
According to their Docker pull numbers and stars
$ API Umbrella
$ APIman.io
$ WSO2 API Manager
$ Fusio
$ Apigility
$ SwaggerHub
$ API Axle
$ IBM Bluemix API
$ Repose
$ SnapLogic Enterprise Integration Cloud
$ DreamFactory
$ 3Scale
$ Gloo Edge
$ Akana
$ Mashery
$ Azure
```

</details>

## Analyzing Current APIM Solutions
| Number | Solution-Name | Docker Image | Language | Docker Pull | GitHub Stars |
| :-----: | :-: | :-: | :-: | :-: | :-: |
| 1 | [Kong Enterprise](https://konghq.com/) | [Image](https://hub.docker.com/_/kong) | Lua | 100M+ | [33.9k](https://github.com/Kong/kong) |
| 2 | [Tyk](https://tyk.io/) | [Image](https://hub.docker.com/r/tykio/tyk-gateway) | Go | 5M+ | [8.1k](https://github.com/TykTechnologies/tyk) |
| 3 | [KrakenD](https://www.krakend.io/) | [Image](https://hub.docker.com/r/devopsfaith/krakend) | Go | 10M+ | [1.1k](https://github.com/krakendio/krakend-ce) |
| 4 | [Gravitee](https://www.gravitee.io/) | [Image](https://hub.docker.com/r/graviteeio/management-api) | Java | 5M+ | [1.6k](https://github.com/gravitee-io/gravitee-api-management) |
| 5 | [Apache APISIX](https://apisix.apache.org/) | [Image](https://hub.docker.com/r/apache/apisix) | Lua, Go, Python, Perl | 5M+ | [11.1k](https://github.com/apache/apisix) |
| 6 | [Ambassador](https://www.getambassador.io/) | [Image](https://hub.docker.com/r/datawire/ambassador) | Kubernetes-native built on Envoy  | 10M+ | - |
| 7 | [Ocelot](https://learn.microsoft.com/en-us/dotnet/architecture/microservices/multi-container-microservice-net-applications/implement-api-gateways-with-ocelot) | [Image](https://hub.docker.com/r/datawire/ambassador) | .NET Core  | 100K+ | [7.5k](https://github.com/ThreeMammals/Ocelot) |


## Kong Enterprise

<details>
<summary>Click me</summary>
<br/>

| Solution-Name | Docker Image | Language | Docker Pull | GitHub Stars |
| :-: | :-: | :-: | :-: | :-: |
| [Kong Enterprise](https://konghq.com/) | [Image](https://hub.docker.com/_/kong) | Lua | 100M+ | [33.9k](https://github.com/Kong/kong) |
  
# Kong quick summary

```Console
Why Kong Enterprise ?
$ Popular 100M+ Docker Pull, open-source, and advanced cloud-native API gateway built for universal deployment. 
$ Run on any platform. 
$ Written in Lua programming language and supports hybrid and multi-cloud infrastructure.
$ Optimized for microservices and distributed architectures.
$ At its core, Kong is built for high performance, extensibility, and portability. 
$ Kong is also lightweight, fast, and scalable. 
$ Supports declarative configuration without a database, using in-memory storage only, and native Kubernative CRDs.
$ Kong features: 
- Load balancing (with different algorithms). 
- Logging.
- Authentication (support for OAuth2.0)
- Rate-limiting
- Transformations
- Live monitoring
- Service discovery
- Caching
- Failure detection and recovery
- Clustering. Most importantly, Kong supports the clustering of nodes and serverless functions.
$ Supports the configuration of proxies for your services, and serve them over SSL, or use WebSockets. 
$ Can load balance traffic through replicas of your upstream services. 
$ Can monitor the availability of your services, and adjust its load balancing accordingly.
$ Additionally, Kong ships with a command-line interface that allows you to manage a Kong cluster from the command line. 
$ Kong is highly extensible using plugins and different kinds of integrations. It can be managed with its RESTful API for maximum flexibility.
```

# Kong detailed analysis "DB-less mode" - [deployed](http://registry.sreboy.com/)

```Console
*** We have three routes for our gateway
$ /quotes >>> quotes-service
$ /jokes >>> jokes-service
$ /random >>> loadbalancer "One of the randomizer servers"
*** You can not make more than 10 requests/minute to any endpoint
```

### Why Dbless mode
```Console
*** Kong has DB mode "Cassandra, Postgresql" and Dbless mode "kong-config/kong.yaml"
$ In Dbless mode you can version control your configuration files "I saw it as a plus"
$ Note: Some plugins still doesn't support dbless mode and require DB mode tho
```

![kong_detailed_analysis](https://user-images.githubusercontent.com/116031573/214587039-403ec2ab-98cb-4ed1-8d55-1701d60cfb23.png)


### In this [example](https://github.com/ZiadMansourM/OS-APIM/tree/main/kong), we have three services:
  - Quotes Service
  - Jokes Service
  - Randomizer Service "Load balance between two containers"
 
```Console
(venv) ziadh@Ziads-MacBook-Air kong % tree -I venv -I pics -I diagrams -I go.sum   
.
├── README.md
├── docker-compose.yml
├── jokes-service
│   ├── Dockerfile
│   ├── go.mod
│   └── main.go
├── kong-config
│   └── kong.yaml
├── quotes-service
│   ├── Dockerfile
│   ├── go.mod
│   └── main.go
└── randomizer-service
    ├── server-one
    │   ├── Dockerfile
    │   ├── go.mod
    │   └── main.go
    └── server-two
        ├── Dockerfile
        ├── go.mod
        └── main.go

6 directories, 15 files
```


## 🐳 docker-compose

```yaml
version: '3.7'

services:
  kong:
    image: kong:3.1.1-alpine
    hostname: kong
    container_name: kong
    environment:
      KONG_DATABASE: "off"
      KONG_PROXY_ACCESS_LOG: "/dev/stdout"
      KONG_ADMIN_ACCESS_LOG: "/dev/stdout"
      KONG_PROXY_ERROR_LOG: "/dev/stderr"
      KONG_ADMIN_ERROR_LOG: "/dev/stderr"
      KONG_ADMIN_LISTEN: "0.0.0.0:8001, 0.0.0.0:8444 ssl"
      KONG_DECLARATIVE_CONFIG: "/opt/kong/kong.yaml"
    command: "kong start"
    ports:
      - "8000:8000" # incoming HTTP traffic
      - "8443:8443" # incoming HTTPS traffic
      - "8001:8001" # Admin API HTTP
      - "8444:8444" # Admin API HTTPS
    volumes:
      - ./kong-config:/opt/kong
    depends_on:
      - quotes
      - jokes 
      - random-one
      - random-two
  quotes:
    build:
      context: quotes-service/
      dockerfile: Dockerfile
  jokes:
    build:
      context: jokes-service/
      dockerfile: Dockerfile
  random-one:
    build:
      context: randomizer-service/server-one
      dockerfile: Dockerfile
  random-two:
    build:
      context: randomizer-service/server-two
      dockerfile: Dockerfile
```

## 🦍 Kong configuration file "/kong-config/kong.yaml"

```yaml
_format_version: "2.1"
_transform: true

services:
  - name: quotes-service
    url: http://quotes:3000
    routes:
      - name: quotes-service-routes
        paths:
          - /quotes
        strip_path: true
  - name: jokes-service
    url: http://jokes:3000
    routes:
      - name: jokes-service-routes
        paths:
          - /jokes
        strip_path: true
  - name: randomizer-service
    host: randomizer-upstream
    protocol: http
    routes:
      - name: randomizer-service-routes
        paths:
          - /random
        strip_path: true

upstreams:
  - name: randomizer-upstream
    targets:
      - target: random-one:3000
        weight: 100
      - target: random-two:3000
        weight: 100

plugins:
  - name: rate-limiting
    config:
      minute: 10
      policy: local
```

## Kong Gateway Router in action - [try-quotes](http://registry.sreboy.com/quotes) - [try-jokes](http://registry.sreboy.com/jokes)

Quotes Service |  Jokes Service
:--:|:--:
![quotes](https://user-images.githubusercontent.com/116031573/214593294-401b6b5d-ac03-49d9-9d1b-daef4477cb37.png)  |  ![jokes](https://user-images.githubusercontent.com/116031573/214593334-7fa42c94-0b1e-4d31-bba7-c22fced5253c.png)


## Kong load-balancer in action - [try-it](http://registry.sreboy.com/random)

Server One |  Server Two
:--:|:--:
![server-one](https://user-images.githubusercontent.com/116031573/214591408-f125b466-38bf-420c-8e3e-5cb93afa4e82.png)  |  ![server-two](https://user-images.githubusercontent.com/116031573/214591454-7b1f543a-ee5d-4efa-ae18-b03690b9a5da.png)

## Kong Rate Limiter in action - [try-it](http://registry.sreboy.com/random)


<img width="1440" alt="kong-rate-limit" src="https://user-images.githubusercontent.com/116031573/214588034-19f9de77-a758-43dc-a0f5-3b19a90caf75.png">


## Kong Plugin Hub - [link](https://docs.konghq.com/hub/) - [custom-plugins](https://docs.konghq.com/gateway/latest/plugin-development/)
```Console
*** Kong has these functionality-categories of Plugins:
$ Authentication - "Basic Authentication, JWT, OAuth 2.0, ...etc".
$ Security - "CORS, Bot detection, IP Restriction, ...etc".
$ Traffic Control - "Proxy Cache, Rate Limiting, Request Size Limiting, ...etc".
$ Serverless - "AWS Lambda, Azure Functions, ...etc".
$ Analytics & Monitoring - "Datadog, Prometheus, OpenTelemetry".
$ Transformations - "gRPC-gateway, gRPC-Web, ...etc".
$ Logging - "File Log, HTTP Log, Loggy, ...etc".
$ Deployment - "AWS, Azure, KongMap".
*** Also Plugins are categorized on:
$ Free.
$ Plus.
$ Enterprise.
```

## How to include a plugin e.g. [Request Size Limiting](https://docs.konghq.com/hub/kong-inc/request-size-limiting/) plugin
Such a super important plugin is strongly encouraged to be enabled for any Service added to Kong Gateway to prevent a DOS (Denial of Service) attack. You can add it by only appending those lines to Kong.yaml file in case you are using Dbless mode or view [other-example-confg](https://docs.konghq.com/hub/kong-inc/request-size-limiting/#example-config):
```yaml
plugins:
- name: request-size-limiting
  service: SERVICE_NAME|SERVICE_ID
  config: 
    allowed_payload_size: 128
    require_content_length: false
```

### Sadly the Kong [DevPortal](https://docs.konghq.com/gateway/latest/kong-enterprise/dev-portal/) is only included in [Enterprise edition](https://konghq.com/pricing).

## ToDo
- [ ] Enable [Kong Manger](https://docs.konghq.com/gateway/latest/kong-manager/) GUI.

</details>

## Apache APISIX

<details>
<summary>Click me</summary>
<br/>
  
  - [X] Dashboad [deployed](http://registry.sreboy.com:9000/dashboard)
  - [X] Gateway [deployed](http://registry.sreboy.com:9080)
  ```console
  $ /random "loadbalncer"
  $ /nginx "loadbalncer"
  $ /jokes
  $ /quotes
  ```
  - [X] Usefull and important links:
  - Plugins - [link](https://apisix.apache.org/plugins/) e.g. keyclock auth supported.
  - Deployment modes - [link](https://apisix.apache.org/docs/apisix/deployment-modes/).
  - example - [link](https://github.com/ZiadMansourM/OS-APIM).

  | Solution-Name | Docker Image | Language | Docker Pull | GitHub Stars |
  | :-: | :-: | :-: | :-: | :-: |
  | [Apache APISIX](https://apisix.apache.org/) | [Image](https://hub.docker.com/r/apache/apisix) | Lua, Go, Python, Perl | 5M+ | [11.1k](https://github.com/apache/apisix) |
  
  # Apache APISIX Quick Summary
  ```console
  $ Apache APISIX has three deploymet moods
  @https://apisix.apache.org/docs/apisix/deployment-modes/
  1- Traditional 
  2- Decoupled
  3- Standalone
  $ We will be using Traditional mode, for example see:
  @https://github.com/ZiadMansourM/OS-APIM/tree/main/apache-apisix
  ```
  
  # Apache Traditional Mode
  ```Console
  *** We have four routes for our gateway
  $ /quotes >>> quotes-service
  $ /jokes >>> jokes-service
  $ /random >>> loadbalancer "One of the randomizer servers"
  $ /nginx >>> loadbalancer "One of the Nginx servers"
  *** You can not make more than 10 requests/minute to any endpoint
  ```
  
  ![apache_apisix_detailed_analysis](https://user-images.githubusercontent.com/64917739/217309653-a2a4686a-4aa6-4e98-9080-296dc17a1b15.png)
  
  ### In this [example](https://github.com/ZiadMansourM/OS-APIM/tree/main/apache-apisix/traditional), we have four services:
  - Quotes Service.
  - Jokes Service.
  - Randomizer Service.
  - Web Service.
  
  ```console
  ziadh@Ziads-MacBook-Air traditional % tree -I diagrams -I go.sum -I grafana -I etcd -I prometheus
  .
  ├── README.md
  ├── docker-compose.yml
  └── services
      ├── apisix
      │   ├── config
      │   │   └── config.yaml
      │   └── logs
      │       ├── access.log
      │       ├── error.log
      │       └── nginx.pid
      ├── apisix-dashboard
      │   ├── config
      │   │   └── conf.yaml
      │   └── logs
      ├── jokes-service
      │   ├── Dockerfile
      │   ├── go.mod
      │   └── main.go
      ├── quotes-service
      │   ├── Dockerfile
      │   ├── go.mod
      │   └── main.go
      ├── randomizer-service
      │   ├── server-one
      │   │   ├── Dockerfile
      │   │   ├── go.mod
      │   │   └── main.go
      │   └── server-two
      │       ├── Dockerfile
      │       ├── go.mod
      │       └── main.go
      └── upstream
          ├── web-one.conf
          └── web-two.conf

  13 directories, 21 files
  ```
  
  ## 🐳 docker-compose
  ```yaml
  version: "3"

  services:
    apisix-dashboard:
      image: apache/apisix-dashboard:3.0.0-alpine
      restart: always
      volumes:
      -  ./services/apisix-dashboard/config/conf.yaml:/usr/local/apisix-dashboard/conf/conf.yaml
      ports:
      - "9000:9000"
      networks:
        apisix:

    apisix:
      image: apache/apisix:latest
      restart: always
      volumes:
        - ./services/apisix/logs:/usr/local/apisix/logs
        # ro: read only
        - ./services/apisix/config/config.yaml:/usr/local/apisix/conf/config.yaml:ro
      depends_on:
        - etcd
      # network_mode: host
      ports:
        - "9180:9180/tcp" # Access Admin API
        - "9080:9080/tcp" # HTTP Traffic
        - "9443:9443/tcp" # HTTPs Traffic
        - "9091:9091/tcp" # ~Not~Sure~ prometheus port
        - "9092:9092/tcp" # ~Not~Sure~ control port see: 
      networks:
        apisix:

    etcd:
      image: bitnami/etcd:3.4.15
      restart: always
      volumes:
        - ./services/etcd/data:/bitnami/etcd
      environment:
        ETCD_ENABLE_V2: "true"
        ALLOW_NONE_AUTHENTICATION: "yes"
        ETCD_ADVERTISE_CLIENT_URLS: "http://etcd:2379"
        ETCD_LISTEN_CLIENT_URLS: "http://0.0.0.0:2379"
      ports:
        - "2379:2379/tcp"
      networks:
        apisix:

    quotes:
      build:
        context: ./services/quotes-service/
        dockerfile: Dockerfile
      networks:
        apisix:

    jokes:
      build:
        context: ./services/jokes-service/
        dockerfile: Dockerfile
      networks:
        apisix:

    random-one:
      build:
        context: ./services/randomizer-service/server-one
        dockerfile: Dockerfile
      networks:
        apisix:

    random-two:
      build:
        context: ./services/randomizer-service/server-two
        dockerfile: Dockerfile
      networks:
        apisix:

    web-one:
      image: nginx:1.19.0-alpine
      restart: always
      volumes:
        - ./services/upstream/web-one.conf:/etc/nginx/nginx.conf
      ports:
        - "9081:80/tcp"
      environment:
        - NGINX_PORT=80
      networks:
        apisix:

    web-two:
      image: nginx:1.19.0-alpine
      restart: always
      volumes:
        - ./services/upstream/web-two.conf:/etc/nginx/nginx.conf
      ports:
        - "9082:80/tcp"
      environment:
        - NGINX_PORT=80
      networks:
        apisix:

    prometheus:
      image: prom/prometheus:v2.25.0
      restart: always
      volumes:
        - ./services/prometheus/config/config.yaml:/etc/prometheus/prometheus.yml
      ports:
        - "9090:9090"
      networks:
        apisix:

    grafana:
      image: grafana/grafana:7.3.7
      restart: always
      ports:
        - "3000:3000"
      volumes:
        - "./services/grafana/provisioning:/etc/grafana/provisioning"
        - "./services/grafana/dashboards:/var/lib/grafana/dashboards"
        - "./services/grafana/config/grafana.ini:/etc/grafana/grafana.ini"
      networks:
        apisix:

  networks:
    apisix:
      driver: bridge

  volumes:
    etcd_data:
      driver: local
  ```
  
  ## APISIX config.yaml
  
  ```yaml
  apisix:
    node_listen: 9080 # APISIX Gateway listening port for HTTP Traffic
    enable_ipv6: false
    enable_control: true
    control:
      ip: "0.0.0.0"
      port: 9092

  deployment:
    admin:
      allow_admin:
        # http://nginx.org/en/docs/http/ngx_http_access_module.html#allow
        # We need to restrict ip access rules for security. 0.0.0.0/0 is for test.
        - 0.0.0.0/0
      admin_key:
        - name: "admin"
          key: edd1c9f034335f136f87ad84b625c8f1
          # admin: manage all configuration data
          role: admin
        - name: "viewer"
          key: 4054f7cf07e344346cd3f287985e76a2
          role: viewer
    etcd:
      host:
        # it's possible to define multiple etcd hosts addresses of the same etcd cluster.
        - "http://etcd:2379"
      prefix: "/apisix" # apisix configurations prefix
      timeout: 30 # 30 seconds

  plugin_attr:
    prometheus:
      export_addr:
        ip: "0.0.0.0"
        port: 9091
  ```
  
  ## APISIX Dashboard conf.yaml
  
  ```yaml
  conf:
    listen:
      host: 0.0.0.0 # `manager api` listening ip or host name
      port: 9000 # `manager api` listening port
    allow_list: # If we don't set any IP list, then any IP access is allowed by default.
      - 0.0.0.0/0
    etcd:
      endpoints: # supports defining multiple etcd host addresses for an etcd cluster
        - "http://etcd:2379"
      # etcd basic auth info
      # username: "root"    # ignore etcd username if not enable etcd auth
      # password: "123456"  # ignore etcd password if not enable etcd auth
      mtls:
        key_file: "" # Path of your self-signed client side key
        cert_file: "" # Path of your self-signed client side cert
        ca_file: "" # Path of your self-signed ca cert, the CA is used to sign callers' certificates
      prefix: /apisix # apisix config's prefix in etcd, /apisix by default
    log:
      error_log:
        level: warn # supports levels, lower to higher: debug, info, warn, error, panic, fatal
        file_path:
          # supports relative path, absolute path, standard output
          # such as: logs/error.log, /tmp/logs/error.log, /dev/stdout, /dev/stderr
          logs/error.log
      access_log:
        file_path:
          # supports relative path, absolute path, standard output
          # such as: logs/access.log, /tmp/logs/access.log, /dev/stdout, /dev/stderr
          # log example: 2020-12-09T16:38:09.039+0800	INFO	filter/logging.go:46	/apisix/admin/routes/r1	{"status": 401, "host": "127.0.0.1:9000", "query": "asdfsafd=adf&a=a", "requestId": "3d50ecb8-758c-46d1-af5b-cd9d1c820156", "latency": 0, "remoteIP": "127.0.0.1", "method": "PUT", "errs": []}
          logs/access.log
    security:
        # access_control_allow_origin: "http://httpbin.org"
        # access_control_allow_credentials: true # support using custom cors configuration
        # access_control_allow_headers: "Authorization"
        # access_control-allow_methods: "*"
        # x_frame_options: "deny"
        content_security_policy: "default-src 'self'; script-src 'self' 'unsafe-eval' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; frame-src *"  # You can set frame-src to provide content for your grafana panel.

  authentication:
    secret:
      # secret for jwt token generation.
      # NOTE: Highly recommended to modify this value to protect `manager api`.
      # if it's default value, when `manager api` start, it will generate a random string to replace it.
      secret
    expire_time: 3600 # jwt token expire time, in second
    users:
      # username and password for login `manager api`
      - username: admin
        password: admin
      - username: user
        password: user

  plugins:
    - api-breaker
    - authz-keycloak
    - basic-auth
    - batch-requests
    - consumer-restriction
    - cors
    # - dubbo-proxy
    - echo
    # - error-log-logger
    # - example-plugin
    - fault-injection
    - grpc-transcode
    - hmac-auth
    - http-logger
    - ip-restriction
    - jwt-auth
    - kafka-logger
    - key-auth
    - limit-conn
    - limit-count
    - limit-req
    # - log-rotate
    # - node-status
    - openid-connect
    - prometheus
    - proxy-cache
    - proxy-mirror
    - proxy-rewrite
    - redirect
    - referer-restriction
    - request-id
    - request-validation
    - response-rewrite
    - serverless-post-function
    - serverless-pre-function
    # - skywalking
    - sls-logger
    - syslog
    - tcp-logger
    - udp-logger
    - uri-blocker
    - wolf-rbac
    - zipkin
    - server-info
    - traffic-split
  ```
  
  ## Create Upstreams and Routes using Admin API
  ```console
  $ Admin API use http communication.
  $ You can store http put requests to version control them, like below.
  $ You can also use Dashboard to create Routes and Upstreams.
  $ In standalone deployment there is no control domain hence no Admin API, see:
  - https://github.com/ZiadMansourM/OS-APIM/tree/main/apache-apisix
  - https://apisix.apache.org/docs/apisix/deployment-modes/
  - https://apisix.apache.org/docs/apisix/stand-alone/
  ```
  
  ```sh
  #!/bin/bash

  HOST=registry.sreboy.com

  # Create Upstreams
  curl "http://$HOST:9180/apisix/admin/upstreams/1" \
  -H "X-API-KEY: edd1c9f034335f136f87ad84b625c8f1" -X PUT -d '
  {
    "name": "randomizer-upstream",
    "desc": "Upstream for randomizer service",
    "type": "roundrobin",
    "scheme": "http",
    "nodes": {
      "random-one:3000": 1,
      "random-two:3000": 1
    }
  }'

  curl "http://$HOST:9180/apisix/admin/upstreams/2" \
  -H "X-API-KEY: edd1c9f034335f136f87ad84b625c8f1" -X PUT -d '
  {
    "name": "jokes-upstream",
    "desc": "Upstream for jokes service",
    "type": "roundrobin",
    "scheme": "http",
    "nodes": [
      {
        "host": "jokes",
        "port": 3000,
        "weight": 1
      }
    ]
  }'

  curl "http://$HOST:9180/apisix/admin/upstreams/3" \
  -H "X-API-KEY: edd1c9f034335f136f87ad84b625c8f1" -X PUT -d '
  {
    "name": "quotes-upstream",
    "desc": "Upstream for quotes service",
    "type": "roundrobin",
    "scheme": "http",
    "nodes": {
      "quotes:3000": 1
    }
  }'

  curl "http://$HOST:9180/apisix/admin/upstreams/4" \
  -H "X-API-KEY: edd1c9f034335f136f87ad84b625c8f1" -X PUT -d '
  {
    "name": "nginx-upstream",
    "desc": "Upstream for multiple Nginx service",
    "type": "roundrobin",
    "scheme": "http",
    "nodes": {
      "web-one:80": 1,
      "web-two:80": 1
    }
  }'

  # Create Routes
  curl "http://$HOST:9180/apisix/admin/routes/1" \
  -H "X-API-KEY: edd1c9f034335f136f87ad84b625c8f1" -X PUT -d '
  {
    "name": "randomizer-service-routes",
    "desc": "Get a random number",
    "uri": "/random",
    "plugins": {
      "proxy-rewrite": {
        "uri": "/"
      }
    },
    "upstream_id": "1",
    "methods": ["GET"],
    "status": 1
  }'

  curl "http://$HOST:9180/apisix/admin/routes/2" \
  -H "X-API-KEY: edd1c9f034335f136f87ad84b625c8f1" -X PUT -d '
  {
    "name": "jokes-service-routes",
    "desc": "List all jokes",
    "uri": "/jokes",
    "plugins": {
      "proxy-rewrite": {
        "uri": "/"
      }
    },
    "upstream_id": "2",
    "methods": ["GET"],
    "status": 1
  }'

  curl "http://$HOST:9180/apisix/admin/routes/3" \
  -H "X-API-KEY: edd1c9f034335f136f87ad84b625c8f1" -X PUT -d '
  {
    "name": "quotes-service-routes",
    "desc": "List all quotes",
    "uri": "/quotes",
    "plugins": {
      "proxy-rewrite": {
        "uri": "/"
      }
    },
    "upstream_id": "3",
    "methods": ["GET"],
    "status": 1
  }'

  curl "http://$HOST:9180/apisix/admin/routes/4" \
  -H "X-API-KEY: edd1c9f034335f136f87ad84b625c8f1" -X PUT -d '
  {
    "name": "nginx-services-routes",
    "desc": "Route to multiple Nginx upstreams",
    "uri": "/nginx",
    "plugins": {
      "proxy-rewrite": {
        "uri": "/"
      }
    },
    "upstream_id": "4",
    "methods": ["GET"],
    "status": 1
  }'
  ```
  
  ## 🧐 APISIX Gateway router in action
  
  [Quotes Service](http://registry.sreboy.com:9080/quotes) |  [Jokes Service](http://registry.sreboy.com:9080/jokes)
:--:|:--:
![quotes](https://user-images.githubusercontent.com/64917739/217319686-b3240822-6a31-488c-ba5c-47aa1009da8e.png) | ![jokes](https://user-images.githubusercontent.com/64917739/217319717-5c4f54da-8bf5-4b7e-8523-00327e7678af.png)

  ## 🧐 APISIX Gateway router + loadbalancer in action - [try-randomizer](http://registry.sreboy.com:9080/random) - [try-nginx](http://registry.sreboy.com:9080/nginx)
  
  Randomizer Server One |  Randomizer Server Two
  :--:|:--:
  ![Server One](https://user-images.githubusercontent.com/64917739/217321131-03417e58-32f8-4ce9-9fee-a93e64909cb4.png) | ![Server Two](https://user-images.githubusercontent.com/64917739/217321360-cea694f4-1204-4e9d-a66d-055ac4e0e071.png)

  
  [Nginx Upstream One](https://github.com/ZiadMansourM/OS-APIM/blob/main/apache-apisix/traditional/services/upstream/web-one.conf) |  [Nginx Upstream Two](https://github.com/ZiadMansourM/OS-APIM/blob/main/apache-apisix/traditional/services/upstream/web-two.conf)
  :--:|:--:
  ![Nginx One](https://user-images.githubusercontent.com/64917739/217322024-3b40927d-8807-4f72-a3f9-8849eecbaf87.png) | ![Nginx Two](https://user-images.githubusercontent.com/64917739/217322064-beeec2d7-079a-407f-9732-80f5b41b7f3f.png)
  
  ## 🧐 APISIX Dashboard - [link-live](http://registry.sreboy.com:9000/)
  
  ```console
  *** credentials
  $ user: admin
  $ password: admin
  ```
  
  <img width="1440" alt="login" src="https://user-images.githubusercontent.com/64917739/217326628-1851fc6d-8543-47dc-8e3a-c1f71954f3b3.png">

  
  [Upstream Page]() |  Route Page
:--:|:--:
![Upstream](https://user-images.githubusercontent.com/64917739/217326759-c7d3acab-f3cd-4015-ad46-a3b8d49da193.png) | ![Route](https://user-images.githubusercontent.com/64917739/217326946-c360ed39-6355-4ef9-a713-0b46381e22c4.png)

  
</details>


## Tyk

<details>
<summary>Click me</summary>
<br/>

| Language | Docker Pull |
| :-----: | :-: |
| Go | 5M+ |

```Console
Why Tyk ?
$ Tyk (pronounced Taik) is an open-source, powerful, lightweight, and fully-featured API gateway.
$ Written from the ground up using the Go programming language. 
$ It is cloud-native, highly performant with an easily extensible and pluggable architecture based on open standards.
$ Tyk can run independently and only requires Redis as a data store. 
$ Allows users to securely publish and manage a variety of services including:
- legacy
- REST
- GraphQL
$ Baked with so many features which include a variety of: 
- Authentication methods
- Quotas
- Rate-limiting
- Version control
- Notifications and events
- Monitoring, and analytics
$ Supports service discovery, on-the-fly transforms, and virtual endpoints.
$ Allows for creating mock out APIs before release.
$ Tyk also supports API documentation and offers an:
- API Developer portal, a CMS-like system where you can publish your managed APIs and third-party developers 
sign up, enroll to your APIs, and can manage their own keys.
$ Importantly, there is only one version of the Tyk API Gateway and it is 100% Open Source. 
- Whether you are a Community Edition user or an enterprise user, you get the same API Gateway. It ships with all possible
parts required for full usability, with no feature lockout and no black box. With Tyk, you get to 
know exactly how your data is being processed.
```

</details>

## KrakenD

<details>
<summary>Click me</summary>
<br/>

| Language | Docker Pull |
| :-----: | :-: |
| Go | 10M+ |

```Console
Why KrakenD ?
$ Written in Go.
$ Built with performance in mind. KrakenD is a high performance open-source, simple, and pluggable API gateway.
$ Designed with stateless architecture. 
$ Can run everywhere and requires no database to run. 
$ Has a simple configuration and supports unlimited endpoints and backends.
$ KrakenD features:
- Monitoring
- Caching
- User quota
- Rate limiting
- Quality of service (concurrent calls, circuit breaker, and grained timeout) 
- Transformation
- Aggregation, merge sources 
- Filtering (whitelisting and blacklisting)
- Decoding. 
$ KrakenD offers proxy features such as: 
- Load balancing
- Protocol translation
- Oauth
- Security features such as SSL and security policies.
$ You can configure the API gateway behavior by hand or using the KrakenDesigner, a GUI that allows you visually 
design your API from scratch or resume an existing one. 
$ Furthermore, KrakenD’s extensible architecture allows for adding additional functionalities, plug-ins, 
embedded scripts, and middlewares without modifying its source code.
```

</details>

## Gravitee

<details>
<summary>Click me</summary>
<br/>
  
| Language | Docker Pull |
| :-----: | :-: |
| Java | 5M+ |

```Console
Why Gravitee ?
$ Java-based, easy-to-use API management platform that helps organizations to secure, publish, analyze, and document their APIs. 
$ Comes with three major modules, which are:
- [1] API Management (APIM): an open-source, simple yet powerful, flexible, lightweight, and blazing-fast API management (APIM) solution 
designed to give your organization full control over who accesses your APIs, when, and how.
- [2] Access Management (AM): a flexible, lightweight, versatile, and easy-to-use Open Source Identity And Access Management solution. 
It is based on OAuth2/OpenID Connect protocols and acts as an identity provider broker. It features a centralized Authentication 
and Authorization Service to secure your applications and your APIs.
- [3] Alert Engine (AE): a module that allows users to configure alerts and receive notifications to easily and efficiently monitor 
their API platform. It supports multi-channel notifications and suspicious behavior detection, and more.
$ Gravitee.io ships with Cockpit: 
Tool that helps you design your APIs and publishes them across all your environments with fully-featured multi-tenancy support. 
It enables you to scale your Gravitee.io deployment from the platform itself. And graviteeio-cli, a simple command-line tool used 
to manage the Gravitee.io eco-system.
```

</details>
