/**
 * @Author: cyj19
 * @Date: 2021/11/23 9:25
 */

package main

import (
	"github.com/cyj19/crow/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
