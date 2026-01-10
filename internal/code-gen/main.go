// This is part of an internal code generation tool of Gost-DOM. It's in an
// internal package and not used by production code.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gost-dom/code-gen/events"
	htmlelements "github.com/gost-dom/code-gen/html-elements"
	"github.com/gost-dom/code-gen/interfaces"
	"github.com/gost-dom/code-gen/scripting"
)

func getWriter(output string) io.Writer {
	if output == "stdout" {
		return os.Stdout
	}
	file, err := os.Create(output)
	if err != nil {
		fmt.Println("Error creating output file")
		os.Exit(1)
	}
	return file
}

var generators = map[string]func(io.Writer) error{
	"html-elements": generateHtmlElements,
}

func main() {
	var (
		generatorType string
		packageName   string
		outputFile    string
		debug         bool
	)
	flag.BoolVar(&debug, "d", false, "Debug")
	flag.StringVar(&outputFile, "o", "", "Output file to write")
	flag.StringVar(&generatorType, "g", "", "Generator type")
	flag.StringVar(&packageName, "p", "", "Package to generate")
	flag.Parse()

	globals := []string{"Window" /*, "DedicatedWorkerGlobalScope" */}

	switch generatorType {
	case "script":
		exitOnError(scripting.CreateJavaScriptMappings(packageName))
		os.Exit(0)
	case "script-bootstrap":
		exitOnError(scripting.GenerateRegisterFunctions(packageName, globals))
		os.Exit(0)
	case "event-init-decoders":
		if packageName == "" {
			panic("Missing package spec")
		}
		exitOnError(scripting.CreateEventInitDecoders(packageName))
		os.Exit(0)
	case "gotypes":
		if packageName == "" {
			panic("Missing package spec")
		}
		exitOnError(htmlelements.CreateImplementationPackage(packageName))
		os.Exit(0)
	case "interfaces":
		if packageName == "" {
			panic("Missing package spec")
		}
		exitOnError(interfaces.CreateInterfaces(packageName))
		os.Exit(0)
	case "eventTypes":
		if packageName == "" {
			panic("Missing package spec")
		}
		exitOnError(events.CreateEventGenerators(packageName))
		os.Exit(0)
	case "eventInitTypes":
		if packageName == "" {
			panic("Missing package spec")
		}
		exitOnError(events.CreateEventDicts(packageName))
		os.Exit(0)
	}

	if outputFile == "" || generatorType == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if debug {
		fmt.Println(strings.Join(os.Args, " "))
		fmt.Println("--------")
	}

	file := getWriter(outputFile)

	generator, ok := generators[generatorType]
	if !ok {
		os.Exit(1)
	}
	err := generator(file)
	exitOnError(err)
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println("Error running generator")
		fmt.Println(err)
		os.Exit(1)
	}
}
