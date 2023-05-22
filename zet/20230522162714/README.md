# Java

## Interfaces

* Interfejsy sa pewnego rodzaju kontraktem, mowiacym o tym jakie metody, pola chcemy miec w naszym programie. Interfejsy pomagaja nam w lepszym rozumieniu oprogramowania, poprzez sygnatury metod wiemy jakie metody chcemy zrobic i jakie metody ktos nam dostarcza

* interfejsy moga zostac rozszerzone o inne interfejsy, robi sie to tak samo jak przy klasach

* interfejsy moga zawierac, abstract, default i static metody (abstract metoda to metoda ktora nie zawiera implementacji) sa one **defaultowo public**, wiec mozemy ominac modyfikator

* interfejs moze zawierac stale, wszystkie stale sa **defaultowo public static final**

* gdy chcemy ewoluowac interfejs, po jego ewoluowaniu zawsze jakakolwiek klasa ktore implementuje ten interfejs, bedzie musiala dopisac implementacje tej metody. Jest to dosyc uciazliwe, dlatego pojawil sie koncept rozszerzania interfejsow, po prostu twozrymy nowy interfejs, dajemy extend tego poprzedniego i tyle. Jest jeszcze drugi sposob, stworzenie metody **default / static**, przy takich metodach poprzednie klasy nie musza byc nadpisywane.

### Default method

* metoda z modyfikatorem **default**, musi zawrzec od razu implementacje. Gdy **poszerzamy zakres swojego interfejsu o inny interfejs**, nie musimy deklarowac tej samej **default metody, dzieki temu dziedziczymy ta metode**.

* Redeklaracja tej metody, sprawia ze staje sie ona *abstract*

* Redefinicja metody **default**, **overrides** (nadpisuje) ja.

### Static methods

* W interfejsie mozemy tworzyc statyczne metody, co sprawia ze mozemy je trzymac w interfejsie a nie w klasie implementujaca ta metode



# Related to: 

* https://dev.java/learn/interfaces/defining-interfaces/


#java #static #default #interfaces
