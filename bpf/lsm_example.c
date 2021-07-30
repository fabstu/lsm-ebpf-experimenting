#include "vmlinux.h"
#include "bpf_helpers.h"
#include "bpf_helper_defs.h"
#include "my_helpers.h"

#include <errno.h>
#include <string.h>

char str[20] = "DEBUG:LSM! \n";

char __license[] SEC("license") = "Dual MIT/GPL";

struct {
	__uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
} events SEC(".maps");

SEC("lsm_file_mprotect")
int BPF_PROG(mprotect_audit, struct vm_area_struct *vma,
             unsigned long reqprot, unsigned long prot, int ret)
{
        bpf_trace_printk(str, strlen(str));
        /* ret is the return value from the previous BPF program
         * or 0 if it's the first hook.
         */
        if (ret != 0)
                return ret;

        int is_heap;
        is_heap = (vma->vm_start >= vma->vm_mm->start_brk &&
                   vma->vm_end <= vma->vm_mm->brk);

        /* Return an -EPERM or write information to the perf events buffer
         * for auditing
         */
        if (is_heap)
                return -EPERM;
}
