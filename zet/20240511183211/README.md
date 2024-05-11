### Good practices

## Why we use DDD
* To make code readable by business
* To map one to one domain to code

## Separation of concerns
* Clean architecture helps in this, basically is hexagonal architecture extended via dependency flow (it flows in one way [outer to inner])

## Tests
* Create helpers, and abstract assertions to make code more readable. It is not bad to abstract

## DRY

* Dont repeat yourself, but leave this principle in data entities. Entity shouldnt be the same when pushed onto database and pushed onto external world.

* Expose methods, not internal implementation. Keep things private, expose when needed. Be immutable

## Domain objects

* Create API which is more "BEHAVIOR-ORIENTED" rather than CRUD


* Create multiple small functions instead of one huge elephant

---
### Note
* Teams with good performance are independent of another teams

* Note of a day:
    * Code will be used in ways i cannot anticipate, in ways it was not designed, and for longer than it was ever intended.


