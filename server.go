package logr

import (
	"log"
	"net/http"
	"strconv"
)

// StartServe - starts a Logr server [OPTIONAL]
func StartServe() {
	http.HandleFunc("/write/severe", severeWriteHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/read", readHandler)
	panic(http.ListenAndServe(":8080", nil))
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

	vals := r.Form
	priorityStr := vals.Get("priority")
	priority, err := strconv.Atoi(priorityStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error encountered while trying to handle write: %v\n", err)
	}

	msg := vals.Get("msg")

	if err := writer.Log(priority, msg); err != nil {
		w.WriteHeader(http.StatusNotModified)
		log.Printf("Error encountered while trying to handle write: %v\n", err)
	}

	w.Write([]byte("Log successfully written"))
	w.WriteHeader(http.StatusCreated)
}

func severeWriteHandler(w http.ResponseWriter, r *http.Request) {
	writer, err := NewSevereLogWriter()
	if err != nil {
		log.Fatal(err)
	}

	vals := r.Form
	msg := vals.Get("msg")

	if err := writer.Log(msg); err != nil {
		w.WriteHeader(http.StatusNotModified)
		log.Printf("Error encountered while trying to handle write: %v\n", err)
	}

	w.Write([]byte("Log successfully written"))
	w.WriteHeader(http.StatusCreated)
}
