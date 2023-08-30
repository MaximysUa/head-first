package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}
func newHandler(writer http.ResponseWriter, request *http.Request) {
	files, err := template.ParseFiles("templates/new.html")
	check(err)
	err = files.Execute(writer, nil)
	check(err)
}
func createHandler(writer http.ResponseWriter, request *http.Request) {
	value := request.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	defer file.Close()
	check(err)
	_, err = fmt.Fprintln(file, value)
	check(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

func deleteHandler(writer http.ResponseWriter, request *http.Request) {

}
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	files, err := template.ParseFiles("templates/view.html")
	check(err)
	err = files.Execute(writer, guestbook)
	check(err)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
