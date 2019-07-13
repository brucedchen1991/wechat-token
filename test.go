package main


func main()  {
	s := small(12)
	println(s)
}
const smallsString = "00010203040506070809" +
	"10111213141516171819" +
	"20212223242526272829" +
	"30313233343536373839" +
	"40414243444546474849" +
	"50515253545556575859" +
	"60616263646566676869" +
	"70717273747576777879" +
	"80818283848586878889" +
	"90919293949596979899"

//const host32bit = ^uint(0)>>32 == 0

const digits = "0123456789abcdefghijklmnopqrstuvwxyz"
func small(i int) string {
	if i < 10 {
		return digits[i : i+1]
	}
	return smallsString[i*2 : i*2+2]
}
