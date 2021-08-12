package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

//Exits - Check if the directory exists, if it doesn't it will be created
func Exits(path string) error {

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {

			erro := os.Mkdir(path, 0755)
			if erro != nil {
				fmt.Println(erro)
				return erro
			}
		} else {
			return err
		}
	}

	return nil
}

// Wget download the file using the linux wget command
//
// It is necessary to inform the url, filename and extension without the '.'
//
// The downloads are saved in the tmp directory created in the project's root.
func Wget(url, filepath, extension string) error {

	path := "tmp"

	erro := Exits(path)

	if erro != nil {
		log.Fatal(erro)
	}

	out, err := os.Create(fmt.Sprintf("./%s/%s.%s", path, filepath, extension))
	if err != nil {
		return err
	}
	defer out.Close()

	// run shell `wget URL -O filepath`
	cmd := exec.Command("wget", url, "-O", out.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {

	url := flag.String("url", "", "URL to download")

	filename := flag.String("file", "", "Inform the file name")

	extension := flag.String("ext", "", "Inform the file extension without the '.'")

	flag.Parse()

	if *url == "" || *filename == "" || *extension == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println("Download Started")

	erro := Wget(*url, *filename, *extension)

	if erro != nil {
		panic(erro)
	}

	fmt.Println("Download Finished")

}
