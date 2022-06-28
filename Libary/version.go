package Libary

import "runtime"

func Version() string {
	return runtime.Version()
}
