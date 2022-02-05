package error

import "log"

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func NameSpaceCheckErr(nameSpace string, booler bool) {
	if booler == false {
		log.Fatalf("'%v' Is Not In The Namespaces List", nameSpace)
	}
}
