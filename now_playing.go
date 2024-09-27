package main

import (
    "os/exec"
    "fmt"
    "strings"
)

func main(){
    
    var collection Media

    artist := exec.Command("sh", "-c", "playerctl metadata artist")
    output, err := artist.Output()
    if err != nil {
        collection.artist = ""
    }else {
        collection.artist = strings.TrimSpace(string(output))
    }


    album := exec.Command("sh", "-c", "playerctl metadata album")
    output, err = album.Output()
    if err != nil {
        collection.album = ""
    }else {
        collection.album = strings.TrimSpace(string(output))
    }

    title := exec.Command("sh", "-c", "playerctl metadata title")
    output, err = title.Output()
    if err != nil {
        collection.title = ""
    }else {
        collection.title = strings.TrimSpace(string(output))
    }

    if collection.album == "" && collection.title == "" && collection.artist == "" {
        fmt.Print("[]")
    }else if collection.album != "" {
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
