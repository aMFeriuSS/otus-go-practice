package ntp

import "fmt"
import "time"
import "github.com/beevik/ntp"

func main() {
	time, err := ntp.Time("time.google.com")
	if err != nil {
		fmt.Println("Something is wrong")
		panic(err)
	}

	fmt.Println(time)
}

func Time() {
	time := time.Now()

	fmt.Println(time)
}

func NtpTime() {
	time, err := ntp.Time("time.google.com")
	if err != nil {
		fmt.Println("Something is wrong")
		panic(err)
	}

	fmt.Println(time)

}
