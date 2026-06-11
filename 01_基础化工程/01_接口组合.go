package main

// 定义多个接口
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

// 接口组合
type ReadWriter interface {
    Reader
    Writer
}

// 使用
type File struct {
    name string
}

func (f *File) Read(data []byte) (int, error) {
    // 实现读操作
    return 0, nil
}

func (f *File) Write(data []byte) (int, error) {
    // 实现写操作
    return 0, nil
}