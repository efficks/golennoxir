package main

import (
	"encoding/json"
	"errors"
	"github.com/efficks/lennoxHeatPump/lennox"
)

type Temperature struct {
	T int
}

func (t *Temperature) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &t.T)
	if err != nil {
		return err
	}
	return nil
}

type IState interface {
	Convert() (*lennox.IState, error)
}

type CoolState struct {
	Temperature int    `json:"temperature"`
	FanSpeed    string `json:"fanSpeed"`
}

func (s CoolState) Convert() (*lennox.CoolState, error) {
	fs, e := lennox.ToFanSpeed(s.FanSpeed)
	if e != nil {
		return nil, e
	}
	if s.Temperature < 18 || s.Temperature > 30 {
		return nil, errors.New("Temperature must be between 18 and 30")
	}
	return &lennox.CoolState{s.Temperature, fs}, nil
}

func (s HeatState) Convert() (*lennox.HeatState, error) {
	fs, e := lennox.ToFanSpeed(s.FanSpeed)
	if e != nil {
		return nil, e
	}
	if s.Temperature < 18 || s.Temperature > 30 {
		return nil, errors.New("Temperature must be between 18 and 30")
	}
	return &lennox.HeatState{s.Temperature, fs}, nil
}

func (s DryState) Convert() (*lennox.DryState, error) {
	if s.Temperature < 18 || s.Temperature > 30 {
		return nil, errors.New("Temperature must be between 18 and 30")
	}
	return &lennox.DryState{s.Temperature}, nil
}

func (s FanState) Convert() (*lennox.FanState, error) {
	fs, e := lennox.ToFanSpeed(s.FanSpeed)
	if e != nil {
		return nil, e
	}
	return &lennox.FanState{fs}, nil
}

type HeatState struct {
	Temperature int    `json:"temperature"`
	FanSpeed    string `json:"fanSpeed"`
}

type FanState struct {
	FanSpeed string `json:"fanSpeed"`
}

type DryState struct {
	Temperature int    `json:"temperature"`
	FanSpeed    string `json:"fanSpeed"`
}

type OffState struct {
}
/*
type AutoState struct {
	Temperature int    `json:"temperature"`
	FanSpeed    string `json:"fanSpeed"`
}*/
