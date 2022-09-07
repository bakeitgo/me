# ReactJS

* We can use props (properties) to pass static values to the components

* We can use state to pass dynamic values (changing values) to the components

* Component is a re-usable element which is made from molecules which are made from atoms (html elements)

* Molecule is a couple of atoms

* Atom is a single html element

* In react there are 2 kinds of components, class and function components

* In class components we can use state by setting `this.state` in their constructors. 

	* All class components should have `super(props)` called when we are defining a constructor.To make props available in the constructor

* When we update state in component, it is updated in children components also.

* Naming convention for props representing events  is on[event] and handle[event] for methods which handle events

* If we want to share state, we need to hold it in the parent component of childs between who we want to share state.

* We should consider not changing data directly, instead make copy of data and change it. This concept is called immutability. It helps in making pure components, and helps react know when to re-render component.













#React #ReactJS #js #javascript #props #state
