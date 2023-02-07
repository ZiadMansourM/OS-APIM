# Traditional deployment mode

To access the Apache APISIX dashboard: [localhost:9000](localhost:9000)

## To create an upstream
```console
*** Randomizer Upstream
curl "http://127.0.0.1:9180/apisix/admin/upstreams/1" \
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
*** Jokes Upstream
curl "http://127.0.0.1:9180/apisix/admin/upstreams/2" \
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
*** Quotes Upstream
curl "http://127.0.0.1:9180/apisix/admin/upstreams/3" \
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
*** Nginx Upstream
curl "http://127.0.0.1:9180/apisix/admin/upstreams/4" \
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
```

## To create a route
```console
*** Randomizer Router
curl "http://127.0.0.1:9180/apisix/admin/routes/1" \
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
*** Jokes Router
curl "http://127.0.0.1:9180/apisix/admin/routes/2" \
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
*** Quotes Router
curl "http://127.0.0.1:9180/apisix/admin/routes/3" \
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
*** Nginx Router
curl "http://127.0.0.1:9180/apisix/admin/routes/4" \
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

```console
*** To Test the route
curl -i -X GET "http://127.0.0.1:9080/get?foo1=bar1&foo2=bar2" -H "Host: example.com"
```


```console
*** Access Admin API
curl "http://127.0.0.1:9180/apisix/admin/routes/" -H 'X-API-KEY: edd1c9f034335f136f87ad84b625c8f1'
```