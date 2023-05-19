# Java

## Strings

* string to nic innego a sekwencja znakow xd

* string jest wyjatkowa klasa, bo jest klasa a zachowuje sie jak prymityw, jak podamy bezposrednio w kodize wartosc literalna np. "hello" to wtedy kompilator sam za nas stworzy instancje Stringa z ta wartoscia xd (pieklo stringa)

* jak kazda inna kalse mozemy stworzyc instnacje stringa uzywajac slowa new etc. co wiecej stringa mozemy stworzyc z tablicy znakow, po prostu podajac ja do argumentu konstruktora stringa xd

* **ATENZIONE** klasa *String* jest immutable tak samo jak *Character*, a jednak **String** ma metody ktore umozliwaija te zmiany, w rzeczywistosci nie modyfikuja stringa same w sobie tylko kopiuja go, podmieniaja i zwracaja zmieniona wartosc xD

* podobnie jak na sout mozna decydowac o formacie outputu to w stringu tez mozna to robic, jest taka metoda format, ktora wyglada identycznie jak printf

* Mozemy pobrac okreslony znak uzywajac indeksu, bo stringi sa index-based

## String Builders

* ta klase uzywamy gdy musimy polaczycc ze soba duza liczbe stringow, instancja tej klasy to nie jest zwykly string tylko tablica ktora zawiera ciag znakow

* string builder oprocz length ma jeszcze capacity, z reguly cap ma wiekszy niz length

* gdy appendujemy do stringa innego stringa i length bedzie wieksze od cap, to cap od razu ulega powiekszeniu

* jest jeszcze klasa **StringBuffer** ktora jest odpowiednikiem klasy **StringBuilder** ale ona jest oznaczona jako *thread-safe*, bo ma metody ktore sa zsynchronizowane.


## Autoboxing

* Jest to automatyczna konwersja ktora java robi w locie z prymitywa na referencyjny typ

## Unboxing

* to samo co autoboxing ale w druga strone xd

# Related to:

* https://dev.java/learn/numbers-strings/strings/

#Java #Strings #String Builders #String Buffers #Autoboxing #unboxing
