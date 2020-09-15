package main

import (
    "bufio"
    "fmt"
    rice "github.com/GeertJohan/go.rice"
    "io"
    "io/ioutil"
)

func main() {
    rice.Debug = true

    box, err := rice.FindBox("./sql")
    if err != nil {
        panic(err)
    }

    file, err := box.Open("20200915113656_test1.up.sql")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    bufferSize := 200

    b := bufio.NewReaderSize(file, bufferSize)
    if _, err = b.Peek(bufferSize); err != nil && err != io.EOF {
        panic(err)
    }

    br, bw := io.Pipe()
    go func() {
        fmt.Println("Begin to write to pipe bw")
        n, err := b.WriteTo(bw)
        fmt.Printf("WriteTo wrote %v bytes, err=%v\n", n, err)
    }()

    content, err := ioutil.ReadAll(br)
    if err != nil {
        fmt.Printf("Error when reading all content through pipe: %v\n", err)
    }
    fmt.Printf("ReadAll got content: %v", content)

    fmt.Println("Test finished")
}
