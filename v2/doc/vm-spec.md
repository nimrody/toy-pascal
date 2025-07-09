# **Toy Pascal Virtual Machine Specification**

### **1\. Architecture Overview**

The virtual machine (VM) is a stack-based machine designed for simplicity and educational clarity. It consists of four primary memory regions and two essential registers.

* **Code Memory**: A read-only region storing the program's bytecode instructions.  
* **Global Memory**: A static data area for all declared global variables. The size is determined at compile time.  
* **The Heap**: A dynamic memory region for records allocated via the New procedure. This is the memory space managed by the garbage collector.  
* **The Stack**: A unified last-in, first-out (LIFO) data structure used for:  
  * Passing arguments to functions.  
  * Storing local variables.  
  * Holding intermediate operands during expression evaluation.  
  * Managing function call context (return addresses, frame pointers).

The VM's state is controlled by two registers:

* **Instruction Pointer (IP)**: Holds the address in Code Memory of the next instruction to be executed.  
* **Frame Pointer (FP)**: Holds the base address of the current function's activation frame on the Stack.

### **2\. Data Representation**

* **Integer**: A 32-bit signed integer.  
* **Address/Pointer**: A 32-bit unsigned integer representing a memory address. This applies to pointers into the Heap, addresses in Global Memory, and addresses in Code Memory.  
* **nil**: The null pointer constant is represented by the integer value 0\.

### **3\. Stack Frame Layout**

When a user-defined function is called, a new **stack frame** is pushed onto the stack. The stack grows from lower to higher memory addresses. The Frame Pointer (FP) always points to the base of the current frame, allowing for consistent access to locals and arguments.

      \+--------------------+  
      | Operand            |  \<-- Top of Stack (TOS)  
      | ...                |  
      \+--------------------+  
      | Local Variable N   |  (Offset: FP \+ 3 \+ N)  
      | ...                |  
      | Local Variable 1   |  (Offset: FP \+ 3\)  
      \+--------------------+  
FP \-\> | Old Frame Pointer  |  (Offset: FP \+ 2\)  
      \+--------------------+  
      | Return Address     |  (Offset: FP \+ 1\)  
      \+--------------------+  
      | Argument N         |  
      | ...                |  
      | Argument 1         |  (Offset: FP \- NumArgs)  
      \+====================+  
      | Caller's Frame     |  
      | ...                |

### **4\. Instruction Set Format**

Each instruction consists of a **1-byte opcode** followed by an optional **4-byte operand**.

#### **Stack & Constant Operations**

| Opcode | Mnemonic | Operand | Description | Stack Effect |
| :---- | :---- | :---- | :---- | :---- |
| 0x01 | PUSH\_CONST | \<value\> | Pushes a 32-bit constant integer onto the stack. | \[...\] \-\> \[... value\] |
| 0x02 | POP |  | Discards the top value from the stack. | \[... value\] \-\> \[...\] |

#### **Arithmetic & Comparison Operations**

These opcodes pop their operands from the stack (popping b then a) and push the result.

| Opcode | Mnemonic | Description | Stack Effect |
| :---- | :---- | :---- | :---- |
| 0x10 | ADD | Pushes a \+ b. | \[... a, b\] \-\> \[... a+b\] |
| 0x11 | SUB | Pushes a \- b. | \[... a, b\] \-\> \[... a-b\] |
| 0x12 | MUL | Pushes a \* b. | \[... a, b\] \-\> \[... a\*b\] |
| 0x13 | DIV | Pushes a / b (integer division). | \[... a, b\] \-\> \[... a/b\] |
| 0x14 | CMP\_EQ | Pushes 1 if a \= b, else 0\. | \[... a, b\] \-\> \[... 0 or 1\] |
| 0x15 | CMP\_NEQ | Pushes 1 if a \<\> b, else 0\. | \[... a, b\] \-\> \[... 0 or 1\] |
| 0x16 | CMP\_LT | Pushes 1 if a \< b, else 0\. | \[... a, b\] \-\> \[... 0 or 1\] |
| 0x17 | CMP\_GT | Pushes 1 if a \> b, else 0\. | \[... a, b\] \-\> \[... 0 or 1\] |
| 0x18 | CMP\_LE | Pushes 1 if a \<= b, else 0\. | \[... a, b\] \-\> \[... 0 or 1\] |
| 0x19 | CMP\_GE | Pushes 1 if a \>= b, else 0\. | \[... a, b\] \-\> \[... 0 or 1\] |

