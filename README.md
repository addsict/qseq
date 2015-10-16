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

This sequence server achieved over 60K requests per second on my MacBook Pro (2.3 GHz Core i7).

```
$ wrk -c 100 -t 10 -d 10 http://127.0.0.1:9000/sequences/foo
Running 10s test @ http://127.0.0.1:9000/sequences/foo
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.64ms  602.02us  31.82ms   95.61%
    Req/Sec     6.17k     1.08k   17.42k    77.87%
  616108 requests in 10.10s, 71.58MB read
Requests/sec:  61004.27
Transfer/sec:      7.09MB
```
