package main

import (
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.Create("/tmp/access.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for {
		if _, err := f.Write([]byte("127.0.0.1 - james [09/May/2018:16:00:39 +0000] \"GET /report/t HTTP/1.0\" 200 123\n")); err != nil {
			log.Fatal(err)
		}

		time.Sleep(10 * time.Millisecond)

		if _, err := f.Write([]byte("127.0.0.1 - jill [09/May/2018:16:00:41 +0000] \"GET /api/user HTTP/1.0\" 200 234\n")); err != nil {
			log.Fatal(err)
		}

		time.Sleep(10 * time.Millisecond)
		if _, err := f.Write([]byte("127.0.0.1 - frank [09/May/2018:16:00:42 +0000] \"POST /api/user HTTP/1.0\" 200 34\n")); err != nil {
			log.Fatal(err)
		}
		time.Sleep(10 * time.Millisecond)
	}
}
