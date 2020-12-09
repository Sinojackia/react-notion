package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"log"
	"math/rand"
	"src/model"
	"time"
)

type queryHook struct {
}

func (q queryHook) BeforeQuery(ctx context.Context, event *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (q queryHook) AfterQuery(ctx context.Context, event *pg.QueryEvent) error {
	// 记录SQL语句
	stmt, err := event.FormattedQuery()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(string(stmt))
	return nil
}

func NewNode() (*snowflake.Node, error) {
	snowflake.Epoch = time.Now().Unix()
	rand.Seed(rand.Int63n(time.Now().UnixNano())) // nolint
	node := 0 + rand.Int63n(1023-0)               // nolint
	return snowflake.NewNode(node)
}

func main() {
	var (
		err  error
		node *snowflake.Node
	)

	node, err = NewNode()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(node.Generate().Base36())
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "123456",
		Database: "alayou",
	})
	db.AddQueryHook(queryHook{})

	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err.Error())
		}
	}()
	err = db.Model(&model.BlockValue{}).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Model(&model.PageValue{}).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	//bytes, err := ioutil.ReadFile("./data/2e22de6b770e4166be301490f6ffd420.json")
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//v := new(model.BlockMap)
	//err = json.Unmarshal(bytes, v)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//for _, block := range *v {
	//	_, err = db.Model(&block.Value).Insert()
	//	if err != nil {
	//		log.Fatal(err)
	//		return
	//	}
	//}
	pageSize := 20
	page := 1

	m := make([]model.BlockValue, pageSize)
	query := db.Model(&m)
	query.Offset((page - 1) * pageSize).
		Where("type='toggle'").
		Limit(pageSize).Order("created_time")
	count, err := query.SelectAndCount()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
	resBytes, err := json.Marshal(m[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(resBytes))
	m2 := new(model.ToggleValue)
	err = json.Unmarshal(resBytes, m2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m2.CreatedById)
	fmt.Println(m2.Format)
}
