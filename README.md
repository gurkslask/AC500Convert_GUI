# AC500 Converter
## Introduction
This GUI is used to add an communication address for variables used in communication from a ABB AC500 PLC. Can work with other Codesys systems as well.
## Usage
In the left pane resides the variables that are to be addressed. By default it is already populated with some test variables. Both UINT and BOOL are supported.

If variables already has addresses then the program will keep that address for the variable and continue incrementing from that point for the variables that follows.
## Jumps
Functionality for making jumps in addresses is accessed via writing "RJUMP XXX" for jumping registers (UINT), or "BJUMP XXX" for jumping bits (BOOL).
## Output
In the left pane resides the choice for which protocol that is chosen. Modbus and Comli are supported as of now (2020-09-09).

In the top field resides the choice of what device the output is made for, either PLC code or csv format for importing to ix Developer or ABB Panel builder 800.