### Linked List

    * single where we keep reference to the head node

    * double where we keep reference both to head and tail nodes

    * O(1) for adding item at the beginnning of the list, O(n) for adding item to the end of the list (single linked list), for doubly linked list its constant operation to add to the end

    * Screwes up retrieving data, O(n) operation, we need to linear search it


### Queue

    * FIFO structure

    * We enqueue (add item to the end of list) O(1) operation

    * We dequeue (pop item from the beginning of the list O(1) operation

    * We maintain reference to both tail and head of the list

    * Bad for data retrieve

### Stack

    * LIFO structure

    * We maintain reference to the head of the list 

    * Everything is inserted / removed to / from the head of the list O(1) operation

### Array

    * Good for data retrieve O(1)operation if we know index

    * Good for push operation (if capacity is not extended), if there is need for capacity increase, we need to change underlying array, copy all values and update slice pointer which is O(n) operation

    * Better than linked list based structures when we need to search for data

### Ring buffer

    * An array buffer which have specified capacity, head is not 0 and tail is not len like in normal array, but head and tail starts from the middle and moves left / right with indices when adding, right / left when removing

    * something like this

        * ```
        [nil,nil,nil,nil,0,1,nil,nil,nil,nil]
        //after adding head
        [nil,nil,nil,-1,0,1,nil,nil,nil,nil]

        ```
