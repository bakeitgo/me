# Java

## Overrided Instance methods

* Mozemy tworzyc nadpisane metody superklasy, pod warunkiem ze metoda zwraca ten sam typ, lub subtype (zwane covariant return type'm), ma taka sama nazwe oraz przyjmuje takie same parametry. Wtedy mamy do czynienia z metoda @override.

* Mozemy stosowac adnotacje @Override, dzieki ktorej w przypadku gdy w superklasie nie bedzie takiej metody, kompilator zwroci nam error.

## Overrided Static Methods

* W przypadku nadpisanych statycznych metod, mamy doczynienia z **ukrywaniem** a nie z **nadpisywaniem**. Statyczna metoda ktora ma taka sama sygnature(tj. nazwa, zwracany typ, liczba paramsow), jest ukrywana nizeli nadpisywana. To moze prowadzic do roznych implikacji, np. w przypadku gdy stworzymy instancje subklasy deklarujac jej typ jako parent klasa, np. Animal animal = new Cat(); sprawi ze **odwolanie do statycznej metody Animal, uzyje metody ktora jest w klasie Animal**, natomiast metoda **instancji zostanie wywolana ta ktora jest napisana w klasie Cat, a nie *Animal*** 

## Interface methods

* Metody o oznaczeniu **default** oraz **abstract** sa zawsze nadpisywane przez metode **instancji** klasy.
```java
public class Horse {
    public String identifyMyself() {
        return "I am a horse.";
    }
}

public interface Flyer {
    default public String identifyMyself() {
        return "I am able to fly.";
    }
}

public interface Mythical {
    default public String identifyMyself() {
        return "I am a mythical creature.";
    }
}

public class Pegasus extends Horse implements Flyer, Mythical {
    public static void main(String... args) {
        Pegasus myApp = new Pegasus();
        System.out.println(myApp.identifyMyself());
    }
}

// na outpucie pojawi sie "I am a horse", w koncu metoda instancji ma najwieksze prio ;)
```

* Subinterfejsy (interfejs ktory ma ojca interfejsa [tak to se nazwalem xd]) zawsze nadpisuja implementacje parenta, wiec gdy uzyjemy metody ktora pochodzi z zaimplementowanego interfejsu w naszej klasie, ktorej to interfejsem jest subinterfejs, to zawsze ta metoda z subinterfejsu bedzie wykorzystywana, nie ta z parent interfejsu
```java
public interface Animal {
    default public String identifyMyself() {
        return "I am an animal.";
    }
}

public interface EggLayer extends Animal {
    default public String identifyMyself() {
        return "I am able to lay eggs.";
    }
}

public interface FireBreather extends Animal { }

public class Dragon implements EggLayer, FireBreather {
    public static void main (String... args) {
        Dragon myApp = new Dragon();
        System.out.println(myApp.identifyMyself());
    }
}
```

* W momencie gdy uzyjemy dwoch **NIE POWIAZANYCH** ze soba interfejsow, ktore maja takie same sygnatury metod, to w przypadku zaimplementowania takich interfejsow przez klase, kompilator zwroci nam error. Aby temu zapobiec, trzeba zewnetrznie nadpisac taka metode, poprzez uzycie keyworda **super**, dokladnie to tak **nazwa_interfejsu.super.nazwa_metody**. Ponizej przyklad:
```java
public interface OperateCar {
    // ...
    default public int startEngine(EncryptedKey key) {
        // Implementation
    }
}

public interface FlyCar {
    // ...
    default public int startEngine(EncryptedKey key) {
        // Implementation
    }
}

public class FlyingCar implements OperateCar, FlyCar {
    // ...
    public int startEngine(EncryptedKey key) {
        FlyCar.super.startEngine(key); //tutaj mamy ten explicit override takiej metody
        OperateCar.super.startEngine(key);
    }
}
```

* Odziedziczone przez klase abstrakcyjne metody interfejsu, sa przez nia nadpisywane. Ponizej przyklad:

```java
public interface Mammal {
    String identifyMyself();
}

public class Horse {
    public String identifyMyself() {
        return "I am a horse.";
    }
}

public class Mustang extends Horse implements Mammal {
    public static void main(String... args) {
        Mustang myApp = new Mustang();
        System.out.println(myApp.identifyMyself());
    }
}
// na outpucie pojawia sie "I am a horse."
```

## Overrided method Modifiers

* W przpyadku modyfikatorow, odziedziczona metoda moze zdefiniowac modyfikator o mniejszej restrykcji, nigdy wiekszej, np. gdy odziedziczona metoda jest **protected**, subklasa nie moze zrobic z niej **private** ani package private, jedynie **public** i **protected**. Przy probie dokonania takiej akcji, kompilator zwroci error

## Summary:

| | Superclass Instance Method | Superclass Static Method |
| --- | --- | --- |
| Subclass Instance Method | Overrides | Generates a compile-time error |
| Subclass Static Method | Generates a compile error | Hides |


# Related to: 

* https://dev.java/learn/inheritance/overriding/


#Java #Override #instance #static #interface #modifiers
