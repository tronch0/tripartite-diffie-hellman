package main

import (
	"crypto/rand"
	"golang.org/x/crypto/bn256"
)

func man(){
	// Generate private keys
	a, _ := rand.Int(rand.Reader, bn256.Order)
	b, _ := rand.Int(rand.Reader, bn256.Order)
	c, _ := rand.Int(rand.Reader, bn256.Order)

	ap := new(bn256.G1).ScalarBaseMult(a)
	aq := new(bn256.G2).ScalarBaseMult(a)

	bp := new(bn256.G1).ScalarBaseMult(b)
	bq := new(bn256.G2).ScalarBaseMult(b)

	cp := new(bn256.G1).ScalarBaseMult(c)
	cq := new(bn256.G2).ScalarBaseMult(c)

	k1 := bn256.Pair(bp, cq)
	k1.ScalarMult(k1, a)

	k2 := bn256.Pair(cp, aq)
	k2.ScalarMult(k2, b)

	k3 := bn256.Pair(ap, bq)
	k3.ScalarMult(k3, c)

	// k1 = k2 = k3

}
