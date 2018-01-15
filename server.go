package logr

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fatih/color"
)

// Serve - starts a Logr server [OPTIONAL]
func Serve() {
	http.HandleFunc("/write/severe", severeWriteHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/read", readHandler)

	color.Green("Starting logr server...")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		color.Red(err.Error())
	}
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	reader, err := NewLogReader()
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(reader.Get()))
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	writer, err := NewGenericLogWriter()
	if err != nil {
		log.Fatal(err)
	}

	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	vals := r.PostForm

	priorityStr := vals.Get("priority")
	priority, err := strconv.Atoi(priorityStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error encountered while trying to handle write: %v\n", err)
		return
	}

	msg := vals.Get("msg")

	if err := writer.Log(priority, msg); err != nil {
		w.WriteHeader(http.StatusNotModified)
		log.Printf("Error encountered while trying to handle write: %v\n", err)
		return
	}

	w.Write([]byte("Log successfully written"))
}

func severeWriteHandler(w http.ResponseWriter, r *http.Request) {
	writer, err := NewSevereLogWriter()
	if err != nil {
		log.Fatal(err)
	}

	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	vals := r.PostForm

	msg := vals.Get("msg")

	if err := writer.Log(msg); err != nil {
		w.WriteHeader(http.StatusNotModified)
		log.Printf("Error encountered while trying to handle write: %v\n", err)
		return
	}

	w.Write([]byte("Log successfully written"))
}
