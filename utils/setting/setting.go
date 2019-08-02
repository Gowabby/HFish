package setting

import (
	"HFish/core/exec"
	"HFish/utils/color"
	"HFish/view"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"net/http"
	"time"
)

func RunHtml(project string, url string) http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())

	// 引入html资源
	r.LoadHTMLGlob("web/" + project + "/*")

	// 引入静态资源
	r.Static("/static", "./static/"+project)

	r.GET(url, func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	return r
}

func RunWeb() http.Handler {
	gin.DisableConsoleColor()
	f, _ := os.Create("./logs/hfish.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// 引入gin
	r := gin.Default()

	r.Use(gin.Recovery())
	// 引入html资源
	r.LoadHTMLGlob("admin/*")

	// 引入静态资源
	r.Static("/static", "./static")

	// 加载路由
	view.LoadUrl(r)

	return r
}

func Run(project string, url string, types string) {
	server01 := &http.Server{
		Addr:         ":9001",
		Handler:      RunWeb(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":9000",
		Handler:      RunHtml(project, url),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	switch types {
	case "all":
		// 前后端全部启动
		go server01.ListenAndServe()
		server02.ListenAndServe()
	case "d":
		// 启动后端
		server02.ListenAndServe()
	}
}

func Init() {
	fmt.Println("test")
}

func Help() {
	exec.Execute("clear")
	logo := ` o
  \_/\o
 ( Oo)                    \|/
 (_=-)  .===O- ~~~Z~A~P~~ -O-
 /   \_/U'                /|\
 ||  |_/
 \\  |	     ~ By: HackLC Team
 {K ||       __ _______     __
  | PP      / // / __(_)__ / /
  | ||     / _  / _// (_-</ _ \
  (__\\   /_//_/_/ /_/___/_//_/ v0.1
`
	fmt.Println(color.Yellow(logo))
	fmt.Println(color.White(" An Active Attack Honeypot System for Fishing."))
	fmt.Println("")
	fmt.Println(color.Yellow(" + [ ABOUT ]----------------------------------------------------------- +"))
	fmt.Println("")
	fmt.Println(color.Green("   - Github:"), color.White("https://github.com/hacklcs/HFish"), color.Green(" - Team:"), color.White("https://hack.lc"))
	fmt.Println("")
	fmt.Println(color.Yellow(" + [ ARGUMENTS ]------------------------------------------------------- +"))
	fmt.Println("")
	fmt.Println(color.Cyan("   run,--run"), color.White("	       Start up service"))
	//fmt.Println(color.Cyan("   init,--init"), color.White("		   Initialization, Wipe data"))
	fmt.Println(color.Cyan("   version,--version"), color.White("  HFish Version"))
	fmt.Println(color.Cyan("   help,--help"), color.White("	       Help"))
	fmt.Println("")
	fmt.Println(color.Yellow(" + -------------------------------------------------------------------- +"))
	fmt.Println("")
}













