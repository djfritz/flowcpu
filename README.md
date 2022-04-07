# Gravwell Flows CPU

This is a simple, 5 stage pipelined, subset of the MIPS ISA implemented as a set of Gravwell Flows. 

It supports the following instructions:

- nop (no operation)
- addiu (add immediate unsigned)
- lb (load byte)
- sb (store byte)
- beq (branch if equal)
- j (jump)

It additionally has a character addressed "terminal" memory mapped at address 0x40. When run with the "terminal" program (running on the same host as the webserver), it will display characters written to it.

The included program in memory.csv is a simple "Hello world!" loop that writes to the terminal.

# To setup

Quick version - just load the included backup file and enable the schedule on all the flows.

Less quick version - Load each of the flow JSON blobs, as well as each of the csv files as resources (named by their filename, eg "memory.csv" should be named "memory"). Then schedule all the flows.

# To reset

Reupload the memory, rf, and pc resources. The pipeline stages do not need to be cleared.
