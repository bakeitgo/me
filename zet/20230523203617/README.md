# Kotlin

## Primitive types

* Kotlin sam dedukuje typ, numery defaultowo sa intem, zmiennoprezcinkowe defaultowo sa doublem, ponizej opis jak mozemy ustawic swoj typ

### Numbers

```kotlin
val zero = 0 // Int
val one = 1  // Int
val oneMillion = 1_000_000  // Int

val twoMillion = 2_000_000L           // Long because it is tagged with L
val bigNumber = 1_000_000_000_000_000 // Long, Kotlin automatically chooses it (Int is too small)
val ten: Long = 10                    // Long because the type is specified

val shortNumber: Short = 15 // Short because the type is specified
val byteNumber: Byte = 15   // Byte because the type is specified

val pi = 3.1415              // Double
val e = 2.71828f             // Float because it is tagged with f
val fraction: Float = 1.51f  // Float because the type is specified
```

* mozemy wyswietlic max / min value kazdego typu poprzez:

```kotlin
println(Int.MIN_VALUE)  // -2147483648
println(Int.MAX_VALUE)  // 2147483647
println(Long.MIN_VALUE) // -9223372036854775808
println(Long.MAX_VALUE) // 9223372036854775807
```

* jak i jego wielkosc w bitach  /bajtach 

```kotlin
println(Int.SIZE_BYTES) // 4
println(Int.SIZE_BITS)  // 32
```

### Characters

* to co pomiedzy single quote'ami, nie moze byc wiecej niz jednego znaku, size to 2 bajty

```kotlin
val lowerCaseLetter = 'a'
val upperCaseLetter = 'Q'
val number = '1'
val space = ' '
val dollar = '$'
```

### Booleans

* jak wszedzie

```kotlin
val enabled = true
val bugFound = false
```

### Strings

* jak wszedzie

```kotlin
val creditCardNumber = "1234 5678 9012 3456"
val message = "Learn Kotlin instead of Java."
```

#kotlin #strings #characters #integers #numbers #double #float #short #char #int #string #boolean #long

