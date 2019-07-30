package main

import (
	//"HFish/utils/setting"
	//"fmt"
	//"os"
	"HFish/core/protocol/ssh"
)

func main() {
	ssh.Start()
	//setting.Run("weibo", "/", "all")
	//
	//args := os.Args
	//if args == nil || len(args) < 2 {
	//	setting.Help()
	//} else {
	//	if args[1] == "help" || args[1] == "--help" {
	//		setting.Help()
	//	} else if args[1] == "init" || args[1] == "--init" {
	//		setting.Init()/**/
	//	} else if args[1] == "version" || args[1] == "--version" {
	//		fmt.Println("0.1")
	//	} else if args[1] == "run" || args[1] == "--run" {
	//		if len(args) >= 3 {
	//			setting.Run("weibo", "/", "all")
	//		} else {
	//			setting.Run("weibo", "/", "all")
	//		}
	//	} else {
	//		setting.Help()
	//	}
	//}
}
