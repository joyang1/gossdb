# gossdb

## Description

go client for [SSDB](https://github.com/ideawu/ssdb/)

## TODO
support all commands

## Usage

    import "gossdb"

    db := ssdb.Conn("127.0.0.1", 8888)
    defer db.Close()
    if db.Err != nil {
        fmt.Println(db.Err.Error())
        os.Exit(1)
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
