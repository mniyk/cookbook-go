package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	inDir := "./"
	outDir := "./output"

	if f, err := os.Stat(outDir); os.IsNotExist(err) || !f.IsDir() {
		err := os.Mkdir(outDir, 0777)
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}
	}

	_ = filepath.Walk(inDir, func(path string, info os.FileInfo, err error) error {
		ex := filepath.Ext(path)

		if path != inDir && ex == ".zip" {
			zr, err := zip.OpenReader(path)
			if err != nil {
				fmt.Printf("Error: %s", err)
				os.Exit(1)
			}

			for _, f := range zr.File {
				r, err := f.Open()
				if err != nil {
					fmt.Printf("Error: %s", err)
					os.Exit(1)
				}

				outPath := filepath.Join(outDir, f.Name)

				if f.FileInfo().IsDir() {
					os.MkdirAll(outPath, f.Mode())
					continue
				}

				w, err := os.OpenFile(outPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
				if err != nil {
					fmt.Printf("Error: %s", err)
					os.Exit(1)
				}

				_, err = io.Copy(w, r)
				if err != nil {
					fmt.Printf("Error: %s", err)
					os.Exit(1)
				}

				w.Close()
				r.Close()
			}
			zr.Close()
		}
		return nil
	})
}
