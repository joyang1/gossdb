# gossdb

## Description

go client for [SSDB](https://github.com/ideawu/ssdb/)

## TODO
support all commands

## Usage

    import "gossdb"

    db, err := Conn("127.0.0.1", 8888)
	if err != nil {
		panic(err)
	}

    db.Set("test", "456")

    ret, err := db.Get("test")
    fmt.Println(ret, err)

    db.Incr("test", 100)

    ret, err = db.Get("test")
    fmt.Println(ret, err)

    
    //batch mode
    db.Batch()
    db.Set("test", "123")
    db.Incr("test", 123)
    db.Get("test")
    db.Zrange("test", 0, 10)
    ret, err := db.Exec()
