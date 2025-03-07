# Go_Enumer
Go Based Enumeration on Windows and Linux. 

## Lin_Enumer
Linpeasish

## Win_Enumer
WinPeasish


## Advanced - Nefarious
May perform additional invasive attacks. Not ran by default. Use at your own risk.
Examples:
 #### Linux:
 ---
 - GTFObins attempts (timed)
 - 

 #### Windows:
 ---
 - DLL Injection
 - 



## Disclaimer
Intended as a educational and/or sanctioned tool for enumerating and some light attacks. Use this tool only on systems you own or with explicit **written permission** during Security Assesments (**Signed Rules of Engagement**).

*Obfuscate for a sneakier binary (AV/EDR). Actions this tool runs, should light up SIEM or other monitoring tools.*

`garble -literals build .`


# TODO
- Makefile for easy builds
 - Flag to create benign payloads pre-build, save shellcode to a configfile (TOML, YML, JSON)(config file uploadable)
 - Build with configfile to inject payloads at build time
 - Obfuscate flag to Garble the code
- Add Benign attacks/POC code for enhanced Detection options
- Ensure code/project includes any applicable Licenses from other projects used as motivation or code.  
