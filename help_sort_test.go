package flaggy

import (
    "os"
    "strings"
    "testing"
)

func TestHelpFlagsSortedWhenEnabled(t *testing.T) {
    // Use default parser and functions to enable sort
    ResetParser()
    DefaultParser.ShowHelpWithHFlag = false
    DefaultParser.ShowVersionWithVersionFlag = false
    SortFlagsByLongName()

    var a, b, z string
    // Intentionally add in non-sorted order
    String(&z, "z", "zeta", "")
    String(&a, "a", "alpha", "")
    String(&b, "b", "beta", "")

    rd, wr, err := os.Pipe()
    if err != nil { t.Fatalf("pipe error: %v", err) }
    saved := os.Stderr
    os.Stderr = wr
    defer func(){ os.Stderr = saved }()

    DefaultParser.ShowHelp()

    buf := make([]byte, 4096)
    n, err := rd.Read(buf)
    if err != nil { t.Fatalf("read error: %v", err) }
    lines := strings.Split(string(buf[:n]), "\n")

    // collect just the flag lines (start with two spaces then a dash or spaces then --)
    var flagLines []string
    inFlags := false
    for _, l := range lines {
        if strings.HasPrefix(l, "  Flags:") { inFlags = true; continue }
        if inFlags {
            if strings.TrimSpace(l) == "" { break }
            flagLines = append(flagLines, l)
        }
    }
    if len(flagLines) < 3 {
        t.Fatalf("expected at least 3 flag lines, got %d: %q", len(flagLines), flagLines)
    }

    // find the three of interest
    var idxAlpha, idxBeta, idxZeta = -1, -1, -1
    for i, l := range flagLines {
        if strings.Contains(l, "--alpha") { idxAlpha = i }
        if strings.Contains(l, "--beta") { idxBeta = i }
        if strings.Contains(l, "--zeta") { idxZeta = i }
    }
    if idxAlpha == -1 || idxBeta == -1 || idxZeta == -1 {
        t.Fatalf("expected to find alpha, beta, zeta in flags; got: %q", flagLines)
    }
    if !(idxAlpha < idxBeta && idxBeta < idxZeta) {
        t.Fatalf("flags not sorted: alpha=%d beta=%d zeta=%d; lines=%q", idxAlpha, idxBeta, idxZeta, flagLines)
    }
}

func TestHelpFlagsSortedReversed(t *testing.T) {
    // Use default parser and reversed sort
    ResetParser()
    DefaultParser.ShowHelpWithHFlag = false
    DefaultParser.ShowVersionWithVersionFlag = false
    SortFlagsByLongNameReversed()

    var a, b, z string
    // Intentionally add in non-sorted order
    String(&z, "z", "zeta", "")
    String(&a, "a", "alpha", "")
    String(&b, "b", "beta", "")

    rd, wr, err := os.Pipe()
    if err != nil { t.Fatalf("pipe error: %v", err) }
    saved := os.Stderr
    os.Stderr = wr
    defer func(){ os.Stderr = saved }()

    DefaultParser.ShowHelp()

    buf := make([]byte, 4096)
    n, err := rd.Read(buf)
    if err != nil { t.Fatalf("read error: %v", err) }
    lines := strings.Split(string(buf[:n]), "\n")

    var flagLines []string
    inFlags := false
    for _, l := range lines {
        if strings.HasPrefix(l, "  Flags:") { inFlags = true; continue }
        if inFlags {
            if strings.TrimSpace(l) == "" { break }
            flagLines = append(flagLines, l)
        }
    }
    if len(flagLines) < 3 {
        t.Fatalf("expected at least 3 flag lines, got %d: %q", len(flagLines), flagLines)
    }

    var idxAlpha, idxBeta, idxZeta = -1, -1, -1
    for i, l := range flagLines {
        if strings.Contains(l, "--alpha") { idxAlpha = i }
        if strings.Contains(l, "--beta") { idxBeta = i }
        if strings.Contains(l, "--zeta") { idxZeta = i }
    }
    if idxAlpha == -1 || idxBeta == -1 || idxZeta == -1 {
        t.Fatalf("expected to find alpha, beta, zeta in flags; got: %q", flagLines)
    }
    if !(idxZeta < idxBeta && idxBeta < idxAlpha) {
        t.Fatalf("flags not reverse-sorted: alpha=%d beta=%d zeta=%d; lines=%q", idxAlpha, idxBeta, idxZeta, flagLines)
    }
}
