// #include <linux/fs.h>
#include <errno.h>
// #include <linux/bpf.h>
#include "bcc_helper.h"

LSM_PROBE(settime, const struct timespec64 *ts, const struct timezone *tz)
{
    return -EPERM;
}
