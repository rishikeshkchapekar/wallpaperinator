package server

import(
	"net/http"
	"encoding/json"
	"fmt"
	"log"
	"bytes"
	"io"
	"io/ioutil"
	"math/rand"
	"time"
	"wallpaperinator/data"
	"github.com/google/uuid"
	"os"
	"os/exec"
	"strings"
)

func GetImages() *data.Response{
	var respData *data.Response

	unsplash_key := os.Getenv("UNSPLASH_API")

	queryName := "black abstract"
	_a := strings.Split(queryName," ")
	if len(_a)==1{
		queryName = _a[0]
	}else{
		queryName = strings.Join(_a[:], "+")
	}

	url := "https://api.unsplash.com/search/photos?page=1&query="+queryName+"&per_page=100&client_id="+unsplash_key
	resp,err := http.Get(url)
	if err!=nil{
		log.Fatalf("Error occured %v",err)
	}
	defer resp.Body.Close()

	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Fatalln(err)
	}
	err = json.Unmarshal(body,&respData)

	if err!=nil{
		log.Panic(err)
	}

	return respData
}

func DownloadImage(resp_data *data.Response){
	allRes := resp_data.Results

	rand.Seed(time.Now().UnixNano())
	
	num := len(allRes)
	index := rand.Intn(num)
	
	first := allRes[index]

	rawUrl := first.Urls.Raw
	fmt.Println("Image URL: ",rawUrl)

	downloadImageFromUrl(rawUrl)
}

func downloadImageFromUrl(url string){
	resp,err := http.Get(url)
	if err!=nil{
		log.Fatal(err)
	}

	defer resp.Body.Close()

	path := "wallpaperinator_images"
	if _, err := os.Stat(path); os.IsNotExist(err) {
	    os.Mkdir(path, 0777)
	}
	fname := getUuid()

	completePath := path+"/"+fname+".jpg"
	
	file,err:=os.Create(completePath)
	if err!=nil{
		log.Fatal("Error creating file",err)
	}
	defer file.Close()

	_,err=io.Copy(file,resp.Body)
	if err!=nil{
		log.Fatal("Error doing that copy thing",err)
	}

	fmt.Println("Downloaded Image Successfully")
	setWallpaper(completePath)
}

func setWallpaper(imgPath string){

	var pwd bytes.Buffer
	
	cmd := exec.Command("pwd")
	
	cmd.Stdout = &pwd
	
	err := cmd.Run()
	if err!=nil{
		log.Fatal(err)
	}

	curpath := pwd.String()
	curpath = strings.Replace(curpath,"\n","",-1)

	arg:="file://"+curpath+"/"+imgPath

	cmd = exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", arg)
	err = cmd.Run()
	if err!=nil{
		log.Fatal("Error running wallpaper command: ",err)
	}

}

func getUuid() string{
	uuidWithHyphen := uuid.New()
    fmt.Println(uuidWithHyphen)
    uuidNew := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
    return uuidNew
}