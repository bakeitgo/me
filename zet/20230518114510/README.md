# Java

## Classes

### Classes return values

* Klasa w metodach moze zwracac jedynie swoj typ, czyli np. klasa Siema moze w metodzie zwracac typ Siema, oraz subklasy i interfejsy ktorych uzywa.

* subklasa nie moze zwracac parent klasy

### Static fields / methods

* Klasa moze miec metody / pola klasowe (static) oraz metody / pola (non static) instancji (obiektu). Jak skonsumowac takowa metode? uzywajac nazwy klasy oraz nazwy tej metody. **ATENZIONE** mozna odwolac sie do takowej statycznej metody rowniez z poziomu instancji klasy, ale jest to odradzana praktyka bo zmniejsza czytelnosc kodu.

* Metody klasowe nie moga korzystac z zmiennych instancji bezporsednio, musza odwolywac sie do nich poprzez instancje, oraz nie moga uzywac **this** (w koncu this odwoluje sie do instancji).

* Instancje klasy ktore maja static fieldy / metody moga ich uzywac, ale jest to odradzane (patrz wyzej)

### Initializing fields

* Powinnismy zawsze deklarowac zmienne na poczatku klasy - dobra praktyka

* Jesli chcemy zainicjalizowac pole static przy uzyciu logiki, mozemy uzyc do tego inicjalizatora bloku, taki blok bedzie zawsze wykonywany w miejscu gdzie zostal zainicjalizowany. Mozemy uzyc tez jego alternatywy, prywatnej statycznej metody, plusem takiego podejscia jest to ze takowa metode mozemy uzyc ponownie gdy chcemy jeszcze raz zainicjalizowac klasowa zmienna:

```java
public static publicznaStatyczna = initializeStaticMethod();

private static initializeStaticMethod() {

}
```

* **ATENZIONE** zadeklarowanego statycznego bloku nie da sie zredeklarowac oraz zatrzymac przed uruchamianiem. Jesli takowy statyczny  blok nie moze zostac wykonany przez program, to wtedy nie bedzie zadnej mozliwosci stworzenia instancji klasy zawierajacej static block. Moze sie to wytdarzyc gdy statyczny blok ma kod ktory odwoluje sie do **sieci** lub **zewnetrznego zrodla**

* Mozna tez inicjalizowac w podobny sposob pola instancji, w nast. sposob:

```java
{
 // to samo co przy static bloku tylko bez slowa static xD
}

private varType myVar = initializeInstanceVariable();

protected final varType initializeInstanceVariable() { } 

// drugie podejscie jest fajne w momencie gdy subklasy chca ponownie uzyc tej metody. Metoda jest final poniewaz metody initializacji pol instancji ktore nie sa stale moze prowadzic do problemow
```

### this keyword

* keyword **this** odnosi sie do odnoszenia sie do pol obiektu w klasie, mozna go uzywac w konstruktorze np. gdy mamy kilka konstruktorow mozemy uzyc **explicit constructor invocation** czyli odwolania sie do konstruktora klasy poprzez slowo this()

```java
public class Rectangle {
    private int x, y;
    private int width, height;
        
    public Rectangle() {
        this(0, 0, 1, 1);
    }
    public Rectangle(int width, int height) {
        this(0, 0, width, height);
    }
    public Rectangle(int x, int y, int width, int height) {
        this.x = x;
        this.y = y;
        this.width = width;
        this.height = height;
    }
    ...
}

```

### Control access Modifiers

* Wystepuja tu 3 slowa kluczowe i brak slowa hehe, public protected (non-modifier [brak]) private, kazda definiuje co moze siedostac do danej metody / pola.

| Modifier | Class | Package | Subclass | World |
| ------ | ------ | ------ |  ------ | ------ |
| public | Y | Y | Y | Y |
| protected | Y | Y | Y | X |
| non-modifier | Y | Y | X | X |
| private | Y | X | X | X |

* Powinnismy zawsze uzywac modifiera **private**, **public** zawsze trzeba uzywac przy stalych **final**

# Related to:

* https://dev.java/learn/classes-objects/more-on-classes/



#Java #Modifiers #Static #Initialization #static code block #static initialization methods #instance initialization methods #return #return type #subclass #this
