package main

import (
    "os/exec"
    "fmt"
    "strings"
)

func main(){
    artist := exec.Command("sh", "-c", "playerctl metadata artist")
    output, err := artist.Output()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    var collection Media;

    collection.artist = strings.TrimSpace(string(output))

    album := exec.Command("sh", "-c", "playerctl metadata album")
    output, err = album.Output()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    collection.album = strings.TrimSpace(string(output))

    title := exec.Command("sh", "-c", "playerctl metadata title")
    output, err = title.Output()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    collection.title = strings.TrimSpace(string(output))

    if collection.album != "" {
        fmt.Print("[",collection.artist,"]"," [",collection.album,"] ","[",collection.title,"]")
    }else {
        fmt.Print("[",collection.artist,"]"," [",collection.title,"]")
    }
}

type Media struct {
    artist string
    album string
    title string
}
