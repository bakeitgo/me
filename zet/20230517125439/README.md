# Java

## Object

* Obiekt, to zbior stanow (fields) i zachowan (methods),  metody operuja na ich wewnetrznym stanie (ich polach). Poprzez metody ukrywamy stan obiektow, co niesie sie z konceptem enkapsulacji (ukrywanie wewnetrznej implementacji). Dzieki konceptowi obiektow zyskujemy reuzywalnosc, ukrywanie wew. impl., modularnosc (jeden obiekt moze byc uyzyty w innym obiekcie), latwosc w debugowaniu i podlaczalnosc (latwosc w zamienianiu jednego zepsutego obiektu na drugi, jak w mechanice, np. gdy wymieniamy opony letnie na zimowe)

## Class

* Obiekt wymieniony wyzej jest instancja klasy, klasa to nic innego jak template / schemat wg. ktorego konstruujemy obiekty

## Inheritance

* Inaczej dziedziczenie, dzieki temu konceptowi klasa moze moze dziedziczyc od rodzica jego zachowania i stany. Np. tworzymy bardziej ogolna klase **Bicycle**, i dziedziczymy po niej wszystkie zdefiniowane metody / stany dla szczegolnego przypadku, np. **Mountain Bike**, **Road Bike**, **Tandem Bike**, w ktorych to definiujemy specyficzne dla okreslonego przypadku stany / zachowania. 
> ```class MountainBike extends Bicycle```

## Interface

* Jest to interfejs ktory definiuje nam jakie metody musza znalezc sie w klasie.
> ```interface Bicycle {
	void metoda1(int wartosc);

	void metoda2(int wartosc);

	void metoda3(int wartosc);
}
public class ACMEBicycle implements Bicycle {
	...

	// Gdy metody okreslone w interfejsie nie znajda sie w klasie, kompilacja programu sie nie uda
	// nazwy zmiennych w parametrach tez musza byc takie same jak okreslone w interfejsie
}
```

## Package

* Zbior klas i interfejsow, JAVA platform oferuje taki zbior, znany jako **API**, _Application Programming Interface_




# Related to:

* https://dev.java/learn/oop/


#Java #Class #Object #Interface #Inheritance #Encapsulation
