## map

### Instructions

Écrire une fonction `Map` qui, pour un tableau d'`int`, applique une fonction de type `func(int) bool` sur chaque éléments de ce tableau et retournes un tableau de toutes les valeurs de retour.

### Fonction attendue

```go
func Map(f func(int) bool, arr []int) []bool {
}
```

### Utilisation

Voici un éventuel [programme](TODO-LINK) pour tester votre fonction :

```go
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	arr = Map(piscine.IsPrime, arr)
	fmt.Println(arr)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
[false true true false true false]
student@ubuntu:~/piscine/test$
```