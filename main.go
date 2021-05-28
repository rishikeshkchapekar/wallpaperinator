package main

import(
	"wallpaperinator/server"
)

func main(){
	resp := server.GetImages()
	server.DownloadImage(resp)
}