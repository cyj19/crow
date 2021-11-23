/**
 * @Author: vagary
 * @Date: 2021/11/23 9:25
 */

package main

import (
	"github.com/vagaryer/yucli/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
