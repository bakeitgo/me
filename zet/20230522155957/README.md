# Java

## Abstract methods and classes

* Zamiast interfejsu mozna stworzyc abstrakcyjna klase, ma ono podobne lecz bardziej rozszerzone zastosowanie niz interfejs, w klasie abstrakcyjnej mozemy definiowac metody / pola ktore maja modyfikatory inne niz **public static final** (bo takie wlasnie sa w interfejsie). Przykladowe zastosowanie takiej klasy to np. *generyczne(ogolne)* polaczenie klas, np. abstrakcyjna klasa Vehicle, laczy w sobie wszystkie klasy *Bicycle, Car, Scooter*. Taka klase definiujemy uzywajac modyfikatora *abstract*. W takiej klasie mozna definiowac metody ktore sa abstrakcyjne poprzez uzycie modyfikatora *abstract* oraz bez takiego modyfikatora. Metody z takim modyfikatorem nie maja bloku (implementacji) natomiast te drugie juz maja.

* Jesli chodzi o uzycie takiej abstrakcyjnej klasy, to restrykcje sa tutaj takie same jak w zwyklych klasach, subklasa moze miec jednego parenta

* Dodatkowe zastosowanie takeij abstrakcyjnej klasy, jest takie ze taka abstrakcyjna klasa moze stanowic pewnego rodzaju bypass dla interfejsow, dzieki czemu nie musimy implementowac wszystkich metod zawartych w interfejsie, a tylko te ktore nas interesuja. Krocej mowiac, abstrakcyjna klasa nie musi implementowac wszystkich metod interfejsu,a tylko te ktore developer chce. Ponizej zastosowanie:

```java
abstract class X implements Y {
  // implements all but one method of Y
}

class XX extends X {
  // implements the remaining method in Y
}
```

* Abstrakcyjna klasa moze miec pola / metody **static**, do ktorych mozna bezposrednio sie odwolywac.

## Polymorphism

* Polymorphism - od lacinskiego slowa (poly - wiele, morph - form), glownie chodzi o to ze subklasa moze uzywac metod / pol superklasy przy uzyciu slowa super, hiduje jej pola gdy zadeklaruje je u siebie

* do metod parent klasy w subklasie odwolujemy sie poprzez (.) np. *super.jakasMetoda()*

* w konstruktorze subklasy mozna uzyc *super()*, wtedy uzywamy bezargumentowego konstruktora parenta, lub *super(jakies sobie tam paramsy)* wtedy uzywamy okreslonego konstruktora parenta

* W parencie mozna zainicjalizowac **stala metode ktora nie moze zostac nadpisana przez subklasy** poprzez uzycie modyfikatora **final**

# Related to:

* https://dev.java/learn/inheritance/abstract-classes/

#Java #abstract #interface #polymorphism
