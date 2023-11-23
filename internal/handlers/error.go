/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import "log"

func CheckError(err error, msg string) {
	if err != nil {
		log.Println(msg)
		log.Fatalf("ERROR: %v", err)
	}
}
