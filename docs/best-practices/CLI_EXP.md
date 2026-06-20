# Respect POSIX args
Basic form: utility_name [options] [operands]
Example:
ls -l /home/user
- utility: ls
- options: -l
- operands: /home/user

1. Short options (-<letter>)
2. Options can be grouped together 
3. -- separator: stop parsing options; everything after this is an operand

# Empathic
Use little message to help user recover from failure and know what to do next.