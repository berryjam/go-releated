package main

import (
	"os"
	"fmt"
)

/**
In Go, most of the file operation functions are located in the os package. Here are some directory functions:

1.func Mkdir(name string, perm FileMode) error
Create a directory with name. perm is the directory permissions, i.e 0777.

2.func MkdirAll(path string, perm FileMode) error
Create multiple directories according to path, like astaxie/test1/test2.

3.func Remove(name string) error
Removes directory with name. Returns error if it's not a directory or not empty.

4.func RemoveAll(path string) error
Removes multiple directories according to path. Directories will not be deleted if path is a single path.
 */
func Directories() {
	os.Mkdir("astaxie", 0777)
	os.MkdirAll("astaxie/test1/test2", 0777)
	err := os.Remove("astaxie")
	if err != nil {
		fmt.Println(err)
	}
}

/**
There are two functions for creating files:

1.func Create(name string) (file *File, err Error)
Create a file with name and return a read-writable file object with permission 0666.

2.func NewFile(fd uintptr, name string) *File
Create a file and return a file object.

There are also two functions to open files:

1.func Open(name string) (file *File, err Error)
Opens a file called name with read-only access, calling OpenFile under the covers.

2.func OpenFile(name string, flag int, perm uint32) (file *File, err Error)
Opens a file called name. flag is open mode like read-only, read-write, etc. perm are the file permissions.
 */
func Files() {
	/**
	Create and open files
	 */
	userFile := "astaxie.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test!\r\n")
		fout.Write([]byte("Just a test!\r\n"))
	}

	/**
	Read files

	Functions for reading files:

	1/func (file *File) Read(b []byte) (n int, err Error)
	Read data to b.

	2.func (file *File) ReadAt(b []byte, off int64) (n int, err Error)
	Read data from position off to b.
	 */
	f1, err := os.Open("astaxie.txt")
	if err != nil {
		fmt.Println(userFile, err)
		return
	}

	defer f1.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := f1.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}

	/**
	Delete files

	Go uses the same function for removing files and directories:

	1.func Remove(name string) Error
	Remove a file or directory called name.( a name ending with / signifies that it's a directory )
	 */
	removeErr := os.Remove("astaxie.txt")
	if removeErr != nil {
		fmt.Printf("error %v", removeErr)
	}
}

func main() {
	//Directories()
	Files()
}
