package main

import (
	"fmt"
	"github.com/cheggaaa/pb"
	"github.com/smcduck/xdsa/xrandom"
	"github.com/smcduck/xdsa/xvolume"
	"github.com/smcduck/xsys/xhdd"
	"github.com/smcduck/xsys/xproc"
	"os"
)

func main() {

	_, _, mydir, err := xproc.SelfPath()
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

	vi, err := xhdd.GetVolumeInfo(mydir)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("this program will use random string to erase the volume where itself is")
	fmt.Println(vi.Available.String(), "[", vi.Available.MBytes(), " MB", "]", "disk space to erase")

	MBs := int(vi.Available.MBytes())
	OneMBString := xrandom.RandomString(int(xvolume.MB.Bytes()))
	bar := pb.New(MBs)
	bar.Start()
	for i := 0; i < MBs; i++ {
		bar.Add(1)
		f.WriteString(OneMBString)
	}
	bar.Finish()
}
