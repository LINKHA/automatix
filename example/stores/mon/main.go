package main

import (
	"context"
	"fmt"

	"github.com/linkha/automatix/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// 	1.初始化 Model
	// 在你的应用中，首先需要创建一个 Model 的实例。这可以通过调用 NewModel 或 MustNewModel 函数来完成。
	uri := "mongodb://localhost:27017"
	db := "your_database_name"
	collection := "your_collection_name"
	model, err := mon.NewModel(uri, db, collection)
	if err != nil {
		// 处理错误
	}

	// 	2.使用 Model 进行数据库操作
	// 一旦你有了一个 Model 的实例，就可以使用它来对数据库执行各种操作。

	// a. 插入文档
	document := bson.M{"key": "value"}
	insertResult, err := model.InsertOne(context.Background(), document)
	if err != nil {
		fmt.Println(insertResult)
	}

	// b. 查找文档
	filter := bson.M{"key": "value"}
	var result bson.D
	err = model.FindOne(context.Background(), &result, filter)
	if err != nil {
		fmt.Println(err)
	}

	// c. 更新文档
	filter2 := bson.M{"key": "value"}
	update := bson.M{"$set": bson.M{"another_key": "another_value"}}
	updateResult, err := model.UpdateOne(context.Background(), filter2, update)
	if err != nil {
		fmt.Println(updateResult)
	}

	//d. 删除文档
	filter3 := bson.M{"key": "value"}
	deleteResult, err := model.DeleteOne(context.Background(), filter3)
	if err != nil {
		fmt.Println(deleteResult)
	}

	// // 4. 管理会话和事务
	// // 该代码还提供了会话和事务的管理。例如，开始一个新会话：
	// sess, err := model.StartSession()
	// if err != nil {
	// 	// 处理错误
	// }
	// // 之后可以使用会话来执行事务：
	// result, err = sess.WithTransaction(context.Background(), func(sessCtx mongo.SessionContext) (any, error) {
	// 	// 在此处执行事务内的操作
	// 	return nil, nil
	// })
	// if err != nil {
	// 	fmt.Println(result)
	// }
}
