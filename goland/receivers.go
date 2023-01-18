package main

import (
	"fmt"
)

type Kullanici struct {
	Ad      string `json:"adi"`
	Soyad   string `json:"-"`
	Takipci []Kullanici
}

func (k Kullanici) TakipciSayisi() int {
	return len(k.Takipci)
}

func (k *Kullanici) TakipciEkle(t Kullanici) {
	if k.Takipci == nil {
		k.Takipci = []Kullanici{}
	}

	k.Takipci = append(k.Takipci, t)
}

func main() {
	k := Kullanici{
		Ad:    "Go",
		Soyad: "Turkıye",
		Takipci: []Kullanici{
			{
				Ad:    "Takipci",
				Soyad: "1",
			},
			{
				Ad:    "Takipci",
				Soyad: "2",
			},
		},
	}

	t := Kullanici{
		Ad:    "Takipci",
		Soyad: "-",
	}

	fmt.Println(k.TakipciSayisi())

	k.TakipciEkle(t)

	fmt.Println(k.TakipciSayisi())
}
