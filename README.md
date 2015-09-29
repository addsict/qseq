qseq
===
(under development)
qseq is a simple and fast sequence number generator with RESTful HTTP API


# RESTful API

## Create new sequence

```
PUT /sequences/foo
```

## Get next sequence value

```
GET /sequences/foo
```

## Update sequence value

```
PUT /sequences/foo
body: 1000
```

## Delete sequence

```
DELETE /sequences/foo
```

## List sequences

```
GET /sequences
```
