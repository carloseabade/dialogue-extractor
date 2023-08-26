package main

import (
  "fmt"
  "io"
  "log"
  "net/http"
  "os"
)

func main() {
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/upload", uploadHandler)

  fmt.Println("Running server on port 8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  data, err := os.ReadFile("./public/index.html")
  if err != nil {
    log.Fatal(err)
    return
  }
  w.WriteHeader(http.StatusOK)
  w.Header().Add("Content-Type", "text/html")
  w.Header().Add("Content-Length", string(len(data)))
  w.Write(data)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    r.ParseMultipartForm(1 << 20)

    uploadedFile, handler, err := r.FormFile("uploadFile")
    if err != nil {
      log.Fatal(err)
      return
    }
    defer uploadedFile.Close()

    localFile, err := os.OpenFile("./sub.ass", os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
      log.Fatal(err)
      return
    }
    defer localFile.Close()

    _, err = io.Copy(localFile, uploadedFile)
    if err != nil {
      log.Fatal(err)
      return
    }

    dialogues := extractDialogue("sub.ass")
    dialogueFile, err := os.OpenFile("./d"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
      log.Fatal(err)
      return
    }
    defer dialogueFile.Close()

    _, err = dialogueFile.WriteString(dialogues)
    if err != nil {
      log.Fatal(err)
      return
    }

    w.Header().Set("Content-Disposition", "attachment; filename=d"+handler.Filename)
    w.Header().Set("Content-Type", "application/octet-stream")
    http.ServeFile(w, r, "d"+handler.Filename)

    err = os.Remove("d"+handler.Filename)
    if err != nil {
      log.Fatal(err)
      return
    }
    err = os.Remove("sub.ass")
    if err != nil {
      log.Fatal(err)
      return
    }
  } else {
    w.WriteHeader(http.StatusMethodNotAllowed)
    w.Write([]byte(fmt.Sprintf("{\"error\":\"method not allowed\"}")))
  }
}

