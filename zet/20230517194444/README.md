# Java

# Classes and Objects

## Classes

```java
// ojciec

public class Bicycle {

}

// synek

public class MountainBike extends Bicycle {

}

// synek sobie dziedziczy po starym jego cechy, czyli pola i metody i oczywiscie moze sobie dodac swoje cechy


public class MountainBike extends Bicycle implements BikeInterface {
	// synek dziedziczy i implementuje interfejs
}
```

* Deklaracja klasy ma nastepujace komponenty, w kolejnosci:

	* Modyfikatory (Modifiers) tj. **public / private**, prywatna klasa moze byc zadeklarowana tylko w przypadku jak jest nested

	* Nazwa klasy, zawsze capital letter

	* dziedziczenie klasy **extend**, _Klasa moze dziedziczyc tylko jedna klase_

	* korzystanie z interfejsow, **implements**, klasa moze korzystac z wielu interfejsow, oddzielane przecinkiem

	* body klasy, pomiedzy curly braces { jakies body se tu jes }

* Deklaracja zmiennych wewnetrznych

	* Member variables (zmienne wew.) sa znane jako **fields** (pola)

	* Zmienne w metodzie klasy, to **local variables** (lokalne zmienne)

	* Zmienne w deklaracji metody - parametry

	* zmienna klasowa z modyfikatorem **public** oznacza, ze instancja obiektu ma do niej direct access, a jak modyfikator **private** to nie ma dostepu (wtedy tylko klasa ma) [enkapsulacja support eses],jak damy modyfikator **private** to mozemy dostac sie do tej zmiennej poprzez tzw. getter, to jest metoda ktora zwraca nam ten field dzieki temu podejsciu uniemozliwiamy zmiane tego pola - ale i to jest mozliwe dzieki setterowi, to kolejna metoda ktora modyfikuje to pole XD, na poczatku mozna pomyslec ze ktos kto to wymyslil byl napierdolony, ale by far to zwieksza czytelnosc

	* Kazda zmienna musi miec typ

	* Konwencja nazewnictwa to: **Pierwszy znak klasy powinien byc z duzej litery** oraz **pierwsze slowo opisujace metode powinno byc czasownikiem**


* Deklaracja metod

	* Tutaj tez mozna uzywac modyfikatorow, jak wyzej

	* return type moze byc void

	* Metody maja swoja sygnature, np. metoda public double calculate(double value1, double value2, int value3, float value4), to sygnatura jest nast. **calculate(double, double, int, float)**, jasne? jasne
	* Konwencja nazewnictwa to: **camelCase, pierwsze slowo czasownik, next przymiotnik, rzeczownik**

	* ATENZIONE, metoda moze byc overloaded, czyli moze byc zadeklarowanych z inicjalizacja kilka metod, ktore maja rozne typy w parametrach lub rozna liczbe parametrow, co wiecej - sygnatury tez maja inne, powinny byc uzywane z uwaga, bo moga zmniejszyc czytelnosc, jak np. tutaj [przeladowaniec](https://docs.oracle.com/en/java/javase/20/docs/api/java.base/java/util/Arrays.html)

* Konstruktor

	* Konstruktor, wyglada jak zwykla metoda, ale jest ekscytujaco niezwykla, bo to ona odpowiada zastworzenie naszej instancji
```java
public Bicycle(int startCadence, int startSpeed, int startGear) {
    gear = startGear;
    cadence = startCadence;
    speed = startSpeed;
} // to jest konstruktorek

//a to jest konsumpcja konstruktora ktora tworzy instancje

Bicycle myBike = new Bicycle(30, 0, 8);
```

	* Konstruktorkow moze byc wiecej niz 1, moze byc nawet bezargumentowy
```java
public Bicycle() {
    gear = 1;
    cadence = 10;
    speed = 0;
}
// takie konstruktorki maja swoje ograniczenia, tj. overloaded metody czyli nie moze byc dwoch takich samych konstruktorow, musza miec rozne zmienne w parametrach, lub roznic sie w liczbie tych zmiennych
```

	* Klasa nie musi miec konstruktora, ale **ATENZIONE** wtedy ta klasa pzryjmuje default konstruktor, czyli taki bezargumentowy, ktory calluje super klase i sprawdza czy tam nie ma zadeklarowanego bezargumentowego konstruktora, ale jesli ojciec nie ma konstruktora, to wtedy twoja klasa przyjmuje bezargumentowy konstruktor zadeklarowany w klasie Object **Klasa Object jest defualtowo ojcem wszystkich klas**

	* Mozna uzywac modyfikatorow w deklaracji kosntruktora, aby kontrolowac kto moze uzywac danego konstruktora
	

# Related to: 

* https://dev.java/learn/classes-objects/

#Java #Class #Constructor #Method #Variable #Modifiers
