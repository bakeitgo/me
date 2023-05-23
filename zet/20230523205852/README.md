# Kotlin

## Data types

* W kotlinie mamy cos takiego jak type inference, nie muisimy deklarowac typu jak definiujemy zmienna, kompilator sam sobie go dedukuje

* W kotlinie zawsze musimy inicjalizowac zmeinna

* Nie mozemy mismatchowac typow, czyli nie mozemy przypisywac wartosci ktora wksazuje na inny typ do zmiennej

## Type conversion

* niektore funkcje wymagaja okreslonych typow, aby nie bylo type mismatcha mozemy przekonwertowac liczby uzywajac toDouble, toLong etc.

* **char** nie jest numeric tyypem ale mozemy przekonwertowac inta w char i wicewersa - toChar()

* **atenzione** w przypadku jak dana wartosc przekracza pojemnosc jakiegos typu, jego konwersja na ten typ sprawi ze zmienna przyjmie dziwna wartosc - nzywa sie to - *type overflow*

* konwersja stringa 
```kotlin
val n = 8     // Int
val d = 10.09 // Double
val c = '@'   // Char
val b = true  // Boolean

val s1 = n.toString() // "8"
val s2 = d.toString() // "10.09"
val s3 = c.toString() // "@"
val s4 = b.toString() // "true"

// mozna dokonywac tez konwersji w druga strone ze stirnga, w przpyadku booleana konwersja jest taka ze jak wartosc boola nie jest true / TRUE to wtedy jest false, ponizej przyklady

val b1 = "false".toBoolean() // false
val b2 = "tru".toBoolean()   // false
val b3 = "true".toBoolean()  // true
val b4 = "TRUE".toBoolean()  // true
```

### Konwersja na Short / Byte

* Wg. kotlina powinnismy trzymac numery zawsze w incie, a jak juz chcemy byte / short z double / float to musimy dokonywac step konwersji (ja to tak nazywam)

```kotlin
val floatNumber = 10f
val doubleNumber = 1.0

val shortNumber = floatNumber.toShort() // avoid this
val byteNumber = doubleNumber.toByte()  // avoid this

val shortNumber = floatNumber.toInt().toShort() // correct way
val byteNumber = doubleNumber.toInt().toByte()  // correct way
```

### String 

* konkatenacja stringow jest jak w js, mozna konkatenowac rozne typy i one zostana przekonwertowane na stringa, **condition jest taki ze string musi byc pierwszy w kolejnosci** *wyjatek jest gdy pierwsza konkatenowana wartoscia jest character, wtedy normalnie nastepuje konkatenacja*

* istnieje taki koncept jak **raw strings**, to cos takiego jak tempalte literal w js, deklarujemy go poprzez """ jakis string sobie tam """
```kotlin
val largeString = """
    This is the house that Jack built.
      
    This is the malt that lay in the house that Jack built.
       
    This is the rat that ate the malt
    That lay in the house that Jack built.
       
    This is the cat
    That killed the rat that ate the malt
    That lay in the house that Jack built.
""".trimIndent() // removes the first and the last lines and trims indents
print(largeString)
```

* string templates -- mozemy odwolywac sie do zmiennych w stringu poprzez $zmienna, jesli chcemy uzyc **template dla expression** to musimy odwolac sie jak w js, ${zmiennaString.length}

### Boolean

* w kotlinie 0 i >1 to nie jest false / true

* kazdy bool jest true w momencie gdy podany do niego string ma wartosc "true" lub "TRUE" "TrUE" "TRue" etc. czyli po prostu musi byc true nie zwazajac na wielkosc liter


#### Logical operators

```kotlin

//LOGICAL NOT

val f = false // f is false
val t = !f    // t is true

//LOGICAL AND

val b1 = false && false // false
val b2 = false && true // false
val b3 = true && false // false
val b4 = true && true  // true

//LOGICAL OR

val b1 = false || false // false
val b2 = false || true  // true
val b3 = true || false  // true
val b4 = true || true   // true

//LOGICAL XOR

val b1 = false xor false // false
val b2 = false xor true  // true
val b3 = true xor false  // true
val b4 = true xor true   // false
```

* kolejnosc logicznych operatorow - !, xor, &&, ||

```kotlin
val bool = true && !false // true because !false is evaluated first
```

### Type coercion

* wynik kalkulacji byte + byte = int

* wynik kalkulacji short + short = int

* wynik kalkulacji short + byte = int

* wynik kalkulacji long + int = long

* wynik klakulacji long + float = float

* wynik kalklacji long + double =  double


### Unsigned integers

* w kotlinie mamy unsigned integery, takie jak UByte, UShort, UInt, ULong

* tworzymy je poprzez dodanie do nich litery u / U na koncu wartosci lub uL / UL (dla unsigned longa)

* type overflow - juz wiem jak dziala, gdy zmienna przekroczy swoj zakres, jej wartosc binarna zostaje nadpisana, w przypadku konwersji po prostu liczba binarna jest ucinana w zakresie bitow, np przy konwersji z logna na inta, to wartosc liczbowa ktora jest w 64 bitach "zostaje przepolowiona" i tylko 32 bity od prawej zostaja przypisane do zmiennej xD



#type overflow #uint #type coercion #data types #int #byte #short #long #boolean #string #char #raw string #kotlin




