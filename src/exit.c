#include <stdlib.h>
#include <stdio.h>
#include <getopt.h>
#include <unistd.h>
#define _POSIX_C_SOURCE 2

int main(int argc, char **argv) {
	extern char *optarg;
	int ch, i;

	i = 0;
	while ((ch = getopt(argc, argv, "i:")) != -1) {
		switch (ch) {
			case 'i':
				i = atoi(optarg);
				break;
			case '?':
			default:
				break;
		}
	}

	exit(i);
}
