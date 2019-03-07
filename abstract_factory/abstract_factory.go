package main

import "fmt"

// interface
type Army interface {
	getDescription() string
}

type Castle interface {
	getDescription() string
}

type King interface {
	getDescription() string
}

// Elven
type ElfArmy struct {
}
func (e *ElfArmy) getDescription() string {
	return "This is the Elven Army!"
}

type ElfCastle struct {
}
func (e *ElfCastle) getDescription() string {
	return "This is the Elven Castle!"
}

type ElfKing struct {
}
func (e *ElfKing) getDescription() string {
	return "This is the Eleven King!"
}

// Orc
type OrcArmy struct {
}
func (e *OrcArmy) getDescription() string {
	return "This is the Orc Army!"
}

type OrcCastle struct {
}
func (e *OrcCastle) getDescription() string {
	return "This is the Orc Castle!"
}

type OrcKing struct {
}
func (e *OrcKing) getDescription() string {
	return "This is the Orc King!"
}

// Kingdom
type KingdomFactory interface {
	createCastle() Castle
	createKing()   King
	createArmy()   Army
}

type elfKingdomFactory struct {
}

func NewElfKingdomFactory() KingdomFactory {
	return &elfKingdomFactory{}
}

func(e *elfKingdomFactory) createCastle() Castle {
	return &ElfCastle{}
}

func(e *elfKingdomFactory) createKing() King {
	return &ElfKing{}
}

func(e *elfKingdomFactory) createArmy() Army {
	return &ElfArmy{}
}

// Concrete ファクトリー
type orcKingdomFactory struct {
}

func NewOrcKingdomFactory() KingdomFactory {
	return &orcKingdomFactory{}
}

func(e *orcKingdomFactory) createCastle() Castle {
	return &OrcCastle{}
}

func(e *orcKingdomFactory) createKing() King {
	return &OrcKing{}
}

func(e *orcKingdomFactory) createArmy() Army {
	return &OrcArmy{}
}

func main() {
	elfKingdam := NewElfKingdomFactory()
	elfCastle  := elfKingdam.createCastle()
	elfKng     := elfKingdam.createKing()
	elfArmy    := elfKingdam.createArmy()
	fmt.Println(elfCastle.getDescription())
	fmt.Println(elfKng.getDescription())
	fmt.Println(elfArmy.getDescription())

	orkKingdam := NewOrcKingdomFactory()
	orkCastle  := orkKingdam.createCastle()
	orkKng     := orkKingdam.createKing()
	orkArmy    := orkKingdam.createArmy()
	fmt.Println(orkCastle.getDescription())
	fmt.Println(orkKng.getDescription())
	fmt.Println(orkArmy.getDescription())
}