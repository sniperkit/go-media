#include "go_readline.h"

int keyCallback(int arg0, int invokingKey) {
	return goKeyCallback(arg0, invokingKey);
}
