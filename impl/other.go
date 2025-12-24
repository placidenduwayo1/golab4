package impl

import (
	"fmt"
	"time"
)

// DoOtherWorks independent function
func (u User) DoOtherWorks() {
	time.Sleep(30 * time.Second)
	fmt.Printf("%v %v do other actions\n", u.FirstName, u.LastName)
	Wg.Done()
}
