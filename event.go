
package dirscanner

const (
    DEBUG = "DEBUG"
    INFO = "INFO"
    WARNING = "WARNING"
    ERROR = "ERROR"
    PANIC = "PANIC"
)

type Event struct {
    M_level string
    M_msg string
    M_err error
}
