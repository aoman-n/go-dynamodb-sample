## go-dynamodb-sample

DynamoDB Localを使ってみる

## AWS CLI

### Create Table Command

```
aws dynamodb create-table \
  --endpoint-url http://localhost:8000 \
  --table-name MyFirstTable \
  --attribute-definitions AttributeName=MyHashKey,AttributeType=S AttributeName=MyRangeKey,AttributeType=N \
  --key-schema AttributeName=MyHashKey,KeyType=HASH AttributeName=MyRangeKey,KeyType=RANGE \
  --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1
```

### Scan Command

```
aws dynamodb scan --table-name MyFirstTable --endpoint-url http://localhost:8000
```

## Reference

### scan

https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/Scan.html

### query

https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/Query.html
