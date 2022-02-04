package error

import "log"

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func NsCheckErr(nameSpace string) {
	log.Fatalf("%v is not in the namespaces list", nameSpace)
}
