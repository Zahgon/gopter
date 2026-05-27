package gopter

import (
	"io"
)

const newLine = "\n"

// FormatedReporter reports test results in a human readable manager.
type FormatedReporter struct {
	verbose bool
	width   int
	output  io.Writer
}

// NewFormatedReporter create a new formated reporter
// verbose toggles verbose output of the property results
// width is the maximal width per line
// output is the writer were the report will be written to
func NewFormatedReporter(verbose bool, width int, output io.Writer) Reporter {
	_ = "STUB: not implemented"
	return *new(Reporter)
}

// ConsoleReporter creates a FormatedReporter writing to the console (i.e. stdout)
func ConsoleReporter(verbose bool) Reporter { _ = "STUB: not implemented"; return *new(Reporter) }

// ReportTestResult reports a single property result
func (r *FormatedReporter) ReportTestResult(propName string, result *TestResult) {
	_ = "STUB: not implemented"
	return
}

func (r *FormatedReporter) reportResult(result *TestResult) string {
	_ = "STUB: not implemented"
	return ""
}

func (r *FormatedReporter) reportLabels(labels []string) string {
	_ = "STUB: not implemented"
	return ""
}

func (r *FormatedReporter) reportPropArgs(p PropArgs) string { _ = "STUB: not implemented"; return "" }

func (r *FormatedReporter) reportPropArg(idx int, propArg *PropArg) string {
	_ = "STUB: not implemented"
	return ""
}

func (r *FormatedReporter) formatLines(str, lead, trail string) string {
	_ = "STUB: not implemented"
	return ""
}

func (r *FormatedReporter) breakLine(str, lead string) string { _ = "STUB: not implemented"; return "" }

func concatLines(strs ...string) string { _ = "STUB: not implemented"; return "" }
