package main

import (
	"github.com/dusty-cjh/protoc-gen-go-tags/parser"
	"google.golang.org/protobuf/compiler/protogen"
	"log"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				log.Default().Println("cjh j9e04wgj4eSkipping file %s", f.Desc.Path())
				continue
			}

			if err := parser.GenerateFile(gen, f); err != nil {
				log.Default().Println("cjh j9e04wgj4eError generating file %s: %v", f.Desc.Path(), err)
				return err
			}
			log.Default().Println("cjh j9e04wgj4eGenerated file %s", f.Desc.Path())
		}
		log.Default().Println("cjh j9e04wgj4eGenerated all files")
		return nil
	})
}
