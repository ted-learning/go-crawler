### start elasticsearch

```shell
docker network create crawler
docker run -d --name elasticsearch --net crawler -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.17.2
```

### data saver
```shell
docker build . -t hataketed/go-crawler-data-saver:0.1

docker run -d --name go-crawler-data-saver --net=host -p 1234:1234 hataketed/go-crawler-data-saver:0.1
```
