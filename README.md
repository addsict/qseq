qseq
===
qseq is a simple and fast sequence number generator with RESTful HTTP API

## Run the server

```bash
$ qseq --datadir=. --port=9000
```

## RESTful API

### Create a new sequence

```bash
$ curl -X PUT http://127.0.0.1:8080/sequences/foo
```

### Get the sequence value

```bash
# next value
$ curl http://127.0.0.1:8080/sequences/foo

# next 10 value
$ curl http://127.0.0.1:8080/sequences/foo?increment=10

# current value
$ curl http://127.0.0.1:8080/sequences/foo?increment=0
```

### Update the sequence value

```bash
$ curl -X PUT -d 100 http://127.0.0.1:8080/sequences/foo
```

### Delete the sequence

```bash
$ curl -X DELETE http://127.0.0.1:8080/sequences/foo
```

## Performance

This sequence server achieved 27K requests per second on my MacBook Air 11 inch machine (1.3 GHz Core i5).

```
$ wrk -c 100 -t 10 -d 10 http://127.0.0.1:8080/sequences/foo
Running 10s test @ http://127.0.0.1:8080/sequences/foo
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.71ms    1.14ms  25.48ms   82.34%
    Req/Sec     2.90k   475.10     4.55k    73.70%
  274189 requests in 10.00s, 31.80MB read
Requests/sec:  27417.81
Transfer/sec:      3.18MB
```
