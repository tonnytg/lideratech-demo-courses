package main

import (
    "log"
    "os"
    "bufio"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {

        worlds := string(scanner.Bytes())
        if worlds == "exit" {
            os.Exit(0)
        }

        log.Println(worlds)

    }
}
