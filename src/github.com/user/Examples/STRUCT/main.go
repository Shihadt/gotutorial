package main

func main() {
	shihad := person{
		fname: "shihad",
		lname: "thayyil",
		contactInfo: contactInfo{
			email:   "tshihad9@gmail.com",
			zipCode: 2,
		},
	}
	shihad.print()
	shihad.updateFName("noname")
	shihad.print()
}
