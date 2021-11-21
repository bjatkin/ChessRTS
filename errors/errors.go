package errors

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"text/tabwriter"
)

// CodeLocation is a spot in the code
type CodeLocation struct {
	file string
	line int
}

func (c CodeLocation) String() string {
	return fmt.Sprintf("%s:%d", c.file, c.line)
}

// Err is the custome error type
type Err struct {
	child    error
	location CodeLocation

	Message string
	Code    int
	Tags    []Tag
}

func (tag Tag) String() string {
	// use tabs to set up for use with a tab writer
	return fmt.Sprintf("%s\t%v", tag.key, tag.val)
}

// Tag is an error tag
type Tag struct {
	key string
	val interface{}
}

// Wrap wraps one error into another
func Wrap(child error, code int, message string) *Err {
	parent := New(code, message)
	parent.child = child
	// we need to reset the code location since we call new from inside the werror package
	_, file, line, ok := runtime.Caller(1)
	parent.location = CodeLocation{}
	if ok {
		parent.location = CodeLocation{file: file, line: line}
	}

	return parent
}

// New creates a new custom error type
func New(code int, message string) *Err {
	err := &Err{
		Message: message,
		Code:    code,
	}

	_, file, line, ok := runtime.Caller(1)
	if ok {
		err.location = CodeLocation{file: file, line: line}
	}

	return err
}

func (err *Err) Tag(key string, val interface{}) *Err {
	err.Tags = append(err.Tags, Tag{key: key, val: val})
	return err
}

func (err *Err) Error() string {
	tags := &strings.Builder{}
	if len(err.Tags) > 0 {
		w := tabwriter.NewWriter(tags, 0, 1, 1, ' ', 0)

		fmt.Fprintln(w, "TAG_NAME\tTAG_VALUE")
		fmt.Fprintln(w, "--------\t---------")
		for _, tag := range err.Tags {
			fmt.Fprintln(w, tag)
		}

		flushErr := w.Flush()
		if flushErr != nil {
			log.Fatalf("writing error tags failed\n%s\n", flushErr)
		}
	}

	if err.child == nil {
		return fmt.Sprintf("ROOT_ERR: %s\nCode: %d\nLocation: %s\n%s",
			err.Message,
			err.Code,
			err.location,
			tags.String(),
		)
	}

	return fmt.Sprintf("ERR: %s\nCode: %d\nLocation: %s\n%s\n%s",
		err.Message,
		err.Code,
		err.location,
		tags.String(),
		err.child,
	)
}
