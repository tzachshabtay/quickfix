package trdgsesgrp

func New() *TrdgSesGrp {
	var m TrdgSesGrp
	return &m
}

//NoTradingSessions is a repeating group in TrdgSesGrp
type NoTradingSessions struct {
	//TradingSessionID is a non-required field for NoTradingSessions.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessions.
	TradingSessionSubID *string `fix:"625"`
}

//TrdgSesGrp is a fix50sp1 Component
type TrdgSesGrp struct {
	//NoTradingSessions is a non-required field for TrdgSesGrp.
	NoTradingSessions []NoTradingSessions `fix:"386,omitempty"`
}

func (m *TrdgSesGrp) SetNoTradingSessions(v []NoTradingSessions) { m.NoTradingSessions = v }