package main

import (
	"bytes"
  "bufio"
	"fmt"
	"log"
  "os"
	"os/exec"
	"strings"
	"time"
)

func main() {
  text_arr := read_commands_from_file("commands_sim/commands1.in")
  fmt.Println(strings.Join(text_arr, "\n"))

	// output := run_command("ls -lah")
	// delay(1)
	// fmt.Println(output)
}

func executer(command string) {
  
}

func read_commands_from_file(filepath string) []string {
  file, err := os.Open(filepath)
  c(err)
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    if len(scanner.Text()) > 0 {
      lines = append(lines, scanner.Text())
    }

  }
  c(scanner.Err())
  return lines
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

func alert() {
	send_text_alert()
}
