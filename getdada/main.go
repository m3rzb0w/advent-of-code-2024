package getdata

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Getdata(url string, sessionCookie string) string {

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Cookie", sessionCookie)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("[INFO] Req status : ", res.Status)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)
}

func Grabsession() string {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	session := os.Getenv("COOKIE")
	fmt.Println("[INFO] Current cookie : ", session)
	return session
}

func Filename(urlStr string) string {
	u, err := url.Parse(urlStr)
	if err != nil {
		log.Fatal("[INFO] Err parsing url")
	}
	urlSplit := strings.Split(u.Path, "/")
	year := urlSplit[1]
	day := urlSplit[3]
	fileName := fmt.Sprintf("%s_day_%s.txt", year, day)
	return fileName
}

func CheckFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func Getfile(fileExists bool, fileName string, url string, sessionCookie string) string {
	var data string
	if !fileExists {
		data = Getdata(url, sessionCookie)
		f, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		_, err = f.WriteString(data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[INFO] Writting file : %s completed!\n", fileName)
		return data
	}
	fmt.Printf("[INFO] Reading file : %s...\n", fileName)
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
