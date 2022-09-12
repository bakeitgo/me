# Redux

* Redux is used to manage state of whole app

* Redux uses `reducer` function to manage app state

* `reducer` contains `state` (current) and `action` 

	* `action` is a plain js object which takes `type` field (string with descriptive name *domain/eventName*) and `payload` field (additional info about action). to simplify action, we use `action creator` *func which returns object with specified fields/values*

	* They make immutable updates (not mutating state directly, make copy and work on it)

	* They must not do any `async` operations, calculate `random values` or cause other `side effects`


* Redux app state lives in an object called *store*

* The only way to update state is to use `dispatch` function with action object passed in

	* Dispatch is like triggering event

* Selectors are used to extract specific piece from store state value.




