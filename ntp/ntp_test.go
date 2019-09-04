package ntp_test

import (
	"testing"

	"github.com/amferiuss/otus-go-practice/ntp"
)

func Test_Time(t *testing.T) {
	ntp.Time()
	ntp.NtpTime()
}