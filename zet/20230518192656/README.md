# Java

## Nested Classes

* Wyrozniamy dwa typy nested klas, static nested klase i non-static nested klase. 
	
	* **Non-static Nested class** jest nazywana **Inner Class**

	* **Static nested class** jest nazywana **nested class**

* inner klasa ma dostep do pol / metod outer klasy, nawet jesli sa one zadeklarowane z modyfikatorem **private**, nawet jesli sa **static**

* nested klasa nie ma dostepu do zadnych pol / metod outer klasy

* **Outer klasa** moze byc zadeklarowana jako **public lub package-private (non-modifier)**

* Dlaczego uzywamy nested  / inner klas?

	* Jesli klasa jest uzyteczna w jednym miejscu, to jest to dobry moment aby stworzyc z niej nested / inner klase, dzieki temu nasze paczki kawowskie sa bardziej czytelne.

	* Poprawia enkapsulacje, dzieki temu ze inner klasa ma dostep nawet do prywatnych metod / pol outer klasy

	* zwieksza czytelnosc, utrzymywanie kodu. nestowanie klas sprawia ze kod bardziej odwzorowywuje co chcemy osiagnac

	* Aby uzyc inner klasy w innym miejscu niz w outer klasie, najpierw trzeba zainicjalizowac outer klase, nastepnie z niej inner klase 

	* **ATENZIONE** inner klasa nie moze posiadac pol / metod **static**, w koncu mozemy sie z nia komunikowac tylko i wylacznie poprzez instancje outer klasy

```java
OuterClass outerObject = new OuterClass();
OuterClass.InnerClass innerObject = outerObject.new InnerClass();
```

* **Static nested klasa** - nie moze odwolywac sie do pol klasy innych niz static, analogicznie jak **static** metody zwyklej klasy, ale moze odwolywac sie do nich poprzez stworzona instancje, a to juz sie da zrobic.

```java
public class OuterClass {

    String outerField = "Outer field";
    static String staticOuterField = "Static outer field";

    class InnerClass {
        void accessMembers() {
            System.out.println(outerField);
            System.out.println(staticOuterField);
        }
    }

    static class StaticNestedClass {
        void accessMembers(OuterClass outer) {
            // Compiler error: Cannot make a static reference to the non-static
            //     field outerField
            // System.out.println(outerField);
            System.out.println(outer.outerField);
            System.out.println(staticOuterField);
        }
    }

    public static void main(String[] args) {
        System.out.println("Inner class:");
        System.out.println("------------");
        OuterClass outerObject = new OuterClass();
        OuterClass.InnerClass innerObject = outerObject.new InnerClass();
        innerObject.accessMembers();

        System.out.println("\nStatic nested class:");
        System.out.println("--------------------");
        StaticNestedClass staticNestedObject = new StaticNestedClass();
        staticNestedObject.accessMembers(outerObject);

        System.out.println("\nTop-level class:");
        System.out.println("--------------------");
        TopLevelClass topLevelObject = new TopLevelClass();
        topLevelObject.accessMembers(outerObject);
    }
}
```

* Jak dziala this? this dziala w zakresie klasy w ktorej wystepuje, this nie moze byc w statycznej klasie wiec thisa tam nie uzyjemy, ale w innerclasie juz moize byc wtedy odwoluje sie do inner klasy. tak samo deklaracja zmiennej o takiej samej nazwie jak zmienna w outer classie jest zshadowowana. Mozna sie odwolac do konkretnego pola outer klasy poprzez taki syntax. **NAZWAKLASY.this.NAZWAPOLA**

* Serializacja klas w ktorych sa inner klasy jest mocno odradzana

* Przyklad uzycia inner klasy

```java
public class DataStructure {

    // Create an array
    private final static int SIZE = 15;
    private int[] arrayOfInts = new int[SIZE];

    public DataStructure() {
        // fill the array with ascending integer values
        for (int i = 0; i < SIZE; i++) {
            arrayOfInts[i] = i;
        }
    }

    public void printEven() {

        // Print out values of even indices of the array
        DataStructureIterator iterator = this.new EvenIterator();
        while (iterator.hasNext()) {
            System.out.print(iterator.next() + " ");
        }
        System.out.println();
    }

    interface DataStructureIterator extends java.util.Iterator<Integer> { }

    // Inner class implements the DataStructureIterator interface,
    // which extends the Iterator<Integer> interface

    private class EvenIterator implements DataStructureIterator {

        // Start stepping through the array from the beginning
        private int nextIndex = 0;

        public boolean hasNext() {

            // Check if the current element is the last in the array
            return (nextIndex <= SIZE - 1);
        }

        public Integer next() {

            // Record a value of an even index of the array
            Integer retValue = Integer.valueOf(arrayOfInts[nextIndex]);

            // Get the next even element
            nextIndex += 2;
            return retValue;
        }
    }

    public static void main(String s[]) {

        // Fill the array with integer values and print out only
        // values of even indices
        DataStructure ds = new DataStructure();
        ds.printEven();
    }
}
```

