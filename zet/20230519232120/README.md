# Java

## Inheritance

* Dziedziczenie - 

	* Klasa ktora dziedziczy po innej klasie nazywana jest subklasa / child klasa / derived class / extended class

	* Klasa ktora jest ojcem nazywana jest superclass / base class / parent class

	* Klasy moga tworzyc drzewo klas, czyli miec nieskonczona ilosc parent class

	* Klasa ktora dziedziczy moze uzywac pol i metod superklasy

	* Klasa ktora dziedziczy pobiera wszystkie pola metody superklasy, oprocz konstruktorow, ale konstruktor moze byc wywolany przez subklase

	* Na samym przedzie hierarchi jest klasa Object, w paczce java.lang i te klasy ktore my tworzymy

```java
public class Bicycle {
        
    // the Bicycle class has three fields
    public int cadence;
    public int gear;
    public int speed;
        
    // the Bicycle class has one constructor
    public Bicycle(int startCadence, int startSpeed, int startGear) {
        gear = startGear;
        cadence = startCadence;
        speed = startSpeed;
    }
        
    // the Bicycle class has four methods
    public void setCadence(int newValue) {
        cadence = newValue;
    }
        
    public void setGear(int newValue) {
        gear = newValue;
    }
        
    public void applyBrake(int decrement) {
        speed -= decrement;
    }
        
    public void speedUp(int increment) {
        speed += increment;
    }
}

// inna klasa  - inny plik


public class MountainBike extends Bicycle {

    // the MountainBike subclass adds one field
    public int seatHeight;

    // the MountainBike subclass has one constructor
    public MountainBike(int startHeight,
                        int startCadence,
                        int startSpeed,
                        int startGear) {
        super(startCadence, startSpeed, startGear); // pobranie i nadpisanie pol ojca
        seatHeight = startHeight;
    }

    // the MountainBike subclass adds one method
    public void setHeight(int newValue) {
        seatHeight = newValue;
    }
}
```

* Subklasa inherituje wszystkie pola ktore sa oznaczone modyfikatorami **protected** i **public**, jesli subklasa jest w tej samej paczce co parent klasa, inherituje tez od niej **package-private** memberow

* Odradza sie deklarowania tych samych zmiennych w subklasach (ich shadowing)

* W subklasach mozna overridowac zainheritowane metody

* W subklasie mozna napisac nowa **static** metode ktora jest w parencie, shadowuje ona metode parenta

## Private members in a superclass

* jak wiadomo subklasa nie inherituje prywatnych memberow superklasy, ale jak superlasa ma nested klase, to ta nested klasa ma juz dostep do prywatnych memberoww, a jak taka nested klasa jest protected / private albo pckg private (ale wtedy subklasa musi byc w tej samej paczce) to subklasa inherituje takowa klase i moze miec dostep do tych pol z tej wlasnie klasy xd

## Casting objects

* Obiekty moga byc skastowane, zlaozmy ze mamy nast. drzewo klas, MountainBike - dzieciak Bicycle - dzieciak Object. Przy takim polaczeniu, mozemy zainicjalizowac zmienna w nast. sposob: **Object object = new MountainBike();** taka operacja nazwya sie **implicit casting**. warto zaznaczyc ze nie dziala to w druga strone. W druga strone jakbysmy chcieli przypisac do zmiennej musimy zrobic **explicit casting** w koncu kompilator nie wie ze ojciec moze miec taki sam typ jak jego syn (brak polaczenia bezposredniego). przy explicit castingu kompilator nam zaufa ze wiemy co robimy, oto przykald: **MountainBike bike = (MountainBike)object;**, mozemy zrobic logical test ktory nie wywali nam exceptiona gdybysmy przypadkiem sie pomylili, oto on: **if (obj instanceof MountainBike) { MountainBike myBike = (MountainBike)object}**

# Related to: 

* https://dev.java/learn/inheritance/



#Java #inheritance #subclass #superclass #super #implicit cast #explicit cast #private
