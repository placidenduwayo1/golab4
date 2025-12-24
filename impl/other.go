package impl

import (
	"fmt"
	"time"
)

// DoOtherWorks independent function
func DoOtherWorks(u User) {
	time.Sleep(30 * time.Second)
	fmt.Printf("%v %v do other actions\n", u.FirstName, u.LastName)
	Wg.Done()
}
