package power

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func GetCPU(power chan<- float64) {
	cmd := "sudo powermetrics --samplers cpu_power -n 1 -i 500 | grep \"CPU Power:\""
	out, err := exec.Command("bash","-c",cmd).Output()
	
	if err != nil {
        log.Fatal(err)
    }
	num := strings.Split(string(out), " ")[2]
    p, err := strconv.ParseFloat(num,64)
	power<-p


}


func GetGPU(power chan<- float64) {
	cmd := "sudo powermetrics --samplers cpu_power -n 1 -i 500 | grep \"GPU Power:\""
	out, err := exec.Command("bash","-c",cmd).Output()
	
	if err != nil {
        log.Fatal(err)
    }
	num := strings.Split(string(out), " ")[2]
    p, err := strconv.ParseFloat(num,64)
	power<-p


}

