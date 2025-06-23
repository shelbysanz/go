package main

import (
	"fmt"
	"reflect"
)

/*
 Structs use memory in contiguous blocks, so the ordering of the structs
 can have a big impact on memory usage. Its best practice to align the
 fields by size, largest to smallest. As a reminder, I can use the reflect
 package to debug the memory layout of the struct.
*/

type contact struct {
	userID       string
	sendingLimit int32
	age          int32
}

type perms struct {
	permissionLevel int
	canSend         bool
	canReceive      bool
	canManage       bool
}

func main() {
	contact_type := reflect.TypeOf(contact{})
	fmt.Printf("Contact is %d bytes\n", contact_type.Size())

	perms_type := reflect.TypeOf(perms{})
	fmt.Printf("Perms is %d bytes\n", perms_type.Size())
}
