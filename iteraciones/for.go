package iteraciones

import (
	"fmt"
)

func Iterar() {
	for i := 100; i > 10; i-=5 {
		if i == 20 {
			continue
		}
		fmt.Println(i)
	}
}
