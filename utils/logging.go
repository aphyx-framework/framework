package utils

import (
	"github.com/TwiN/go-color"
	"log"
	"os"
)

// Logging ---
//
// This function is used to log messages to the console
// It will log the message to the console with the appropriate color
// on *nix systems or in the console based on the type of message
var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

// LoadLogger ---
//
// This function is used to load the logger
func LoadLogger() {
	InfoLogger = log.New(os.Stdout, color.Black+color.CyanBackground+" INFO:     "+color.Reset+" ", log.Ldate|log.Ltime)
	WarningLogger = log.New(os.Stdout, color.Black+color.YellowBackground+" WARNING:  "+color.Reset+" ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(os.Stderr, color.Black+color.RedBackground+" ERROR:    "+color.Reset+" ", log.Ldate|log.Ltime)
}
