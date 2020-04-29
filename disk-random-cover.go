package main

import (
	"fmt"
	"github.com/cheggaaa/pb"
	"github.com/shawnwyckoff/gpkg/container/grandom"
	"github.com/shawnwyckoff/gpkg/container/gvolume"
	"github.com/shawnwyckoff/gpkg/sys/ghdd"
	"github.com/shawnwyckoff/gpkg/sys/gproc"
	"os"
)

func main() {

	mydir, err := gproc.SelfDir()
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Create("disk-random-cover.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	vi, err := ghdd.GetVolumeInfo(mydir)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("this program will use random string to erase the volume where itself is")
	fmt.Println(vi.Available.String(), "[", vi.Available.MBytes(), " MB", "]", "disk space to erase")

	MBs := int(vi.Available.MBytes())
	OneMBString := grandom.RandomString(int(gvolume.MB.Bytes()))
	bar := pb.New(MBs)
	bar.Start()
	for i := 0; i < MBs; i++ {
		bar.Add(1)
		f.WriteString(OneMBString)
	}
	bar.Finish()
}
