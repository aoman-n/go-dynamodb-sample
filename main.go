package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// MyFirstTable
// key: MyHashKey - S(文字列型)
// key: MyRangeKey - N(数値型)
func setUpDB() *dynamo.DB {
	dynamoDbRegion := os.Getenv("AWS_REGION")
	disableSsl := false

	dynamoDBEndpoint := os.Getenv("DYNAMO_ENDPOINT")
	if len(dynamoDBEndpoint) != 0 {
		disableSsl = true
	}

	if len(dynamoDbRegion) == 0 {
		dynamoDbRegion = "ap-northeast-1"
	}

	return dynamo.New(session.New(), &aws.Config{
		Region:     aws.String(dynamoDbRegion),
		Endpoint:   aws.String(dynamoDBEndpoint),
		DisableSSL: aws.Bool(disableSsl),
	})
}

type DynamoOperator struct {
	Db    *dynamo.DB
	Table *dynamo.Table
}

func NewDynamoOperator() *DynamoOperator {
	db := setUpDB()
	table := db.Table("MyFirstTable")

	return &DynamoOperator{
		Db:    db,
		Table: &table,
	}
}

var (
	PutErrMsg = "failed to put item %v"
	GetErrMsg = "failed to get item %v"
)

// Create 項目を作成
func (d *DynamoOperator) Create(item *Item) error {
	if err := d.Table.Put(item).Run(); err != nil {
		return fmt.Errorf(PutErrMsg, err)
	}
	return nil
}

// Read 項目を取得
func (d *DynamoOperator) GetByHashKey(key string) (*Item, error) {
	var result Item
	if err := d.Table.Get("MyHashKey", key).One(&result); err != nil {
		return nil, fmt.Errorf(GetErrMsg, err)
	}

	return &result, nil
}

// Item DynamoDB格納用構造体
type Item struct {
	MyHashKey  string
	MyRangeKey int
	MyText     string
}

func main() {
	dynamoOperator := NewDynamoOperator()

	item := Item{
		MyHashKey:  "MyHash2",
		MyRangeKey: 2,
		MyText:     "Hello, My Text2",
	}
	err := dynamoOperator.Create(&item)
	if err != nil {
		log.Println("dynamoOperator Create error: ", err)
	}

	resultItem, err := dynamoOperator.GetByHashKey(item.MyHashKey)
	if err != nil {
		log.Println("dynamoOperator GetByHashKey error: ", err)
	}
	log.Println("resultItem: ", resultItem)
}
