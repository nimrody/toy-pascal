# Toy Pascal Language Specification

### 1. General Characteristics

* Typing: The language is *statically and strongly typed*. Type checking is
  performed at compile time.  

* Memory Management: Memory is managed automatically via a *pluggable garbage
  collector*. The programmer allocates memory for records but does not need to
  deallocate it.  

* Structure: A program consists of a type declaration section, a var
  declaration section, and a sequence of function and procedure definitions.

### 2. Types

The language supports three kinds of types: primitive types, record types, and
pointer types.

#### 2.1 Primitive Type

* Integer: A 32-bit signed integer.

#### 2.2 Pointer Types

A pointer type defines a reference to a value of another type. Pointers are
necessary for creating dynamic data structures like linked lists and trees.

* Syntax: ^TypeName  
* Example:
  type  
    NodePtr = ^Node; // A pointer to a Node record

* Special Value: The keyword nil represents a null pointer that does not point
  to any object.

#### 2.3 Record Types

A record is a composite data type that groups variables (fields) under a single
name.

* Syntax:
  type  
    TypeName = record  
      FieldName1: Type1;  
      FieldName2: Type2;  
      ...  
    end;

* Fields: Fields can be of type Integer, a pointer type, or another (previously
  defined) record type.

### 3. Declarations

#### 3.1 Type Declarations

All custom types (records and pointers) must be declared in a type block before they are used.

* Example:
  type  
    NodePtr = ^Node;  
    Node = record  
      Data: Integer;  
      Next: NodePtr;  
    end;

#### 3.2 Variable Declarations

All global and local variables must be declared in a var block.

* Syntax:
  var  
    VariableName1: TypeName;  
    VariableName2, VariableName3: AnotherType;

* Scope: Variables can be declared globally for the entire program or locally
  within a function/procedure.

#### 3.3 Function and Procedure Declarations

* A *procedure* is a subprogram that does not return a value.  
* A *function* is a subprogram that returns a value.  
* Procedure Syntax:  
  procedure ProcedureName(param1: Type1; param2: Type2);  
  var  
    // local variables  
  begin  
    // procedure body  
  end;

* Function Syntax:
  function FunctionName(param1: Type1): ReturnType;  
  var  
    // local variables  
  begin  
    // function body  
    FunctionName := returnValue; // Return value is assigned to the function's name  
  end;

* Parameters: Parameters are *passed by value*.

### 4. Statements

#### 4.1 Assignment

* Syntax: variable := expression;  
* Rules: The type of the expression must be compatible with the type of the
  variable.

#### 4.2 Compound Statements

A sequence of statements can be grouped into a single block using begin and end.

* Syntax:
  begin  
    statement1;  
    statement2;  
    ...  
  end

#### 4.3 Conditional Statements (if)

* Syntax:  
  if condition then  
    statement1  
  else  
    statement2;

* The else part is optional. The condition must evaluate to a boolean result.

#### 4.4 Loop Statements (while)

* Syntax:
  while condition do  
    statement;

* The statement is executed repeatedly as long as the condition is true.

### 5. Expressions

#### 5.1 Literals

* Integer Literals: e.g., 10, 0, -5.  
* Pointer Literal: nil.

#### 5.2. Variable and Field Access

* Variable: myVariable  
* Dereferencing and Field Access: To access the fields of a record that a
  pointer refers to, use the ^. notation:
  * Syntax: pointerVariable^.fieldName  
  * Example: Current^.Data

#### 5.3. Operators

* Arithmetic (for Integer): +, -, \*, / (integer division).  
* Comparison: =, <>, <, >, <=, >=. These can be used with integers and pointers
  (for checking equality with nil or other pointers).

### 6. Memory Management and Garbage Collection

The core of the language's runtime is its automated memory management system.

#### 6.1. Allocation (New)

* Syntax: New(pointerVariable);  
* Action: The New procedure allocates memory on the heap for the type the
  pointerVariable points to. The pointer is then updated to reference this new
  memory block.

#### 6.2 Deallocation (Automatic)

* There is *no explicit Dispose or free procedure*.
* The garbage collector automatically identifies and reclaims memory that is no longer in use.

#### 6.3 Reachability and the Root Set

For the garbage collector to function, it must identify all "live" objects. An
object is considered live if it is reachable by following a chain of pointers
from a "root."

* The *root set* consists of:  
  1. All *global variables*.  
  2. All *local variables* and parameters on the active call stack.

#### 6.4 Pluggable Garbage Collection Architecture

The language runtime is designed to allow different garbage collection
algorithms to be *swapped in and out*. The initial implementation will use
the Mark-and-Sweep algorithm.

* Default Algorithm: Stop-the-World Mark-and-Sweep  
  The default garbage collector is a stop-the-world mark-and-sweep collector.
  Program execution is paused entirely while the GC cycle runs.  

  1. Mark Phase: The collector starts from the root set and traverses every
     reachable object in the heap, marking each one as "in-use."  

  2. Sweep Phase: The collector scans the entire heap. Any object that was not
     marked during the Mark phase is considered garbage and its memory is
     reclaimed.

