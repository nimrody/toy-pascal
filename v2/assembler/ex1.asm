PUSH_CONST 100
PUSH_CONST 20
ADD

PUSH_CONST 2
MUL

PUSH_CONST 50
SUB

# Store result in global 0
STORE_GLOBAL 0

# Load from global and print
LOAD_GLOBAL 0
CALL -1 1  # Call built-in print

# Print a constant directly
PUSH_CONST 42
CALL -1 1  # Call built-in print

HALT
