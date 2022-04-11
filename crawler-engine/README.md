### run data-saver first

```shell
docker run -d --name go-crawler-data-saver --net=host -p 1234:1234 hataketed/go-crawler-data-saver:0.1
```

### crawler-engine
```shell
docker build . -t hataketed/go-crawler-engine:0.1

docker run -d --name go-crawler-engine --net=host -p 7500:7500 hataketed/go-crawler-engine:0.1
```
