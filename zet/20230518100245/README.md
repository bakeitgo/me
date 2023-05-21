# Java

## Creating and using objects

### What objects are?

* Czym sa obiekty? Obiekty to instancje klasy, konkretne byty. Sa tworzone na podstawie klas, a klasa to po prostu taki template obiektu (blueprint). Jak tworzone sa obiekty?

```java
// Jak tworzone sa obiekty?

public class Main {
	class Point {}
	public static void main() {
		Point point = new Point(20, 30);	
	}

}
```

* W jaki sposob tworzone sa obiekty? 3 etapy:

	* Deklaracja - **Point point**, tutaj okreslamy typ referencyjny oraz nazywamy zmienna **point**
	* Instantacja - slowo kluczowe **new** sluzy za tworzenie obiektu (czyli instancji)

	* Initializacja - odwolanie sie do konstruktora klasy **Point(20, 30)**

* Sam etap deklaracji obiektu jest nie mozliwy, trzeba zainicjalizowac obiekt, inaczej dostajemy **compile error**

* Zamiast tworzyc caly obiekt mozna odwolac sie do konkretnego pola / metody obiektu przy jego inicjalizacji oraz przypisac jego wartosc do okreslonej zmiennej. Takie podejscie  sprawia ze po odwolaniu sie do konkretnego pola instancji obiektu, nie potrzebujemy juz calego obiektu wiec JVM oznacza ten obiekt jako nieosiagalny, przez co pamiec ktora zostala zaalokowana do tego obiektu moze zostac usunieta przez **garbage collector**

```java
int x = new Point().x;
```

### Garbage collector

* Garbage collector odpowiada za czyszczenie pamieci, usuwa z pamieci obiekty ktore sa juz nieosiagalne, czyli gdy referencje do tych obiektow wychodza poza scope programu, zmienna ktora byla referencja zostanie nadpisana, lub przypiszemy do zmiennej ktoar byla referencja wartosc **null**

# Related to:

* https://dev.java/learn/classes-objects/creating-objects/


#java #garbage #garbagecollector #objects #constructors #instantation #initialization #declaration
