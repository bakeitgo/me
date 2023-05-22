# Java

## Enums

* wrapper na stałe
```java
public enums DaysOfTheWeek {
	SUNDAY,
	MONDAY,
	TUESDAY,
	WEDNESDAY,
	THURSDAY,
	FRIDAY,
	SATURDAY
}

// jak uzyc?

public class EnumTutorial {

	public static void main (String[] args) {

		DaysOfTheWeek day = DaysOfTheWeek.SUNDAY // uzycie enuma


		for (DaysOfTheWeek myDay : DaysOfTheWeek.values()) {
			System.out.println(myDay); // wypisze wszystkie wartosci w naszym enumie
		}
	}

}
```


* enum moze przyjmowac jeszce wiecej wartosci po chainie, o to przyklad implementacji, tj konstruktor w klasie:

```java
public enum Cereals {
	CAPATAIN(50),
	FROOT(60),
	REESES(100),
	COCOA(75);

	final int levelOfDeliciousness; // warto zrobic final, aby nikt nei byl w stanie nadpisac tej wartosci

	Cereals(int levelOfDeliciousness) {
		this.levelOfDeliciousness = levelOfDeliciousness;
	}

//uzyciew kodzie

public class EnumTutorial {

        public static void main (String[] args) {
	Cereals.FROOT.levelOfDeliciousness
}

}
```
# Related to: 

* https://www.youtube.com/watch?v=wq9SJb8VeyM&list=PLkeaG1zpPTHiMjczpmZ6ALd46VjjiQJ_8&index=4


#Java #enums

