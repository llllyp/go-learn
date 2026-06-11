package main

import "fmt"

// 定义一个数据库接口
type Database interface {
    Connect() error
    Query(sql string) (interface{}, error)
    Close() error
}

// 实现MySQL数据库
type MySQL struct {
    connection string
}

func (m *MySQL) Connect() error {
    // 连接MySQL
    return nil
}

func (m *MySQL) Query(sql string) (interface{}, error) {
    // 执行查询
    return []string{"result1", "result2"}, nil
}

func (m *MySQL) Close() error {
    // 关闭连接
    return nil
}

// 实现PostgreSQL数据库
type PostgreSQL struct {
    connection string
}

func (p *PostgreSQL) Connect() error {
    // 连接PostgreSQL
    return nil
}

func (p *PostgreSQL) Query(sql string) (interface{}, error) {
    // 执行查询
    return []string{"result1", "result2"}, nil
}

func (p *PostgreSQL) Close() error {
    // 关闭连接
    return nil
}

// 多态使用
func ExecuteQuery(db Database, sql string) {
    db.Connect()
    defer db.Close()
    
    result, err := db.Query(sql)
    if err != nil {
        panic(err)
    }
    fmt.Println("Result:", result)
}