#Java

## Variables

* Od Java SE 10 mozna deklarowac zmienne przy uzyciu ```var```, kompilator decyduje jakiego zmienna jest wtedy typu

* Mozna je uzywac w metodach, kosntruktorach i blokach

* zmienna zadeklarowana przez ```var``` musi byc zainicjalizowana

## Operators

* operatory ewaluuja od lewej do prawej, z wyjatkiem assignment operatora (=) ktory ewaluuje od prawej do lewej

* Operatory sa ewaluowane w nastepujacej kolejnosci:

| Operators | Precedence |
| ------------ | ------------ |
| postfix | expr++ expr-- |
| unary | ++expr --expr +expr -expr ~ ! |
| multiplicative | * / % |
| additive | + - |
| shift | << >> >>> |
| relational | < > <= >= instanceof |
| equality | == != |
| bitwise AND | & |
| bitwise XOR | ^ |
| bitwise OR | | |
| logical AND | && |
| logical OR | || |
| ternary | ? : |
| assignment | = += -= %= /= &= ^= |= <<= >>= >>>= |

* operator ```instanceof``` sprawdza czy dany obiekt jest instancja okreslonej klasy, subklasy, lub korzysta z okreslonego interfejsu 
```
class Parent {}
class Child extends Parent implements MyInterface {}
interface MyInterface {}

class InstanceofDemo {
    public static void main(String[] args) {

        Parent obj1 = new Parent();
        Parent obj2 = new Child();

        System.out.println("obj1 instanceof Parent: "
            + (obj1 instanceof Parent));
        System.out.println("obj1 instanceof Child: "
            + (obj1 instanceof Child));
        System.out.println("obj1 instanceof MyInterface: "
            + (obj1 instanceof MyInterface));
        System.out.println("obj2 instanceof Parent: "
            + (obj2 instanceof Parent));
        System.out.println("obj2 instanceof Child: "
            + (obj2 instanceof Child));
        System.out.println("obj2 instanceof MyInterface: "
            + (obj2 instanceof MyInterface));
    }
}

// output: 
//obj1 instanceof Parent: true
//obj1 instanceof Child: false
//obj1 instanceof MyInterface: false
//obj2 instanceof Parent: true
//obj2 instanceof Child: true
//obj2 instanceof MyInterface: true
```

# Related to: 

* https://dev.java/learn/language-basics/using-operators/


#c #operators #instanceof #evaluation
