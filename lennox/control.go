package lennox

import "strconv"
import "fmt"
import "strings"
import "io/ioutil"
import "os"
import "io"
import "os/exec"

const TIME_SHORT = 550
const TIME_LONG = 1550
const TIME_4000 = 4350
const TIME_5000 = 5150

const COOL_MODE = 0
const DRY_MODE = 1
const AUTO_MODE = 2
const HEAT_MODE = 3
const FAN_MODE = 4

type IState interface {
    Data() string
}

type CoolState struct {
	Temperature int
	FanSpeed    FanSpeed
}
type HeatState struct {
	Temperature int
	FanSpeed    FanSpeed
}
type OffState struct {
}
type FanState struct {
	FanSpeed    FanSpeed
}
type DryState struct {
	Temperature int
}


func (f FanSpeed) Data() string {
	return fmt.Sprintf("%03s", strconv.FormatInt(int64(f), 2))
}

func (s CoolState) Data() string {
	f := s.FanSpeed.Data()
	m := strconv.FormatInt(int64(COOL_MODE), 2)
	t := strconv.FormatInt(int64(s.Temperature-17), 2)
	d := fmt.Sprintf("1010000110%03s%03s0100%04s1111111111111111", f, m, t)
	return d
}

func (s HeatState) Data() string {
	f := s.FanSpeed.Data()
	m := strconv.FormatInt(int64(HEAT_MODE), 2)
	t := strconv.FormatInt(int64(s.Temperature-17), 2)
	d := fmt.Sprintf("1010000110%03s%03s0100%04s1111111111111111", f, m, t)
	return d
}

func flip(s string) string {
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

func checksum(s string) (string,error) {
	var sum uint64 = 0

	for i := 0; i < 5; i++ {
		b := s[8*i : 8*i+8]
		b = reverse(b)
		i, err := strconv.ParseUint(b, 2, 8)

		if err != nil {
			return "",err
		}
		sum += i
	}
	sum = 256 - sum%256

	out := strconv.FormatUint(sum, 2)
	out = fmt.Sprintf("%08s", out)
	out = reverse(out)

	return out,nil
}

func encode(data string) []uint {
	var s []uint
	s = append(s, TIME_4000, TIME_4000)

	for _, v := range data {
		s = append(s, TIME_SHORT)
		switch v {
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
		switch v {
		case '0':
			s = append(s, TIME_LONG)
		case '1':
			s = append(s, TIME_SHORT)
		}
	}
	s = append(s, TIME_SHORT)

	return s
}

func Apply(state IState) error {
	data := state.Data()
	chk,err := checksum(data)
	if err != nil {
		return err
	}
	data += chk

	encodedData := encode(data)
	fmt.Printf("%s\n",encodedData)

	tmpfile, err := ioutil.TempFile("", "lennox")
	if err != nil {
		return err
	}
	defer os.Remove(tmpfile.Name())

	var on bool = false
	var s string
	for _, t := range encodedData{
		if(on) {
			s = fmt.Sprintf("space %d\n", t)
		} else {
			s = fmt.Sprintf("pulse %d\n", t)
		}

		_,err=io.WriteString(tmpfile,s)
		if err != nil {
			return err
		}
		on = !on
	}

	cmd := exec.Command("ir-ctl", "--send", tmpfile.Name())
	err = cmd.Run()

	return err
}
