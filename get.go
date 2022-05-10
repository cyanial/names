package names

import (
	"math/rand"
	"time"
)

func Get() string {
	rand.Seed(time.Now().UnixNano())
	return names[rand.Intn(len(names))]
}
