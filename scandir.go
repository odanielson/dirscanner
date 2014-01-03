
package dirscanner

import (
    "fmt"
    "io/ioutil"
    "path"
)

func ScanDir(a_path string, a_fileMsgs chan FileMsg,
    a_events chan Result) {

    if files, err := ioutil.ReadDir(a_path); err == nil {
        for _, file := range files {
            if file.IsDir() {
                ScanDir(path.Join(a_path, file.Name()), a_fileMsgs, a_events)
            } else {
                a_fileMsgs <- FileMsg{a_path, file.Name(), CREATED}
                a_events <- Result{DEBUG, fmt.Sprintf("Found file %s",
                    path.Join(a_path, file.Name())), nil}
            }
        }
    } else {
        a_events <- Result{ERROR,
            fmt.Sprintf("Failed to scan path %s", a_path), err}
    }
}
