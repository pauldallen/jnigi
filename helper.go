package jnigi

import (
	"os"
	"runtime"
	"log"
	"path/filepath"
)

// AttemptToFindJVMLibPath tries to find the full path to the JVM shared library file
func AttemptToFindJVMLibPath() string {
	prefix := os.Getenv("JAVA_HOME")
	log.Println("got prefix" + prefix)
	if prefix == "" {
		if runtime.GOOS == "windows" {
			prefix = filepath.Join("c:", "Program Files", "Java", "jdk")
		} else if runtime.GOOS == "darwin" {
			prefix = "/Library/Java/Home"
		} else {
			prefix = "/usr/lib/jvm/default-java"
		}
	}
	dirPath := filepath.Join(prefix, "jre", "lib", runtime.GOARCH, "server")
	log.Println("got dirPath" + dirPath)
	var libPath string
	if runtime.GOOS == "windows" {
		libPath = filepath.Join(dirPath, "jvm.dll")
	} else if runtime.GOOS == "darwin" {
		libPath = filepath.Join(dirPath, "libjvm.dylib")
	} else {
		libPath = filepath.Join(dirPath, "libjvm.so")
	}
	log.Println("returning libPath" + libPath)
	
	return libPath
}
