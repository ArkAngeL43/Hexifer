package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var (
	flagHelp = flag.Bool("h", false, `Print the help menu and exit`)
	flagFile = flag.String("file", "", `File`)
)

func isflagged() {
	flag.Parse()
	//
	if *flagHelp {
		fmt.Println("-> Usages | go run main.go -file <filename> -byte <1-1000> ")
		os.Exit(1)
	}
	//
}

func e(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func banner() {
	content, err := os.ReadFile("banner.txt")
	e(err)
	fmt.Println("\033[35m", string(content))
}

func clear() {
	if runtime.GOOS == "windows" {
		cls, err := exec.Command("cls").Output()
		if err != nil {
			log.Fatal(err)
		}
		output := string(cls[:])
		fmt.Println(output)
	}
	if runtime.GOOS == "linux" {
		clear, err := exec.Command("clear").Output()
		e(err)
		output := string(clear[:])
		fmt.Println(output)
	}
}

func main() {
	clear()
	banner()
	flag.Parse()
	isflagged()
	r, err := os.Stat(*flagFile)
	e(err)
	re := r.Size()
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("\033[34m[ + ] File Being Read ~> ", *flagFile)
	fmt.Println("\033[34m[ + ] File Byte Size  ~> ", re)
	fmt.Println("\033[33m[ * ] WARN: To prevent stack/buffer overflow hex dump value will use read bytes of the current file")
	fmt.Println("\033[34m[ + ] Logged Out File ~> log.txt")
	fmt.Println("\033[35m")
	time.Sleep(2 * time.Second)
	file, err := os.Open(*flagFile)
	e(err)
	defer file.Close()
	//
	//
	read := bufio.NewReader(file)
	buf := make([]byte, re)
	for {
		_, err := read.Read(buf)
		e(err)
		break
	}
	fmt.Printf("%s", hex.Dump(buf))
	cr, err := os.Create("log.txt")
	e(err)
	defer cr.Close()
	_, err2 := cr.WriteString(hex.Dump(buf))
	e(err2)
}

////////////////////// filing

//	cr, err := os.Create("log.txt")
//	e(err)
//	defer cr.Close()
//	_, err2 := cr.WriteString(hex.Dump(buf))
//	e(err2)
//	fmt.Printf("%s", hex.Dump(buf))
//	filename := "log.txt"
//	fmt.Println("\033[32m[+] Data logged to a file ~> ", filename)
