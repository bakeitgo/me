# Java

## Numbers

* uzywamy **wrapperow** prymitywow aby konwertowac wartosci pomiedzy roznymi typami, zarowno prymitywnymi jak i referencyjnymi, aby uzywac stalych np. MIN\_VALUE, MAX\_VALUE, jako argument metody ktory wymaga objektu, np. gdy manipulujemy kolekcja numerow

* Mozna konwertowac z referencyjengo wrappera do prymitywu

	* byte byteValue()

	* short shortValue()

	* int intValue()

	* long longValue()

	* float floatValue()

	* double doubleValue()

* jest tez metoda to komparowania int compareTo(Number wrapper xx)

* boolean equals(Object obj); -- ta metoda zwraca okreslony bool gdy podany argumetr jest w tym samym typie oraz ma taka sama numeryczna wartosc. Dla zmiennoprzecinkowych **Double** i **Float** potrzebne sa wieksze wymagania, wszystko opisane w API javy

* Kazda podklasa numerowa posiada swoje metody do konwertowania z stringa i przekonwertowywania z innych systemow numerowych, **np. short do int**

* Warto zaznaczyc ze nie musimy konwertowac wrapperow do prymitywow, bo java robi to za nas przy uzyciu autoboxingu i unboxingu (te dwa kocnepty odpowiadaja za przekonwertowywanie z wrappera od inta i wicewersa)

## Printf and Format methods

* java.io pckg ma klase PrintStream ktora odpowiada za instancje znanego **System.out**, obsluguje takie metody jak format printf println etc.

* formattery sa opisane w **java.util.Formatter**

```java
import java.util.Calendar;
import java.util.Locale;

public class TestFormat {
    
    public static void main(String[] args) {
      long n = 461012;
      System.out.format("%d%n", n);      //  -->  "461012"
      System.out.format("%08d%n", n);    //  -->  "00461012"
      System.out.format("%+8d%n", n);    //  -->  " +461012"
      System.out.format("%,8d%n", n);    // -->  " 461,012"
      System.out.format("%+,8d%n%n", n); //  -->  "+461,012"
      
      double pi = Math.PI;

      System.out.format("%f%n", pi);       // -->  "3.141593"
      System.out.format("%.3f%n", pi);     // -->  "3.142"
      System.out.format("%10.3f%n", pi);   // -->  "     3.142"
      System.out.format("%-10.3f%n", pi);  // -->  "3.142"
      System.out.format(Locale.FRANCE,
                        "%-10.4f%n%n", pi); // -->  "3,1416"

      Calendar c = Calendar.getInstance();
      System.out.format("%tB %te, %tY%n", c, c, c); // -->  "May 29, 2006"

      System.out.format("%tl:%tM %tp%n", c, c, c);  // -->  "2:34 am"

      System.out.format("%tD%n", c);    // -->  "05/29/06"
    }
}
// o to przykladowe formattery w prezyklaach
```

## DecimalFormat Class

* ta klasa pozwala nam na okreslnie formatu wszystkich zer prefixow, suffixow grupowania separatorow etc.

```java
import java.text.*;

public class DecimalFormatDemo {

   static public void customFormat(String pattern, double value ) {
      DecimalFormat myFormatter = new DecimalFormat(pattern);
      String output = myFormatter.format(value);
      System.out.println(value + "  " + pattern + "  " + output);
   }

   static public void main(String[] args) {

      customFormat("###,###.###", 123456.789);
      customFormat("###.##", 123456.789);
      customFormat("000000.000", 123.78);
      customFormat("$###,###.###", 12345.67);  
   }
}
```

## Math class

* Ta klasa jak sama nazwa wskazuje sluzy do operacji arytmetycznych

* JEST TAKI KONCEPT *STATIC IMPORT LANGUAGE FEATURE* TO **TAKI GAMECHANGER DZIEKI TKOREMU MOZEMY UZYWAC STATYCZNYCH METOD JAK ZAIMPORTUJEMY PACZKE W NASTEPUJACY SPOSOB: import static java.lang.Math.(wildcard); wtesdy nie musimy uzywac metod tak Math.cos, tylko po prostu cos**

* Ta klasa ma dwie stale

	* **E** (bazowa naturalnych logarytmow)

	* **PI** stos. obw. kola do jego srednicy

```java
public class BasicMathDemo {
    public static void main(String[] args) {
        double a = -191.635;
        double b = 43.74;
        int c = 16, d = 45;

        System.out.printf("The absolute value " + "of %.3f is %.3f%n", 
                          a, Math.abs(a));

        System.out.printf("The ceiling of " + "%.2f is %.0f%n", 
                          b, Math.ceil(b));

        System.out.printf("The floor of " + "%.2f is %.0f%n", 
                          b, Math.floor(b));

        System.out.printf("The rint of %.2f " + "is %.0f%n", 
                          b, Math.rint(b));

        System.out.printf("The max of %d and " + "%d is %d%n",
                          c, d, Math.max(c, d));

        System.out.printf("The min of of %d " + "and %d is %d%n",
                          c, d, Math.min(c, d));
    }
}
```

## Random numbers

* Mozemy generowac liczby pseudoloswe uzywajac metody Math.random()

```java
int number = (int)(Math.random() * 10);
```

* **ATENZIONE** jesli musimy wygenerowac wieksza ilosc liczb randomowych, zaleca sie uzycia instancji *java.util.Random* i uzycia metod tej instancji

# Related to:

* https://dev.java/learn/numbers-strings/numbers/


#Java #random number #math class #numbers #wrappers #printf #format methods
