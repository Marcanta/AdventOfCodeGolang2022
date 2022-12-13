package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Dir struct {
	name      string
	files     []File
	parentDir *Dir
	dirs      []*Dir
	size      int
}

type File struct {
	name string
	size int
}

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func (d Dir) search(name string) *Dir {
	for _, dir := range d.dirs {
		if dir.name == name {
			return dir
		}
	}
	return nil
}

func (d *Dir) getTotalSizeOfLilDirs() *Dir {
	// if 100000 d.getSize
	if size := d.parentDir.getSize(); size >= 100000 {
		// fmt.Printf("d: %s, size: %d, files: %v, dirs: %v\n", d.name, d.size, d.files, d.getSubDirsName())
		return d
	}
	// some code goes here
	return d.parentDir.getTotalSizeOfLilDirs()
}

func (d *Dir) getSize() int {
	size := 0
	for _, file := range d.files {
		size += file.size
	}
	for _, dir := range d.dirs {
		size += dir.size
	}
	d.size = size
	return size
}

func (d *Dir) getLastDir() []*Dir {
	if len(d.dirs) == 0 {
		return []*Dir{d}
	}
	resDirs := []*Dir{}
	for _, dir := range d.dirs {
		resDirs = append(resDirs, dir.getLastDir()...)
	}
	return resDirs
}

func retrieveFileTree(rawTermOuput string) (fileTree Dir) {
	termOutput := strings.Split(rawTermOuput, "\n")

	fileTree = Dir{
		name:  "/",
		files: []File{},
		dirs:  []*Dir{},
	}
	currentDir := &fileTree
	for i, line := range termOutput {
		if i == 0 {
			continue
		}
		if strings.Contains(line, "$ cd") {
			name := strings.Split(line, " ")[len(strings.Split(line, " "))-1]
			if name == ".." {
				currentDir = currentDir.parentDir
				continue
			}
			if dir := currentDir.search(name); dir != nil {
				currentDir = dir
			}
		} else if strings.Contains(line, "$ ls") {
			continue
		} else {
			entity := strings.Split(line, " ")
			if entity[0] == "dir" {
				currentDir.dirs = append(currentDir.dirs, &Dir{name: entity[1], parentDir: currentDir})
			} else {
				size, _ := strconv.Atoi(entity[0])
				currentDir.files = append(currentDir.files, File{name: entity[1], size: size})
			}
		}
	}
	return
}

func unique(dirSlice []*Dir) []*Dir {
	keys := make(map[*Dir]bool)
	list := []*Dir{}
	for _, entry := range dirSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	data := readFile("../../inputs/07.in")

	fileTree := retrieveFileTree(data)

	dirs := []*Dir{}
	for _, dir := range fileTree.getLastDir() {
		dir.getSize()
		dirs = append(dirs, dir.getTotalSizeOfLilDirs())
		// fmt.Printf("name: %s, parentDir: { name: %s size: %d }\n", dir.name, dir.parentDir.name, dir.parentDir.size)
	}
	sum := 0
	dirs = unique(dirs)
	for _, dir := range dirs {
		sum += dir.size
	}
	fmt.Printf("sum: %d\n", sum)

	// test2 := Dir{name: "test2", size: 0, files: []File{{name: "sfsdfsdf", size: 134}, {name: "sfsdfsdf", size: 100}}}
	// test := Dir{name: "test", size: 0, dirs: []*Dir{&test2}}

	// fmt.Printf("test2.getSize(): %v\n", test2.getSize())
	// fmt.Printf("test.getSize(): %v\n", test.getSize())
	fmt.Printf("fileTree: name: %s, dirs: %v, files: %v, parentDir: %v\n", fileTree.name, fileTree.dirs, fileTree.files, fileTree.parentDir)
}
