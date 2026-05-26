package cpu

func readFile(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func parseUint(s string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// 1. Handle negative values greater than MinInt64 (and)
// 2. Handle negative values lesser than MinInt64

// ParseUintList parses and validates the specified string as the value
// found in some cgroup file (e.g. cpuset.cpus, cpuset.mems), which could be
// one of the formats below. Note that duplicates are actually allowed in the
// input string. It returns a map[int]bool with available elements from val
// set to true.
// Supported formats:
// 7
// 1-6
// 0,3-4,7,8-10
// 0-0,0,1-7
// 03,1-3 <- this is gonna get parsed as [1,2,3]
// 3,2,1
// 0-2,3,1
func ParseUintList(val string) (map[int]bool, error) { _ = "STUB: not implemented"; return nil, nil }

// ReadLines reads contents from a file and splits them by new lines.
// A convenience wrapper to ReadLinesOffsetN(filename, 0, -1).
func readLines(filename string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// ReadLinesOffsetN reads contents from file and splits them by new line.
// The offset tells at which line number to start.
// The count determines the number of lines to read (starting from offset):
//
//	n >= 0: at most n lines
//	n < 0: whole file
func readLinesOffsetN(filename string, offset uint, n int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint
