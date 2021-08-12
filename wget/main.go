package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Wget download the file using the linux wget command
//
// It is necessary to inform the url, filename and extension without the '.'
//
// The downloads are saved in the tmp directory created in the project's root.
func Wget(url, filepath, extension string) error {

	erro := os.Mkdir("tmp", 0755)
	if erro != nil {
		log.Fatal(erro)
	}

	out, err := os.Create(fmt.Sprintf("./tmp/%s.%s", filepath, extension))
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
