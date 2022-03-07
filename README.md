##Imports

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

engine -> scheduler -> workers -> engine