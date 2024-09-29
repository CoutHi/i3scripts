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
        artistStr := strings.TrimSpace(string(output))
        if len(artistStr) > 15 {
            collection.artist = artistStr[:15]+"..."
        }else {
            collection.artist = artistStr
        }
    }


    album := exec.Command("sh", "-c", "playerctl metadata album")
    output, err = album.Output()
    if err != nil {
        collection.album = ""
    }else {
        albumStr := strings.TrimSpace(string(output))
        if len(albumStr) > 15 {
            collection.album = albumStr[:15]+"..."
        }else {
            collection.album = albumStr
        }
    }

    title := exec.Command("sh", "-c", "playerctl metadata title")
    output, err = title.Output()
    if err != nil {
        collection.title = ""
    }else {
        titleStr := strings.TrimSpace(string(output))
        if len(titleStr) > 15 {
            collection.title = titleStr[:15]+"..."
        }else {
            collection.title = titleStr
        }
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
