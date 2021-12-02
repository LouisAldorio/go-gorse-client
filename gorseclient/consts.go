package gorseclient

type DebugMode int

const (
	DEBUG_INFO DebugMode = iota + 1
	DEBUG_ERROR
	DEBUG_SILENT
)