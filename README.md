## Imports

```
# gbk to utf8
go get -u golang.org/x/text

# check html charset
go get -u golang.org/x/net/html
```

## Simple

main -> engine -> loop(fetcher -> parser)

index -> teams -> rosters -> players


## Concurrence

engine(result chan) -> scheduler(master chan) -> workers(master, result) -> engine


## Helm integrate
TODO: health check for each service