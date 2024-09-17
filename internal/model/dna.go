package model

import (
	"math/rand"
	"strconv"
)

type DNA struct {
	PointerDNA int
	Array      [lengthDNA]int
}

func (dna DNA) GoString() (stringDNA string) {
	for i := 0; i < len(dna.Array); i++ {
		stringDNA += strconv.Itoa(dna.Array[i]) + ", "
	}
	return stringDNA
}

func Random() DNA {
	var dna DNA
	for i := 0; i < lengthDNA; i++ {
		dna.Array[i] = rand.Intn(lengthDNA - 1)
	}
	dna.PointerDNA = rand.Intn(lengthDNA - 1)
	return dna
}

func (d1 *DNA) Set(d2 DNA) {
	*d1 = d2
}

func (e *Entity) Mutation(count int) {
	for i := 0; i < count; i++ {
		e.DNA.Array[rand.Intn(lengthDNA-1)] = rand.Intn(8)
	}
}
