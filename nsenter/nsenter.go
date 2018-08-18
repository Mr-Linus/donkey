package nsenter

/*
#include <errno.h>
#include <sched.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>
__attribute__((constructor)) void enter_namespace(void) {
	char *donkey_pid;
	donkey_pid = getenv("donkey_pid");
	if (donkey_pid) {
		//fprintf(stdout, "got donkey_pid=%s\n", donkey_pid);
	} else {
		//fprintf(stdout, "missing donkey_pid env skip nsenter");
		return;
	}
	char *donkey_cmd;
	donkey_cmd = getenv("donkey_cmd");
	if (donkey_cmd) {
		//fprintf(stdout, "got donkey_cmd=%s\n", donkey_cmd);
	} else {
		//fprintf(stdout, "missing donkey_cmd env skip nsenter");
		return;
	}
	int i;
	char nspath[1024];
	char *namespaces[] = { "ipc", "uts", "net", "pid", "mnt" };
	for (i=0; i<5; i++) {
		sprintf(nspath, "/proc/%s/ns/%s", donkey_pid, namespaces[i]);
		int fd = open(nspath, O_RDONLY);
		if (setns(fd, 0) == -1) {
			//fprintf(stderr, "setns on %s namespace failed: %s\n", namespaces[i], strerror(errno));
		} else {
			//fprintf(stdout, "setns on %s namespace succeeded\n", namespaces[i]);
		}
		close(fd);
	}
	int res = system(donkey_cmd);
	exit(0);
	return;
}
*/
import "C"