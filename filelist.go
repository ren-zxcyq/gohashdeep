// this file contains a far away cousin of ls
// call as:	walkdir()
package main

import (
	"fmt"
	"os"
	"log"
	"path/filepath"
	"syscall"
	"text/tabwriter"
	"crypto/sha256"
	"io"
	"encoding/hex"
)

func main() {
	walkdir()
}

// Function walkdir lists contents of the current directory, as shown below
// The Hash column contains the sha256 of the file.
//
// Path              UID    GID    Size   Hash
// ----              ---    ---    ----   ----
// .                 1000   1000   4096   
// filelist.go       1000   1000   2128   60a2d712cb16e92ceaeb1c494150b28e32567bcba99238e17c33f409ad97858f
func walkdir() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ',0)
		//tabwriter.AlignRight)
	fmt.Fprintf(w, "Path\tUID\tGID\tSize\tHash\n")
	fmt.Fprintf(w, "----\t---\t---\t----\t----\n")
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			//	fmt.Println(path, info.Size())
			info, _ = os.Stat(path)

			var UID int
			var GID int
			if stat, ok := info.Sys().(*syscall.Stat_t); ok {
				UID = int(stat.Uid)
				GID = int(stat.Gid)
			} else {
				// we are not in linux, this won't work anyway in windows, 
				// but maybe you want to log warnings
				UID = os.Getuid()
				GID = os.Getgid()
			}

			// test hashfilecontents
			var hash string
			if fileisfile(path) {
			//if (path != "." && path != "..") {
				cwd,_ := os.Getwd()
				hash = hashfilecontents(cwd+"/"+path)
			}

			//fmt.Printf("%s\t%d\t%d\t%d\r\n", path, UID, GID, info.Size())
			//var width int = 20

			// Substitute this for tabwriter:
			//fmt.Printf("%-20s %-20d %-20d %-20d\n", path, UID, GID, info.Size())

			fmt.Fprintf(w, "%s\t%d\t%d\t%d\t%s\n",path,UID,GID,info.Size(),hash)

			//fix space via https://stackoverflow.com/a/59364558
			return nil
		})
	w.Flush()
	if err != nil {
		log.Println(err)
	}
}

func fileisfile(filepath string) bool {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		log.Fatal("File does not exist.", filepath)
	}
	if info.IsDir() {
		return false
	} else {
		return true
	}
}

func hashfilecontents(filepath string) string {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%x", h.Sum(nil))
	//fmt.Println("Hash is:", hex.EncodeToString(h.Sum(nil)))
	return hex.EncodeToString(h.Sum(nil))
}
