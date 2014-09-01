#include <stdlib.h>
#include <libguile.h>

#include "_cgo_export.h"

static SCM
log_now (void)
{
    LogNow();
    return SCM_BOOL_F;
}

static void
shell_main (void *data, int argc, char **argv)
{
    scm_c_define_gsubr ("log_now", 0, 0, 0, log_now);
    scm_shell (argc, argv);
}

void
LaunchGuile ()
{
    scm_boot_guile (0, NULL, shell_main, 0);
}
