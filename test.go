package gossdb

import (
	"fmt"
)

func main() {
	db, err := Conn("127.0.0.1", 8888)
	if err != nil {
		panic(err)
	}

	db.Set("test", "456")

	ret, err := db.Get("test")
	fmt.Println(ret, err)

	db.Incr("test", 100)

	ret, err = db.Get("test")
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)

	//batch mode
	db.Batch()
	db.Set("test", "123")
	db.Incr("test", 123)
	db.Get("test")
	db.Zrange("test", 0, 10)
	res, err := db.Exec()

	if err != nil {
		panic(err)
	}
	for _, item := range res {
		fmt.Println(item)
	}
}
