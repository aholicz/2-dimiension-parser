package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type Person struct {
	Name  string
	Age   int64
	Speed float64
}

type Speed struct {
	SpeedFloat64 float64
}

func (s *Speed) UnmarshalJSON(bytes []byte) error {
	var speed string
	err := json.Unmarshal(bytes, &speed)
	if err != nil {
		return err
	}

	f, err := strconv.ParseFloat(speed, 64)
	if err != nil {
		return err
	}

	s.SpeedFloat64 = f

	return nil
}

func (p *Person) UnmarshalJSON(body []byte) error {
	var tmp []json.RawMessage
	if err := json.Unmarshal(body, &tmp); err != nil {
		return err
	}

	if err := json.Unmarshal(tmp[0], &p.Name); err != nil {
		return err
	}

	if err := json.Unmarshal(tmp[1], &p.Age); err != nil {
		return err
	}

	speed := new(Speed)
	if err := json.Unmarshal(tmp[2], &speed); err != nil {
		return err
	}

	p.Speed = speed.SpeedFloat64

	return nil
}

func main() {
	var data = []byte(`[[ "Alice", 18, "1.234" ],[ "Bob", 40, "4.567"]]`)
	var persons []Person
	if err := json.Unmarshal(data, &persons); err != nil {
		log.Fatal(err)
	}

	fmt.Print(persons)
}
