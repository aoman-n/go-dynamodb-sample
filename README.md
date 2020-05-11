## go-dynamodb

DynamoDBローカルを使ってアクセスしてみる

## aws cli command

テーブル内すべて取得

```
aws dynamodb scan --table-name MyFirstTable --endpoint-url http://localhost:8000
```

## ..

### scan

https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/Scan.html

### query

https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/Query.html