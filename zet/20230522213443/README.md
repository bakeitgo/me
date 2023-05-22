# Java

## Generics

### Generic classes

* Dlaczego generics powstalo? Wyobraz sobie taki przypadek, dla kazdej klasy ktora implementujesz, implementujesz te same metody, pola jedyne cyzm sie roznia od siebie to typem, i robisz klasa po klasie, aby obsluzyc kazdy typ. Przez to aby obsluzyc 10 typow trzeba zrobic 10 klas, niezbyt czytelny syf. Dzieki temu powstaly generyki ktore umozliwiaja zdefiniowanie *generycznego* typu. Ale okej, mozna stwierdzic ze  robienie to jak wyzej to glupota, przeciez mozna stworzyc po prostu klase ktora jest o typie Object, wtedy wszystko tam wepchniemy. No tak, tylko co jak mamy np. tablice, wpychamy tam obiekty o typie Cat, a nagle przychodzi programista wariacik i sobie wrzuca tam obiekt o innym typie, wyobrazasz sobie ze w runtime probujesz sie odwolac do niego myslac ze to kot a to pies? i co wtedy, runtime error, wez to debuguj chlopie, tak to masz generyczki i masz esse bo jest type safety.

* Typy mozna rozszerzac, wtedy mowimy o **Bounded Generic**, jest to klasa ktora przyjmuje typ ktory jest rozszerzony, np public class Printer <T extends Animal> { }. Przy tego typu generyku, jestesmy w stanie bezposrednio odwolywac sie poprzez instancje T do metod / pol zdefiniowanych w klasie rozszerzajacej nasz typ, czyli w tym wypadku **Animal** *MOZNA TEGO UZYWAC TEZ INTERFEJSAMI, TEZ Z UZYCIEM SLOWA extends* 

* Typy mozna rozszerzac o *wiele typow*, mowimy wtedy o **Multiple Bounds**, jego syntax wyglada tak: **public class Printer <T extends Animal & SomeInterface>**, *Przy tym syntaxie gdy klasa jest rozszerzana i o klase i o dowolna liste interfejsow, musimy pamietac o tym ze klasa musi zawsze znalezc sie pierwsza w tym chainie, oraz taka klase mozemy rozszerzyc zawsze tylko o 1 klase*

### Generic methods

* Generyczne metody to takie metody ktore przyjmuja jako parametry generyczne typy, oraz moga zwrocic generyczny typ, ich syntax wyglada np. tak: (nie zwracaj uwage na modifiery)

```java
private static <T,V> void (T param1, V param2) {
	// implementacja

	//ta generyczna metoda przyjmuje dwa typy generyczne, T oraz V
}

private static <T,V> T (T param1, V param2) {
	// implementacja

	// tutaj widzimy podobna metode jak wyuzej, tylko ze w tym przypadku metoda zwraca typ T
}
```


### Wildcards

* W przypadku gdy nie mamy bladego pojecia jaki typ wrzucimy do paramsow powinnismy miec mozliwosc wrzucenia unknown typu, i tutaj wlasnie przychodzi wildcard, wildcard moze byc rozszerzony o okreslona klase / interfejs (podobnie jak typee parameter

```java
private static void (List<?> someList) {
	// wildcard instead of Type param
}
```

* Wildcard moze zostac rozszerzony tak samo jak type parameter, taki sam syntax

# Related to:



#Java #Generics #wildcards #generic methods #generic classes
