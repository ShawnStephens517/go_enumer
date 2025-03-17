// +build windows

//package windows_actions
/*
The intent of this action is to "Embed" a secondary Go binary into the memory of the Current App
The idea is to be extra stealthy with our execution. Run additional "nefarious" actions from memory.
Should be used to help enhance Detection rules for these types of scenarios outside of the default
Atomic Red Team atomics (Powershell IEX/ IWR;IEX)

TODO
[] Write secondary application - Possibly use CTF_Serve. This is a simple Web Server that hosts the files in a specific Directory or could use default settings
[] Review github.com/amenzhinsky/go-memexec
[] Call the secondary binary at runtime to be executed similar to ATR powershell version (if possible)
*/
