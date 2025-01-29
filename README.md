# Silent Errors from Accessing Non-Existent Map Keys in Go

This example demonstrates a common pitfall in Go when working with maps: accessing a non-existent key silently returns the zero value.  This can lead to subtle, hard-to-debug errors.

The `bug.go` file shows code that attempts to access a key that might not be present.  It illustrates how the program continues without a runtime panic, potentially producing incorrect results.

The `bugSolution.go` file demonstrates how to safely handle this situation by checking for the key's existence before attempting to access its value.