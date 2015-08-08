// +build !bin
// http://blog.golang.org/profiling-go-programs

package commands

import (
	"flag"
	"os"
	"runtime/pprof"
)

type ProfileHeap struct {
	Log    *Logger
	Output string
	File   *os.File
}

func ProfileHeapCLI(cli *flag.FlagSet) CommandLineHandler {
	ph := &ProfileHeap{
		Log: NewLogger("[ostent profile-heap] "),
	}
	cli.StringVar(&ph.Output, "profile-heap", "", "Profiling heap output `filename`")
	return func() (AtexitHandler, bool, error) {
		if ph.Output == "" {
			return nil, false, nil
		}
		return ph.Atexit, false, ph.Run()
	}
}

func (ph *ProfileHeap) Atexit() {
	ph.Log.Print("Writing heap profile")
	if err := pprof.Lookup("heap").WriteTo(ph.File, 1); err != nil {
		ph.Log.Print(err) // just print
	}
	if err := ph.File.Close(); err != nil {
		ph.Log.Print(err) // just print
	}
}

func (ph *ProfileHeap) Run() (err error) {
	ph.File, err = os.Create(ph.Output)
	if err != nil {
		ph.Log.Fatal(err)
	}
	return err
}

/* ******************************************************************************** */

type ProfileCPU struct {
	Log    *Logger
	Output string
	File   *os.File
}

func ProfileCPUCLI(cli *flag.FlagSet) CommandLineHandler {
	pc := &ProfileCPU{
		Log: NewLogger("[ostent profile-cpu] "),
	}
	cli.StringVar(&pc.Output, "profile-cpu", "", "Profiling CPU output `filename`")
	return func() (AtexitHandler, bool, error) {
		if pc.Output == "" {
			return nil, false, nil
		}
		return pc.Atexit, false, pc.Run()
	}
}

func (pc *ProfileCPU) Atexit() {
	pc.Log.Print("Writing CPU profile")
	pprof.StopCPUProfile()
	if err := pc.File.Close(); err != nil {
		pc.Log.Print(err) // just print
	}
}

func (pc *ProfileCPU) Run() (err error) {
	pc.File, err = os.Create(pc.Output)
	if err == nil {
		err = pprof.StartCPUProfile(pc.File)
	}
	if err != nil {
		pc.Log.Fatal(err)
	}
	return err
}

func init() {
	AddCommandLine(ProfileCPUCLI)
	AddCommandLine(ProfileHeapCLI)
}
