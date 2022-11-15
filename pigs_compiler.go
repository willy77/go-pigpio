package pigpio

import (
	"encoding/json"
	"fmt"
	"github.com/BxNiom/go-pigpio/assets"
	"strings"
)

type CompileError struct {
	line int
	msg  string
}

func NewCompileError(line int, msgFmt string, a ...any) *CompileError {
	return &CompileError{line: line, msg: fmt.Sprintf(msgFmt, a...)}
}

func (ce *CompileError) Error() string {
	return fmt.Sprintf("(%d): %s", ce.line, ce.msg)
}

type CompilerMacro struct {
	Name    string
	Params  int
	Results int
	Code    string
}

type Compiler struct {
	result []string
	macros []CompilerMacro
}

func NewCompiler() *Compiler {
	return &Compiler{
		macros: make([]CompilerMacro, 0),
	}
}

func (c *Compiler) append(line string) {
	c.result = append(c.result, line)
}

func (c *Compiler) LoadMacros(j string) error {
	var macros []CompilerMacro
	err := json.Unmarshal([]byte(j), &macros)
	if err != nil {
		return err
	}

	c.macros = append(c.macros, macros...)
	return nil
}

func (c *Compiler) LoadDefaultMacros() error {
	var e error
	if e = c.LoadMacros(assets.DefaultMacros); e != nil {
		return e
	}

	return nil
}

func (c *Compiler) Compile(src string) (string, error) {
	var i int
	var parts []string
	var usedMacros []int
	var lastMacro *CompilerMacro

	c.result = make([]string, 0)

	lastMacro = nil
	usedMacros = make([]int, 0)
	srcLines := strings.Split(src, "\n")

	for idx, line := range srcLines {
		line = strings.TrimSpace(line)
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		if ci := strings.Index(line, "#"); ci >= 0 {
			line = line[:ci]
			line = strings.TrimSpace(line)
		}

		parts = strings.Split(line, " ")
		cmd := strings.ToLower(parts[0])
		if cmd == "ldr" {
			if len(parts) <= 1 {
				return "", NewCompileError(idx, "ldr has to few parameters")
			}

			if lastMacro != nil && len(parts)-1 > lastMacro.Results {
				return "", NewCompileError(idx, "ldr has to many parameters")
			}

			for i = 1; i < len(parts); i++ {
				c.append(fmt.Sprintf("ld %s v%d", parts[i], i+139))
			}

			continue
		}

		macroFound := false
		for mi, macro := range c.macros {
			if macro.Name == cmd {
				if len(parts)-1 != macro.Params {
					return "", NewCompileError(idx, "macro has to few parameters")
				}

				for p := 1; p < len(parts); p++ {
					c.append(fmt.Sprintf("ld v%d %s", p+129, parts[p]))
				}
				c.append(fmt.Sprintf("call %d", 1100+mi))

				if !contains(usedMacros, mi) {
					usedMacros = append(usedMacros, mi)
				}
				macroFound = true
				break
			}
		}

		if !macroFound {
			c.append(line)
		}
	}

	for _, mid := range usedMacros {
		c.append(fmt.Sprintf("tag %d %s ret", 1100+mid, c.macros[mid].Code))
	}

	return strings.Join(c.result, " "), nil
}

func contains(list []int, val int) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == val {
			return true
		}
	}

	return false
}
