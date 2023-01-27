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

## Apache APISIX

<details>
<summary>Click me</summary>
<br/>
  
</details>
