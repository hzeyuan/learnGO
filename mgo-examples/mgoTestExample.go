package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Posts struct {
	Title string
	Content string
	Date time.Time
}

var (
	Session,_ = mgo.Dial("localhost")
	Database = "mgo"
	Collection = "posts"
	Coll = Session.DB(Database).C(Collection)
)

// Drop Database
func DropDatebase(){
	fmt.Println("drop Database -->")
	if err :=Session.DB(Database).DropDatabase();err!=nil{
		fmt.Println("drop Datebase fail!!")
	}
}

// 添加一个文档
func TestInsert(){
	fmt.Println("insert document to mongo DB -->")
	post1 := &Posts{
		Title:   "post1",
		Content: "post1-content",
		Date:    time.Now(),
	}
	Coll.Insert(post1)

}
// 添加多个文档
func TestMultipleInsert(){
	t:=time.Now()
	fmt.Println("insert Multi document -->")
	var multiPosts []interface{}
	for i:=1;i<5001;i++{
		multiPosts = append(multiPosts,&Posts{
			Title:   fmt.Sprintf("post-%d",i),
			Content: fmt.Sprintf("post-%d-content",i),
			Date:    time.Now(),
		})
	}
	Coll.Insert(multiPosts...)
	fmt.Println(time.Since(t))
}

//批量插入
func TestBulkInsert(){
	t:=time.Now()
	fmt.Println("Bulk Insert -->")
	b :=Coll.Bulk()
	var bulkPosts []interface{}
	for i:=10;i<5010;i++{
		bulkPosts = append(bulkPosts,&Posts{
			Title:   fmt.Sprintf("post-%d",i),
			Content: fmt.Sprintf("post-%d-content",i),
			Date:    time.Now(),
		})
	}
	b.Insert(bulkPosts...)
	if _,err:=b.Run();err!=nil{
		fmt.Println(err)
	}
	fmt.Println(time.Since(t))
}

//更新文档操作
func TestUpdate(){
	fmt.Println("Test Update in mongo DB -->")
	selector := bson.M{"title":"post1"}
	update :=bson.M{"$set":bson.M{"title":"post1-update"}}
	if err := Coll.Update(selector,update);err!=nil{
		fmt.Println(err)
	}
}

//添加或更新文档
func TestUpsert() {
	fmt.Println("Test Upsert in Mongo DB -->")

	//添加或者更新文档
	update := bson.M{"$set": bson.M{"content": "post-upsert-content"}}
	selector := bson.M{"title": "post-upsert-title"}

	_, err := Coll.Upsert(selector, update)
	if err != nil {
		panic(err)
	}
}

//查询文档
func TestSelect(){
	fmt.Println("Test Select in Mongo DB -->")
	var result Posts
	var results []Posts
	if err:=Coll.Find(bson.M{"title":"post1-update"}).One(&result);err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("find one:%v\n",result)
	if err:=Coll.Find(bson.M{"title":"post1-update"}).All(&results);err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("find all:%v\n",results)
	if err:=Coll.Find(bson.M{"title":"post1-update"}).All(&results);err!=nil{
		fmt.Println(err)
	}
	if err:=Coll.Find(bson.M{"title":"post1-update"}).Limit(1).All(&results);err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("find limit:%v\n",results)
}

//聚合操作
func TestAggregate(){
	pipeline := []bson.M{
		{"$match": bson.M{"title": "post1-update" }},
	}
	pipe := Coll.Pipe(pipeline)

	result := []bson.M{}
	//err := pipe.AllowDiskUse().All(&result) //allow disk use
	err := pipe.All(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("find TestAggregate result:", result)
}

//保存json到mongo
func saveJsonToDB(){
	var f interface{}
	j:=[]byte(`{"posts": {
		"title": "post1",
		"content": "post1-content"
	}
}`)
	if err:=json.Unmarshal(j,&f);err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("%v",&f)
	if err:=Coll.Insert(&f);err!=nil{
		fmt.Println(err)
	}

}
func main(){
	TestInsert()
	TestUpdate()
	TestUpsert()
	TestSelect()
	TestAggregate()
	TestMultipleInsert()
	TestBulkInsert()
	saveJsonToDB()
	defer Session.Close()
}


