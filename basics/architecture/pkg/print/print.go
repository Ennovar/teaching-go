// This package helps print things to various
// mediums.
package print

import "log"

func PrintToConsole(print interface{}) {
	log.Println(print)
}