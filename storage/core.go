package storage

import "os"

var APIUsername = os.Getenv("LISX_API_USERNAME")
var APIPassword = os.Getenv("LISX_API_PASSWORD")
var APIServer = os.Getenv("LISX_API_SERVEER")
var Token string