package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"aduu.dev/utils/errors2"
	"golang.org/x/sys/unix"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang-11 LSM ./bpf/lsm_example.c -- -I../headers

const mapKey uint32 = 0

func run() (err error) {

	// Name of the kernel function to trace.
	// fn := "lsm/file_mprotect"

	// Increase the rlimit of the current process to provide sufficient space
	// for locking memory for the eBPF map.
	if err := unix.Setrlimit(unix.RLIMIT_MEMLOCK, &unix.Rlimit{
		Cur: unix.RLIM_INFINITY,
		Max: unix.RLIM_INFINITY,
	}); err != nil {
		log.Fatalf("failed to set temporary rlimit: %v", err)
	}

	// Load pre-compiled programs and maps into the kernel.
	objs := LSMObjects{}
	if err := LoadLSMObjects(&objs, nil); err != nil {
		log.Fatalf("loading objects: %v", err)
	}
	defer func() {
		if err2 := objs.Close(); err2 != nil {
			err = errors2.CombineErrors(err, err2)
		}
	}()

	_, err = LoadLSM()
	if err != nil {
		log.Fatalf("loading lsm: %v", err)
	}

	/*prog, err := ebpf.NewProgram(&ebpf.ProgramSpec{
		AttachTo:   "lsm/file_mprotect",
		AttachType: ebpf.AttachLSMMac,
		Instructions: asm.Instructions{
			asm.LoadImm(asm.R0, 0, asm.DWord),
			asm.Return(),
		},
		License: "GPL",
		Type:    ebpf.LSM,
	})
	if err != nil {
		return
	}

	defer func() {
		if err2 := prog.Close(); err2 != nil {
			err = errors2.CombineErrors(err, err2)
		}
	}()
	*/

	// Read loop reporting the total amount of times the kernel
	// function was entered, once per second.
	ticker := time.NewTicker(1 * time.Second)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	log.Println("Waiting for events..")

	for {
		select {
		case <-ticker.C:
			// var value uint64
			// if err := objs.KprobeMap.Lookup(mapKey, &value); err != nil {
			// 	log.Fatalf("reading map: %v", err)
			// }
			// log.Printf("%s called %d times\n", fn, value)
			fmt.Println("Small update")
		case <-ctx.Done():
			return
		}
	}
}
