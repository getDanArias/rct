package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	argsWithoutProg := os.Args[1:]

	action := argsWithoutProg[0]
	actionBody := argsWithoutProg[1:]

	var arguments []string
	var options []string

	for i := 0; i < len(actionBody); i++ {
		if strings.Contains(actionBody[i], "--") {
			options = append(options, actionBody[i])
		} else {
			arguments = append(arguments, actionBody[i])
		}
	}

	wdPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if action == "g" || action == "generate" {
		construct := arguments[0]
		name := arguments[1]

		if construct == "c" || construct == "component" {
			componentDir := wdPath + "/" + name

			if _, err := os.Stat(componentDir);

				os.IsNotExist(err) {
				os.Mkdir(componentDir, os.ModePerm)
			}

			stateful := true
			stylish := true

			jsFilePath := componentDir + "/" + name + ".js"

			jsFile, err := os.Create(jsFilePath)
			if err != nil {
				panic(err)
			}

			importReactString := `import React from 'react';`
			importStyleString := "import './" + name + ".css';"

			for i:= 0; i < len(options); i++ {
				if strings.Contains(options[i], "--stateless") {
					stateful = false
				} else if strings.Contains(options[i], "--no-style") {
					stylish = false
				}
			}

			if stylish {
				styleFilePath := componentDir + "/" + name + ".scss"
				styleCompiledFilePath := componentDir + "/" + name + ".css"

				styleFile, err := os.Create(styleFilePath)
				if err != nil {
					panic(err)
				}

				styleCompiledFile, err := os.Create(styleCompiledFilePath)
				if err != nil {
					panic(err)
				}

				className := "." + name + " {\n"
				defaultProp1 := "  display: flex;\n"
				defaultProp2 := "  flex-direction: column;\n"
				defaultProp3 := "  justify-content: center;\n"
				defaultProp4 := "  align-items: center;\n"
				classEnd := "}\n"

				finalStyleString := className + defaultProp1 + defaultProp2 + defaultProp3 + defaultProp4 + classEnd

				_, err = styleFile.WriteString(finalStyleString)
				if err != nil {
					panic(err)
				}

				_, err = styleCompiledFile.WriteString(finalStyleString)
				if err != nil {
					panic(err)
				}

				defer styleFile.Close()

			}

			if stateful {
				classExport := "export default class " + name + " extends React.Component {"
				render := "  render() {"
				renderReturn := "    return <div className={`" + name + "`}>" + name +"</div>;"
				renderClose := "  }"
				classClose := "}"

				stateCompString := importReactString + "\n\n"

				if stylish {
					stateCompString += importStyleString + "\n\n"
				}

				stateCompString +=  classExport + "\n"
				stateCompString += render + "\n"
				stateCompString += renderReturn + "\n"
				stateCompString += renderClose + "\n"
				stateCompString += classClose + "\n"

				_, err = jsFile.WriteString(stateCompString)

				if err != nil {
					panic(err)
				}

			} else {
				functionHeader := "const " + name + "= props => {"
				functionReturn := "return <div className={`" + name + "`}>" + name + "</div>"
				functionClose := "};"
				functionExport := "export default " + name + ";"

				statelessCompString := importReactString + "\n\n"

				if stylish {
					statelessCompString += importStyleString + "\n\n"
				}

				statelessCompString += functionHeader + "\n"
				statelessCompString += functionReturn + "\n"
				statelessCompString += functionClose + "\n"
				statelessCompString += functionExport + "\n"

				_, err = jsFile.WriteString(statelessCompString)
				if err != nil {
					panic(err)
				}
			}

			defer jsFile.Close()
		}
	} else if action == "h" || action == "help" {
		fmt.Println(`  rct generate <construct> <options...> <name>`)
		fmt.Println(`  aliases: g`)
		fmt.Println(`  component: Create a directory with <name> to hold a React component`)
		fmt.Println(`    aliases: c`)
		fmt.Println(`    --state: Create React stateful component with <name>`)
		fmt.Println(`    --stateless: Create React stateless component with <name>`)
		fmt.Println(`    --style: Create .scss file with <name> of component`)
		fmt.Println(`    --no-style: Skips creating '.scss file for component`)

	}
}
