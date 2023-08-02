package main

import (
	"github.com/gostaticanalysis/templates/inspect"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(inspect.Analyzer) }
