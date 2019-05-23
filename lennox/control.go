package lennox

import "strconv"
import "fmt"
import "log"
import "strings"

const TIME_SHORT = 550
const TIME_LONG = 1550
const TIME_4000 = 4350
const TIME_5000 = 5150

const COOL_MODE = 0
type CoolState struct {
	Temperature int
	FanSpeed    FanSpeed
}

func (f FanSpeed) Data() (string) {
	return fmt.Sprintf("%03s", strconv.FormatInt(int64(f), 2))
}

func (s *CoolState) Data() (string) {
	f := s.FanSpeed.Data()
	m := strconv.FormatInt(int64(COOL_MODE), 2)
	t := strconv.FormatInt(int64(s.Temperature-17), 2)
	d := fmt.Sprintf("1010000110%03s%03s0100%04s1111111111111111",f,m,t)
	return d
}

func flip(s string) (string) {
	newString := strings.Replace(s, "0", "2", -1)
	newString = strings.Replace(newString, "1", "0", -1)
	newString = strings.Replace(newString, "2", "1", -1)
	return newString
}

func reverse(s string) string {
  runes := []rune(s)
  for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
  }
  return string(runes)
}

func checksum(s string) (string) {
	var sum uint64 = 0

	for i := 0; i < 5; i++ {
		b := s[8*i:8*i+8]
		b = reverse(b)
		i, err := strconv.ParseUint(b, 2, 8)

		if err != nil {
			log.Fatal(err)
		}
		sum += i
	}
	sum = 256 - sum % 256

	out := strconv.FormatUint(sum, 2)
	out = fmt.Sprintf("%08s",out)
	out = reverse(out)

	return out
}

func encode(data string) ([]uint) {
	var s []uint
	s = append(s, TIME_4000, TIME_4000)

	for _, v := range data {
		s = append(s, TIME_SHORT)
		switch(v) {
		case '0':
			s = append(s, TIME_SHORT)
		case '1':
			s = append(s, TIME_LONG)
		}
	}
	s = append(s, TIME_SHORT)
	s = append(s, TIME_5000, TIME_4000, TIME_4000)

	for _, v := range data {
		s = append(s, TIME_SHORT)
		switch(v) {
		case '0':
			s = append(s, TIME_LONG)
		case '1':
			s = append(s, TIME_SHORT)
		}
	}
	s = append(s, TIME_SHORT)

	return s
}

func Apply(state *CoolState) (error) {
	data := state.Data()
	data += checksum(data)

	encodedData := encode(data)
	fmt.Printf("%s\n", encodedData)
	return nil
}
