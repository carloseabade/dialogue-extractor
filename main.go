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
  err = removePreDialogueInfo(scanner)
  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
  dialogueLines := getDialogues(scanner)
  fmt.Println(strings.Join(dialogueLines, "\n"))
}

func removePreDialogueInfo(s *bufio.Scanner) error {
  for s.Scan() {
    if s.Text() == "[Events]" {
      s.Scan()
      return nil
    }
  }
  return errors.New("There is no '[Events]' tag in the file.")
}

func getDialoguePosition(s *bufio.Scanner) int {
  pos := strings.Count(s.Text(), ",")
  return pos
}

func getDialogues(s *bufio.Scanner) []string {
  dPos := getDialoguePosition(s)

  var lines []string
  var line string
  for s.Scan() {
    line = strings.Join(strings.Split(s.Text(), ",")[dPos:] ,"")
    lines = append(lines, line)
  }
  return lines
}
