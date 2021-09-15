package main

import "fmt"

type Hierarchy interface {
	Name() string
	Add(child Hierarchy) Hierarchy
	Remove(child Hierarchy) Hierarchy
	Children() []*Hierarchy
}

type Directory struct {
	children []*Hierarchy
	name     string
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) Add(child Hierarchy) Hierarchy {
	d.children = append(d.children, &child)
	return d
}

func (d *Directory) Remove(child Hierarchy) Hierarchy {
	// find and remove an item by name
	for i, hierarchy := range d.children {
		if (*hierarchy).Name() == child.Name() {
			d.children[i] = d.children[len(d.children)-1]
			d.children = d.children[:len(d.children)-1]
			return d
		}
	}
	return d
}

func (d *Directory) Children() []*Hierarchy {
	return d.children
}

func NewDirectory(name string) Hierarchy {
	return &Directory{
		children: []*Hierarchy{},
		name:     name,
	}
}

// file is in the hierarchy structure but does not contain children
// we can create another interface for file structure as a leaf
type File struct {
	name string
}

func (d *File) Name() string {
	return d.name
}

func (d *File) Add(_ Hierarchy) Hierarchy {
	return d
}

func (d *File) Remove(_ Hierarchy) Hierarchy {
	return d
}

func (d *File) Children() []*Hierarchy {
	return nil
}

func NewFile(name string) Hierarchy {
	return &File{name: name}
}

func PrintDirectory(dir Hierarchy, indent int) {
	if dir == nil {
		return
	}
	for i := 0; i < indent; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%s\n", dir.Name())
	for _, child := range dir.Children() {
		PrintDirectory(*child, indent+2)
	}
}

func main() {
	documentsDir := NewDirectory("Documents")

	imagesDir := NewDirectory("Images")
	bookDir := NewDirectory("Books")
	downloadDir := NewDirectory("Downloads")
	// document directory contains 3 other directories
	documentsDir.Add(imagesDir).Add(downloadDir).Add(bookDir)
	// images directory contain another directory
	myImages := NewDirectory("MyImages")
	myImages.Add(NewFile("my-image.png"))
	imagesDir.Add(myImages).Add(NewFile("background.jpg")).Add(NewFile("nero.jpg"))

	bookDir.Add(NewFile("book1.pdf")).Add(NewFile("book2.pdf")).Add(NewFile("book3.pdf"))
	downloadDir.Add(NewFile("Video.mp4"))

	PrintDirectory(documentsDir, 2)
	fmt.Println("----------------------------")
	// you can now move hole image directory to another folder like download folder
	documentsDir.Remove(imagesDir)
	downloadDir.Add(imagesDir)
	PrintDirectory(documentsDir, 2)
}
