#C Programming

## Lifetime

* Period during program execution when variable or function exists. Storage duration determines it

* ```static``` storage class specifier have static duration for the whole duration of program (called global), it is initialized before program startup

* identifier declared without ```static``` has automatic (local) storage duration, it is block-scoped, it loses its storage after leaving block and creating new one after entering it every time

* All functions have global lifetime, same as identifiers declared in file-scope (outside of block)

* If a local variable is declared, it is reinitialized every time program reaches block (EXCEPT IT IS DECLARED AS STATIC, THEN IT IS INITIALIZED ONLY ONCE [FIRST RUN OF BLOCK CODE] AND IT REMAINS DURING WHOLE PROGRAM









#c #programming #lifetime #scope #visibility #linkage
