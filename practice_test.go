package practice

import "testing"
import "github.com/amferiuss/otus-go-practice/ntp"

// Week-1: ntp.
func Test_NtpTime(t *testing.T) {
	ntp.Time()
	ntp.NtpTime()
}