# Golang

![Golang Image](golang.png)

---------------------------------------------------------------------

# Tries

Trie is a tree type data structure which is used to store words. Like Trees, it has also nodes, root, children etc. One of the well known application of trie data structure is **Autocomplete**. For instance, as we type in a search words in search bar, Google tries to guess what we are trying to type. This application and its algorithm depends on Trie data structure.

* Storing words in Trie data structure which let's us to search words very fast.

* For a standard English alphabet, we have 26 letters. Therefore, in Trie data structure we have 26 children for each node as well. 

![Trie Image](https://upload.wikimedia.org/wikipedia/commons/b/be/Trie_example.svg)

* The image is taken from **Wikipedia.org**

--------------------------------------------------------------

## Code Implementation

* There are two structs in implementation: Node and Trie. Node is simply an element in Trie structure and Trie refers to whole structure.

* There are two ADT operations which are implemented: **Insert()** and **Search()**.

* The implementations can be seen in the code **Tries.go** in Tries directory.

---------------------------------------------------------

## Code sample  and usage

* The insertion and search operations can be seen below.

```go
    countryTrie := InitTrie()

    toAdd := []string{
        "england",
        "france",
        "germany",
        "netherlands",
        "sweden",
        "norway",
        "spain",
        "denmark",
        "turkey",
        "greece",
    }

    for _, v := range toAdd{
        countryTrie.Insert(v)
    }
    fmt.Println(countryTrie.Search("germany"))
```

```[terminal]
england is inserted to trie
france is inserted to trie
germany is inserted to trie
netherlands is inserted to trie
sweden is inserted to trie
norway is inserted to trie
spain is inserted to trie
denmark is inserted to trie
turkey is inserted to trie
greece is inserted to trie
Gotcha!, germany is included in the trie
true
```

* As you can see above, all of the words are inserted to the trie. For every letter, a **Node** is assigned. 

* In search method, nodes are simply checked if they are nil or not.

-----------------------------------

## Licence

[MIT](LICENCE)
