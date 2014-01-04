
package dirscanner

type DirScanner struct {
    m_path string
    m_fileMsgs chan FileMsg
    m_events chan Event
    m_monitor *Monitor
}

func CreateDirScanner(a_path string, a_fileMsgs chan FileMsg,
    a_events chan Event) *DirScanner {

    d := DirScanner{}
    d.m_path = a_path
    d.m_fileMsgs = a_fileMsgs
    d.m_events = a_events
    d.m_monitor = CreateMonitor(a_path, a_fileMsgs, a_events)

    return &d
}

func (d *DirScanner) Start() {

    // First,  start monitoring the filesystem for changes
    d.m_monitor.Start()

    // Second, make an initial scan of the file system
    go func() {
        ScanDir(d.m_path, d.m_fileMsgs, d.m_events)
        d.m_events <- Event{INFO, "Initial filescan completed.", nil}
    } ()

}

func (d *DirScanner) Stop() {
    d.m_monitor.Stop()
}
