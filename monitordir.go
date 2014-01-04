
package dirscanner

import (
    "fmt"
    "io/ioutil"
    "os"
    "path"

    "github.com/howeyc/fsnotify"
)

type Monitor struct {
    m_path string
    m_fileMsgs chan FileMsg
    m_events chan Event
    m_watcher *fsnotify.Watcher
}

func (m *Monitor) handleEvents() {

    for {
        select {
        case ev := <-m.m_watcher.Event:
            if ev != nil {
                op := ""
                switch {
                case ev.IsCreate():
                    if file, err := os.Lstat(ev.Name); err == nil {
                        if file.IsDir() {
                            m.monitorDir(ev.Name)
                        }
                    }
                    op = CREATED
                case ev.IsDelete():
                    op = DELETED
                case ev.IsModify():
                    op = MODIFIED
                case ev.IsRename():
                    op = RENAMED
                }
                if op != "" {
                    m.m_fileMsgs <- FileMsg{M_name: path.Base(ev.Name),
                        M_path: path.Dir(ev.Name),  M_operation: op}
                }

            } else {
                m.m_events <- Event{ERROR, "Got empty event from fsnotify.",
                    nil}
            }
        case err := <-m.m_watcher.Error:
            m.m_events <- Event{ERROR, "Got error from fsnotify.", err}
        }
    }
}

func (m *Monitor) monitorDir(a_path string) {

    if files, err := ioutil.ReadDir(a_path); err == nil {
        for _, file := range files {
            if file.IsDir() {
                m.monitorDir(fmt.Sprintf("%s/%s", a_path, file.Name()))
            }
        }
    } else {
        m.m_events <- Event{ERROR, fmt.Sprintf(
            "Failed to list files at path %s\n", a_path), err}
    }

    if err := m.m_watcher.Watch(a_path); err != nil {
        m.m_events <- Event{ERROR, fmt.Sprintf(
            "Failed to watch path %s\n", a_path), err}
        return
    }
}

func CreateMonitor(a_path string, a_fileMsgs chan FileMsg,
    a_events chan Event) *Monitor {

    m := Monitor{}
    m.m_path = a_path
    m.m_fileMsgs = a_fileMsgs
    m.m_events = a_events
    var err error
    if m.m_watcher, err = fsnotify.NewWatcher(); err != nil {
        m.m_events <- Event{ERROR, "Failed to create fsnotify watcher.",
            nil}
        return nil
    }

    return &m
}

func (m *Monitor) Start() {
    go m.handleEvents()
    go m.monitorDir(m.m_path)
}

func (m *Monitor) Stop() {
    m.m_watcher.Close()
}
