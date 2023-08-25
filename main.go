package main

import (
  "bufio"
  "errors"
  "fmt"
  "log"
  "os"
  "strings"
)

func main() {
  extractDialogue("sub.ass")
}

func extractDialogue(path string) {
  f, err := os.Open(path)
  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
  scanner := bufio.NewScanner(f)
  scanner, err = removePreDialogueInfo(scanner)
  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
  dPos := getDialoguePosition(scanner)
  fmt.Println(dPos, scanner.Text())
}

func removePreDialogueInfo(s *bufio.Scanner) (*bufio.Scanner, error) {
  for s.Scan() {
    if s.Text() == "[Events]" {
      s.Scan()
      return s, nil
    }
  }
  return nil, errors.New("There is no '[Events]' tag in the file.")
}

func getDialoguePosition(s *bufio.Scanner) int {
  pos := strings.Count(s.Text(), ",")
  s.Scan()
  return pos
}