#### **Variable Access Operations**

| Opcode | Mnemonic | Operand | Description | Stack Effect |
| :---- | :---- | :---- | :---- | :---- |
| 0x20 | LOAD\_GLOBAL | \<address\> | Pushes the value from Global Memory at \<address\>. | \[...\] \-\> \[... value\] |
| 0x21 | STORE\_GLOBAL | \<address\> | Pops a value and stores it in Global Memory at \<address\>. | \[... value\] \-\> \[...\] |
| 0x22 | LOAD\_LOCAL | \<offset\> | Pushes the value from the stack at address FP \+ \<offset\>. | \[...\] \-\> \[... value\] |
| 0x23 | STORE\_LOCAL | \<offset\> | Pops a value and stores it on the stack at address FP \+ \<offset\>. | \[... value\] \-\> \[...\] |

#### **Control Flow Operations**

| Opcode | Mnemonic | Operand | Description | Stack Effect |
| :---- | :---- | :---- | :---- | :---- |
| 0x30 | JUMP | \<address\> | Sets the Instruction Pointer (IP) to \<address\>. | (none) |
| 0x31 | JUMP\_IF\_FALSE | \<address\> | Pops a value. If it is 0, sets IP to \<address\>. | \[... value\] \-\> \[...\] |

#### **Function & Procedure Call Operations**

| Opcode | Mnemonic | Operand | Description | Stack Effect |
| :---- | :---- | :---- | :---- | :---- |
| 0x40 | CALL | \<addr\> \<nArgs\> | Calls a user function or a built-in function. See Section 6\. | (varies) |
| 0x41 | RET |  | Restores caller's stack frame and returns control from a user function. | (cleans frame) |

#### **Heap & Pointer Operations**

| Opcode | Mnemonic | Operand | Description | Stack Effect |
| :---- | :---- | :---- | :---- | :---- |
| 0x50 | NEW | \<size\> | Allocates \<size\> bytes on the heap, pushes the starting address (pointer). | \[...\] \-\> \[... pointer\] |
| 0x51 | LOAD\_INDIRECT |  | Pops an address, reads value from that heap address, pushes value. | \[... addr\] \-\> \[... value\] |
| 0x52 | STORE\_INDIRECT |  | Pops a value, then an address. Stores value at the address on the heap. | \[... val, addr\] \-\> \[...\] |

#### **Machine Control**

| Opcode | Mnemonic | Description |
| :---- | :---- | :---- |
| 0xFF | HALT | Stops program execution. |

### **5\. Garbage Collection Interface**

The garbage collector is triggered during the execution of the NEW instruction. To perform its work, the garbage collector requires read-access to the entire root set:

1. **The Global Memory Area**: To scan for root pointers in global variables.  
2. **The Stack**: To scan the entire stack for root pointers in arguments and local variables of all active functions.

### **6\. Built-in Functions**

To avoid creating a new opcode for every special operation, the VM supports a set of built-in functions that are invoked using the standard CALL instruction.

* **Invocation**: A built-in function is called by using a special **negative address** in the CALL instruction's \<addr\> operand.  
* **Execution**: When the VM executes CALL with a negative address, it does not jump. Instead, it executes native Go code corresponding to that built-in's ID. It does not create a new stack frame in the same way a user function call does.  
* **Argument Passing**: Arguments are pushed onto the stack just like for a regular function call. The native Go code is responsible for popping its arguments from the stack.

#### **Defined Built-ins**

| ID / Address | Name | Arguments | Description |
| :---- | :---- | :---- | :---- |
| \-1 | print | value: Integer | Pops an integer from the stack and prints it to standard output. |

