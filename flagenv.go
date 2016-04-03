package flagenv

import (
	"flag"
	"strings"
	"os"
	"strconv"
	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

func String(name string, value string, usage string) *string {
	return flag.String(name, envString(parameterNameToEnvName(name), value), usage)
}

func Int(name string, value int, usage string) *int {
	return flag.Int(name, envInt(parameterNameToEnvName(name), value), usage)
}

func Bool(name string, value bool, usage string) *bool {
	return flag.Bool(name, envBool(parameterNameToEnvName(name), value), usage)
}

func envString(key, def string) string {
	if env := os.Getenv(key); env != "" {
		return env
	}
	return def
}

func envBool(key string, def bool) bool {
	if env := os.Getenv(key); env != "" {
		res, err := strconv.ParseBool(env)
		if err != nil {
			return def
		}

		return res
	}
	return def
}

func envInt(key string, def int) int {
	if env := os.Getenv(key); env != "" {
		val, err := strconv.Atoi(env)
		if err != nil {
			logger.Debugf("invalid value for %q: using default: %q", key, def)
			return def
		}
		return val
	}
	return def
}

func parameterNameToEnvName(name string) string {
	return strings.Replace(strings.ToUpper(name), "-", "_", -1)
}

func Parse() {
	flag.Parse()
}