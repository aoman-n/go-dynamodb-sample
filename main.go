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

func NewDynamoOperator(tableName string) *DynamoOperator {
	db := setUpDB()
	table := db.Table(tableName)

	return &DynamoOperator{
		Db:    db,
		Table: &table,
	}
}

var (
	PutErrMsg    = "failed to put item %s"
	GetErrMsg    = "failed to get item %s"
	UpdateErrMsg = "failed to update item %s"
	DeleteErrMsg = "failed to delete item %s"
)

// Create 項目を作成
func (d *DynamoOperator) Create(item *Item) error {
	if err := d.Table.Put(item).Run(); err != nil {
		return fmt.Errorf(PutErrMsg, err)
	}
	return nil
}

// Read 項目を取得
func (d *DynamoOperator) GetByHashKey(hashKey string, rangeKey int) (*Item, error) {
	var result Item
	err := d.Table.
		Get("MyHashKey", hashKey).
		Range("MyRangeKey", dynamo.Equal, rangeKey).
		One(&result)
	if err != nil {
		return nil, fmt.Errorf(GetErrMsg, err)
	}

	return &result, nil
}

// Update 項目を更新
func (d *DynamoOperator) Update(item *Item) error {
	err := d.Table.
		Update("MyHashKey", item.MyHashKey).
		Range("MyRangeKey", item.MyRangeKey).
		Set("MyText", item.MyText).
		Value(item)
	if err != nil {
		return fmt.Errorf(UpdateErrMsg, err)
	}
	return nil
}

func (d *DynamoOperator) Delete(item *Item) error {
	err := d.Table.
		Delete("MyHashKey", item.MyHashKey).
		Range("MyRangeKey", item.MyRangeKey).
		Run()
	if err != nil {
		return fmt.Errorf(DeleteErrMsg, err)
	}
	return nil
}

// Item DynamoDB格納用構造体
type Item struct {
	MyHashKey  string
	MyRangeKey int
	MyText     string
}

func main() {
	dynamoOperator := NewDynamoOperator("MyFirstTable")

	item := Item{
		MyHashKey:  "my hash key 1",
		MyRangeKey: 3,
		MyText:     "text1",
	}
	err := dynamoOperator.Create(&item)
	if err != nil {
		log.Println("dynamoOperator Create error: ", err)
	}

	resultItem, err := dynamoOperator.GetByHashKey(item.MyHashKey, item.MyRangeKey)
	if err != nil {
		log.Println("dynamoOperator GetByHashKey error: ", err)
	}
	log.Println("resultItem: ", resultItem)

	item.MyText = "Hello, Updated!!!!"
	err = dynamoOperator.Update(&item)
	if err != nil {
		log.Println("dynamoOperator Update error: ", err)
	}

	err = dynamoOperator.Delete(&item)
	if err != nil {
		log.Println("dynamoOperator Delete error: ", err)
	}
}
