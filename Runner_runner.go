package main

import (
  "bytes"
  "fmt"
  "log"
  "os/exec"
  "strings"
  "time"
  "gopkg.in/gomail.v2"
)

func main() {
  output := run_command("ls -lah")
  // delay(1)
  fmt.Println(output)
}

func read_command_from_file() {

}

func run_command(command string) string {
  commands_split := strings.Split(command, " ")
  cmd := exec.Command(commands_split[0], commands_split[1:]...)
  var out bytes.Buffer
  cmd.Stdout = &out
  err := cmd.Run()
  c(err)
  return strings.TrimSpace(out.String())
}

func delay(minutes int) {
  time.Sleep(time.Duration(minutes) * time.Minute)
}

func c(err error) {
  if err != nil {
		log.Fatal(err)
	}
}
