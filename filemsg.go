
package dirscanner

const (
    CREATED = "CREATED"
    MODIFIED = "MODIFIED"
    DELETED = "DELETED"
    RENAMED = "RENAMED"
)

type FileMsg struct {
    M_path string
    M_name string
    M_operation string
}
