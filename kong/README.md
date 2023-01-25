# Kong analysis "DB-less mode" - [live](http://registry.sreboy.com/)

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
$ In Dbless mode you can version control your configuration easier "I saw it as a plus"
$ Note: Some plugins still doesn't support dbless mode and require DB mode tho
```

![architecture](https://github.com/ZiadMansourM/OS-APIM/blob/main/kong/diagrams/final.png)

### In this example we have three services:
  - Quotes Service
  - Jokes Service
  - Randomizer Service "Load balance between two containers"


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


## Kong Rate Limiter in action

![rate-limiter](https://github.com/ZiadMansourM/OS-APIM/blob/main/kong/pics/kong-rate-limit.png)
