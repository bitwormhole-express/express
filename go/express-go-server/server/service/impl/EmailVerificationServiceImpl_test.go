package impl

import "testing"

func TestEmailVerificationServiceImpl(t *testing.T) {

	x := EmailVerificationServiceImpl{}
	// x.Init()
	x.secret = "123456789"
	x.mailContentType = "text/plain"
	x.mailContentTemplate = "hello"

	nonce := x.generateNonce()

	vcode := x.generateVeriCode()

	t.Log(x.secret)
	t.Log(nonce)
	t.Log(vcode)
}

func TestEmailVerificationServiceImplVCode(t *testing.T) {
	table := make(map[string]int64)
	x := EmailVerificationServiceImpl{}
	for count := 100000; count > 0; count-- {
		vcode := x.generateVeriCode()
		ch0 := vcode[0:1]
		table[ch0]++
	}
	t.Log(table)
}
