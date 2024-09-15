# nodelay
understanding Nagle's algo

## How do you test this?

1. Run server like so:

```bash
NO_DELAY=false go run server/main.go
```

`NO_DELAY` can be set to true or false

2. Run the client like this:

```bash
go run client/main.go
```

3. Plot the graphs if you hate looking at logs like me:

```bash
python plot_latency.py 
```