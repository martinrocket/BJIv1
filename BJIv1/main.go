package main

import (
    "fmt"
    "os/exec"
    "os"
    "runtime"
)

func main(){
    clearScreen(runtime.GOOS)
    fmt.Println(startApp())
}

func startApp() string {
    return "Starting..."
}

func clearScreen(s string) string {
    switch s {
        case "linux":
            cmd := exec.Command("clear")
            cmd.Stdout = os.Stdout
            cmd.Run()
            fmt.Println(s)
            return s
        case "darwin":
            cmd := exec.Command("clear")
            cmd.Stdout = os.Stdout
            cmd.Run()
            fmt.Println(s)
            return s
        case "windows":
            cmd := exec.Command("cls")
            cmd.Stdout = os.Stdout
            cmd.Run()
            fmt.Println(s)
            return s
        default:
		    fmt.Println("Unknown OS, cannot clear screen")
            return "unknown OS"

    }
    
}