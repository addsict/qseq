qseq
===
(under development)
qseq is a simple and fast sequence number generator with RESTful HTTP API

何を解決しようとしたものなのか
* sequenceテーブルだと、一つのテーブル毎に毎回定義するのが開発時には特に煩雑だし、シーケンス値取得のためだけにMySQL使うのはどうなのかと(実際は他のDBと共存する運用だが)
* 理想は何も準備せずに、このシーケンスをくださいって言えば勝手にシーケンスが作られて、シーケンス値が取得できるといい。


# RESTful API

## Create new sequence

```
PUT /sequences/foo
```

## Get next sequence value

```
GET /sequences/foo
GET /sequences/foo?increment=10
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

## List all keys of sequence

```
GET /sequences
```
