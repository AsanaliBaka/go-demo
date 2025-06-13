package link

import (
	"go/adv-demo/internal/stat"
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url   string      `json:"url"`
	Hash  string      `json:"Hash" gorm:"uniqueIndex"`
	Stats []stat.Stat `gorm:"constraints:OnUpdate:CASCADE, OnDelete:SET NULL"`
}

func NewLink(url string) *Link {
	return &Link{
		Url:  url,
		Hash: RandStringsRunes(6),
	}
}

var letterRumes = []rune("qwertyuiopasdfghjklzxcvbnm")

func RandStringsRunes(n int) string {

	b := make([]rune, n)

	for i := range b {
		b[i] = letterRumes[rand.Intn(len(letterRumes))]
	}

	return string(b)
}
