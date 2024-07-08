## MaxHeap

* Every child is smaller than parent

* The parent is the MAX in the tree

## MinHeap

* Every child is larger than parent

* The parent is the MIN in the tree


## Properties which apply to BOTH

* Weak ordered structure:

    * We add elements from the left to right in a row (kinda like BFS)

* How inserting works?
    * We add value to the end of the tree (from the left), and we compare it to the previous node.
    * For e.g. in MinHeap (the min is the top one) we add value where it should be placed as the top one value (because its the smallest), we add it to the end, and we compare it to the previous node. IF prev node value is GREATER THAN our inserted value, we swap. And we do this operation until prev node value < inserted value, or the prev node do not exists (then our inserted node is the top one node). Thats O(log h) operation, where h is the height of our tree.

* How deleting works?
    * We delete node, then we take the very last node in our tree, counting from the left ofc, shift node up to the deleted node position, then we bubble each position down and do compare/swap like in insert, but in reversed direction and if there are two nodes down, we take the smallest / greatest value (based on heap) to compare/swap operation. O(log h) operation.

* Implementing it in array data structure
    * How do we maintain from parent to childs relationship?
        a) Via index:
            Left child is: 2i+1
            Right child is: 2i+2
    * How do we maintain from child to parent?
        a) Via index, in the opposite:
            Parent is: (i-1)/2


