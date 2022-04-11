### start elasticsearch & crawler-engine

```shell
docker network create crawler
docker run -d --name elasticsearch --net crawler -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.17.2

docker run -d --name go-crawler-engine --net=host -p 7500:7500 hataketed/go-crawler-engine:0.1
```

### start portal
```shell
docker build . -t hataketed/go-crawler-portal:0.2

docker run -d --name go-crawler-portal --net=host -p 8888:8888 hataketed/go-crawler-portal:0.2
```
