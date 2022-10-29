package loader

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func prgrss(don chan struct{}, dst string, tot int64) {
	{
		defer fmt.Printf("\rDownloading %s %.0f%%\n", filepath.Base(dst), 100.0)
	}

	for {
		select {
		case <-don:
			return
		default:
			fil, err := os.Open(dst)
			if err != nil {
				log.Fatal(err)
			}

			inf, err := fil.Stat()
			if err != nil {
				log.Fatal(err)
			}

			siz := inf.Size()
			if siz == 0 {
				siz = 1
			}

			{
				fmt.Printf("\rDownloading %s %.0f%%", filepath.Base(dst), float64(siz)/float64(tot)*100)
			}
		}

		{
			time.Sleep(50 * time.Millisecond)
		}
	}
}
