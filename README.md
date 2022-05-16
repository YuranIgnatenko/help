# help

Step 1: go install github.com.YuranIgnatenko/help
Step 2: used package

Example:
// used test module : modulet.go
// init structure
p := fmt.Println
h := Help{}

// used functions structure
// get all functions struct 'Cat' 
p(h.Dir("/github.com/YuranIgnatenko/help/modulet.go", "Cat")) 
// return []string [Say, GetName, SetName]

// get all functions struct 'Cat' and annotation
p(h.DirVerbose("/github.com/YuranIgnatenko/help/modulet.go", "Cat"))
// return []string [Say(), GetName(), SetName(name string)]

// get all functions struct 'Dog'
p(h.Dir("/github.com/YuranIgnatenko/help/modulet.go", "Dog"))
// return []string [Say, GetName, SetName]

// get all functions struct 'Dog' and annotation
p(h.DirVerbose("/github.com/YuranIgnatenko/help/modulet.go", "Dog"))
// return []string [Say(), GetName(), SetName(name string)]

Thank you!
