# Java

## Expressions

* do while

* if 

* if elseif

* while

* for ( initialize; condition; incrementation) {}

	* przy 1 iteracji nie wykonuje sie inkrementacja / dekrementacja dopiero przy 2 wzw

	* kazde z wyrazen jest opcjonalne, e.g. (for (; ; ) {} // infinite loop

* petla for dla Collections i arrays:

	* for (initialize : array / collection) {}

* break / continue w loopach

* break / continue moze byc **labeled** albo **unlabeled**, labeled break przydaje sie w przypadku nested loopow, gdy chcemy kompletnie wyjsc z calej iteracji a jestesmy w nested loopie. przyklad: 
```java
class BreakWithLabelDemo {
    public static void main(String[] args) {

        int[][] arrayOfInts = {
            {  32,   87,    3, 589 },
            {  12, 1076, 2000,   8 },
            { 622,  127,   77, 955 }
        };
        int searchfor = 12;

        int i;
        int j = 0;
        boolean foundIt = false;

    search:
        for (i = 0; i < arrayOfInts.length; i++) {
            for (j = 0; j < arrayOfInts[i].length;
                 j++) {
                if (arrayOfInts[i][j] == searchfor) {
                    foundIt = true;
                    break search;
                }
            }
        }

        if (foundIt) {
            System.out.println("Found " + searchfor + " at " + i + ", " + j);
        } else {
            System.out.println(searchfor + " not in the array");
        }
    }
}
//output
//Found 12 at 1, 0
```

* return 

* yield - ten statement nie moze byc voidem, zawsze musi cos zwrocic, wykorzystywany w **switch(java 14)**

* switch

	* switch moze byc uzywany z nast. typami:
	> byte, short, char, and int primitive data types
	Character, Byte, Short, and Integer wrapper types
	enumerated types
	the String type.

	* od java se 14 syntax switcha mozna pisac w prostszy i czytelniejszy sposob:
	```java
// kiedys

int day = ...; // any day
int len = 0;
switch (day) {
    case MONDAY:
    case FRIDAY:
    case SUNDAY:
        len = 6;
        break;
    case TUESDAY:
        len = 7;
        break;
    case THURSDAY:
    case SATURDAY:
        len = 8;
        break;
    case WEDNESDAY:
        len = 9;
        break;
}
System.out.println("len = " + len);

//teraz

int day = ...; // any day
int len =
    switch (day) {
        case MONDAY, FRIDAY, SUNDAY -> 6;
        case TUESDAY                -> 7;
        case THURSDAY, SATURDAY     -> 8;
        case WEDNESDAY              -> 9;
    }
System.out.println("len = " + len);
	```



# Related to: 

* https://dev.java/learn/language-basics/expressions-statements-blocks/

#Expressions #Java #Switch #Java14 #yield #return #if #while #dowhile #for #fore 
