# Java

## Parameter Types


* Po uzyciu tej metody, w runtime zmienne w parametrach zamieniane sa na wartosci, wazne aby byly w odpowiedniej kolejnosci

* Mozemy uzywac kazdego typu w parametrach, nawet reference types czyli obiekty i tablice

```java
public Polygon polygonFrom(Point[] corners) {
    // method body goes here
}
// w tym przypadku metoda tworzy nowy obiekt Polygon oraz przyjmuje tablice obiektow Point jako argument
```

* Do parametru mozna wpuscic nie skonczona liczbe argumentow, koncept ten nazywa sie **varargs**, uzywamy go wtedy gdy nie wiemy ile argumentow o okreslonym typie wrzucimy do metody, to shortcut aby nie tworzyc tablicy manualnie

```java
public Polygon polygonFrom(Point... corners) {
    int numberOfSides = corners.length;
    double squareOfSide1, lengthOfSide1;
    squareOfSide1 = (corners[1].x - corners[0].x)
                     * (corners[1].x - corners[0].x) 
                     + (corners[1].y - corners[0].y)
                     * (corners[1].y - corners[0].y);
    lengthOfSide1 = Math.sqrt(squareOfSide1);

    // more method body code follows that creates and returns a 
    // polygon connecting the Points
}
// w srodku metody varargs jest traktowana jako tablica
```

### Parameter Names

* Parametr musi miec unikalna nazwe, nie moze sie nazywac tak samo jak cos w jego scopie, czyli inny parametr lub lokalna zmienna.

* Parametr moze miec nazwe taka sama jak jedno z pol klasy, jesli cos takiego wystepuje zachodzi cos takiego jak ukrycie pola (**Field Shadowing**), powinno sie tego unikac poniewaz zmniejsza czytelnosc kodu.**Konwencjonalnie jest to uzywane tylko w _setterach_**

```java
public class Circle {
    private int x, y, radius;
    public void setOrigin(int x, int y) {
        ...
    }
}
// wystepuje tutaj field shadowing, ktory w scopie metody ukrywa zmienne zainicjalizowane wyzej, i uzywa tych uzytych w parametrze. Aby uzyc pola ktore jest zadeklarowane w scopie klasy, nalezy odwolac sie do niego poprzez slowo kluczowe this
```

### Passing primitive data type arguments

* Podanie prymitywnego typu do parametru wystepuje poprzez wartosc, to znaczy ze tylko w scopie tej metody ten parametr moze ulec zmianie, czyli zmienna x zainicjalizowana wartoscia 3 w global scope, gdy zostanie przekazana do metody ktora zmienia ta wartosc na inna niz 3, finalnie zmieni ta wartosc tylko w scope tej metody, ta zmienna w global scope nadal bedzie 3

### Passing reference data type arguments

* Podanie referencyjnego typu danych (tablicy / obiektu) do parametru tez nastepuje poprzez wartosc, alew zasadzie dokladniej to sama w sobie referencja jest ta wartoscia a w zasadzie to kopia referencji do tego obiektu xD Co daje nam referencja tego obiektu ? Sprawia ze mozemy zmodyfikowac okreslone pola ktore sa dla nas dostepne (o dostepnosci decyduja modyfikatory tych pol, ale przeciez i tak nie modyfikujemy pol bezposrednio bo to niezgodne z praktyka, a settery i gettery z reguly sa publiczne wiec whatever). Jak metoda zakonczy dzialanie to obiekt w global scope nadal jest ten sam, ma ten sam adres, i ma zmienione okreslone pola. **ATENZIONE** proba ponownego przypisania nowego obiektu do kopi referencji podanej do parametru, nie ma jakiegokolwiek wplywu na oryginalny obiekt, w koncu to kopia referencji xD

# Related to: 

* https://dev.java/learn/classes-objects/calling-methods-constructors/


#Class #Reference #Parameter #Primitive #parameternames #names #referencetype #java #types #parametertypes
