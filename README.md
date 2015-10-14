qseq
===
qseq is a simple and fast sequence number generator with RESTful HTTP API

# Run server

```bash
$ qseq --datadir=. --port=9000
```

# RESTful API

## Create a new sequence

```bash
$ curl -X PUT http://127.0.0.1:8080/sequences/foo
```

## Get the next sequence value

```bash
$ curl http://127.0.0.1:8080/sequences/foo
```

## Update the sequence value

```bash
$ curl -X PUT -d 100 http://127.0.0.1:8080/sequences/foo
```

## Delete the sequence

```bash
$ curl -X DELETE http://127.0.0.1:8080/sequences/foo
```
