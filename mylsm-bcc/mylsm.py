import os
import sys
import time

from bcc import BPF, libbcc

if not BPF.support_lsm():
        print("LSM not supported")

src1 = """
#include <linux/fs.h>
#include <uapi/asm-generic/errno-base.h>
LSM_PROBE(settime, const struct timespec64 *ts, const struct timezone *tz)
{
    return -EPERM;
}
"""

b = BPF(text=src1)

while True:
    time.sleep(1)

