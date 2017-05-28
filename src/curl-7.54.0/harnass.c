
#include <stdbool.h>
#include "lib/url.h"

char buf[4096];
int main() {

  char **user = NULL;
  char **passwd = NULL;

  fread(buf, 1, 4096, stdin);
  parse_login_details(buf, strlen(buf), user, passwd, NULL);
}
