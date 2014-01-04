
package main

import (
    "fmt"
    "path"

    "github.com/odanielson/dirscanner"
)

func main() {

    // Create a channel for file events
    channel := make(chan dirscanner.FileMsg, 1000)
    // Create another channel for log and error events
    events := make(chan dirscanner.Event, 1000)

    // Create and start the scanner
    scanner := dirscanner.CreateDirScanner("./", channel, events)
    scanner.Start()

    // Watch the channels for FileMsgs and other Events
    var fileMsg dirscanner.FileMsg
    var event dirscanner.Event

    for {
        select {
        case fileMsg = <- channel:
            fmt.Printf("FileMsg[%s] %s\n", fileMsg.M_operation,
                path.Join(fileMsg.M_path, fileMsg.M_name))
        case event = <- events:
            fmt.Printf("Event[%s] %s\n", event.M_level, event.M_msg)
            if event.M_err != nil {
                fmt.Println(event.M_err)
            }
        }
    }
}
