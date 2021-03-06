
package dirscanner

import (
    "fmt"
    "io/ioutil"
    "path"
)

func ScanDir(a_path string, a_fileMsgs chan FileMsg,
    a_events chan Event) {

    if files, err := ioutil.ReadDir(a_path); err == nil {
        for _, file := range files {
            if file.IsDir() {
                a_fileMsgs <- FileMsg{a_path, file.Name(), CREATED}
                a_events <- Event{DEBUG, fmt.Sprintf("Found dir %s",
                    path.Join(a_path, file.Name())), nil}
                ScanDir(path.Join(a_path, file.Name()), a_fileMsgs, a_events)
            } else {
                a_fileMsgs <- FileMsg{a_path, file.Name(), CREATED}
                a_events <- Event{DEBUG, fmt.Sprintf("Found file %s",
                    path.Join(a_path, file.Name())), nil}
            }
        }
    } else {
        a_events <- Event{ERROR,
            fmt.Sprintf("Failed to scan path %s", a_path), err}
    }
}
