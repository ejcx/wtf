#include <stdlib.h>
#include <stdio.h>
#define _POSIX_C_SOURCE 2

int main(int argc, char **argv) {
	int i;
	char buf[4096];

	if (read(0, buf, 4096) < 1) {
		exit(1);
	}
	i = atoi(buf);
	exit(i);
}
