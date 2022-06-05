package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	//out := runCommand("schtasks", `/create`, `/f`, `/tn`, `tnetupdater`, `/tr`, `"\"C:\Program Files\test.exe\" run"`, `/sc`, `minute`, `/mo`, `60`)
	//log.Println(out)
	//f, err := ini.Load(`D:\test.ini`)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//command := f.Section("").Key("command").MustString("")
	//log.Println("Command: ", command)
	execpath, err := os.Executable()
	if err != nil {
		log.Println(err)
	}
	stdout:= runCommand("schtasks", `/create`, `/f`, `/tn`, `tnetupdater`, `/tr`, fmt.Sprintf(`"%st" run`, execpath), `/sc`, `minute`, `/mo`, `60`)
	log.Println(stdout)
}

func run(command ...string) string{
	//append([]string{"/c", "chcp", "437", "&&"}, command...)...
	stdout, err := exec.Command("C:\\Windows\\System32\\cmd.exe", append([]string{"/c", "chcp", "437", "&&"}, command...)...).Output()
	if err != nil {
		log.Println(err)
	}
	return string(stdout)
}

func runCommand(command string, args ...string) string{
	stdout, err := exec.Command("C:\\Windows\\System32\\cmd.exe", append([]string{"/c", "chcp", "437", "&&", command}, args...)...).Output()
	if err != nil {
		log.Println(err)
	}
	return string(stdout)
}
