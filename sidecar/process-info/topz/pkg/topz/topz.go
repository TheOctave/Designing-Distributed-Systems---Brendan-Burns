package topz

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"text/tabwriter"

	"github.com/shirou/gopsutil/process"
)

func handleError(res http.ResponseWriter, err error) {
	res.WriteHeader(http.StatusInternalServerError)
	res.Write([]byte(err.Error()))
}

type ProcInfo struct {
	PID           int32
	MemoryPercent float32
	MemoryInfo    *process.MemoryInfoStat
	CPUPercent    float64
	Command       string
}

func HandleRequest(res http.ResponseWriter, req *http.Request) {
	pids, err := process.Pids()
	if err != nil {
		handleError(res, err)
		return
	}

	procs := []*ProcInfo{}
	wg := sync.WaitGroup{}
	for idx, pid := range pids {
		fmt.Println(pid)
		proc, err := process.NewProcess(pid)
		if err != nil {
			continue
		}

		wg.Add(1)
		go func(i int) {
			var err error
			p := &ProcInfo{}
			p.PID = pids[i]
			if p.Command, err = proc.Cmdline(); err != nil {
				log.Printf("Error getting Command Line: %v", err)
			}
			if p.MemoryInfo, err = proc.MemoryInfo(); err != nil {
				log.Printf("Error getting memory info: %v", err)
			}
			if p.MemoryPercent, err = proc.MemoryPercent(); err != nil {
				log.Printf("Error getting Memory usage: %v", err)
			}
			if p.CPUPercent, err = proc.CPUPercent(); err != nil {
				log.Printf("Error getting CPU usage: %v", err)
			}
			if len(p.Command) > 0 {
				procs = append(procs, p)
			}
			wg.Done()
		}(idx)
	}
	wg.Wait()

	res.WriteHeader(http.StatusOK)
	w := tabwriter.NewWriter(res, 0, 0, 1, ' ', 0)

	for _, proc := range procs {
		if proc == nil {
			continue
		}
		fmt.Fprintf(w, "%d\t%g\t%g\t%s\n", proc.PID, proc.CPUPercent, proc.MemoryPercent, proc.Command)
	}
	w.Flush()
}
