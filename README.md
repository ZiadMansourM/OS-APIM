# 🧐 API Management Solutions / Platforms
Introduction to the `process of API Management` a crucial topic for every organization. Plus, analysis to `Open-Source APIM Solutions and Platforms`.

🦦 Table of contents
--------------------
| Number | Topic  |
| :-----: | :-: |
| 1 | [What is API Managment](https://github.com/ZiadMansourM/OS-APIM#what-is-api-managment) |
| 2 | [Popular APIM Solutions features](https://github.com/ZiadMansourM/OS-APIM#popular-apim-solutions-features) |
| 3 | [How to Pick the Right APIM Solution](https://github.com/ZiadMansourM/OS-APIM#how-to-pick-the-right-open-source-api-management-solution) |
| 4 | [Popular OS-APIM Soltions](https://github.com/ZiadMansourM/OS-APIM#popular-os-apim-soltions) |
| 5 | [Analyzing Current APIM Solutions](https://github.com/ZiadMansourM/OS-APIM#analyzing-current-apim-solutions) |

## What is API Managment:
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

## Popular APIM Solutions features:
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


## How to Pick the Right Open-Source API Management Solution
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

## Popular OS APIM Soltions
```Console
$ Apache APISIX
$ Kong Enterprise
$ API Umbrella
$ Gravitee.io
$ APIman.io
$ WSO2 API Manager
$ Tyk.io
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
$ KrakenD
```

## Analyzing Current APIM Solutions
| Number | Solution-Name | Docker Iamge | Summary |
| :-----: | :-: | :-: | :-: |
| 1 | [Kong Enterprise](https://konghq.com/) | [Image](https://hub.docker.com/_/kong) | [Kong Enterprise](https://github.com/ZiadMansourM/OS-APIM#kong-enterprise) |
| 2 | [Tyk](https://tyk.io/) | [Image](https://hub.docker.com/r/tykio/tyk-gateway) | [Tyk](https://github.com/ZiadMansourM/OS-APIM#tyk) |
| 3 | [KrakenD](https://www.krakend.io/) | [Image](https://hub.docker.com/r/devopsfaith/krakend) | [KrakenD](https://github.com/ZiadMansourM/OS-APIM#krakend) |
| 4 | [Apache APISIX](https://apisix.apache.org/) | [Image](https://hub.docker.com/r/apache/apisix) | [Apache APISIX](https://github.com/ZiadMansourM/OS-APIM#apache-apisix) |


## Kong Enterprise

| Language | Docker Pull |
| :-----: | :-: |
| Lua | 100M+ |

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

## Tyk

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

## KrakenD

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

## Apache APISIX


