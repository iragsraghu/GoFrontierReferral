package serialNumber

import (
	"fmt"
	"os/exec"
	"strings"
)

func SerialNumber(num string) string {
	var serialNumber string
	out, _ := exec.Command("/usr/sbin/ioreg", "-l").Output() // err ignored for brevity
	for _, l := range strings.Split(string(out), "\n") {
		if strings.Contains(l, "IOPlatformSerialNumber") {
			s := strings.Split(l, " ")
			serialNumber = s[len(s)-1]
			fmt.Println(serialNumber)
			break
		}
	}
	fmt.Println("serial:", serialNumber)
	return serialNumber
}