### Local and Anonymous Classes

* Prezentacja dwoch roznych inner klas, **local** klase mozna zadeklarowac przy uzyciu body metody, natomiast **anonymous** to nic innego jak po prostu inner klasa ale bez nazwy

#### Local class

* Deklarowanie klasy lokalnej moze nastapic w kazdym bloku **Expressions, Statements, Blocks**

* Klasa lokalna ma dostep do lokalnych zmiennych TYLKO jak sa one **final** (constant) lub **effective final** czyli taka do ktorej nie jest przypisywana w ktoryms miejscu w kodzie jakas wartosc. Klasa lokalna ma rowniez dostep do zmiennych podanych w parametrze enclosing metody (od java 8)

* Deklarowanie zmiennej k tora ma taka sama nazwe w parametrze jest shadowowana

* W lokalnych klasach nie mozna definiowac pol **static**, dlatego nie mozna definiowac rowniez interface (interface jest wlasciwie **static**), tak samo w enclosing metodach lokalnej klasy. **CHYBA ZE TO POLE STATIC JEST STALE, CZYLI JEST FINAL WTEDY SIE DA**

* Zmienna stala jest zmienna ktora ma prymitywny typ badz typ String, i jest bezposrednio przypisana lub obliczana przy compile time

#### Anonymous Classes

* Anonimowa klasa jest podobna do lokalnej klasy, z tym ze anonimowa klasa jest brana jako wyrazenie. Wyobraz sobie ze masz interfejs HelloWorld() i zamiast deklarowac np. w metodzie klaseclass nazwa implements esese { jakas implementacja } to inicjalizujesz ten interfejs tak jak bys chcial stworzyc obiekt HelloWorld siema = new HelloWorld() - po tej inicjalizacji dajesz curly braces i definiujesz wszystko tak samo jak w klasie lokalnej xD przyklad n izje

```java
public class HelloWorldAnonymousClasses {

    interface HelloWorld {
        public void greet();
        public void greetSomeone(String someone);
    }

    public void sayHello() {

        class EnglishGreeting implements HelloWorld {
            String name = "world";
            public void greet() {
                greetSomeone("world");
            }
            public void greetSomeone(String someone) {
                name = someone;
                System.out.println("Hello " + name);
            }
        }

        HelloWorld englishGreeting = new EnglishGreeting();

        HelloWorld frenchGreeting = new HelloWorld() {
            String name = "tout le monde";
            public void greet() {
                greetSomeone("tout le monde");
            }
            public void greetSomeone(String someone) {
                name = someone;
                System.out.println("Salut " + name);
            }
        };

        HelloWorld spanishGreeting = new HelloWorld() {
            String name = "mundo";
            public void greet() {
                greetSomeone("mundo");
            }
            public void greetSomeone(String someone) {
                name = someone;
                System.out.println("Hola, " + name);
            }
        };
        englishGreeting.greet();
        frenchGreeting.greetSomeone("Fred");
        spanishGreeting.greet();
    }

    public static void main(String... args) {
        HelloWorldAnonymousClasses myApp =
            new HelloWorldAnonymousClasses();
        myApp.sayHello();
    }
}
```

* W anonimowej klasie statements nie sa dopuszczalne

* Anonimowa klasa ma dostep do zmiennych ktore sa **final** lub **effective final**, ma dostepdo zmiennych klasy wyzej. deklaracja zmiennej shadowuje zmienne enclosing scope, moze miec pola **static** w przypadku jak sa one **final**. Nie da sie tworzyc konstruktorow w takiej klasie,


# Related to:

* https://dev.java/learn/classes-objects/nested-classes/



#java #nested #class #local class #anonymous class #inner class #nested class #scope
