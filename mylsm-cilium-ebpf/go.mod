module aduu.dev/acs/ebpf-rawlinux/mylsm

go 1.16

require (
	aduu.dev/utils v0.3.1
	github.com/cilium/ebpf v0.6.2
	github.com/spf13/cobra v1.2.1
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c
)

replace github.com/cilium/ebpf => ../../ebpf
