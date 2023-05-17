#Java

## Basics

### Variables

* Instance variables (Non-Static Fields) - te pola sa tworzone przy kazdej nowej instancji klasy (zmienne ktore nie sa zadeklarowane ze slowem **static**

* Class variables (Static fields) -  Te pola sa dedykowane dla klasy (dla kazdej instancji klasy ta zmienna jest taka sama) [zmienne zadeklarowane ze slowem **static**]. Przy uzyciu slowa **final** deklarujemy ta zmienna jako stala

* Local variables - Zmienne zadeklarowane w metodzie, sa widoczne i dostepne tylko w zakresie tej metody

* Parameters - zmienne ktore sa przekazywane jako parametr funkcji / metody

### Naming variables

* Konwencja jest taka aby zawsze zaczynac od znaku [a-z], w camelCase


### Primitive types

* byte - 8 bits (-128 - +127)

* short - 16 bits (-32,768 - + 32,767)

* int - 32 bits (-2^31 - +2^31-1)

* long - 64 bits (-2^63 - +2^63-1)

* float - 32 bits

* double - 64 bits

* boolean - size undefined

* char - 16bit unicode character (\u0000 [0] - \uffff [65,535])

### Declaring and initializing variable

* Deklaracja zmiennej sprawi ze zmienna przyjmie defualt value (dla okreslonego typu sa inne, z reguly 0, false, null (dla Stringa lub innego obiektu) - **DEKLARACJA ZMIENNEJ BEZ JEJ INICJALIZACJI JEST BRANA JAKO ZLA PRAKTYKA**, *Zmienna nie przyjmuje default value gdy jest to local variable (zadeklarowana w zakresie metody), gdy nie zainicjalizujemy takiej zmiennej, proba dostania sie do niej zwroci blad w trakcie kompilacji*

### Integer Literals

* Wartosc ktora chcemy aby byla o typie long musi byc rowniez zakonczona litera l / L **Uzywamy uppercase, bo lowercase to zla praktyka ze wzgl na czytelnosc**

* Literal inta moze byc reprezentowany przez wartosc decymalna, binarna, hexadecymalna.

### Floating point literals

*  Jesli zmeinna jest typu float, jego zainicjalizowana artosc musi byc zakonczona litera f / F

* W przypadku zmiennej typu double, jego zainicjalizowana wartosc **moze** byc zakonczona litera d / D

* Przy obydwu typach mozemy uzywac notacji wykladniczej *e*

### Character and String literals

* Te literaly moga pomiescic Unicode Characters (UTF-16), jesli edytor na to pozwala mozemy uzywac ich bezposrednio w kodzie zrodlowym, jesli nie to trzeba uzywac **Unicode Escape** e.g. *\u0108* (capital C)

* Dla **char** literalow uzywaj single quotes **' '** 

* Dla **string** Literalow uzywaj double quotes **" "**

* Java supportuje escape sequences tj. **\b, \t, \n, \f, \r, \", \', \\ ,**

* W przypadku javy jest jeszcze specjalny literal **null**, moze byc on przyznany do kazdej zmiennej ktora niejest prymitywem. Z reugly uzywa sie jej aby pokazac ze obiekt jest niedostepny

* Jest jeszcze specjalny literal klasowy ktorego uzywa sie jako append do typename - **.class**, e.g. **String.class**

### Underscore character in Numeric Literals

* Mozemy dividowac numer co 3 cyfry w decimal oraz co 4 znaki w hexie, binarnym uzywajac underscore \_

### Array

* Tablice deklarujemy tak int[] anArray

* Tablice mozna zadeklarowac i od raz zainicjalizowac uzywajac slowa **new** ```int[] anArray = new int[10];``` lub tak ```int[] anArray = {100, 200, 300, 
400, 500, 600, 
700, 800, 900, 1000 
}```

* Tablica tablic deklarowana jest w nast. sposob ```String[][] names = { {"Mr.", "Mrs."}, {"Smith", "Jones"} }```

* Tablica moze byc kazdego typu, nie tylko prymitywnego

* Tablice mozna kopiowac uzywajac metody **arraycopy()** z klasy **System**

# Related to: 

* https://dev.java/learn/language-basics/



#Array #Table #Variable #Literal #integer #string #char #byte #int #double #float #long #boolean #short
